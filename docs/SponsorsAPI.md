# \SponsorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsors**](SponsorsAPI.md#CreateSponsors) | **Post** /v2/sponsors | Create Sponsor (v2)
[**GetSponsors**](SponsorsAPI.md#GetSponsors) | **Get** /v2/sponsors | Get Sponsors (v2)
[**GetSponsorsById**](SponsorsAPI.md#GetSponsorsById) | **Get** /v2/sponsors/{id} | Get Sponsor by ID (v2)
[**GetSponsorsByIdSponsorships**](SponsorsAPI.md#GetSponsorsByIdSponsorships) | **Get** /v2/sponsors/{id}/sponsorships | Get Sponsor Sponsorships (v2)



## CreateSponsors

> DomainSponsorResponse CreateSponsors(ctx).RequestCreateSponsorRequest(requestCreateSponsorRequest).XTenantUserId(xTenantUserId).Execute()

Create Sponsor (v2)



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
	requestCreateSponsorRequest := *openapiclient.NewRequestCreateSponsorRequest("Name_example", "Url_example") // RequestCreateSponsorRequest | Create Sponsor Request
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.CreateSponsors(context.Background()).RequestCreateSponsorRequest(requestCreateSponsorRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.CreateSponsors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsors`: DomainSponsorResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.CreateSponsors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorRequest** | [**RequestCreateSponsorRequest**](RequestCreateSponsorRequest.md) | Create Sponsor Request | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorResponse**](DomainSponsorResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsors

> DomainSponsorListResponse GetSponsors(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Sponsors (v2)



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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.GetSponsors(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsors`: DomainSponsorListResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorListResponse**](DomainSponsorListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorsById

> DomainSponsorResponse GetSponsorsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get Sponsor by ID (v2)



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
	id := "id_example" // string | Sponsor ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.GetSponsorsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorsById`: DomainSponsorResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorResponse**](DomainSponsorResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorsByIdSponsorships

> DomainSponsorshipListResponse GetSponsorsByIdSponsorships(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Sponsor Sponsorships (v2)



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
	id := "id_example" // string | Sponsor ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.GetSponsorsByIdSponsorships(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsorsByIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorsByIdSponsorships`: DomainSponsorshipListResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsorsByIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsByIdSponsorshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorshipListResponse**](DomainSponsorshipListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

