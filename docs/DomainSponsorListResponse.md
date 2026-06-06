# DomainSponsorListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSponsor**](DomainSponsor.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainSponsorListResponse

`func NewDomainSponsorListResponse(data []DomainSponsor, ) *DomainSponsorListResponse`

NewDomainSponsorListResponse instantiates a new DomainSponsorListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorListResponseWithDefaults

`func NewDomainSponsorListResponseWithDefaults() *DomainSponsorListResponse`

NewDomainSponsorListResponseWithDefaults instantiates a new DomainSponsorListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainSponsorListResponse) GetData() []DomainSponsor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainSponsorListResponse) GetDataOk() (*[]DomainSponsor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainSponsorListResponse) SetData(v []DomainSponsor)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainSponsorListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainSponsorListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainSponsorListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainSponsorListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


