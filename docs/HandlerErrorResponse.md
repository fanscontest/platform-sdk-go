# HandlerErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**HandlerErrorDetail**](HandlerErrorDetail.md) |  | [optional] 

## Methods

### NewHandlerErrorResponse

`func NewHandlerErrorResponse() *HandlerErrorResponse`

NewHandlerErrorResponse instantiates a new HandlerErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerErrorResponseWithDefaults

`func NewHandlerErrorResponseWithDefaults() *HandlerErrorResponse`

NewHandlerErrorResponseWithDefaults instantiates a new HandlerErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *HandlerErrorResponse) GetError() HandlerErrorDetail`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HandlerErrorResponse) GetErrorOk() (*HandlerErrorDetail, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HandlerErrorResponse) SetError(v HandlerErrorDetail)`

SetError sets Error field to given value.

### HasError

`func (o *HandlerErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


