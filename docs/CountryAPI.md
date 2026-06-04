# \CountryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicCountries**](CountryAPI.md#GetPublicCountries) | **Get** /v2/public/countries | List enabled countries (sorted, with native names for non-English Accept-Language)
[**GetPublicLocation**](CountryAPI.md#GetPublicLocation) | **Get** /v2/public/location | Resolve the caller&#39;s country by IP (CF-Ipcountry, then Maxmind)



## GetPublicCountries

> []DomainCountry GetPublicCountries(ctx).Execute()

List enabled countries (sorted, with native names for non-English Accept-Language)

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
	resp, r, err := apiClient.CountryAPI.GetPublicCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.GetPublicCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicCountries`: []DomainCountry
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.GetPublicCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicCountriesRequest struct via the builder pattern


### Return type

[**[]DomainCountry**](DomainCountry.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicLocation

> DomainCountry GetPublicLocation(ctx).Execute()

Resolve the caller's country by IP (CF-Ipcountry, then Maxmind)

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
	resp, r, err := apiClient.CountryAPI.GetPublicLocation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.GetPublicLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicLocation`: DomainCountry
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.GetPublicLocation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicLocationRequest struct via the builder pattern


### Return type

[**DomainCountry**](DomainCountry.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

