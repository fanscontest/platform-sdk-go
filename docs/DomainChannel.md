# DomainChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCode** | Pointer to **string** |  | [optional] 
**AccessType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | Unique identifier of channel | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**HeaderImageUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsFanChannel** | Pointer to **bool** |  | [optional] 
**MemberSubscription** | Pointer to [**DomainMemberSubscriptionSummary**](DomainMemberSubscriptionSummary.md) | Set when authenticated; omitted when not subscribed | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Regions** | Pointer to [**[]DomainRegion**](DomainRegion.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Stats** | Pointer to [**DomainStats**](DomainStats.md) |  | [optional] 
**VenueId** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainChannel

`func NewDomainChannel() *DomainChannel`

NewDomainChannel instantiates a new DomainChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainChannelWithDefaults

`func NewDomainChannelWithDefaults() *DomainChannel`

NewDomainChannelWithDefaults instantiates a new DomainChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCode

`func (o *DomainChannel) GetAccessCode() string`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *DomainChannel) GetAccessCodeOk() (*string, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *DomainChannel) SetAccessCode(v string)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *DomainChannel) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetAccessType

`func (o *DomainChannel) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *DomainChannel) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *DomainChannel) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *DomainChannel) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetCode

`func (o *DomainChannel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DomainChannel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DomainChannel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DomainChannel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainChannel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainChannel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainChannel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHeaderImageUrl

`func (o *DomainChannel) GetHeaderImageUrl() string`

GetHeaderImageUrl returns the HeaderImageUrl field if non-nil, zero value otherwise.

### GetHeaderImageUrlOk

`func (o *DomainChannel) GetHeaderImageUrlOk() (*string, bool)`

GetHeaderImageUrlOk returns a tuple with the HeaderImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderImageUrl

`func (o *DomainChannel) SetHeaderImageUrl(v string)`

SetHeaderImageUrl sets HeaderImageUrl field to given value.

### HasHeaderImageUrl

`func (o *DomainChannel) HasHeaderImageUrl() bool`

HasHeaderImageUrl returns a boolean if a field has been set.

### GetId

`func (o *DomainChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFanChannel

`func (o *DomainChannel) GetIsFanChannel() bool`

GetIsFanChannel returns the IsFanChannel field if non-nil, zero value otherwise.

### GetIsFanChannelOk

`func (o *DomainChannel) GetIsFanChannelOk() (*bool, bool)`

GetIsFanChannelOk returns a tuple with the IsFanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFanChannel

`func (o *DomainChannel) SetIsFanChannel(v bool)`

SetIsFanChannel sets IsFanChannel field to given value.

### HasIsFanChannel

`func (o *DomainChannel) HasIsFanChannel() bool`

HasIsFanChannel returns a boolean if a field has been set.

### GetMemberSubscription

`func (o *DomainChannel) GetMemberSubscription() DomainMemberSubscriptionSummary`

GetMemberSubscription returns the MemberSubscription field if non-nil, zero value otherwise.

### GetMemberSubscriptionOk

`func (o *DomainChannel) GetMemberSubscriptionOk() (*DomainMemberSubscriptionSummary, bool)`

GetMemberSubscriptionOk returns a tuple with the MemberSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSubscription

`func (o *DomainChannel) SetMemberSubscription(v DomainMemberSubscriptionSummary)`

SetMemberSubscription sets MemberSubscription field to given value.

### HasMemberSubscription

`func (o *DomainChannel) HasMemberSubscription() bool`

HasMemberSubscription returns a boolean if a field has been set.

### GetName

`func (o *DomainChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *DomainChannel) GetOwner() DomainPlatformIdentity`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DomainChannel) GetOwnerOk() (*DomainPlatformIdentity, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DomainChannel) SetOwner(v DomainPlatformIdentity)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DomainChannel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRegions

`func (o *DomainChannel) GetRegions() []DomainRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *DomainChannel) GetRegionsOk() (*[]DomainRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *DomainChannel) SetRegions(v []DomainRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *DomainChannel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *DomainChannel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DomainChannel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DomainChannel) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DomainChannel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStats

`func (o *DomainChannel) GetStats() DomainStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *DomainChannel) GetStatsOk() (*DomainStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *DomainChannel) SetStats(v DomainStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *DomainChannel) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetVenueId

`func (o *DomainChannel) GetVenueId() string`

GetVenueId returns the VenueId field if non-nil, zero value otherwise.

### GetVenueIdOk

`func (o *DomainChannel) GetVenueIdOk() (*string, bool)`

GetVenueIdOk returns a tuple with the VenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueId

`func (o *DomainChannel) SetVenueId(v string)`

SetVenueId sets VenueId field to given value.

### HasVenueId

`func (o *DomainChannel) HasVenueId() bool`

HasVenueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


