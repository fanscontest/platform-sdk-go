# \SponsorsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsors**](SponsorsAPI.md#CreateSponsors) | **Post** /v2/sponsors | Create Sponsor (v2)
[**GetSponsors**](SponsorsAPI.md#GetSponsors) | **Get** /v2/sponsors | Get Sponsors (v2)
[**GetSponsorsById**](SponsorsAPI.md#GetSponsorsById) | **Get** /v2/sponsors/{id} | Get Sponsor by ID (v2)
[**GetSponsorsByIdSponsorships**](SponsorsAPI.md#GetSponsorsByIdSponsorships) | **Get** /v2/sponsors/{id}/sponsorships | Get Sponsor Sponsorships (v2)



## CreateSponsors

> DomainSponsor CreateSponsors(ctx).RequestCreateSponsorRequest(requestCreateSponsorRequest).Execute()

Create Sponsor (v2)



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
	requestCreateSponsorRequest := *openapiclient.NewRequestCreateSponsorRequest("Name_example", "Url_example") // RequestCreateSponsorRequest | Create Sponsor Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.CreateSponsors(context.Background()).RequestCreateSponsorRequest(requestCreateSponsorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.CreateSponsors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsors`: DomainSponsor
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.CreateSponsors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorRequest** | [**RequestCreateSponsorRequest**](RequestCreateSponsorRequest.md) | Create Sponsor Request | 

### Return type

[**DomainSponsor**](DomainSponsor.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsors

> []DomainSponsor GetSponsors(ctx).Execute()

Get Sponsors (v2)



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
	resp, r, err := apiClient.SponsorsAPI.GetSponsors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsors`: []DomainSponsor
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsRequest struct via the builder pattern


### Return type

[**[]DomainSponsor**](DomainSponsor.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorsById

> DomainSponsor GetSponsorsById(ctx, id).Execute()

Get Sponsor by ID (v2)



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
	id := "id_example" // string | Sponsor ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.GetSponsorsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorsById`: DomainSponsor
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsor**](DomainSponsor.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorsByIdSponsorships

> []DomainSponsorship GetSponsorsByIdSponsorships(ctx, id).Execute()

Get Sponsor Sponsorships (v2)



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
	id := "id_example" // string | Sponsor ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorsAPI.GetSponsorsByIdSponsorships(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorsAPI.GetSponsorsByIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorsByIdSponsorships`: []DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorsAPI.GetSponsorsByIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorsByIdSponsorshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainSponsorship**](DomainSponsorship.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

