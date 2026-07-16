# \CreativeContestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdBallots**](CreativeContestsAPI.md#CreateContestsByIdBallots) | **Post** /v2/contests/{id}/ballots | Serve a community-judging matchup
[**CreateContestsByIdEntries**](CreativeContestsAPI.md#CreateContestsByIdEntries) | **Post** /v2/contests/{id}/entries | Submit entry for creative contest
[**CreateContestsByIdEntriesByEntryIdScore**](CreativeContestsAPI.md#CreateContestsByIdEntriesByEntryIdScore) | **Post** /v2/contests/{id}/entries/{entryId}/score | Score an entry
[**CreateContestsByIdForceProceed**](CreativeContestsAPI.md#CreateContestsByIdForceProceed) | **Post** /v2/contests/{id}/force-proceed | Force proceed from paused state
[**CreateContestsByIdJudgments**](CreativeContestsAPI.md#CreateContestsByIdJudgments) | **Post** /v2/contests/{id}/judgments | Submit a community-judging pairwise choice
[**CreateContestsByIdJurorsInvite**](CreativeContestsAPI.md#CreateContestsByIdJurorsInvite) | **Post** /v2/contests/{id}/jurors/invite | Invite jurors to judge a contest
[**CreateJurorsInvitationsByInvitationIdAccept**](CreativeContestsAPI.md#CreateJurorsInvitationsByInvitationIdAccept) | **Post** /v2/jurors/invitations/{invitationId}/accept | Accept a juror invitation
[**CreateJurorsInvitationsByInvitationIdDecline**](CreativeContestsAPI.md#CreateJurorsInvitationsByInvitationIdDecline) | **Post** /v2/jurors/invitations/{invitationId}/decline | Decline a juror invitation
[**CreateJurorsOptIn**](CreativeContestsAPI.md#CreateJurorsOptIn) | **Post** /v2/jurors/opt-in | Opt into jury pool
[**DeleteContestsByIdEntries**](CreativeContestsAPI.md#DeleteContestsByIdEntries) | **Delete** /v2/contests/{id}/entries | Withdraw entry from creative contest
[**GetContestsByIdEntries**](CreativeContestsAPI.md#GetContestsByIdEntries) | **Get** /v2/contests/{id}/entries | Get entries for a contest
[**GetContestsByIdJurors**](CreativeContestsAPI.md#GetContestsByIdJurors) | **Get** /v2/contests/{id}/jurors | Get jurors for a contest
[**GetContestsByIdJurorsInvitations**](CreativeContestsAPI.md#GetContestsByIdJurorsInvitations) | **Get** /v2/contests/{id}/jurors/invitations | Get invitations for a contest
[**GetJurorsAssignments**](CreativeContestsAPI.md#GetJurorsAssignments) | **Get** /v2/jurors/assignments | Get juror assignments (cursor-paginated)
[**GetJurorsInvitations**](CreativeContestsAPI.md#GetJurorsInvitations) | **Get** /v2/jurors/invitations | Get my pending invitations
[**UpdateContestsByIdEntriesByEntryIdScore**](CreativeContestsAPI.md#UpdateContestsByIdEntriesByEntryIdScore) | **Put** /v2/contests/{id}/entries/{entryId}/score | Score an entry



## CreateContestsByIdBallots

> HandlerMatchupResponseResponse CreateContestsByIdBallots(ctx, id).XActingAs(xActingAs).Execute()

Serve a community-judging matchup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdBallots(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdBallots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdBallots`: HandlerMatchupResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdBallots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdBallotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerMatchupResponseResponse**](HandlerMatchupResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdEntries

> DomainEntryResponse CreateContestsByIdEntries(ctx, id).RequestSubmitEntryRequest(requestSubmitEntryRequest).XActingAs(xActingAs).Execute()

Submit entry for creative contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	requestSubmitEntryRequest := *openapiclient.NewRequestSubmitEntryRequest("MediaType_example", "MediaUrl_example") // RequestSubmitEntryRequest | Entry submission
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdEntries(context.Background(), id).RequestSubmitEntryRequest(requestSubmitEntryRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntries`: DomainEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestSubmitEntryRequest** | [**RequestSubmitEntryRequest**](RequestSubmitEntryRequest.md) | Entry submission | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainEntryResponse**](DomainEntryResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdEntriesByEntryIdScore

> DomainJurorScoreResponse CreateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XActingAs(xActingAs).Execute()

Score an entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	entryId := "entryId_example" // string | Entry ID
	requestScoreEntryRequest := *openapiclient.NewRequestScoreEntryRequest(int32(123)) // RequestScoreEntryRequest | Score
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntriesByEntryIdScore`: DomainJurorScoreResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdEntriesByEntryIdScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestScoreEntryRequest** | [**RequestScoreEntryRequest**](RequestScoreEntryRequest.md) | Score | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorScoreResponse**](DomainJurorScoreResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdForceProceed

> HandlerStatusResponseResponse CreateContestsByIdForceProceed(ctx, id).XActingAs(xActingAs).Execute()

Force proceed from paused state



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdForceProceed(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdForceProceed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdForceProceed`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdForceProceed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdForceProceedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdJudgments

> DomainCommunityBallotResponse CreateContestsByIdJudgments(ctx, id).RequestSubmitJudgmentRequest(requestSubmitJudgmentRequest).XActingAs(xActingAs).Execute()

Submit a community-judging pairwise choice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	requestSubmitJudgmentRequest := *openapiclient.NewRequestSubmitJudgmentRequest("BallotToken_example", "Choice_example") // RequestSubmitJudgmentRequest | Ballot token + choice
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdJudgments(context.Background(), id).RequestSubmitJudgmentRequest(requestSubmitJudgmentRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdJudgments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdJudgments`: DomainCommunityBallotResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdJudgments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdJudgmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestSubmitJudgmentRequest** | [**RequestSubmitJudgmentRequest**](RequestSubmitJudgmentRequest.md) | Ballot token + choice | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainCommunityBallotResponse**](DomainCommunityBallotResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdJurorsInvite

> HandlerStatusResponseResponse CreateContestsByIdJurorsInvite(ctx, id).RequestInviteJurorsRequest(requestInviteJurorsRequest).XActingAs(xActingAs).Execute()

Invite jurors to judge a contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	requestInviteJurorsRequest := *openapiclient.NewRequestInviteJurorsRequest([]string{"PlatformIdentityIds_example"}) // RequestInviteJurorsRequest | Member IDs to invite
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdJurorsInvite(context.Background(), id).RequestInviteJurorsRequest(requestInviteJurorsRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdJurorsInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdJurorsInvite`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdJurorsInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdJurorsInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInviteJurorsRequest** | [**RequestInviteJurorsRequest**](RequestInviteJurorsRequest.md) | Member IDs to invite | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsInvitationsByInvitationIdAccept

> DomainJurorAssignmentResponse CreateJurorsInvitationsByInvitationIdAccept(ctx, invitationId).XActingAs(xActingAs).Execute()

Accept a juror invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	invitationId := "invitationId_example" // string | Invitation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept(context.Background(), invitationId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdAccept`: DomainJurorAssignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsInvitationsByInvitationIdAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorAssignmentResponse**](DomainJurorAssignmentResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsInvitationsByInvitationIdDecline

> HandlerStatusResponseResponse CreateJurorsInvitationsByInvitationIdDecline(ctx, invitationId).XActingAs(xActingAs).Execute()

Decline a juror invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	invitationId := "invitationId_example" // string | Invitation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline(context.Background(), invitationId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdDecline`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsInvitationsByInvitationIdDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsOptIn

> DomainJurorResponse CreateJurorsOptIn(ctx).RequestOptInJurorRequest(requestOptInJurorRequest).XActingAs(xActingAs).Execute()

Opt into jury pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	requestOptInJurorRequest := *openapiclient.NewRequestOptInJurorRequest([]string{"Categories_example"}) // RequestOptInJurorRequest | Categories
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsOptIn(context.Background()).RequestOptInJurorRequest(requestOptInJurorRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsOptIn`: DomainJurorResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsOptIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestOptInJurorRequest** | [**RequestOptInJurorRequest**](RequestOptInJurorRequest.md) | Categories | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorResponse**](DomainJurorResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsByIdEntries

> HandlerStatusResponseResponse DeleteContestsByIdEntries(ctx, id).XActingAs(xActingAs).Execute()

Withdraw entry from creative contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.DeleteContestsByIdEntries(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.DeleteContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContestsByIdEntries`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.DeleteContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdEntries

> DomainEntryListResponse GetContestsByIdEntries(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get entries for a contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdEntries(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdEntries`: DomainEntryListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainEntryListResponse**](DomainEntryListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurors

> DomainJurorListResponse GetContestsByIdJurors(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get jurors for a contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, derived) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, derived) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdJurors(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdJurors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurors`: DomainJurorListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdJurors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdJurorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, derived) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, derived) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorListResponse**](DomainJurorListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurorsInvitations

> DomainJurorAssignmentListResponse GetContestsByIdJurorsInvitations(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get invitations for a contest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdJurorsInvitations(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurorsInvitations`: DomainJurorAssignmentListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdJurorsInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdJurorsInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorAssignmentListResponse**](DomainJurorAssignmentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsAssignments

> DomainJurorAssignmentListResponse GetJurorsAssignments(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get juror assignments (cursor-paginated)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetJurorsAssignments(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetJurorsAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsAssignments`: DomainJurorAssignmentListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetJurorsAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorAssignmentListResponse**](DomainJurorAssignmentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsInvitations

> DomainJurorAssignmentListResponse GetJurorsInvitations(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get my pending invitations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, service-assembled) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, service-assembled) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetJurorsInvitations(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsInvitations`: DomainJurorAssignmentListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetJurorsInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, service-assembled) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, service-assembled) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorAssignmentListResponse**](DomainJurorAssignmentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdEntriesByEntryIdScore

> DomainJurorScoreResponse UpdateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XActingAs(xActingAs).Execute()

Score an entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	id := "id_example" // string | Contest ID
	entryId := "entryId_example" // string | Entry ID
	requestScoreEntryRequest := *openapiclient.NewRequestScoreEntryRequest(int32(123)) // RequestScoreEntryRequest | Score
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdEntriesByEntryIdScore`: DomainJurorScoreResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContestsByIdEntriesByEntryIdScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestScoreEntryRequest** | [**RequestScoreEntryRequest**](RequestScoreEntryRequest.md) | Score | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainJurorScoreResponse**](DomainJurorScoreResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

