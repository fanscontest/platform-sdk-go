# \ParticipationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContestsByIdEnrollment**](ParticipationAPI.md#CreateContestsByIdEnrollment) | **Post** /v2/contests/{id}/enrollment | Enroll in a Contest
[**CreateContestsByIdParticipate**](ParticipationAPI.md#CreateContestsByIdParticipate) | **Post** /v2/contests/{id}/participate | Create Participation
[**CreateParticipationByIdStart**](ParticipationAPI.md#CreateParticipationByIdStart) | **Post** /v2/participation/{id}/start | Log the start of a participation
[**DeleteContestsByIdEnrollment**](ParticipationAPI.md#DeleteContestsByIdEnrollment) | **Delete** /v2/contests/{id}/enrollment | Withdraw the caller&#39;s enrollment from a contest
[**GetContestsByIdScoreboard**](ParticipationAPI.md#GetContestsByIdScoreboard) | **Get** /v2/contests/{id}/scoreboard | Get a contest scoreboard
[**GetIdentitiesByPiidParticipation**](ParticipationAPI.md#GetIdentitiesByPiidParticipation) | **Get** /v2/identities/{piid}/participation | Get all participation for a platform identity (cursor-paginated)
[**GetIdentitiesByPiidSchedule**](ParticipationAPI.md#GetIdentitiesByPiidSchedule) | **Get** /v2/identities/{piid}/schedule | Get a platform identity&#39;s contest schedule (cached read model)
[**GetParticipationById**](ParticipationAPI.md#GetParticipationById) | **Get** /v2/participation/{id} | Get participation by id
[**GetParticipationContestById**](ParticipationAPI.md#GetParticipationContestById) | **Get** /v2/participation/contest/{id} | Get member participation by contest id
[**GetPublicContestsByIdGroupedScoreboard**](ParticipationAPI.md#GetPublicContestsByIdGroupedScoreboard) | **Get** /v2/public/contests/{id}/grouped-scoreboard | Get ScoreBoard
[**GetPublicContestsByIdLiveRanking**](ParticipationAPI.md#GetPublicContestsByIdLiveRanking) | **Get** /v2/public/contests/{id}/live-ranking | Get live ranking snapshot
[**GetPublicContestsByIdScoreboard**](ParticipationAPI.md#GetPublicContestsByIdScoreboard) | **Get** /v2/public/contests/{id}/scoreboard | Get a contest scoreboard
[**PatchParticipationByParticipationIdGroup**](ParticipationAPI.md#PatchParticipationByParticipationIdGroup) | **Patch** /v2/participation/{participationId}/group | Select group for open grouping contest



## CreateContestsByIdEnrollment

> DomainParticipationResponse CreateContestsByIdEnrollment(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdEnrollment(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> DomainParticipationResponse CreateContestsByIdParticipate(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdParticipate(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> HandlerStatusResponseResponse CreateParticipationByIdStart(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateParticipationByIdStart(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> DeleteContestsByIdEnrollment(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParticipationAPI.DeleteContestsByIdEnrollment(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetContestsByIdScoreboard

> DomainScoreBoardResponse GetContestsByIdScoreboard(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetContestsByIdScoreboard(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> DomainParticipationListResponse GetIdentitiesByPiidParticipation(ctx, piid).Limit(limit).Cursor(cursor).State(state).XTenantUserId(xTenantUserId).Execute()

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
	state := "state_example" // string | Contest state filter (created|started|closed|completed) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidParticipation(context.Background(), piid).Limit(limit).Cursor(cursor).State(state).XTenantUserId(xTenantUserId).Execute()
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
 **state** | **string** | Contest state filter (created|started|closed|completed) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> DomainParticipationListResponse GetIdentitiesByPiidSchedule(ctx, piid).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get a platform identity's contest schedule (cached read model)



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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page derived data) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page derived data) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidSchedule(context.Background(), piid).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetIdentitiesByPiidSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidSchedule`: DomainParticipationListResponse
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

 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page derived data) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page derived data) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetParticipationById

> DomainParticipationResponse GetParticipationById(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

> DomainParticipationResponse GetParticipationContestById(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationContestById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
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

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetPublicContestsByIdGroupedScoreboard

> DomainGroupedScoreBoardResponse GetPublicContestsByIdGroupedScoreboard(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdGroupedScoreboard(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdGroupedScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdGroupedScoreboard`: DomainGroupedScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetPublicContestsByIdGroupedScoreboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdGroupedScoreboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetPublicContestsByIdLiveRanking

> DomainGroupedScoreBoardResponse GetPublicContestsByIdLiveRanking(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdLiveRanking(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdLiveRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdLiveRanking`: DomainGroupedScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetPublicContestsByIdLiveRanking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdLiveRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetPublicContestsByIdScoreboard

> DomainScoreBoardResponse GetPublicContestsByIdScoreboard(ctx, id).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdScoreboard(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdScoreboard`: DomainScoreBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `ParticipationAPI.GetPublicContestsByIdScoreboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdScoreboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## PatchParticipationByParticipationIdGroup

> DomainSelectGroupResponseResponse PatchParticipationByParticipationIdGroup(ctx, participationId).RequestSelectGroupRequest(requestSelectGroupRequest).XTenantUserId(xTenantUserId).Execute()

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
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.PatchParticipationByParticipationIdGroup(context.Background(), participationId).RequestSelectGroupRequest(requestSelectGroupRequest).XTenantUserId(xTenantUserId).Execute()
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
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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

