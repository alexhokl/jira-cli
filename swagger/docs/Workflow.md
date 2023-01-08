# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The creation date of the workflow. | [optional] 
**Description** | **string** | The description of the workflow. | 
**HasDraftWorkflow** | Pointer to **bool** | Whether the workflow has a draft version. | [optional] 
**Id** | [**PublishedWorkflowId**](PublishedWorkflowId.md) |  | 
**IsDefault** | Pointer to **bool** | Whether this is the default workflow. | [optional] 
**Operations** | Pointer to [**WorkflowOperations**](WorkflowOperations.md) |  | [optional] 
**Projects** | Pointer to [**[]ProjectDetails**](ProjectDetails.md) | The projects the workflow is assigned to, through workflow schemes. | [optional] 
**Schemes** | Pointer to [**[]WorkflowSchemeIdName**](WorkflowSchemeIdName.md) | The workflow schemes the workflow is assigned to. | [optional] 
**Statuses** | Pointer to [**[]WorkflowStatus**](WorkflowStatus.md) | The statuses of the workflow. | [optional] 
**Transitions** | Pointer to [**[]Transition**](Transition.md) | The transitions of the workflow. | [optional] 
**Updated** | Pointer to **time.Time** | The last edited date of the workflow. | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow(description string, id PublishedWorkflowId, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Workflow) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Workflow) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Workflow) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Workflow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Workflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Workflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Workflow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHasDraftWorkflow

`func (o *Workflow) GetHasDraftWorkflow() bool`

GetHasDraftWorkflow returns the HasDraftWorkflow field if non-nil, zero value otherwise.

### GetHasDraftWorkflowOk

`func (o *Workflow) GetHasDraftWorkflowOk() (*bool, bool)`

GetHasDraftWorkflowOk returns a tuple with the HasDraftWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDraftWorkflow

`func (o *Workflow) SetHasDraftWorkflow(v bool)`

SetHasDraftWorkflow sets HasDraftWorkflow field to given value.

### HasHasDraftWorkflow

`func (o *Workflow) HasHasDraftWorkflow() bool`

HasHasDraftWorkflow returns a boolean if a field has been set.

### GetId

`func (o *Workflow) GetId() PublishedWorkflowId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*PublishedWorkflowId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v PublishedWorkflowId)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *Workflow) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Workflow) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Workflow) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Workflow) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOperations

`func (o *Workflow) GetOperations() WorkflowOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Workflow) GetOperationsOk() (*WorkflowOperations, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Workflow) SetOperations(v WorkflowOperations)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Workflow) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetProjects

`func (o *Workflow) GetProjects() []ProjectDetails`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Workflow) GetProjectsOk() (*[]ProjectDetails, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Workflow) SetProjects(v []ProjectDetails)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Workflow) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSchemes

`func (o *Workflow) GetSchemes() []WorkflowSchemeIdName`

GetSchemes returns the Schemes field if non-nil, zero value otherwise.

### GetSchemesOk

`func (o *Workflow) GetSchemesOk() (*[]WorkflowSchemeIdName, bool)`

GetSchemesOk returns a tuple with the Schemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemes

`func (o *Workflow) SetSchemes(v []WorkflowSchemeIdName)`

SetSchemes sets Schemes field to given value.

### HasSchemes

`func (o *Workflow) HasSchemes() bool`

HasSchemes returns a boolean if a field has been set.

### GetStatuses

`func (o *Workflow) GetStatuses() []WorkflowStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *Workflow) GetStatusesOk() (*[]WorkflowStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *Workflow) SetStatuses(v []WorkflowStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *Workflow) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTransitions

`func (o *Workflow) GetTransitions() []Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *Workflow) GetTransitionsOk() (*[]Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *Workflow) SetTransitions(v []Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *Workflow) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetUpdated

`func (o *Workflow) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Workflow) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Workflow) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Workflow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


