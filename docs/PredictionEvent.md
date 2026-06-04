# PredictionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**EsEventId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]PredictionItem**](PredictionItem.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPredictionEvent

`func NewPredictionEvent() *PredictionEvent`

NewPredictionEvent instantiates a new PredictionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionEventWithDefaults

`func NewPredictionEventWithDefaults() *PredictionEvent`

NewPredictionEventWithDefaults instantiates a new PredictionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PredictionEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PredictionEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PredictionEvent) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PredictionEvent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *PredictionEvent) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *PredictionEvent) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *PredictionEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PredictionEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PredictionEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PredictionEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PredictionEvent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PredictionEvent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTime

`func (o *PredictionEvent) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *PredictionEvent) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *PredictionEvent) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *PredictionEvent) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *PredictionEvent) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *PredictionEvent) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetEsEventId

`func (o *PredictionEvent) GetEsEventId() string`

GetEsEventId returns the EsEventId field if non-nil, zero value otherwise.

### GetEsEventIdOk

`func (o *PredictionEvent) GetEsEventIdOk() (*string, bool)`

GetEsEventIdOk returns a tuple with the EsEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEventId

`func (o *PredictionEvent) SetEsEventId(v string)`

SetEsEventId sets EsEventId field to given value.

### HasEsEventId

`func (o *PredictionEvent) HasEsEventId() bool`

HasEsEventId returns a boolean if a field has been set.

### SetEsEventIdNil

`func (o *PredictionEvent) SetEsEventIdNil(b bool)`

 SetEsEventIdNil sets the value for EsEventId to be an explicit nil

### UnsetEsEventId
`func (o *PredictionEvent) UnsetEsEventId()`

UnsetEsEventId ensures that no value is present for EsEventId, not even an explicit nil
### GetId

`func (o *PredictionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PredictionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PredictionEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PredictionEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *PredictionEvent) GetItems() []PredictionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PredictionEvent) GetItemsOk() (*[]PredictionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PredictionEvent) SetItems(v []PredictionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *PredictionEvent) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *PredictionEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredictionEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredictionEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PredictionEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *PredictionEvent) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PredictionEvent) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PredictionEvent) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PredictionEvent) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *PredictionEvent) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *PredictionEvent) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


