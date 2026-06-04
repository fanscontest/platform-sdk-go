# DomainBuddyBoardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuddyBoard** | Pointer to [**DomainBuddyBoard**](DomainBuddyBoard.md) |  | [optional] 
**Contest** | Pointer to [**DomainContest**](DomainContest.md) |  | [optional] 
**MemberParticipation** | Pointer to [**DomainParticipation**](DomainParticipation.md) |  | [optional] 
**Scores** | Pointer to [**[]DomainMemberResult**](DomainMemberResult.md) |  | [optional] 
**Winners** | Pointer to [**[]DomainMemberResult**](DomainMemberResult.md) |  | [optional] 

## Methods

### NewDomainBuddyBoardView

`func NewDomainBuddyBoardView() *DomainBuddyBoardView`

NewDomainBuddyBoardView instantiates a new DomainBuddyBoardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBuddyBoardViewWithDefaults

`func NewDomainBuddyBoardViewWithDefaults() *DomainBuddyBoardView`

NewDomainBuddyBoardViewWithDefaults instantiates a new DomainBuddyBoardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuddyBoard

`func (o *DomainBuddyBoardView) GetBuddyBoard() DomainBuddyBoard`

GetBuddyBoard returns the BuddyBoard field if non-nil, zero value otherwise.

### GetBuddyBoardOk

`func (o *DomainBuddyBoardView) GetBuddyBoardOk() (*DomainBuddyBoard, bool)`

GetBuddyBoardOk returns a tuple with the BuddyBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuddyBoard

`func (o *DomainBuddyBoardView) SetBuddyBoard(v DomainBuddyBoard)`

SetBuddyBoard sets BuddyBoard field to given value.

### HasBuddyBoard

`func (o *DomainBuddyBoardView) HasBuddyBoard() bool`

HasBuddyBoard returns a boolean if a field has been set.

### GetContest

`func (o *DomainBuddyBoardView) GetContest() DomainContest`

GetContest returns the Contest field if non-nil, zero value otherwise.

### GetContestOk

`func (o *DomainBuddyBoardView) GetContestOk() (*DomainContest, bool)`

GetContestOk returns a tuple with the Contest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContest

`func (o *DomainBuddyBoardView) SetContest(v DomainContest)`

SetContest sets Contest field to given value.

### HasContest

`func (o *DomainBuddyBoardView) HasContest() bool`

HasContest returns a boolean if a field has been set.

### GetMemberParticipation

`func (o *DomainBuddyBoardView) GetMemberParticipation() DomainParticipation`

GetMemberParticipation returns the MemberParticipation field if non-nil, zero value otherwise.

### GetMemberParticipationOk

`func (o *DomainBuddyBoardView) GetMemberParticipationOk() (*DomainParticipation, bool)`

GetMemberParticipationOk returns a tuple with the MemberParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberParticipation

`func (o *DomainBuddyBoardView) SetMemberParticipation(v DomainParticipation)`

SetMemberParticipation sets MemberParticipation field to given value.

### HasMemberParticipation

`func (o *DomainBuddyBoardView) HasMemberParticipation() bool`

HasMemberParticipation returns a boolean if a field has been set.

### GetScores

`func (o *DomainBuddyBoardView) GetScores() []DomainMemberResult`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *DomainBuddyBoardView) GetScoresOk() (*[]DomainMemberResult, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *DomainBuddyBoardView) SetScores(v []DomainMemberResult)`

SetScores sets Scores field to given value.

### HasScores

`func (o *DomainBuddyBoardView) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetWinners

`func (o *DomainBuddyBoardView) GetWinners() []DomainMemberResult`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *DomainBuddyBoardView) GetWinnersOk() (*[]DomainMemberResult, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *DomainBuddyBoardView) SetWinners(v []DomainMemberResult)`

SetWinners sets Winners field to given value.

### HasWinners

`func (o *DomainBuddyBoardView) HasWinners() bool`

HasWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


