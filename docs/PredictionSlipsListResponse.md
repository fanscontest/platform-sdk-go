# PredictionSlipsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PredictionSlip**](PredictionSlip.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewPredictionSlipsListResponse

`func NewPredictionSlipsListResponse() *PredictionSlipsListResponse`

NewPredictionSlipsListResponse instantiates a new PredictionSlipsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionSlipsListResponseWithDefaults

`func NewPredictionSlipsListResponseWithDefaults() *PredictionSlipsListResponse`

NewPredictionSlipsListResponseWithDefaults instantiates a new PredictionSlipsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PredictionSlipsListResponse) GetData() []PredictionSlip`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PredictionSlipsListResponse) GetDataOk() (*[]PredictionSlip, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PredictionSlipsListResponse) SetData(v []PredictionSlip)`

SetData sets Data field to given value.

### HasData

`func (o *PredictionSlipsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PredictionSlipsListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PredictionSlipsListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PredictionSlipsListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PredictionSlipsListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


