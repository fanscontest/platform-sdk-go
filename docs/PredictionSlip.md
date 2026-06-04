# PredictionSlip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CuratedTags** | Pointer to **[]string** |  | [optional] 
**Events** | Pointer to [**[]PredictionEvent**](PredictionEvent.md) |  | [optional] 
**Id** | **string** |  | 
**IsCurated** | Pointer to **NullableBool** |  | [optional] 
**IsPublic** | Pointer to **NullableBool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPredictionSlip

`func NewPredictionSlip(id string, name string, ) *PredictionSlip`

NewPredictionSlip instantiates a new PredictionSlip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionSlipWithDefaults

`func NewPredictionSlipWithDefaults() *PredictionSlip`

NewPredictionSlipWithDefaults instantiates a new PredictionSlip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PredictionSlip) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PredictionSlip) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PredictionSlip) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PredictionSlip) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCuratedTags

`func (o *PredictionSlip) GetCuratedTags() []string`

GetCuratedTags returns the CuratedTags field if non-nil, zero value otherwise.

### GetCuratedTagsOk

`func (o *PredictionSlip) GetCuratedTagsOk() (*[]string, bool)`

GetCuratedTagsOk returns a tuple with the CuratedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuratedTags

`func (o *PredictionSlip) SetCuratedTags(v []string)`

SetCuratedTags sets CuratedTags field to given value.

### HasCuratedTags

`func (o *PredictionSlip) HasCuratedTags() bool`

HasCuratedTags returns a boolean if a field has been set.

### GetEvents

`func (o *PredictionSlip) GetEvents() []PredictionEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *PredictionSlip) GetEventsOk() (*[]PredictionEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *PredictionSlip) SetEvents(v []PredictionEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *PredictionSlip) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *PredictionSlip) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PredictionSlip) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PredictionSlip) SetId(v string)`

SetId sets Id field to given value.


### GetIsCurated

`func (o *PredictionSlip) GetIsCurated() bool`

GetIsCurated returns the IsCurated field if non-nil, zero value otherwise.

### GetIsCuratedOk

`func (o *PredictionSlip) GetIsCuratedOk() (*bool, bool)`

GetIsCuratedOk returns a tuple with the IsCurated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurated

`func (o *PredictionSlip) SetIsCurated(v bool)`

SetIsCurated sets IsCurated field to given value.

### HasIsCurated

`func (o *PredictionSlip) HasIsCurated() bool`

HasIsCurated returns a boolean if a field has been set.

### SetIsCuratedNil

`func (o *PredictionSlip) SetIsCuratedNil(b bool)`

 SetIsCuratedNil sets the value for IsCurated to be an explicit nil

### UnsetIsCurated
`func (o *PredictionSlip) UnsetIsCurated()`

UnsetIsCurated ensures that no value is present for IsCurated, not even an explicit nil
### GetIsPublic

`func (o *PredictionSlip) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PredictionSlip) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PredictionSlip) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *PredictionSlip) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *PredictionSlip) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *PredictionSlip) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil
### GetLocked

`func (o *PredictionSlip) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PredictionSlip) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PredictionSlip) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *PredictionSlip) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *PredictionSlip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredictionSlip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredictionSlip) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *PredictionSlip) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PredictionSlip) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PredictionSlip) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PredictionSlip) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


