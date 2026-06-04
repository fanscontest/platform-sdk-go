# RequestUpdateContestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**GroupingType** | Pointer to **string** |  | [optional] 
**NoOfGroups** | **int32** |  | 
**Reward** | [**DomainReward**](DomainReward.md) |  | 
**Timing** | [**DomainTiming**](DomainTiming.md) |  | 
**Title** | **string** |  | 
**WinnersPerGroup** | **int32** |  | 

## Methods

### NewRequestUpdateContestRequest

`func NewRequestUpdateContestRequest(description string, noOfGroups int32, reward DomainReward, timing DomainTiming, title string, winnersPerGroup int32, ) *RequestUpdateContestRequest`

NewRequestUpdateContestRequest instantiates a new RequestUpdateContestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdateContestRequestWithDefaults

`func NewRequestUpdateContestRequestWithDefaults() *RequestUpdateContestRequest`

NewRequestUpdateContestRequestWithDefaults instantiates a new RequestUpdateContestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RequestUpdateContestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestUpdateContestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestUpdateContestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupingType

`func (o *RequestUpdateContestRequest) GetGroupingType() string`

GetGroupingType returns the GroupingType field if non-nil, zero value otherwise.

### GetGroupingTypeOk

`func (o *RequestUpdateContestRequest) GetGroupingTypeOk() (*string, bool)`

GetGroupingTypeOk returns a tuple with the GroupingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingType

`func (o *RequestUpdateContestRequest) SetGroupingType(v string)`

SetGroupingType sets GroupingType field to given value.

### HasGroupingType

`func (o *RequestUpdateContestRequest) HasGroupingType() bool`

HasGroupingType returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *RequestUpdateContestRequest) GetNoOfGroups() int32`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *RequestUpdateContestRequest) GetNoOfGroupsOk() (*int32, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *RequestUpdateContestRequest) SetNoOfGroups(v int32)`

SetNoOfGroups sets NoOfGroups field to given value.


### GetReward

`func (o *RequestUpdateContestRequest) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *RequestUpdateContestRequest) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *RequestUpdateContestRequest) SetReward(v DomainReward)`

SetReward sets Reward field to given value.


### GetTiming

`func (o *RequestUpdateContestRequest) GetTiming() DomainTiming`

GetTiming returns the Timing field if non-nil, zero value otherwise.

### GetTimingOk

`func (o *RequestUpdateContestRequest) GetTimingOk() (*DomainTiming, bool)`

GetTimingOk returns a tuple with the Timing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiming

`func (o *RequestUpdateContestRequest) SetTiming(v DomainTiming)`

SetTiming sets Timing field to given value.


### GetTitle

`func (o *RequestUpdateContestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestUpdateContestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestUpdateContestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWinnersPerGroup

`func (o *RequestUpdateContestRequest) GetWinnersPerGroup() int32`

GetWinnersPerGroup returns the WinnersPerGroup field if non-nil, zero value otherwise.

### GetWinnersPerGroupOk

`func (o *RequestUpdateContestRequest) GetWinnersPerGroupOk() (*int32, bool)`

GetWinnersPerGroupOk returns a tuple with the WinnersPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnersPerGroup

`func (o *RequestUpdateContestRequest) SetWinnersPerGroup(v int32)`

SetWinnersPerGroup sets WinnersPerGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


