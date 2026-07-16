# RequestSubmitJudgmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BallotToken** | **string** |  | 
**Choice** | **string** |  | 

## Methods

### NewRequestSubmitJudgmentRequest

`func NewRequestSubmitJudgmentRequest(ballotToken string, choice string, ) *RequestSubmitJudgmentRequest`

NewRequestSubmitJudgmentRequest instantiates a new RequestSubmitJudgmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSubmitJudgmentRequestWithDefaults

`func NewRequestSubmitJudgmentRequestWithDefaults() *RequestSubmitJudgmentRequest`

NewRequestSubmitJudgmentRequestWithDefaults instantiates a new RequestSubmitJudgmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBallotToken

`func (o *RequestSubmitJudgmentRequest) GetBallotToken() string`

GetBallotToken returns the BallotToken field if non-nil, zero value otherwise.

### GetBallotTokenOk

`func (o *RequestSubmitJudgmentRequest) GetBallotTokenOk() (*string, bool)`

GetBallotTokenOk returns a tuple with the BallotToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBallotToken

`func (o *RequestSubmitJudgmentRequest) SetBallotToken(v string)`

SetBallotToken sets BallotToken field to given value.


### GetChoice

`func (o *RequestSubmitJudgmentRequest) GetChoice() string`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *RequestSubmitJudgmentRequest) GetChoiceOk() (*string, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *RequestSubmitJudgmentRequest) SetChoice(v string)`

SetChoice sets Choice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


