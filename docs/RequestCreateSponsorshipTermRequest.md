# RequestCreateSponsorshipTermRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Title** | **string** |  | 
**Version** | Pointer to **string** | Optional, defaults to \&quot;1.0\&quot; | [optional] 

## Methods

### NewRequestCreateSponsorshipTermRequest

`func NewRequestCreateSponsorshipTermRequest(description string, title string, ) *RequestCreateSponsorshipTermRequest`

NewRequestCreateSponsorshipTermRequest instantiates a new RequestCreateSponsorshipTermRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSponsorshipTermRequestWithDefaults

`func NewRequestCreateSponsorshipTermRequestWithDefaults() *RequestCreateSponsorshipTermRequest`

NewRequestCreateSponsorshipTermRequestWithDefaults instantiates a new RequestCreateSponsorshipTermRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RequestCreateSponsorshipTermRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestCreateSponsorshipTermRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestCreateSponsorshipTermRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *RequestCreateSponsorshipTermRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestCreateSponsorshipTermRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestCreateSponsorshipTermRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVersion

`func (o *RequestCreateSponsorshipTermRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RequestCreateSponsorshipTermRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RequestCreateSponsorshipTermRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RequestCreateSponsorshipTermRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


