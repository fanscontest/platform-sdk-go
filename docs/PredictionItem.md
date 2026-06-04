# PredictionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**PredictionItemConfiguration**](PredictionItemConfiguration.md) |  | [optional] 
**CorrectAnswer** | Pointer to **map[string]interface{}** | Populated post-grading; null otherwise. | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]PredictionItemOptionsInner**](PredictionItemOptionsInner.md) |  | [optional] 

## Methods

### NewPredictionItem

`func NewPredictionItem() *PredictionItem`

NewPredictionItem instantiates a new PredictionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionItemWithDefaults

`func NewPredictionItemWithDefaults() *PredictionItem`

NewPredictionItemWithDefaults instantiates a new PredictionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PredictionItem) GetConfiguration() PredictionItemConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PredictionItem) GetConfigurationOk() (*PredictionItemConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PredictionItem) SetConfiguration(v PredictionItemConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PredictionItem) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCorrectAnswer

`func (o *PredictionItem) GetCorrectAnswer() map[string]interface{}`

GetCorrectAnswer returns the CorrectAnswer field if non-nil, zero value otherwise.

### GetCorrectAnswerOk

`func (o *PredictionItem) GetCorrectAnswerOk() (*map[string]interface{}, bool)`

GetCorrectAnswerOk returns a tuple with the CorrectAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectAnswer

`func (o *PredictionItem) SetCorrectAnswer(v map[string]interface{})`

SetCorrectAnswer sets CorrectAnswer field to given value.

### HasCorrectAnswer

`func (o *PredictionItem) HasCorrectAnswer() bool`

HasCorrectAnswer returns a boolean if a field has been set.

### SetCorrectAnswerNil

`func (o *PredictionItem) SetCorrectAnswerNil(b bool)`

 SetCorrectAnswerNil sets the value for CorrectAnswer to be an explicit nil

### UnsetCorrectAnswer
`func (o *PredictionItem) UnsetCorrectAnswer()`

UnsetCorrectAnswer ensures that no value is present for CorrectAnswer, not even an explicit nil
### GetDescription

`func (o *PredictionItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PredictionItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PredictionItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PredictionItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PredictionItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PredictionItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *PredictionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PredictionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PredictionItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PredictionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PredictionItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PredictionItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PredictionItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PredictionItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PredictionItem) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PredictionItem) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOptions

`func (o *PredictionItem) GetOptions() []PredictionItemOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PredictionItem) GetOptionsOk() (*[]PredictionItemOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PredictionItem) SetOptions(v []PredictionItemOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PredictionItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


