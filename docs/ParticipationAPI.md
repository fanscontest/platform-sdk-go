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

> DomainParticipation CreateContestsByIdEnrollment(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdEnrollment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateContestsByIdEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdEnrollment`: DomainParticipation
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


### Return type

[**DomainParticipation**](DomainParticipation.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdParticipate

> DomainParticipation CreateContestsByIdParticipate(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateContestsByIdParticipate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateContestsByIdParticipate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdParticipate`: DomainParticipation
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


### Return type

[**DomainParticipation**](DomainParticipation.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipationByIdStart

> HandlerStatusResponse CreateParticipationByIdStart(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.CreateParticipationByIdStart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.CreateParticipationByIdStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateParticipationByIdStart`: HandlerStatusResponse
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


## DeleteContestsByIdEnrollment

> DeleteContestsByIdEnrollment(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ParticipationAPI.DeleteContestsByIdEnrollment(context.Background(), id).Execute()
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

> DomainScoreBoard GetContestsByIdScoreboard(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetContestsByIdScoreboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetContestsByIdScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdScoreboard`: DomainScoreBoard
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


### Return type

[**DomainScoreBoard**](DomainScoreBoard.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidParticipation

> []DomainParticipation GetIdentitiesByPiidParticipation(ctx, piid).Limit(limit).Cursor(cursor).State(state).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidParticipation(context.Background(), piid).Limit(limit).Cursor(cursor).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetIdentitiesByPiidParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidParticipation`: []DomainParticipation
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


## GetIdentitiesByPiidSchedule

> []DomainParticipation GetIdentitiesByPiidSchedule(ctx, piid).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetIdentitiesByPiidSchedule(context.Background(), piid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetIdentitiesByPiidSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidSchedule`: []DomainParticipation
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


## GetParticipationById

> DomainParticipation GetParticipationById(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetParticipationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipationById`: DomainParticipation
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


### Return type

[**DomainParticipation**](DomainParticipation.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipationContestById

> DomainParticipation GetParticipationContestById(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetParticipationContestById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetParticipationContestById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipationContestById`: DomainParticipation
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


### Return type

[**DomainParticipation**](DomainParticipation.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByIdGroupedScoreboard

> DomainGroupedScoreBoard GetPublicContestsByIdGroupedScoreboard(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdGroupedScoreboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdGroupedScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdGroupedScoreboard`: DomainGroupedScoreBoard
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


### Return type

[**DomainGroupedScoreBoard**](DomainGroupedScoreBoard.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByIdLiveRanking

> DomainGroupedScoreBoard GetPublicContestsByIdLiveRanking(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdLiveRanking(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdLiveRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdLiveRanking`: DomainGroupedScoreBoard
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


### Return type

[**DomainGroupedScoreBoard**](DomainGroupedScoreBoard.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByIdScoreboard

> DomainScoreBoard GetPublicContestsByIdScoreboard(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.GetPublicContestsByIdScoreboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.GetPublicContestsByIdScoreboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdScoreboard`: DomainScoreBoard
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


### Return type

[**DomainScoreBoard**](DomainScoreBoard.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchParticipationByParticipationIdGroup

> DomainSelectGroupResponse PatchParticipationByParticipationIdGroup(ctx, participationId).RequestSelectGroupRequest(requestSelectGroupRequest).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipationAPI.PatchParticipationByParticipationIdGroup(context.Background(), participationId).RequestSelectGroupRequest(requestSelectGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipationAPI.PatchParticipationByParticipationIdGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchParticipationByParticipationIdGroup`: DomainSelectGroupResponse
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

### Return type

[**DomainSelectGroupResponse**](DomainSelectGroupResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

