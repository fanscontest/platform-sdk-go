# \PuzzleSheetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet) | **Post** /v2/puzzle-sheets/{id}/clone | Clone puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerCreate**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerCreate) | **Post** /v2/puzzle-sheets | Create puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet) | **Post** /v2/puzzle-sheets/adhoc-community-sheet | Create adhoc community puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet) | **Post** /v2/puzzle-sheets/adhoc-user-sheet | Create adhoc user puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerDelete**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerDelete) | **Delete** /v2/puzzle-sheets/{id} | Delete puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerIndex**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerIndex) | **Get** /v2/puzzle-sheets | List puzzle sheets
[**PuzzleWebV2PuzzleSheetControllerShow**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerShow) | **Get** /v2/puzzle-sheets/{id} | Get puzzle sheet
[**PuzzleWebV2PuzzleSheetControllerUpdate**](PuzzleSheetsAPI.md#PuzzleWebV2PuzzleSheetControllerUpdate) | **Patch** /v2/puzzle-sheets/{id} | Update puzzle sheet



## PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet(ctx, id).ClonePuzzleSheetRequest(clonePuzzleSheetRequest).Execute()

Clone puzzle sheet

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
	clonePuzzleSheetRequest := *openapiclient.NewClonePuzzleSheetRequest() // ClonePuzzleSheetRequest | Clone payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet(context.Background(), id).ClonePuzzleSheetRequest(clonePuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerClonePuzzleSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerClonePuzzleSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clonePuzzleSheetRequest** | [**ClonePuzzleSheetRequest**](ClonePuzzleSheetRequest.md) | Clone payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerCreate

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerCreate(ctx).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()

Create puzzle sheet

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
	createPuzzleSheetRequest := *openapiclient.NewCreatePuzzleSheetRequest("Name_example", "private") // CreatePuzzleSheetRequest | Puzzle sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreate(context.Background()).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerCreate`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPuzzleSheetRequest** | [**CreatePuzzleSheetRequest**](CreatePuzzleSheetRequest.md) | Puzzle sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet(ctx).CreateAdhocCommunitySheetRequest(createAdhocCommunitySheetRequest).Execute()

Create adhoc community puzzle sheet



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
	createAdhocCommunitySheetRequest := *openapiclient.NewCreateAdhocCommunitySheetRequest("Difficulty_example", int32(123), "Topic_example") // CreateAdhocCommunitySheetRequest | Adhoc community sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet(context.Background()).CreateAdhocCommunitySheetRequest(createAdhocCommunitySheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerCreateAdhocCommunitySheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdhocCommunitySheetRequest** | [**CreateAdhocCommunitySheetRequest**](CreateAdhocCommunitySheetRequest.md) | Adhoc community sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet(ctx).CreateAdhocUserSheetRequest(createAdhocUserSheetRequest).Execute()

Create adhoc user puzzle sheet



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
	createAdhocUserSheetRequest := *openapiclient.NewCreateAdhocUserSheetRequest(int32(123), "PuzzleSheetId_example") // CreateAdhocUserSheetRequest | Adhoc user sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet(context.Background()).CreateAdhocUserSheetRequest(createAdhocUserSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerCreateAdhocUserSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdhocUserSheetRequest** | [**CreateAdhocUserSheetRequest**](CreateAdhocUserSheetRequest.md) | Adhoc user sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerDelete

> PuzzleWebV2PuzzleSheetControllerDelete(ctx, id).Execute()

Delete puzzle sheet

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PuzzleWebV2PuzzleSheetControllerIndex

> PuzzleSheetsListResponse PuzzleWebV2PuzzleSheetControllerIndex(ctx).Cursor(cursor).Limit(limit).Execute()

List puzzle sheets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerIndex(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerIndex`: PuzzleSheetsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**PuzzleSheetsListResponse**](PuzzleSheetsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerShow

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerShow(ctx, id).Execute()

Get puzzle sheet

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerShow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerShow`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PuzzleSheetControllerUpdate

> PuzzleSheetResponse PuzzleWebV2PuzzleSheetControllerUpdate(ctx, id).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()

Update puzzle sheet

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
	createPuzzleSheetRequest := *openapiclient.NewCreatePuzzleSheetRequest("Name_example", "private") // CreatePuzzleSheetRequest | Puzzle sheet payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerUpdate(context.Background(), id).CreatePuzzleSheetRequest(createPuzzleSheetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PuzzleSheetControllerUpdate`: PuzzleSheetResponse
	fmt.Fprintf(os.Stdout, "Response from `PuzzleSheetsAPI.PuzzleWebV2PuzzleSheetControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PuzzleSheetControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPuzzleSheetRequest** | [**CreatePuzzleSheetRequest**](CreatePuzzleSheetRequest.md) | Puzzle sheet payload | 

### Return type

[**PuzzleSheetResponse**](PuzzleSheetResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

