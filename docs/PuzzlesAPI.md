# \PuzzlesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PuzzleWebV2PuzzleControllerCreate**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerCreate) | **Post** /v2/puzzles | Create puzzle
[**PuzzleWebV2PuzzleControllerDelete**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerDelete) | **Delete** /v2/puzzles/{id} | Delete puzzle
[**PuzzleWebV2PuzzleControllerGetTopicConfiguration**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerGetTopicConfiguration) | **Get** /v2/puzzles/get-topic-configuration | Get topic configuration
[**PuzzleWebV2PuzzleControllerGetUniqueCategories**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerGetUniqueCategories) | **Get** /v2/puzzles/categories | List unique categories
[**PuzzleWebV2PuzzleControllerIndex**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerIndex) | **Get** /v2/puzzles | List puzzles
[**PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch) | **Get** /v2/puzzles/migrate | Migrate puzzles to OpenSearch
[**PuzzleWebV2PuzzleControllerSearchPuzzleBank**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerSearchPuzzleBank) | **Get** /v2/puzzles/search-puzzle-bank | Search puzzle bank
[**PuzzleWebV2PuzzleControllerShow**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerShow) | **Get** /v2/puzzles/{id} | Get puzzle
[**PuzzleWebV2PuzzleControllerUpdate**](PuzzlesAPI.md#PuzzleWebV2PuzzleControllerUpdate) | **Patch** /v2/puzzles/{id} | Update puzzle



## PuzzleWebV2PuzzleControllerCreate

> PuzzleResponse PuzzleWebV2PuzzleControllerCreate(ctx).XActingAs(xActingAs).CreatePuzzleRequest(createPuzzleRequest).Execute()

Create puzzle

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	createPuzzleRequest := *openapiclient.NewCreatePuzzleRequest(map[string]interface{}{"key": interface{}(123)}) // CreatePuzzleRequest | Create puzzle request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerCreate(context.Background()).XActingAs(xActingAs).CreatePuzzleRequest(createPuzzleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerCreate`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createPuzzleRequest** | [**CreatePuzzleRequest**](CreatePuzzleRequest.md) | Create puzzle request | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerDelete

> PuzzleWebV2PuzzleControllerDelete(ctx, id).XActingAs(xActingAs).Execute()

Delete puzzle

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerDelete(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

 (empty response body)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerGetTopicConfiguration

> TopicConfigurationResponse PuzzleWebV2PuzzleControllerGetTopicConfiguration(ctx).Topic(topic).XActingAs(xActingAs).Execute()

Get topic configuration



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
	topic := "topic_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerGetTopicConfiguration(context.Background()).Topic(topic).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerGetTopicConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerGetTopicConfiguration`: TopicConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerGetTopicConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerGetTopicConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**TopicConfigurationResponse**](TopicConfigurationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerGetUniqueCategories

> CategoriesResponse PuzzleWebV2PuzzleControllerGetUniqueCategories(ctx).XActingAs(xActingAs).Execute()

List unique categories



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerGetUniqueCategories(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerGetUniqueCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerGetUniqueCategories`: CategoriesResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerGetUniqueCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerGetUniqueCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CategoriesResponse**](CategoriesResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerIndex

> PuzzlesListResponse PuzzleWebV2PuzzleControllerIndex(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List puzzles



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
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerIndex(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerIndex`: PuzzlesListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzlesListResponse**](PuzzlesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch

> MessageResponse PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch(ctx).XActingAs(xActingAs).Execute()

Migrate puzzles to OpenSearch



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch`: MessageResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerMigratePuzzlesToOpensearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerSearchPuzzleBank

> PuzzlesListResponse PuzzleWebV2PuzzleControllerSearchPuzzleBank(ctx).Topic(topic).Difficulty(difficulty).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Search puzzle bank



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
	topic := "topic_example" // string | 
	difficulty := "difficulty_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerSearchPuzzleBank(context.Background()).Topic(topic).Difficulty(difficulty).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerSearchPuzzleBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerSearchPuzzleBank`: PuzzlesListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerSearchPuzzleBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerSearchPuzzleBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **string** |  | 
 **difficulty** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzlesListResponse**](PuzzlesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerShow

> PuzzleResponse PuzzleWebV2PuzzleControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get puzzle

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerShow`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleControllerUpdate

> PuzzleResponse PuzzleWebV2PuzzleControllerUpdate(ctx, id).XActingAs(xActingAs).UpdatePuzzleRequest(updatePuzzleRequest).Execute()

Update puzzle

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	updatePuzzleRequest := *openapiclient.NewUpdatePuzzleRequest() // UpdatePuzzleRequest | Update puzzle request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzlesAPI.PuzzleWebV2PuzzleControllerUpdate(context.Background(), id).XActingAs(xActingAs).UpdatePuzzleRequest(updatePuzzleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzlesAPI.PuzzleWebV2PuzzleControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleControllerUpdate`: PuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzlesAPI.PuzzleWebV2PuzzleControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **updatePuzzleRequest** | [**UpdatePuzzleRequest**](UpdatePuzzleRequest.md) | Update puzzle request | 

### Return type

[**PuzzleResponse**](PuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

