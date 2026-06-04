# RequestCreateBuddyBoardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoOfWinners** | **int32** |  | 
**Reward** | [**DomainReward**](DomainReward.md) |  | 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestCreateBuddyBoardRequest

`func NewRequestCreateBuddyBoardRequest(noOfWinners int32, reward DomainReward, ) *RequestCreateBuddyBoardRequest`

NewRequestCreateBuddyBoardRequest instantiates a new RequestCreateBuddyBoardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateBuddyBoardRequestWithDefaults

`func NewRequestCreateBuddyBoardRequestWithDefaults() *RequestCreateBuddyBoardRequest`

NewRequestCreateBuddyBoardRequestWithDefaults instantiates a new RequestCreateBuddyBoardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoOfWinners

`func (o *RequestCreateBuddyBoardRequest) GetNoOfWinners() int32`

GetNoOfWinners returns the NoOfWinners field if non-nil, zero value otherwise.

### GetNoOfWinnersOk

`func (o *RequestCreateBuddyBoardRequest) GetNoOfWinnersOk() (*int32, bool)`

GetNoOfWinnersOk returns a tuple with the NoOfWinners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfWinners

`func (o *RequestCreateBuddyBoardRequest) SetNoOfWinners(v int32)`

SetNoOfWinners sets NoOfWinners field to given value.


### GetReward

`func (o *RequestCreateBuddyBoardRequest) GetReward() DomainReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *RequestCreateBuddyBoardRequest) GetRewardOk() (*DomainReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *RequestCreateBuddyBoardRequest) SetReward(v DomainReward)`

SetReward sets Reward field to given value.


### GetTitle

`func (o *RequestCreateBuddyBoardRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestCreateBuddyBoardRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestCreateBuddyBoardRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequestCreateBuddyBoardRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


