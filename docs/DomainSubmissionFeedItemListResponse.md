# DomainSubmissionFeedItemListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSubmissionFeedItem**](DomainSubmissionFeedItem.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainSubmissionFeedItemListResponse

`func NewDomainSubmissionFeedItemListResponse(data []DomainSubmissionFeedItem, ) *DomainSubmissionFeedItemListResponse`

NewDomainSubmissionFeedItemListResponse instantiates a new DomainSubmissionFeedItemListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSubmissionFeedItemListResponseWithDefaults

`func NewDomainSubmissionFeedItemListResponseWithDefaults() *DomainSubmissionFeedItemListResponse`

NewDomainSubmissionFeedItemListResponseWithDefaults instantiates a new DomainSubmissionFeedItemListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainSubmissionFeedItemListResponse) GetData() []DomainSubmissionFeedItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainSubmissionFeedItemListResponse) GetDataOk() (*[]DomainSubmissionFeedItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainSubmissionFeedItemListResponse) SetData(v []DomainSubmissionFeedItem)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainSubmissionFeedItemListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainSubmissionFeedItemListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainSubmissionFeedItemListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainSubmissionFeedItemListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


