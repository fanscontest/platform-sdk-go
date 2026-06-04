# HandlerCreatePlatformChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | **string** |  | 
**ExternalRef** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Size** | **int32** |  | 

## Methods

### NewHandlerCreatePlatformChannelRequest

`func NewHandlerCreatePlatformChannelRequest(accessType string, name string, size int32, ) *HandlerCreatePlatformChannelRequest`

NewHandlerCreatePlatformChannelRequest instantiates a new HandlerCreatePlatformChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreatePlatformChannelRequestWithDefaults

`func NewHandlerCreatePlatformChannelRequestWithDefaults() *HandlerCreatePlatformChannelRequest`

NewHandlerCreatePlatformChannelRequestWithDefaults instantiates a new HandlerCreatePlatformChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *HandlerCreatePlatformChannelRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *HandlerCreatePlatformChannelRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *HandlerCreatePlatformChannelRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetExternalRef

`func (o *HandlerCreatePlatformChannelRequest) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *HandlerCreatePlatformChannelRequest) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *HandlerCreatePlatformChannelRequest) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *HandlerCreatePlatformChannelRequest) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetName

`func (o *HandlerCreatePlatformChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlerCreatePlatformChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlerCreatePlatformChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *HandlerCreatePlatformChannelRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HandlerCreatePlatformChannelRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HandlerCreatePlatformChannelRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


