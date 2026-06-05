# \PredictionSlipsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PuzzleWebV2PredictionControllerClonePredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerClonePredictionSlip) | **Post** /v2/prediction-slips/{id}/clone | Clone prediction slip
[**PuzzleWebV2PredictionControllerCreatePredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerCreatePredictionSlip) | **Post** /v2/prediction-slips | Create prediction slip
[**PuzzleWebV2PredictionControllerDeletePredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerDeletePredictionSlip) | **Delete** /v2/prediction-slips/{id} | Delete prediction slip
[**PuzzleWebV2PredictionControllerListCuratedSlips**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerListCuratedSlips) | **Get** /v2/public/curated-slips | List curated prediction slips
[**PuzzleWebV2PredictionControllerListPredictionSlips**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerListPredictionSlips) | **Get** /v2/prediction-slips | List the caller&#39;s prediction slips
[**PuzzleWebV2PredictionControllerShowOpenPredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerShowOpenPredictionSlip) | **Get** /v2/prediction-slips/open | Get the caller&#39;s open prediction slip
[**PuzzleWebV2PredictionControllerShowPredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerShowPredictionSlip) | **Get** /v2/prediction-slips/{id} | Get prediction slip
[**PuzzleWebV2PredictionControllerUpdatePredictionSlip**](PredictionSlipsAPI.md#PuzzleWebV2PredictionControllerUpdatePredictionSlip) | **Put** /v2/prediction-slips/{id} | Update prediction slip



## PuzzleWebV2PredictionControllerClonePredictionSlip

> PredictionSlipResponse PuzzleWebV2PredictionControllerClonePredictionSlip(ctx, id).XTenantUserId(xTenantUserId).ClonePredictionSlipRequest(clonePredictionSlipRequest).Execute()

Clone prediction slip

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)
	clonePredictionSlipRequest := *openapiclient.NewClonePredictionSlipRequest() // ClonePredictionSlipRequest | Clone payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerClonePredictionSlip(context.Background(), id).XTenantUserId(xTenantUserId).ClonePredictionSlipRequest(clonePredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerClonePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerClonePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerClonePredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerClonePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 
 **clonePredictionSlipRequest** | [**ClonePredictionSlipRequest**](ClonePredictionSlipRequest.md) | Clone payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerCreatePredictionSlip

> PredictionSlipResponse PuzzleWebV2PredictionControllerCreatePredictionSlip(ctx).XTenantUserId(xTenantUserId).CreatePredictionSlipRequest(createPredictionSlipRequest).Execute()

Create prediction slip

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)
	createPredictionSlipRequest := *openapiclient.NewCreatePredictionSlipRequest([]openapiclient.CreatePredictionSlipRequestEventsInner{*openapiclient.NewCreatePredictionSlipRequestEventsInner("EsEventId_example", []openapiclient.CreatePredictionSlipRequestEventsInnerItemsInner{*openapiclient.NewCreatePredictionSlipRequestEventsInnerItemsInner()})}, "Name_example") // CreatePredictionSlipRequest | Prediction slip payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerCreatePredictionSlip(context.Background()).XTenantUserId(xTenantUserId).CreatePredictionSlipRequest(createPredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerCreatePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerCreatePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerCreatePredictionSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerCreatePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 
 **createPredictionSlipRequest** | [**CreatePredictionSlipRequest**](CreatePredictionSlipRequest.md) | Prediction slip payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerDeletePredictionSlip

> PuzzleWebV2PredictionControllerDeletePredictionSlip(ctx, id).XTenantUserId(xTenantUserId).Execute()

Delete prediction slip

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerDeletePredictionSlip(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerDeletePredictionSlip``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerDeletePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## PuzzleWebV2PredictionControllerListCuratedSlips

> PredictionSlipsListResponse PuzzleWebV2PredictionControllerListCuratedSlips(ctx).Family(family).Tags(tags).IsPublic(isPublic).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List curated prediction slips

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
	family := "family_example" // string |  (optional)
	tags := "tags_example" // string |  (optional)
	isPublic := true // bool |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerListCuratedSlips(context.Background()).Family(family).Tags(tags).IsPublic(isPublic).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerListCuratedSlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerListCuratedSlips`: PredictionSlipsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerListCuratedSlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerListCuratedSlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **family** | **string** |  | 
 **tags** | **string** |  | 
 **isPublic** | **bool** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipsListResponse**](PredictionSlipsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerListPredictionSlips

> PredictionSlipsListResponse PuzzleWebV2PredictionControllerListPredictionSlips(ctx).Locked(locked).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List the caller's prediction slips

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
	locked := true // bool |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerListPredictionSlips(context.Background()).Locked(locked).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerListPredictionSlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerListPredictionSlips`: PredictionSlipsListResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerListPredictionSlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerListPredictionSlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locked** | **bool** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipsListResponse**](PredictionSlipsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerShowOpenPredictionSlip

> OpenPredictionSlipResponse PuzzleWebV2PredictionControllerShowOpenPredictionSlip(ctx).XTenantUserId(xTenantUserId).Execute()

Get the caller's open prediction slip



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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowOpenPredictionSlip(context.Background()).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowOpenPredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerShowOpenPredictionSlip`: OpenPredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowOpenPredictionSlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerShowOpenPredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**OpenPredictionSlipResponse**](OpenPredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerShowPredictionSlip

> PredictionSlipResponse PuzzleWebV2PredictionControllerShowPredictionSlip(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get prediction slip

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowPredictionSlip(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowPredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerShowPredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerShowPredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerShowPredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PuzzleWebV2PredictionControllerUpdatePredictionSlip

> PredictionSlipResponse PuzzleWebV2PredictionControllerUpdatePredictionSlip(ctx, id).XTenantUserId(xTenantUserId).UpdatePredictionSlipRequest(updatePredictionSlipRequest).Execute()

Update prediction slip

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)
	updatePredictionSlipRequest := *openapiclient.NewUpdatePredictionSlipRequest() // UpdatePredictionSlipRequest | Prediction slip update payload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PredictionSlipsAPI.PuzzleWebV2PredictionControllerUpdatePredictionSlip(context.Background(), id).XTenantUserId(xTenantUserId).UpdatePredictionSlipRequest(updatePredictionSlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PredictionSlipsAPI.PuzzleWebV2PredictionControllerUpdatePredictionSlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PuzzleWebV2PredictionControllerUpdatePredictionSlip`: PredictionSlipResponse
	fmt.Fprintf(os.Stdout, "Response from `PredictionSlipsAPI.PuzzleWebV2PredictionControllerUpdatePredictionSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPuzzleWebV2PredictionControllerUpdatePredictionSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 
 **updatePredictionSlipRequest** | [**UpdatePredictionSlipRequest**](UpdatePredictionSlipRequest.md) | Prediction slip update payload | 

### Return type

[**PredictionSlipResponse**](PredictionSlipResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

