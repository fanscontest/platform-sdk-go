# EventItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**DoubleUp** | Pointer to **NullableBool** |  | [optional] 
**EventTemplateItemId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsSelected** | Pointer to **NullableBool** |  | [optional] 
**ItemKind** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]EventOption**](EventOption.md) |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventItem

`func NewEventItem() *EventItem`

NewEventItem instantiates a new EventItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventItemWithDefaults

`func NewEventItemWithDefaults() *EventItem`

NewEventItemWithDefaults instantiates a new EventItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EventItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDoubleUp

`func (o *EventItem) GetDoubleUp() bool`

GetDoubleUp returns the DoubleUp field if non-nil, zero value otherwise.

### GetDoubleUpOk

`func (o *EventItem) GetDoubleUpOk() (*bool, bool)`

GetDoubleUpOk returns a tuple with the DoubleUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleUp

`func (o *EventItem) SetDoubleUp(v bool)`

SetDoubleUp sets DoubleUp field to given value.

### HasDoubleUp

`func (o *EventItem) HasDoubleUp() bool`

HasDoubleUp returns a boolean if a field has been set.

### SetDoubleUpNil

`func (o *EventItem) SetDoubleUpNil(b bool)`

 SetDoubleUpNil sets the value for DoubleUp to be an explicit nil

### UnsetDoubleUp
`func (o *EventItem) UnsetDoubleUp()`

UnsetDoubleUp ensures that no value is present for DoubleUp, not even an explicit nil
### GetEventTemplateItemId

`func (o *EventItem) GetEventTemplateItemId() string`

GetEventTemplateItemId returns the EventTemplateItemId field if non-nil, zero value otherwise.

### GetEventTemplateItemIdOk

`func (o *EventItem) GetEventTemplateItemIdOk() (*string, bool)`

GetEventTemplateItemIdOk returns a tuple with the EventTemplateItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTemplateItemId

`func (o *EventItem) SetEventTemplateItemId(v string)`

SetEventTemplateItemId sets EventTemplateItemId field to given value.

### HasEventTemplateItemId

`func (o *EventItem) HasEventTemplateItemId() bool`

HasEventTemplateItemId returns a boolean if a field has been set.

### SetEventTemplateItemIdNil

`func (o *EventItem) SetEventTemplateItemIdNil(b bool)`

 SetEventTemplateItemIdNil sets the value for EventTemplateItemId to be an explicit nil

### UnsetEventTemplateItemId
`func (o *EventItem) UnsetEventTemplateItemId()`

UnsetEventTemplateItemId ensures that no value is present for EventTemplateItemId, not even an explicit nil
### GetId

`func (o *EventItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSelected

`func (o *EventItem) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *EventItem) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *EventItem) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.

### HasIsSelected

`func (o *EventItem) HasIsSelected() bool`

HasIsSelected returns a boolean if a field has been set.

### SetIsSelectedNil

`func (o *EventItem) SetIsSelectedNil(b bool)`

 SetIsSelectedNil sets the value for IsSelected to be an explicit nil

### UnsetIsSelected
`func (o *EventItem) UnsetIsSelected()`

UnsetIsSelected ensures that no value is present for IsSelected, not even an explicit nil
### GetItemKind

`func (o *EventItem) GetItemKind() string`

GetItemKind returns the ItemKind field if non-nil, zero value otherwise.

### GetItemKindOk

`func (o *EventItem) GetItemKindOk() (*string, bool)`

GetItemKindOk returns a tuple with the ItemKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKind

`func (o *EventItem) SetItemKind(v string)`

SetItemKind sets ItemKind field to given value.

### HasItemKind

`func (o *EventItem) HasItemKind() bool`

HasItemKind returns a boolean if a field has been set.

### SetItemKindNil

`func (o *EventItem) SetItemKindNil(b bool)`

 SetItemKindNil sets the value for ItemKind to be an explicit nil

### UnsetItemKind
`func (o *EventItem) UnsetItemKind()`

UnsetItemKind ensures that no value is present for ItemKind, not even an explicit nil
### GetLabel

`func (o *EventItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EventItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EventItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EventItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *EventItem) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *EventItem) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOptions

`func (o *EventItem) GetOptions() []EventOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EventItem) GetOptionsOk() (*[]EventOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EventItem) SetOptions(v []EventOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EventItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrder

`func (o *EventItem) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EventItem) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EventItem) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *EventItem) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *EventItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EventItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EventItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


