# RequestCreateTeamSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Teams** | [**[]RequestCreateTeamRequest**](RequestCreateTeamRequest.md) |  | 

## Methods

### NewRequestCreateTeamSetRequest

`func NewRequestCreateTeamSetRequest(name string, teams []RequestCreateTeamRequest, ) *RequestCreateTeamSetRequest`

NewRequestCreateTeamSetRequest instantiates a new RequestCreateTeamSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateTeamSetRequestWithDefaults

`func NewRequestCreateTeamSetRequestWithDefaults() *RequestCreateTeamSetRequest`

NewRequestCreateTeamSetRequestWithDefaults instantiates a new RequestCreateTeamSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *RequestCreateTeamSetRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RequestCreateTeamSetRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RequestCreateTeamSetRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RequestCreateTeamSetRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *RequestCreateTeamSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestCreateTeamSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestCreateTeamSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *RequestCreateTeamSetRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequestCreateTeamSetRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequestCreateTeamSetRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequestCreateTeamSetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeams

`func (o *RequestCreateTeamSetRequest) GetTeams() []RequestCreateTeamRequest`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *RequestCreateTeamSetRequest) GetTeamsOk() (*[]RequestCreateTeamRequest, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *RequestCreateTeamSetRequest) SetTeams(v []RequestCreateTeamRequest)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


