# UserEventStatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GradingRequestedAt** | Pointer to **NullableTime** |  | [optional] 
**HasResult** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ResultApproved** | Pointer to **NullableBool** |  | [optional] 
**ResultEnteredAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserEventStatusResponseData

`func NewUserEventStatusResponseData() *UserEventStatusResponseData`

NewUserEventStatusResponseData instantiates a new UserEventStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEventStatusResponseDataWithDefaults

`func NewUserEventStatusResponseDataWithDefaults() *UserEventStatusResponseData`

NewUserEventStatusResponseDataWithDefaults instantiates a new UserEventStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGradingRequestedAt

`func (o *UserEventStatusResponseData) GetGradingRequestedAt() time.Time`

GetGradingRequestedAt returns the GradingRequestedAt field if non-nil, zero value otherwise.

### GetGradingRequestedAtOk

`func (o *UserEventStatusResponseData) GetGradingRequestedAtOk() (*time.Time, bool)`

GetGradingRequestedAtOk returns a tuple with the GradingRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradingRequestedAt

`func (o *UserEventStatusResponseData) SetGradingRequestedAt(v time.Time)`

SetGradingRequestedAt sets GradingRequestedAt field to given value.

### HasGradingRequestedAt

`func (o *UserEventStatusResponseData) HasGradingRequestedAt() bool`

HasGradingRequestedAt returns a boolean if a field has been set.

### SetGradingRequestedAtNil

`func (o *UserEventStatusResponseData) SetGradingRequestedAtNil(b bool)`

 SetGradingRequestedAtNil sets the value for GradingRequestedAt to be an explicit nil

### UnsetGradingRequestedAt
`func (o *UserEventStatusResponseData) UnsetGradingRequestedAt()`

UnsetGradingRequestedAt ensures that no value is present for GradingRequestedAt, not even an explicit nil
### GetHasResult

`func (o *UserEventStatusResponseData) GetHasResult() bool`

GetHasResult returns the HasResult field if non-nil, zero value otherwise.

### GetHasResultOk

`func (o *UserEventStatusResponseData) GetHasResultOk() (*bool, bool)`

GetHasResultOk returns a tuple with the HasResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResult

`func (o *UserEventStatusResponseData) SetHasResult(v bool)`

SetHasResult sets HasResult field to given value.

### HasHasResult

`func (o *UserEventStatusResponseData) HasHasResult() bool`

HasHasResult returns a boolean if a field has been set.

### GetId

`func (o *UserEventStatusResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserEventStatusResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserEventStatusResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserEventStatusResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResultApproved

`func (o *UserEventStatusResponseData) GetResultApproved() bool`

GetResultApproved returns the ResultApproved field if non-nil, zero value otherwise.

### GetResultApprovedOk

`func (o *UserEventStatusResponseData) GetResultApprovedOk() (*bool, bool)`

GetResultApprovedOk returns a tuple with the ResultApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultApproved

`func (o *UserEventStatusResponseData) SetResultApproved(v bool)`

SetResultApproved sets ResultApproved field to given value.

### HasResultApproved

`func (o *UserEventStatusResponseData) HasResultApproved() bool`

HasResultApproved returns a boolean if a field has been set.

### SetResultApprovedNil

`func (o *UserEventStatusResponseData) SetResultApprovedNil(b bool)`

 SetResultApprovedNil sets the value for ResultApproved to be an explicit nil

### UnsetResultApproved
`func (o *UserEventStatusResponseData) UnsetResultApproved()`

UnsetResultApproved ensures that no value is present for ResultApproved, not even an explicit nil
### GetResultEnteredAt

`func (o *UserEventStatusResponseData) GetResultEnteredAt() time.Time`

GetResultEnteredAt returns the ResultEnteredAt field if non-nil, zero value otherwise.

### GetResultEnteredAtOk

`func (o *UserEventStatusResponseData) GetResultEnteredAtOk() (*time.Time, bool)`

GetResultEnteredAtOk returns a tuple with the ResultEnteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultEnteredAt

`func (o *UserEventStatusResponseData) SetResultEnteredAt(v time.Time)`

SetResultEnteredAt sets ResultEnteredAt field to given value.

### HasResultEnteredAt

`func (o *UserEventStatusResponseData) HasResultEnteredAt() bool`

HasResultEnteredAt returns a boolean if a field has been set.

### SetResultEnteredAtNil

`func (o *UserEventStatusResponseData) SetResultEnteredAtNil(b bool)`

 SetResultEnteredAtNil sets the value for ResultEnteredAt to be an explicit nil

### UnsetResultEnteredAt
`func (o *UserEventStatusResponseData) UnsetResultEnteredAt()`

UnsetResultEnteredAt ensures that no value is present for ResultEnteredAt, not even an explicit nil
### GetStatus

`func (o *UserEventStatusResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserEventStatusResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserEventStatusResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserEventStatusResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *UserEventStatusResponseData) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UserEventStatusResponseData) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UserEventStatusResponseData) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UserEventStatusResponseData) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *UserEventStatusResponseData) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *UserEventStatusResponseData) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


