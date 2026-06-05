# GetContestsByContestIdSponsorships200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainSponsorship**](DomainSponsorship.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetContestsByContestIdSponsorships200Response

`func NewGetContestsByContestIdSponsorships200Response(data []DomainSponsorship, ) *GetContestsByContestIdSponsorships200Response`

NewGetContestsByContestIdSponsorships200Response instantiates a new GetContestsByContestIdSponsorships200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContestsByContestIdSponsorships200ResponseWithDefaults

`func NewGetContestsByContestIdSponsorships200ResponseWithDefaults() *GetContestsByContestIdSponsorships200Response`

NewGetContestsByContestIdSponsorships200ResponseWithDefaults instantiates a new GetContestsByContestIdSponsorships200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContestsByContestIdSponsorships200Response) GetData() []DomainSponsorship`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContestsByContestIdSponsorships200Response) GetDataOk() (*[]DomainSponsorship, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContestsByContestIdSponsorships200Response) SetData(v []DomainSponsorship)`

SetData sets Data field to given value.


### GetPagination

`func (o *GetContestsByContestIdSponsorships200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetContestsByContestIdSponsorships200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetContestsByContestIdSponsorships200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetContestsByContestIdSponsorships200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


