/*
FansContest outbound-webhook verification (platform ADR 0011).

This is a HAND-WRITTEN helper injected into the generated SDK by
tools/sdk-gen/generate.sh — it is NOT emitted by openapi-generator (webhook
verification is a client-side helper, not a REST operation). It is the single,
canonical Go definition of the scheme, so a tenant calls one function instead of
copy-pasting the crypto. The TypeScript SDK ships the byte-for-byte equivalent
(constructEvent). It mirrors the platform's Util/webhooksig exactly.

Wire format (Stripe-style): every delivery carries

	X-FC-Signature: t=<unix-seconds>,v1=<hex HMAC-SHA256 of "<t>.<body>">

keyed by the subscription's signing secret (whsec_…). The HMAC is over the exact
raw request bytes; the receiver also rejects deliveries whose timestamp falls
outside a tolerance window (replay guard). Changing anything here is a
cross-service contract change (it must stay in step with the platform's
Util/webhooksig and the TS SDK).
*/
package platform

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

// Webhook header names. WebhookEventTypeHeader is a non-authenticated convenience
// so a receiver can route or log without parsing the body first; the body's
// "type" field is authoritative.
const (
	WebhookSignatureHeader = "X-FC-Signature"
	WebhookEventTypeHeader = "X-FC-Event-Type"
)

// DefaultWebhookTolerance is the replay-guard window used by ConstructWebhookEvent
// and by VerifyWebhookSignature when tolerance <= 0 is passed.
const DefaultWebhookTolerance = 5 * time.Minute

// Verification errors. Distinct so a receiver can tell "this isn't from us"
// (mismatch) from "this is stale" (tolerance).
var (
	ErrWebhookMissingSecret        = errors.New("platform: webhook signing secret is empty")
	ErrWebhookMalformedSignature   = errors.New("platform: malformed webhook signature header")
	ErrWebhookSignatureMismatch    = errors.New("platform: webhook signature mismatch")
	ErrWebhookTimestampOutOfWindow = errors.New("platform: webhook timestamp outside tolerance")
	ErrWebhookInvalidPayload       = errors.New("platform: webhook payload is not a valid event envelope")
)

// WebhookEvent is the delivered event envelope (the platform's Domain.WebhookEvent).
// Data is event-type-specific raw JSON; unmarshal it based on Type.
type WebhookEvent struct {
	Id        string          `json:"id"`
	Type      string          `json:"type"`
	TenantId  string          `json:"tenant_id"`
	CreatedAt time.Time       `json:"created_at"`
	Data      json.RawMessage `json:"data"`
}

// ConstructWebhookEvent verifies a delivery's signature AND parses it into a
// typed WebhookEvent — the one call a tenant needs (mirrors Stripe's
// webhook.ConstructEvent). body MUST be the exact raw request bytes. A tolerance
// <= 0 uses DefaultWebhookTolerance; pass a negative sentinel only if you mean
// the default. Returns the event on success, or a verification error.
func ConstructWebhookEvent(secret, signatureHeader string, body []byte, tolerance time.Duration, now time.Time) (*WebhookEvent, error) {
	if tolerance <= 0 {
		tolerance = DefaultWebhookTolerance
	}
	if err := VerifyWebhookSignature(secret, signatureHeader, body, tolerance, now); err != nil {
		return nil, err
	}
	var evt WebhookEvent
	if err := json.Unmarshal(body, &evt); err != nil || evt.Type == "" {
		return nil, ErrWebhookInvalidPayload
	}
	return &evt, nil
}

// VerifyWebhookSignature checks an X-FC-Signature header value against body using
// secret. When tolerance > 0, the header timestamp must be within tolerance of
// now. Comparison is constant-time (hmac.Equal). Returns nil on success. Use when
// you handle the body yourself; most callers want ConstructWebhookEvent.
func VerifyWebhookSignature(secret, signatureHeader string, body []byte, tolerance time.Duration, now time.Time) error {
	if secret == "" {
		return ErrWebhookMissingSecret
	}

	ts, sig, err := parseWebhookSignature(signatureHeader)
	if err != nil {
		return err
	}

	if tolerance > 0 {
		delta := now.Unix() - ts
		if delta < 0 {
			delta = -delta
		}
		if time.Duration(delta)*time.Second > tolerance {
			return ErrWebhookTimestampOutOfWindow
		}
	}

	expected := signWebhook(secret, ts, body)
	if !hmac.Equal([]byte(expected), []byte(sig)) {
		return ErrWebhookSignatureMismatch
	}
	return nil
}

// signWebhook returns the hex-encoded HMAC-SHA256 of "<ts>.<body>" keyed by secret.
func signWebhook(secret string, ts int64, body []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(ts, 10)))
	mac.Write([]byte("."))
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

// parseWebhookSignature splits "t=<ts>,v1=<hex>" into its parts. Order-independent;
// unknown fields are ignored so the scheme can grow new versions (v2, …) without
// breaking older parsers.
func parseWebhookSignature(header string) (ts int64, sig string, err error) {
	var tPart, vPart string
	for _, part := range strings.Split(header, ",") {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "t":
			tPart = kv[1]
		case "v1":
			vPart = kv[1]
		}
	}
	if tPart == "" || vPart == "" {
		return 0, "", ErrWebhookMalformedSignature
	}
	ts, err = strconv.ParseInt(tPart, 10, 64)
	if err != nil {
		return 0, "", ErrWebhookMalformedSignature
	}
	return ts, vPart, nil
}
