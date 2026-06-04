# RequestUpdateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | **string** |  | 
**IsFanChannel** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Regions** | Pointer to [**[]RequestRegion**](RequestRegion.md) |  | [optional] 
**Size** | **int32** |  | 

## Methods

### NewRequestUpdateChannelRequest

`func NewRequestUpdateChannelRequest(headerImageUrl string, name string, size int32, ) *RequestUpdateChannelRequest`

NewRequestUpdateChannelRequest instantiates a new RequestUpdateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdateChannelRequestWithDefaults

`func NewRequestUpdateChannelRequestWithDefaults() *RequestUpdateChannelRequest`

NewRequestUpdateChannelRequestWithDefaults instantiates a new RequestUpdateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *RequestUpdateChannelRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *RequestUpdateChannelRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *RequestUpdateChannelRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *RequestUpdateChannelRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *RequestUpdateChannelRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *RequestUpdateChannelRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *RequestUpdateChannelRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.


### GetIsFanChannel

`func (o *RequestUpdateChannelRequest) GetIsFanChannel() bool`

GetIsFanChannel returns the IsFanChannel field if non-nil, zero value otherwise.

### GetIsFanChannelOk

`func (o *RequestUpdateChannelRequest) GetIsFanChannelOk() (*bool, bool)`

GetIsFanChannelOk returns a tuple with the IsFanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanChannel

`func (o *RequestUpdateChannelRequest) SetIsFanChannel(v bool)`

SetIsFanChannel sets IsFanChannel field to given value.

### HasIsFanChannel

`func (o *RequestUpdateChannelRequest) HasIsFanChannel() bool`

HasIsFanChannel returns a boolean if a field has been set.

### GetName

`func (o *RequestUpdateChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestUpdateChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestUpdateChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegions

`func (o *RequestUpdateChannelRequest) GetRegions() []RequestRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *RequestUpdateChannelRequest) GetRegionsOk() (*[]RequestRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *RequestUpdateChannelRequest) SetRegions(v []RequestRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *RequestUpdateChannelRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *RequestUpdateChannelRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RequestUpdateChannelRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RequestUpdateChannelRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


