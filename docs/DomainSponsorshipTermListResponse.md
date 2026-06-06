# DomainSponsorshipTermListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSponsorshipTerm**](DomainSponsorshipTerm.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainSponsorshipTermListResponse

`func NewDomainSponsorshipTermListResponse(data []DomainSponsorshipTerm, ) *DomainSponsorshipTermListResponse`

NewDomainSponsorshipTermListResponse instantiates a new DomainSponsorshipTermListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorshipTermListResponseWithDefaults

`func NewDomainSponsorshipTermListResponseWithDefaults() *DomainSponsorshipTermListResponse`

NewDomainSponsorshipTermListResponseWithDefaults instantiates a new DomainSponsorshipTermListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainSponsorshipTermListResponse) GetData() []DomainSponsorshipTerm`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainSponsorshipTermListResponse) GetDataOk() (*[]DomainSponsorshipTerm, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainSponsorshipTermListResponse) SetData(v []DomainSponsorshipTerm)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainSponsorshipTermListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainSponsorshipTermListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainSponsorshipTermListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainSponsorshipTermListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


