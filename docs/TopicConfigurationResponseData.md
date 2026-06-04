# TopicConfigurationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DifficultySettings** | Pointer to [**[]TopicConfigurationResponseDataDifficultySettingsInner**](TopicConfigurationResponseDataDifficultySettingsInner.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 

## Methods

### NewTopicConfigurationResponseData

`func NewTopicConfigurationResponseData() *TopicConfigurationResponseData`

NewTopicConfigurationResponseData instantiates a new TopicConfigurationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicConfigurationResponseDataWithDefaults

`func NewTopicConfigurationResponseDataWithDefaults() *TopicConfigurationResponseData`

NewTopicConfigurationResponseDataWithDefaults instantiates a new TopicConfigurationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifficultySettings

`func (o *TopicConfigurationResponseData) GetDifficultySettings() []TopicConfigurationResponseDataDifficultySettingsInner`

GetDifficultySettings returns the DifficultySettings field if non-nil, zero value otherwise.

### GetDifficultySettingsOk

`func (o *TopicConfigurationResponseData) GetDifficultySettingsOk() (*[]TopicConfigurationResponseDataDifficultySettingsInner, bool)`

GetDifficultySettingsOk returns a tuple with the DifficultySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficultySettings

`func (o *TopicConfigurationResponseData) SetDifficultySettings(v []TopicConfigurationResponseDataDifficultySettingsInner)`

SetDifficultySettings sets DifficultySettings field to given value.

### HasDifficultySettings

`func (o *TopicConfigurationResponseData) HasDifficultySettings() bool`

HasDifficultySettings returns a boolean if a field has been set.

### GetTopic

`func (o *TopicConfigurationResponseData) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *TopicConfigurationResponseData) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *TopicConfigurationResponseData) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *TopicConfigurationResponseData) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


