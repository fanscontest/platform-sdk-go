# DomainBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Blocker** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Channel** | Pointer to [**DomainChannel**](DomainChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainBlock

`func NewDomainBlock() *DomainBlock`

NewDomainBlock instantiates a new DomainBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBlockWithDefaults

`func NewDomainBlockWithDefaults() *DomainBlock`

NewDomainBlockWithDefaults instantiates a new DomainBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *DomainBlock) GetBlocked() DomainPlatformIdentity`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *DomainBlock) GetBlockedOk() (*DomainPlatformIdentity, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *DomainBlock) SetBlocked(v DomainPlatformIdentity)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *DomainBlock) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlocker

`func (o *DomainBlock) GetBlocker() DomainPlatformIdentity`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *DomainBlock) GetBlockerOk() (*DomainPlatformIdentity, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *DomainBlock) SetBlocker(v DomainPlatformIdentity)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *DomainBlock) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetChannel

`func (o *DomainBlock) GetChannel() DomainChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DomainBlock) GetChannelOk() (*DomainChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DomainBlock) SetChannel(v DomainChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DomainBlock) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainBlock) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainBlock) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainBlock) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainBlock) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


