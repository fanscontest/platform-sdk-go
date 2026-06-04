# DomainEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContestId** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | nullable, for live recording (seconds) | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** | \&quot;image\&quot; or \&quot;video\&quot; | [optional] 
**MediaUrl** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **[]int32** | for future use | [optional] 
**ParticipantId** | Pointer to **string** |  | [optional] 
**SubmittedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainEntry

`func NewDomainEntry() *DomainEntry`

NewDomainEntry instantiates a new DomainEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntryWithDefaults

`func NewDomainEntryWithDefaults() *DomainEntry`

NewDomainEntryWithDefaults instantiates a new DomainEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContestId

`func (o *DomainEntry) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainEntry) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainEntry) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainEntry) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetDuration

`func (o *DomainEntry) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DomainEntry) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DomainEntry) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DomainEntry) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *DomainEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaType

`func (o *DomainEntry) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *DomainEntry) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *DomainEntry) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *DomainEntry) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaUrl

`func (o *DomainEntry) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *DomainEntry) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *DomainEntry) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *DomainEntry) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *DomainEntry) GetMetadata() []int32`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DomainEntry) GetMetadataOk() (*[]int32, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DomainEntry) SetMetadata(v []int32)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DomainEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParticipantId

`func (o *DomainEntry) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *DomainEntry) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *DomainEntry) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *DomainEntry) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *DomainEntry) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DomainEntry) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DomainEntry) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DomainEntry) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


