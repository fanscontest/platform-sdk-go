# DomainSubmissionFeedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestId** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**SubmissionNumber** | Pointer to **int32** |  | [optional] 
**SubmittedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainSubmissionFeedItem

`func NewDomainSubmissionFeedItem() *DomainSubmissionFeedItem`

NewDomainSubmissionFeedItem instantiates a new DomainSubmissionFeedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSubmissionFeedItemWithDefaults

`func NewDomainSubmissionFeedItemWithDefaults() *DomainSubmissionFeedItem`

NewDomainSubmissionFeedItemWithDefaults instantiates a new DomainSubmissionFeedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestId

`func (o *DomainSubmissionFeedItem) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainSubmissionFeedItem) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainSubmissionFeedItem) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainSubmissionFeedItem) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetGroup

`func (o *DomainSubmissionFeedItem) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DomainSubmissionFeedItem) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DomainSubmissionFeedItem) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DomainSubmissionFeedItem) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMember

`func (o *DomainSubmissionFeedItem) GetMember() DomainPlatformIdentity`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *DomainSubmissionFeedItem) GetMemberOk() (*DomainPlatformIdentity, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *DomainSubmissionFeedItem) SetMember(v DomainPlatformIdentity)`

SetMember sets Member field to given value.

### HasMember

`func (o *DomainSubmissionFeedItem) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetSubmissionNumber

`func (o *DomainSubmissionFeedItem) GetSubmissionNumber() int32`

GetSubmissionNumber returns the SubmissionNumber field if non-nil, zero value otherwise.

### GetSubmissionNumberOk

`func (o *DomainSubmissionFeedItem) GetSubmissionNumberOk() (*int32, bool)`

GetSubmissionNumberOk returns a tuple with the SubmissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionNumber

`func (o *DomainSubmissionFeedItem) SetSubmissionNumber(v int32)`

SetSubmissionNumber sets SubmissionNumber field to given value.

### HasSubmissionNumber

`func (o *DomainSubmissionFeedItem) HasSubmissionNumber() bool`

HasSubmissionNumber returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *DomainSubmissionFeedItem) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DomainSubmissionFeedItem) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DomainSubmissionFeedItem) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DomainSubmissionFeedItem) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


