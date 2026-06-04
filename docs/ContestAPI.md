# \ContestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicContestsByIdGroups**](ContestAPI.md#GetPublicContestsByIdGroups) | **Get** /v2/public/contests/{id}/groups | Get Contest Groups



## GetPublicContestsByIdGroups

> DomainContestGroups GetPublicContestsByIdGroups(ctx, id).Execute()

Get Contest Groups



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
	id := "id_example" // string | Contest ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestAPI.GetPublicContestsByIdGroups(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestAPI.GetPublicContestsByIdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdGroups`: DomainContestGroups
	fmt.Fprintf(os.Stdout, "Response from `ContestAPI.GetPublicContestsByIdGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainContestGroups**](DomainContestGroups.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

