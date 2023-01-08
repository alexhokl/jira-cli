# WorkflowCapabilityPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]StatusPayload**](StatusPayload.md) | The statuses for the workflow | [optional] 
**WorkflowScheme** | Pointer to [**WorkflowSchemePayload**](WorkflowSchemePayload.md) |  | [optional] 
**Workflows** | Pointer to [**[]WorkflowPayload**](WorkflowPayload.md) | The transitions for the workflow | [optional] 

## Methods

### NewWorkflowCapabilityPayload

`func NewWorkflowCapabilityPayload() *WorkflowCapabilityPayload`

NewWorkflowCapabilityPayload instantiates a new WorkflowCapabilityPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCapabilityPayloadWithDefaults

`func NewWorkflowCapabilityPayloadWithDefaults() *WorkflowCapabilityPayload`

NewWorkflowCapabilityPayloadWithDefaults instantiates a new WorkflowCapabilityPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowCapabilityPayload) GetStatuses() []StatusPayload`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowCapabilityPayload) GetStatusesOk() (*[]StatusPayload, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowCapabilityPayload) SetStatuses(v []StatusPayload)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowCapabilityPayload) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflowScheme

`func (o *WorkflowCapabilityPayload) GetWorkflowScheme() WorkflowSchemePayload`

GetWorkflowScheme returns the WorkflowScheme field if non-nil, zero value otherwise.

### GetWorkflowSchemeOk

`func (o *WorkflowCapabilityPayload) GetWorkflowSchemeOk() (*WorkflowSchemePayload, bool)`

GetWorkflowSchemeOk returns a tuple with the WorkflowScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowScheme

`func (o *WorkflowCapabilityPayload) SetWorkflowScheme(v WorkflowSchemePayload)`

SetWorkflowScheme sets WorkflowScheme field to given value.

### HasWorkflowScheme

`func (o *WorkflowCapabilityPayload) HasWorkflowScheme() bool`

HasWorkflowScheme returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowCapabilityPayload) GetWorkflows() []WorkflowPayload`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowCapabilityPayload) GetWorkflowsOk() (*[]WorkflowPayload, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowCapabilityPayload) SetWorkflows(v []WorkflowPayload)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowCapabilityPayload) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


