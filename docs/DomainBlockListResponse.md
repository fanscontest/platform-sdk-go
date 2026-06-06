# DomainBlockListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainBlock**](DomainBlock.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainBlockListResponse

`func NewDomainBlockListResponse(data []DomainBlock, ) *DomainBlockListResponse`

NewDomainBlockListResponse instantiates a new DomainBlockListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBlockListResponseWithDefaults

`func NewDomainBlockListResponseWithDefaults() *DomainBlockListResponse`

NewDomainBlockListResponseWithDefaults instantiates a new DomainBlockListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainBlockListResponse) GetData() []DomainBlock`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainBlockListResponse) GetDataOk() (*[]DomainBlock, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainBlockListResponse) SetData(v []DomainBlock)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainBlockListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainBlockListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainBlockListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainBlockListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


