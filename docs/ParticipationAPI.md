# \ParticipationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdEnrollment**](ParticipationAPI.md#CreateContestsByIdEnrollment) | **Post** /v2/contests/{id}/enrollment | Enroll in a Contest
[**CreateContestsByIdParticipate**](ParticipationAPI.md#CreateContestsByIdParticipate) | **Post** /v2/contests/{id}/participate | Create Participation
[**CreateParticipationByIdStart**](ParticipationAPI.md#CreateParticipationByIdStart) | **Post** /v2/participation/{id}/start | Log the start of a participation
[**DeleteContestsByIdEnrollment**](ParticipationAPI.md#DeleteContestsByIdEnrollment) | **Delete** /v2/contests/{id}/enrollment | Withdraw the caller&#39;s enrollment from a contest
[**GetContestsByIdGroupedScoreboard**](ParticipationAPI.md#GetContestsByIdGroupedScoreboard) | **Get** /v2/contests/{id}/grouped-scoreboard | Get ScoreBoard
[**GetContestsByIdLiveRanking**](ParticipationAPI.md#GetContestsByIdLiveRanking) | **Get** /v2/contests/{id}/live-ranking | Get live ranking snapshot
[**GetContestsByIdScoreboard**](ParticipationAPI.md#GetContestsByIdScoreboard) | **Get** /v2/contests/{id}/scoreboard | Get a contest scoreboard
[**GetIdentitiesByPiidParticipation**](ParticipationAPI.md#GetIdentitiesByPiidParticipation) | **Get** /v2/identities/{piid}/participation | Get all participation for a platform identity (cursor-paginated)
[**GetIdentitiesByPiidSchedule**](ParticipationAPI.md#GetIdentitiesByPiidSchedule) | **Get** /v2/identities/{piid}/schedule | Get a platform identity&#39;s contest schedule (upcoming, eligibility-sourced, cursor-paginated)
[**GetParticipationById**](ParticipationAPI.md#GetParticipationById) | **Get** /v2/participation/{id} | Get participation by id
[**GetParticipationContestById**](ParticipationAPI.md#GetParticipationContestById) | **Get** /v2/participation/contest/{id} | Get member participation by contest id
[**PatchParticipationByParticipationIdGroup**](ParticipationAPI.md#PatchParticipationByParticipationIdGroup) | **Patch** /v2/participation/{participationId}/group | Select group for open grouping contest



## CreateContestsByIdEnrollment

> DomainParticipationResponse CreateContestsByIdEnrollment(ctx, id).XActingAs(xActingAs).Execute()

Enroll in a Contest



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdEnrollment(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateContestsByIdEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEnrollment`: DomainParticipationResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.CreateContestsByIdEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationResponse**](DomainParticipationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdParticipate

> DomainParticipationResponse CreateContestsByIdParticipate(ctx, id).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()

Create Participation



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
	idempotencyKey := "idempotencyKey_example" // string | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdParticipate(context.Background(), id).IdempotencyKey(idempotencyKey).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateContestsByIdParticipate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdParticipate`: DomainParticipationResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.CreateContestsByIdParticipate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdParticipateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Retry key (~24h). Same key + same body replays the original response; a different body returns 422 (ADR 0055). | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationResponse**](DomainParticipationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipationByIdStart

> HandlerStatusResponseResponse CreateParticipationByIdStart(ctx, id).XActingAs(xActingAs).Execute()

Log the start of a participation

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
	id := "id_example" // string | Participation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateParticipationByIdStart(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateParticipationByIdStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateParticipationByIdStart`: HandlerStatusResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.CreateParticipationByIdStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Participation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateParticipationByIdStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsByIdEnrollment

> DeleteContestsByIdEnrollment(ctx, id).XActingAs(xActingAs).Execute()

Withdraw the caller's enrollment from a contest

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParticipationAPI.DeleteContestsByIdEnrollment(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.DeleteContestsByIdEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContestsByIdEnrollmentRequest struct via the builder pattern


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


## GetContestsByIdGroupedScoreboard

> DomainGroupedScoreBoardResponse GetContestsByIdGroupedScoreboard(ctx, id).XActingAs(xActingAs).Execute()

Get ScoreBoard



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetContestsByIdGroupedScoreboard(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetContestsByIdGroupedScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdGroupedScoreboard`: DomainGroupedScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetContestsByIdGroupedScoreboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdGroupedScoreboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainGroupedScoreBoardResponse**](DomainGroupedScoreBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdLiveRanking

> DomainGroupedScoreBoardResponse GetContestsByIdLiveRanking(ctx, id).XActingAs(xActingAs).Execute()

Get live ranking snapshot



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetContestsByIdLiveRanking(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetContestsByIdLiveRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdLiveRanking`: DomainGroupedScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetContestsByIdLiveRanking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdLiveRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainGroupedScoreBoardResponse**](DomainGroupedScoreBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdScoreboard

> DomainScoreBoardResponse GetContestsByIdScoreboard(ctx, id).XActingAs(xActingAs).Execute()

Get a contest scoreboard

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetContestsByIdScoreboard(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetContestsByIdScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdScoreboard`: DomainScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetContestsByIdScoreboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdScoreboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainScoreBoardResponse**](DomainScoreBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidParticipation

> DomainParticipationListResponse GetIdentitiesByPiidParticipation(ctx, piid).Limit(limit).Cursor(cursor).Status(status).XActingAs(xActingAs).Execute()

Get all participation for a platform identity (cursor-paginated)



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
	limit := int32(56) // int32 | Page size (1-200, default 50) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	status := "status_example" // string | Participation status filter, comma-separated (enrolled|participated|timed_out); default all. The schedule history view passes participated,timed_out. (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidParticipation(context.Background(), piid).Limit(limit).Cursor(cursor).Status(status).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetIdentitiesByPiidParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidParticipation`: DomainParticipationListResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetIdentitiesByPiidParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size (1-200, default 50) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **status** | **string** | Participation status filter, comma-separated (enrolled|participated|timed_out); default all. The schedule history view passes participated,timed_out. | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationListResponse**](DomainParticipationListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidSchedule

> DomainContestListResponse GetIdentitiesByPiidSchedule(ctx, piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

Get a platform identity's contest schedule (upcoming, eligibility-sourced, cursor-paginated)



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (1-200, default 50) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidSchedule(context.Background(), piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetIdentitiesByPiidSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidSchedule`: DomainContestListResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetIdentitiesByPiidSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (1-200, default 50) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestListResponse**](DomainContestListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipationById

> DomainParticipationResponse GetParticipationById(ctx, id).XActingAs(xActingAs).Execute()

Get participation by id



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
	id := "id_example" // string | Participation ID
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetParticipationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipationById`: DomainParticipationResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetParticipationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Participation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationResponse**](DomainParticipationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipationContestById

> DomainParticipationResponse GetParticipationContestById(ctx, id).XActingAs(xActingAs).Execute()

Get member participation by contest id



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationContestById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetParticipationContestById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipationContestById`: DomainParticipationResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetParticipationContestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipationContestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainParticipationResponse**](DomainParticipationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchParticipationByParticipationIdGroup

> DomainSelectGroupResponseResponse PatchParticipationByParticipationIdGroup(ctx, participationId).RequestSelectGroupRequest(requestSelectGroupRequest).XActingAs(xActingAs).Execute()

Select group for open grouping contest



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
	requestSelectGroupRequest := *openapiclient.NewRequestSelectGroupRequest("Group_example") // RequestSelectGroupRequest | Select Group Request
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.PatchParticipationByParticipationIdGroup(context.Background(), participationId).RequestSelectGroupRequest(requestSelectGroupRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.PatchParticipationByParticipationIdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchParticipationByParticipationIdGroup`: DomainSelectGroupResponseResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.PatchParticipationByParticipationIdGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**participationId** | **string** | Participation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchParticipationByParticipationIdGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestSelectGroupRequest** | [**RequestSelectGroupRequest**](RequestSelectGroupRequest.md) | Select Group Request | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainSelectGroupResponseResponse**](DomainSelectGroupResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

