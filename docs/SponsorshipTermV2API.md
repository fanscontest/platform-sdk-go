# \SponsorshipTermV2API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsorshipTerms**](SponsorshipTermV2API.md#CreateSponsorshipTerms) | **Post** /v2/sponsorship-terms | Create Sponsorship Term (v2)
[**DeleteSponsorshipTermsById**](SponsorshipTermV2API.md#DeleteSponsorshipTermsById) | **Delete** /v2/sponsorship-terms/{id} | Deactivate Sponsorship Term (v2)
[**GetSponsorshipTerms**](SponsorshipTermV2API.md#GetSponsorshipTerms) | **Get** /v2/sponsorship-terms | Get Sponsorship Terms (v2)
[**GetSponsorshipTermsById**](SponsorshipTermV2API.md#GetSponsorshipTermsById) | **Get** /v2/sponsorship-terms/{id} | Get Sponsorship Term (v2)
[**GetSponsorshipTermsPlatform**](SponsorshipTermV2API.md#GetSponsorshipTermsPlatform) | **Get** /v2/sponsorship-terms/platform | Get Platform Terms (v2)
[**UpdateSponsorshipTermsById**](SponsorshipTermV2API.md#UpdateSponsorshipTermsById) | **Put** /v2/sponsorship-terms/{id} | Update Sponsorship Term (v2)



## CreateSponsorshipTerms

> DomainSponsorshipTerm CreateSponsorshipTerms(ctx).RequestCreateSponsorshipTermRequest(requestCreateSponsorshipTermRequest).Execute()

Create Sponsorship Term (v2)



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
	requestCreateSponsorshipTermRequest := *openapiclient.NewRequestCreateSponsorshipTermRequest("Description_example", "Title_example") // RequestCreateSponsorshipTermRequest | Create Term Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermV2API.CreateSponsorshipTerms(context.Background()).RequestCreateSponsorshipTermRequest(requestCreateSponsorshipTermRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.CreateSponsorshipTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipTerms`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.CreateSponsorshipTerms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorshipTermRequest** | [**RequestCreateSponsorshipTermRequest**](RequestCreateSponsorshipTermRequest.md) | Create Term Request | 

### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSponsorshipTermsById

> HandlerStatusResponse DeleteSponsorshipTermsById(ctx, id).Execute()

Deactivate Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermV2API.DeleteSponsorshipTermsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.DeleteSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSponsorshipTermsById`: HandlerStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.DeleteSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSponsorshipTermsByIdRequest struct via the builder pattern


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


## GetSponsorshipTerms

> []DomainSponsorshipTerm GetSponsorshipTerms(ctx).Execute()

Get Sponsorship Terms (v2)



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
	resp, r, err := apiClient.SponsorshipTermV2API.GetSponsorshipTerms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.GetSponsorshipTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTerms`: []DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.GetSponsorshipTerms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsRequest struct via the builder pattern


### Return type

[**[]DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipTermsById

> DomainSponsorshipTerm GetSponsorshipTermsById(ctx, id).Execute()

Get Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermV2API.GetSponsorshipTermsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.GetSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTermsById`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.GetSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipTermsPlatform

> []DomainSponsorshipTerm GetSponsorshipTermsPlatform(ctx).Execute()

Get Platform Terms (v2)



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
	resp, r, err := apiClient.SponsorshipTermV2API.GetSponsorshipTermsPlatform(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.GetSponsorshipTermsPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipTermsPlatform`: []DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.GetSponsorshipTermsPlatform`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipTermsPlatformRequest struct via the builder pattern


### Return type

[**[]DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSponsorshipTermsById

> DomainSponsorshipTerm UpdateSponsorshipTermsById(ctx, id).RequestUpdateSponsorshipTermRequest(requestUpdateSponsorshipTermRequest).Execute()

Update Sponsorship Term (v2)



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
	id := "id_example" // string | Term ID
	requestUpdateSponsorshipTermRequest := *openapiclient.NewRequestUpdateSponsorshipTermRequest("Description_example", "Title_example", "Version_example") // RequestUpdateSponsorshipTermRequest | Update Term Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipTermV2API.UpdateSponsorshipTermsById(context.Background(), id).RequestUpdateSponsorshipTermRequest(requestUpdateSponsorshipTermRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipTermV2API.UpdateSponsorshipTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSponsorshipTermsById`: DomainSponsorshipTerm
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipTermV2API.UpdateSponsorshipTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Term ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsorshipTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateSponsorshipTermRequest** | [**RequestUpdateSponsorshipTermRequest**](RequestUpdateSponsorshipTermRequest.md) | Update Term Request | 

### Return type

[**DomainSponsorshipTerm**](DomainSponsorshipTerm.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

