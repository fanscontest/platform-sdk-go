# DomainMemberParticipationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanParticipate** | Pointer to **bool** |  | [optional] 
**CanParticipateReasons** | Pointer to **[]string** |  | [optional] 
**Enrolled** | Pointer to **bool** | deprecated: always true when struct is present | [optional] 
**ParticipationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | \&quot;participated\&quot; | \&quot;timed_out\&quot; | [optional] 
**Team** | Pointer to [**DomainTeam**](DomainTeam.md) |  | [optional] 

## Methods

### NewDomainMemberParticipationSummary

`func NewDomainMemberParticipationSummary() *DomainMemberParticipationSummary`

NewDomainMemberParticipationSummary instantiates a new DomainMemberParticipationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainMemberParticipationSummaryWithDefaults

`func NewDomainMemberParticipationSummaryWithDefaults() *DomainMemberParticipationSummary`

NewDomainMemberParticipationSummaryWithDefaults instantiates a new DomainMemberParticipationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanParticipate

`func (o *DomainMemberParticipationSummary) GetCanParticipate() bool`

GetCanParticipate returns the CanParticipate field if non-nil, zero value otherwise.

### GetCanParticipateOk

`func (o *DomainMemberParticipationSummary) GetCanParticipateOk() (*bool, bool)`

GetCanParticipateOk returns a tuple with the CanParticipate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanParticipate

`func (o *DomainMemberParticipationSummary) SetCanParticipate(v bool)`

SetCanParticipate sets CanParticipate field to given value.

### HasCanParticipate

`func (o *DomainMemberParticipationSummary) HasCanParticipate() bool`

HasCanParticipate returns a boolean if a field has been set.

### GetCanParticipateReasons

`func (o *DomainMemberParticipationSummary) GetCanParticipateReasons() []string`

GetCanParticipateReasons returns the CanParticipateReasons field if non-nil, zero value otherwise.

### GetCanParticipateReasonsOk

`func (o *DomainMemberParticipationSummary) GetCanParticipateReasonsOk() (*[]string, bool)`

GetCanParticipateReasonsOk returns a tuple with the CanParticipateReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanParticipateReasons

`func (o *DomainMemberParticipationSummary) SetCanParticipateReasons(v []string)`

SetCanParticipateReasons sets CanParticipateReasons field to given value.

### HasCanParticipateReasons

`func (o *DomainMemberParticipationSummary) HasCanParticipateReasons() bool`

HasCanParticipateReasons returns a boolean if a field has been set.

### GetEnrolled

`func (o *DomainMemberParticipationSummary) GetEnrolled() bool`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *DomainMemberParticipationSummary) GetEnrolledOk() (*bool, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *DomainMemberParticipationSummary) SetEnrolled(v bool)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *DomainMemberParticipationSummary) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetParticipationId

`func (o *DomainMemberParticipationSummary) GetParticipationId() string`

GetParticipationId returns the ParticipationId field if non-nil, zero value otherwise.

### GetParticipationIdOk

`func (o *DomainMemberParticipationSummary) GetParticipationIdOk() (*string, bool)`

GetParticipationIdOk returns a tuple with the ParticipationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationId

`func (o *DomainMemberParticipationSummary) SetParticipationId(v string)`

SetParticipationId sets ParticipationId field to given value.

### HasParticipationId

`func (o *DomainMemberParticipationSummary) HasParticipationId() bool`

HasParticipationId returns a boolean if a field has been set.

### GetStatus

`func (o *DomainMemberParticipationSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainMemberParticipationSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainMemberParticipationSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainMemberParticipationSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeam

`func (o *DomainMemberParticipationSummary) GetTeam() DomainTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *DomainMemberParticipationSummary) GetTeamOk() (*DomainTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *DomainMemberParticipationSummary) SetTeam(v DomainTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *DomainMemberParticipationSummary) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


