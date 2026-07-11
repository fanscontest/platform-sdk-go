# \LeaderboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannelsByIdStandings**](LeaderboardAPI.md#GetChannelsByIdStandings) | **Get** /v2/channels/{id}/standings | Get a channel&#39;s season standings (ranked by cumulative score)



## GetChannelsByIdStandings

> DomainLeaderboardResponse GetChannelsByIdStandings(ctx, id).Window(window).Actor(actor).XActingAs(xActingAs).Execute()

Get a channel's season standings (ranked by cumulative score)



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
	window := "window_example" // string | Standings window: week | all_time (default all_time) (optional)
	actor := "actor_example" // string | member (truncated view) | owner (full board, default) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetChannelsByIdStandings(context.Background(), id).Window(window).Actor(actor).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetChannelsByIdStandings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdStandings`: DomainLeaderboardResponse
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetChannelsByIdStandings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsByIdStandingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **window** | **string** | Standings window: week | all_time (default all_time) | 
 **actor** | **string** | member (truncated view) | owner (full board, default) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainLeaderboardResponse**](DomainLeaderboardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

