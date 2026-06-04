# RequestCreateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCode** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | **string** |  | 
**IsFanChannel** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Regions** | Pointer to [**[]RequestRegion**](RequestRegion.md) |  | [optional] 
**Size** | **int32** |  | 

## Methods

### NewRequestCreateChannelRequest

`func NewRequestCreateChannelRequest(headerImageUrl string, name string, size int32, ) *RequestCreateChannelRequest`

NewRequestCreateChannelRequest instantiates a new RequestCreateChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateChannelRequestWithDefaults

`func NewRequestCreateChannelRequestWithDefaults() *RequestCreateChannelRequest`

NewRequestCreateChannelRequestWithDefaults instantiates a new RequestCreateChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCode

`func (o *RequestCreateChannelRequest) GetAccessCode() string`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *RequestCreateChannelRequest) GetAccessCodeOk() (*string, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *RequestCreateChannelRequest) SetAccessCode(v string)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *RequestCreateChannelRequest) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetAccessType

`func (o *RequestCreateChannelRequest) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *RequestCreateChannelRequest) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *RequestCreateChannelRequest) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *RequestCreateChannelRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCode

`func (o *RequestCreateChannelRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RequestCreateChannelRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RequestCreateChannelRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RequestCreateChannelRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *RequestCreateChannelRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *RequestCreateChannelRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *RequestCreateChannelRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.


### GetIsFanChannel

`func (o *RequestCreateChannelRequest) GetIsFanChannel() bool`

GetIsFanChannel returns the IsFanChannel field if non-nil, zero value otherwise.

### GetIsFanChannelOk

`func (o *RequestCreateChannelRequest) GetIsFanChannelOk() (*bool, bool)`

GetIsFanChannelOk returns a tuple with the IsFanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanChannel

`func (o *RequestCreateChannelRequest) SetIsFanChannel(v bool)`

SetIsFanChannel sets IsFanChannel field to given value.

### HasIsFanChannel

`func (o *RequestCreateChannelRequest) HasIsFanChannel() bool`

HasIsFanChannel returns a boolean if a field has been set.

### GetName

`func (o *RequestCreateChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestCreateChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestCreateChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegions

`func (o *RequestCreateChannelRequest) GetRegions() []RequestRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *RequestCreateChannelRequest) GetRegionsOk() (*[]RequestRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *RequestCreateChannelRequest) SetRegions(v []RequestRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *RequestCreateChannelRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *RequestCreateChannelRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RequestCreateChannelRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RequestCreateChannelRequest) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


