# GetVenues200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainVenue**](DomainVenue.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetVenues200Response

`func NewGetVenues200Response(data []DomainVenue, ) *GetVenues200Response`

NewGetVenues200Response instantiates a new GetVenues200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVenues200ResponseWithDefaults

`func NewGetVenues200ResponseWithDefaults() *GetVenues200Response`

NewGetVenues200ResponseWithDefaults instantiates a new GetVenues200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVenues200Response) GetData() []DomainVenue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVenues200Response) GetDataOk() (*[]DomainVenue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVenues200Response) SetData(v []DomainVenue)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetVenues200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetVenues200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetVenues200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetVenues200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


