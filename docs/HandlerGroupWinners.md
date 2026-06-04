# HandlerGroupWinners

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** |  | [optional] 
**Winners** | Pointer to [**[]HandlerWinner**](HandlerWinner.md) |  | [optional] 

## Methods

### NewHandlerGroupWinners

`func NewHandlerGroupWinners() *HandlerGroupWinners`

NewHandlerGroupWinners instantiates a new HandlerGroupWinners object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerGroupWinnersWithDefaults

`func NewHandlerGroupWinnersWithDefaults() *HandlerGroupWinners`

NewHandlerGroupWinnersWithDefaults instantiates a new HandlerGroupWinners object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *HandlerGroupWinners) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *HandlerGroupWinners) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *HandlerGroupWinners) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *HandlerGroupWinners) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetWinners

`func (o *HandlerGroupWinners) GetWinners() []HandlerWinner`

GetWinners returns the Winners field if non-nil, zero value otherwise.

### GetWinnersOk

`func (o *HandlerGroupWinners) GetWinnersOk() (*[]HandlerWinner, bool)`

GetWinnersOk returns a tuple with the Winners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinners

`func (o *HandlerGroupWinners) SetWinners(v []HandlerWinner)`

SetWinners sets Winners field to given value.

### HasWinners

`func (o *HandlerGroupWinners) HasWinners() bool`

HasWinners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


