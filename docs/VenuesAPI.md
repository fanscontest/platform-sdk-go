# \VenuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVenues**](VenuesAPI.md#CreateVenues) | **Post** /v2/venues | Create a venue under the calling tenant
[**CreateVenuesByVenueIdChannels**](VenuesAPI.md#CreateVenuesByVenueIdChannels) | **Post** /v2/venues/{venueId}/channels | Create a channel under one of the tenant&#39;s venues
[**GetVenues**](VenuesAPI.md#GetVenues) | **Get** /v2/venues | List venues owned by the calling tenant
[**GetVenuesByVenueIdChannels**](VenuesAPI.md#GetVenuesByVenueIdChannels) | **Get** /v2/venues/{venueId}/channels | List channels under one of the tenant&#39;s venues



## CreateVenues

> HandlerStatusResponseResponse CreateVenues(ctx).HandlerCreatePlatformVenueRequest(handlerCreatePlatformVenueRequest).XTenantUserId(xTenantUserId).Execute()

Create a venue under the calling tenant

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
	handlerCreatePlatformVenueRequest := *openapiclient.NewHandlerCreatePlatformVenueRequest("Name_example", "Type_example") // HandlerCreatePlatformVenueRequest | Venue payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.CreateVenues(context.Background()).HandlerCreatePlatformVenueRequest(handlerCreatePlatformVenueRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.CreateVenues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVenues`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.CreateVenues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVenuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlerCreatePlatformVenueRequest** | [**HandlerCreatePlatformVenueRequest**](HandlerCreatePlatformVenueRequest.md) | Venue payload | 
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


## CreateVenuesByVenueIdChannels

> HandlerStatusResponseResponse CreateVenuesByVenueIdChannels(ctx, venueId).HandlerCreatePlatformChannelRequest(handlerCreatePlatformChannelRequest).XTenantUserId(xTenantUserId).Execute()

Create a channel under one of the tenant's venues

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
	venueId := "venueId_example" // string | Venue ID
	handlerCreatePlatformChannelRequest := *openapiclient.NewHandlerCreatePlatformChannelRequest("AccessType_example", "Name_example", int32(123)) // HandlerCreatePlatformChannelRequest | Channel payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.CreateVenuesByVenueIdChannels(context.Background(), venueId).HandlerCreatePlatformChannelRequest(handlerCreatePlatformChannelRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.CreateVenuesByVenueIdChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVenuesByVenueIdChannels`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.CreateVenuesByVenueIdChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**venueId** | **string** | Venue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVenuesByVenueIdChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerCreatePlatformChannelRequest** | [**HandlerCreatePlatformChannelRequest**](HandlerCreatePlatformChannelRequest.md) | Channel payload | 
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


## GetVenues

> DomainVenueListResponse GetVenues(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List venues owned by the calling tenant

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
	resp, r, err := apiClient.VenuesAPI.GetVenues(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.GetVenues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVenues`: DomainVenueListResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.GetVenues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVenuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainVenueListResponse**](DomainVenueListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVenuesByVenueIdChannels

> DomainChannelListResponse GetVenuesByVenueIdChannels(ctx, venueId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List channels under one of the tenant's venues

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
	venueId := "venueId_example" // string | Venue ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VenuesAPI.GetVenuesByVenueIdChannels(context.Background(), venueId).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VenuesAPI.GetVenuesByVenueIdChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVenuesByVenueIdChannels`: DomainChannelListResponse
	fmt.Fprintf(os.Stdout, "Response from `VenuesAPI.GetVenuesByVenueIdChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**venueId** | **string** | Venue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVenuesByVenueIdChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainChannelListResponse**](DomainChannelListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

