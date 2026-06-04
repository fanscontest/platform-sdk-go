# DomainTournamentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCreateTournament** | Pointer to **bool** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**MaxRoundsPossible** | Pointer to **int32** |  | [optional] 
**MinParticipantsPerGroup** | Pointer to **int32** |  | [optional] 
**MinRoundsRequired** | Pointer to **int32** |  | [optional] 
**SubscriberCount** | Pointer to **int32** |  | [optional] 
**Templates** | Pointer to [**[]DomainTournamentTemplate**](DomainTournamentTemplate.md) |  | [optional] 

## Methods

### NewDomainTournamentConfiguration

`func NewDomainTournamentConfiguration() *DomainTournamentConfiguration`

NewDomainTournamentConfiguration instantiates a new DomainTournamentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTournamentConfigurationWithDefaults

`func NewDomainTournamentConfigurationWithDefaults() *DomainTournamentConfiguration`

NewDomainTournamentConfigurationWithDefaults instantiates a new DomainTournamentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCreateTournament

`func (o *DomainTournamentConfiguration) GetCanCreateTournament() bool`

GetCanCreateTournament returns the CanCreateTournament field if non-nil, zero value otherwise.

### GetCanCreateTournamentOk

`func (o *DomainTournamentConfiguration) GetCanCreateTournamentOk() (*bool, bool)`

GetCanCreateTournamentOk returns a tuple with the CanCreateTournament field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateTournament

`func (o *DomainTournamentConfiguration) SetCanCreateTournament(v bool)`

SetCanCreateTournament sets CanCreateTournament field to given value.

### HasCanCreateTournament

`func (o *DomainTournamentConfiguration) HasCanCreateTournament() bool`

HasCanCreateTournament returns a boolean if a field has been set.

### GetChannelId

`func (o *DomainTournamentConfiguration) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DomainTournamentConfiguration) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DomainTournamentConfiguration) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *DomainTournamentConfiguration) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetMaxRoundsPossible

`func (o *DomainTournamentConfiguration) GetMaxRoundsPossible() int32`

GetMaxRoundsPossible returns the MaxRoundsPossible field if non-nil, zero value otherwise.

### GetMaxRoundsPossibleOk

`func (o *DomainTournamentConfiguration) GetMaxRoundsPossibleOk() (*int32, bool)`

GetMaxRoundsPossibleOk returns a tuple with the MaxRoundsPossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRoundsPossible

`func (o *DomainTournamentConfiguration) SetMaxRoundsPossible(v int32)`

SetMaxRoundsPossible sets MaxRoundsPossible field to given value.

### HasMaxRoundsPossible

`func (o *DomainTournamentConfiguration) HasMaxRoundsPossible() bool`

HasMaxRoundsPossible returns a boolean if a field has been set.

### GetMinParticipantsPerGroup

`func (o *DomainTournamentConfiguration) GetMinParticipantsPerGroup() int32`

GetMinParticipantsPerGroup returns the MinParticipantsPerGroup field if non-nil, zero value otherwise.

### GetMinParticipantsPerGroupOk

`func (o *DomainTournamentConfiguration) GetMinParticipantsPerGroupOk() (*int32, bool)`

GetMinParticipantsPerGroupOk returns a tuple with the MinParticipantsPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinParticipantsPerGroup

`func (o *DomainTournamentConfiguration) SetMinParticipantsPerGroup(v int32)`

SetMinParticipantsPerGroup sets MinParticipantsPerGroup field to given value.

### HasMinParticipantsPerGroup

`func (o *DomainTournamentConfiguration) HasMinParticipantsPerGroup() bool`

HasMinParticipantsPerGroup returns a boolean if a field has been set.

### GetMinRoundsRequired

`func (o *DomainTournamentConfiguration) GetMinRoundsRequired() int32`

GetMinRoundsRequired returns the MinRoundsRequired field if non-nil, zero value otherwise.

### GetMinRoundsRequiredOk

`func (o *DomainTournamentConfiguration) GetMinRoundsRequiredOk() (*int32, bool)`

GetMinRoundsRequiredOk returns a tuple with the MinRoundsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRoundsRequired

`func (o *DomainTournamentConfiguration) SetMinRoundsRequired(v int32)`

SetMinRoundsRequired sets MinRoundsRequired field to given value.

### HasMinRoundsRequired

`func (o *DomainTournamentConfiguration) HasMinRoundsRequired() bool`

HasMinRoundsRequired returns a boolean if a field has been set.

### GetSubscriberCount

`func (o *DomainTournamentConfiguration) GetSubscriberCount() int32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *DomainTournamentConfiguration) GetSubscriberCountOk() (*int32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *DomainTournamentConfiguration) SetSubscriberCount(v int32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *DomainTournamentConfiguration) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### GetTemplates

`func (o *DomainTournamentConfiguration) GetTemplates() []DomainTournamentTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *DomainTournamentConfiguration) GetTemplatesOk() (*[]DomainTournamentTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *DomainTournamentConfiguration) SetTemplates(v []DomainTournamentTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *DomainTournamentConfiguration) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


