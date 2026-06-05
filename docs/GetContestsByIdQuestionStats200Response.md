# GetContestsByIdQuestionStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainQuestionStats**](DomainQuestionStats.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetContestsByIdQuestionStats200Response

`func NewGetContestsByIdQuestionStats200Response(data []DomainQuestionStats, ) *GetContestsByIdQuestionStats200Response`

NewGetContestsByIdQuestionStats200Response instantiates a new GetContestsByIdQuestionStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContestsByIdQuestionStats200ResponseWithDefaults

`func NewGetContestsByIdQuestionStats200ResponseWithDefaults() *GetContestsByIdQuestionStats200Response`

NewGetContestsByIdQuestionStats200ResponseWithDefaults instantiates a new GetContestsByIdQuestionStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContestsByIdQuestionStats200Response) GetData() []DomainQuestionStats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContestsByIdQuestionStats200Response) GetDataOk() (*[]DomainQuestionStats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContestsByIdQuestionStats200Response) SetData(v []DomainQuestionStats)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetContestsByIdQuestionStats200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetContestsByIdQuestionStats200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetContestsByIdQuestionStats200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetContestsByIdQuestionStats200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


