# DomainEntryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainEntry**](DomainEntry.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainEntryListResponse

`func NewDomainEntryListResponse(data []DomainEntry, ) *DomainEntryListResponse`

NewDomainEntryListResponse instantiates a new DomainEntryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntryListResponseWithDefaults

`func NewDomainEntryListResponseWithDefaults() *DomainEntryListResponse`

NewDomainEntryListResponseWithDefaults instantiates a new DomainEntryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainEntryListResponse) GetData() []DomainEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainEntryListResponse) GetDataOk() (*[]DomainEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainEntryListResponse) SetData(v []DomainEntry)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainEntryListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainEntryListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainEntryListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainEntryListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


