# HandlerOGParticipationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestId** | Pointer to **string** |  | [optional] 
**ContestTitle** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **[]string** |  | [optional] 
**Participant** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**ParticipantName** | Pointer to **string** |  | [optional] 
**Predictions** | Pointer to [**[]HandlerEventPredictions**](HandlerEventPredictions.md) | Grouped by event | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlerOGParticipationData

`func NewHandlerOGParticipationData() *HandlerOGParticipationData`

NewHandlerOGParticipationData instantiates a new HandlerOGParticipationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerOGParticipationDataWithDefaults

`func NewHandlerOGParticipationDataWithDefaults() *HandlerOGParticipationData`

NewHandlerOGParticipationDataWithDefaults instantiates a new HandlerOGParticipationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestId

`func (o *HandlerOGParticipationData) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *HandlerOGParticipationData) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *HandlerOGParticipationData) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *HandlerOGParticipationData) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetContestTitle

`func (o *HandlerOGParticipationData) GetContestTitle() string`

GetContestTitle returns the ContestTitle field if non-nil, zero value otherwise.

### GetContestTitleOk

`func (o *HandlerOGParticipationData) GetContestTitleOk() (*string, bool)`

GetContestTitleOk returns a tuple with the ContestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestTitle

`func (o *HandlerOGParticipationData) SetContestTitle(v string)`

SetContestTitle sets ContestTitle field to given value.

### HasContestTitle

`func (o *HandlerOGParticipationData) HasContestTitle() bool`

HasContestTitle returns a boolean if a field has been set.

### GetCreator

`func (o *HandlerOGParticipationData) GetCreator() DomainPlatformIdentity`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *HandlerOGParticipationData) GetCreatorOk() (*DomainPlatformIdentity, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *HandlerOGParticipationData) SetCreator(v DomainPlatformIdentity)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *HandlerOGParticipationData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetGroupName

`func (o *HandlerOGParticipationData) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *HandlerOGParticipationData) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *HandlerOGParticipationData) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *HandlerOGParticipationData) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetMetadata

`func (o *HandlerOGParticipationData) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HandlerOGParticipationData) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HandlerOGParticipationData) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HandlerOGParticipationData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParticipant

`func (o *HandlerOGParticipationData) GetParticipant() DomainPlatformIdentity`

GetParticipant returns the Participant field if non-nil, zero value otherwise.

### GetParticipantOk

`func (o *HandlerOGParticipationData) GetParticipantOk() (*DomainPlatformIdentity, bool)`

GetParticipantOk returns a tuple with the Participant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipant

`func (o *HandlerOGParticipationData) SetParticipant(v DomainPlatformIdentity)`

SetParticipant sets Participant field to given value.

### HasParticipant

`func (o *HandlerOGParticipationData) HasParticipant() bool`

HasParticipant returns a boolean if a field has been set.

### GetParticipantName

`func (o *HandlerOGParticipationData) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *HandlerOGParticipationData) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *HandlerOGParticipationData) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *HandlerOGParticipationData) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### GetPredictions

`func (o *HandlerOGParticipationData) GetPredictions() []HandlerEventPredictions`

GetPredictions returns the Predictions field if non-nil, zero value otherwise.

### GetPredictionsOk

`func (o *HandlerOGParticipationData) GetPredictionsOk() (*[]HandlerEventPredictions, bool)`

GetPredictionsOk returns a tuple with the Predictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictions

`func (o *HandlerOGParticipationData) SetPredictions(v []HandlerEventPredictions)`

SetPredictions sets Predictions field to given value.

### HasPredictions

`func (o *HandlerOGParticipationData) HasPredictions() bool`

HasPredictions returns a boolean if a field has been set.

### GetRank

`func (o *HandlerOGParticipationData) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *HandlerOGParticipationData) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *HandlerOGParticipationData) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *HandlerOGParticipationData) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetScore

`func (o *HandlerOGParticipationData) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *HandlerOGParticipationData) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *HandlerOGParticipationData) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *HandlerOGParticipationData) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStatus

`func (o *HandlerOGParticipationData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerOGParticipationData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerOGParticipationData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlerOGParticipationData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


