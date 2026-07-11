# DomainLeaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Context** | Pointer to [**DomainBoardContext**](DomainBoardContext.md) |  | [optional] 
**MemberRank** | Pointer to [**DomainLeaderboardRank**](DomainLeaderboardRank.md) |  | [optional] 
**Period** | Pointer to [**DomainPeriod**](DomainPeriod.md) |  | [optional] 
**Rankings** | Pointer to [**[]DomainLeaderboardRank**](DomainLeaderboardRank.md) |  | [optional] 

## Methods

### NewDomainLeaderboard

`func NewDomainLeaderboard() *DomainLeaderboard`

NewDomainLeaderboard instantiates a new DomainLeaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainLeaderboardWithDefaults

`func NewDomainLeaderboardWithDefaults() *DomainLeaderboard`

NewDomainLeaderboardWithDefaults instantiates a new DomainLeaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *DomainLeaderboard) GetBrand() DomainPlatformIdentity`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *DomainLeaderboard) GetBrandOk() (*DomainPlatformIdentity, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *DomainLeaderboard) SetBrand(v DomainPlatformIdentity)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *DomainLeaderboard) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetContext

`func (o *DomainLeaderboard) GetContext() DomainBoardContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DomainLeaderboard) GetContextOk() (*DomainBoardContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DomainLeaderboard) SetContext(v DomainBoardContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *DomainLeaderboard) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMemberRank

`func (o *DomainLeaderboard) GetMemberRank() DomainLeaderboardRank`

GetMemberRank returns the MemberRank field if non-nil, zero value otherwise.

### GetMemberRankOk

`func (o *DomainLeaderboard) GetMemberRankOk() (*DomainLeaderboardRank, bool)`

GetMemberRankOk returns a tuple with the MemberRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRank

`func (o *DomainLeaderboard) SetMemberRank(v DomainLeaderboardRank)`

SetMemberRank sets MemberRank field to given value.

### HasMemberRank

`func (o *DomainLeaderboard) HasMemberRank() bool`

HasMemberRank returns a boolean if a field has been set.

### GetPeriod

`func (o *DomainLeaderboard) GetPeriod() DomainPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainLeaderboard) GetPeriodOk() (*DomainPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainLeaderboard) SetPeriod(v DomainPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainLeaderboard) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRankings

`func (o *DomainLeaderboard) GetRankings() []DomainLeaderboardRank`

GetRankings returns the Rankings field if non-nil, zero value otherwise.

### GetRankingsOk

`func (o *DomainLeaderboard) GetRankingsOk() (*[]DomainLeaderboardRank, bool)`

GetRankingsOk returns a tuple with the Rankings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankings

`func (o *DomainLeaderboard) SetRankings(v []DomainLeaderboardRank)`

SetRankings sets Rankings field to given value.

### HasRankings

`func (o *DomainLeaderboard) HasRankings() bool`

HasRankings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


