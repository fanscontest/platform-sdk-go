# DomainContest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**DomainChannel**](DomainChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Difficulty** | Pointer to **string** |  | [optional] 
**EntryBasedConfig** | Pointer to [**EntryBasedContestConfig**](EntryBasedContestConfig.md) | Present only for entry-based (creative) contests — juror sizing + judging config (uman#221) | [optional] 
**FairPlayGrouping** | Pointer to **bool** | @deprecated | [optional] 
**GroupingType** | Pointer to **string** | &#39;system&#39; or &#39;open&#39; | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MemberParticipation** | Pointer to [**DomainMemberParticipationSummary**](DomainMemberParticipationSummary.md) |  | [optional] 
**NoOfGroups** | Pointer to **int32** |  | [optional] 
**NoOfPuzzles** | Pointer to **int32** |  | [optional] 
**NoOfWinners** | Pointer to **int32** | @deprecated | [optional] 
**ParticipationRules** | Pointer to [**DomainParticipationRules**](DomainParticipationRules.md) |  | [optional] 
**PuzzleSheetId** | Pointer to **string** | v1; prefer source_id | [optional] 
**QuestionType** | Pointer to [**DomainContestType**](DomainContestType.md) | v1; prefer source_type | [optional] 
**Reward** | Pointer to [**DomainReward**](DomainReward.md) |  | [optional] 
**Sender** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**SourceId** | Pointer to **string** | v2; same as puzzle_sheet_id | [optional] 
**SourceType** | Pointer to [**DomainContestType**](DomainContestType.md) | v2; same as question_type | [optional] 
**Sponsor** | Pointer to [**DomainSponsorInfo**](DomainSponsorInfo.md) | Only present if contest has accepted sponsorship | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubmissionZoneConfiguration** | Pointer to [**DomainSubmissionZoneConfiguration**](DomainSubmissionZoneConfiguration.md) |  | [optional] 
**TeamSet** | Pointer to [**DomainTeamSet**](DomainTeamSet.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TimePenaltyConfiguration** | Pointer to [**[]DomainTimePenaltyConfiguration**](DomainTimePenaltyConfiguration.md) |  | [optional] 
**Timing** | Pointer to [**DomainTiming**](DomainTiming.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**WinnersPerGroup** | Pointer to **int32** |  | [optional] 

## Methods

### NewDomainContest

`func NewDomainContest() *DomainContest`

NewDomainContest instantiates a new DomainContest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContestWithDefaults

`func NewDomainContestWithDefaults() *DomainContest`

NewDomainContestWithDefaults instantiates a new DomainContest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DomainContest) GetChannel() DomainChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DomainContest) GetChannelOk() (*DomainChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DomainContest) SetChannel(v DomainChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DomainContest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainContest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainContest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainContest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainContest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DomainContest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainContest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainContest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainContest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDifficulty

`func (o *DomainContest) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *DomainContest) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *DomainContest) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *DomainContest) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetEntryBasedConfig

`func (o *DomainContest) GetEntryBasedConfig() EntryBasedContestConfig`

GetEntryBasedConfig returns the EntryBasedConfig field if non-nil, zero value otherwise.

### GetEntryBasedConfigOk

`func (o *DomainContest) GetEntryBasedConfigOk() (*EntryBasedContestConfig, bool)`

GetEntryBasedConfigOk returns a tuple with the EntryBasedConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryBasedConfig

`func (o *DomainContest) SetEntryBasedConfig(v EntryBasedContestConfig)`

SetEntryBasedConfig sets EntryBasedConfig field to given value.

### HasEntryBasedConfig

`func (o *DomainContest) HasEntryBasedConfig() bool`

HasEntryBasedConfig returns a boolean if a field has been set.

### GetFairPlayGrouping

`func (o *DomainContest) GetFairPlayGrouping() bool`

GetFairPlayGrouping returns the FairPlayGrouping field if non-nil, zero value otherwise.

### GetFairPlayGroupingOk

`func (o *DomainContest) GetFairPlayGroupingOk() (*bool, bool)`

GetFairPlayGroupingOk returns a tuple with the FairPlayGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFairPlayGrouping

`func (o *DomainContest) SetFairPlayGrouping(v bool)`

SetFairPlayGrouping sets FairPlayGrouping field to given value.

### HasFairPlayGrouping

`func (o *DomainContest) HasFairPlayGrouping() bool`

HasFairPlayGrouping returns a boolean if a field has been set.

### GetGroupingType

`func (o *DomainContest) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *DomainContest) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *DomainContest) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *DomainContest) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *DomainContest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *DomainContest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *DomainContest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *DomainContest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetId

`func (o *DomainContest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainContest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainContest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainContest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberParticipation

`func (o *DomainContest) GetMemberParticipation() DomainMemberParticipationSummary`

GetMemberParticipation returns the MemberParticipation field if non-nil, zero value otherwise.

### GetMemberParticipationOk

`func (o *DomainContest) GetMemberParticipationOk() (*DomainMemberParticipationSummary, bool)`

GetMemberParticipationOk returns a tuple with the MemberParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberParticipation

`func (o *DomainContest) SetMemberParticipation(v DomainMemberParticipationSummary)`

SetMemberParticipation sets MemberParticipation field to given value.

### HasMemberParticipation

`func (o *DomainContest) HasMemberParticipation() bool`

HasMemberParticipation returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *DomainContest) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *DomainContest) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *DomainContest) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.

### HasNoOfGroups

`func (o *DomainContest) HasNoOfGroups() bool`

HasNoOfGroups returns a boolean if a field has been set.

### GetNoOfPuzzles

`func (o *DomainContest) GetNoOfPuzzles() int32`

GetNoOfPuzzles returns the NoOfPuzzles field if non-nil, zero value otherwise.

### GetNoOfPuzzlesOk

`func (o *DomainContest) GetNoOfPuzzlesOk() (*int32, bool)`

GetNoOfPuzzlesOk returns a tuple with the NoOfPuzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPuzzles

`func (o *DomainContest) SetNoOfPuzzles(v int32)`

SetNoOfPuzzles sets NoOfPuzzles field to given value.

### HasNoOfPuzzles

`func (o *DomainContest) HasNoOfPuzzles() bool`

HasNoOfPuzzles returns a boolean if a field has been set.

### GetNoOfWinners

`func (o *DomainContest) GetNoOfWinners() int32`

GetNoOfWinners returns the NoOfWinners field if non-nil, zero value otherwise.

### GetNoOfWinnersOk

`func (o *DomainContest) GetNoOfWinnersOk() (*int32, bool)`

GetNoOfWinnersOk returns a tuple with the NoOfWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfWinners

`func (o *DomainContest) SetNoOfWinners(v int32)`

SetNoOfWinners sets NoOfWinners field to given value.

### HasNoOfWinners

`func (o *DomainContest) HasNoOfWinners() bool`

HasNoOfWinners returns a boolean if a field has been set.

### GetParticipationRules

`func (o *DomainContest) GetParticipationRules() DomainParticipationRules`

GetParticipationRules returns the ParticipationRules field if non-nil, zero value otherwise.

### GetParticipationRulesOk

`func (o *DomainContest) GetParticipationRulesOk() (*DomainParticipationRules, bool)`

GetParticipationRulesOk returns a tuple with the ParticipationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationRules

`func (o *DomainContest) SetParticipationRules(v DomainParticipationRules)`

SetParticipationRules sets ParticipationRules field to given value.

### HasParticipationRules

`func (o *DomainContest) HasParticipationRules() bool`

HasParticipationRules returns a boolean if a field has been set.

### GetPuzzleSheetId

`func (o *DomainContest) GetPuzzleSheetId() string`

GetPuzzleSheetId returns the PuzzleSheetId field if non-nil, zero value otherwise.

### GetPuzzleSheetIdOk

`func (o *DomainContest) GetPuzzleSheetIdOk() (*string, bool)`

GetPuzzleSheetIdOk returns a tuple with the PuzzleSheetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzleSheetId

`func (o *DomainContest) SetPuzzleSheetId(v string)`

SetPuzzleSheetId sets PuzzleSheetId field to given value.

### HasPuzzleSheetId

`func (o *DomainContest) HasPuzzleSheetId() bool`

HasPuzzleSheetId returns a boolean if a field has been set.

### GetQuestionType

`func (o *DomainContest) GetQuestionType() DomainContestType`

GetQuestionType returns the QuestionType field if non-nil, zero value otherwise.

### GetQuestionTypeOk

`func (o *DomainContest) GetQuestionTypeOk() (*DomainContestType, bool)`

GetQuestionTypeOk returns a tuple with the QuestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionType

`func (o *DomainContest) SetQuestionType(v DomainContestType)`

SetQuestionType sets QuestionType field to given value.

### HasQuestionType

`func (o *DomainContest) HasQuestionType() bool`

HasQuestionType returns a boolean if a field has been set.

### GetReward

`func (o *DomainContest) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *DomainContest) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *DomainContest) SetReward(v DomainReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *DomainContest) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetSender

`func (o *DomainContest) GetSender() DomainPlatformIdentity`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DomainContest) GetSenderOk() (*DomainPlatformIdentity, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DomainContest) SetSender(v DomainPlatformIdentity)`

SetSender sets Sender field to given value.

### HasSender

`func (o *DomainContest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSourceId

`func (o *DomainContest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DomainContest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DomainContest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DomainContest) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceType

`func (o *DomainContest) GetSourceType() DomainContestType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DomainContest) GetSourceTypeOk() (*DomainContestType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DomainContest) SetSourceType(v DomainContestType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *DomainContest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSponsor

`func (o *DomainContest) GetSponsor() DomainSponsorInfo`

GetSponsor returns the Sponsor field if non-nil, zero value otherwise.

### GetSponsorOk

`func (o *DomainContest) GetSponsorOk() (*DomainSponsorInfo, bool)`

GetSponsorOk returns a tuple with the Sponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsor

`func (o *DomainContest) SetSponsor(v DomainSponsorInfo)`

SetSponsor sets Sponsor field to given value.

### HasSponsor

`func (o *DomainContest) HasSponsor() bool`

HasSponsor returns a boolean if a field has been set.

### GetState

`func (o *DomainContest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainContest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainContest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainContest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubmissionZoneConfiguration

`func (o *DomainContest) GetSubmissionZoneConfiguration() DomainSubmissionZoneConfiguration`

GetSubmissionZoneConfiguration returns the SubmissionZoneConfiguration field if non-nil, zero value otherwise.

### GetSubmissionZoneConfigurationOk

`func (o *DomainContest) GetSubmissionZoneConfigurationOk() (*DomainSubmissionZoneConfiguration, bool)`

GetSubmissionZoneConfigurationOk returns a tuple with the SubmissionZoneConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionZoneConfiguration

`func (o *DomainContest) SetSubmissionZoneConfiguration(v DomainSubmissionZoneConfiguration)`

SetSubmissionZoneConfiguration sets SubmissionZoneConfiguration field to given value.

### HasSubmissionZoneConfiguration

`func (o *DomainContest) HasSubmissionZoneConfiguration() bool`

HasSubmissionZoneConfiguration returns a boolean if a field has been set.

### GetTeamSet

`func (o *DomainContest) GetTeamSet() DomainTeamSet`

GetTeamSet returns the TeamSet field if non-nil, zero value otherwise.

### GetTeamSetOk

`func (o *DomainContest) GetTeamSetOk() (*DomainTeamSet, bool)`

GetTeamSetOk returns a tuple with the TeamSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSet

`func (o *DomainContest) SetTeamSet(v DomainTeamSet)`

SetTeamSet sets TeamSet field to given value.

### HasTeamSet

`func (o *DomainContest) HasTeamSet() bool`

HasTeamSet returns a boolean if a field has been set.

### GetTenantId

`func (o *DomainContest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DomainContest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DomainContest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DomainContest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTimePenaltyConfiguration

`func (o *DomainContest) GetTimePenaltyConfiguration() []DomainTimePenaltyConfiguration`

GetTimePenaltyConfiguration returns the TimePenaltyConfiguration field if non-nil, zero value otherwise.

### GetTimePenaltyConfigurationOk

`func (o *DomainContest) GetTimePenaltyConfigurationOk() (*[]DomainTimePenaltyConfiguration, bool)`

GetTimePenaltyConfigurationOk returns a tuple with the TimePenaltyConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePenaltyConfiguration

`func (o *DomainContest) SetTimePenaltyConfiguration(v []DomainTimePenaltyConfiguration)`

SetTimePenaltyConfiguration sets TimePenaltyConfiguration field to given value.

### HasTimePenaltyConfiguration

`func (o *DomainContest) HasTimePenaltyConfiguration() bool`

HasTimePenaltyConfiguration returns a boolean if a field has been set.

### GetTiming

`func (o *DomainContest) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *DomainContest) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *DomainContest) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.

### HasTiming

`func (o *DomainContest) HasTiming() bool`

HasTiming returns a boolean if a field has been set.

### GetTitle

`func (o *DomainContest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DomainContest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DomainContest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DomainContest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWinnersPerGroup

`func (o *DomainContest) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *DomainContest) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *DomainContest) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.

### HasWinnersPerGroup

`func (o *DomainContest) HasWinnersPerGroup() bool`

HasWinnersPerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


