# RequestCreateSponsorshipOfferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCurrency** | **string** |  | 
**AmountValue** | **string** |  | 
**BrandingTerms** | Pointer to **string** |  | [optional] 
**ContestId** | **string** |  | 
**Deadline** | Pointer to **string** |  | [optional] 
**TermIds** | **[]string** | IDs of terms to include | 

## Methods

### NewRequestCreateSponsorshipOfferRequest

`func NewRequestCreateSponsorshipOfferRequest(amountCurrency string, amountValue string, contestId string, termIds []string, ) *RequestCreateSponsorshipOfferRequest`

NewRequestCreateSponsorshipOfferRequest instantiates a new RequestCreateSponsorshipOfferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSponsorshipOfferRequestWithDefaults

`func NewRequestCreateSponsorshipOfferRequestWithDefaults() *RequestCreateSponsorshipOfferRequest`

NewRequestCreateSponsorshipOfferRequestWithDefaults instantiates a new RequestCreateSponsorshipOfferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCurrency

`func (o *RequestCreateSponsorshipOfferRequest) GetAmountCurrency() string`

GetAmountCurrency returns the AmountCurrency field if non-nil, zero value otherwise.

### GetAmountCurrencyOk

`func (o *RequestCreateSponsorshipOfferRequest) GetAmountCurrencyOk() (*string, bool)`

GetAmountCurrencyOk returns a tuple with the AmountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCurrency

`func (o *RequestCreateSponsorshipOfferRequest) SetAmountCurrency(v string)`

SetAmountCurrency sets AmountCurrency field to given value.


### GetAmountValue

`func (o *RequestCreateSponsorshipOfferRequest) GetAmountValue() string`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *RequestCreateSponsorshipOfferRequest) GetAmountValueOk() (*string, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *RequestCreateSponsorshipOfferRequest) SetAmountValue(v string)`

SetAmountValue sets AmountValue field to given value.


### GetBrandingTerms

`func (o *RequestCreateSponsorshipOfferRequest) GetBrandingTerms() string`

GetBrandingTerms returns the BrandingTerms field if non-nil, zero value otherwise.

### GetBrandingTermsOk

`func (o *RequestCreateSponsorshipOfferRequest) GetBrandingTermsOk() (*string, bool)`

GetBrandingTermsOk returns a tuple with the BrandingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTerms

`func (o *RequestCreateSponsorshipOfferRequest) SetBrandingTerms(v string)`

SetBrandingTerms sets BrandingTerms field to given value.

### HasBrandingTerms

`func (o *RequestCreateSponsorshipOfferRequest) HasBrandingTerms() bool`

HasBrandingTerms returns a boolean if a field has been set.

### GetContestId

`func (o *RequestCreateSponsorshipOfferRequest) GetContestId() string`

GetContestId returns the ContestId field if non-nil, zero value otherwise.

### GetContestIdOk

`func (o *RequestCreateSponsorshipOfferRequest) GetContestIdOk() (*string, bool)`

GetContestIdOk returns a tuple with the ContestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContestId

`func (o *RequestCreateSponsorshipOfferRequest) SetContestId(v string)`

SetContestId sets ContestId field to given value.


### GetDeadline

`func (o *RequestCreateSponsorshipOfferRequest) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *RequestCreateSponsorshipOfferRequest) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *RequestCreateSponsorshipOfferRequest) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *RequestCreateSponsorshipOfferRequest) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetTermIds

`func (o *RequestCreateSponsorshipOfferRequest) GetTermIds() []string`

GetTermIds returns the TermIds field if non-nil, zero value otherwise.

### GetTermIdsOk

`func (o *RequestCreateSponsorshipOfferRequest) GetTermIdsOk() (*[]string, bool)`

GetTermIdsOk returns a tuple with the TermIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermIds

`func (o *RequestCreateSponsorshipOfferRequest) SetTermIds(v []string)`

SetTermIds sets TermIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


