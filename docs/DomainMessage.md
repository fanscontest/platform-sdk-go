# DomainMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Sender** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainMessage

`func NewDomainMessage() *DomainMessage`

NewDomainMessage instantiates a new DomainMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainMessageWithDefaults

`func NewDomainMessageWithDefaults() *DomainMessage`

NewDomainMessageWithDefaults instantiates a new DomainMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *DomainMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DomainMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DomainMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DomainMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecipients

`func (o *DomainMessage) GetRecipients() DomainPlatformIdentity`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *DomainMessage) GetRecipientsOk() (*DomainPlatformIdentity, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *DomainMessage) SetRecipients(v DomainPlatformIdentity)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *DomainMessage) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSender

`func (o *DomainMessage) GetSender() DomainPlatformIdentity`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *DomainMessage) GetSenderOk() (*DomainPlatformIdentity, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *DomainMessage) SetSender(v DomainPlatformIdentity)`

SetSender sets Sender field to given value.

### HasSender

`func (o *DomainMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetStatus

`func (o *DomainMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *DomainMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DomainMessage) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DomainMessage) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *DomainMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetType

`func (o *DomainMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DomainMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


