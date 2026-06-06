# DomainWebhookSubscriptionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DomainWebhookSubscription**](DomainWebhookSubscription.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewDomainWebhookSubscriptionListResponse

`func NewDomainWebhookSubscriptionListResponse(data []DomainWebhookSubscription, ) *DomainWebhookSubscriptionListResponse`

NewDomainWebhookSubscriptionListResponse instantiates a new DomainWebhookSubscriptionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWebhookSubscriptionListResponseWithDefaults

`func NewDomainWebhookSubscriptionListResponseWithDefaults() *DomainWebhookSubscriptionListResponse`

NewDomainWebhookSubscriptionListResponseWithDefaults instantiates a new DomainWebhookSubscriptionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainWebhookSubscriptionListResponse) GetData() []DomainWebhookSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainWebhookSubscriptionListResponse) GetDataOk() (*[]DomainWebhookSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainWebhookSubscriptionListResponse) SetData(v []DomainWebhookSubscription)`

SetData sets Data field to given value.


### GetPagination

`func (o *DomainWebhookSubscriptionListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DomainWebhookSubscriptionListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DomainWebhookSubscriptionListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DomainWebhookSubscriptionListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


