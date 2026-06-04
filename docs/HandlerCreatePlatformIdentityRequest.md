# HandlerCreatePlatformIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** | CountryCode is an optional ISO 3166-1 alpha-2 region the tenant supplies to scope geo-eligibility on the fan&#39;s contests and channels. Empty is permitted — geo queries tolerate it by returning the wider unscoped result. Added in slice 22.6 / uman#107 (Member vestige cleanup). | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ProfilePic** | Pointer to **string** |  | [optional] 
**TenantUserId** | **string** |  | 

## Methods

### NewHandlerCreatePlatformIdentityRequest

`func NewHandlerCreatePlatformIdentityRequest(tenantUserId string, ) *HandlerCreatePlatformIdentityRequest`

NewHandlerCreatePlatformIdentityRequest instantiates a new HandlerCreatePlatformIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreatePlatformIdentityRequestWithDefaults

`func NewHandlerCreatePlatformIdentityRequestWithDefaults() *HandlerCreatePlatformIdentityRequest`

NewHandlerCreatePlatformIdentityRequestWithDefaults instantiates a new HandlerCreatePlatformIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *HandlerCreatePlatformIdentityRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *HandlerCreatePlatformIdentityRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *HandlerCreatePlatformIdentityRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *HandlerCreatePlatformIdentityRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *HandlerCreatePlatformIdentityRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *HandlerCreatePlatformIdentityRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *HandlerCreatePlatformIdentityRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *HandlerCreatePlatformIdentityRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *HandlerCreatePlatformIdentityRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HandlerCreatePlatformIdentityRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HandlerCreatePlatformIdentityRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HandlerCreatePlatformIdentityRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFirstName

`func (o *HandlerCreatePlatformIdentityRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *HandlerCreatePlatformIdentityRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *HandlerCreatePlatformIdentityRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *HandlerCreatePlatformIdentityRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *HandlerCreatePlatformIdentityRequest) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *HandlerCreatePlatformIdentityRequest) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *HandlerCreatePlatformIdentityRequest) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *HandlerCreatePlatformIdentityRequest) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetLastName

`func (o *HandlerCreatePlatformIdentityRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *HandlerCreatePlatformIdentityRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *HandlerCreatePlatformIdentityRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *HandlerCreatePlatformIdentityRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetProfilePic

`func (o *HandlerCreatePlatformIdentityRequest) GetProfilePic() string`

GetProfilePic returns the ProfilePic field if non-nil, zero value otherwise.

### GetProfilePicOk

`func (o *HandlerCreatePlatformIdentityRequest) GetProfilePicOk() (*string, bool)`

GetProfilePicOk returns a tuple with the ProfilePic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePic

`func (o *HandlerCreatePlatformIdentityRequest) SetProfilePic(v string)`

SetProfilePic sets ProfilePic field to given value.

### HasProfilePic

`func (o *HandlerCreatePlatformIdentityRequest) HasProfilePic() bool`

HasProfilePic returns a boolean if a field has been set.

### GetTenantUserId

`func (o *HandlerCreatePlatformIdentityRequest) GetTenantUserId() string`

GetTenantUserId returns the TenantUserId field if non-nil, zero value otherwise.

### GetTenantUserIdOk

`func (o *HandlerCreatePlatformIdentityRequest) GetTenantUserIdOk() (*string, bool)`

GetTenantUserIdOk returns a tuple with the TenantUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUserId

`func (o *HandlerCreatePlatformIdentityRequest) SetTenantUserId(v string)`

SetTenantUserId sets TenantUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


