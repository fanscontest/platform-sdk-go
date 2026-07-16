# \TournamentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTournaments**](TournamentsAPI.md#CreateTournaments) | **Post** /v2/tournaments | Create tournament
[**GetChannelsByIdTournamentConfiguration**](TournamentsAPI.md#GetChannelsByIdTournamentConfiguration) | **Get** /v2/channels/{id}/tournament-configuration | Get tournament configuration templates
[**GetChannelsByIdTournaments**](TournamentsAPI.md#GetChannelsByIdTournaments) | **Get** /v2/channels/{id}/tournaments | List channel tournaments (cursor-paginated)
[**GetIdentitiesByPiidTournaments**](TournamentsAPI.md#GetIdentitiesByPiidTournaments) | **Get** /v2/identities/{piid}/tournaments | List tournaments owned by a platform identity (cursor-paginated)
[**GetTournaments**](TournamentsAPI.md#GetTournaments) | **Get** /v2/tournaments | List the caller&#39;s tournaments (cursor-paginated)
[**GetTournamentsById**](TournamentsAPI.md#GetTournamentsById) | **Get** /v2/tournaments/{id} | Get tournament



## CreateTournaments

> DomainTournamentResponse CreateTournaments(ctx).XActingAs(xActingAs).Execute()

Create tournament



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.CreateTournaments(context.Background()).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.CreateTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTournaments`: DomainTournamentResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.CreateTournaments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentResponse**](DomainTournamentResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdTournamentConfiguration

> DomainTournamentConfigurationResponse GetChannelsByIdTournamentConfiguration(ctx, id).XActingAs(xActingAs).Execute()

Get tournament configuration templates



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetChannelsByIdTournamentConfiguration(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetChannelsByIdTournamentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdTournamentConfiguration`: DomainTournamentConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetChannelsByIdTournamentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsByIdTournamentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentConfigurationResponse**](DomainTournamentConfigurationResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdTournaments

> DomainTournamentListResponse GetChannelsByIdTournaments(ctx, id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List channel tournaments (cursor-paginated)

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
	limit := int32(56) // int32 | Page size (optional)
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetChannelsByIdTournaments(context.Background(), id).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetChannelsByIdTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdTournaments`: DomainTournamentListResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetChannelsByIdTournaments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsByIdTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentListResponse**](DomainTournamentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesByPiidTournaments

> DomainTournamentListResponse GetIdentitiesByPiidTournaments(ctx, piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List tournaments owned by a platform identity (cursor-paginated)

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
	resp, r, err := apiClient.TournamentsAPI.GetIdentitiesByPiidTournaments(context.Background(), piid).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetIdentitiesByPiidTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitiesByPiidTournaments`: DomainTournamentListResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetIdentitiesByPiidTournaments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**piid** | **string** | Platform Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesByPiidTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentListResponse**](DomainTournamentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournaments

> DomainTournamentListResponse GetTournaments(ctx).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()

List the caller's tournaments (cursor-paginated)

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
	resp, r, err := apiClient.TournamentsAPI.GetTournaments(context.Background()).Cursor(cursor).Limit(limit).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournaments`: DomainTournamentListResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournaments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size | 
 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentListResponse**](DomainTournamentListResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentsById

> DomainTournamentResponse GetTournamentsById(ctx, id).XActingAs(xActingAs).Execute()

Get tournament



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
	xActingAs := "xActingAs_example" // string | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentsAPI.GetTournamentsById(context.Background(), id).XActingAs(xActingAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentsAPI.GetTournamentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentsById`: DomainTournamentResponse
	fmt.Fprintf(os.Stdout, "Response from `TournamentsAPI.GetTournamentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tournament ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTournamentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xActingAs** | **string** | Acting-as. The platform identity id (piid) this request is on behalf of. The platform verifies the piid belongs to the calling tenant and acts as that identity. Omit for tenant-level calls. | 

### Return type

[**DomainTournamentResponse**](DomainTournamentResponse.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

