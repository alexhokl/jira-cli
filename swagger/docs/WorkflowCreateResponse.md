# WorkflowCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]JiraWorkflowStatus**](JiraWorkflowStatus.md) | List of created statuses. | [optional] 
**Workflows** | Pointer to [**[]JiraWorkflow**](JiraWorkflow.md) | List of created workflows. | [optional] 

## Methods

### NewWorkflowCreateResponse

`func NewWorkflowCreateResponse() *WorkflowCreateResponse`

NewWorkflowCreateResponse instantiates a new WorkflowCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCreateResponseWithDefaults

`func NewWorkflowCreateResponseWithDefaults() *WorkflowCreateResponse`

NewWorkflowCreateResponseWithDefaults instantiates a new WorkflowCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowCreateResponse) GetStatuses() []JiraWorkflowStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowCreateResponse) GetStatusesOk() (*[]JiraWorkflowStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowCreateResponse) SetStatuses(v []JiraWorkflowStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowCreateResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowCreateResponse) GetWorkflows() []JiraWorkflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowCreateResponse) GetWorkflowsOk() (*[]JiraWorkflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowCreateResponse) SetWorkflows(v []JiraWorkflow)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowCreateResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


