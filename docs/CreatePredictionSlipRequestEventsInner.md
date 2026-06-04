# CreatePredictionSlipRequestEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsEventId** | **string** |  | 
**Items** | [**[]CreatePredictionSlipRequestEventsInnerItemsInner**](CreatePredictionSlipRequestEventsInnerItemsInner.md) |  | 

## Methods

### NewCreatePredictionSlipRequestEventsInner

`func NewCreatePredictionSlipRequestEventsInner(esEventId string, items []CreatePredictionSlipRequestEventsInnerItemsInner, ) *CreatePredictionSlipRequestEventsInner`

NewCreatePredictionSlipRequestEventsInner instantiates a new CreatePredictionSlipRequestEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePredictionSlipRequestEventsInnerWithDefaults

`func NewCreatePredictionSlipRequestEventsInnerWithDefaults() *CreatePredictionSlipRequestEventsInner`

NewCreatePredictionSlipRequestEventsInnerWithDefaults instantiates a new CreatePredictionSlipRequestEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsEventId

`func (o *CreatePredictionSlipRequestEventsInner) GetEsEventId() string`

GetEsEventId returns the EsEventId field if non-nil, zero value otherwise.

### GetEsEventIdOk

`func (o *CreatePredictionSlipRequestEventsInner) GetEsEventIdOk() (*string, bool)`

GetEsEventIdOk returns a tuple with the EsEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsEventId

`func (o *CreatePredictionSlipRequestEventsInner) SetEsEventId(v string)`

SetEsEventId sets EsEventId field to given value.


### GetItems

`func (o *CreatePredictionSlipRequestEventsInner) GetItems() []CreatePredictionSlipRequestEventsInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreatePredictionSlipRequestEventsInner) GetItemsOk() (*[]CreatePredictionSlipRequestEventsInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreatePredictionSlipRequestEventsInner) SetItems(v []CreatePredictionSlipRequestEventsInnerItemsInner)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


