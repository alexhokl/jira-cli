# CreatedIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]BulkOperationErrorResult**](BulkOperationErrorResult.md) | Error details for failed issue creation requests. | [optional] [readonly] 
**Issues** | Pointer to [**[]CreatedIssue**](CreatedIssue.md) | Details of the issues created. | [optional] [readonly] 

## Methods

### NewCreatedIssues

`func NewCreatedIssues() *CreatedIssues`

NewCreatedIssues instantiates a new CreatedIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedIssuesWithDefaults

`func NewCreatedIssuesWithDefaults() *CreatedIssues`

NewCreatedIssuesWithDefaults instantiates a new CreatedIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CreatedIssues) GetErrors() []BulkOperationErrorResult`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreatedIssues) GetErrorsOk() (*[]BulkOperationErrorResult, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreatedIssues) SetErrors(v []BulkOperationErrorResult)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreatedIssues) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetIssues

`func (o *CreatedIssues) GetIssues() []CreatedIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *CreatedIssues) GetIssuesOk() (*[]CreatedIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *CreatedIssues) SetIssues(v []CreatedIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *CreatedIssues) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


