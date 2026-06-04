# DomainBuddyBoardInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardId** | Pointer to **string** |  | [optional] 
**ContestId** | Pointer to **string** |  | [optional] 
**InvitedAt** | Pointer to **string** |  | [optional] 
**InvitedById** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**RespondedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainBuddyBoardInvite

`func NewDomainBuddyBoardInvite() *DomainBuddyBoardInvite`

NewDomainBuddyBoardInvite instantiates a new DomainBuddyBoardInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBuddyBoardInviteWithDefaults

`func NewDomainBuddyBoardInviteWithDefaults() *DomainBuddyBoardInvite`

NewDomainBuddyBoardInviteWithDefaults instantiates a new DomainBuddyBoardInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardId

`func (o *DomainBuddyBoardInvite) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *DomainBuddyBoardInvite) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *DomainBuddyBoardInvite) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *DomainBuddyBoardInvite) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetContestId

`func (o *DomainBuddyBoardInvite) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainBuddyBoardInvite) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainBuddyBoardInvite) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainBuddyBoardInvite) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetInvitedAt

`func (o *DomainBuddyBoardInvite) GetInvitedAt() string`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *DomainBuddyBoardInvite) GetInvitedAtOk() (*string, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *DomainBuddyBoardInvite) SetInvitedAt(v string)`

SetInvitedAt sets InvitedAt field to given value.

### HasInvitedAt

`func (o *DomainBuddyBoardInvite) HasInvitedAt() bool`

HasInvitedAt returns a boolean if a field has been set.

### GetInvitedById

`func (o *DomainBuddyBoardInvite) GetInvitedById() string`

GetInvitedById returns the InvitedById field if non-nil, zero value otherwise.

### GetInvitedByIdOk

`func (o *DomainBuddyBoardInvite) GetInvitedByIdOk() (*string, bool)`

GetInvitedByIdOk returns a tuple with the InvitedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedById

`func (o *DomainBuddyBoardInvite) SetInvitedById(v string)`

SetInvitedById sets InvitedById field to given value.

### HasInvitedById

`func (o *DomainBuddyBoardInvite) HasInvitedById() bool`

HasInvitedById returns a boolean if a field has been set.

### GetMember

`func (o *DomainBuddyBoardInvite) GetMember() DomainPlatformIdentity`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *DomainBuddyBoardInvite) GetMemberOk() (*DomainPlatformIdentity, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *DomainBuddyBoardInvite) SetMember(v DomainPlatformIdentity)`

SetMember sets Member field to given value.

### HasMember

`func (o *DomainBuddyBoardInvite) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetRespondedAt

`func (o *DomainBuddyBoardInvite) GetRespondedAt() string`

GetRespondedAt returns the RespondedAt field if non-nil, zero value otherwise.

### GetRespondedAtOk

`func (o *DomainBuddyBoardInvite) GetRespondedAtOk() (*string, bool)`

GetRespondedAtOk returns a tuple with the RespondedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondedAt

`func (o *DomainBuddyBoardInvite) SetRespondedAt(v string)`

SetRespondedAt sets RespondedAt field to given value.

### HasRespondedAt

`func (o *DomainBuddyBoardInvite) HasRespondedAt() bool`

HasRespondedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DomainBuddyBoardInvite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainBuddyBoardInvite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainBuddyBoardInvite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainBuddyBoardInvite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


