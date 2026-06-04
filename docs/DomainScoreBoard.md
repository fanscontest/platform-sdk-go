# DomainScoreBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contest** | Pointer to [**DomainContest**](DomainContest.md) |  | [optional] 
**MemberParticipation** | Pointer to [**DomainParticipation**](DomainParticipation.md) |  | [optional] 
**Scores** | Pointer to [**[]DomainMemberResult**](DomainMemberResult.md) |  | [optional] 

## Methods

### NewDomainScoreBoard

`func NewDomainScoreBoard() *DomainScoreBoard`

NewDomainScoreBoard instantiates a new DomainScoreBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainScoreBoardWithDefaults

`func NewDomainScoreBoardWithDefaults() *DomainScoreBoard`

NewDomainScoreBoardWithDefaults instantiates a new DomainScoreBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContest

`func (o *DomainScoreBoard) GetContest() DomainContest`

GetContest returns the Contest field if non-nil, zero value otherwise.

### GetContestOk

`func (o *DomainScoreBoard) GetContestOk() (*DomainContest, bool)`

GetContestOk returns a tuple with the Contest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContest

`func (o *DomainScoreBoard) SetContest(v DomainContest)`

SetContest sets Contest field to given value.

### HasContest

`func (o *DomainScoreBoard) HasContest() bool`

HasContest returns a boolean if a field has been set.

### GetMemberParticipation

`func (o *DomainScoreBoard) GetMemberParticipation() DomainParticipation`

GetMemberParticipation returns the MemberParticipation field if non-nil, zero value otherwise.

### GetMemberParticipationOk

`func (o *DomainScoreBoard) GetMemberParticipationOk() (*DomainParticipation, bool)`

GetMemberParticipationOk returns a tuple with the MemberParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberParticipation

`func (o *DomainScoreBoard) SetMemberParticipation(v DomainParticipation)`

SetMemberParticipation sets MemberParticipation field to given value.

### HasMemberParticipation

`func (o *DomainScoreBoard) HasMemberParticipation() bool`

HasMemberParticipation returns a boolean if a field has been set.

### GetScores

`func (o *DomainScoreBoard) GetScores() []DomainMemberResult`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *DomainScoreBoard) GetScoresOk() (*[]DomainMemberResult, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *DomainScoreBoard) SetScores(v []DomainMemberResult)`

SetScores sets Scores field to given value.

### HasScores

`func (o *DomainScoreBoard) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


