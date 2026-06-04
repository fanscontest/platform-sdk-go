# HandlerUpdatePlatformIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** | CountryCode is optional. Sent as ISO 3166-1 alpha-2; empty value preserves the current value on the row (mirrors the image-URL preservation pattern below). Added in slice 22.6 / uman#107. | [optional] 
**DisplayName** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ProfilePic** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlerUpdatePlatformIdentityRequest

`func NewHandlerUpdatePlatformIdentityRequest(displayName string, ) *HandlerUpdatePlatformIdentityRequest`

NewHandlerUpdatePlatformIdentityRequest instantiates a new HandlerUpdatePlatformIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdatePlatformIdentityRequestWithDefaults

`func NewHandlerUpdatePlatformIdentityRequestWithDefaults() *HandlerUpdatePlatformIdentityRequest`

NewHandlerUpdatePlatformIdentityRequestWithDefaults instantiates a new HandlerUpdatePlatformIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *HandlerUpdatePlatformIdentityRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *HandlerUpdatePlatformIdentityRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *HandlerUpdatePlatformIdentityRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *HandlerUpdatePlatformIdentityRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *HandlerUpdatePlatformIdentityRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *HandlerUpdatePlatformIdentityRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *HandlerUpdatePlatformIdentityRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HandlerUpdatePlatformIdentityRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFirstName

`func (o *HandlerUpdatePlatformIdentityRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *HandlerUpdatePlatformIdentityRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *HandlerUpdatePlatformIdentityRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *HandlerUpdatePlatformIdentityRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *HandlerUpdatePlatformIdentityRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *HandlerUpdatePlatformIdentityRequest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetLastName

`func (o *HandlerUpdatePlatformIdentityRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *HandlerUpdatePlatformIdentityRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *HandlerUpdatePlatformIdentityRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetProfilePic

`func (o *HandlerUpdatePlatformIdentityRequest) GetProfilePic() string`

GetProfilePic returns the ProfilePic field if non-nil, zero value otherwise.

### GetProfilePicOk

`func (o *HandlerUpdatePlatformIdentityRequest) GetProfilePicOk() (*string, bool)`

GetProfilePicOk returns a tuple with the ProfilePic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePic

`func (o *HandlerUpdatePlatformIdentityRequest) SetProfilePic(v string)`

SetProfilePic sets ProfilePic field to given value.

### HasProfilePic

`func (o *HandlerUpdatePlatformIdentityRequest) HasProfilePic() bool`

HasProfilePic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


