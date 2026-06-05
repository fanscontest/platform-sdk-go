# \CreativeContestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdEntries**](CreativeContestsAPI.md#CreateContestsByIdEntries) | **Post** /v2/contests/{id}/entries | Submit entry for creative contest
[**CreateContestsByIdEntriesByEntryIdScore**](CreativeContestsAPI.md#CreateContestsByIdEntriesByEntryIdScore) | **Post** /v2/contests/{id}/entries/{entryId}/score | Score an entry
[**CreateContestsByIdForceProceed**](CreativeContestsAPI.md#CreateContestsByIdForceProceed) | **Post** /v2/contests/{id}/force-proceed | Force proceed from paused state
[**CreateContestsByIdJurorsInvite**](CreativeContestsAPI.md#CreateContestsByIdJurorsInvite) | **Post** /v2/contests/{id}/jurors/invite | Invite jurors to judge a contest
[**CreateJurorsInvitationsByInvitationIdAccept**](CreativeContestsAPI.md#CreateJurorsInvitationsByInvitationIdAccept) | **Post** /v2/jurors/invitations/{invitationId}/accept | Accept a juror invitation
[**CreateJurorsInvitationsByInvitationIdDecline**](CreativeContestsAPI.md#CreateJurorsInvitationsByInvitationIdDecline) | **Post** /v2/jurors/invitations/{invitationId}/decline | Decline a juror invitation
[**CreateJurorsOptIn**](CreativeContestsAPI.md#CreateJurorsOptIn) | **Post** /v2/jurors/opt-in | Opt into jury pool
[**DeleteContestsByIdEntries**](CreativeContestsAPI.md#DeleteContestsByIdEntries) | **Delete** /v2/contests/{id}/entries | Withdraw entry from creative contest
[**GetContestsByIdEntries**](CreativeContestsAPI.md#GetContestsByIdEntries) | **Get** /v2/contests/{id}/entries | Get entries for a contest
[**GetContestsByIdJurors**](CreativeContestsAPI.md#GetContestsByIdJurors) | **Get** /v2/contests/{id}/jurors | Get jurors for a contest
[**GetContestsByIdJurorsInvitations**](CreativeContestsAPI.md#GetContestsByIdJurorsInvitations) | **Get** /v2/contests/{id}/jurors/invitations | Get invitations for a contest
[**GetJurorsAssignments**](CreativeContestsAPI.md#GetJurorsAssignments) | **Get** /v2/jurors/assignments | Get juror assignments (cursor-paginated)
[**GetJurorsInvitations**](CreativeContestsAPI.md#GetJurorsInvitations) | **Get** /v2/jurors/invitations | Get my pending invitations
[**UpdateContestsByIdEntriesByEntryIdScore**](CreativeContestsAPI.md#UpdateContestsByIdEntriesByEntryIdScore) | **Put** /v2/contests/{id}/entries/{entryId}/score | Score an entry



## CreateContestsByIdEntries

> CreateContestsByIdEntries200Response CreateContestsByIdEntries(ctx, id).RequestSubmitEntryRequest(requestSubmitEntryRequest).XTenantUserId(xTenantUserId).Execute()

Submit entry for creative contest



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
	requestSubmitEntryRequest := *openapiclient.NewRequestSubmitEntryRequest("MediaType_example", "MediaUrl_example") // RequestSubmitEntryRequest | Entry submission
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdEntries(context.Background(), id).RequestSubmitEntryRequest(requestSubmitEntryRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntries`: CreateContestsByIdEntries200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestSubmitEntryRequest** | [**RequestSubmitEntryRequest**](RequestSubmitEntryRequest.md) | Entry submission | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateContestsByIdEntries200Response**](CreateContestsByIdEntries200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdEntriesByEntryIdScore

> UpdateContestsByIdEntriesByEntryIdScore200Response CreateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XTenantUserId(xTenantUserId).Execute()

Score an entry



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
	entryId := "entryId_example" // string | Entry ID
	requestScoreEntryRequest := *openapiclient.NewRequestScoreEntryRequest(int32(123)) // RequestScoreEntryRequest | Score
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntriesByEntryIdScore`: UpdateContestsByIdEntriesByEntryIdScore200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdEntriesByEntryIdScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestScoreEntryRequest** | [**RequestScoreEntryRequest**](RequestScoreEntryRequest.md) | Score | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**UpdateContestsByIdEntriesByEntryIdScore200Response**](UpdateContestsByIdEntriesByEntryIdScore200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdForceProceed

> CreateAnalyticsImpression200Response CreateContestsByIdForceProceed(ctx, id).XTenantUserId(xTenantUserId).Execute()

Force proceed from paused state



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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdForceProceed(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdForceProceed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdForceProceed`: CreateAnalyticsImpression200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdForceProceed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdForceProceedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateAnalyticsImpression200Response**](CreateAnalyticsImpression200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdJurorsInvite

> CreateAnalyticsImpression200Response CreateContestsByIdJurorsInvite(ctx, id).RequestInviteJurorsRequest(requestInviteJurorsRequest).XTenantUserId(xTenantUserId).Execute()

Invite jurors to judge a contest



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
	requestInviteJurorsRequest := *openapiclient.NewRequestInviteJurorsRequest([]string{"PlatformIdentityIds_example"}) // RequestInviteJurorsRequest | Member IDs to invite
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateContestsByIdJurorsInvite(context.Background(), id).RequestInviteJurorsRequest(requestInviteJurorsRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateContestsByIdJurorsInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdJurorsInvite`: CreateAnalyticsImpression200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateContestsByIdJurorsInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdJurorsInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInviteJurorsRequest** | [**RequestInviteJurorsRequest**](RequestInviteJurorsRequest.md) | Member IDs to invite | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateAnalyticsImpression200Response**](CreateAnalyticsImpression200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsInvitationsByInvitationIdAccept

> CreateJurorsInvitationsByInvitationIdAccept200Response CreateJurorsInvitationsByInvitationIdAccept(ctx, invitationId).XTenantUserId(xTenantUserId).Execute()

Accept a juror invitation



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
	invitationId := "invitationId_example" // string | Invitation ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept(context.Background(), invitationId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdAccept`: CreateJurorsInvitationsByInvitationIdAccept200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsInvitationsByInvitationIdAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateJurorsInvitationsByInvitationIdAccept200Response**](CreateJurorsInvitationsByInvitationIdAccept200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsInvitationsByInvitationIdDecline

> CreateAnalyticsImpression200Response CreateJurorsInvitationsByInvitationIdDecline(ctx, invitationId).XTenantUserId(xTenantUserId).Execute()

Decline a juror invitation



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
	invitationId := "invitationId_example" // string | Invitation ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline(context.Background(), invitationId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdDecline`: CreateAnalyticsImpression200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsInvitationsByInvitationIdDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsInvitationsByInvitationIdDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateAnalyticsImpression200Response**](CreateAnalyticsImpression200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsOptIn

> CreateJurorsOptIn200Response CreateJurorsOptIn(ctx).RequestOptInJurorRequest(requestOptInJurorRequest).XTenantUserId(xTenantUserId).Execute()

Opt into jury pool



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
	requestOptInJurorRequest := *openapiclient.NewRequestOptInJurorRequest([]string{"Categories_example"}) // RequestOptInJurorRequest | Categories
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.CreateJurorsOptIn(context.Background()).RequestOptInJurorRequest(requestOptInJurorRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.CreateJurorsOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsOptIn`: CreateJurorsOptIn200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.CreateJurorsOptIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestOptInJurorRequest** | [**RequestOptInJurorRequest**](RequestOptInJurorRequest.md) | Categories | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateJurorsOptIn200Response**](CreateJurorsOptIn200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsByIdEntries

> CreateAnalyticsImpression200Response DeleteContestsByIdEntries(ctx, id).XTenantUserId(xTenantUserId).Execute()

Withdraw entry from creative contest



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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.DeleteContestsByIdEntries(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.DeleteContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContestsByIdEntries`: CreateAnalyticsImpression200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.DeleteContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**CreateAnalyticsImpression200Response**](CreateAnalyticsImpression200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdEntries

> GetContestsByIdEntries200Response GetContestsByIdEntries(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get entries for a contest



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdEntries(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdEntries`: GetContestsByIdEntries200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByIdEntries200Response**](GetContestsByIdEntries200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurors

> GetContestsByIdJurors200Response GetContestsByIdJurors(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get jurors for a contest



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, derived) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, derived) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdJurors(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdJurors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurors`: GetContestsByIdJurors200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdJurors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdJurorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, derived) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, derived) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByIdJurors200Response**](GetContestsByIdJurors200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurorsInvitations

> GetContestsByIdJurorsInvitations200Response GetContestsByIdJurorsInvitations(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get invitations for a contest



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetContestsByIdJurorsInvitations(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetContestsByIdJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurorsInvitations`: GetContestsByIdJurorsInvitations200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetContestsByIdJurorsInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdJurorsInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByIdJurorsInvitations200Response**](GetContestsByIdJurorsInvitations200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsAssignments

> GetContestsByIdJurorsInvitations200Response GetJurorsAssignments(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get juror assignments (cursor-paginated)



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetJurorsAssignments(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetJurorsAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsAssignments`: GetContestsByIdJurorsInvitations200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetJurorsAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByIdJurorsInvitations200Response**](GetContestsByIdJurorsInvitations200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsInvitations

> GetContestsByIdJurorsInvitations200Response GetJurorsInvitations(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get my pending invitations



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page, service-assembled) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page, service-assembled) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.GetJurorsInvitations(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.GetJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsInvitations`: GetContestsByIdJurorsInvitations200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.GetJurorsInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page, service-assembled) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page, service-assembled) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**GetContestsByIdJurorsInvitations200Response**](GetContestsByIdJurorsInvitations200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdEntriesByEntryIdScore

> UpdateContestsByIdEntriesByEntryIdScore200Response UpdateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XTenantUserId(xTenantUserId).Execute()

Score an entry



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
	entryId := "entryId_example" // string | Entry ID
	requestScoreEntryRequest := *openapiclient.NewRequestScoreEntryRequest(int32(123)) // RequestScoreEntryRequest | Score
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdEntriesByEntryIdScore`: UpdateContestsByIdEntriesByEntryIdScore200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestsAPI.UpdateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContestsByIdEntriesByEntryIdScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestScoreEntryRequest** | [**RequestScoreEntryRequest**](RequestScoreEntryRequest.md) | Score | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**UpdateContestsByIdEntriesByEntryIdScore200Response**](UpdateContestsByIdEntriesByEntryIdScore200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

