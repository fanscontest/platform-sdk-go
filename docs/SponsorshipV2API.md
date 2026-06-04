# \SponsorshipV2API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByContestIdSponsorships**](SponsorshipV2API.md#CreateContestsByContestIdSponsorships) | **Post** /v2/contests/{contestId}/sponsorships | Create Sponsorship Offer (v2)
[**CreateSponsorshipsByIdAccept**](SponsorshipV2API.md#CreateSponsorshipsByIdAccept) | **Post** /v2/sponsorships/{id}/accept | Accept Sponsorship (v2)
[**CreateSponsorshipsByIdReject**](SponsorshipV2API.md#CreateSponsorshipsByIdReject) | **Post** /v2/sponsorships/{id}/reject | Reject Sponsorship (v2)
[**GetContestsByContestIdSponsorships**](SponsorshipV2API.md#GetContestsByContestIdSponsorships) | **Get** /v2/contests/{contestId}/sponsorships | Get Contest Sponsorships (v2)



## CreateContestsByContestIdSponsorships

> DomainSponsorship CreateContestsByContestIdSponsorships(ctx, contestId).RequestCreateSponsorshipOfferRequest(requestCreateSponsorshipOfferRequest).Execute()

Create Sponsorship Offer (v2)



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
	contestId := "contestId_example" // string | Contest ID
	requestCreateSponsorshipOfferRequest := *openapiclient.NewRequestCreateSponsorshipOfferRequest("AmountCurrency_example", "AmountValue_example", "ContestId_example", []string{"TermIds_example"}) // RequestCreateSponsorshipOfferRequest | Create Sponsorship Offer Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipV2API.CreateContestsByContestIdSponsorships(context.Background(), contestId).RequestCreateSponsorshipOfferRequest(requestCreateSponsorshipOfferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipV2API.CreateContestsByContestIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdSponsorships`: DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipV2API.CreateContestsByContestIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByContestIdSponsorshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestCreateSponsorshipOfferRequest** | [**RequestCreateSponsorshipOfferRequest**](RequestCreateSponsorshipOfferRequest.md) | Create Sponsorship Offer Request | 

### Return type

[**DomainSponsorship**](DomainSponsorship.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipsByIdAccept

> DomainSponsorship CreateSponsorshipsByIdAccept(ctx, id).Execute()

Accept Sponsorship (v2)



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
	id := "id_example" // string | Sponsorship ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipV2API.CreateSponsorshipsByIdAccept(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipV2API.CreateSponsorshipsByIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipsByIdAccept`: DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipV2API.CreateSponsorshipsByIdAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsorship ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipsByIdAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorship**](DomainSponsorship.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipsByIdReject

> DomainSponsorship CreateSponsorshipsByIdReject(ctx, id).Execute()

Reject Sponsorship (v2)



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
	id := "id_example" // string | Sponsorship ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipV2API.CreateSponsorshipsByIdReject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipV2API.CreateSponsorshipsByIdReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipsByIdReject`: DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipV2API.CreateSponsorshipsByIdReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsorship ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipsByIdRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorship**](DomainSponsorship.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdSponsorships

> []DomainSponsorship GetContestsByContestIdSponsorships(ctx, contestId).Execute()

Get Contest Sponsorships (v2)



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
	contestId := "contestId_example" // string | Contest ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipV2API.GetContestsByContestIdSponsorships(context.Background(), contestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipV2API.GetContestsByContestIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdSponsorships`: []DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipV2API.GetContestsByContestIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByContestIdSponsorshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainSponsorship**](DomainSponsorship.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

