# EventOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | Pointer to **string** |  | [optional] 
**OptionText** | Pointer to **NullableString** |  | [optional] 
**Points** | Pointer to **float32** |  | [optional] 

## Methods

### NewEventOption

`func NewEventOption() *EventOption`

NewEventOption instantiates a new EventOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventOptionWithDefaults

`func NewEventOptionWithDefaults() *EventOption`

NewEventOptionWithDefaults instantiates a new EventOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *EventOption) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *EventOption) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *EventOption) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *EventOption) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetOptionText

`func (o *EventOption) GetOptionText() string`

GetOptionText returns the OptionText field if non-nil, zero value otherwise.

### GetOptionTextOk

`func (o *EventOption) GetOptionTextOk() (*string, bool)`

GetOptionTextOk returns a tuple with the OptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionText

`func (o *EventOption) SetOptionText(v string)`

SetOptionText sets OptionText field to given value.

### HasOptionText

`func (o *EventOption) HasOptionText() bool`

HasOptionText returns a boolean if a field has been set.

### SetOptionTextNil

`func (o *EventOption) SetOptionTextNil(b bool)`

 SetOptionTextNil sets the value for OptionText to be an explicit nil

### UnsetOptionText
`func (o *EventOption) UnsetOptionText()`

UnsetOptionText ensures that no value is present for OptionText, not even an explicit nil
### GetPoints

`func (o *EventOption) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *EventOption) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *EventOption) SetPoints(v float32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *EventOption) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


