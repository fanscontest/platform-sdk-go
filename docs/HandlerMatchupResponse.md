# HandlerMatchupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matchup** | Pointer to [**HandlerServedMatchup**](HandlerServedMatchup.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlerMatchupResponse

`func NewHandlerMatchupResponse() *HandlerMatchupResponse`

NewHandlerMatchupResponse instantiates a new HandlerMatchupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerMatchupResponseWithDefaults

`func NewHandlerMatchupResponseWithDefaults() *HandlerMatchupResponse`

NewHandlerMatchupResponseWithDefaults instantiates a new HandlerMatchupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchup

`func (o *HandlerMatchupResponse) GetMatchup() HandlerServedMatchup`

GetMatchup returns the Matchup field if non-nil, zero value otherwise.

### GetMatchupOk

`func (o *HandlerMatchupResponse) GetMatchupOk() (*HandlerServedMatchup, bool)`

GetMatchupOk returns a tuple with the Matchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchup

`func (o *HandlerMatchupResponse) SetMatchup(v HandlerServedMatchup)`

SetMatchup sets Matchup field to given value.

### HasMatchup

`func (o *HandlerMatchupResponse) HasMatchup() bool`

HasMatchup returns a boolean if a field has been set.

### GetReason

`func (o *HandlerMatchupResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlerMatchupResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlerMatchupResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlerMatchupResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


