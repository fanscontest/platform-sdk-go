# \IdentitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentities**](IdentitiesAPI.md#CreateIdentities) | **Post** /v2/identities | Register a tenant-supplied end-user as a platform identity
[**DeleteIdentitiesById**](IdentitiesAPI.md#DeleteIdentitiesById) | **Delete** /v2/identities/{id} | Soft-delete a platform identity, scoped to the calling tenant
[**GetIdentities**](IdentitiesAPI.md#GetIdentities) | **Get** /v2/identities | Look up a platform identity by tenant_user_id (409 recovery)
[**GetIdentitiesById**](IdentitiesAPI.md#GetIdentitiesById) | **Get** /v2/identities/{id} | Get a platform identity by ID, scoped to the calling tenant
[**UpdateIdentitiesById**](IdentitiesAPI.md#UpdateIdentitiesById) | **Put** /v2/identities/{id} | Replace a platform identity&#39;s profile, scoped to the calling tenant



## CreateIdentities

> DomainPlatformIdentity CreateIdentities(ctx).HandlerCreatePlatformIdentityRequest(handlerCreatePlatformIdentityRequest).Execute()

Register a tenant-supplied end-user as a platform identity



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
	handlerCreatePlatformIdentityRequest := *openapiclient.NewHandlerCreatePlatformIdentityRequest("TenantUserId_example") // HandlerCreatePlatformIdentityRequest | Platform identity payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.CreateIdentities(context.Background()).HandlerCreatePlatformIdentityRequest(handlerCreatePlatformIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.CreateIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentities`: DomainPlatformIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.CreateIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handlerCreatePlatformIdentityRequest** | [**HandlerCreatePlatformIdentityRequest**](HandlerCreatePlatformIdentityRequest.md) | Platform identity payload | 

### Return type

[**DomainPlatformIdentity**](DomainPlatformIdentity.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentitiesById

> DeleteIdentitiesById(ctx, id).Execute()

Soft-delete a platform identity, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitiesAPI.DeleteIdentitiesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.DeleteIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitiesByIdRequest struct via the builder pattern


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


## GetIdentities

> DomainPlatformIdentity GetIdentities(ctx).TenantUserId(tenantUserId).Execute()

Look up a platform identity by tenant_user_id (409 recovery)



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
	tenantUserId := "tenantUserId_example" // string | Tenant-supplied user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentities(context.Background()).TenantUserId(tenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentities`: DomainPlatformIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantUserId** | **string** | Tenant-supplied user ID | 

### Return type

[**DomainPlatformIdentity**](DomainPlatformIdentity.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesById

> DomainPlatformIdentity GetIdentitiesById(ctx, id).Execute()

Get a platform identity by ID, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentitiesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesById`: DomainPlatformIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainPlatformIdentity**](DomainPlatformIdentity.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentitiesById

> DomainPlatformIdentity UpdateIdentitiesById(ctx, id).HandlerUpdatePlatformIdentityRequest(handlerUpdatePlatformIdentityRequest).Execute()

Replace a platform identity's profile, scoped to the calling tenant



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
	id := "id_example" // string | Platform Identity ID
	handlerUpdatePlatformIdentityRequest := *openapiclient.NewHandlerUpdatePlatformIdentityRequest("DisplayName_example") // HandlerUpdatePlatformIdentityRequest | Profile fields

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.UpdateIdentitiesById(context.Background(), id).HandlerUpdatePlatformIdentityRequest(handlerUpdatePlatformIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.UpdateIdentitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentitiesById`: DomainPlatformIdentity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.UpdateIdentitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerUpdatePlatformIdentityRequest** | [**HandlerUpdatePlatformIdentityRequest**](HandlerUpdatePlatformIdentityRequest.md) | Profile fields | 

### Return type

[**DomainPlatformIdentity**](DomainPlatformIdentity.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

