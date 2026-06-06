# DomainTeamSetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainTeamSet**](DomainTeamSet.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainTeamSetListResponse

`func NewDomainTeamSetListResponse(data []DomainTeamSet, ) *DomainTeamSetListResponse`

NewDomainTeamSetListResponse instantiates a new DomainTeamSetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTeamSetListResponseWithDefaults

`func NewDomainTeamSetListResponseWithDefaults() *DomainTeamSetListResponse`

NewDomainTeamSetListResponseWithDefaults instantiates a new DomainTeamSetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainTeamSetListResponse) GetData() []DomainTeamSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainTeamSetListResponse) GetDataOk() (*[]DomainTeamSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainTeamSetListResponse) SetData(v []DomainTeamSet)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainTeamSetListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainTeamSetListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainTeamSetListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainTeamSetListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


