# DomainTournament

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**DomainChannel**](DomainChannel.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to [**DomainPlatformIdentity**](DomainPlatformIdentity.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**RevealGroupsLastNRounds** | Pointer to **int32** | RevealGroupsLastNRounds is how many trailing rounds get their groups pre-assigned from the advancers (revealed ahead of play). See ADR 0045. | [optional] 
**Rounds** | Pointer to [**[]DomainTournamentRound**](DomainTournamentRound.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainTournament

`func NewDomainTournament() *DomainTournament`

NewDomainTournament instantiates a new DomainTournament object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTournamentWithDefaults

`func NewDomainTournamentWithDefaults() *DomainTournament`

NewDomainTournamentWithDefaults instantiates a new DomainTournament object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DomainTournament) GetChannel() DomainChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DomainTournament) GetChannelOk() (*DomainChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DomainTournament) SetChannel(v DomainChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DomainTournament) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainTournament) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainTournament) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainTournament) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DomainTournament) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *DomainTournament) GetCreator() DomainPlatformIdentity`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *DomainTournament) GetCreatorOk() (*DomainPlatformIdentity, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *DomainTournament) SetCreator(v DomainPlatformIdentity)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *DomainTournament) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *DomainTournament) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainTournament) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainTournament) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainTournament) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DomainTournament) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainTournament) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainTournament) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainTournament) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRevealGroupsLastNRounds

`func (o *DomainTournament) GetRevealGroupsLastNRounds() int32`

GetRevealGroupsLastNRounds returns the RevealGroupsLastNRounds field if non-nil, zero value otherwise.

### GetRevealGroupsLastNRoundsOk

`func (o *DomainTournament) GetRevealGroupsLastNRoundsOk() (*int32, bool)`

GetRevealGroupsLastNRoundsOk returns a tuple with the RevealGroupsLastNRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevealGroupsLastNRounds

`func (o *DomainTournament) SetRevealGroupsLastNRounds(v int32)`

SetRevealGroupsLastNRounds sets RevealGroupsLastNRounds field to given value.

### HasRevealGroupsLastNRounds

`func (o *DomainTournament) HasRevealGroupsLastNRounds() bool`

HasRevealGroupsLastNRounds returns a boolean if a field has been set.

### GetRounds

`func (o *DomainTournament) GetRounds() []DomainTournamentRound`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *DomainTournament) GetRoundsOk() (*[]DomainTournamentRound, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *DomainTournament) SetRounds(v []DomainTournamentRound)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *DomainTournament) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetState

`func (o *DomainTournament) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DomainTournament) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DomainTournament) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DomainTournament) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *DomainTournament) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DomainTournament) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DomainTournament) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DomainTournament) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


