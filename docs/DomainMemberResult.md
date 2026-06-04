# DomainMemberResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestId** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Result** | Pointer to [**DomainParticipationResult**](DomainParticipationResult.md) |  | [optional] 
**Team** | Pointer to [**DomainTeam**](DomainTeam.md) |  | [optional] 

## Methods

### NewDomainMemberResult

`func NewDomainMemberResult() *DomainMemberResult`

NewDomainMemberResult instantiates a new DomainMemberResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainMemberResultWithDefaults

`func NewDomainMemberResultWithDefaults() *DomainMemberResult`

NewDomainMemberResultWithDefaults instantiates a new DomainMemberResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestId

`func (o *DomainMemberResult) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainMemberResult) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainMemberResult) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainMemberResult) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetGroup

`func (o *DomainMemberResult) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DomainMemberResult) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DomainMemberResult) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DomainMemberResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMember

`func (o *DomainMemberResult) GetMember() DomainPlatformIdentity`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *DomainMemberResult) GetMemberOk() (*DomainPlatformIdentity, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *DomainMemberResult) SetMember(v DomainPlatformIdentity)`

SetMember sets Member field to given value.

### HasMember

`func (o *DomainMemberResult) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetResult

`func (o *DomainMemberResult) GetResult() DomainParticipationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DomainMemberResult) GetResultOk() (*DomainParticipationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DomainMemberResult) SetResult(v DomainParticipationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DomainMemberResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTeam

`func (o *DomainMemberResult) GetTeam() DomainTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DomainMemberResult) GetTeamOk() (*DomainTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DomainMemberResult) SetTeam(v DomainTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DomainMemberResult) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


