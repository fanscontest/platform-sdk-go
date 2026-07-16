# EntryBasedContestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the creative category (e.g. photography, comedy) — required for entry-based contests. See Domain.ValidCategories. | [optional] 
**ExtensionCount** | Pointer to **int32** | default 0; incremented when judging is extended for insufficient jurors | [optional] 
**JudgingDeadline** | Pointer to **string** | set when the contest enters judging (EndAt + 7d) | [optional] 
**JudgingMethod** | Pointer to **string** | Judging method + community-method settlement config (uman#221). | [optional] 
**JurorSelectionMode** | Pointer to **string** | Domain.JurorSelectionMode{Random,Invited,Hybrid} | [optional] 
**JurySize** | Pointer to **int32** | Juror-method sizing. | [optional] 
**MinJurorScores** | Pointer to **int32** | default 3 | [optional] 
**Quorum** | Pointer to **int32** | community: min judgments per entry to settle | [optional] 
**RatingModel** | Pointer to **string** | community: Domain.RatingModelWinFraction | [optional] 
**TieBreak** | Pointer to **string** | community: Domain.TieBreakEarliestSubmission | [optional] 

## Methods

### NewEntryBasedContestConfig

`func NewEntryBasedContestConfig() *EntryBasedContestConfig`

NewEntryBasedContestConfig instantiates a new EntryBasedContestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryBasedContestConfigWithDefaults

`func NewEntryBasedContestConfigWithDefaults() *EntryBasedContestConfig`

NewEntryBasedContestConfigWithDefaults instantiates a new EntryBasedContestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *EntryBasedContestConfig) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EntryBasedContestConfig) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EntryBasedContestConfig) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EntryBasedContestConfig) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExtensionCount

`func (o *EntryBasedContestConfig) GetExtensionCount() int32`

GetExtensionCount returns the ExtensionCount field if non-nil, zero value otherwise.

### GetExtensionCountOk

`func (o *EntryBasedContestConfig) GetExtensionCountOk() (*int32, bool)`

GetExtensionCountOk returns a tuple with the ExtensionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionCount

`func (o *EntryBasedContestConfig) SetExtensionCount(v int32)`

SetExtensionCount sets ExtensionCount field to given value.

### HasExtensionCount

`func (o *EntryBasedContestConfig) HasExtensionCount() bool`

HasExtensionCount returns a boolean if a field has been set.

### GetJudgingDeadline

`func (o *EntryBasedContestConfig) GetJudgingDeadline() string`

GetJudgingDeadline returns the JudgingDeadline field if non-nil, zero value otherwise.

### GetJudgingDeadlineOk

`func (o *EntryBasedContestConfig) GetJudgingDeadlineOk() (*string, bool)`

GetJudgingDeadlineOk returns a tuple with the JudgingDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgingDeadline

`func (o *EntryBasedContestConfig) SetJudgingDeadline(v string)`

SetJudgingDeadline sets JudgingDeadline field to given value.

### HasJudgingDeadline

`func (o *EntryBasedContestConfig) HasJudgingDeadline() bool`

HasJudgingDeadline returns a boolean if a field has been set.

### GetJudgingMethod

`func (o *EntryBasedContestConfig) GetJudgingMethod() string`

GetJudgingMethod returns the JudgingMethod field if non-nil, zero value otherwise.

### GetJudgingMethodOk

`func (o *EntryBasedContestConfig) GetJudgingMethodOk() (*string, bool)`

GetJudgingMethodOk returns a tuple with the JudgingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgingMethod

`func (o *EntryBasedContestConfig) SetJudgingMethod(v string)`

SetJudgingMethod sets JudgingMethod field to given value.

### HasJudgingMethod

`func (o *EntryBasedContestConfig) HasJudgingMethod() bool`

HasJudgingMethod returns a boolean if a field has been set.

### GetJurorSelectionMode

`func (o *EntryBasedContestConfig) GetJurorSelectionMode() string`

GetJurorSelectionMode returns the JurorSelectionMode field if non-nil, zero value otherwise.

### GetJurorSelectionModeOk

`func (o *EntryBasedContestConfig) GetJurorSelectionModeOk() (*string, bool)`

GetJurorSelectionModeOk returns a tuple with the JurorSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurorSelectionMode

`func (o *EntryBasedContestConfig) SetJurorSelectionMode(v string)`

SetJurorSelectionMode sets JurorSelectionMode field to given value.

### HasJurorSelectionMode

`func (o *EntryBasedContestConfig) HasJurorSelectionMode() bool`

HasJurorSelectionMode returns a boolean if a field has been set.

### GetJurySize

`func (o *EntryBasedContestConfig) GetJurySize() int32`

GetJurySize returns the JurySize field if non-nil, zero value otherwise.

### GetJurySizeOk

`func (o *EntryBasedContestConfig) GetJurySizeOk() (*int32, bool)`

GetJurySizeOk returns a tuple with the JurySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurySize

`func (o *EntryBasedContestConfig) SetJurySize(v int32)`

SetJurySize sets JurySize field to given value.

### HasJurySize

`func (o *EntryBasedContestConfig) HasJurySize() bool`

HasJurySize returns a boolean if a field has been set.

### GetMinJurorScores

`func (o *EntryBasedContestConfig) GetMinJurorScores() int32`

GetMinJurorScores returns the MinJurorScores field if non-nil, zero value otherwise.

### GetMinJurorScoresOk

`func (o *EntryBasedContestConfig) GetMinJurorScoresOk() (*int32, bool)`

GetMinJurorScoresOk returns a tuple with the MinJurorScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinJurorScores

`func (o *EntryBasedContestConfig) SetMinJurorScores(v int32)`

SetMinJurorScores sets MinJurorScores field to given value.

### HasMinJurorScores

`func (o *EntryBasedContestConfig) HasMinJurorScores() bool`

HasMinJurorScores returns a boolean if a field has been set.

### GetQuorum

`func (o *EntryBasedContestConfig) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *EntryBasedContestConfig) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *EntryBasedContestConfig) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *EntryBasedContestConfig) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetRatingModel

`func (o *EntryBasedContestConfig) GetRatingModel() string`

GetRatingModel returns the RatingModel field if non-nil, zero value otherwise.

### GetRatingModelOk

`func (o *EntryBasedContestConfig) GetRatingModelOk() (*string, bool)`

GetRatingModelOk returns a tuple with the RatingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingModel

`func (o *EntryBasedContestConfig) SetRatingModel(v string)`

SetRatingModel sets RatingModel field to given value.

### HasRatingModel

`func (o *EntryBasedContestConfig) HasRatingModel() bool`

HasRatingModel returns a boolean if a field has been set.

### GetTieBreak

`func (o *EntryBasedContestConfig) GetTieBreak() string`

GetTieBreak returns the TieBreak field if non-nil, zero value otherwise.

### GetTieBreakOk

`func (o *EntryBasedContestConfig) GetTieBreakOk() (*string, bool)`

GetTieBreakOk returns a tuple with the TieBreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieBreak

`func (o *EntryBasedContestConfig) SetTieBreak(v string)`

SetTieBreak sets TieBreak field to given value.

### HasTieBreak

`func (o *EntryBasedContestConfig) HasTieBreak() bool`

HasTieBreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


