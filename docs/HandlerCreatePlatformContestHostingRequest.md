# HandlerCreatePlatformContestHostingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanParticipate** | Pointer to **bool** |  | [optional] 
**CanView** | Pointer to **bool** |  | [optional] 
**ChannelId** | **string** | venue_id is not a request field (uman#177): the channel determines the venue (channels.venue_id is canonical), so the caller supplies only the channel and the venue is derived from it. | 

## Methods

### NewHandlerCreatePlatformContestHostingRequest

`func NewHandlerCreatePlatformContestHostingRequest(channelId string, ) *HandlerCreatePlatformContestHostingRequest`

NewHandlerCreatePlatformContestHostingRequest instantiates a new HandlerCreatePlatformContestHostingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreatePlatformContestHostingRequestWithDefaults

`func NewHandlerCreatePlatformContestHostingRequestWithDefaults() *HandlerCreatePlatformContestHostingRequest`

NewHandlerCreatePlatformContestHostingRequestWithDefaults instantiates a new HandlerCreatePlatformContestHostingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanParticipate

`func (o *HandlerCreatePlatformContestHostingRequest) GetCanParticipate() bool`

GetCanParticipate returns the CanParticipate field if non-nil, zero value otherwise.

### GetCanParticipateOk

`func (o *HandlerCreatePlatformContestHostingRequest) GetCanParticipateOk() (*bool, bool)`

GetCanParticipateOk returns a tuple with the CanParticipate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanParticipate

`func (o *HandlerCreatePlatformContestHostingRequest) SetCanParticipate(v bool)`

SetCanParticipate sets CanParticipate field to given value.

### HasCanParticipate

`func (o *HandlerCreatePlatformContestHostingRequest) HasCanParticipate() bool`

HasCanParticipate returns a boolean if a field has been set.

### GetCanView

`func (o *HandlerCreatePlatformContestHostingRequest) GetCanView() bool`

GetCanView returns the CanView field if non-nil, zero value otherwise.

### GetCanViewOk

`func (o *HandlerCreatePlatformContestHostingRequest) GetCanViewOk() (*bool, bool)`

GetCanViewOk returns a tuple with the CanView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanView

`func (o *HandlerCreatePlatformContestHostingRequest) SetCanView(v bool)`

SetCanView sets CanView field to given value.

### HasCanView

`func (o *HandlerCreatePlatformContestHostingRequest) HasCanView() bool`

HasCanView returns a boolean if a field has been set.

### GetChannelId

`func (o *HandlerCreatePlatformContestHostingRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *HandlerCreatePlatformContestHostingRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *HandlerCreatePlatformContestHostingRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


