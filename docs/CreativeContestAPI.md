# \CreativeContestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdEntries**](CreativeContestAPI.md#CreateContestsByIdEntries) | **Post** /v2/contests/{id}/entries | Submit entry for creative contest
[**CreateContestsByIdEntriesByEntryIdScore**](CreativeContestAPI.md#CreateContestsByIdEntriesByEntryIdScore) | **Post** /v2/contests/{id}/entries/{entryId}/score | Score an entry
[**CreateContestsByIdForceProceed**](CreativeContestAPI.md#CreateContestsByIdForceProceed) | **Post** /v2/contests/{id}/force-proceed | Force proceed from paused state
[**CreateContestsByIdJurorsInvite**](CreativeContestAPI.md#CreateContestsByIdJurorsInvite) | **Post** /v2/contests/{id}/jurors/invite | Invite jurors to judge a contest
[**CreateJurorsInvitationsByInvitationIdAccept**](CreativeContestAPI.md#CreateJurorsInvitationsByInvitationIdAccept) | **Post** /v2/jurors/invitations/{invitationId}/accept | Accept a juror invitation
[**CreateJurorsInvitationsByInvitationIdDecline**](CreativeContestAPI.md#CreateJurorsInvitationsByInvitationIdDecline) | **Post** /v2/jurors/invitations/{invitationId}/decline | Decline a juror invitation
[**CreateJurorsOptIn**](CreativeContestAPI.md#CreateJurorsOptIn) | **Post** /v2/jurors/opt-in | Opt into jury pool
[**DeleteContestsByIdEntries**](CreativeContestAPI.md#DeleteContestsByIdEntries) | **Delete** /v2/contests/{id}/entries | Withdraw entry from creative contest
[**GetContestsByIdEntries**](CreativeContestAPI.md#GetContestsByIdEntries) | **Get** /v2/contests/{id}/entries | Get entries for a contest
[**GetContestsByIdJurors**](CreativeContestAPI.md#GetContestsByIdJurors) | **Get** /v2/contests/{id}/jurors | Get jurors for a contest
[**GetContestsByIdJurorsInvitations**](CreativeContestAPI.md#GetContestsByIdJurorsInvitations) | **Get** /v2/contests/{id}/jurors/invitations | Get invitations for a contest
[**GetJurorsAssignments**](CreativeContestAPI.md#GetJurorsAssignments) | **Get** /v2/jurors/assignments | Get juror assignments
[**GetJurorsInvitations**](CreativeContestAPI.md#GetJurorsInvitations) | **Get** /v2/jurors/invitations | Get my pending invitations
[**UpdateContestsByIdEntriesByEntryIdScore**](CreativeContestAPI.md#UpdateContestsByIdEntriesByEntryIdScore) | **Put** /v2/contests/{id}/entries/{entryId}/score | Score an entry



## CreateContestsByIdEntries

> DomainEntry CreateContestsByIdEntries(ctx, id).RequestSubmitEntryRequest(requestSubmitEntryRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateContestsByIdEntries(context.Background(), id).RequestSubmitEntryRequest(requestSubmitEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntries`: DomainEntry
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateContestsByIdEntries`: %v\n", resp)
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

### Return type

[**DomainEntry**](DomainEntry.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdEntriesByEntryIdScore

> DomainJurorScore CreateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEntriesByEntryIdScore`: DomainJurorScore
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
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

### Return type

[**DomainJurorScore**](DomainJurorScore.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdForceProceed

> HandlerStatusResponse CreateContestsByIdForceProceed(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateContestsByIdForceProceed(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateContestsByIdForceProceed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdForceProceed`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateContestsByIdForceProceed`: %v\n", resp)
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


### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdJurorsInvite

> HandlerStatusResponse CreateContestsByIdJurorsInvite(ctx, id).RequestInviteJurorsRequest(requestInviteJurorsRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateContestsByIdJurorsInvite(context.Background(), id).RequestInviteJurorsRequest(requestInviteJurorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateContestsByIdJurorsInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdJurorsInvite`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateContestsByIdJurorsInvite`: %v\n", resp)
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


## CreateJurorsInvitationsByInvitationIdAccept

> DomainJurorAssignment CreateJurorsInvitationsByInvitationIdAccept(ctx, invitationId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateJurorsInvitationsByInvitationIdAccept(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateJurorsInvitationsByInvitationIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdAccept`: DomainJurorAssignment
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateJurorsInvitationsByInvitationIdAccept`: %v\n", resp)
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


### Return type

[**DomainJurorAssignment**](DomainJurorAssignment.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsInvitationsByInvitationIdDecline

> HandlerStatusResponse CreateJurorsInvitationsByInvitationIdDecline(ctx, invitationId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateJurorsInvitationsByInvitationIdDecline(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateJurorsInvitationsByInvitationIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsInvitationsByInvitationIdDecline`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateJurorsInvitationsByInvitationIdDecline`: %v\n", resp)
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


### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJurorsOptIn

> DomainJuror CreateJurorsOptIn(ctx).RequestOptInJurorRequest(requestOptInJurorRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.CreateJurorsOptIn(context.Background()).RequestOptInJurorRequest(requestOptInJurorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.CreateJurorsOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJurorsOptIn`: DomainJuror
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.CreateJurorsOptIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJurorsOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestOptInJurorRequest** | [**RequestOptInJurorRequest**](RequestOptInJurorRequest.md) | Categories | 

### Return type

[**DomainJuror**](DomainJuror.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsByIdEntries

> HandlerStatusResponse DeleteContestsByIdEntries(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.DeleteContestsByIdEntries(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.DeleteContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContestsByIdEntries`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.DeleteContestsByIdEntries`: %v\n", resp)
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


### Return type

[**HandlerStatusResponse**](HandlerStatusResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdEntries

> []DomainEntry GetContestsByIdEntries(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.GetContestsByIdEntries(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.GetContestsByIdEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdEntries`: []DomainEntry
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.GetContestsByIdEntries`: %v\n", resp)
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


### Return type

[**[]DomainEntry**](DomainEntry.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurors

> []DomainJuror GetContestsByIdJurors(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.GetContestsByIdJurors(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.GetContestsByIdJurors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurors`: []DomainJuror
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.GetContestsByIdJurors`: %v\n", resp)
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


### Return type

[**[]DomainJuror**](DomainJuror.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdJurorsInvitations

> []DomainJurorAssignment GetContestsByIdJurorsInvitations(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.GetContestsByIdJurorsInvitations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.GetContestsByIdJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdJurorsInvitations`: []DomainJurorAssignment
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.GetContestsByIdJurorsInvitations`: %v\n", resp)
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


### Return type

[**[]DomainJurorAssignment**](DomainJurorAssignment.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsAssignments

> []DomainJurorAssignment GetJurorsAssignments(ctx).Execute()

Get juror assignments



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.GetJurorsAssignments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.GetJurorsAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsAssignments`: []DomainJurorAssignment
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.GetJurorsAssignments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsAssignmentsRequest struct via the builder pattern


### Return type

[**[]DomainJurorAssignment**](DomainJurorAssignment.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJurorsInvitations

> []DomainJurorAssignment GetJurorsInvitations(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.GetJurorsInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.GetJurorsInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJurorsInvitations`: []DomainJurorAssignment
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.GetJurorsInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJurorsInvitationsRequest struct via the builder pattern


### Return type

[**[]DomainJurorAssignment**](DomainJurorAssignment.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdEntriesByEntryIdScore

> DomainJurorScore UpdateContestsByIdEntriesByEntryIdScore(ctx, id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeContestAPI.UpdateContestsByIdEntriesByEntryIdScore(context.Background(), id, entryId).RequestScoreEntryRequest(requestScoreEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeContestAPI.UpdateContestsByIdEntriesByEntryIdScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdEntriesByEntryIdScore`: DomainJurorScore
	fmt.Fprintf(os.Stdout, "Response from `CreativeContestAPI.UpdateContestsByIdEntriesByEntryIdScore`: %v\n", resp)
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

### Return type

[**DomainJurorScore**](DomainJurorScore.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

