# RequestCreateSponsorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SupervisorId** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewRequestCreateSponsorRequest

`func NewRequestCreateSponsorRequest(name string, url string, ) *RequestCreateSponsorRequest`

NewRequestCreateSponsorRequest instantiates a new RequestCreateSponsorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSponsorRequestWithDefaults

`func NewRequestCreateSponsorRequestWithDefaults() *RequestCreateSponsorRequest`

NewRequestCreateSponsorRequestWithDefaults instantiates a new RequestCreateSponsorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestCreateSponsorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestCreateSponsorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestCreateSponsorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSupervisorId

`func (o *RequestCreateSponsorRequest) GetSupervisorId() string`

GetSupervisorId returns the SupervisorId field if non-nil, zero value otherwise.

### GetSupervisorIdOk

`func (o *RequestCreateSponsorRequest) GetSupervisorIdOk() (*string, bool)`

GetSupervisorIdOk returns a tuple with the SupervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorId

`func (o *RequestCreateSponsorRequest) SetSupervisorId(v string)`

SetSupervisorId sets SupervisorId field to given value.

### HasSupervisorId

`func (o *RequestCreateSponsorRequest) HasSupervisorId() bool`

HasSupervisorId returns a boolean if a field has been set.

### GetUrl

`func (o *RequestCreateSponsorRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestCreateSponsorRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestCreateSponsorRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


