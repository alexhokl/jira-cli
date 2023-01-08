# WorkflowValidationErrorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]WorkflowValidationError**](WorkflowValidationError.md) | The list of validation errors. | [optional] 

## Methods

### NewWorkflowValidationErrorList

`func NewWorkflowValidationErrorList() *WorkflowValidationErrorList`

NewWorkflowValidationErrorList instantiates a new WorkflowValidationErrorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowValidationErrorListWithDefaults

`func NewWorkflowValidationErrorListWithDefaults() *WorkflowValidationErrorList`

NewWorkflowValidationErrorListWithDefaults instantiates a new WorkflowValidationErrorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *WorkflowValidationErrorList) GetErrors() []WorkflowValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WorkflowValidationErrorList) GetErrorsOk() (*[]WorkflowValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WorkflowValidationErrorList) SetErrors(v []WorkflowValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WorkflowValidationErrorList) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


