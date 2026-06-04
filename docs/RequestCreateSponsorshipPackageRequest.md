# RequestCreateSponsorshipPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceDeadline** | Pointer to **string** |  | [optional] 
**AllowedContestTopics** | Pointer to **[]string** |  | [optional] 
**BrandingTerms** | Pointer to **string** |  | [optional] 
**BudgetPerContestCurrency** | **string** |  | 
**BudgetPerContestValue** | **string** | Package Agreement fields (auto-created) | 
**MaxContests** | Pointer to **int32** |  | [optional] 
**TermIds** | **[]string** | IDs of terms to include | 
**TotalBudgetCurrency** | **string** |  | 
**TotalBudgetValue** | **string** |  | 
**TotalContests** | **int32** |  | 

## Methods

### NewRequestCreateSponsorshipPackageRequest

`func NewRequestCreateSponsorshipPackageRequest(budgetPerContestCurrency string, budgetPerContestValue string, termIds []string, totalBudgetCurrency string, totalBudgetValue string, totalContests int32, ) *RequestCreateSponsorshipPackageRequest`

NewRequestCreateSponsorshipPackageRequest instantiates a new RequestCreateSponsorshipPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSponsorshipPackageRequestWithDefaults

`func NewRequestCreateSponsorshipPackageRequestWithDefaults() *RequestCreateSponsorshipPackageRequest`

NewRequestCreateSponsorshipPackageRequestWithDefaults instantiates a new RequestCreateSponsorshipPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceDeadline

`func (o *RequestCreateSponsorshipPackageRequest) GetAcceptanceDeadline() string`

GetAcceptanceDeadline returns the AcceptanceDeadline field if non-nil, zero value otherwise.

### GetAcceptanceDeadlineOk

`func (o *RequestCreateSponsorshipPackageRequest) GetAcceptanceDeadlineOk() (*string, bool)`

GetAcceptanceDeadlineOk returns a tuple with the AcceptanceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDeadline

`func (o *RequestCreateSponsorshipPackageRequest) SetAcceptanceDeadline(v string)`

SetAcceptanceDeadline sets AcceptanceDeadline field to given value.

### HasAcceptanceDeadline

`func (o *RequestCreateSponsorshipPackageRequest) HasAcceptanceDeadline() bool`

HasAcceptanceDeadline returns a boolean if a field has been set.

### GetAllowedContestTopics

`func (o *RequestCreateSponsorshipPackageRequest) GetAllowedContestTopics() []string`

GetAllowedContestTopics returns the AllowedContestTopics field if non-nil, zero value otherwise.

### GetAllowedContestTopicsOk

`func (o *RequestCreateSponsorshipPackageRequest) GetAllowedContestTopicsOk() (*[]string, bool)`

GetAllowedContestTopicsOk returns a tuple with the AllowedContestTopics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContestTopics

`func (o *RequestCreateSponsorshipPackageRequest) SetAllowedContestTopics(v []string)`

SetAllowedContestTopics sets AllowedContestTopics field to given value.

### HasAllowedContestTopics

`func (o *RequestCreateSponsorshipPackageRequest) HasAllowedContestTopics() bool`

HasAllowedContestTopics returns a boolean if a field has been set.

### GetBrandingTerms

`func (o *RequestCreateSponsorshipPackageRequest) GetBrandingTerms() string`

GetBrandingTerms returns the BrandingTerms field if non-nil, zero value otherwise.

### GetBrandingTermsOk

`func (o *RequestCreateSponsorshipPackageRequest) GetBrandingTermsOk() (*string, bool)`

GetBrandingTermsOk returns a tuple with the BrandingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandingTerms

`func (o *RequestCreateSponsorshipPackageRequest) SetBrandingTerms(v string)`

SetBrandingTerms sets BrandingTerms field to given value.

### HasBrandingTerms

`func (o *RequestCreateSponsorshipPackageRequest) HasBrandingTerms() bool`

HasBrandingTerms returns a boolean if a field has been set.

### GetBudgetPerContestCurrency

`func (o *RequestCreateSponsorshipPackageRequest) GetBudgetPerContestCurrency() string`

GetBudgetPerContestCurrency returns the BudgetPerContestCurrency field if non-nil, zero value otherwise.

### GetBudgetPerContestCurrencyOk

`func (o *RequestCreateSponsorshipPackageRequest) GetBudgetPerContestCurrencyOk() (*string, bool)`

GetBudgetPerContestCurrencyOk returns a tuple with the BudgetPerContestCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetPerContestCurrency

`func (o *RequestCreateSponsorshipPackageRequest) SetBudgetPerContestCurrency(v string)`

SetBudgetPerContestCurrency sets BudgetPerContestCurrency field to given value.


### GetBudgetPerContestValue

`func (o *RequestCreateSponsorshipPackageRequest) GetBudgetPerContestValue() string`

GetBudgetPerContestValue returns the BudgetPerContestValue field if non-nil, zero value otherwise.

### GetBudgetPerContestValueOk

`func (o *RequestCreateSponsorshipPackageRequest) GetBudgetPerContestValueOk() (*string, bool)`

GetBudgetPerContestValueOk returns a tuple with the BudgetPerContestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetPerContestValue

`func (o *RequestCreateSponsorshipPackageRequest) SetBudgetPerContestValue(v string)`

SetBudgetPerContestValue sets BudgetPerContestValue field to given value.


### GetMaxContests

`func (o *RequestCreateSponsorshipPackageRequest) GetMaxContests() int32`

GetMaxContests returns the MaxContests field if non-nil, zero value otherwise.

### GetMaxContestsOk

`func (o *RequestCreateSponsorshipPackageRequest) GetMaxContestsOk() (*int32, bool)`

GetMaxContestsOk returns a tuple with the MaxContests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContests

`func (o *RequestCreateSponsorshipPackageRequest) SetMaxContests(v int32)`

SetMaxContests sets MaxContests field to given value.

### HasMaxContests

`func (o *RequestCreateSponsorshipPackageRequest) HasMaxContests() bool`

HasMaxContests returns a boolean if a field has been set.

### GetTermIds

`func (o *RequestCreateSponsorshipPackageRequest) GetTermIds() []string`

GetTermIds returns the TermIds field if non-nil, zero value otherwise.

### GetTermIdsOk

`func (o *RequestCreateSponsorshipPackageRequest) GetTermIdsOk() (*[]string, bool)`

GetTermIdsOk returns a tuple with the TermIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermIds

`func (o *RequestCreateSponsorshipPackageRequest) SetTermIds(v []string)`

SetTermIds sets TermIds field to given value.


### GetTotalBudgetCurrency

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalBudgetCurrency() string`

GetTotalBudgetCurrency returns the TotalBudgetCurrency field if non-nil, zero value otherwise.

### GetTotalBudgetCurrencyOk

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalBudgetCurrencyOk() (*string, bool)`

GetTotalBudgetCurrencyOk returns a tuple with the TotalBudgetCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBudgetCurrency

`func (o *RequestCreateSponsorshipPackageRequest) SetTotalBudgetCurrency(v string)`

SetTotalBudgetCurrency sets TotalBudgetCurrency field to given value.


### GetTotalBudgetValue

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalBudgetValue() string`

GetTotalBudgetValue returns the TotalBudgetValue field if non-nil, zero value otherwise.

### GetTotalBudgetValueOk

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalBudgetValueOk() (*string, bool)`

GetTotalBudgetValueOk returns a tuple with the TotalBudgetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBudgetValue

`func (o *RequestCreateSponsorshipPackageRequest) SetTotalBudgetValue(v string)`

SetTotalBudgetValue sets TotalBudgetValue field to given value.


### GetTotalContests

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalContests() int32`

GetTotalContests returns the TotalContests field if non-nil, zero value otherwise.

### GetTotalContestsOk

`func (o *RequestCreateSponsorshipPackageRequest) GetTotalContestsOk() (*int32, bool)`

GetTotalContestsOk returns a tuple with the TotalContests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContests

`func (o *RequestCreateSponsorshipPackageRequest) SetTotalContests(v int32)`

SetTotalContests sets TotalContests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


