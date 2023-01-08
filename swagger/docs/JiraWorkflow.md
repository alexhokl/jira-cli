# JiraWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableString** | The creation date of the workflow. | [optional] 
**Description** | Pointer to **string** | The description of the workflow. | [optional] 
**Id** | Pointer to **string** | The ID of the workflow. | [optional] 
**IsEditable** | Pointer to **bool** | Indicates if the workflow can be edited. | [optional] 
**LoopedTransitionContainerLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the workflow. | [optional] 
**Scope** | Pointer to [**WorkflowScope**](WorkflowScope.md) |  | [optional] 
**StartPointLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowReferenceStatus**](WorkflowReferenceStatus.md) | The statuses referenced in this workflow. | [optional] 
**TaskId** | Pointer to **NullableString** | If there is a current [asynchronous task](#async-operations) operation for this workflow. | [optional] 
**Transitions** | Pointer to [**[]WorkflowTransitions**](WorkflowTransitions.md) | The transitions of the workflow. | [optional] 
**Updated** | Pointer to **NullableString** | The last edited date of the workflow. | [optional] 
**Version** | Pointer to [**DocumentVersion**](DocumentVersion.md) |  | [optional] 

## Methods

### NewJiraWorkflow

`func NewJiraWorkflow() *JiraWorkflow`

NewJiraWorkflow instantiates a new JiraWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraWorkflowWithDefaults

`func NewJiraWorkflowWithDefaults() *JiraWorkflow`

NewJiraWorkflowWithDefaults instantiates a new JiraWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *JiraWorkflow) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JiraWorkflow) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JiraWorkflow) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JiraWorkflow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *JiraWorkflow) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *JiraWorkflow) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetDescription

`func (o *JiraWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JiraWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JiraWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JiraWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *JiraWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JiraWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JiraWorkflow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JiraWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEditable

`func (o *JiraWorkflow) GetIsEditable() bool`

GetIsEditable returns the IsEditable field if non-nil, zero value otherwise.

### GetIsEditableOk

`func (o *JiraWorkflow) GetIsEditableOk() (*bool, bool)`

GetIsEditableOk returns a tuple with the IsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditable

`func (o *JiraWorkflow) SetIsEditable(v bool)`

SetIsEditable sets IsEditable field to given value.

### HasIsEditable

`func (o *JiraWorkflow) HasIsEditable() bool`

HasIsEditable returns a boolean if a field has been set.

### GetLoopedTransitionContainerLayout

`func (o *JiraWorkflow) GetLoopedTransitionContainerLayout() WorkflowLayout`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *JiraWorkflow) GetLoopedTransitionContainerLayoutOk() (*WorkflowLayout, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *JiraWorkflow) SetLoopedTransitionContainerLayout(v WorkflowLayout)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *JiraWorkflow) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### SetLoopedTransitionContainerLayoutNil

`func (o *JiraWorkflow) SetLoopedTransitionContainerLayoutNil(b bool)`

 SetLoopedTransitionContainerLayoutNil sets the value for LoopedTransitionContainerLayout to be an explicit nil

### UnsetLoopedTransitionContainerLayout
`func (o *JiraWorkflow) UnsetLoopedTransitionContainerLayout()`

UnsetLoopedTransitionContainerLayout ensures that no value is present for LoopedTransitionContainerLayout, not even an explicit nil
### GetName

`func (o *JiraWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JiraWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JiraWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JiraWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *JiraWorkflow) GetScope() WorkflowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *JiraWorkflow) GetScopeOk() (*WorkflowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *JiraWorkflow) SetScope(v WorkflowScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *JiraWorkflow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStartPointLayout

`func (o *JiraWorkflow) GetStartPointLayout() WorkflowLayout`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *JiraWorkflow) GetStartPointLayoutOk() (*WorkflowLayout, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *JiraWorkflow) SetStartPointLayout(v WorkflowLayout)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *JiraWorkflow) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### SetStartPointLayoutNil

`func (o *JiraWorkflow) SetStartPointLayoutNil(b bool)`

 SetStartPointLayoutNil sets the value for StartPointLayout to be an explicit nil

### UnsetStartPointLayout
`func (o *JiraWorkflow) UnsetStartPointLayout()`

UnsetStartPointLayout ensures that no value is present for StartPointLayout, not even an explicit nil
### GetStatuses

`func (o *JiraWorkflow) GetStatuses() []WorkflowReferenceStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *JiraWorkflow) GetStatusesOk() (*[]WorkflowReferenceStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *JiraWorkflow) SetStatuses(v []WorkflowReferenceStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *JiraWorkflow) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTaskId

`func (o *JiraWorkflow) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *JiraWorkflow) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *JiraWorkflow) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *JiraWorkflow) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *JiraWorkflow) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *JiraWorkflow) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTransitions

`func (o *JiraWorkflow) GetTransitions() []WorkflowTransitions`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *JiraWorkflow) GetTransitionsOk() (*[]WorkflowTransitions, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *JiraWorkflow) SetTransitions(v []WorkflowTransitions)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *JiraWorkflow) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetUpdated

`func (o *JiraWorkflow) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *JiraWorkflow) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *JiraWorkflow) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *JiraWorkflow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *JiraWorkflow) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *JiraWorkflow) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetVersion

`func (o *JiraWorkflow) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JiraWorkflow) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JiraWorkflow) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *JiraWorkflow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


