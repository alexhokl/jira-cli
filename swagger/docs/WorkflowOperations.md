# WorkflowOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDelete** | **bool** | Whether the workflow can be deleted. | 
**CanEdit** | **bool** | Whether the workflow can be updated. | 

## Methods

### NewWorkflowOperations

`func NewWorkflowOperations(canDelete bool, canEdit bool, ) *WorkflowOperations`

NewWorkflowOperations instantiates a new WorkflowOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowOperationsWithDefaults

`func NewWorkflowOperationsWithDefaults() *WorkflowOperations`

NewWorkflowOperationsWithDefaults instantiates a new WorkflowOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDelete

`func (o *WorkflowOperations) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *WorkflowOperations) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *WorkflowOperations) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.


### GetCanEdit

`func (o *WorkflowOperations) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *WorkflowOperations) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *WorkflowOperations) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


