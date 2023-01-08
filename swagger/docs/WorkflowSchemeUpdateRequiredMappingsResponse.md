# WorkflowSchemeUpdateRequiredMappingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusMappingsByIssueTypes** | Pointer to [**[]RequiredMappingByIssueType**](RequiredMappingByIssueType.md) | The list of required status mappings by issue type. | [optional] 
**StatusMappingsByWorkflows** | Pointer to [**[]RequiredMappingByWorkflows**](RequiredMappingByWorkflows.md) | The list of required status mappings by workflow. | [optional] 
**Statuses** | Pointer to [**[]StatusMetadata**](StatusMetadata.md) | The details of the statuses in the associated workflows. | [optional] 
**StatusesPerWorkflow** | Pointer to [**[]StatusesPerWorkflow**](StatusesPerWorkflow.md) | The statuses associated with each workflow. | [optional] 

## Methods

### NewWorkflowSchemeUpdateRequiredMappingsResponse

`func NewWorkflowSchemeUpdateRequiredMappingsResponse() *WorkflowSchemeUpdateRequiredMappingsResponse`

NewWorkflowSchemeUpdateRequiredMappingsResponse instantiates a new WorkflowSchemeUpdateRequiredMappingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeUpdateRequiredMappingsResponseWithDefaults

`func NewWorkflowSchemeUpdateRequiredMappingsResponseWithDefaults() *WorkflowSchemeUpdateRequiredMappingsResponse`

NewWorkflowSchemeUpdateRequiredMappingsResponseWithDefaults instantiates a new WorkflowSchemeUpdateRequiredMappingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusMappingsByIssueTypes

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusMappingsByIssueTypes() []RequiredMappingByIssueType`

GetStatusMappingsByIssueTypes returns the StatusMappingsByIssueTypes field if non-nil, zero value otherwise.

### GetStatusMappingsByIssueTypesOk

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusMappingsByIssueTypesOk() (*[]RequiredMappingByIssueType, bool)`

GetStatusMappingsByIssueTypesOk returns a tuple with the StatusMappingsByIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappingsByIssueTypes

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) SetStatusMappingsByIssueTypes(v []RequiredMappingByIssueType)`

SetStatusMappingsByIssueTypes sets StatusMappingsByIssueTypes field to given value.

### HasStatusMappingsByIssueTypes

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) HasStatusMappingsByIssueTypes() bool`

HasStatusMappingsByIssueTypes returns a boolean if a field has been set.

### GetStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusMappingsByWorkflows() []RequiredMappingByWorkflows`

GetStatusMappingsByWorkflows returns the StatusMappingsByWorkflows field if non-nil, zero value otherwise.

### GetStatusMappingsByWorkflowsOk

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusMappingsByWorkflowsOk() (*[]RequiredMappingByWorkflows, bool)`

GetStatusMappingsByWorkflowsOk returns a tuple with the StatusMappingsByWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) SetStatusMappingsByWorkflows(v []RequiredMappingByWorkflows)`

SetStatusMappingsByWorkflows sets StatusMappingsByWorkflows field to given value.

### HasStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) HasStatusMappingsByWorkflows() bool`

HasStatusMappingsByWorkflows returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatuses() []StatusMetadata`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusesOk() (*[]StatusMetadata, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) SetStatuses(v []StatusMetadata)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetStatusesPerWorkflow

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusesPerWorkflow() []StatusesPerWorkflow`

GetStatusesPerWorkflow returns the StatusesPerWorkflow field if non-nil, zero value otherwise.

### GetStatusesPerWorkflowOk

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) GetStatusesPerWorkflowOk() (*[]StatusesPerWorkflow, bool)`

GetStatusesPerWorkflowOk returns a tuple with the StatusesPerWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusesPerWorkflow

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) SetStatusesPerWorkflow(v []StatusesPerWorkflow)`

SetStatusesPerWorkflow sets StatusesPerWorkflow field to given value.

### HasStatusesPerWorkflow

`func (o *WorkflowSchemeUpdateRequiredMappingsResponse) HasStatusesPerWorkflow() bool`

HasStatusesPerWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


