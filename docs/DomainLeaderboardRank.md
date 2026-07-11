# DomainLeaderboardRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgSolutionTime** | Pointer to **float32** | Average solution time in seconds (nullable, with decimal precision) | [optional] 
**ContestsPlayed** | Pointer to **int32** | Count of contests | [optional] 
**Current** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Period** | Pointer to [**DomainPeriod**](DomainPeriod.md) | Period type and date range | [optional] 
**Points** | Pointer to **float32** | Participation points with 3 decimal precision (MCQ/zone metric) | [optional] 
**Previous** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **float32** | Cumulative raw score across the period (prediction/creative metric) | [optional] 

## Methods

### NewDomainLeaderboardRank

`func NewDomainLeaderboardRank() *DomainLeaderboardRank`

NewDomainLeaderboardRank instantiates a new DomainLeaderboardRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainLeaderboardRankWithDefaults

`func NewDomainLeaderboardRankWithDefaults() *DomainLeaderboardRank`

NewDomainLeaderboardRankWithDefaults instantiates a new DomainLeaderboardRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgSolutionTime

`func (o *DomainLeaderboardRank) GetAvgSolutionTime() float32`

GetAvgSolutionTime returns the AvgSolutionTime field if non-nil, zero value otherwise.

### GetAvgSolutionTimeOk

`func (o *DomainLeaderboardRank) GetAvgSolutionTimeOk() (*float32, bool)`

GetAvgSolutionTimeOk returns a tuple with the AvgSolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSolutionTime

`func (o *DomainLeaderboardRank) SetAvgSolutionTime(v float32)`

SetAvgSolutionTime sets AvgSolutionTime field to given value.

### HasAvgSolutionTime

`func (o *DomainLeaderboardRank) HasAvgSolutionTime() bool`

HasAvgSolutionTime returns a boolean if a field has been set.

### GetContestsPlayed

`func (o *DomainLeaderboardRank) GetContestsPlayed() int32`

GetContestsPlayed returns the ContestsPlayed field if non-nil, zero value otherwise.

### GetContestsPlayedOk

`func (o *DomainLeaderboardRank) GetContestsPlayedOk() (*int32, bool)`

GetContestsPlayedOk returns a tuple with the ContestsPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestsPlayed

`func (o *DomainLeaderboardRank) SetContestsPlayed(v int32)`

SetContestsPlayed sets ContestsPlayed field to given value.

### HasContestsPlayed

`func (o *DomainLeaderboardRank) HasContestsPlayed() bool`

HasContestsPlayed returns a boolean if a field has been set.

### GetCurrent

`func (o *DomainLeaderboardRank) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *DomainLeaderboardRank) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *DomainLeaderboardRank) SetCurrent(v int32)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *DomainLeaderboardRank) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetMember

`func (o *DomainLeaderboardRank) GetMember() DomainPlatformIdentity`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *DomainLeaderboardRank) GetMemberOk() (*DomainPlatformIdentity, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *DomainLeaderboardRank) SetMember(v DomainPlatformIdentity)`

SetMember sets Member field to given value.

### HasMember

`func (o *DomainLeaderboardRank) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPeriod

`func (o *DomainLeaderboardRank) GetPeriod() DomainPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainLeaderboardRank) GetPeriodOk() (*DomainPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainLeaderboardRank) SetPeriod(v DomainPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainLeaderboardRank) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPoints

`func (o *DomainLeaderboardRank) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DomainLeaderboardRank) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *DomainLeaderboardRank) SetPoints(v float32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *DomainLeaderboardRank) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPrevious

`func (o *DomainLeaderboardRank) GetPrevious() int32`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *DomainLeaderboardRank) GetPreviousOk() (*int32, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *DomainLeaderboardRank) SetPrevious(v int32)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *DomainLeaderboardRank) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetScore

`func (o *DomainLeaderboardRank) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DomainLeaderboardRank) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DomainLeaderboardRank) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *DomainLeaderboardRank) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


