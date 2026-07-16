# DomainCommunityBallot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChoiceEntryId** | Pointer to **string** |  | [optional] 
**ContestId** | Pointer to **string** |  | [optional] 
**EntryIds** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PlatformIdentityId** | Pointer to **string** |  | [optional] 
**ServedAt** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | omitted once redeemed (the secret is burned) | [optional] 

## Methods

### NewDomainCommunityBallot

`func NewDomainCommunityBallot() *DomainCommunityBallot`

NewDomainCommunityBallot instantiates a new DomainCommunityBallot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCommunityBallotWithDefaults

`func NewDomainCommunityBallotWithDefaults() *DomainCommunityBallot`

NewDomainCommunityBallotWithDefaults instantiates a new DomainCommunityBallot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoiceEntryId

`func (o *DomainCommunityBallot) GetChoiceEntryId() string`

GetChoiceEntryId returns the ChoiceEntryId field if non-nil, zero value otherwise.

### GetChoiceEntryIdOk

`func (o *DomainCommunityBallot) GetChoiceEntryIdOk() (*string, bool)`

GetChoiceEntryIdOk returns a tuple with the ChoiceEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoiceEntryId

`func (o *DomainCommunityBallot) SetChoiceEntryId(v string)`

SetChoiceEntryId sets ChoiceEntryId field to given value.

### HasChoiceEntryId

`func (o *DomainCommunityBallot) HasChoiceEntryId() bool`

HasChoiceEntryId returns a boolean if a field has been set.

### GetContestId

`func (o *DomainCommunityBallot) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainCommunityBallot) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainCommunityBallot) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainCommunityBallot) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetEntryIds

`func (o *DomainCommunityBallot) GetEntryIds() []string`

GetEntryIds returns the EntryIds field if non-nil, zero value otherwise.

### GetEntryIdsOk

`func (o *DomainCommunityBallot) GetEntryIdsOk() (*[]string, bool)`

GetEntryIdsOk returns a tuple with the EntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryIds

`func (o *DomainCommunityBallot) SetEntryIds(v []string)`

SetEntryIds sets EntryIds field to given value.

### HasEntryIds

`func (o *DomainCommunityBallot) HasEntryIds() bool`

HasEntryIds returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DomainCommunityBallot) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DomainCommunityBallot) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DomainCommunityBallot) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DomainCommunityBallot) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *DomainCommunityBallot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainCommunityBallot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainCommunityBallot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainCommunityBallot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatformIdentityId

`func (o *DomainCommunityBallot) GetPlatformIdentityId() string`

GetPlatformIdentityId returns the PlatformIdentityId field if non-nil, zero value otherwise.

### GetPlatformIdentityIdOk

`func (o *DomainCommunityBallot) GetPlatformIdentityIdOk() (*string, bool)`

GetPlatformIdentityIdOk returns a tuple with the PlatformIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentityId

`func (o *DomainCommunityBallot) SetPlatformIdentityId(v string)`

SetPlatformIdentityId sets PlatformIdentityId field to given value.

### HasPlatformIdentityId

`func (o *DomainCommunityBallot) HasPlatformIdentityId() bool`

HasPlatformIdentityId returns a boolean if a field has been set.

### GetServedAt

`func (o *DomainCommunityBallot) GetServedAt() string`

GetServedAt returns the ServedAt field if non-nil, zero value otherwise.

### GetServedAtOk

`func (o *DomainCommunityBallot) GetServedAtOk() (*string, bool)`

GetServedAtOk returns a tuple with the ServedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAt

`func (o *DomainCommunityBallot) SetServedAt(v string)`

SetServedAt sets ServedAt field to given value.

### HasServedAt

`func (o *DomainCommunityBallot) HasServedAt() bool`

HasServedAt returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *DomainCommunityBallot) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DomainCommunityBallot) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DomainCommunityBallot) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DomainCommunityBallot) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetToken

`func (o *DomainCommunityBallot) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DomainCommunityBallot) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DomainCommunityBallot) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DomainCommunityBallot) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


