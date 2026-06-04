# DomainSponsorshipPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedContestTopics** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxContests** | Pointer to **int32** |  | [optional] 
**RemainingBudget** | Pointer to [**DomainAmount**](DomainAmount.md) |  | [optional] 
**SponsorId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TotalBudget** | Pointer to [**DomainAmount**](DomainAmount.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainSponsorshipPackage

`func NewDomainSponsorshipPackage() *DomainSponsorshipPackage`

NewDomainSponsorshipPackage instantiates a new DomainSponsorshipPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSponsorshipPackageWithDefaults

`func NewDomainSponsorshipPackageWithDefaults() *DomainSponsorshipPackage`

NewDomainSponsorshipPackageWithDefaults instantiates a new DomainSponsorshipPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedContestTopics

`func (o *DomainSponsorshipPackage) GetAllowedContestTopics() []string`

GetAllowedContestTopics returns the AllowedContestTopics field if non-nil, zero value otherwise.

### GetAllowedContestTopicsOk

`func (o *DomainSponsorshipPackage) GetAllowedContestTopicsOk() (*[]string, bool)`

GetAllowedContestTopicsOk returns a tuple with the AllowedContestTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContestTopics

`func (o *DomainSponsorshipPackage) SetAllowedContestTopics(v []string)`

SetAllowedContestTopics sets AllowedContestTopics field to given value.

### HasAllowedContestTopics

`func (o *DomainSponsorshipPackage) HasAllowedContestTopics() bool`

HasAllowedContestTopics returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSponsorshipPackage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSponsorshipPackage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSponsorshipPackage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainSponsorshipPackage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainSponsorshipPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSponsorshipPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSponsorshipPackage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainSponsorshipPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxContests

`func (o *DomainSponsorshipPackage) GetMaxContests() int32`

GetMaxContests returns the MaxContests field if non-nil, zero value otherwise.

### GetMaxContestsOk

`func (o *DomainSponsorshipPackage) GetMaxContestsOk() (*int32, bool)`

GetMaxContestsOk returns a tuple with the MaxContests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContests

`func (o *DomainSponsorshipPackage) SetMaxContests(v int32)`

SetMaxContests sets MaxContests field to given value.

### HasMaxContests

`func (o *DomainSponsorshipPackage) HasMaxContests() bool`

HasMaxContests returns a boolean if a field has been set.

### GetRemainingBudget

`func (o *DomainSponsorshipPackage) GetRemainingBudget() DomainAmount`

GetRemainingBudget returns the RemainingBudget field if non-nil, zero value otherwise.

### GetRemainingBudgetOk

`func (o *DomainSponsorshipPackage) GetRemainingBudgetOk() (*DomainAmount, bool)`

GetRemainingBudgetOk returns a tuple with the RemainingBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBudget

`func (o *DomainSponsorshipPackage) SetRemainingBudget(v DomainAmount)`

SetRemainingBudget sets RemainingBudget field to given value.

### HasRemainingBudget

`func (o *DomainSponsorshipPackage) HasRemainingBudget() bool`

HasRemainingBudget returns a boolean if a field has been set.

### GetSponsorId

`func (o *DomainSponsorshipPackage) GetSponsorId() string`

GetSponsorId returns the SponsorId field if non-nil, zero value otherwise.

### GetSponsorIdOk

`func (o *DomainSponsorshipPackage) GetSponsorIdOk() (*string, bool)`

GetSponsorIdOk returns a tuple with the SponsorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorId

`func (o *DomainSponsorshipPackage) SetSponsorId(v string)`

SetSponsorId sets SponsorId field to given value.

### HasSponsorId

`func (o *DomainSponsorshipPackage) HasSponsorId() bool`

HasSponsorId returns a boolean if a field has been set.

### GetState

`func (o *DomainSponsorshipPackage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainSponsorshipPackage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainSponsorshipPackage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainSponsorshipPackage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTotalBudget

`func (o *DomainSponsorshipPackage) GetTotalBudget() DomainAmount`

GetTotalBudget returns the TotalBudget field if non-nil, zero value otherwise.

### GetTotalBudgetOk

`func (o *DomainSponsorshipPackage) GetTotalBudgetOk() (*DomainAmount, bool)`

GetTotalBudgetOk returns a tuple with the TotalBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBudget

`func (o *DomainSponsorshipPackage) SetTotalBudget(v DomainAmount)`

SetTotalBudget sets TotalBudget field to given value.

### HasTotalBudget

`func (o *DomainSponsorshipPackage) HasTotalBudget() bool`

HasTotalBudget returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DomainSponsorshipPackage) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DomainSponsorshipPackage) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DomainSponsorshipPackage) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DomainSponsorshipPackage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


