# GetContestsByIdHostings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainContestHosting**](DomainContestHosting.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetContestsByIdHostings200Response

`func NewGetContestsByIdHostings200Response(data []DomainContestHosting, ) *GetContestsByIdHostings200Response`

NewGetContestsByIdHostings200Response instantiates a new GetContestsByIdHostings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContestsByIdHostings200ResponseWithDefaults

`func NewGetContestsByIdHostings200ResponseWithDefaults() *GetContestsByIdHostings200Response`

NewGetContestsByIdHostings200ResponseWithDefaults instantiates a new GetContestsByIdHostings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContestsByIdHostings200Response) GetData() []DomainContestHosting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContestsByIdHostings200Response) GetDataOk() (*[]DomainContestHosting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContestsByIdHostings200Response) SetData(v []DomainContestHosting)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetContestsByIdHostings200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetContestsByIdHostings200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetContestsByIdHostings200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetContestsByIdHostings200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


