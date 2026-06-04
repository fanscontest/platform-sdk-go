# HandlerErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlerErrorDetail

`func NewHandlerErrorDetail() *HandlerErrorDetail`

NewHandlerErrorDetail instantiates a new HandlerErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerErrorDetailWithDefaults

`func NewHandlerErrorDetailWithDefaults() *HandlerErrorDetail`

NewHandlerErrorDetailWithDefaults instantiates a new HandlerErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HandlerErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HandlerErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HandlerErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *HandlerErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *HandlerErrorDetail) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HandlerErrorDetail) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HandlerErrorDetail) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HandlerErrorDetail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *HandlerErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HandlerErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HandlerErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HandlerErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


