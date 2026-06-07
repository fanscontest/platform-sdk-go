# \UserEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PuzzleWebV2UserEventsControllerCreate**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerCreate) | **Post** /v2/user-events | Create user event
[**PuzzleWebV2UserEventsControllerDelete**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerDelete) | **Delete** /v2/user-events/{id} | Delete user event
[**PuzzleWebV2UserEventsControllerGetResultSuggestions**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerGetResultSuggestions) | **Get** /v2/user-events/{user_events_id}/result-suggestions | Get result suggestions for a user event
[**PuzzleWebV2UserEventsControllerGetStatus**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerGetStatus) | **Get** /v2/user-events/{user_events_id}/status | Get user event status
[**PuzzleWebV2UserEventsControllerIndex**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerIndex) | **Get** /v2/user-events | List user events
[**PuzzleWebV2UserEventsControllerListTemplates**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerListTemplates) | **Get** /v2/event-templates | List event templates
[**PuzzleWebV2UserEventsControllerShow**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerShow) | **Get** /v2/user-events/{id} | Get user event
[**PuzzleWebV2UserEventsControllerUpdate**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerUpdate) | **Patch** /v2/user-events/{id} | Update user event
[**PuzzleWebV2UserEventsControllerUseExistingResult**](UserEventsAPI.md#PuzzleWebV2UserEventsControllerUseExistingResult) | **Post** /v2/user-events/{user_events_id}/use-result | Copy result from another user event



## PuzzleWebV2UserEventsControllerCreate

> EventResponse PuzzleWebV2UserEventsControllerCreate(ctx).XActingAs(xActingAs).CreateUserEventRequest(createUserEventRequest).Execute()

Create user event

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
	createUserEventRequest := *openapiclient.NewCreateUserEventRequest("Name_example", "TemplateId_example") // CreateUserEventRequest | User event payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerCreate(context.Background()).XActingAs(xActingAs).CreateUserEventRequest(createUserEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerCreate`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **createUserEventRequest** | [**CreateUserEventRequest**](CreateUserEventRequest.md) | User event payload | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerDelete

> PuzzleWebV2UserEventsControllerDelete(ctx, id).XActingAs(xActingAs).Execute()

Delete user event

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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerDelete(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerDeleteRequest struct via the builder pattern


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


## PuzzleWebV2UserEventsControllerGetResultSuggestions

> ResultSuggestionsResponse PuzzleWebV2UserEventsControllerGetResultSuggestions(ctx, userEventsId).XActingAs(xActingAs).Execute()

Get result suggestions for a user event

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerGetResultSuggestions(context.Background(), userEventsId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerGetResultSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerGetResultSuggestions`: ResultSuggestionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerGetResultSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerGetResultSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**ResultSuggestionsResponse**](ResultSuggestionsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerGetStatus

> UserEventStatusResponse PuzzleWebV2UserEventsControllerGetStatus(ctx, userEventsId).XActingAs(xActingAs).Execute()

Get user event status

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerGetStatus(context.Background(), userEventsId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerGetStatus`: UserEventStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**UserEventStatusResponse**](UserEventStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerIndex

> PuzzleWebV2UserEventsControllerIndex200Response PuzzleWebV2UserEventsControllerIndex(ctx).All(all).Visibility(visibility).DateFrom(dateFrom).DateTo(dateTo).XActingAs(xActingAs).Execute()

List user events

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fanscontest/platform-sdk-go"
)

func main() {
	all := true // bool |  (optional)
	visibility := "visibility_example" // string |  (optional)
	dateFrom := time.Now() // time.Time |  (optional)
	dateTo := time.Now() // time.Time |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerIndex(context.Background()).All(all).Visibility(visibility).DateFrom(dateFrom).DateTo(dateTo).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerIndex`: PuzzleWebV2UserEventsControllerIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** |  | 
 **visibility** | **string** |  | 
 **dateFrom** | **time.Time** |  | 
 **dateTo** | **time.Time** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**PuzzleWebV2UserEventsControllerIndex200Response**](PuzzleWebV2UserEventsControllerIndex200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerListTemplates

> EventTemplatesListResponse PuzzleWebV2UserEventsControllerListTemplates(ctx).Category(category).EventType(eventType).XActingAs(xActingAs).Execute()

List event templates



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
	category := "category_example" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerListTemplates(context.Background()).Category(category).EventType(eventType).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerListTemplates`: EventTemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** |  | 
 **eventType** | **string** |  | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventTemplatesListResponse**](EventTemplatesListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerShow

> EventResponse PuzzleWebV2UserEventsControllerShow(ctx, id).XActingAs(xActingAs).Execute()

Get user event

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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerShow(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerShow`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerUpdate

> EventResponse PuzzleWebV2UserEventsControllerUpdate(ctx, id).XActingAs(xActingAs).UpdateUserEventRequest(updateUserEventRequest).Execute()

Update user event

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
	id := "id_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	updateUserEventRequest := *openapiclient.NewUpdateUserEventRequest() // UpdateUserEventRequest | User event update payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerUpdate(context.Background(), id).XActingAs(xActingAs).UpdateUserEventRequest(updateUserEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerUpdate`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **updateUserEventRequest** | [**UpdateUserEventRequest**](UpdateUserEventRequest.md) | User event update payload | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2UserEventsControllerUseExistingResult

> EventResponse PuzzleWebV2UserEventsControllerUseExistingResult(ctx, userEventsId).XActingAs(xActingAs).UseExistingResultRequest(useExistingResultRequest).Execute()

Copy result from another user event

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
	userEventsId := "userEventsId_example" // string | 
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)
	useExistingResultRequest := *openapiclient.NewUseExistingResultRequest("SourceEventId_example") // UseExistingResultRequest | Source event payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserEventsAPI.PuzzleWebV2UserEventsControllerUseExistingResult(context.Background(), userEventsId).XActingAs(xActingAs).UseExistingResultRequest(useExistingResultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.PuzzleWebV2UserEventsControllerUseExistingResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2UserEventsControllerUseExistingResult`: EventResponse
	fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.PuzzleWebV2UserEventsControllerUseExistingResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEventsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2UserEventsControllerUseExistingResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 
 **useExistingResultRequest** | [**UseExistingResultRequest**](UseExistingResultRequest.md) | Source event payload | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

