# GetChannelsByIdSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSubscription**](DomainSubscription.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetChannelsByIdSubscriptions200Response

`func NewGetChannelsByIdSubscriptions200Response(data []DomainSubscription, ) *GetChannelsByIdSubscriptions200Response`

NewGetChannelsByIdSubscriptions200Response instantiates a new GetChannelsByIdSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChannelsByIdSubscriptions200ResponseWithDefaults

`func NewGetChannelsByIdSubscriptions200ResponseWithDefaults() *GetChannelsByIdSubscriptions200Response`

NewGetChannelsByIdSubscriptions200ResponseWithDefaults instantiates a new GetChannelsByIdSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetChannelsByIdSubscriptions200Response) GetData() []DomainSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChannelsByIdSubscriptions200Response) GetDataOk() (*[]DomainSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChannelsByIdSubscriptions200Response) SetData(v []DomainSubscription)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetChannelsByIdSubscriptions200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetChannelsByIdSubscriptions200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetChannelsByIdSubscriptions200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetChannelsByIdSubscriptions200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


