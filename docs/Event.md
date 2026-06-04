# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**EsEventId** | Pointer to **NullableString** |  | [optional] 
**Family** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**InsertedAt** | Pointer to **time.Time** |  | [optional] 
**Items** | Pointer to [**[]EventItem**](EventItem.md) |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Provider** | **string** |  | 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** | Result payload — shape depends on event item types. | [optional] 
**ResultApproved** | Pointer to **NullableBool** |  | [optional] 
**ResultApprovedAt** | Pointer to **NullableTime** |  | [optional] 
**ResultApprovedBy** | Pointer to **NullableString** |  | [optional] 
**ResultEnteredAt** | Pointer to **NullableTime** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Visibility** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEvent

`func NewEvent(id string, name string, provider string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Event) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Event) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Event) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Event) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Event) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Event) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *Event) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Event) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Event) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Event) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Event) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Event) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndTime

`func (o *Event) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Event) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Event) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Event) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *Event) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *Event) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetEsEventId

`func (o *Event) GetEsEventId() string`

GetEsEventId returns the EsEventId field if non-nil, zero value otherwise.

### GetEsEventIdOk

`func (o *Event) GetEsEventIdOk() (*string, bool)`

GetEsEventIdOk returns a tuple with the EsEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEventId

`func (o *Event) SetEsEventId(v string)`

SetEsEventId sets EsEventId field to given value.

### HasEsEventId

`func (o *Event) HasEsEventId() bool`

HasEsEventId returns a boolean if a field has been set.

### SetEsEventIdNil

`func (o *Event) SetEsEventIdNil(b bool)`

 SetEsEventIdNil sets the value for EsEventId to be an explicit nil

### UnsetEsEventId
`func (o *Event) UnsetEsEventId()`

UnsetEsEventId ensures that no value is present for EsEventId, not even an explicit nil
### GetFamily

`func (o *Event) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Event) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Event) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Event) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *Event) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *Event) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetInsertedAt

`func (o *Event) GetInsertedAt() time.Time`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *Event) GetInsertedAtOk() (*time.Time, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *Event) SetInsertedAt(v time.Time)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *Event) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.

### GetItems

`func (o *Event) GetItems() []EventItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Event) GetItemsOk() (*[]EventItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Event) SetItems(v []EventItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Event) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *Event) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Event) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Event) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *Event) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Event) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Event) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Event) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *Event) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *Event) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetProvider

`func (o *Event) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Event) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Event) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderId

`func (o *Event) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Event) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Event) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Event) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *Event) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *Event) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetResult

`func (o *Event) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Event) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Event) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *Event) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *Event) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *Event) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultApproved

`func (o *Event) GetResultApproved() bool`

GetResultApproved returns the ResultApproved field if non-nil, zero value otherwise.

### GetResultApprovedOk

`func (o *Event) GetResultApprovedOk() (*bool, bool)`

GetResultApprovedOk returns a tuple with the ResultApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultApproved

`func (o *Event) SetResultApproved(v bool)`

SetResultApproved sets ResultApproved field to given value.

### HasResultApproved

`func (o *Event) HasResultApproved() bool`

HasResultApproved returns a boolean if a field has been set.

### SetResultApprovedNil

`func (o *Event) SetResultApprovedNil(b bool)`

 SetResultApprovedNil sets the value for ResultApproved to be an explicit nil

### UnsetResultApproved
`func (o *Event) UnsetResultApproved()`

UnsetResultApproved ensures that no value is present for ResultApproved, not even an explicit nil
### GetResultApprovedAt

`func (o *Event) GetResultApprovedAt() time.Time`

GetResultApprovedAt returns the ResultApprovedAt field if non-nil, zero value otherwise.

### GetResultApprovedAtOk

`func (o *Event) GetResultApprovedAtOk() (*time.Time, bool)`

GetResultApprovedAtOk returns a tuple with the ResultApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultApprovedAt

`func (o *Event) SetResultApprovedAt(v time.Time)`

SetResultApprovedAt sets ResultApprovedAt field to given value.

### HasResultApprovedAt

`func (o *Event) HasResultApprovedAt() bool`

HasResultApprovedAt returns a boolean if a field has been set.

### SetResultApprovedAtNil

`func (o *Event) SetResultApprovedAtNil(b bool)`

 SetResultApprovedAtNil sets the value for ResultApprovedAt to be an explicit nil

### UnsetResultApprovedAt
`func (o *Event) UnsetResultApprovedAt()`

UnsetResultApprovedAt ensures that no value is present for ResultApprovedAt, not even an explicit nil
### GetResultApprovedBy

`func (o *Event) GetResultApprovedBy() string`

GetResultApprovedBy returns the ResultApprovedBy field if non-nil, zero value otherwise.

### GetResultApprovedByOk

`func (o *Event) GetResultApprovedByOk() (*string, bool)`

GetResultApprovedByOk returns a tuple with the ResultApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultApprovedBy

`func (o *Event) SetResultApprovedBy(v string)`

SetResultApprovedBy sets ResultApprovedBy field to given value.

### HasResultApprovedBy

`func (o *Event) HasResultApprovedBy() bool`

HasResultApprovedBy returns a boolean if a field has been set.

### SetResultApprovedByNil

`func (o *Event) SetResultApprovedByNil(b bool)`

 SetResultApprovedByNil sets the value for ResultApprovedBy to be an explicit nil

### UnsetResultApprovedBy
`func (o *Event) UnsetResultApprovedBy()`

UnsetResultApprovedBy ensures that no value is present for ResultApprovedBy, not even an explicit nil
### GetResultEnteredAt

`func (o *Event) GetResultEnteredAt() time.Time`

GetResultEnteredAt returns the ResultEnteredAt field if non-nil, zero value otherwise.

### GetResultEnteredAtOk

`func (o *Event) GetResultEnteredAtOk() (*time.Time, bool)`

GetResultEnteredAtOk returns a tuple with the ResultEnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEnteredAt

`func (o *Event) SetResultEnteredAt(v time.Time)`

SetResultEnteredAt sets ResultEnteredAt field to given value.

### HasResultEnteredAt

`func (o *Event) HasResultEnteredAt() bool`

HasResultEnteredAt returns a boolean if a field has been set.

### SetResultEnteredAtNil

`func (o *Event) SetResultEnteredAtNil(b bool)`

 SetResultEnteredAtNil sets the value for ResultEnteredAt to be an explicit nil

### UnsetResultEnteredAt
`func (o *Event) UnsetResultEnteredAt()`

UnsetResultEnteredAt ensures that no value is present for ResultEnteredAt, not even an explicit nil
### GetStartTime

`func (o *Event) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Event) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *Event) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Event) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStatus

`func (o *Event) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Event) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateId

`func (o *Event) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Event) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Event) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Event) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *Event) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Event) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetUpdatedAt

`func (o *Event) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Event) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Event) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Event) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *Event) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Event) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Event) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Event) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *Event) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *Event) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


