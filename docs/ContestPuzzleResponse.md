# ContestPuzzleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | The contest&#39;s source rendered as-is. One of three shapes by source_type: a PuzzleSheet (see PuzzleSheet schema), a PredictionSlip (see PredictionSlip schema), or a creative-challenge map. Modelled as a free-form object because the bare endpoint returns the source directly; inspect by source_type. | [optional] 

## Methods

### NewContestPuzzleResponse

`func NewContestPuzzleResponse() *ContestPuzzleResponse`

NewContestPuzzleResponse instantiates a new ContestPuzzleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContestPuzzleResponseWithDefaults

`func NewContestPuzzleResponseWithDefaults() *ContestPuzzleResponse`

NewContestPuzzleResponseWithDefaults instantiates a new ContestPuzzleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContestPuzzleResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContestPuzzleResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContestPuzzleResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ContestPuzzleResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


