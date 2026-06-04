# \ContestPuzzlesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PuzzleWebV2ContestPuzzleControllerGetPuzzle**](ContestPuzzlesAPI.md#PuzzleWebV2ContestPuzzleControllerGetPuzzle) | **Get** /v2/contest-puzzles/{contest_id} | Get contest source
[**PuzzleWebV2ContestPuzzleControllerRankings**](ContestPuzzlesAPI.md#PuzzleWebV2ContestPuzzleControllerRankings) | **Get** /v2/contest-puzzles/{contest_id}/rankings | Get contest rankings
[**PuzzleWebV2ContestPuzzleControllerStatus**](ContestPuzzlesAPI.md#PuzzleWebV2ContestPuzzleControllerStatus) | **Get** /v2/contest-puzzles/{contest_id}/status | Get contest status



## PuzzleWebV2ContestPuzzleControllerGetPuzzle

> ContestPuzzleResponse PuzzleWebV2ContestPuzzleControllerGetPuzzle(ctx, contestId).Execute()

Get contest source



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
	contestId := "contestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerGetPuzzle(context.Background(), contestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerGetPuzzle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2ContestPuzzleControllerGetPuzzle`: ContestPuzzleResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerGetPuzzle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2ContestPuzzleControllerGetPuzzleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContestPuzzleResponse**](ContestPuzzleResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2ContestPuzzleControllerRankings

> RankingsResponse PuzzleWebV2ContestPuzzleControllerRankings(ctx, contestId).Execute()

Get contest rankings



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
	contestId := "contestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerRankings(context.Background(), contestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerRankings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2ContestPuzzleControllerRankings`: RankingsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerRankings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2ContestPuzzleControllerRankingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RankingsResponse**](RankingsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2ContestPuzzleControllerStatus

> ContestStatusResponse PuzzleWebV2ContestPuzzleControllerStatus(ctx, contestId).Execute()

Get contest status



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
	contestId := "contestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerStatus(context.Background(), contestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2ContestPuzzleControllerStatus`: ContestStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestPuzzlesAPI.PuzzleWebV2ContestPuzzleControllerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2ContestPuzzleControllerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContestStatusResponse**](ContestStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

