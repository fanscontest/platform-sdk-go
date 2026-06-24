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
[**GetAdminHostingInvitations**](ContestsAPI.md#GetAdminHostingInvitations) | **Get** /v2/admin/hosting-invitations | List pending hosting invitations for the calling tenant&#39;s venues
[**GetChannelsByIdContests**](ContestsAPI.md#GetChannelsByIdContests) | **Get** /v2/channels/{id}/contests | Get channel contests (cursor-paginated)
[**GetContests**](ContestsAPI.md#GetContests) | **Get** /v2/contests | List the caller&#39;s contests (cursor-paginated)
[**GetContestsByContestIdBuddyBoardsByBoardId**](ContestsAPI.md#GetContestsByContestIdBuddyBoardsByBoardId) | **Get** /v2/contests/{contestId}/buddy-boards/{boardId} | Get buddy board view (metadata + scores filtered to board members)
[**GetContestsByContestIdBuddyBoardsByBoardIdMembers**](ContestsAPI.md#GetContestsByContestIdBuddyBoardsByBoardIdMembers) | **Get** /v2/contests/{contestId}/buddy-boards/{boardId}/members | List buddy board members (public or authed)
[**GetContestsByContestIdParticipantsByMemberIdResult**](ContestsAPI.md#GetContestsByContestIdParticipantsByMemberIdResult) | **Get** /v2/contests/{contestId}/participants/{memberId}/result | Get participant result for a contest
[**GetContestsById**](ContestsAPI.md#GetContestsById) | **Get** /v2/contests/{id} | Get a contest by ID
[**GetContestsByIdGroups**](ContestsAPI.md#GetContestsByIdGroups) | **Get** /v2/contests/{id}/groups | Get Contest Groups
[**GetContestsByIdHostings**](ContestsAPI.md#GetContestsByIdHostings) | **Get** /v2/contests/{id}/hostings | List hostings for a contest the calling tenant owns
[**GetContestsByIdParticipation**](ContestsAPI.md#GetContestsByIdParticipation) | **Get** /v2/contests/{id}/participation | Get contest participation result (cursor-paginated)
[**GetContestsByIdQuestionStats**](ContestsAPI.md#GetContestsByIdQuestionStats) | **Get** /v2/contests/{id}/question-stats | Get question statistics for a contest (Contest Owner or Admin Only)
[**GetContestsByIdSubmissionFeed**](ContestsAPI.md#GetContestsByIdSubmissionFeed) | **Get** /v2/contests/{id}/submission-feed | Get the real-time feed of contest submissions (cursor-paginated)
[**GetContestsByRegion**](ContestsAPI.md#GetContestsByRegion) | **Get** /v2/contests/by-region | List contests by region (tenant-scoped, cursor-paginated)
[**GetContestsSearch**](ContestsAPI.md#GetContestsSearch) | **Get** /v2/contests/search | Search contests not owned by the caller (cursor-paginated)
[**GetContestsSearchKeywordByKeyword**](ContestsAPI.md#GetContestsSearchKeywordByKeyword) | **Get** /v2/contests/searchKeyword/{keyword} | Search contests by keyword (tenant-scoped, cursor-paginated)
[**GetIdentitiesByPiidBuddyBoardInvites**](ContestsAPI.md#GetIdentitiesByPiidBuddyBoardInvites) | **Get** /v2/identities/{piid}/buddy-board-invites | List buddy board invites for a platform identity
[**GetIdentitiesByPiidContests**](ContestsAPI.md#GetIdentitiesByPiidContests) | **Get** /v2/identities/{piid}/contests | List contests for a platform identity (cursor-paginated)
[**PatchContestsById**](ContestsAPI.md#PatchContestsById) | **Patch** /v2/contests/{id} | Update a contest
[**UpdateContestsByIdHeaderImage**](ContestsAPI.md#UpdateContestsByIdHeaderImage) | **Put** /v2/contests/{id}/header-image | Update a contest&#39;s header image



## CreateContests

> DomainContestResponse CreateContests(ctx).RequestCreateContestRequest(requestCreateContestRequest).XActingAs(xActingAs).Execute()

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
	requestCreateContestRequest := *openapiclient.NewRequestCreateContestRequest("ChannelId_example", int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example"), "SourceId_example", "SourceType_example", *openapiclient.NewDomainTiming("EndAt_example", "EnrollBy_example", "StartAt_example"), "Title_example", int32(123), int32(123)) // RequestCreateContestRequest | Contest payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContests(context.Background()).RequestCreateContestRequest(requestCreateContestRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContests`: DomainContestResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.CreateContests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateContestRequest** | [**RequestCreateContestRequest**](RequestCreateContestRequest.md) | Contest payload | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestResponse**](DomainContestResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByContestIdBuddyBoards

> DomainBuddyBoardResponse CreateContestsByContestIdBuddyBoards(ctx, contestId).RequestCreateBuddyBoardRequest(requestCreateBuddyBoardRequest).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoards(context.Background(), contestId).RequestCreateBuddyBoardRequest(requestCreateBuddyBoardRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoards`: DomainBuddyBoardResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardResponse**](DomainBuddyBoardResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept

> HandlerStatusResponseResponse CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept(ctx, contestId, boardId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept(context.Background(), contestId, boardId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdInvitesAccept`: HandlerStatusResponseResponse
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


## CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline

> HandlerStatusResponseResponse CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline(ctx, contestId, boardId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline(context.Background(), contestId, boardId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdInvitesDecline`: HandlerStatusResponseResponse
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


## CreateContestsByContestIdBuddyBoardsByBoardIdMembers

> HandlerStatusResponseResponse CreateContestsByContestIdBuddyBoardsByBoardIdMembers(ctx, contestId, boardId).RequestAddBuddyBoardMembersRequest(requestAddBuddyBoardMembersRequest).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdMembers(context.Background(), contestId, boardId).RequestAddBuddyBoardMembersRequest(requestAddBuddyBoardMembersRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByContestIdBuddyBoardsByBoardIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByContestIdBuddyBoardsByBoardIdMembers`: HandlerStatusResponseResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdHostings

> HandlerStatusResponseResponse CreateContestsByIdHostings(ctx, id).HandlerCreatePlatformContestHostingRequest(handlerCreatePlatformContestHostingRequest).XActingAs(xActingAs).Execute()

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
	handlerCreatePlatformContestHostingRequest := *openapiclient.NewHandlerCreatePlatformContestHostingRequest("ChannelId_example") // HandlerCreatePlatformContestHostingRequest | Hosting payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostings(context.Background(), id).HandlerCreatePlatformContestHostingRequest(handlerCreatePlatformContestHostingRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostings`: HandlerStatusResponseResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContestsByIdHostingsByHostingIdAccept

> HandlerStatusResponseResponse CreateContestsByIdHostingsByHostingIdAccept(ctx, id, hostingId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostingsByHostingIdAccept(context.Background(), id, hostingId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostingsByHostingIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostingsByHostingIdAccept`: HandlerStatusResponseResponse
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


## CreateContestsByIdHostingsByHostingIdDecline

> HandlerStatusResponseResponse CreateContestsByIdHostingsByHostingIdDecline(ctx, id, hostingId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdHostingsByHostingIdDecline(context.Background(), id, hostingId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdHostingsByHostingIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdHostingsByHostingIdDecline`: HandlerStatusResponseResponse
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


## CreateContestsByIdParticipations

> HandlerStatusResponseResponse CreateContestsByIdParticipations(ctx, id).HandlerCreatePlatformParticipationRequest(handlerCreatePlatformParticipationRequest).XActingAs(xActingAs).Execute()

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
	handlerCreatePlatformParticipationRequest := *openapiclient.NewHandlerCreatePlatformParticipationRequest("ChannelId_example", "DisplayName_example", "PlatformIdentityId_example") // HandlerCreatePlatformParticipationRequest | Participation payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateContestsByIdParticipations(context.Background(), id).HandlerCreatePlatformParticipationRequest(handlerCreatePlatformParticipationRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateContestsByIdParticipations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContestsByIdParticipations`: HandlerStatusResponseResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**HandlerStatusResponseResponse**](HandlerStatusResponseResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTournamentsByIdContest

> DomainContestResponse CreateTournamentsByIdContest(ctx, id).RequestCreateContestRequest(requestCreateContestRequest).XActingAs(xActingAs).Execute()

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
	requestCreateContestRequest := *openapiclient.NewRequestCreateContestRequest("ChannelId_example", int32(123), *openapiclient.NewDomainReward("RewardType_example", "Value_example"), "SourceId_example", "SourceType_example", *openapiclient.NewDomainTiming("EndAt_example", "EnrollBy_example", "StartAt_example"), "Title_example", int32(123), int32(123)) // RequestCreateContestRequest | Contest payload
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.CreateTournamentsByIdContest(context.Background(), id).RequestCreateContestRequest(requestCreateContestRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.CreateTournamentsByIdContest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTournamentsByIdContest`: DomainContestResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestResponse**](DomainContestResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContestsById

> CreateChannelsByIdReport200Response DeleteContestsById(ctx, id).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.DeleteContestsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.DeleteContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContestsById`: CreateChannelsByIdReport200Response
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

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**CreateChannelsByIdReport200Response**](CreateChannelsByIdReport200Response.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminHostingInvitations

> DomainContestHostingListResponse GetAdminHostingInvitations(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List pending hosting invitations for the calling tenant's venues

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetAdminHostingInvitations(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetAdminHostingInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminHostingInvitations`: DomainContestHostingListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetAdminHostingInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminHostingInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestHostingListResponse**](DomainContestHostingListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdContests

> DomainContestListResponse GetChannelsByIdContests(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetChannelsByIdContests(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetChannelsByIdContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdContests`: DomainContestListResponse
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


## GetContests

> DomainContestListResponse GetContests(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContests(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContests`: DomainContestListResponse
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


## GetContestsByContestIdBuddyBoardsByBoardId

> DomainBuddyBoardViewResponse GetContestsByContestIdBuddyBoardsByBoardId(ctx, contestId, boardId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardId(context.Background(), contestId, boardId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdBuddyBoardsByBoardId`: DomainBuddyBoardViewResponse
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


 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardViewResponse**](DomainBuddyBoardViewResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdBuddyBoardsByBoardIdMembers

> DomainPlatformIdentityListResponse GetContestsByContestIdBuddyBoardsByBoardIdMembers(ctx, contestId, boardId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardIdMembers(context.Background(), contestId, boardId).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdBuddyBoardsByBoardIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdBuddyBoardsByBoardIdMembers`: DomainPlatformIdentityListResponse
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


 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainPlatformIdentityListResponse**](DomainPlatformIdentityListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByContestIdParticipantsByMemberIdResult

> DomainMemberResultResponse GetContestsByContestIdParticipantsByMemberIdResult(ctx, contestId, memberId).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByContestIdParticipantsByMemberIdResult(context.Background(), contestId, memberId).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByContestIdParticipantsByMemberIdResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByContestIdParticipantsByMemberIdResult`: DomainMemberResultResponse
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


 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainMemberResultResponse**](DomainMemberResultResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsById

> DomainContestResponse GetContestsById(ctx, id).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsById`: DomainContestResponse
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

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestResponse**](DomainContestResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdGroups

> DomainContestGroupsResponse GetContestsByIdGroups(ctx, id).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdGroups(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdGroups`: DomainContestGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByIdGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestGroupsResponse**](DomainContestGroupsResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdHostings

> DomainContestHostingListResponse GetContestsByIdHostings(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdHostings(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdHostings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdHostings`: DomainContestHostingListResponse
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

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestHostingListResponse**](DomainContestHostingListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdParticipation

> DomainParticipationListResponse GetContestsByIdParticipation(ctx, id).Cursor(cursor).Limit(limit).Winners(winners).Status(status).XActingAs(xActingAs).Execute()

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
	winners := true // bool | Restrict to winning participations only (optional)
	status := "status_example" // string | Restrict to a single participation status (enrolled, participated, timed_out); omit for all (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdParticipation(context.Background(), id).Cursor(cursor).Limit(limit).Winners(winners).Status(status).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdParticipation`: DomainParticipationListResponse
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
 **winners** | **bool** | Restrict to winning participations only | 
 **status** | **string** | Restrict to a single participation status (enrolled, participated, timed_out); omit for all | 
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


## GetContestsByIdQuestionStats

> DomainQuestionStatsListResponse GetContestsByIdQuestionStats(ctx, id).Sort(sort).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	cursor := "cursor_example" // string | Accepted for shape uniformity; ignored (single-page aggregate) (optional)
	limit := int32(56) // int32 | Accepted for shape uniformity; ignored (single-page aggregate) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdQuestionStats(context.Background(), id).Sort(sort).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdQuestionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdQuestionStats`: DomainQuestionStatsListResponse
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
 **cursor** | **string** | Accepted for shape uniformity; ignored (single-page aggregate) | 
 **limit** | **int32** | Accepted for shape uniformity; ignored (single-page aggregate) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainQuestionStatsListResponse**](DomainQuestionStatsListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByIdSubmissionFeed

> DomainSubmissionFeedItemListResponse GetContestsByIdSubmissionFeed(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByIdSubmissionFeed(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByIdSubmissionFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByIdSubmissionFeed`: DomainSubmissionFeedItemListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByIdSubmissionFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Contest ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByIdSubmissionFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 20, max 100) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainSubmissionFeedItemListResponse**](DomainSubmissionFeedItemListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContestsByRegion

> DomainContestListResponse GetContestsByRegion(ctx).Cursor(cursor).Limit(limit).Ended(ended).Region(region).XActingAs(xActingAs).Execute()

List contests by region (tenant-scoped, cursor-paginated)

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsByRegion(context.Background()).Cursor(cursor).Limit(limit).Ended(ended).Region(region).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsByRegion`: DomainContestListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsByRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50) | 
 **ended** | **bool** | Include ended contests | 
 **region** | **string** | ISO 3166-1 alpha-2 region filter | 
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


## GetContestsSearch

> DomainContestListResponse GetContestsSearch(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsSearch(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsSearch`: DomainContestListResponse
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


## GetContestsSearchKeywordByKeyword

> DomainContestListResponse GetContestsSearchKeywordByKeyword(ctx, keyword).Cursor(cursor).Limit(limit).Ended(ended).XActingAs(xActingAs).Execute()

Search contests by keyword (tenant-scoped, cursor-paginated)

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetContestsSearchKeywordByKeyword(context.Background(), keyword).Cursor(cursor).Limit(limit).Ended(ended).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetContestsSearchKeywordByKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContestsSearchKeywordByKeyword`: DomainContestListResponse
	fmt.Fprintf(os.Stdout, "Response from `ContestsAPI.GetContestsSearchKeywordByKeyword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string** | Search keyword | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContestsSearchKeywordByKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50) | 
 **ended** | **bool** | Include ended contests | 
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


## GetIdentitiesByPiidBuddyBoardInvites

> DomainBuddyBoardInviteListResponse GetIdentitiesByPiidBuddyBoardInvites(ctx, piid).Status(status).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetIdentitiesByPiidBuddyBoardInvites(context.Background(), piid).Status(status).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetIdentitiesByPiidBuddyBoardInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidBuddyBoardInvites`: DomainBuddyBoardInviteListResponse
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
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainBuddyBoardInviteListResponse**](DomainBuddyBoardInviteListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidContests

> DomainContestListResponse GetIdentitiesByPiidContests(ctx, piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.GetIdentitiesByPiidContests(context.Background(), piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.GetIdentitiesByPiidContests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidContests`: DomainContestListResponse
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


## PatchContestsById

> DomainContestResponse PatchContestsById(ctx, id).RequestUpdateContestRequest(requestUpdateContestRequest).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.PatchContestsById(context.Background(), id).RequestUpdateContestRequest(requestUpdateContestRequest).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.PatchContestsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchContestsById`: DomainContestResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestResponse**](DomainContestResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContestsByIdHeaderImage

> DomainContestResponse UpdateContestsByIdHeaderImage(ctx, id).RequestUpdateContestHeaderImage(requestUpdateContestHeaderImage).XActingAs(xActingAs).Execute()

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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContestsAPI.UpdateContestsByIdHeaderImage(context.Background(), id).RequestUpdateContestHeaderImage(requestUpdateContestHeaderImage).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContestsAPI.UpdateContestsByIdHeaderImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContestsByIdHeaderImage`: DomainContestResponse
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
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainContestResponse**](DomainContestResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

