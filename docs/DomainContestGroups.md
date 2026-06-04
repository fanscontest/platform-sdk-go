# DomainContestGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contest** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**map[string][]DomainPlatformIdentity**](array.md) |  | [optional] 
**MemberGroup** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainContestGroups

`func NewDomainContestGroups() *DomainContestGroups`

NewDomainContestGroups instantiates a new DomainContestGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContestGroupsWithDefaults

`func NewDomainContestGroupsWithDefaults() *DomainContestGroups`

NewDomainContestGroupsWithDefaults instantiates a new DomainContestGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContest

`func (o *DomainContestGroups) GetContest() string`

GetContest returns the Contest field if non-nil, zero value otherwise.

### GetContestOk

`func (o *DomainContestGroups) GetContestOk() (*string, bool)`

GetContestOk returns a tuple with the Contest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContest

`func (o *DomainContestGroups) SetContest(v string)`

SetContest sets Contest field to given value.

### HasContest

`func (o *DomainContestGroups) HasContest() bool`

HasContest returns a boolean if a field has been set.

### GetGroups

`func (o *DomainContestGroups) GetGroups() map[string][]DomainPlatformIdentity`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DomainContestGroups) GetGroupsOk() (*map[string][]DomainPlatformIdentity, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DomainContestGroups) SetGroups(v map[string][]DomainPlatformIdentity)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DomainContestGroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMemberGroup

`func (o *DomainContestGroups) GetMemberGroup() string`

GetMemberGroup returns the MemberGroup field if non-nil, zero value otherwise.

### GetMemberGroupOk

`func (o *DomainContestGroups) GetMemberGroupOk() (*string, bool)`

GetMemberGroupOk returns a tuple with the MemberGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroup

`func (o *DomainContestGroups) SetMemberGroup(v string)`

SetMemberGroup sets MemberGroup field to given value.

### HasMemberGroup

`func (o *DomainContestGroups) HasMemberGroup() bool`

HasMemberGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


