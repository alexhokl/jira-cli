# WorkflowPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]JiraWorkflowPreviewStatus**](JiraWorkflowPreviewStatus.md) | The list of statuses referenced by the workflows. | [optional] 
**Workflows** | Pointer to [**[]WorkflowPreview**](WorkflowPreview.md) | The list of workflows. The workflows are returned in the same order as specified in the request. | [optional] 

## Methods

### NewWorkflowPreviewResponse

`func NewWorkflowPreviewResponse() *WorkflowPreviewResponse`

NewWorkflowPreviewResponse instantiates a new WorkflowPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPreviewResponseWithDefaults

`func NewWorkflowPreviewResponseWithDefaults() *WorkflowPreviewResponse`

NewWorkflowPreviewResponseWithDefaults instantiates a new WorkflowPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowPreviewResponse) GetStatuses() []JiraWorkflowPreviewStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowPreviewResponse) GetStatusesOk() (*[]JiraWorkflowPreviewStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowPreviewResponse) SetStatuses(v []JiraWorkflowPreviewStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowPreviewResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowPreviewResponse) GetWorkflows() []WorkflowPreview`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowPreviewResponse) GetWorkflowsOk() (*[]WorkflowPreview, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowPreviewResponse) SetWorkflows(v []WorkflowPreview)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowPreviewResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


