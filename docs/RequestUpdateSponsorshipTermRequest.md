# RequestUpdateSponsorshipTermRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Title** | **string** |  | 
**Version** | **string** | Required, must be incremented from current version | 

## Methods

### NewRequestUpdateSponsorshipTermRequest

`func NewRequestUpdateSponsorshipTermRequest(description string, title string, version string, ) *RequestUpdateSponsorshipTermRequest`

NewRequestUpdateSponsorshipTermRequest instantiates a new RequestUpdateSponsorshipTermRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdateSponsorshipTermRequestWithDefaults

`func NewRequestUpdateSponsorshipTermRequestWithDefaults() *RequestUpdateSponsorshipTermRequest`

NewRequestUpdateSponsorshipTermRequestWithDefaults instantiates a new RequestUpdateSponsorshipTermRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RequestUpdateSponsorshipTermRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestUpdateSponsorshipTermRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestUpdateSponsorshipTermRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTitle

`func (o *RequestUpdateSponsorshipTermRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestUpdateSponsorshipTermRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestUpdateSponsorshipTermRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVersion

`func (o *RequestUpdateSponsorshipTermRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RequestUpdateSponsorshipTermRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RequestUpdateSponsorshipTermRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


