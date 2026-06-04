# \TeamPreferenceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateTeamPreferenceByTeamSetId**](TeamPreferenceAPI.md#UpdateTeamPreferenceByTeamSetId) | **Put** /v2/team-preference/{teamSetId} | Set the caller&#39;s default team for a logical team set



## UpdateTeamPreferenceByTeamSetId

> HandlerStatusResponse UpdateTeamPreferenceByTeamSetId(ctx, teamSetId).RequestUpsertTeamPreferenceRequest(requestUpsertTeamPreferenceRequest).Execute()

Set the caller's default team for a logical team set

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
	teamSetId := "teamSetId_example" // string | Team Set ID
	requestUpsertTeamPreferenceRequest := *openapiclient.NewRequestUpsertTeamPreferenceRequest() // RequestUpsertTeamPreferenceRequest | Preference payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamPreferenceAPI.UpdateTeamPreferenceByTeamSetId(context.Background(), teamSetId).RequestUpsertTeamPreferenceRequest(requestUpsertTeamPreferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamPreferenceAPI.UpdateTeamPreferenceByTeamSetId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamPreferenceByTeamSetId`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamPreferenceAPI.UpdateTeamPreferenceByTeamSetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamSetId** | **string** | Team Set ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamPreferenceByTeamSetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpsertTeamPreferenceRequest** | [**RequestUpsertTeamPreferenceRequest**](RequestUpsertTeamPreferenceRequest.md) | Preference payload | 

### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

