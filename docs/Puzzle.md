# Puzzle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **[]string** |  | [optional] 
**Difficulty** | Pointer to **NullableString** |  | [optional] 
**ExtRefId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**InsertedAt** | Pointer to **time.Time** |  | [optional] 
**OwnerId** | **string** |  | 
**Puzzle** | [**PuzzleBody**](PuzzleBody.md) |  | 
**PuzzleType** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Visibility** | **string** |  | 

## Methods

### NewPuzzle

`func NewPuzzle(id string, ownerId string, puzzle PuzzleBody, puzzleType string, visibility string, ) *Puzzle`

NewPuzzle instantiates a new Puzzle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleWithDefaults

`func NewPuzzleWithDefaults() *Puzzle`

NewPuzzleWithDefaults instantiates a new Puzzle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Puzzle) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Puzzle) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Puzzle) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Puzzle) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDifficulty

`func (o *Puzzle) GetDifficulty() string`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *Puzzle) GetDifficultyOk() (*string, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *Puzzle) SetDifficulty(v string)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *Puzzle) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### SetDifficultyNil

`func (o *Puzzle) SetDifficultyNil(b bool)`

 SetDifficultyNil sets the value for Difficulty to be an explicit nil

### UnsetDifficulty
`func (o *Puzzle) UnsetDifficulty()`

UnsetDifficulty ensures that no value is present for Difficulty, not even an explicit nil
### GetExtRefId

`func (o *Puzzle) GetExtRefId() string`

GetExtRefId returns the ExtRefId field if non-nil, zero value otherwise.

### GetExtRefIdOk

`func (o *Puzzle) GetExtRefIdOk() (*string, bool)`

GetExtRefIdOk returns a tuple with the ExtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtRefId

`func (o *Puzzle) SetExtRefId(v string)`

SetExtRefId sets ExtRefId field to given value.

### HasExtRefId

`func (o *Puzzle) HasExtRefId() bool`

HasExtRefId returns a boolean if a field has been set.

### SetExtRefIdNil

`func (o *Puzzle) SetExtRefIdNil(b bool)`

 SetExtRefIdNil sets the value for ExtRefId to be an explicit nil

### UnsetExtRefId
`func (o *Puzzle) UnsetExtRefId()`

UnsetExtRefId ensures that no value is present for ExtRefId, not even an explicit nil
### GetId

`func (o *Puzzle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Puzzle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Puzzle) SetId(v string)`

SetId sets Id field to given value.


### GetInsertedAt

`func (o *Puzzle) GetInsertedAt() time.Time`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *Puzzle) GetInsertedAtOk() (*time.Time, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *Puzzle) SetInsertedAt(v time.Time)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *Puzzle) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.

### GetOwnerId

`func (o *Puzzle) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Puzzle) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Puzzle) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetPuzzle

`func (o *Puzzle) GetPuzzle() PuzzleBody`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *Puzzle) GetPuzzleOk() (*PuzzleBody, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *Puzzle) SetPuzzle(v PuzzleBody)`

SetPuzzle sets Puzzle field to given value.


### GetPuzzleType

`func (o *Puzzle) GetPuzzleType() string`

GetPuzzleType returns the PuzzleType field if non-nil, zero value otherwise.

### GetPuzzleTypeOk

`func (o *Puzzle) GetPuzzleTypeOk() (*string, bool)`

GetPuzzleTypeOk returns a tuple with the PuzzleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzleType

`func (o *Puzzle) SetPuzzleType(v string)`

SetPuzzleType sets PuzzleType field to given value.


### GetTags

`func (o *Puzzle) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Puzzle) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Puzzle) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Puzzle) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Puzzle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Puzzle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Puzzle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Puzzle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *Puzzle) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Puzzle) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Puzzle) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


