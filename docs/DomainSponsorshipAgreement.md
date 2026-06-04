# DomainSponsorshipAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to [**DomainAmount**](DomainAmount.md) |  | [optional] 
**BrandingTerms** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PackageId** | Pointer to **string** | For package-based agreements | [optional] 
**SponsorshipId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TermsSnapshot** | Pointer to [**[]DomainSponsorshipTermSnapshot**](DomainSponsorshipTermSnapshot.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainSponsorshipAgreement

`func NewDomainSponsorshipAgreement() *DomainSponsorshipAgreement`

NewDomainSponsorshipAgreement instantiates a new DomainSponsorshipAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorshipAgreementWithDefaults

`func NewDomainSponsorshipAgreementWithDefaults() *DomainSponsorshipAgreement`

NewDomainSponsorshipAgreementWithDefaults instantiates a new DomainSponsorshipAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *DomainSponsorshipAgreement) GetAcceptedAt() string`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *DomainSponsorshipAgreement) GetAcceptedAtOk() (*string, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *DomainSponsorshipAgreement) SetAcceptedAt(v string)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *DomainSponsorshipAgreement) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### GetAmount

`func (o *DomainSponsorshipAgreement) GetAmount() DomainAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DomainSponsorshipAgreement) GetAmountOk() (*DomainAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DomainSponsorshipAgreement) SetAmount(v DomainAmount)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DomainSponsorshipAgreement) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBrandingTerms

`func (o *DomainSponsorshipAgreement) GetBrandingTerms() string`

GetBrandingTerms returns the BrandingTerms field if non-nil, zero value otherwise.

### GetBrandingTermsOk

`func (o *DomainSponsorshipAgreement) GetBrandingTermsOk() (*string, bool)`

GetBrandingTermsOk returns a tuple with the BrandingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTerms

`func (o *DomainSponsorshipAgreement) SetBrandingTerms(v string)`

SetBrandingTerms sets BrandingTerms field to given value.

### HasBrandingTerms

`func (o *DomainSponsorshipAgreement) HasBrandingTerms() bool`

HasBrandingTerms returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSponsorshipAgreement) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSponsorshipAgreement) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSponsorshipAgreement) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainSponsorshipAgreement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeadline

`func (o *DomainSponsorshipAgreement) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *DomainSponsorshipAgreement) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *DomainSponsorshipAgreement) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *DomainSponsorshipAgreement) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetId

`func (o *DomainSponsorshipAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSponsorshipAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSponsorshipAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainSponsorshipAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackageId

`func (o *DomainSponsorshipAgreement) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *DomainSponsorshipAgreement) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *DomainSponsorshipAgreement) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *DomainSponsorshipAgreement) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetSponsorshipId

`func (o *DomainSponsorshipAgreement) GetSponsorshipId() string`

GetSponsorshipId returns the SponsorshipId field if non-nil, zero value otherwise.

### GetSponsorshipIdOk

`func (o *DomainSponsorshipAgreement) GetSponsorshipIdOk() (*string, bool)`

GetSponsorshipIdOk returns a tuple with the SponsorshipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorshipId

`func (o *DomainSponsorshipAgreement) SetSponsorshipId(v string)`

SetSponsorshipId sets SponsorshipId field to given value.

### HasSponsorshipId

`func (o *DomainSponsorshipAgreement) HasSponsorshipId() bool`

HasSponsorshipId returns a boolean if a field has been set.

### GetState

`func (o *DomainSponsorshipAgreement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainSponsorshipAgreement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainSponsorshipAgreement) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainSponsorshipAgreement) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTermsSnapshot

`func (o *DomainSponsorshipAgreement) GetTermsSnapshot() []DomainSponsorshipTermSnapshot`

GetTermsSnapshot returns the TermsSnapshot field if non-nil, zero value otherwise.

### GetTermsSnapshotOk

`func (o *DomainSponsorshipAgreement) GetTermsSnapshotOk() (*[]DomainSponsorshipTermSnapshot, bool)`

GetTermsSnapshotOk returns a tuple with the TermsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsSnapshot

`func (o *DomainSponsorshipAgreement) SetTermsSnapshot(v []DomainSponsorshipTermSnapshot)`

SetTermsSnapshot sets TermsSnapshot field to given value.

### HasTermsSnapshot

`func (o *DomainSponsorshipAgreement) HasTermsSnapshot() bool`

HasTermsSnapshot returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DomainSponsorshipAgreement) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DomainSponsorshipAgreement) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DomainSponsorshipAgreement) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DomainSponsorshipAgreement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


