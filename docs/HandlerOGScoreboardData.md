# HandlerOGScoreboardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestTitle** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Leaderboard** | Pointer to [**[]HandlerGroupWinners**](HandlerGroupWinners.md) | Winners per group | [optional] 
**Metadata** | Pointer to **[]string** |  | [optional] 
**NumberOfGroups** | Pointer to **int32** |  | [optional] 
**ParticipantCount** | Pointer to **int32** |  | [optional] 
**TotalWinners** | Pointer to **int32** |  | [optional] 
**WinnersPerGroup** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlerOGScoreboardData

`func NewHandlerOGScoreboardData() *HandlerOGScoreboardData`

NewHandlerOGScoreboardData instantiates a new HandlerOGScoreboardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerOGScoreboardDataWithDefaults

`func NewHandlerOGScoreboardDataWithDefaults() *HandlerOGScoreboardData`

NewHandlerOGScoreboardDataWithDefaults instantiates a new HandlerOGScoreboardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestTitle

`func (o *HandlerOGScoreboardData) GetContestTitle() string`

GetContestTitle returns the ContestTitle field if non-nil, zero value otherwise.

### GetContestTitleOk

`func (o *HandlerOGScoreboardData) GetContestTitleOk() (*string, bool)`

GetContestTitleOk returns a tuple with the ContestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestTitle

`func (o *HandlerOGScoreboardData) SetContestTitle(v string)`

SetContestTitle sets ContestTitle field to given value.

### HasContestTitle

`func (o *HandlerOGScoreboardData) HasContestTitle() bool`

HasContestTitle returns a boolean if a field has been set.

### GetCreator

`func (o *HandlerOGScoreboardData) GetCreator() DomainPlatformIdentity`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *HandlerOGScoreboardData) GetCreatorOk() (*DomainPlatformIdentity, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *HandlerOGScoreboardData) SetCreator(v DomainPlatformIdentity)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *HandlerOGScoreboardData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLeaderboard

`func (o *HandlerOGScoreboardData) GetLeaderboard() []HandlerGroupWinners`

GetLeaderboard returns the Leaderboard field if non-nil, zero value otherwise.

### GetLeaderboardOk

`func (o *HandlerOGScoreboardData) GetLeaderboardOk() (*[]HandlerGroupWinners, bool)`

GetLeaderboardOk returns a tuple with the Leaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboard

`func (o *HandlerOGScoreboardData) SetLeaderboard(v []HandlerGroupWinners)`

SetLeaderboard sets Leaderboard field to given value.

### HasLeaderboard

`func (o *HandlerOGScoreboardData) HasLeaderboard() bool`

HasLeaderboard returns a boolean if a field has been set.

### GetMetadata

`func (o *HandlerOGScoreboardData) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HandlerOGScoreboardData) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HandlerOGScoreboardData) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HandlerOGScoreboardData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNumberOfGroups

`func (o *HandlerOGScoreboardData) GetNumberOfGroups() int32`

GetNumberOfGroups returns the NumberOfGroups field if non-nil, zero value otherwise.

### GetNumberOfGroupsOk

`func (o *HandlerOGScoreboardData) GetNumberOfGroupsOk() (*int32, bool)`

GetNumberOfGroupsOk returns a tuple with the NumberOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGroups

`func (o *HandlerOGScoreboardData) SetNumberOfGroups(v int32)`

SetNumberOfGroups sets NumberOfGroups field to given value.

### HasNumberOfGroups

`func (o *HandlerOGScoreboardData) HasNumberOfGroups() bool`

HasNumberOfGroups returns a boolean if a field has been set.

### GetParticipantCount

`func (o *HandlerOGScoreboardData) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *HandlerOGScoreboardData) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *HandlerOGScoreboardData) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *HandlerOGScoreboardData) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetTotalWinners

`func (o *HandlerOGScoreboardData) GetTotalWinners() int32`

GetTotalWinners returns the TotalWinners field if non-nil, zero value otherwise.

### GetTotalWinnersOk

`func (o *HandlerOGScoreboardData) GetTotalWinnersOk() (*int32, bool)`

GetTotalWinnersOk returns a tuple with the TotalWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWinners

`func (o *HandlerOGScoreboardData) SetTotalWinners(v int32)`

SetTotalWinners sets TotalWinners field to given value.

### HasTotalWinners

`func (o *HandlerOGScoreboardData) HasTotalWinners() bool`

HasTotalWinners returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *HandlerOGScoreboardData) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *HandlerOGScoreboardData) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *HandlerOGScoreboardData) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *HandlerOGScoreboardData) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


