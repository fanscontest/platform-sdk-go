# \FeedsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeed**](FeedsAPI.md#GetFeed) | **Get** /v2/feed | Get the trending feed (tenant-scoped, cursor-paginated)
[**GetIdentitiesByPiidFeedPersonalized**](FeedsAPI.md#GetIdentitiesByPiidFeedPersonalized) | **Get** /v2/identities/{piid}/feed-personalized | Get the personalized feed for a platform identity



## GetFeed

> DomainFeedItemListResponse GetFeed(ctx).Limit(limit).Cursor(cursor).Region(region).XActingAs(xActingAs).Execute()

Get the trending feed (tenant-scoped, cursor-paginated)

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetFeed(context.Background()).Limit(limit).Cursor(cursor).Region(region).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeed`: DomainFeedItemListResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedsAPI.GetFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Page size (1-200, default 20) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **region** | **string** | Optional ISO 3166-1 alpha-2 region filter | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainFeedItemListResponse**](DomainFeedItemListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidFeedPersonalized

> DomainFeedItemListResponse GetIdentitiesByPiidFeedPersonalized(ctx, piid).Limit(limit).Cursor(cursor).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedsAPI.GetIdentitiesByPiidFeedPersonalized(context.Background(), piid).Limit(limit).Cursor(cursor).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedsAPI.GetIdentitiesByPiidFeedPersonalized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidFeedPersonalized`: DomainFeedItemListResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainFeedItemListResponse**](DomainFeedItemListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

