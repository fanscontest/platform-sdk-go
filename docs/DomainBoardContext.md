# DomainBoardContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardType** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to [**DomainChannel**](DomainChannel.md) |  | [optional] 

## Methods

### NewDomainBoardContext

`func NewDomainBoardContext() *DomainBoardContext`

NewDomainBoardContext instantiates a new DomainBoardContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBoardContextWithDefaults

`func NewDomainBoardContextWithDefaults() *DomainBoardContext`

NewDomainBoardContextWithDefaults instantiates a new DomainBoardContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardType

`func (o *DomainBoardContext) GetBoardType() string`

GetBoardType returns the BoardType field if non-nil, zero value otherwise.

### GetBoardTypeOk

`func (o *DomainBoardContext) GetBoardTypeOk() (*string, bool)`

GetBoardTypeOk returns a tuple with the BoardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardType

`func (o *DomainBoardContext) SetBoardType(v string)`

SetBoardType sets BoardType field to given value.

### HasBoardType

`func (o *DomainBoardContext) HasBoardType() bool`

HasBoardType returns a boolean if a field has been set.

### GetChannel

`func (o *DomainBoardContext) GetChannel() DomainChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DomainBoardContext) GetChannelOk() (*DomainChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DomainBoardContext) SetChannel(v DomainChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DomainBoardContext) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


