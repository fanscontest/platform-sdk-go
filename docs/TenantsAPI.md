# \TenantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenants**](TenantsAPI.md#CreateTenants) | **Post** /v2/tenants | Apply to become a tenant (vetted-signup)
[**CreateTenantsByIdApiKeys**](TenantsAPI.md#CreateTenantsByIdApiKeys) | **Post** /v2/tenants/{id}/api-keys | Mint an API key for an approved tenant
[**CreateTenantsByIdWebhookSubscriptions**](TenantsAPI.md#CreateTenantsByIdWebhookSubscriptions) | **Post** /v2/tenants/{id}/webhook-subscriptions | Register a webhook subscription for a tenant
[**DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId**](TenantsAPI.md#DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId) | **Delete** /v2/tenants/{id}/webhook-subscriptions/{subscriptionId} | Delete a webhook subscription
[**GetTenantsByIdWebhookSubscriptions**](TenantsAPI.md#GetTenantsByIdWebhookSubscriptions) | **Get** /v2/tenants/{id}/webhook-subscriptions | List a tenant&#39;s webhook subscriptions



## CreateTenants

> DomainTenantResponse CreateTenants(ctx).HandlerCreatePlatformRequest(handlerCreatePlatformRequest).XActingAs(xActingAs).Execute()

Apply to become a tenant (vetted-signup)



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
	handlerCreatePlatformRequest := *openapiclient.NewHandlerCreatePlatformRequest("Name_example") // HandlerCreatePlatformRequest | Tenant application payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.CreateTenants(context.Background()).HandlerCreatePlatformRequest(handlerCreatePlatformRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenants`: DomainTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlerCreatePlatformRequest** | [**HandlerCreatePlatformRequest**](HandlerCreatePlatformRequest.md) | Tenant application payload | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTenantResponse**](DomainTenantResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantsByIdApiKeys

> DomainTenantApiKeyResponse CreateTenantsByIdApiKeys(ctx, id).HandlerCreateApiKeyRequest(handlerCreateApiKeyRequest).XActingAs(xActingAs).Execute()

Mint an API key for an approved tenant



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
	id := "id_example" // string | Tenant ID
	handlerCreateApiKeyRequest := *openapiclient.NewHandlerCreateApiKeyRequest("Env_example") // HandlerCreateApiKeyRequest | API key request (env + name)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.CreateTenantsByIdApiKeys(context.Background(), id).HandlerCreateApiKeyRequest(handlerCreateApiKeyRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenantsByIdApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantsByIdApiKeys`: DomainTenantApiKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenantsByIdApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantsByIdApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerCreateApiKeyRequest** | [**HandlerCreateApiKeyRequest**](HandlerCreateApiKeyRequest.md) | API key request (env + name) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTenantApiKeyResponse**](DomainTenantApiKeyResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantsByIdWebhookSubscriptions

> DomainWebhookSubscriptionCreatedResponse CreateTenantsByIdWebhookSubscriptions(ctx, id).HandlerCreateWebhookSubscriptionRequest(handlerCreateWebhookSubscriptionRequest).XActingAs(xActingAs).Execute()

Register a webhook subscription for a tenant



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
	id := "id_example" // string | Tenant ID
	handlerCreateWebhookSubscriptionRequest := *openapiclient.NewHandlerCreateWebhookSubscriptionRequest([]string{"EventTypes_example"}, "Url_example") // HandlerCreateWebhookSubscriptionRequest | Subscription payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.CreateTenantsByIdWebhookSubscriptions(context.Background(), id).HandlerCreateWebhookSubscriptionRequest(handlerCreateWebhookSubscriptionRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenantsByIdWebhookSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantsByIdWebhookSubscriptions`: DomainWebhookSubscriptionCreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenantsByIdWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantsByIdWebhookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerCreateWebhookSubscriptionRequest** | [**HandlerCreateWebhookSubscriptionRequest**](HandlerCreateWebhookSubscriptionRequest.md) | Subscription payload | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainWebhookSubscriptionCreatedResponse**](DomainWebhookSubscriptionCreatedResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId

> DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId(ctx, id, subscriptionId).XActingAs(xActingAs).Execute()

Delete a webhook subscription



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
	id := "id_example" // string | Tenant ID
	subscriptionId := "subscriptionId_example" // string | Webhook Subscription ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantsAPI.DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId(context.Background(), id, subscriptionId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.DeleteTenantsByIdWebhookSubscriptionsBySubscriptionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID | 
**subscriptionId** | **string** | Webhook Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantsByIdWebhookSubscriptionsBySubscriptionIdRequest struct via the builder pattern


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


## GetTenantsByIdWebhookSubscriptions

> DomainWebhookSubscriptionListResponse GetTenantsByIdWebhookSubscriptions(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List a tenant's webhook subscriptions



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
	id := "id_example" // string | Tenant ID
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.GetTenantsByIdWebhookSubscriptions(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenantsByIdWebhookSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantsByIdWebhookSubscriptions`: DomainWebhookSubscriptionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenantsByIdWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsByIdWebhookSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainWebhookSubscriptionListResponse**](DomainWebhookSubscriptionListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

