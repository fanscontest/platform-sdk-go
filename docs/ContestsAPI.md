# \ContestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContests**](ContestsAPI.md#CreateContests) | **Post** /v2/contests | Create a contest
[**CreateContestsByContestIdBuddyBoards**](ContestsAPI.md#CreateContestsByContestIdBuddyBoards) | **Post** /v2/contests/{contestId}/buddy-boards | Create a buddy board on a contest (owner &#x3D; caller)
[**CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept**](ContestsAPI.md#CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept) | **Post** /v2/contests/{contestId}/buddy-boards/{boardId}/invites/accept | Accept a pending buddy board invite (invitee only)
[**CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline**](ContestsAPI.md#CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline) | **Post** /v2/contests/{contestId}/buddy-boards/{boardId}/invites/decline | Decline a pending buddy board invite (invitee only)
[**CreateContestsByContestIdBuddyBoardsByBoardIdMembers**](ContestsAPI.md#CreateContestsByContestIdBuddyBoardsByBoardIdMembers) | **Post** /v2/contests/{contestId}/buddy-boards/{boardId}/members | Invite members to a buddy board (owner only)
[**CreateContestsByIdHostings**](ContestsAPI.md#CreateContestsByIdHostings) | **Post** /v2/contests/{id}/hostings | Invite another venue/channel to host a contest
[**CreateContestsByIdHostingsByHostingIdAccept**](ContestsAPI.md#CreateContestsByIdHostingsByHostingIdAccept) | **Post** /v2/contests/{id}/hostings/{hostingId}/accept | Accept a contest hosting invitation
[**CreateContestsByIdHostingsByHostingIdDecline**](ContestsAPI.md#CreateContestsByIdHostingsByHostingIdDecline) | **Post** /v2/contests/{id}/hostings/{hostingId}/decline | Decline a contest hosting invitation
[**CreateContestsByIdParticipations**](ContestsAPI.md#CreateContestsByIdParticipations) | **Post** /v2/contests/{id}/participations | Submit a participation on behalf of a tenant-side fan
[**CreateTournamentsByIdContest**](ContestsAPI.md#CreateTournamentsByIdContest) | **Post** /v2/tournaments/{id}/contest | Create a contest for the next pending round in a tournament
[**DeleteContestsById**](ContestsAPI.md#DeleteContestsById) | **Delete** /v2/contests/{id} | Delete Contest
[**GetChannelsByIdContests**](ContestsAPI.md#GetChannelsByIdContests) | **Get** /v2/channels/{id}/contests | Get channel contests (cursor-paginated)
[**GetContests**](ContestsAPI.md#GetContests) | **Get** /v2/contests | List the caller&#39;s contests (cursor-paginated)
[**GetContestsByContestIdBuddyBoardsByBoardId**](ContestsAPI.md#GetContestsByContestIdBuddyBoardsByBoardId) | **Get** /v2/contests/{contestId}/buddy-boards/{boardId} | Get buddy board view (metadata + scores filtered to board members)
[**GetContestsByContestIdBuddyBoardsByBoardIdMembers**](ContestsAPI.md#GetContestsByContestIdBuddyBoardsByBoardIdMembers) | **Get** /v2/contests/{contestId}/buddy-boards/{boardId}/members | List buddy board members (public or authed)
[**GetContestsByContestIdParticipantsByMemberIdResult**](ContestsAPI.md#GetContestsByContestIdParticipantsByMemberIdResult) | **Get** /v2/contests/{contestId}/participants/{memberId}/result | Get participant result for a contest
[**GetContestsById**](ContestsAPI.md#GetContestsById) | **Get** /v2/contests/{id} | Get a contest by ID
[**GetContestsByIdHostings**](ContestsAPI.md#GetContestsByIdHostings) | **Get** /v2/contests/{id}/hostings | List hostings for a contest the calling tenant owns
[**GetContestsByIdParticipation**](ContestsAPI.md#GetContestsByIdParticipation) | **Get** /v2/contests/{id}/participation | Get contest participation result (cursor-paginated)
[**GetContestsByIdQuestionStats**](ContestsAPI.md#GetContestsByIdQuestionStats) | **Get** /v2/contests/{id}/question-stats | Get question statistics for a contest (Contest Owner or Admin Only)
[**GetContestsSearch**](ContestsAPI.md#GetContestsSearch) | **Get** /v2/contests/search | Search contests not owned by the caller (cursor-paginated)
[**GetIdentitiesByPiidBuddyBoardInvites**](ContestsAPI.md#GetIdentitiesByPiidBuddyBoardInvites) | **Get** /v2/identities/{piid}/buddy-board-invites | List buddy board invites for a platform identity
[**GetIdentitiesByPiidContests**](ContestsAPI.md#GetIdentitiesByPiidContests) | **Get** /v2/identities/{piid}/contests | List contests for a platform identity (cursor-paginated)
[**GetPublicChannelsByIdContests**](ContestsAPI.md#GetPublicChannelsByIdContests) | **Get** /v2/public/channels/{id}/contests | Get channel contests (cursor-paginated)
[**GetPublicContests**](ContestsAPI.md#GetPublicContests) | **Get** /v2/public/contests | List public contests (region-scoped, cursor-paginated)
[**GetPublicContestsByContestIdBuddyBoardsByBoardId**](ContestsAPI.md#GetPublicContestsByContestIdBuddyBoardsByBoardId) | **Get** /v2/public/contests/{contestId}/buddy-boards/{boardId} | Get buddy board view (metadata + scores filtered to board members)
[**GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers**](ContestsAPI.md#GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers) | **Get** /v2/public/contests/{contestId}/buddy-boards/{boardId}/members | List buddy board members (public or authed)
[**GetPublicContestsById**](ContestsAPI.md#GetPublicContestsById) | **Get** /v2/public/contests/{id} | Get a contest by ID
[**GetPublicContestsByIdGroups**](ContestsAPI.md#GetPublicContestsByIdGroups) | **Get** /v2/public/contests/{id}/groups | Get Contest Groups
[**GetPublicContestsByIdSubmissionFeed**](ContestsAPI.md#GetPublicContestsByIdSubmissionFeed) | **Get** /v2/public/contests/{id}/submission-feed | Get the real-time feed of contest submissions (cursor-paginated)
[**GetPublicContestsSearchKeywordByKeyword**](ContestsAPI.md#GetPublicContestsSearchKeywordByKeyword) | **Get** /v2/public/contests/searchKeyword/{keyword} | Search public contests by keyword (cursor-paginated)
[**PatchContestsById**](ContestsAPI.md#PatchContestsById) | **Patch** /v2/contests/{id} | Update a contest
[**UpdateContestsByIdHeaderImage**](ContestsAPI.md#UpdateContestsByIdHeaderImage) | **Put** /v2/contests/{id}/header-image | Update a contest&#39;s header image



## CreateContests

> DomainContest CreateContests(ctx).RequestCreateContestRequest(requestCreateContestRequest).XTenantUserId(xTenantUserId).Execute()

Create a contest

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
	requestCreateContestRequest := *openapiclient.NewRequestCreateContestRequest(int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example"), "SourceId_example", "SourceType_example", *openapiclient.NewDomainTiming("EndAt_example", "EnrollBy_example", "StartAt_example"), "Title_example", "VenueId_example", int32(123), int32(123)) // RequestCreateContestRequest | Contest payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContests(context.Background()).RequestCreateContestRequest(requestCreateContestRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContests`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateContestRequest** | [**RequestCreateContestRequest**](RequestCreateContestRequest.md) | Contest payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByContestIdBuddyBoards

> DomainBuddyBoard CreateContestsByContestIdBuddyBoards(ctx, contestId).RequestCreateBuddyBoardRequest(requestCreateBuddyBoardRequest).XTenantUserId(xTenantUserId).Execute()

Create a buddy board on a contest (owner = caller)

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
	requestCreateBuddyBoardRequest := *openapiclient.NewRequestCreateBuddyBoardRequest(int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example")) // RequestCreateBuddyBoardRequest | Buddy board payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoards(context.Background(), contestId).RequestCreateBuddyBoardRequest(requestCreateBuddyBoardRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoards`: DomainBuddyBoard
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByContestIdBuddyBoards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByContestIdBuddyBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestCreateBuddyBoardRequest** | [**RequestCreateBuddyBoardRequest**](RequestCreateBuddyBoardRequest.md) | Buddy board payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoard**](DomainBuddyBoard.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept

> HandlerStatusResponse CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

Accept a pending buddy board invite (invitee only)



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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByContestIdBuddyBoardsByBoardIdInvitesAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline

> HandlerStatusResponse CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

Decline a pending buddy board invite (invitee only)

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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByContestIdBuddyBoardsByBoardIdInvitesDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByContestIdBuddyBoardsByBoardIdMembers

> HandlerStatusResponse CreateContestsByContestIdBuddyBoardsByBoardIdMembers(ctx, contestId, boardId).RequestAddBuddyBoardMembersRequest(requestAddBuddyBoardMembersRequest).XTenantUserId(xTenantUserId).Execute()

Invite members to a buddy board (owner only)



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
	boardId := "boardId_example" // string | Buddy Board ID
	requestAddBuddyBoardMembersRequest := *openapiclient.NewRequestAddBuddyBoardMembersRequest([]string{"PlatformIdentityIds_example"}) // RequestAddBuddyBoardMembersRequest | Invitee IDs
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdMembers(context.Background(), contestId, boardId).RequestAddBuddyBoardMembersRequest(requestAddBuddyBoardMembersRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdMembers`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByContestIdBuddyBoardsByBoardIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestAddBuddyBoardMembersRequest** | [**RequestAddBuddyBoardMembersRequest**](RequestAddBuddyBoardMembersRequest.md) | Invitee IDs | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByIdHostings

> HandlerStatusResponse CreateContestsByIdHostings(ctx, id).HandlerCreatePlatformContestHostingRequest(handlerCreatePlatformContestHostingRequest).XTenantUserId(xTenantUserId).Execute()

Invite another venue/channel to host a contest



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
	handlerCreatePlatformContestHostingRequest := *openapiclient.NewHandlerCreatePlatformContestHostingRequest("ChannelId_example", "VenueId_example") // HandlerCreatePlatformContestHostingRequest | Hosting payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostings(context.Background(), id).HandlerCreatePlatformContestHostingRequest(handlerCreatePlatformContestHostingRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostings`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByIdHostings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdHostingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerCreatePlatformContestHostingRequest** | [**HandlerCreatePlatformContestHostingRequest**](HandlerCreatePlatformContestHostingRequest.md) | Hosting payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByIdHostingsByHostingIdAccept

> HandlerStatusResponse CreateContestsByIdHostingsByHostingIdAccept(ctx, id, hostingId).XTenantUserId(xTenantUserId).Execute()

Accept a contest hosting invitation



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
	hostingId := "hostingId_example" // string | Contest Hosting ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostingsByHostingIdAccept(context.Background(), id, hostingId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostingsByHostingIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostingsByHostingIdAccept`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByIdHostingsByHostingIdAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**hostingId** | **string** | Contest Hosting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdHostingsByHostingIdAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByIdHostingsByHostingIdDecline

> HandlerStatusResponse CreateContestsByIdHostingsByHostingIdDecline(ctx, id, hostingId).XTenantUserId(xTenantUserId).Execute()

Decline a contest hosting invitation



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
	hostingId := "hostingId_example" // string | Contest Hosting ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostingsByHostingIdDecline(context.Background(), id, hostingId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostingsByHostingIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostingsByHostingIdDecline`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByIdHostingsByHostingIdDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 
**hostingId** | **string** | Contest Hosting ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdHostingsByHostingIdDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateContestsByIdParticipations

> HandlerStatusResponse CreateContestsByIdParticipations(ctx, id).HandlerCreatePlatformParticipationRequest(handlerCreatePlatformParticipationRequest).XTenantUserId(xTenantUserId).Execute()

Submit a participation on behalf of a tenant-side fan



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
	handlerCreatePlatformParticipationRequest := *openapiclient.NewHandlerCreatePlatformParticipationRequest("ChannelId_example", "DisplayName_example", "TenantUserId_example", "VenueId_example") // HandlerCreatePlatformParticipationRequest | Participation payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdParticipations(context.Background(), id).HandlerCreatePlatformParticipationRequest(handlerCreatePlatformParticipationRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdParticipations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdParticipations`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContestsByIdParticipations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsByIdParticipationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **handlerCreatePlatformParticipationRequest** | [**HandlerCreatePlatformParticipationRequest**](HandlerCreatePlatformParticipationRequest.md) | Participation payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## CreateTournamentsByIdContest

> DomainContest CreateTournamentsByIdContest(ctx, id).RequestCreateContestRequest(requestCreateContestRequest).XTenantUserId(xTenantUserId).Execute()

Create a contest for the next pending round in a tournament

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
	id := "id_example" // string | Tournament ID
	requestCreateContestRequest := *openapiclient.NewRequestCreateContestRequest(int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example"), "SourceId_example", "SourceType_example", *openapiclient.NewDomainTiming("EndAt_example", "EnrollBy_example", "StartAt_example"), "Title_example", "VenueId_example", int32(123), int32(123)) // RequestCreateContestRequest | Contest payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateTournamentsByIdContest(context.Background(), id).RequestCreateContestRequest(requestCreateContestRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateTournamentsByIdContest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTournamentsByIdContest`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateTournamentsByIdContest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTournamentsByIdContestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestCreateContestRequest** | [**RequestCreateContestRequest**](RequestCreateContestRequest.md) | Contest payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsById

> string DeleteContestsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Delete Contest



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
	resp, r, err := apiClient.ContestsAPI.DeleteContestsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.DeleteContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContestsById`: string
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.DeleteContestsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContestsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

**string**

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdContests

> []DomainContest GetChannelsByIdContests(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get channel contests (cursor-paginated)



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 20) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetChannelsByIdContests(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetChannelsByIdContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdContests`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetChannelsByIdContests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsByIdContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 20) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContests

> []DomainContest GetContests(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List the caller's contests (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContests(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContests`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdBuddyBoardsByBoardId

> DomainBuddyBoardView GetContestsByContestIdBuddyBoardsByBoardId(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

Get buddy board view (metadata + scores filtered to board members)



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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardId(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdBuddyBoardsByBoardId`: DomainBuddyBoardView
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByContestIdBuddyBoardsByBoardIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardView**](DomainBuddyBoardView.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdBuddyBoardsByBoardIdMembers

> DomainBuddyBoardMembersResponse GetContestsByContestIdBuddyBoardsByBoardIdMembers(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

List buddy board members (public or authed)



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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardIdMembers(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdBuddyBoardsByBoardIdMembers`: DomainBuddyBoardMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByContestIdBuddyBoardsByBoardIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardMembersResponse**](DomainBuddyBoardMembersResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdParticipantsByMemberIdResult

> DomainMemberResult GetContestsByContestIdParticipantsByMemberIdResult(ctx, contestId, memberId).XTenantUserId(xTenantUserId).Execute()

Get participant result for a contest



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
	memberId := "memberId_example" // string | Member ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdParticipantsByMemberIdResult(context.Background(), contestId, memberId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdParticipantsByMemberIdResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdParticipantsByMemberIdResult`: DomainMemberResult
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByContestIdParticipantsByMemberIdResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**memberId** | **string** | Member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByContestIdParticipantsByMemberIdResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainMemberResult**](DomainMemberResult.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsById

> DomainContest GetContestsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get a contest by ID



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
	resp, r, err := apiClient.ContestsAPI.GetContestsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsById`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdHostings

> []DomainContestHosting GetContestsByIdHostings(ctx, id).XTenantUserId(xTenantUserId).Execute()

List hostings for a contest the calling tenant owns

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
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdHostings(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdHostings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdHostings`: []DomainContestHosting
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByIdHostings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdHostingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContestHosting**](DomainContestHosting.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdParticipation

> []DomainParticipation GetContestsByIdParticipation(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get contest participation result (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdParticipation(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdParticipation`: []DomainParticipation
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByIdParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

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


## GetContestsByIdQuestionStats

> []DomainQuestionStats GetContestsByIdQuestionStats(ctx, id).Sort(sort).XTenantUserId(xTenantUserId).Execute()

Get question statistics for a contest (Contest Owner or Admin Only)



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
	sort := "sort_example" // string | Sort order: failures, successes, failure_rate, success_rate (default: question number) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdQuestionStats(context.Background(), id).Sort(sort).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdQuestionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdQuestionStats`: []DomainQuestionStats
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByIdQuestionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdQuestionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort order: failures, successes, failure_rate, success_rate (default: question number) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainQuestionStats**](DomainQuestionStats.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsSearch

> []DomainContest GetContestsSearch(ctx).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Search contests not owned by the caller (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsSearch(context.Background()).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsSearch`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidBuddyBoardInvites

> DomainBuddyBoardInvitesResponse GetIdentitiesByPiidBuddyBoardInvites(ctx, piid).Status(status).XTenantUserId(xTenantUserId).Execute()

List buddy board invites for a platform identity



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
	status := "status_example" // string | Filter by invite status (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetIdentitiesByPiidBuddyBoardInvites(context.Background(), piid).Status(status).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetIdentitiesByPiidBuddyBoardInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidBuddyBoardInvites`: DomainBuddyBoardInvitesResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetIdentitiesByPiidBuddyBoardInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidBuddyBoardInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by invite status | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardInvitesResponse**](DomainBuddyBoardInvitesResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidContests

> []DomainContest GetIdentitiesByPiidContests(ctx, piid).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

List contests for a platform identity (cursor-paginated)



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
	limit := int32(56) // int32 | Page size (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetIdentitiesByPiidContests(context.Background(), piid).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetIdentitiesByPiidContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidContests`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetIdentitiesByPiidContests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicChannelsByIdContests

> []DomainContest GetPublicChannelsByIdContests(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get channel contests (cursor-paginated)



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 20) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicChannelsByIdContests(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicChannelsByIdContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicChannelsByIdContests`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicChannelsByIdContests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicChannelsByIdContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 20) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContests

> []DomainContest GetPublicContests(ctx).Cursor(cursor).Limit(limit).Ended(ended).Region(region).XTenantUserId(xTenantUserId).Execute()

List public contests (region-scoped, cursor-paginated)

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
	limit := int32(56) // int32 | Page size (default 50) (optional)
	ended := true // bool | Include ended contests (optional)
	region := "region_example" // string | ISO 3166-1 alpha-2 region filter (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicContests(context.Background()).Cursor(cursor).Limit(limit).Ended(ended).Region(region).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContests`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50) | 
 **ended** | **bool** | Include ended contests | 
 **region** | **string** | ISO 3166-1 alpha-2 region filter | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByContestIdBuddyBoardsByBoardId

> DomainBuddyBoardView GetPublicContestsByContestIdBuddyBoardsByBoardId(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

Get buddy board view (metadata + scores filtered to board members)



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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardId(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByContestIdBuddyBoardsByBoardId`: DomainBuddyBoardView
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByContestIdBuddyBoardsByBoardIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardView**](DomainBuddyBoardView.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers

> DomainBuddyBoardMembersResponse GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers(ctx, contestId, boardId).XTenantUserId(xTenantUserId).Execute()

List buddy board members (public or authed)



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
	boardId := "boardId_example" // string | Buddy Board ID
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers(context.Background(), contestId, boardId).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers`: DomainBuddyBoardMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsByContestIdBuddyBoardsByBoardIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contestId** | **string** | Contest ID | 
**boardId** | **string** | Buddy Board ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByContestIdBuddyBoardsByBoardIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardMembersResponse**](DomainBuddyBoardMembersResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsById

> DomainContest GetPublicContestsById(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get a contest by ID



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
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsById(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsById`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByIdGroups

> DomainContestGroups GetPublicContestsByIdGroups(ctx, id).XTenantUserId(xTenantUserId).Execute()

Get Contest Groups



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
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsByIdGroups(context.Background(), id).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsByIdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdGroups`: DomainContestGroups
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsByIdGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestGroups**](DomainContestGroups.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsByIdSubmissionFeed

> []DomainSubmissionFeedItem GetPublicContestsByIdSubmissionFeed(ctx, id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()

Get the real-time feed of contest submissions (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (default 20, max 100) (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsByIdSubmissionFeed(context.Background(), id).Cursor(cursor).Limit(limit).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsByIdSubmissionFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsByIdSubmissionFeed`: []DomainSubmissionFeedItem
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsByIdSubmissionFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsByIdSubmissionFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 20, max 100) | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainSubmissionFeedItem**](DomainSubmissionFeedItem.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicContestsSearchKeywordByKeyword

> []DomainContest GetPublicContestsSearchKeywordByKeyword(ctx, keyword).Cursor(cursor).Limit(limit).Ended(ended).XTenantUserId(xTenantUserId).Execute()

Search public contests by keyword (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (default 50) (optional)
	ended := true // bool | Include ended contests (optional)
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetPublicContestsSearchKeywordByKeyword(context.Background(), keyword).Cursor(cursor).Limit(limit).Ended(ended).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetPublicContestsSearchKeywordByKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicContestsSearchKeywordByKeyword`: []DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetPublicContestsSearchKeywordByKeyword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string** | Search keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicContestsSearchKeywordByKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50) | 
 **ended** | **bool** | Include ended contests | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**[]DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchContestsById

> DomainContest PatchContestsById(ctx, id).RequestUpdateContestRequest(requestUpdateContestRequest).XTenantUserId(xTenantUserId).Execute()

Update a contest

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
	requestUpdateContestRequest := *openapiclient.NewRequestUpdateContestRequest("Description_example", int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example"), *openapiclient.NewDomainTiming("EndAt_example", "EnrollBy_example", "StartAt_example"), "Title_example", int32(123)) // RequestUpdateContestRequest | Contest payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.PatchContestsById(context.Background(), id).RequestUpdateContestRequest(requestUpdateContestRequest).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.PatchContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContestsById`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.PatchContestsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContestsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateContestRequest** | [**RequestUpdateContestRequest**](RequestUpdateContestRequest.md) | Contest payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdHeaderImage

> DomainContest UpdateContestsByIdHeaderImage(ctx, id).RequestUpdateContestHeaderImage(requestUpdateContestHeaderImage).XTenantUserId(xTenantUserId).Execute()

Update a contest's header image

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
	requestUpdateContestHeaderImage := *openapiclient.NewRequestUpdateContestHeaderImage("HeaderImageUrl_example") // RequestUpdateContestHeaderImage | Header image payload
	xTenantUserId := "xTenantUserId_example" // string | Acting-as. The tenant's own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.UpdateContestsByIdHeaderImage(context.Background(), id).RequestUpdateContestHeaderImage(requestUpdateContestHeaderImage).XTenantUserId(xTenantUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.UpdateContestsByIdHeaderImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdHeaderImage`: DomainContest
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.UpdateContestsByIdHeaderImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContestsByIdHeaderImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateContestHeaderImage** | [**RequestUpdateContestHeaderImage**](RequestUpdateContestHeaderImage.md) | Header image payload | 
 **xTenantUserId** | **string** | Acting-as. The tenant&#39;s own identifier for the fan this request is on behalf of. The platform resolves (tenant, X-Tenant-User-Id) to a platform identity. Omit for tenant-level calls. | 

### Return type

[**DomainContest**](DomainContest.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

