# \CountriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicCountries**](CountriesAPI.md#GetPublicCountries) | **Get** /v2/public/countries | List enabled countries (sorted, with native names for non-English Accept-Language)
[**GetPublicLocation**](CountriesAPI.md#GetPublicLocation) | **Get** /v2/public/location | Resolve the caller&#39;s country by IP (CF-Ipcountry, then Maxmind)



## GetPublicCountries

> DomainCountryListResponse GetPublicCountries(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List enabled countries (sorted, with native names for non-English Accept-Language)



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page reference data) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page reference data) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetPublicCountries(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetPublicCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicCountries`: DomainCountryListResponse
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetPublicCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page reference data) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page reference data) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainCountryListResponse**](DomainCountryListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicLocation

> DomainCountryResponse GetPublicLocation(ctx).XTenantUserId(xTenantUserId).Execute()

Resolve the caller's country by IP (CF-Ipcountry, then Maxmind)

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetPublicLocation(context.Background()).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetPublicLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicLocation`: DomainCountryResponse
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetPublicLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainCountryResponse**](DomainCountryResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

