# \SponsorshipAgreementV2API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSponsorshipsByIdAgreement**](SponsorshipAgreementV2API.md#GetSponsorshipsByIdAgreement) | **Get** /v2/sponsorships/{id}/agreement | Get Sponsorship Agreement (v2)



## GetSponsorshipsByIdAgreement

> DomainSponsorshipAgreement GetSponsorshipsByIdAgreement(ctx, id).Execute()

Get Sponsorship Agreement (v2)



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
	id := "id_example" // string | Sponsorship ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SponsorshipAgreementV2API.GetSponsorshipsByIdAgreement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SponsorshipAgreementV2API.GetSponsorshipsByIdAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsorshipsByIdAgreement`: DomainSponsorshipAgreement
	fmt.Fprintf(os.Stdout, "Response from `SponsorshipAgreementV2API.GetSponsorshipsByIdAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Sponsorship ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorshipsByIdAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainSponsorshipAgreement**](DomainSponsorshipAgreement.md)

### Authorization

[TenantApiKey](../README.md#TenantApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

