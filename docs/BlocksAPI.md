# \BlocksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlocks**](BlocksAPI.md#CreateBlocks) | **Post** /v2/blocks | Create Block
[**DeleteBlocksById**](BlocksAPI.md#DeleteBlocksById) | **Delete** /v2/blocks/{id} | Delete Block
[**GetBlocks**](BlocksAPI.md#GetBlocks) | **Get** /v2/blocks | Get Blocked Members (cursor-paginated)



## CreateBlocks

> DomainBlockResponse CreateBlocks(ctx).HandlerCreateBlockRequest(handlerCreateBlockRequest).XTenantUserId(xTenantUserId).Execute()

Create Block



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
	handlerCreateBlockRequest := *openapiclient.NewHandlerCreateBlockRequest("PlatformIdentityId_example") // HandlerCreateBlockRequest | Block payload; optional channel_id scopes the block to a channel
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.CreateBlocks(context.Background()).HandlerCreateBlockRequest(handlerCreateBlockRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.CreateBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlocks`: DomainBlockResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.CreateBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlerCreateBlockRequest** | [**HandlerCreateBlockRequest**](HandlerCreateBlockRequest.md) | Block payload; optional channel_id scopes the block to a channel | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBlockResponse**](DomainBlockResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlocksById

> DeleteBlocksById200Response DeleteBlocksById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Delete Block



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
	id := "id_example" // string | Blocked PlatformIdentity ID to unblock
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.DeleteBlocksById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.DeleteBlocksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlocksById`: DeleteBlocksById200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.DeleteBlocksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Blocked PlatformIdentity ID to unblock | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlocksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DeleteBlocksById200Response**](DeleteBlocksById200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocks

> DomainBlockListResponse GetBlocks(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get Blocked Members (cursor-paginated)



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
	resp, r, err := apiClient.BlocksAPI.GetBlocks(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocks`: DomainBlockListResponse
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBlockListResponse**](DomainBlockListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

