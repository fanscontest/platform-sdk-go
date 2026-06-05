# \OGImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicOgChannelByChannelId**](OGImagesAPI.md#GetPublicOgChannelByChannelId) | **Get** /v2/public/og/channel/{channelId} | Get channel data for OG image generation
[**GetPublicOgContestByContestId**](OGImagesAPI.md#GetPublicOgContestByContestId) | **Get** /v2/public/og/contest/{contestId} | Get contest data for OG image generation
[**GetPublicOgScoreboardByContestId**](OGImagesAPI.md#GetPublicOgScoreboardByContestId) | **Get** /v2/public/og/scoreboard/{contestId} | Get scoreboard data for OG image generation
[**GetPublicOgSubmissionsByParticipationId**](OGImagesAPI.md#GetPublicOgSubmissionsByParticipationId) | **Get** /v2/public/og/submissions/{participationId} | Get participation data for OG image generation



## GetPublicOgChannelByChannelId

> HandlerOGChannelData GetPublicOgChannelByChannelId(ctx, channelId).AccessCode(accessCode).XTenantUserId(xTenantUserId).Execute()

Get channel data for OG image generation



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
	channelId := "channelId_example" // string | Channel ID
	accessCode := "accessCode_example" // string | Access code for restricted channels (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OGImagesAPI.GetPublicOgChannelByChannelId(context.Background(), channelId).AccessCode(accessCode).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OGImagesAPI.GetPublicOgChannelByChannelId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOgChannelByChannelId`: HandlerOGChannelData
	fmt.Fprintf(os.Stdout, "Response from `OGImagesAPI.GetPublicOgChannelByChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOgChannelByChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessCode** | **string** | Access code for restricted channels | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**HandlerOGChannelData**](HandlerOGChannelData.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOgContestByContestId

> HandlerOGContestData GetPublicOgContestByContestId(ctx, contestId).XTenantUserId(xTenantUserId).Execute()

Get contest data for OG image generation



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
	contestId := "contestId_example" // string | Contest ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OGImagesAPI.GetPublicOgContestByContestId(context.Background(), contestId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OGImagesAPI.GetPublicOgContestByContestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOgContestByContestId`: HandlerOGContestData
	fmt.Fprintf(os.Stdout, "Response from `OGImagesAPI.GetPublicOgContestByContestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOgContestByContestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**HandlerOGContestData**](HandlerOGContestData.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOgScoreboardByContestId

> HandlerOGScoreboardData GetPublicOgScoreboardByContestId(ctx, contestId).XTenantUserId(xTenantUserId).Execute()

Get scoreboard data for OG image generation



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
	contestId := "contestId_example" // string | Contest ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OGImagesAPI.GetPublicOgScoreboardByContestId(context.Background(), contestId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OGImagesAPI.GetPublicOgScoreboardByContestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOgScoreboardByContestId`: HandlerOGScoreboardData
	fmt.Fprintf(os.Stdout, "Response from `OGImagesAPI.GetPublicOgScoreboardByContestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOgScoreboardByContestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**HandlerOGScoreboardData**](HandlerOGScoreboardData.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOgSubmissionsByParticipationId

> HandlerOGParticipationData GetPublicOgSubmissionsByParticipationId(ctx, participationId).XTenantUserId(xTenantUserId).Execute()

Get participation data for OG image generation



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
	participationId := "participationId_example" // string | Participation ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OGImagesAPI.GetPublicOgSubmissionsByParticipationId(context.Background(), participationId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OGImagesAPI.GetPublicOgSubmissionsByParticipationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOgSubmissionsByParticipationId`: HandlerOGParticipationData
	fmt.Fprintf(os.Stdout, "Response from `OGImagesAPI.GetPublicOgSubmissionsByParticipationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**participationId** | **string** | Participation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOgSubmissionsByParticipationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**HandlerOGParticipationData**](HandlerOGParticipationData.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

