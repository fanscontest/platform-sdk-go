# DomainGroupedScoreBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contest** | Pointer to [**DomainContest**](DomainContest.md) |  | [optional] 
**MemberParticipation** | Pointer to [**DomainParticipation**](DomainParticipation.md) |  | [optional] 
**ScoresByGroup** | Pointer to [**map[string][]DomainMemberResult**](array.md) |  | [optional] 

## Methods

### NewDomainGroupedScoreBoard

`func NewDomainGroupedScoreBoard() *DomainGroupedScoreBoard`

NewDomainGroupedScoreBoard instantiates a new DomainGroupedScoreBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainGroupedScoreBoardWithDefaults

`func NewDomainGroupedScoreBoardWithDefaults() *DomainGroupedScoreBoard`

NewDomainGroupedScoreBoardWithDefaults instantiates a new DomainGroupedScoreBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContest

`func (o *DomainGroupedScoreBoard) GetContest() DomainContest`

GetContest returns the Contest field if non-nil, zero value otherwise.

### GetContestOk

`func (o *DomainGroupedScoreBoard) GetContestOk() (*DomainContest, bool)`

GetContestOk returns a tuple with the Contest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContest

`func (o *DomainGroupedScoreBoard) SetContest(v DomainContest)`

SetContest sets Contest field to given value.

### HasContest

`func (o *DomainGroupedScoreBoard) HasContest() bool`

HasContest returns a boolean if a field has been set.

### GetMemberParticipation

`func (o *DomainGroupedScoreBoard) GetMemberParticipation() DomainParticipation`

GetMemberParticipation returns the MemberParticipation field if non-nil, zero value otherwise.

### GetMemberParticipationOk

`func (o *DomainGroupedScoreBoard) GetMemberParticipationOk() (*DomainParticipation, bool)`

GetMemberParticipationOk returns a tuple with the MemberParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberParticipation

`func (o *DomainGroupedScoreBoard) SetMemberParticipation(v DomainParticipation)`

SetMemberParticipation sets MemberParticipation field to given value.

### HasMemberParticipation

`func (o *DomainGroupedScoreBoard) HasMemberParticipation() bool`

HasMemberParticipation returns a boolean if a field has been set.

### GetScoresByGroup

`func (o *DomainGroupedScoreBoard) GetScoresByGroup() map[string][]DomainMemberResult`

GetScoresByGroup returns the ScoresByGroup field if non-nil, zero value otherwise.

### GetScoresByGroupOk

`func (o *DomainGroupedScoreBoard) GetScoresByGroupOk() (*map[string][]DomainMemberResult, bool)`

GetScoresByGroupOk returns a tuple with the ScoresByGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoresByGroup

`func (o *DomainGroupedScoreBoard) SetScoresByGroup(v map[string][]DomainMemberResult)`

SetScoresByGroup sets ScoresByGroup field to given value.

### HasScoresByGroup

`func (o *DomainGroupedScoreBoard) HasScoresByGroup() bool`

HasScoresByGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


