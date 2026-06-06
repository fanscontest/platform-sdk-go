# DomainJurorAssignmentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainJurorAssignment**](DomainJurorAssignment.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainJurorAssignmentListResponse

`func NewDomainJurorAssignmentListResponse(data []DomainJurorAssignment, ) *DomainJurorAssignmentListResponse`

NewDomainJurorAssignmentListResponse instantiates a new DomainJurorAssignmentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainJurorAssignmentListResponseWithDefaults

`func NewDomainJurorAssignmentListResponseWithDefaults() *DomainJurorAssignmentListResponse`

NewDomainJurorAssignmentListResponseWithDefaults instantiates a new DomainJurorAssignmentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainJurorAssignmentListResponse) GetData() []DomainJurorAssignment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainJurorAssignmentListResponse) GetDataOk() (*[]DomainJurorAssignment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainJurorAssignmentListResponse) SetData(v []DomainJurorAssignment)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainJurorAssignmentListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainJurorAssignmentListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainJurorAssignmentListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainJurorAssignmentListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


