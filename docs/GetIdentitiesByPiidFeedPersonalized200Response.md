# GetIdentitiesByPiidFeedPersonalized200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainFeedItem**](DomainFeedItem.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetIdentitiesByPiidFeedPersonalized200Response

`func NewGetIdentitiesByPiidFeedPersonalized200Response(data []DomainFeedItem, ) *GetIdentitiesByPiidFeedPersonalized200Response`

NewGetIdentitiesByPiidFeedPersonalized200Response instantiates a new GetIdentitiesByPiidFeedPersonalized200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentitiesByPiidFeedPersonalized200ResponseWithDefaults

`func NewGetIdentitiesByPiidFeedPersonalized200ResponseWithDefaults() *GetIdentitiesByPiidFeedPersonalized200Response`

NewGetIdentitiesByPiidFeedPersonalized200ResponseWithDefaults instantiates a new GetIdentitiesByPiidFeedPersonalized200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) GetData() []DomainFeedItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) GetDataOk() (*[]DomainFeedItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) SetData(v []DomainFeedItem)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetIdentitiesByPiidFeedPersonalized200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


