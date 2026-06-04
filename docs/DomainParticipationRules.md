# DomainParticipationRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximumParticipation** | Pointer to [**DomainIntervalOccurrence**](DomainIntervalOccurrence.md) |  | [optional] 
**MaximumParticipationWon** | Pointer to [**DomainIntervalOccurrence**](DomainIntervalOccurrence.md) |  | [optional] 
**MemberCanEnroll** | Pointer to **bool** |  | [optional] 
**MinimumParticipation** | Pointer to [**DomainIntervalOccurrence**](DomainIntervalOccurrence.md) |  | [optional] 
**MinimumParticipationWon** | Pointer to [**DomainIntervalOccurrence**](DomainIntervalOccurrence.md) |  | [optional] 
**NewMembersOnly** | Pointer to **bool** |  | [optional] 
**Preset** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainParticipationRules

`func NewDomainParticipationRules() *DomainParticipationRules`

NewDomainParticipationRules instantiates a new DomainParticipationRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainParticipationRulesWithDefaults

`func NewDomainParticipationRulesWithDefaults() *DomainParticipationRules`

NewDomainParticipationRulesWithDefaults instantiates a new DomainParticipationRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestId

`func (o *DomainParticipationRules) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainParticipationRules) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainParticipationRules) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainParticipationRules) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetDescription

`func (o *DomainParticipationRules) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainParticipationRules) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainParticipationRules) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainParticipationRules) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DomainParticipationRules) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainParticipationRules) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainParticipationRules) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainParticipationRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximumParticipation

`func (o *DomainParticipationRules) GetMaximumParticipation() DomainIntervalOccurrence`

GetMaximumParticipation returns the MaximumParticipation field if non-nil, zero value otherwise.

### GetMaximumParticipationOk

`func (o *DomainParticipationRules) GetMaximumParticipationOk() (*DomainIntervalOccurrence, bool)`

GetMaximumParticipationOk returns a tuple with the MaximumParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumParticipation

`func (o *DomainParticipationRules) SetMaximumParticipation(v DomainIntervalOccurrence)`

SetMaximumParticipation sets MaximumParticipation field to given value.

### HasMaximumParticipation

`func (o *DomainParticipationRules) HasMaximumParticipation() bool`

HasMaximumParticipation returns a boolean if a field has been set.

### GetMaximumParticipationWon

`func (o *DomainParticipationRules) GetMaximumParticipationWon() DomainIntervalOccurrence`

GetMaximumParticipationWon returns the MaximumParticipationWon field if non-nil, zero value otherwise.

### GetMaximumParticipationWonOk

`func (o *DomainParticipationRules) GetMaximumParticipationWonOk() (*DomainIntervalOccurrence, bool)`

GetMaximumParticipationWonOk returns a tuple with the MaximumParticipationWon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumParticipationWon

`func (o *DomainParticipationRules) SetMaximumParticipationWon(v DomainIntervalOccurrence)`

SetMaximumParticipationWon sets MaximumParticipationWon field to given value.

### HasMaximumParticipationWon

`func (o *DomainParticipationRules) HasMaximumParticipationWon() bool`

HasMaximumParticipationWon returns a boolean if a field has been set.

### GetMemberCanEnroll

`func (o *DomainParticipationRules) GetMemberCanEnroll() bool`

GetMemberCanEnroll returns the MemberCanEnroll field if non-nil, zero value otherwise.

### GetMemberCanEnrollOk

`func (o *DomainParticipationRules) GetMemberCanEnrollOk() (*bool, bool)`

GetMemberCanEnrollOk returns a tuple with the MemberCanEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCanEnroll

`func (o *DomainParticipationRules) SetMemberCanEnroll(v bool)`

SetMemberCanEnroll sets MemberCanEnroll field to given value.

### HasMemberCanEnroll

`func (o *DomainParticipationRules) HasMemberCanEnroll() bool`

HasMemberCanEnroll returns a boolean if a field has been set.

### GetMinimumParticipation

`func (o *DomainParticipationRules) GetMinimumParticipation() DomainIntervalOccurrence`

GetMinimumParticipation returns the MinimumParticipation field if non-nil, zero value otherwise.

### GetMinimumParticipationOk

`func (o *DomainParticipationRules) GetMinimumParticipationOk() (*DomainIntervalOccurrence, bool)`

GetMinimumParticipationOk returns a tuple with the MinimumParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumParticipation

`func (o *DomainParticipationRules) SetMinimumParticipation(v DomainIntervalOccurrence)`

SetMinimumParticipation sets MinimumParticipation field to given value.

### HasMinimumParticipation

`func (o *DomainParticipationRules) HasMinimumParticipation() bool`

HasMinimumParticipation returns a boolean if a field has been set.

### GetMinimumParticipationWon

`func (o *DomainParticipationRules) GetMinimumParticipationWon() DomainIntervalOccurrence`

GetMinimumParticipationWon returns the MinimumParticipationWon field if non-nil, zero value otherwise.

### GetMinimumParticipationWonOk

`func (o *DomainParticipationRules) GetMinimumParticipationWonOk() (*DomainIntervalOccurrence, bool)`

GetMinimumParticipationWonOk returns a tuple with the MinimumParticipationWon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumParticipationWon

`func (o *DomainParticipationRules) SetMinimumParticipationWon(v DomainIntervalOccurrence)`

SetMinimumParticipationWon sets MinimumParticipationWon field to given value.

### HasMinimumParticipationWon

`func (o *DomainParticipationRules) HasMinimumParticipationWon() bool`

HasMinimumParticipationWon returns a boolean if a field has been set.

### GetNewMembersOnly

`func (o *DomainParticipationRules) GetNewMembersOnly() bool`

GetNewMembersOnly returns the NewMembersOnly field if non-nil, zero value otherwise.

### GetNewMembersOnlyOk

`func (o *DomainParticipationRules) GetNewMembersOnlyOk() (*bool, bool)`

GetNewMembersOnlyOk returns a tuple with the NewMembersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMembersOnly

`func (o *DomainParticipationRules) SetNewMembersOnly(v bool)`

SetNewMembersOnly sets NewMembersOnly field to given value.

### HasNewMembersOnly

`func (o *DomainParticipationRules) HasNewMembersOnly() bool`

HasNewMembersOnly returns a boolean if a field has been set.

### GetPreset

`func (o *DomainParticipationRules) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *DomainParticipationRules) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *DomainParticipationRules) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *DomainParticipationRules) HasPreset() bool`

HasPreset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


