# PuzzleSheetsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PuzzleSheet**](PuzzleSheet.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewPuzzleSheetsListResponse

`func NewPuzzleSheetsListResponse() *PuzzleSheetsListResponse`

NewPuzzleSheetsListResponse instantiates a new PuzzleSheetsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleSheetsListResponseWithDefaults

`func NewPuzzleSheetsListResponseWithDefaults() *PuzzleSheetsListResponse`

NewPuzzleSheetsListResponseWithDefaults instantiates a new PuzzleSheetsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PuzzleSheetsListResponse) GetData() []PuzzleSheet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PuzzleSheetsListResponse) GetDataOk() (*[]PuzzleSheet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PuzzleSheetsListResponse) SetData(v []PuzzleSheet)`

SetData sets Data field to given value.

### HasData

`func (o *PuzzleSheetsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PuzzleSheetsListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PuzzleSheetsListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PuzzleSheetsListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PuzzleSheetsListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


