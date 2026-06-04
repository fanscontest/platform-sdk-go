# RequestCreateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewRequestCreateTeamRequest

`func NewRequestCreateTeamRequest(name string, ) *RequestCreateTeamRequest`

NewRequestCreateTeamRequest instantiates a new RequestCreateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateTeamRequestWithDefaults

`func NewRequestCreateTeamRequestWithDefaults() *RequestCreateTeamRequest`

NewRequestCreateTeamRequestWithDefaults instantiates a new RequestCreateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *RequestCreateTeamRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *RequestCreateTeamRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *RequestCreateTeamRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *RequestCreateTeamRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetName

`func (o *RequestCreateTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestCreateTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestCreateTeamRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


