# DomainQuestionStatsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainQuestionStats**](DomainQuestionStats.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainQuestionStatsListResponse

`func NewDomainQuestionStatsListResponse(data []DomainQuestionStats, ) *DomainQuestionStatsListResponse`

NewDomainQuestionStatsListResponse instantiates a new DomainQuestionStatsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainQuestionStatsListResponseWithDefaults

`func NewDomainQuestionStatsListResponseWithDefaults() *DomainQuestionStatsListResponse`

NewDomainQuestionStatsListResponseWithDefaults instantiates a new DomainQuestionStatsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainQuestionStatsListResponse) GetData() []DomainQuestionStats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainQuestionStatsListResponse) GetDataOk() (*[]DomainQuestionStats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainQuestionStatsListResponse) SetData(v []DomainQuestionStats)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainQuestionStatsListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainQuestionStatsListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainQuestionStatsListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainQuestionStatsListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


