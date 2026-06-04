# DomainParticipationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**[]DomainParticipationResultBreakdownInner**](DomainParticipationResultBreakdownInner.md) | Legacy field used to migrate pre-v2 records that lack graded_sheet. Stripped before v2 responses. | [optional] 
**GlobalRank** | Pointer to **int32** |  | [optional] 
**GradedAt** | Pointer to **string** | Grading timestamp | [optional] 
**GradedSheet** | Pointer to **[]int32** | Complete graded sheet (v2 preferred) | [optional] 
**GroupRank** | Pointer to **int32** |  | [optional] 
**ParticipationPoint** | Pointer to **float32** | Participation points with 3 decimal precision | [optional] 
**ResultsIncomplete** | Pointer to **bool** | True if &lt; MinJurorScores for creative contests | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**ScorePct** | Pointer to **string** |  | [optional] 
**SubmissionZone** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**TimePenalty** | Pointer to **string** | Legacy fields used to backfill ZoneDeduction/SubmissionZone for pre-v2 records. | [optional] 
**TimePenaltyZone** | Pointer to **string** |  | [optional] 
**Won** | Pointer to **bool** |  | [optional] 
**ZoneDeduction** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainParticipationResult

`func NewDomainParticipationResult() *DomainParticipationResult`

NewDomainParticipationResult instantiates a new DomainParticipationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainParticipationResultWithDefaults

`func NewDomainParticipationResultWithDefaults() *DomainParticipationResult`

NewDomainParticipationResultWithDefaults instantiates a new DomainParticipationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *DomainParticipationResult) GetBreakdown() []DomainParticipationResultBreakdownInner`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *DomainParticipationResult) GetBreakdownOk() (*[]DomainParticipationResultBreakdownInner, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *DomainParticipationResult) SetBreakdown(v []DomainParticipationResultBreakdownInner)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *DomainParticipationResult) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetGlobalRank

`func (o *DomainParticipationResult) GetGlobalRank() int32`

GetGlobalRank returns the GlobalRank field if non-nil, zero value otherwise.

### GetGlobalRankOk

`func (o *DomainParticipationResult) GetGlobalRankOk() (*int32, bool)`

GetGlobalRankOk returns a tuple with the GlobalRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRank

`func (o *DomainParticipationResult) SetGlobalRank(v int32)`

SetGlobalRank sets GlobalRank field to given value.

### HasGlobalRank

`func (o *DomainParticipationResult) HasGlobalRank() bool`

HasGlobalRank returns a boolean if a field has been set.

### GetGradedAt

`func (o *DomainParticipationResult) GetGradedAt() string`

GetGradedAt returns the GradedAt field if non-nil, zero value otherwise.

### GetGradedAtOk

`func (o *DomainParticipationResult) GetGradedAtOk() (*string, bool)`

GetGradedAtOk returns a tuple with the GradedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradedAt

`func (o *DomainParticipationResult) SetGradedAt(v string)`

SetGradedAt sets GradedAt field to given value.

### HasGradedAt

`func (o *DomainParticipationResult) HasGradedAt() bool`

HasGradedAt returns a boolean if a field has been set.

### GetGradedSheet

`func (o *DomainParticipationResult) GetGradedSheet() []int32`

GetGradedSheet returns the GradedSheet field if non-nil, zero value otherwise.

### GetGradedSheetOk

`func (o *DomainParticipationResult) GetGradedSheetOk() (*[]int32, bool)`

GetGradedSheetOk returns a tuple with the GradedSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradedSheet

`func (o *DomainParticipationResult) SetGradedSheet(v []int32)`

SetGradedSheet sets GradedSheet field to given value.

### HasGradedSheet

`func (o *DomainParticipationResult) HasGradedSheet() bool`

HasGradedSheet returns a boolean if a field has been set.

### GetGroupRank

`func (o *DomainParticipationResult) GetGroupRank() int32`

GetGroupRank returns the GroupRank field if non-nil, zero value otherwise.

### GetGroupRankOk

`func (o *DomainParticipationResult) GetGroupRankOk() (*int32, bool)`

GetGroupRankOk returns a tuple with the GroupRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRank

`func (o *DomainParticipationResult) SetGroupRank(v int32)`

SetGroupRank sets GroupRank field to given value.

### HasGroupRank

`func (o *DomainParticipationResult) HasGroupRank() bool`

HasGroupRank returns a boolean if a field has been set.

### GetParticipationPoint

`func (o *DomainParticipationResult) GetParticipationPoint() float32`

GetParticipationPoint returns the ParticipationPoint field if non-nil, zero value otherwise.

### GetParticipationPointOk

`func (o *DomainParticipationResult) GetParticipationPointOk() (*float32, bool)`

GetParticipationPointOk returns a tuple with the ParticipationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationPoint

`func (o *DomainParticipationResult) SetParticipationPoint(v float32)`

SetParticipationPoint sets ParticipationPoint field to given value.

### HasParticipationPoint

`func (o *DomainParticipationResult) HasParticipationPoint() bool`

HasParticipationPoint returns a boolean if a field has been set.

### GetResultsIncomplete

`func (o *DomainParticipationResult) GetResultsIncomplete() bool`

GetResultsIncomplete returns the ResultsIncomplete field if non-nil, zero value otherwise.

### GetResultsIncompleteOk

`func (o *DomainParticipationResult) GetResultsIncompleteOk() (*bool, bool)`

GetResultsIncompleteOk returns a tuple with the ResultsIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsIncomplete

`func (o *DomainParticipationResult) SetResultsIncomplete(v bool)`

SetResultsIncomplete sets ResultsIncomplete field to given value.

### HasResultsIncomplete

`func (o *DomainParticipationResult) HasResultsIncomplete() bool`

HasResultsIncomplete returns a boolean if a field has been set.

### GetScore

`func (o *DomainParticipationResult) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DomainParticipationResult) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DomainParticipationResult) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DomainParticipationResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScorePct

`func (o *DomainParticipationResult) GetScorePct() string`

GetScorePct returns the ScorePct field if non-nil, zero value otherwise.

### GetScorePctOk

`func (o *DomainParticipationResult) GetScorePctOk() (*string, bool)`

GetScorePctOk returns a tuple with the ScorePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorePct

`func (o *DomainParticipationResult) SetScorePct(v string)`

SetScorePct sets ScorePct field to given value.

### HasScorePct

`func (o *DomainParticipationResult) HasScorePct() bool`

HasScorePct returns a boolean if a field has been set.

### GetSubmissionZone

`func (o *DomainParticipationResult) GetSubmissionZone() string`

GetSubmissionZone returns the SubmissionZone field if non-nil, zero value otherwise.

### GetSubmissionZoneOk

`func (o *DomainParticipationResult) GetSubmissionZoneOk() (*string, bool)`

GetSubmissionZoneOk returns a tuple with the SubmissionZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionZone

`func (o *DomainParticipationResult) SetSubmissionZone(v string)`

SetSubmissionZone sets SubmissionZone field to given value.

### HasSubmissionZone

`func (o *DomainParticipationResult) HasSubmissionZone() bool`

HasSubmissionZone returns a boolean if a field has been set.

### GetTime

`func (o *DomainParticipationResult) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DomainParticipationResult) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DomainParticipationResult) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *DomainParticipationResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimePenalty

`func (o *DomainParticipationResult) GetTimePenalty() string`

GetTimePenalty returns the TimePenalty field if non-nil, zero value otherwise.

### GetTimePenaltyOk

`func (o *DomainParticipationResult) GetTimePenaltyOk() (*string, bool)`

GetTimePenaltyOk returns a tuple with the TimePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePenalty

`func (o *DomainParticipationResult) SetTimePenalty(v string)`

SetTimePenalty sets TimePenalty field to given value.

### HasTimePenalty

`func (o *DomainParticipationResult) HasTimePenalty() bool`

HasTimePenalty returns a boolean if a field has been set.

### GetTimePenaltyZone

`func (o *DomainParticipationResult) GetTimePenaltyZone() string`

GetTimePenaltyZone returns the TimePenaltyZone field if non-nil, zero value otherwise.

### GetTimePenaltyZoneOk

`func (o *DomainParticipationResult) GetTimePenaltyZoneOk() (*string, bool)`

GetTimePenaltyZoneOk returns a tuple with the TimePenaltyZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePenaltyZone

`func (o *DomainParticipationResult) SetTimePenaltyZone(v string)`

SetTimePenaltyZone sets TimePenaltyZone field to given value.

### HasTimePenaltyZone

`func (o *DomainParticipationResult) HasTimePenaltyZone() bool`

HasTimePenaltyZone returns a boolean if a field has been set.

### GetWon

`func (o *DomainParticipationResult) GetWon() bool`

GetWon returns the Won field if non-nil, zero value otherwise.

### GetWonOk

`func (o *DomainParticipationResult) GetWonOk() (*bool, bool)`

GetWonOk returns a tuple with the Won field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWon

`func (o *DomainParticipationResult) SetWon(v bool)`

SetWon sets Won field to given value.

### HasWon

`func (o *DomainParticipationResult) HasWon() bool`

HasWon returns a boolean if a field has been set.

### GetZoneDeduction

`func (o *DomainParticipationResult) GetZoneDeduction() string`

GetZoneDeduction returns the ZoneDeduction field if non-nil, zero value otherwise.

### GetZoneDeductionOk

`func (o *DomainParticipationResult) GetZoneDeductionOk() (*string, bool)`

GetZoneDeductionOk returns a tuple with the ZoneDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDeduction

`func (o *DomainParticipationResult) SetZoneDeduction(v string)`

SetZoneDeduction sets ZoneDeduction field to given value.

### HasZoneDeduction

`func (o *DomainParticipationResult) HasZoneDeduction() bool`

HasZoneDeduction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


