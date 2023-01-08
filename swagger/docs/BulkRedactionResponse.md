# BulkRedactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SingleRedactionResponse**](SingleRedactionResponse.md) | Result for requested redactions | 

## Methods

### NewBulkRedactionResponse

`func NewBulkRedactionResponse(results []SingleRedactionResponse, ) *BulkRedactionResponse`

NewBulkRedactionResponse instantiates a new BulkRedactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkRedactionResponseWithDefaults

`func NewBulkRedactionResponseWithDefaults() *BulkRedactionResponse`

NewBulkRedactionResponseWithDefaults instantiates a new BulkRedactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BulkRedactionResponse) GetResults() []SingleRedactionResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkRedactionResponse) GetResultsOk() (*[]SingleRedactionResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkRedactionResponse) SetResults(v []SingleRedactionResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


