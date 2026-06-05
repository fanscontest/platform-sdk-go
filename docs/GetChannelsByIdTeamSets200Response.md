# GetChannelsByIdTeamSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainTeamSet**](DomainTeamSet.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetChannelsByIdTeamSets200Response

`func NewGetChannelsByIdTeamSets200Response(data []DomainTeamSet, ) *GetChannelsByIdTeamSets200Response`

NewGetChannelsByIdTeamSets200Response instantiates a new GetChannelsByIdTeamSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannelsByIdTeamSets200ResponseWithDefaults

`func NewGetChannelsByIdTeamSets200ResponseWithDefaults() *GetChannelsByIdTeamSets200Response`

NewGetChannelsByIdTeamSets200ResponseWithDefaults instantiates a new GetChannelsByIdTeamSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetChannelsByIdTeamSets200Response) GetData() []DomainTeamSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChannelsByIdTeamSets200Response) GetDataOk() (*[]DomainTeamSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChannelsByIdTeamSets200Response) SetData(v []DomainTeamSet)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetChannelsByIdTeamSets200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetChannelsByIdTeamSets200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetChannelsByIdTeamSets200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetChannelsByIdTeamSets200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


