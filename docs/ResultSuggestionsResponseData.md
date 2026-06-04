# ResultSuggestionsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | Pointer to [**[]ResultSuggestionsResponseDataSuggestionsInner**](ResultSuggestionsResponseDataSuggestionsInner.md) |  | [optional] 
**SuggestionsPresent** | Pointer to **bool** |  | [optional] 

## Methods

### NewResultSuggestionsResponseData

`func NewResultSuggestionsResponseData() *ResultSuggestionsResponseData`

NewResultSuggestionsResponseData instantiates a new ResultSuggestionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultSuggestionsResponseDataWithDefaults

`func NewResultSuggestionsResponseDataWithDefaults() *ResultSuggestionsResponseData`

NewResultSuggestionsResponseDataWithDefaults instantiates a new ResultSuggestionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *ResultSuggestionsResponseData) GetSuggestions() []ResultSuggestionsResponseDataSuggestionsInner`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *ResultSuggestionsResponseData) GetSuggestionsOk() (*[]ResultSuggestionsResponseDataSuggestionsInner, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *ResultSuggestionsResponseData) SetSuggestions(v []ResultSuggestionsResponseDataSuggestionsInner)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *ResultSuggestionsResponseData) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetSuggestionsPresent

`func (o *ResultSuggestionsResponseData) GetSuggestionsPresent() bool`

GetSuggestionsPresent returns the SuggestionsPresent field if non-nil, zero value otherwise.

### GetSuggestionsPresentOk

`func (o *ResultSuggestionsResponseData) GetSuggestionsPresentOk() (*bool, bool)`

GetSuggestionsPresentOk returns a tuple with the SuggestionsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionsPresent

`func (o *ResultSuggestionsResponseData) SetSuggestionsPresent(v bool)`

SetSuggestionsPresent sets SuggestionsPresent field to given value.

### HasSuggestionsPresent

`func (o *ResultSuggestionsResponseData) HasSuggestionsPresent() bool`

HasSuggestionsPresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


