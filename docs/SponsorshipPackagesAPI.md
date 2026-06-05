# \SponsorshipPackagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSponsorshipPackages**](SponsorshipPackagesAPI.md#CreateSponsorshipPackages) | **Post** /v2/sponsorship-packages | Create Sponsorship Package (v2)
[**CreateSponsorshipPackagesByIdCancel**](SponsorshipPackagesAPI.md#CreateSponsorshipPackagesByIdCancel) | **Post** /v2/sponsorship-packages/{id}/cancel | Cancel Sponsorship Package (v2)
[**GetSponsorshipPackages**](SponsorshipPackagesAPI.md#GetSponsorshipPackages) | **Get** /v2/sponsorship-packages | Get Sponsorship Packages (v2)
[**GetSponsorshipPackagesAvailable**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesAvailable) | **Get** /v2/sponsorship-packages/available | Get Available Packages (v2)
[**GetSponsorshipPackagesByIdAgreement**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesByIdAgreement) | **Get** /v2/sponsorship-packages/{id}/agreement | Get Package Agreement (v2)
[**GetSponsorshipPackagesByIdSponsorships**](SponsorshipPackagesAPI.md#GetSponsorshipPackagesByIdSponsorships) | **Get** /v2/sponsorship-packages/{id}/sponsorships | Get Package Sponsorships (v2)



## CreateSponsorshipPackages

> DomainSponsorshipPackage CreateSponsorshipPackages(ctx).RequestCreateSponsorshipPackageRequest(requestCreateSponsorshipPackageRequest).Execute()

Create Sponsorship Package (v2)



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
	requestCreateSponsorshipPackageRequest := *openapiclient.NewRequestCreateSponsorshipPackageRequest("BudgetPerContestCurrency_example", "BudgetPerContestValue_example", []string{"TermIds_example"}, "TotalBudgetCurrency_example", "TotalBudgetValue_example", int32(123)) // RequestCreateSponsorshipPackageRequest | Create Sponsorship Package Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.CreateSponsorshipPackages(context.Background()).RequestCreateSponsorshipPackageRequest(requestCreateSponsorshipPackageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.CreateSponsorshipPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipPackages`: DomainSponsorshipPackage
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.CreateSponsorshipPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateSponsorshipPackageRequest** | [**RequestCreateSponsorshipPackageRequest**](RequestCreateSponsorshipPackageRequest.md) | Create Sponsorship Package Request | 

### Return type

[**DomainSponsorshipPackage**](DomainSponsorshipPackage.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSponsorshipPackagesByIdCancel

> DomainSponsorshipPackage CreateSponsorshipPackagesByIdCancel(ctx, id).Execute()

Cancel Sponsorship Package (v2)



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
	id := "id_example" // string | Package ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSponsorshipPackagesByIdCancel`: DomainSponsorshipPackage
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.CreateSponsorshipPackagesByIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsorshipPackagesByIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorshipPackage**](DomainSponsorshipPackage.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackages

> []DomainSponsorshipPackage GetSponsorshipPackages(ctx).Execute()

Get Sponsorship Packages (v2)



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
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackages`: []DomainSponsorshipPackage
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesRequest struct via the builder pattern


### Return type

[**[]DomainSponsorshipPackage**](DomainSponsorshipPackage.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesAvailable

> []DomainSponsorshipPackage GetSponsorshipPackagesAvailable(ctx).SponsorId(sponsorId).PackageId(packageId).Execute()

Get Available Packages (v2)



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
	sponsorId := "sponsorId_example" // string | Filter by sponsor ID (optional)
	packageId := "packageId_example" // string | Get specific package by ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable(context.Background()).SponsorId(sponsorId).PackageId(packageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesAvailable`: []DomainSponsorshipPackage
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesAvailable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sponsorId** | **string** | Filter by sponsor ID | 
 **packageId** | **string** | Get specific package by ID | 

### Return type

[**[]DomainSponsorshipPackage**](DomainSponsorshipPackage.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesByIdAgreement

> DomainSponsorshipPackageAgreement GetSponsorshipPackagesByIdAgreement(ctx, id).Execute()

Get Package Agreement (v2)



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
	id := "id_example" // string | Package ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesByIdAgreement`: DomainSponsorshipPackageAgreement
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesByIdAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorshipPackageAgreement**](DomainSponsorshipPackageAgreement.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsorshipPackagesByIdSponsorships

> []DomainSponsorship GetSponsorshipPackagesByIdSponsorships(ctx, id).Execute()

Get Package Sponsorships (v2)



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
	id := "id_example" // string | Package ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipPackagesByIdSponsorships`: []DomainSponsorship
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipPackagesAPI.GetSponsorshipPackagesByIdSponsorships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Package ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipPackagesByIdSponsorshipsRequest struct via the builder pattern


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

