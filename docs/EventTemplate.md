# EventTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EventType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]EventTemplateItem**](EventTemplateItem.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewEventTemplate

`func NewEventTemplate() *EventTemplate`

NewEventTemplate instantiates a new EventTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTemplateWithDefaults

`func NewEventTemplateWithDefaults() *EventTemplate`

NewEventTemplateWithDefaults instantiates a new EventTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *EventTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *EventTemplate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *EventTemplate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *EventTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventType

`func (o *EventTemplate) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventTemplate) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventTemplate) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventTemplate) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *EventTemplate) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *EventTemplate) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetId

`func (o *EventTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *EventTemplate) GetItems() []EventTemplateItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventTemplate) GetItemsOk() (*[]EventTemplateItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventTemplate) SetItems(v []EventTemplateItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventTemplate) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *EventTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


