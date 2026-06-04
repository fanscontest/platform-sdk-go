# CreatePuzzleSheetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Puzzles** | Pointer to [**[]PuzzleBody**](PuzzleBody.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Visibility** | **string** |  | 

## Methods

### NewCreatePuzzleSheetRequest

`func NewCreatePuzzleSheetRequest(name string, visibility string, ) *CreatePuzzleSheetRequest`

NewCreatePuzzleSheetRequest instantiates a new CreatePuzzleSheetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePuzzleSheetRequestWithDefaults

`func NewCreatePuzzleSheetRequestWithDefaults() *CreatePuzzleSheetRequest`

NewCreatePuzzleSheetRequestWithDefaults instantiates a new CreatePuzzleSheetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreatePuzzleSheetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePuzzleSheetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePuzzleSheetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePuzzleSheetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatePuzzleSheetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatePuzzleSheetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *CreatePuzzleSheetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePuzzleSheetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePuzzleSheetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPuzzles

`func (o *CreatePuzzleSheetRequest) GetPuzzles() []PuzzleBody`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *CreatePuzzleSheetRequest) GetPuzzlesOk() (*[]PuzzleBody, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *CreatePuzzleSheetRequest) SetPuzzles(v []PuzzleBody)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *CreatePuzzleSheetRequest) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetTags

`func (o *CreatePuzzleSheetRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePuzzleSheetRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePuzzleSheetRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePuzzleSheetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVisibility

`func (o *CreatePuzzleSheetRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreatePuzzleSheetRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreatePuzzleSheetRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


