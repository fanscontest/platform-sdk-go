# CreatePuzzleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puzzle** | **map[string]interface{}** | Puzzle attributes. owner_id is taken from the authenticated identity, not the body. | 

## Methods

### NewCreatePuzzleRequest

`func NewCreatePuzzleRequest(puzzle map[string]interface{}, ) *CreatePuzzleRequest`

NewCreatePuzzleRequest instantiates a new CreatePuzzleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePuzzleRequestWithDefaults

`func NewCreatePuzzleRequestWithDefaults() *CreatePuzzleRequest`

NewCreatePuzzleRequestWithDefaults instantiates a new CreatePuzzleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuzzle

`func (o *CreatePuzzleRequest) GetPuzzle() map[string]interface{}`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *CreatePuzzleRequest) GetPuzzleOk() (*map[string]interface{}, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *CreatePuzzleRequest) SetPuzzle(v map[string]interface{})`

SetPuzzle sets Puzzle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


