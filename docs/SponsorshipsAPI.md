# \SponsorshipsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByContestIdSponsorships**](SponsorshipsAPI.md#CreateContestsByContestIdSponsorships) | **Post** /v2/contests/{contestId}/sponsorships | Create Sponsorship Offer (v2)
[**CreateSponsorshipsByIdAccept**](SponsorshipsAPI.md#CreateSponsorshipsByIdAccept) | **Post** /v2/sponsorships/{id}/accept | Accept Sponsorship (v2)
[**CreateSponsorshipsByIdReject**](SponsorshipsAPI.md#CreateSponsorshipsByIdReject) | **Post** /v2/sponsorships/{id}/reject | Reject Sponsorship (v2)
[**GetContestsByContestIdSponsorships**](SponsorshipsAPI.md#GetContestsByContestIdSponsorships) | **Get** /v2/contests/{contestId}/sponsorships | Get Contest Sponsorships (v2)
[**GetSponsorshipsByIdAgreement**](SponsorshipsAPI.md#GetSponsorshipsByIdAgreement) | **Get** /v2/sponsorships/{id}/agreement | Get Sponsorship Agreement (v2)



## CreateContestsByContestIdSponsorships

> CreateContestsByContestIdSponsorships201Response CreateContestsByContestIdSponsorships(ctx, contestId).RequestCreateSponsorshipOfferRequest(requestCreateSponsorshipOfferRequest).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipsAPI.CreateContestsByContestIdSponsorships(context.Background(), contestId).RequestCreateSponsorshipOfferRequest(requestCreateSponsorshipOfferRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipsAPI.CreateContestsByContestIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdSponsorships`: CreateContestsByContestIdSponsorships201Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipsAPI.CreateContestsByContestIdSponsorships`: %v\n", resp)
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
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateContestsByContestIdSponsorships201Response**](CreateContestsByContestIdSponsorships201Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipsByIdAccept

> CreateContestsByContestIdSponsorships201Response CreateSponsorshipsByIdAccept(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipsAPI.CreateSponsorshipsByIdAccept(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipsAPI.CreateSponsorshipsByIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipsByIdAccept`: CreateContestsByContestIdSponsorships201Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipsAPI.CreateSponsorshipsByIdAccept`: %v\n", resp)
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateContestsByContestIdSponsorships201Response**](CreateContestsByContestIdSponsorships201Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipsByIdReject

> CreateContestsByContestIdSponsorships201Response CreateSponsorshipsByIdReject(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipsAPI.CreateSponsorshipsByIdReject(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipsAPI.CreateSponsorshipsByIdReject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipsByIdReject`: CreateContestsByContestIdSponsorships201Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipsAPI.CreateSponsorshipsByIdReject`: %v\n", resp)
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateContestsByContestIdSponsorships201Response**](CreateContestsByContestIdSponsorships201Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdSponsorships

> GetContestsByContestIdSponsorships200Response GetContestsByContestIdSponsorships(ctx, contestId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipsAPI.GetContestsByContestIdSponsorships(context.Background(), contestId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipsAPI.GetContestsByContestIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdSponsorships`: GetContestsByContestIdSponsorships200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipsAPI.GetContestsByContestIdSponsorships`: %v\n", resp)
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

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByContestIdSponsorships200Response**](GetContestsByContestIdSponsorships200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipsByIdAgreement

> GetSponsorshipsByIdAgreement200Response GetSponsorshipsByIdAgreement(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get Sponsorship Agreement (v2)



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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipsAPI.GetSponsorshipsByIdAgreement(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipsAPI.GetSponsorshipsByIdAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipsByIdAgreement`: GetSponsorshipsByIdAgreement200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipsAPI.GetSponsorshipsByIdAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsorship ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipsByIdAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetSponsorshipsByIdAgreement200Response**](GetSponsorshipsByIdAgreement200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

