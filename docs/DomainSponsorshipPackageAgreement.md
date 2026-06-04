# DomainSponsorshipPackageAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceDeadline** | Pointer to **string** |  | [optional] 
**AllowedContestTopics** | Pointer to **[]string** |  | [optional] 
**BrandingTerms** | Pointer to **string** |  | [optional] 
**BudgetPerContest** | Pointer to [**DomainAmount**](DomainAmount.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PackageId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TermsSnapshot** | Pointer to [**[]DomainSponsorshipTermSnapshot**](DomainSponsorshipTermSnapshot.md) |  | [optional] 
**TotalContests** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainSponsorshipPackageAgreement

`func NewDomainSponsorshipPackageAgreement() *DomainSponsorshipPackageAgreement`

NewDomainSponsorshipPackageAgreement instantiates a new DomainSponsorshipPackageAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorshipPackageAgreementWithDefaults

`func NewDomainSponsorshipPackageAgreementWithDefaults() *DomainSponsorshipPackageAgreement`

NewDomainSponsorshipPackageAgreementWithDefaults instantiates a new DomainSponsorshipPackageAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceDeadline

`func (o *DomainSponsorshipPackageAgreement) GetAcceptanceDeadline() string`

GetAcceptanceDeadline returns the AcceptanceDeadline field if non-nil, zero value otherwise.

### GetAcceptanceDeadlineOk

`func (o *DomainSponsorshipPackageAgreement) GetAcceptanceDeadlineOk() (*string, bool)`

GetAcceptanceDeadlineOk returns a tuple with the AcceptanceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDeadline

`func (o *DomainSponsorshipPackageAgreement) SetAcceptanceDeadline(v string)`

SetAcceptanceDeadline sets AcceptanceDeadline field to given value.

### HasAcceptanceDeadline

`func (o *DomainSponsorshipPackageAgreement) HasAcceptanceDeadline() bool`

HasAcceptanceDeadline returns a boolean if a field has been set.

### GetAllowedContestTopics

`func (o *DomainSponsorshipPackageAgreement) GetAllowedContestTopics() []string`

GetAllowedContestTopics returns the AllowedContestTopics field if non-nil, zero value otherwise.

### GetAllowedContestTopicsOk

`func (o *DomainSponsorshipPackageAgreement) GetAllowedContestTopicsOk() (*[]string, bool)`

GetAllowedContestTopicsOk returns a tuple with the AllowedContestTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContestTopics

`func (o *DomainSponsorshipPackageAgreement) SetAllowedContestTopics(v []string)`

SetAllowedContestTopics sets AllowedContestTopics field to given value.

### HasAllowedContestTopics

`func (o *DomainSponsorshipPackageAgreement) HasAllowedContestTopics() bool`

HasAllowedContestTopics returns a boolean if a field has been set.

### GetBrandingTerms

`func (o *DomainSponsorshipPackageAgreement) GetBrandingTerms() string`

GetBrandingTerms returns the BrandingTerms field if non-nil, zero value otherwise.

### GetBrandingTermsOk

`func (o *DomainSponsorshipPackageAgreement) GetBrandingTermsOk() (*string, bool)`

GetBrandingTermsOk returns a tuple with the BrandingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTerms

`func (o *DomainSponsorshipPackageAgreement) SetBrandingTerms(v string)`

SetBrandingTerms sets BrandingTerms field to given value.

### HasBrandingTerms

`func (o *DomainSponsorshipPackageAgreement) HasBrandingTerms() bool`

HasBrandingTerms returns a boolean if a field has been set.

### GetBudgetPerContest

`func (o *DomainSponsorshipPackageAgreement) GetBudgetPerContest() DomainAmount`

GetBudgetPerContest returns the BudgetPerContest field if non-nil, zero value otherwise.

### GetBudgetPerContestOk

`func (o *DomainSponsorshipPackageAgreement) GetBudgetPerContestOk() (*DomainAmount, bool)`

GetBudgetPerContestOk returns a tuple with the BudgetPerContest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetPerContest

`func (o *DomainSponsorshipPackageAgreement) SetBudgetPerContest(v DomainAmount)`

SetBudgetPerContest sets BudgetPerContest field to given value.

### HasBudgetPerContest

`func (o *DomainSponsorshipPackageAgreement) HasBudgetPerContest() bool`

HasBudgetPerContest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSponsorshipPackageAgreement) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSponsorshipPackageAgreement) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSponsorshipPackageAgreement) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainSponsorshipPackageAgreement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainSponsorshipPackageAgreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSponsorshipPackageAgreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSponsorshipPackageAgreement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainSponsorshipPackageAgreement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackageId

`func (o *DomainSponsorshipPackageAgreement) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *DomainSponsorshipPackageAgreement) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *DomainSponsorshipPackageAgreement) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *DomainSponsorshipPackageAgreement) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetState

`func (o *DomainSponsorshipPackageAgreement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainSponsorshipPackageAgreement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainSponsorshipPackageAgreement) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainSponsorshipPackageAgreement) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTermsSnapshot

`func (o *DomainSponsorshipPackageAgreement) GetTermsSnapshot() []DomainSponsorshipTermSnapshot`

GetTermsSnapshot returns the TermsSnapshot field if non-nil, zero value otherwise.

### GetTermsSnapshotOk

`func (o *DomainSponsorshipPackageAgreement) GetTermsSnapshotOk() (*[]DomainSponsorshipTermSnapshot, bool)`

GetTermsSnapshotOk returns a tuple with the TermsSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsSnapshot

`func (o *DomainSponsorshipPackageAgreement) SetTermsSnapshot(v []DomainSponsorshipTermSnapshot)`

SetTermsSnapshot sets TermsSnapshot field to given value.

### HasTermsSnapshot

`func (o *DomainSponsorshipPackageAgreement) HasTermsSnapshot() bool`

HasTermsSnapshot returns a boolean if a field has been set.

### GetTotalContests

`func (o *DomainSponsorshipPackageAgreement) GetTotalContests() int32`

GetTotalContests returns the TotalContests field if non-nil, zero value otherwise.

### GetTotalContestsOk

`func (o *DomainSponsorshipPackageAgreement) GetTotalContestsOk() (*int32, bool)`

GetTotalContestsOk returns a tuple with the TotalContests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContests

`func (o *DomainSponsorshipPackageAgreement) SetTotalContests(v int32)`

SetTotalContests sets TotalContests field to given value.

### HasTotalContests

`func (o *DomainSponsorshipPackageAgreement) HasTotalContests() bool`

HasTotalContests returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DomainSponsorshipPackageAgreement) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DomainSponsorshipPackageAgreement) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DomainSponsorshipPackageAgreement) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DomainSponsorshipPackageAgreement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


