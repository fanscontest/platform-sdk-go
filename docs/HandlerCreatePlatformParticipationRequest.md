# HandlerCreatePlatformParticipationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**DisplayName** | **string** |  | 
**PlatformIdentityId** | **string** |  | 
**Solution** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewHandlerCreatePlatformParticipationRequest

`func NewHandlerCreatePlatformParticipationRequest(channelId string, displayName string, platformIdentityId string, ) *HandlerCreatePlatformParticipationRequest`

NewHandlerCreatePlatformParticipationRequest instantiates a new HandlerCreatePlatformParticipationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreatePlatformParticipationRequestWithDefaults

`func NewHandlerCreatePlatformParticipationRequestWithDefaults() *HandlerCreatePlatformParticipationRequest`

NewHandlerCreatePlatformParticipationRequestWithDefaults instantiates a new HandlerCreatePlatformParticipationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *HandlerCreatePlatformParticipationRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *HandlerCreatePlatformParticipationRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *HandlerCreatePlatformParticipationRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetDisplayName

`func (o *HandlerCreatePlatformParticipationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HandlerCreatePlatformParticipationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HandlerCreatePlatformParticipationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPlatformIdentityId

`func (o *HandlerCreatePlatformParticipationRequest) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *HandlerCreatePlatformParticipationRequest) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *HandlerCreatePlatformParticipationRequest) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.


### GetSolution

`func (o *HandlerCreatePlatformParticipationRequest) GetSolution() []int32`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *HandlerCreatePlatformParticipationRequest) GetSolutionOk() (*[]int32, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *HandlerCreatePlatformParticipationRequest) SetSolution(v []int32)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *HandlerCreatePlatformParticipationRequest) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


