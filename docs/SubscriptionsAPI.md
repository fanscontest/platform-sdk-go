# \SubscriptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptions**](SubscriptionsAPI.md#CreateSubscriptions) | **Post** /v2/subscriptions | Subscribe the caller to a channel
[**DeleteSubscriptionsById**](SubscriptionsAPI.md#DeleteSubscriptionsById) | **Delete** /v2/subscriptions/{id} | Unsubscribe from Channel
[**GetChannelsByIdSubscriptions**](SubscriptionsAPI.md#GetChannelsByIdSubscriptions) | **Get** /v2/channels/{id}/subscriptions | Get channel subscribers
[**GetIdentitiesByPiidSubscriptions**](SubscriptionsAPI.md#GetIdentitiesByPiidSubscriptions) | **Get** /v2/identities/{piid}/subscriptions | Get a platform identity&#39;s channel subscriptions (v2)
[**GetPublicChannelsByIdSubscriptions**](SubscriptionsAPI.md#GetPublicChannelsByIdSubscriptions) | **Get** /v2/public/channels/{id}/subscriptions | Get channel subscribers
[**GetSubscriptionsByIdLogs**](SubscriptionsAPI.md#GetSubscriptionsByIdLogs) | **Get** /v2/subscriptions/{id}/logs | Get the participation log for a subscription
[**GetSubscriptionsSearchKeywordByKeyword**](SubscriptionsAPI.md#GetSubscriptionsSearchKeywordByKeyword) | **Get** /v2/subscriptions/searchKeyword/{keyword} | Search subscriptions by keyword (cursor-paginated)
[**UpdateSubscriptionsByIdStatus**](SubscriptionsAPI.md#UpdateSubscriptionsByIdStatus) | **Put** /v2/subscriptions/{id}/status | Approve or deny subscription



## CreateSubscriptions

> DomainSubscription CreateSubscriptions(ctx).RequestCreateSubscriptionRequest(requestCreateSubscriptionRequest).Execute()

Subscribe the caller to a channel

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
	requestCreateSubscriptionRequest := *openapiclient.NewRequestCreateSubscriptionRequest("ChannelId_example") // RequestCreateSubscriptionRequest | Subscription payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CreateSubscriptions(context.Background()).RequestCreateSubscriptionRequest(requestCreateSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CreateSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscriptions`: DomainSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CreateSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSubscriptionRequest** | [**RequestCreateSubscriptionRequest**](RequestCreateSubscriptionRequest.md) | Subscription payload | 

### Return type

[**DomainSubscription**](DomainSubscription.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionsById

> DeleteSubscriptionsById(ctx, id).Execute()

Unsubscribe from Channel



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
	id := "id_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.DeleteSubscriptionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.DeleteSubscriptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionsByIdRequest struct via the builder pattern


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


## GetChannelsByIdSubscriptions

> []DomainSubscription GetChannelsByIdSubscriptions(ctx, id).Execute()

Get channel subscribers



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
	resp, r, err := apiClient.SubscriptionsAPI.GetChannelsByIdSubscriptions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetChannelsByIdSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdSubscriptions`: []DomainSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetChannelsByIdSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsByIdSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainSubscription**](DomainSubscription.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidSubscriptions

> []DomainSubscription GetIdentitiesByPiidSubscriptions(ctx, piid).Cursor(cursor).Limit(limit).Execute()

Get a platform identity's channel subscriptions (v2)



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
	piid := "piid_example" // string | Platform Identity ID
	cursor := "cursor_example" // string | Cursor for pagination (optional)
	limit := int32(56) // int32 | Items per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetIdentitiesByPiidSubscriptions(context.Background(), piid).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetIdentitiesByPiidSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidSubscriptions`: []DomainSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetIdentitiesByPiidSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Cursor for pagination | 
 **limit** | **int32** | Items per page | 

### Return type

[**[]DomainSubscription**](DomainSubscription.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicChannelsByIdSubscriptions

> []DomainSubscription GetPublicChannelsByIdSubscriptions(ctx, id).Execute()

Get channel subscribers



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
	resp, r, err := apiClient.SubscriptionsAPI.GetPublicChannelsByIdSubscriptions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetPublicChannelsByIdSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicChannelsByIdSubscriptions`: []DomainSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetPublicChannelsByIdSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicChannelsByIdSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainSubscription**](DomainSubscription.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionsByIdLogs

> []DomainParticipation GetSubscriptionsByIdLogs(ctx, id).Execute()

Get the participation log for a subscription

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
	id := "id_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetSubscriptionsByIdLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscriptionsByIdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionsByIdLogs`: []DomainParticipation
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscriptionsByIdLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsByIdLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainParticipation**](DomainParticipation.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionsSearchKeywordByKeyword

> []DomainSubscription GetSubscriptionsSearchKeywordByKeyword(ctx, keyword).Cursor(cursor).Limit(limit).Execute()

Search subscriptions by keyword (cursor-paginated)

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
	keyword := "keyword_example" // string | Search keyword
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.GetSubscriptionsSearchKeywordByKeyword(context.Background(), keyword).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscriptionsSearchKeywordByKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionsSearchKeywordByKeyword`: []DomainSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscriptionsSearchKeywordByKeyword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string** | Search keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsSearchKeywordByKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 

### Return type

[**[]DomainSubscription**](DomainSubscription.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscriptionsByIdStatus

> UpdateSubscriptionsByIdStatus(ctx, id).RequestUpdateSubscriptionStatusRequest(requestUpdateSubscriptionStatusRequest).Execute()

Approve or deny subscription



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
	id := "id_example" // string | Subscription ID
	requestUpdateSubscriptionStatusRequest := *openapiclient.NewRequestUpdateSubscriptionStatusRequest("Status_example") // RequestUpdateSubscriptionStatusRequest | request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionsAPI.UpdateSubscriptionsByIdStatus(context.Background(), id).RequestUpdateSubscriptionStatusRequest(requestUpdateSubscriptionStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UpdateSubscriptionsByIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionsByIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateSubscriptionStatusRequest** | [**RequestUpdateSubscriptionStatusRequest**](RequestUpdateSubscriptionStatusRequest.md) | request | 

### Return type

 (empty response body)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

