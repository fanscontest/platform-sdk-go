# HandlerCreatePlatformParticipationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** |  | 
**DisplayName** | **string** |  | 
**Solution** | Pointer to **[]int32** |  | [optional] 
**TenantUserId** | **string** |  | 
**VenueId** | **string** |  | 

## Methods

### NewHandlerCreatePlatformParticipationRequest

`func NewHandlerCreatePlatformParticipationRequest(channelId string, displayName string, tenantUserId string, venueId string, ) *HandlerCreatePlatformParticipationRequest`

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

### GetTenantUserId

`func (o *HandlerCreatePlatformParticipationRequest) GetTenantUserId() string`

GetTenantUserId returns the TenantUserId field if non-nil, zero value otherwise.

### GetTenantUserIdOk

`func (o *HandlerCreatePlatformParticipationRequest) GetTenantUserIdOk() (*string, bool)`

GetTenantUserIdOk returns a tuple with the TenantUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUserId

`func (o *HandlerCreatePlatformParticipationRequest) SetTenantUserId(v string)`

SetTenantUserId sets TenantUserId field to given value.


### GetVenueId

`func (o *HandlerCreatePlatformParticipationRequest) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *HandlerCreatePlatformParticipationRequest) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *HandlerCreatePlatformParticipationRequest) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


