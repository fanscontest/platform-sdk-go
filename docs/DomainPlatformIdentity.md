# DomainPlatformIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**DomainCountry**](DomainCountry.md) |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsOfficial** | Pointer to **bool** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**ProfilePic** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TenantUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainPlatformIdentity

`func NewDomainPlatformIdentity() *DomainPlatformIdentity`

NewDomainPlatformIdentity instantiates a new DomainPlatformIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPlatformIdentityWithDefaults

`func NewDomainPlatformIdentityWithDefaults() *DomainPlatformIdentity`

NewDomainPlatformIdentityWithDefaults instantiates a new DomainPlatformIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *DomainPlatformIdentity) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DomainPlatformIdentity) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DomainPlatformIdentity) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DomainPlatformIdentity) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *DomainPlatformIdentity) GetCountry() DomainCountry`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DomainPlatformIdentity) GetCountryOk() (*DomainCountry, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DomainPlatformIdentity) SetCountry(v DomainCountry)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DomainPlatformIdentity) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *DomainPlatformIdentity) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DomainPlatformIdentity) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DomainPlatformIdentity) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DomainPlatformIdentity) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainPlatformIdentity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainPlatformIdentity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainPlatformIdentity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainPlatformIdentity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *DomainPlatformIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DomainPlatformIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DomainPlatformIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DomainPlatformIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFirstName

`func (o *DomainPlatformIdentity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DomainPlatformIdentity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DomainPlatformIdentity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DomainPlatformIdentity) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *DomainPlatformIdentity) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *DomainPlatformIdentity) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *DomainPlatformIdentity) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *DomainPlatformIdentity) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetId

`func (o *DomainPlatformIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainPlatformIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainPlatformIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainPlatformIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOfficial

`func (o *DomainPlatformIdentity) GetIsOfficial() bool`

GetIsOfficial returns the IsOfficial field if non-nil, zero value otherwise.

### GetIsOfficialOk

`func (o *DomainPlatformIdentity) GetIsOfficialOk() (*bool, bool)`

GetIsOfficialOk returns a tuple with the IsOfficial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficial

`func (o *DomainPlatformIdentity) SetIsOfficial(v bool)`

SetIsOfficial sets IsOfficial field to given value.

### HasIsOfficial

`func (o *DomainPlatformIdentity) HasIsOfficial() bool`

HasIsOfficial returns a boolean if a field has been set.

### GetLastName

`func (o *DomainPlatformIdentity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DomainPlatformIdentity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DomainPlatformIdentity) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DomainPlatformIdentity) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetProfilePic

`func (o *DomainPlatformIdentity) GetProfilePic() string`

GetProfilePic returns the ProfilePic field if non-nil, zero value otherwise.

### GetProfilePicOk

`func (o *DomainPlatformIdentity) GetProfilePicOk() (*string, bool)`

GetProfilePicOk returns a tuple with the ProfilePic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePic

`func (o *DomainPlatformIdentity) SetProfilePic(v string)`

SetProfilePic sets ProfilePic field to given value.

### HasProfilePic

`func (o *DomainPlatformIdentity) HasProfilePic() bool`

HasProfilePic returns a boolean if a field has been set.

### GetTenantId

`func (o *DomainPlatformIdentity) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DomainPlatformIdentity) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DomainPlatformIdentity) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DomainPlatformIdentity) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantUserId

`func (o *DomainPlatformIdentity) GetTenantUserId() string`

GetTenantUserId returns the TenantUserId field if non-nil, zero value otherwise.

### GetTenantUserIdOk

`func (o *DomainPlatformIdentity) GetTenantUserIdOk() (*string, bool)`

GetTenantUserIdOk returns a tuple with the TenantUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUserId

`func (o *DomainPlatformIdentity) SetTenantUserId(v string)`

SetTenantUserId sets TenantUserId field to given value.

### HasTenantUserId

`func (o *DomainPlatformIdentity) HasTenantUserId() bool`

HasTenantUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


