# \AnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyticsImpression**](AnalyticsAPI.md#CreateAnalyticsImpression) | **Post** /v2/analytics/impression | Record a click/impression event
[**CreateTrackerClick**](AnalyticsAPI.md#CreateTrackerClick) | **Post** /v2/tracker/click | Record a click/impression event
[**GetAnalyticsOverview**](AnalyticsAPI.md#GetAnalyticsOverview) | **Get** /v2/analytics/overview | Get the caller&#39;s analytics overview



## CreateAnalyticsImpression

> HandlerStatusResponse CreateAnalyticsImpression(ctx).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CreateAnalyticsImpression(context.Background()).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CreateAnalyticsImpression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalyticsImpression`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CreateAnalyticsImpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticsImpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestUpdateClickCountRequest** | [**RequestUpdateClickCountRequest**](RequestUpdateClickCountRequest.md) | Impression payload | 

### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackerClick

> HandlerStatusResponse CreateTrackerClick(ctx).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CreateTrackerClick(context.Background()).RequestUpdateClickCountRequest(requestUpdateClickCountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CreateTrackerClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrackerClick`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CreateTrackerClick`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrackerClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestUpdateClickCountRequest** | [**RequestUpdateClickCountRequest**](RequestUpdateClickCountRequest.md) | Impression payload | 

### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsOverview

> DomainStatsOverview GetAnalyticsOverview(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsOverview(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsOverview`: DomainStatsOverview
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOverviewRequest struct via the builder pattern


### Return type

[**DomainStatsOverview**](DomainStatsOverview.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

