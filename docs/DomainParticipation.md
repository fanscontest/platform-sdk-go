# DomainParticipation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** |  | [optional] 
**Contest** | Pointer to [**DomainContest**](DomainContest.md) |  | [optional] 
**ContestGroupName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedAt** | Pointer to **string** |  | [optional] 
**Participant** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**PlatformIdentityId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**DomainParticipationResult**](DomainParticipationResult.md) |  | [optional] 
**Solutions** | Pointer to **[]int32** | Unified field - can be []Solution or []Prediction | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubmissionNumber** | Pointer to **int32** |  | [optional] 
**Team** | Pointer to [**DomainTeam**](DomainTeam.md) |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainParticipation

`func NewDomainParticipation() *DomainParticipation`

NewDomainParticipation instantiates a new DomainParticipation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainParticipationWithDefaults

`func NewDomainParticipationWithDefaults() *DomainParticipation`

NewDomainParticipationWithDefaults instantiates a new DomainParticipation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *DomainParticipation) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DomainParticipation) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DomainParticipation) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *DomainParticipation) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetContest

`func (o *DomainParticipation) GetContest() DomainContest`

GetContest returns the Contest field if non-nil, zero value otherwise.

### GetContestOk

`func (o *DomainParticipation) GetContestOk() (*DomainContest, bool)`

GetContestOk returns a tuple with the Contest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContest

`func (o *DomainParticipation) SetContest(v DomainContest)`

SetContest sets Contest field to given value.

### HasContest

`func (o *DomainParticipation) HasContest() bool`

HasContest returns a boolean if a field has been set.

### GetContestGroupName

`func (o *DomainParticipation) GetContestGroupName() string`

GetContestGroupName returns the ContestGroupName field if non-nil, zero value otherwise.

### GetContestGroupNameOk

`func (o *DomainParticipation) GetContestGroupNameOk() (*string, bool)`

GetContestGroupNameOk returns a tuple with the ContestGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestGroupName

`func (o *DomainParticipation) SetContestGroupName(v string)`

SetContestGroupName sets ContestGroupName field to given value.

### HasContestGroupName

`func (o *DomainParticipation) HasContestGroupName() bool`

HasContestGroupName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainParticipation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainParticipation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainParticipation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainParticipation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainParticipation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainParticipation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainParticipation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainParticipation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *DomainParticipation) GetLastModifiedAt() string`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *DomainParticipation) GetLastModifiedAtOk() (*string, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *DomainParticipation) SetLastModifiedAt(v string)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *DomainParticipation) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetParticipant

`func (o *DomainParticipation) GetParticipant() DomainPlatformIdentity`

GetParticipant returns the Participant field if non-nil, zero value otherwise.

### GetParticipantOk

`func (o *DomainParticipation) GetParticipantOk() (*DomainPlatformIdentity, bool)`

GetParticipantOk returns a tuple with the Participant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipant

`func (o *DomainParticipation) SetParticipant(v DomainPlatformIdentity)`

SetParticipant sets Participant field to given value.

### HasParticipant

`func (o *DomainParticipation) HasParticipant() bool`

HasParticipant returns a boolean if a field has been set.

### GetPlatformIdentityId

`func (o *DomainParticipation) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *DomainParticipation) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *DomainParticipation) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.

### HasPlatformIdentityId

`func (o *DomainParticipation) HasPlatformIdentityId() bool`

HasPlatformIdentityId returns a boolean if a field has been set.

### GetResult

`func (o *DomainParticipation) GetResult() DomainParticipationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DomainParticipation) GetResultOk() (*DomainParticipationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DomainParticipation) SetResult(v DomainParticipationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DomainParticipation) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSolutions

`func (o *DomainParticipation) GetSolutions() []int32`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *DomainParticipation) GetSolutionsOk() (*[]int32, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *DomainParticipation) SetSolutions(v []int32)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *DomainParticipation) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetStatus

`func (o *DomainParticipation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainParticipation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainParticipation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainParticipation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmissionNumber

`func (o *DomainParticipation) GetSubmissionNumber() int32`

GetSubmissionNumber returns the SubmissionNumber field if non-nil, zero value otherwise.

### GetSubmissionNumberOk

`func (o *DomainParticipation) GetSubmissionNumberOk() (*int32, bool)`

GetSubmissionNumberOk returns a tuple with the SubmissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionNumber

`func (o *DomainParticipation) SetSubmissionNumber(v int32)`

SetSubmissionNumber sets SubmissionNumber field to given value.

### HasSubmissionNumber

`func (o *DomainParticipation) HasSubmissionNumber() bool`

HasSubmissionNumber returns a boolean if a field has been set.

### GetTeam

`func (o *DomainParticipation) GetTeam() DomainTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DomainParticipation) GetTeamOk() (*DomainTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DomainParticipation) SetTeam(v DomainTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DomainParticipation) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTime

`func (o *DomainParticipation) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DomainParticipation) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DomainParticipation) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *DomainParticipation) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


