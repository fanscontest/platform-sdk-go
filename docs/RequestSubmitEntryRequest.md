# RequestSubmitEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Optional, for live recording | [optional] 
**MediaType** | **string** |  | 
**MediaUrl** | **string** |  | 

## Methods

### NewRequestSubmitEntryRequest

`func NewRequestSubmitEntryRequest(mediaType string, mediaUrl string, ) *RequestSubmitEntryRequest`

NewRequestSubmitEntryRequest instantiates a new RequestSubmitEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSubmitEntryRequestWithDefaults

`func NewRequestSubmitEntryRequestWithDefaults() *RequestSubmitEntryRequest`

NewRequestSubmitEntryRequestWithDefaults instantiates a new RequestSubmitEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *RequestSubmitEntryRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RequestSubmitEntryRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RequestSubmitEntryRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RequestSubmitEntryRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetMediaType

`func (o *RequestSubmitEntryRequest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RequestSubmitEntryRequest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RequestSubmitEntryRequest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMediaUrl

`func (o *RequestSubmitEntryRequest) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *RequestSubmitEntryRequest) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *RequestSubmitEntryRequest) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


