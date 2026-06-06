# \AnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyticsImpression**](AnalyticsAPI.md#CreateAnalyticsImpression) | **Post** /v2/analytics/impression | Record a click/impression event
[**CreateTrackerClick**](AnalyticsAPI.md#CreateTrackerClick) | **Post** /v2/tracker/click | Record a click/impression event
[**GetAnalyticsOverview**](AnalyticsAPI.md#GetAnalyticsOverview) | **Get** /v2/analytics/overview | Get the caller&#39;s analytics overview



## CreateAnalyticsImpression

> HandlerStatusResponseResponse CreateAnalyticsImpression(ctx).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).XTenantUserId(xTenantUserId).Execute()

Record a click/impression event



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
	requestUpdateClickCountRequest := *openapiclient.NewRequestUpdateClickCountRequest("RefId_example", "Type_example") // RequestUpdateClickCountRequest | Impression payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CreateAnalyticsImpression(context.Background()).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CreateAnalyticsImpression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalyticsImpression`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CreateAnalyticsImpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticsImpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestUpdateClickCountRequest** | [**RequestUpdateClickCountRequest**](RequestUpdateClickCountRequest.md) | Impression payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateTrackerClick

> HandlerStatusResponseResponse CreateTrackerClick(ctx).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).XTenantUserId(xTenantUserId).Execute()

Record a click/impression event



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
	requestUpdateClickCountRequest := *openapiclient.NewRequestUpdateClickCountRequest("RefId_example", "Type_example") // RequestUpdateClickCountRequest | Impression payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CreateTrackerClick(context.Background()).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CreateTrackerClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrackerClick`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CreateTrackerClick`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackerClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestUpdateClickCountRequest** | [**RequestUpdateClickCountRequest**](RequestUpdateClickCountRequest.md) | Impression payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetAnalyticsOverview

> DomainStatsOverviewResponse GetAnalyticsOverview(ctx).XTenantUserId(xTenantUserId).Execute()

Get the caller's analytics overview

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
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsOverview(context.Background()).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsOverview`: DomainStatsOverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainStatsOverviewResponse**](DomainStatsOverviewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

