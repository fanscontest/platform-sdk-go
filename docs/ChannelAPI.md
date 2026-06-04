# \ChannelAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannelsByIdReport**](ChannelAPI.md#CreateChannelsByIdReport) | **Post** /v2/channels/{id}/report | Report a channel
[**PatchChannelsByIdAccessCode**](ChannelAPI.md#PatchChannelsByIdAccessCode) | **Patch** /v2/channels/{id}/access-code | Update Channel Access Code



## CreateChannelsByIdReport

> string CreateChannelsByIdReport(ctx, id).Execute()

Report a channel



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
	id := "id_example" // string | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.CreateChannelsByIdReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.CreateChannelsByIdReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateChannelsByIdReport`: string
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.CreateChannelsByIdReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChannelsByIdReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchChannelsByIdAccessCode

> DomainChannel PatchChannelsByIdAccessCode(ctx, id).RequestUpdateChannelAccessCodeRequest(requestUpdateChannelAccessCodeRequest).Execute()

Update Channel Access Code



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
	id := "id_example" // string | Channel ID
	requestUpdateChannelAccessCodeRequest := *openapiclient.NewRequestUpdateChannelAccessCodeRequest("AccessCode_example") // RequestUpdateChannelAccessCodeRequest | request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PatchChannelsByIdAccessCode(context.Background(), id).RequestUpdateChannelAccessCodeRequest(requestUpdateChannelAccessCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PatchChannelsByIdAccessCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchChannelsByIdAccessCode`: DomainChannel
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PatchChannelsByIdAccessCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchChannelsByIdAccessCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateChannelAccessCodeRequest** | [**RequestUpdateChannelAccessCodeRequest**](RequestUpdateChannelAccessCodeRequest.md) | request | 

### Return type

[**DomainChannel**](DomainChannel.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

