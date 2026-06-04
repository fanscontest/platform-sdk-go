# RequestCreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCode** | Pointer to **string** |  | [optional] 
**ChannelId** | **string** |  | 

## Methods

### NewRequestCreateSubscriptionRequest

`func NewRequestCreateSubscriptionRequest(channelId string, ) *RequestCreateSubscriptionRequest`

NewRequestCreateSubscriptionRequest instantiates a new RequestCreateSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSubscriptionRequestWithDefaults

`func NewRequestCreateSubscriptionRequestWithDefaults() *RequestCreateSubscriptionRequest`

NewRequestCreateSubscriptionRequestWithDefaults instantiates a new RequestCreateSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCode

`func (o *RequestCreateSubscriptionRequest) GetAccessCode() string`

GetAccessCode returns the AccessCode field if non-nil, zero value otherwise.

### GetAccessCodeOk

`func (o *RequestCreateSubscriptionRequest) GetAccessCodeOk() (*string, bool)`

GetAccessCodeOk returns a tuple with the AccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCode

`func (o *RequestCreateSubscriptionRequest) SetAccessCode(v string)`

SetAccessCode sets AccessCode field to given value.

### HasAccessCode

`func (o *RequestCreateSubscriptionRequest) HasAccessCode() bool`

HasAccessCode returns a boolean if a field has been set.

### GetChannelId

`func (o *RequestCreateSubscriptionRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *RequestCreateSubscriptionRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *RequestCreateSubscriptionRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


