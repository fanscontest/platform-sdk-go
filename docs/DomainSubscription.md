# DomainSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**DomainChannel**](DomainChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subscriber** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 

## Methods

### NewDomainSubscription

`func NewDomainSubscription() *DomainSubscription`

NewDomainSubscription instantiates a new DomainSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSubscriptionWithDefaults

`func NewDomainSubscriptionWithDefaults() *DomainSubscription`

NewDomainSubscriptionWithDefaults instantiates a new DomainSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DomainSubscription) GetChannel() DomainChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DomainSubscription) GetChannelOk() (*DomainChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DomainSubscription) SetChannel(v DomainChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DomainSubscription) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSubscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSubscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSubscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DomainSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriber

`func (o *DomainSubscription) GetSubscriber() DomainPlatformIdentity`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *DomainSubscription) GetSubscriberOk() (*DomainPlatformIdentity, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *DomainSubscription) SetSubscriber(v DomainPlatformIdentity)`

SetSubscriber sets Subscriber field to given value.

### HasSubscriber

`func (o *DomainSubscription) HasSubscriber() bool`

HasSubscriber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


