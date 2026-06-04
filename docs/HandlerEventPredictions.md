# HandlerEventPredictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]HandlerPredictionItem**](HandlerPredictionItem.md) |  | [optional] 

## Methods

### NewHandlerEventPredictions

`func NewHandlerEventPredictions() *HandlerEventPredictions`

NewHandlerEventPredictions instantiates a new HandlerEventPredictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerEventPredictionsWithDefaults

`func NewHandlerEventPredictionsWithDefaults() *HandlerEventPredictions`

NewHandlerEventPredictionsWithDefaults instantiates a new HandlerEventPredictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *HandlerEventPredictions) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *HandlerEventPredictions) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *HandlerEventPredictions) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *HandlerEventPredictions) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetItems

`func (o *HandlerEventPredictions) GetItems() []HandlerPredictionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HandlerEventPredictions) GetItemsOk() (*[]HandlerPredictionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HandlerEventPredictions) SetItems(v []HandlerPredictionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *HandlerEventPredictions) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


