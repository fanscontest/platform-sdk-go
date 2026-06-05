# GetPublicContestsByIdSubmissionFeed200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSubmissionFeedItem**](DomainSubmissionFeedItem.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetPublicContestsByIdSubmissionFeed200Response

`func NewGetPublicContestsByIdSubmissionFeed200Response(data []DomainSubmissionFeedItem, ) *GetPublicContestsByIdSubmissionFeed200Response`

NewGetPublicContestsByIdSubmissionFeed200Response instantiates a new GetPublicContestsByIdSubmissionFeed200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicContestsByIdSubmissionFeed200ResponseWithDefaults

`func NewGetPublicContestsByIdSubmissionFeed200ResponseWithDefaults() *GetPublicContestsByIdSubmissionFeed200Response`

NewGetPublicContestsByIdSubmissionFeed200ResponseWithDefaults instantiates a new GetPublicContestsByIdSubmissionFeed200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPublicContestsByIdSubmissionFeed200Response) GetData() []DomainSubmissionFeedItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPublicContestsByIdSubmissionFeed200Response) GetDataOk() (*[]DomainSubmissionFeedItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPublicContestsByIdSubmissionFeed200Response) SetData(v []DomainSubmissionFeedItem)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetPublicContestsByIdSubmissionFeed200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetPublicContestsByIdSubmissionFeed200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetPublicContestsByIdSubmissionFeed200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetPublicContestsByIdSubmissionFeed200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


