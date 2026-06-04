# \TournamentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTournaments**](TournamentAPI.md#CreateTournaments) | **Post** /v2/tournaments | Create tournament
[**GetChannelsByIdTournamentConfiguration**](TournamentAPI.md#GetChannelsByIdTournamentConfiguration) | **Get** /v2/channels/{id}/tournament-configuration | Get tournament configuration templates
[**GetChannelsByIdTournaments**](TournamentAPI.md#GetChannelsByIdTournaments) | **Get** /v2/channels/{id}/tournaments | List channel tournaments (cursor-paginated)
[**GetTournamentsById**](TournamentAPI.md#GetTournamentsById) | **Get** /v2/tournaments/{id} | Get tournament



## CreateTournaments

> DomainTournament CreateTournaments(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentAPI.CreateTournaments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentAPI.CreateTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTournaments`: DomainTournament
	fmt.Fprintf(os.Stdout, "Response from `TournamentAPI.CreateTournaments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTournamentsRequest struct via the builder pattern


### Return type

[**DomainTournament**](DomainTournament.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdTournamentConfiguration

> DomainTournamentConfiguration GetChannelsByIdTournamentConfiguration(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentAPI.GetChannelsByIdTournamentConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentAPI.GetChannelsByIdTournamentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdTournamentConfiguration`: DomainTournamentConfiguration
	fmt.Fprintf(os.Stdout, "Response from `TournamentAPI.GetChannelsByIdTournamentConfiguration`: %v\n", resp)
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


### Return type

[**DomainTournamentConfiguration**](DomainTournamentConfiguration.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelsByIdTournaments

> []DomainTournament GetChannelsByIdTournaments(ctx, id).Cursor(cursor).Limit(limit).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentAPI.GetChannelsByIdTournaments(context.Background(), id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentAPI.GetChannelsByIdTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsByIdTournaments`: []DomainTournament
	fmt.Fprintf(os.Stdout, "Response from `TournamentAPI.GetChannelsByIdTournaments`: %v\n", resp)
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

### Return type

[**[]DomainTournament**](DomainTournament.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTournamentsById

> DomainTournament GetTournamentsById(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TournamentAPI.GetTournamentsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TournamentAPI.GetTournamentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTournamentsById`: DomainTournament
	fmt.Fprintf(os.Stdout, "Response from `TournamentAPI.GetTournamentsById`: %v\n", resp)
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


### Return type

[**DomainTournament**](DomainTournament.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

