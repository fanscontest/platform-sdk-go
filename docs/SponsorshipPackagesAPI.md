# \SponsorshipPackagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsorshipPackages**](SponsorshipPackagesAPI.md#CreateSponsorshipPackages) | **Post** /v2/sponsorship-packages | Create Sponsorship Package (v2)
[**CreateSponsorshipPackagesByIdCancel**](SponsorshipPackagesAPI.md#CreateSponsorshipPackagesByIdCancel) | **Post** /v2/sponsorship-packages/{id}/cancel | Cancel Sponsorship Package (v2)
[**GetSponsorshipPackages**](SponsorshipPackagesAPI.md#GetSponsorshipPackages) | **Get** /v2/sponsorship-packages | Get Sponsorship Packages (v2)
[**GetSponsorshipPackagesAvailable**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesAvailable) | **Get** /v2/sponsorship-packages/available | Get Available Packages (v2)
[**GetSponsorshipPackagesByIdAgreement**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesByIdAgreement) | **Get** /v2/sponsorship-packages/{id}/agreement | Get Package Agreement (v2)
[**GetSponsorshipPackagesByIdSponsorships**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesByIdSponsorships) | **Get** /v2/sponsorship-packages/{id}/sponsorships | Get Package Sponsorships (v2)



## CreateSponsorshipPackages

> CreateSponsorshipPackages201Response CreateSponsorshipPackages(ctx).RequestCreateSponsorshipPackageRequest(requestCreateSponsorshipPackageRequest).XTenantUserId(xTenantUserId).Execute()

Create Sponsorship Package (v2)



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
	requestCreateSponsorshipPackageRequest := *openapiclient.NewRequestCreateSponsorshipPackageRequest("BudgetPerContestCurrency_example", "BudgetPerContestValue_example", []string{"TermIds_example"}, "TotalBudgetCurrency_example", "TotalBudgetValue_example", int32(123)) // RequestCreateSponsorshipPackageRequest | Create Sponsorship Package Request
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.CreateSponsorshipPackages(context.Background()).RequestCreateSponsorshipPackageRequest(requestCreateSponsorshipPackageRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.CreateSponsorshipPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipPackages`: CreateSponsorshipPackages201Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.CreateSponsorshipPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorshipPackageRequest** | [**RequestCreateSponsorshipPackageRequest**](RequestCreateSponsorshipPackageRequest.md) | Create Sponsorship Package Request | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateSponsorshipPackages201Response**](CreateSponsorshipPackages201Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipPackagesByIdCancel

> CreateSponsorshipPackages201Response CreateSponsorshipPackagesByIdCancel(ctx, id).XTenantUserId(xTenantUserId).Execute()

Cancel Sponsorship Package (v2)



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
	id := "id_example" // string | Package ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipPackagesByIdCancel`: CreateSponsorshipPackages201Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipPackagesByIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateSponsorshipPackages201Response**](CreateSponsorshipPackages201Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackages

> GetSponsorshipPackages200Response GetSponsorshipPackages(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Sponsorship Packages (v2)



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page endpoint; see uman#132) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page endpoint; see uman#132) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackages(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackages`: GetSponsorshipPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page endpoint; see uman#132) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page endpoint; see uman#132) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetSponsorshipPackages200Response**](GetSponsorshipPackages200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesAvailable

> GetSponsorshipPackages200Response GetSponsorshipPackagesAvailable(ctx).SponsorId(sponsorId).PackageId(packageId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Available Packages (v2)



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
	sponsorId := "sponsorId_example" // string | Filter by sponsor ID (optional)
	packageId := "packageId_example" // string | Get specific package by ID (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable(context.Background()).SponsorId(sponsorId).PackageId(packageId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesAvailable`: GetSponsorshipPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sponsorId** | **string** | Filter by sponsor ID | 
 **packageId** | **string** | Get specific package by ID | 
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetSponsorshipPackages200Response**](GetSponsorshipPackages200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesByIdAgreement

> GetSponsorshipPackagesByIdAgreement200Response GetSponsorshipPackagesByIdAgreement(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get Package Agreement (v2)



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
	id := "id_example" // string | Package ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesByIdAgreement`: GetSponsorshipPackagesByIdAgreement200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesByIdAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetSponsorshipPackagesByIdAgreement200Response**](GetSponsorshipPackagesByIdAgreement200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesByIdSponsorships

> GetContestsByContestIdSponsorships200Response GetSponsorshipPackagesByIdSponsorships(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Package Sponsorships (v2)



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
	id := "id_example" // string | Package ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesByIdSponsorships`: GetContestsByContestIdSponsorships200Response
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesByIdSponsorshipsRequest struct via the builder pattern


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

