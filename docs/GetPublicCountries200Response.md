# GetPublicCountries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainCountry**](DomainCountry.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetPublicCountries200Response

`func NewGetPublicCountries200Response(data []DomainCountry, ) *GetPublicCountries200Response`

NewGetPublicCountries200Response instantiates a new GetPublicCountries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicCountries200ResponseWithDefaults

`func NewGetPublicCountries200ResponseWithDefaults() *GetPublicCountries200Response`

NewGetPublicCountries200ResponseWithDefaults instantiates a new GetPublicCountries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPublicCountries200Response) GetData() []DomainCountry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPublicCountries200Response) GetDataOk() (*[]DomainCountry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPublicCountries200Response) SetData(v []DomainCountry)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetPublicCountries200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetPublicCountries200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetPublicCountries200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetPublicCountries200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


