# DomainSelectGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**ParticipationId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | \&quot;selected\&quot;, \&quot;already_selected_same\&quot;, \&quot;already_selected_other\&quot; | [optional] 

## Methods

### NewDomainSelectGroupResponse

`func NewDomainSelectGroupResponse() *DomainSelectGroupResponse`

NewDomainSelectGroupResponse instantiates a new DomainSelectGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSelectGroupResponseWithDefaults

`func NewDomainSelectGroupResponseWithDefaults() *DomainSelectGroupResponse`

NewDomainSelectGroupResponseWithDefaults instantiates a new DomainSelectGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DomainSelectGroupResponse) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DomainSelectGroupResponse) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DomainSelectGroupResponse) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DomainSelectGroupResponse) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetParticipationId

`func (o *DomainSelectGroupResponse) GetParticipationId() string`

GetParticipationId returns the ParticipationId field if non-nil, zero value otherwise.

### GetParticipationIdOk

`func (o *DomainSelectGroupResponse) GetParticipationIdOk() (*string, bool)`

GetParticipationIdOk returns a tuple with the ParticipationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationId

`func (o *DomainSelectGroupResponse) SetParticipationId(v string)`

SetParticipationId sets ParticipationId field to given value.

### HasParticipationId

`func (o *DomainSelectGroupResponse) HasParticipationId() bool`

HasParticipationId returns a boolean if a field has been set.

### GetState

`func (o *DomainSelectGroupResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainSelectGroupResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainSelectGroupResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainSelectGroupResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


