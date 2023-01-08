# WorkflowUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]WorkflowStatusUpdate**](WorkflowStatusUpdate.md) | The statuses to associate with the workflows. | [optional] 
**Workflows** | Pointer to [**[]WorkflowUpdate**](WorkflowUpdate.md) | The details of the workflows to update. | [optional] 

## Methods

### NewWorkflowUpdateRequest

`func NewWorkflowUpdateRequest() *WorkflowUpdateRequest`

NewWorkflowUpdateRequest instantiates a new WorkflowUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUpdateRequestWithDefaults

`func NewWorkflowUpdateRequestWithDefaults() *WorkflowUpdateRequest`

NewWorkflowUpdateRequestWithDefaults instantiates a new WorkflowUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowUpdateRequest) GetStatuses() []WorkflowStatusUpdate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowUpdateRequest) GetStatusesOk() (*[]WorkflowStatusUpdate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowUpdateRequest) SetStatuses(v []WorkflowStatusUpdate)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowUpdateRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowUpdateRequest) GetWorkflows() []WorkflowUpdate`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowUpdateRequest) GetWorkflowsOk() (*[]WorkflowUpdate, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowUpdateRequest) SetWorkflows(v []WorkflowUpdate)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowUpdateRequest) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


