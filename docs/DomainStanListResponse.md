# DomainStanListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainStan**](DomainStan.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainStanListResponse

`func NewDomainStanListResponse(data []DomainStan, ) *DomainStanListResponse`

NewDomainStanListResponse instantiates a new DomainStanListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainStanListResponseWithDefaults

`func NewDomainStanListResponseWithDefaults() *DomainStanListResponse`

NewDomainStanListResponseWithDefaults instantiates a new DomainStanListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainStanListResponse) GetData() []DomainStan`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainStanListResponse) GetDataOk() (*[]DomainStan, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainStanListResponse) SetData(v []DomainStan)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainStanListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainStanListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainStanListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainStanListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


