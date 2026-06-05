# \FeedsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentitiesByPiidFeedPersonalized**](FeedsAPI.md#GetIdentitiesByPiidFeedPersonalized) | **Get** /v2/identities/{piid}/feed-personalized | Get the personalized feed for a platform identity
[**GetPublicFeed**](FeedsAPI.md#GetPublicFeed) | **Get** /v2/public/feed | Get the public feed (trending content, unauthenticated)



## GetIdentitiesByPiidFeedPersonalized

> GetIdentitiesByPiidFeedPersonalized200Response GetIdentitiesByPiidFeedPersonalized(ctx, piid).Limit(limit).Cursor(cursor).XTenantUserId(xTenantUserId).Execute()

Get the personalized feed for a platform identity



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
	piid := "piid_example" // string | Platform Identity ID
	limit := int32(56) // int32 | Page size (1-200, default 20) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetIdentitiesByPiidFeedPersonalized(context.Background(), piid).Limit(limit).Cursor(cursor).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetIdentitiesByPiidFeedPersonalized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidFeedPersonalized`: GetIdentitiesByPiidFeedPersonalized200Response
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetIdentitiesByPiidFeedPersonalized`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidFeedPersonalizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size (1-200, default 20) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetIdentitiesByPiidFeedPersonalized200Response**](GetIdentitiesByPiidFeedPersonalized200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicFeed

> GetIdentitiesByPiidFeedPersonalized200Response GetPublicFeed(ctx).Limit(limit).Cursor(cursor).Region(region).XTenantUserId(xTenantUserId).Execute()

Get the public feed (trending content, unauthenticated)

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
	limit := int32(56) // int32 | Page size (1-200, default 20) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	region := "region_example" // string | Optional ISO 3166-1 alpha-2 region filter (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetPublicFeed(context.Background()).Limit(limit).Cursor(cursor).Region(region).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetPublicFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicFeed`: GetIdentitiesByPiidFeedPersonalized200Response
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetPublicFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Page size (1-200, default 20) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **region** | **string** | Optional ISO 3166-1 alpha-2 region filter | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetIdentitiesByPiidFeedPersonalized200Response**](GetIdentitiesByPiidFeedPersonalized200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

