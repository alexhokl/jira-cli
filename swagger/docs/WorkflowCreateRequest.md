# WorkflowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**WorkflowScope**](WorkflowScope.md) |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowStatusUpdate**](WorkflowStatusUpdate.md) | The statuses to associate with the workflows. | [optional] 
**Workflows** | Pointer to [**[]WorkflowCreate**](WorkflowCreate.md) | The details of the workflows to create. | [optional] 

## Methods

### NewWorkflowCreateRequest

`func NewWorkflowCreateRequest() *WorkflowCreateRequest`

NewWorkflowCreateRequest instantiates a new WorkflowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCreateRequestWithDefaults

`func NewWorkflowCreateRequestWithDefaults() *WorkflowCreateRequest`

NewWorkflowCreateRequestWithDefaults instantiates a new WorkflowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *WorkflowCreateRequest) GetScope() WorkflowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowCreateRequest) GetScopeOk() (*WorkflowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowCreateRequest) SetScope(v WorkflowScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WorkflowCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowCreateRequest) GetStatuses() []WorkflowStatusUpdate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowCreateRequest) GetStatusesOk() (*[]WorkflowStatusUpdate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowCreateRequest) SetStatuses(v []WorkflowStatusUpdate)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowCreateRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowCreateRequest) GetWorkflows() []WorkflowCreate`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowCreateRequest) GetWorkflowsOk() (*[]WorkflowCreate, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowCreateRequest) SetWorkflows(v []WorkflowCreate)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowCreateRequest) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


