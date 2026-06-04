# \BlocksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlocks**](BlocksAPI.md#CreateBlocks) | **Post** /v2/blocks | Create Block
[**DeleteBlocksById**](BlocksAPI.md#DeleteBlocksById) | **Delete** /v2/blocks/{id} | Delete Block
[**GetBlocks**](BlocksAPI.md#GetBlocks) | **Get** /v2/blocks | Get Blocked Members



## CreateBlocks

> DomainBlock CreateBlocks(ctx).HandlerCreateBlockRequest(handlerCreateBlockRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.CreateBlocks(context.Background()).HandlerCreateBlockRequest(handlerCreateBlockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.CreateBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlocks`: DomainBlock
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.CreateBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlerCreateBlockRequest** | [**HandlerCreateBlockRequest**](HandlerCreateBlockRequest.md) | Block payload; optional channel_id scopes the block to a channel | 

### Return type

[**DomainBlock**](DomainBlock.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlocksById

> bool DeleteBlocksById(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocksAPI.DeleteBlocksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.DeleteBlocksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlocksById`: bool
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


### Return type

**bool**

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlocks

> []DomainBlock GetBlocks(ctx).Execute()

Get Blocked Members



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
	resp, r, err := apiClient.BlocksAPI.GetBlocks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlocks`: []DomainBlock
	fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetBlocks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocksRequest struct via the builder pattern


### Return type

[**[]DomainBlock**](DomainBlock.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

