# DomainSponsorship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to [**DomainAmount**](DomainAmount.md) |  | [optional] 
**ContestId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PackageId** | Pointer to **string** |  | [optional] 
**PlatformFee** | Pointer to **float32** |  | [optional] 
**SponsorId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainSponsorship

`func NewDomainSponsorship() *DomainSponsorship`

NewDomainSponsorship instantiates a new DomainSponsorship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorshipWithDefaults

`func NewDomainSponsorshipWithDefaults() *DomainSponsorship`

NewDomainSponsorshipWithDefaults instantiates a new DomainSponsorship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreementId

`func (o *DomainSponsorship) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *DomainSponsorship) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *DomainSponsorship) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *DomainSponsorship) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetAmount

`func (o *DomainSponsorship) GetAmount() DomainAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DomainSponsorship) GetAmountOk() (*DomainAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DomainSponsorship) SetAmount(v DomainAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DomainSponsorship) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetContestId

`func (o *DomainSponsorship) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainSponsorship) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainSponsorship) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainSponsorship) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSponsorship) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSponsorship) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSponsorship) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainSponsorship) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainSponsorship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSponsorship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSponsorship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainSponsorship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackageId

`func (o *DomainSponsorship) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *DomainSponsorship) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *DomainSponsorship) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *DomainSponsorship) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPlatformFee

`func (o *DomainSponsorship) GetPlatformFee() float32`

GetPlatformFee returns the PlatformFee field if non-nil, zero value otherwise.

### GetPlatformFeeOk

`func (o *DomainSponsorship) GetPlatformFeeOk() (*float32, bool)`

GetPlatformFeeOk returns a tuple with the PlatformFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformFee

`func (o *DomainSponsorship) SetPlatformFee(v float32)`

SetPlatformFee sets PlatformFee field to given value.

### HasPlatformFee

`func (o *DomainSponsorship) HasPlatformFee() bool`

HasPlatformFee returns a boolean if a field has been set.

### GetSponsorId

`func (o *DomainSponsorship) GetSponsorId() string`

GetSponsorId returns the SponsorId field if non-nil, zero value otherwise.

### GetSponsorIdOk

`func (o *DomainSponsorship) GetSponsorIdOk() (*string, bool)`

GetSponsorIdOk returns a tuple with the SponsorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorId

`func (o *DomainSponsorship) SetSponsorId(v string)`

SetSponsorId sets SponsorId field to given value.

### HasSponsorId

`func (o *DomainSponsorship) HasSponsorId() bool`

HasSponsorId returns a boolean if a field has been set.

### GetState

`func (o *DomainSponsorship) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainSponsorship) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainSponsorship) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainSponsorship) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DomainSponsorship) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DomainSponsorship) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DomainSponsorship) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DomainSponsorship) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


