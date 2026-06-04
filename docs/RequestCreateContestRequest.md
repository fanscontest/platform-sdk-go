# RequestCreateContestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**GroupingType** | Pointer to **string** |  | [optional] 
**JurorSelectionMode** | Pointer to **string** |  | [optional] 
**NoOfGroups** | **int32** |  | 
**NoOfPuzzles** | Pointer to **int32** |  | [optional] 
**Reward** | [**DomainReward**](DomainReward.md) |  | 
**RulePreset** | Pointer to **string** |  | [optional] 
**SourceId** | **string** |  | 
**SourceType** | **string** |  | 
**SponsorshipPackageId** | Pointer to **string** |  | [optional] 
**TeamSetId** | Pointer to **string** |  | [optional] 
**Timing** | [**DomainTiming**](DomainTiming.md) |  | 
**Title** | **string** |  | 
**VenueId** | **string** |  | 
**WinnersPerGroup** | **int32** |  | 
**ZoneAdvantage** | **int32** |  | 

## Methods

### NewRequestCreateContestRequest

`func NewRequestCreateContestRequest(noOfGroups int32, reward DomainReward, sourceId string, sourceType string, timing DomainTiming, title string, venueId string, winnersPerGroup int32, zoneAdvantage int32, ) *RequestCreateContestRequest`

NewRequestCreateContestRequest instantiates a new RequestCreateContestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateContestRequestWithDefaults

`func NewRequestCreateContestRequestWithDefaults() *RequestCreateContestRequest`

NewRequestCreateContestRequestWithDefaults instantiates a new RequestCreateContestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *RequestCreateContestRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RequestCreateContestRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RequestCreateContestRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RequestCreateContestRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChannelId

`func (o *RequestCreateContestRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *RequestCreateContestRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *RequestCreateContestRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *RequestCreateContestRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetDescription

`func (o *RequestCreateContestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestCreateContestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestCreateContestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestCreateContestRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDifficulty

`func (o *RequestCreateContestRequest) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *RequestCreateContestRequest) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *RequestCreateContestRequest) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *RequestCreateContestRequest) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetGroupingType

`func (o *RequestCreateContestRequest) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *RequestCreateContestRequest) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *RequestCreateContestRequest) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *RequestCreateContestRequest) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetJurorSelectionMode

`func (o *RequestCreateContestRequest) GetJurorSelectionMode() string`

GetJurorSelectionMode returns the JurorSelectionMode field if non-nil, zero value otherwise.

### GetJurorSelectionModeOk

`func (o *RequestCreateContestRequest) GetJurorSelectionModeOk() (*string, bool)`

GetJurorSelectionModeOk returns a tuple with the JurorSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurorSelectionMode

`func (o *RequestCreateContestRequest) SetJurorSelectionMode(v string)`

SetJurorSelectionMode sets JurorSelectionMode field to given value.

### HasJurorSelectionMode

`func (o *RequestCreateContestRequest) HasJurorSelectionMode() bool`

HasJurorSelectionMode returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *RequestCreateContestRequest) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *RequestCreateContestRequest) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *RequestCreateContestRequest) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.


### GetNoOfPuzzles

`func (o *RequestCreateContestRequest) GetNoOfPuzzles() int32`

GetNoOfPuzzles returns the NoOfPuzzles field if non-nil, zero value otherwise.

### GetNoOfPuzzlesOk

`func (o *RequestCreateContestRequest) GetNoOfPuzzlesOk() (*int32, bool)`

GetNoOfPuzzlesOk returns a tuple with the NoOfPuzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPuzzles

`func (o *RequestCreateContestRequest) SetNoOfPuzzles(v int32)`

SetNoOfPuzzles sets NoOfPuzzles field to given value.

### HasNoOfPuzzles

`func (o *RequestCreateContestRequest) HasNoOfPuzzles() bool`

HasNoOfPuzzles returns a boolean if a field has been set.

### GetReward

`func (o *RequestCreateContestRequest) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *RequestCreateContestRequest) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *RequestCreateContestRequest) SetReward(v DomainReward)`

SetReward sets Reward field to given value.


### GetRulePreset

`func (o *RequestCreateContestRequest) GetRulePreset() string`

GetRulePreset returns the RulePreset field if non-nil, zero value otherwise.

### GetRulePresetOk

`func (o *RequestCreateContestRequest) GetRulePresetOk() (*string, bool)`

GetRulePresetOk returns a tuple with the RulePreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulePreset

`func (o *RequestCreateContestRequest) SetRulePreset(v string)`

SetRulePreset sets RulePreset field to given value.

### HasRulePreset

`func (o *RequestCreateContestRequest) HasRulePreset() bool`

HasRulePreset returns a boolean if a field has been set.

### GetSourceId

`func (o *RequestCreateContestRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RequestCreateContestRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RequestCreateContestRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *RequestCreateContestRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RequestCreateContestRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RequestCreateContestRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSponsorshipPackageId

`func (o *RequestCreateContestRequest) GetSponsorshipPackageId() string`

GetSponsorshipPackageId returns the SponsorshipPackageId field if non-nil, zero value otherwise.

### GetSponsorshipPackageIdOk

`func (o *RequestCreateContestRequest) GetSponsorshipPackageIdOk() (*string, bool)`

GetSponsorshipPackageIdOk returns a tuple with the SponsorshipPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorshipPackageId

`func (o *RequestCreateContestRequest) SetSponsorshipPackageId(v string)`

SetSponsorshipPackageId sets SponsorshipPackageId field to given value.

### HasSponsorshipPackageId

`func (o *RequestCreateContestRequest) HasSponsorshipPackageId() bool`

HasSponsorshipPackageId returns a boolean if a field has been set.

### GetTeamSetId

`func (o *RequestCreateContestRequest) GetTeamSetId() string`

GetTeamSetId returns the TeamSetId field if non-nil, zero value otherwise.

### GetTeamSetIdOk

`func (o *RequestCreateContestRequest) GetTeamSetIdOk() (*string, bool)`

GetTeamSetIdOk returns a tuple with the TeamSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSetId

`func (o *RequestCreateContestRequest) SetTeamSetId(v string)`

SetTeamSetId sets TeamSetId field to given value.

### HasTeamSetId

`func (o *RequestCreateContestRequest) HasTeamSetId() bool`

HasTeamSetId returns a boolean if a field has been set.

### GetTiming

`func (o *RequestCreateContestRequest) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *RequestCreateContestRequest) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *RequestCreateContestRequest) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.


### GetTitle

`func (o *RequestCreateContestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestCreateContestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestCreateContestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVenueId

`func (o *RequestCreateContestRequest) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *RequestCreateContestRequest) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *RequestCreateContestRequest) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.


### GetWinnersPerGroup

`func (o *RequestCreateContestRequest) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *RequestCreateContestRequest) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *RequestCreateContestRequest) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.


### GetZoneAdvantage

`func (o *RequestCreateContestRequest) GetZoneAdvantage() int32`

GetZoneAdvantage returns the ZoneAdvantage field if non-nil, zero value otherwise.

### GetZoneAdvantageOk

`func (o *RequestCreateContestRequest) GetZoneAdvantageOk() (*int32, bool)`

GetZoneAdvantageOk returns a tuple with the ZoneAdvantage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdvantage

`func (o *RequestCreateContestRequest) SetZoneAdvantage(v int32)`

SetZoneAdvantage sets ZoneAdvantage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


