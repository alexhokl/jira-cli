# StatusesPerWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialStatusId** | Pointer to **string** | The ID of the initial status for the workflow. | [optional] 
**Statuses** | Pointer to **[]string** | The status IDs associated with the workflow. | [optional] 
**WorkflowId** | Pointer to **string** | The ID of the workflow. | [optional] 

## Methods

### NewStatusesPerWorkflow

`func NewStatusesPerWorkflow() *StatusesPerWorkflow`

NewStatusesPerWorkflow instantiates a new StatusesPerWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusesPerWorkflowWithDefaults

`func NewStatusesPerWorkflowWithDefaults() *StatusesPerWorkflow`

NewStatusesPerWorkflowWithDefaults instantiates a new StatusesPerWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialStatusId

`func (o *StatusesPerWorkflow) GetInitialStatusId() string`

GetInitialStatusId returns the InitialStatusId field if non-nil, zero value otherwise.

### GetInitialStatusIdOk

`func (o *StatusesPerWorkflow) GetInitialStatusIdOk() (*string, bool)`

GetInitialStatusIdOk returns a tuple with the InitialStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialStatusId

`func (o *StatusesPerWorkflow) SetInitialStatusId(v string)`

SetInitialStatusId sets InitialStatusId field to given value.

### HasInitialStatusId

`func (o *StatusesPerWorkflow) HasInitialStatusId() bool`

HasInitialStatusId returns a boolean if a field has been set.

### GetStatuses

`func (o *StatusesPerWorkflow) GetStatuses() []string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *StatusesPerWorkflow) GetStatusesOk() (*[]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *StatusesPerWorkflow) SetStatuses(v []string)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *StatusesPerWorkflow) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflowId

`func (o *StatusesPerWorkflow) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *StatusesPerWorkflow) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *StatusesPerWorkflow) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *StatusesPerWorkflow) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


