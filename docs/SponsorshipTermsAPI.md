# \SponsorshipTermsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsorshipTerms**](SponsorshipTermsAPI.md#CreateSponsorshipTerms) | **Post** /v2/sponsorship-terms | Create Sponsorship Term (v2)
[**DeleteSponsorshipTermsById**](SponsorshipTermsAPI.md#DeleteSponsorshipTermsById) | **Delete** /v2/sponsorship-terms/{id} | Deactivate Sponsorship Term (v2)
[**GetSponsorshipTerms**](SponsorshipTermsAPI.md#GetSponsorshipTerms) | **Get** /v2/sponsorship-terms | Get Sponsorship Terms (v2)
[**GetSponsorshipTermsById**](SponsorshipTermsAPI.md#GetSponsorshipTermsById) | **Get** /v2/sponsorship-terms/{id} | Get Sponsorship Term (v2)
[**GetSponsorshipTermsPlatform**](SponsorshipTermsAPI.md#GetSponsorshipTermsPlatform) | **Get** /v2/sponsorship-terms/platform | Get Platform Terms (v2)
[**UpdateSponsorshipTermsById**](SponsorshipTermsAPI.md#UpdateSponsorshipTermsById) | **Put** /v2/sponsorship-terms/{id} | Update Sponsorship Term (v2)



## CreateSponsorshipTerms

> DomainSponsorshipTerm CreateSponsorshipTerms(ctx).RequestCreateSponsorshipTermRequest(requestCreateSponsorshipTermRequest).XTenantUserId(xTenantUserId).Execute()

Create Sponsorship Term (v2)



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
	requestCreateSponsorshipTermRequest := *openapiclient.NewRequestCreateSponsorshipTermRequest("Description_example", "Title_example") // RequestCreateSponsorshipTermRequest | Create Term Request
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermsAPI.CreateSponsorshipTerms(context.Background()).RequestCreateSponsorshipTermRequest(requestCreateSponsorshipTermRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.CreateSponsorshipTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipTerms`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.CreateSponsorshipTerms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorshipTermRequest** | [**RequestCreateSponsorshipTermRequest**](RequestCreateSponsorshipTermRequest.md) | Create Term Request | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSponsorshipTermsById

> HandlerStatusResponse DeleteSponsorshipTermsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Deactivate Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermsAPI.DeleteSponsorshipTermsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.DeleteSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSponsorshipTermsById`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.DeleteSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSponsorshipTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipTerms

> []DomainSponsorshipTerm GetSponsorshipTerms(ctx).XTenantUserId(xTenantUserId).Execute()

Get Sponsorship Terms (v2)



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
	resp, r, err := apiClient.SponsorshipTermsAPI.GetSponsorshipTerms(context.Background()).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.GetSponsorshipTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTerms`: []DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.GetSponsorshipTerms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipTermsById

> DomainSponsorshipTerm GetSponsorshipTermsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermsAPI.GetSponsorshipTermsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.GetSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTermsById`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.GetSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipTermsPlatform

> []DomainSponsorshipTerm GetSponsorshipTermsPlatform(ctx).XTenantUserId(xTenantUserId).Execute()

Get Platform Terms (v2)



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
	resp, r, err := apiClient.SponsorshipTermsAPI.GetSponsorshipTermsPlatform(context.Background()).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.GetSponsorshipTermsPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTermsPlatform`: []DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.GetSponsorshipTermsPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSponsorshipTermsById

> DomainSponsorshipTerm UpdateSponsorshipTermsById(ctx, id).RequestUpdateSponsorshipTermRequest(requestUpdateSponsorshipTermRequest).XTenantUserId(xTenantUserId).Execute()

Update Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID
	requestUpdateSponsorshipTermRequest := *openapiclient.NewRequestUpdateSponsorshipTermRequest("Description_example", "Title_example", "Version_example") // RequestUpdateSponsorshipTermRequest | Update Term Request
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermsAPI.UpdateSponsorshipTermsById(context.Background(), id).RequestUpdateSponsorshipTermRequest(requestUpdateSponsorshipTermRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermsAPI.UpdateSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSponsorshipTermsById`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermsAPI.UpdateSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsorshipTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateSponsorshipTermRequest** | [**RequestUpdateSponsorshipTermRequest**](RequestUpdateSponsorshipTermRequest.md) | Update Term Request | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

