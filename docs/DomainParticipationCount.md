# DomainParticipationCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | Can be integer count or float sum of participation points | [optional] 
**Profile** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 

## Methods

### NewDomainParticipationCount

`func NewDomainParticipationCount() *DomainParticipationCount`

NewDomainParticipationCount instantiates a new DomainParticipationCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainParticipationCountWithDefaults

`func NewDomainParticipationCountWithDefaults() *DomainParticipationCount`

NewDomainParticipationCountWithDefaults instantiates a new DomainParticipationCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DomainParticipationCount) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DomainParticipationCount) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DomainParticipationCount) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DomainParticipationCount) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetProfile

`func (o *DomainParticipationCount) GetProfile() DomainPlatformIdentity`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DomainParticipationCount) GetProfileOk() (*DomainPlatformIdentity, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DomainParticipationCount) SetProfile(v DomainPlatformIdentity)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DomainParticipationCount) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


