# PuzzleSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**InsertedAt** | Pointer to **time.Time** |  | [optional] 
**IsAdhocSheet** | Pointer to **NullableBool** |  | [optional] 
**IsCommunityQuestions** | Pointer to **NullableBool** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **string** |  | 
**Puzzles** | Pointer to [**[]PuzzleBody**](PuzzleBody.md) | Embedded puzzle bodies. Empty for adhoc sheet stubs. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Visibility** | **string** |  | 

## Methods

### NewPuzzleSheet

`func NewPuzzleSheet(id string, name string, ownerId string, visibility string, ) *PuzzleSheet`

NewPuzzleSheet instantiates a new PuzzleSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleSheetWithDefaults

`func NewPuzzleSheetWithDefaults() *PuzzleSheet`

NewPuzzleSheetWithDefaults instantiates a new PuzzleSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PuzzleSheet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PuzzleSheet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PuzzleSheet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PuzzleSheet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PuzzleSheet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PuzzleSheet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *PuzzleSheet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleSheet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleSheet) SetId(v string)`

SetId sets Id field to given value.


### GetInsertedAt

`func (o *PuzzleSheet) GetInsertedAt() time.Time`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *PuzzleSheet) GetInsertedAtOk() (*time.Time, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *PuzzleSheet) SetInsertedAt(v time.Time)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *PuzzleSheet) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.

### GetIsAdhocSheet

`func (o *PuzzleSheet) GetIsAdhocSheet() bool`

GetIsAdhocSheet returns the IsAdhocSheet field if non-nil, zero value otherwise.

### GetIsAdhocSheetOk

`func (o *PuzzleSheet) GetIsAdhocSheetOk() (*bool, bool)`

GetIsAdhocSheetOk returns a tuple with the IsAdhocSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdhocSheet

`func (o *PuzzleSheet) SetIsAdhocSheet(v bool)`

SetIsAdhocSheet sets IsAdhocSheet field to given value.

### HasIsAdhocSheet

`func (o *PuzzleSheet) HasIsAdhocSheet() bool`

HasIsAdhocSheet returns a boolean if a field has been set.

### SetIsAdhocSheetNil

`func (o *PuzzleSheet) SetIsAdhocSheetNil(b bool)`

 SetIsAdhocSheetNil sets the value for IsAdhocSheet to be an explicit nil

### UnsetIsAdhocSheet
`func (o *PuzzleSheet) UnsetIsAdhocSheet()`

UnsetIsAdhocSheet ensures that no value is present for IsAdhocSheet, not even an explicit nil
### GetIsCommunityQuestions

`func (o *PuzzleSheet) GetIsCommunityQuestions() bool`

GetIsCommunityQuestions returns the IsCommunityQuestions field if non-nil, zero value otherwise.

### GetIsCommunityQuestionsOk

`func (o *PuzzleSheet) GetIsCommunityQuestionsOk() (*bool, bool)`

GetIsCommunityQuestionsOk returns a tuple with the IsCommunityQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommunityQuestions

`func (o *PuzzleSheet) SetIsCommunityQuestions(v bool)`

SetIsCommunityQuestions sets IsCommunityQuestions field to given value.

### HasIsCommunityQuestions

`func (o *PuzzleSheet) HasIsCommunityQuestions() bool`

HasIsCommunityQuestions returns a boolean if a field has been set.

### SetIsCommunityQuestionsNil

`func (o *PuzzleSheet) SetIsCommunityQuestionsNil(b bool)`

 SetIsCommunityQuestionsNil sets the value for IsCommunityQuestions to be an explicit nil

### UnsetIsCommunityQuestions
`func (o *PuzzleSheet) UnsetIsCommunityQuestions()`

UnsetIsCommunityQuestions ensures that no value is present for IsCommunityQuestions, not even an explicit nil
### GetName

`func (o *PuzzleSheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PuzzleSheet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PuzzleSheet) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *PuzzleSheet) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PuzzleSheet) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PuzzleSheet) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetPuzzles

`func (o *PuzzleSheet) GetPuzzles() []PuzzleBody`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *PuzzleSheet) GetPuzzlesOk() (*[]PuzzleBody, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *PuzzleSheet) SetPuzzles(v []PuzzleBody)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *PuzzleSheet) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetTags

`func (o *PuzzleSheet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PuzzleSheet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PuzzleSheet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PuzzleSheet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PuzzleSheet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PuzzleSheet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PuzzleSheet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PuzzleSheet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *PuzzleSheet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PuzzleSheet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PuzzleSheet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


