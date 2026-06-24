# DomainPlatformIdentityListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainPlatformIdentityListResponse

`func NewDomainPlatformIdentityListResponse(data []DomainPlatformIdentity, ) *DomainPlatformIdentityListResponse`

NewDomainPlatformIdentityListResponse instantiates a new DomainPlatformIdentityListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPlatformIdentityListResponseWithDefaults

`func NewDomainPlatformIdentityListResponseWithDefaults() *DomainPlatformIdentityListResponse`

NewDomainPlatformIdentityListResponseWithDefaults instantiates a new DomainPlatformIdentityListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainPlatformIdentityListResponse) GetData() []DomainPlatformIdentity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainPlatformIdentityListResponse) GetDataOk() (*[]DomainPlatformIdentity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainPlatformIdentityListResponse) SetData(v []DomainPlatformIdentity)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainPlatformIdentityListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainPlatformIdentityListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainPlatformIdentityListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainPlatformIdentityListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


