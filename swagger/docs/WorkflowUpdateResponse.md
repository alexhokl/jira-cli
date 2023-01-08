# WorkflowUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]JiraWorkflowStatus**](JiraWorkflowStatus.md) | List of updated statuses. | [optional] 
**TaskId** | Pointer to **NullableString** | If there is a [asynchronous task](#async-operations) operation, as a result of this update. | [optional] 
**Workflows** | Pointer to [**[]JiraWorkflow**](JiraWorkflow.md) | List of updated workflows. | [optional] 

## Methods

### NewWorkflowUpdateResponse

`func NewWorkflowUpdateResponse() *WorkflowUpdateResponse`

NewWorkflowUpdateResponse instantiates a new WorkflowUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUpdateResponseWithDefaults

`func NewWorkflowUpdateResponseWithDefaults() *WorkflowUpdateResponse`

NewWorkflowUpdateResponseWithDefaults instantiates a new WorkflowUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowUpdateResponse) GetStatuses() []JiraWorkflowStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowUpdateResponse) GetStatusesOk() (*[]JiraWorkflowStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowUpdateResponse) SetStatuses(v []JiraWorkflowStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowUpdateResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTaskId

`func (o *WorkflowUpdateResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowUpdateResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowUpdateResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *WorkflowUpdateResponse) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *WorkflowUpdateResponse) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *WorkflowUpdateResponse) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetWorkflows

`func (o *WorkflowUpdateResponse) GetWorkflows() []JiraWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowUpdateResponse) GetWorkflowsOk() (*[]JiraWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowUpdateResponse) SetWorkflows(v []JiraWorkflow)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowUpdateResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


