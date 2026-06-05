# DomainJurorAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedAt** | Pointer to **string** |  | [optional] 
**ContestId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | keyset pagination sort key (uman#132) | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JurorId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | Use JurorAssignmentSourceRandom or JurorAssignmentSourceInvited | [optional] 
**Status** | Pointer to **string** | Use JurorAssignmentStatusPending, Confirmed, Declined, or Expired | [optional] 

## Methods

### NewDomainJurorAssignment

`func NewDomainJurorAssignment() *DomainJurorAssignment`

NewDomainJurorAssignment instantiates a new DomainJurorAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainJurorAssignmentWithDefaults

`func NewDomainJurorAssignmentWithDefaults() *DomainJurorAssignment`

NewDomainJurorAssignmentWithDefaults instantiates a new DomainJurorAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedAt

`func (o *DomainJurorAssignment) GetAssignedAt() string`

GetAssignedAt returns the AssignedAt field if non-nil, zero value otherwise.

### GetAssignedAtOk

`func (o *DomainJurorAssignment) GetAssignedAtOk() (*string, bool)`

GetAssignedAtOk returns a tuple with the AssignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAt

`func (o *DomainJurorAssignment) SetAssignedAt(v string)`

SetAssignedAt sets AssignedAt field to given value.

### HasAssignedAt

`func (o *DomainJurorAssignment) HasAssignedAt() bool`

HasAssignedAt returns a boolean if a field has been set.

### GetContestId

`func (o *DomainJurorAssignment) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *DomainJurorAssignment) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *DomainJurorAssignment) SetContestId(v string)`

SetContestId sets ContestId field to given value.

### HasContestId

`func (o *DomainJurorAssignment) HasContestId() bool`

HasContestId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainJurorAssignment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainJurorAssignment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainJurorAssignment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainJurorAssignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *DomainJurorAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainJurorAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainJurorAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainJurorAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJurorId

`func (o *DomainJurorAssignment) GetJurorId() string`

GetJurorId returns the JurorId field if non-nil, zero value otherwise.

### GetJurorIdOk

`func (o *DomainJurorAssignment) GetJurorIdOk() (*string, bool)`

GetJurorIdOk returns a tuple with the JurorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurorId

`func (o *DomainJurorAssignment) SetJurorId(v string)`

SetJurorId sets JurorId field to given value.

### HasJurorId

`func (o *DomainJurorAssignment) HasJurorId() bool`

HasJurorId returns a boolean if a field has been set.

### GetSource

`func (o *DomainJurorAssignment) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DomainJurorAssignment) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DomainJurorAssignment) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DomainJurorAssignment) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *DomainJurorAssignment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainJurorAssignment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainJurorAssignment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainJurorAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


