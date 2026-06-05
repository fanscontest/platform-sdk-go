# GetChannelsByIdContests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainContest**](DomainContest.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetChannelsByIdContests200Response

`func NewGetChannelsByIdContests200Response(data []DomainContest, ) *GetChannelsByIdContests200Response`

NewGetChannelsByIdContests200Response instantiates a new GetChannelsByIdContests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannelsByIdContests200ResponseWithDefaults

`func NewGetChannelsByIdContests200ResponseWithDefaults() *GetChannelsByIdContests200Response`

NewGetChannelsByIdContests200ResponseWithDefaults instantiates a new GetChannelsByIdContests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetChannelsByIdContests200Response) GetData() []DomainContest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChannelsByIdContests200Response) GetDataOk() (*[]DomainContest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChannelsByIdContests200Response) SetData(v []DomainContest)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetChannelsByIdContests200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetChannelsByIdContests200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetChannelsByIdContests200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetChannelsByIdContests200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


