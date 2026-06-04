# CreatePredictionSlipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]CreatePredictionSlipRequestEventsInner**](CreatePredictionSlipRequestEventsInner.md) |  | 
**Name** | **string** |  | 

## Methods

### NewCreatePredictionSlipRequest

`func NewCreatePredictionSlipRequest(events []CreatePredictionSlipRequestEventsInner, name string, ) *CreatePredictionSlipRequest`

NewCreatePredictionSlipRequest instantiates a new CreatePredictionSlipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePredictionSlipRequestWithDefaults

`func NewCreatePredictionSlipRequestWithDefaults() *CreatePredictionSlipRequest`

NewCreatePredictionSlipRequestWithDefaults instantiates a new CreatePredictionSlipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *CreatePredictionSlipRequest) GetEvents() []CreatePredictionSlipRequestEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreatePredictionSlipRequest) GetEventsOk() (*[]CreatePredictionSlipRequestEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreatePredictionSlipRequest) SetEvents(v []CreatePredictionSlipRequestEventsInner)`

SetEvents sets Events field to given value.


### GetName

`func (o *CreatePredictionSlipRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePredictionSlipRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePredictionSlipRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


