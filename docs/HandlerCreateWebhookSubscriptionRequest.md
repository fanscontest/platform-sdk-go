# HandlerCreateWebhookSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypes** | **[]string** |  | 
**Url** | **string** |  | 

## Methods

### NewHandlerCreateWebhookSubscriptionRequest

`func NewHandlerCreateWebhookSubscriptionRequest(eventTypes []string, url string, ) *HandlerCreateWebhookSubscriptionRequest`

NewHandlerCreateWebhookSubscriptionRequest instantiates a new HandlerCreateWebhookSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateWebhookSubscriptionRequestWithDefaults

`func NewHandlerCreateWebhookSubscriptionRequestWithDefaults() *HandlerCreateWebhookSubscriptionRequest`

NewHandlerCreateWebhookSubscriptionRequestWithDefaults instantiates a new HandlerCreateWebhookSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypes

`func (o *HandlerCreateWebhookSubscriptionRequest) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *HandlerCreateWebhookSubscriptionRequest) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *HandlerCreateWebhookSubscriptionRequest) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetUrl

`func (o *HandlerCreateWebhookSubscriptionRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HandlerCreateWebhookSubscriptionRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HandlerCreateWebhookSubscriptionRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


