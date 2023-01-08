# WorkflowPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the workflow | [optional] 
**LoopedTransitionContainerLayout** | Pointer to [**WorkflowStatusLayoutPayload**](WorkflowStatusLayoutPayload.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the workflow | [optional] 
**OnConflict** | Pointer to **string** | The strategy to use if there is a conflict with another workflow | [optional] [default to "NEW"]
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**StartPointLayout** | Pointer to [**WorkflowStatusLayoutPayload**](WorkflowStatusLayoutPayload.md) |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowStatusPayload**](WorkflowStatusPayload.md) | The statuses to be used in the workflow | [optional] 
**Transitions** | Pointer to [**[]TransitionPayload**](TransitionPayload.md) | The transitions for the workflow | [optional] 

## Methods

### NewWorkflowPayload

`func NewWorkflowPayload() *WorkflowPayload`

NewWorkflowPayload instantiates a new WorkflowPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPayloadWithDefaults

`func NewWorkflowPayloadWithDefaults() *WorkflowPayload`

NewWorkflowPayloadWithDefaults instantiates a new WorkflowPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoopedTransitionContainerLayout

`func (o *WorkflowPayload) GetLoopedTransitionContainerLayout() WorkflowStatusLayoutPayload`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *WorkflowPayload) GetLoopedTransitionContainerLayoutOk() (*WorkflowStatusLayoutPayload, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *WorkflowPayload) SetLoopedTransitionContainerLayout(v WorkflowStatusLayoutPayload)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *WorkflowPayload) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### GetName

`func (o *WorkflowPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *WorkflowPayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *WorkflowPayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *WorkflowPayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *WorkflowPayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *WorkflowPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *WorkflowPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *WorkflowPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *WorkflowPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetStartPointLayout

`func (o *WorkflowPayload) GetStartPointLayout() WorkflowStatusLayoutPayload`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *WorkflowPayload) GetStartPointLayoutOk() (*WorkflowStatusLayoutPayload, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *WorkflowPayload) SetStartPointLayout(v WorkflowStatusLayoutPayload)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *WorkflowPayload) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowPayload) GetStatuses() []WorkflowStatusPayload`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowPayload) GetStatusesOk() (*[]WorkflowStatusPayload, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowPayload) SetStatuses(v []WorkflowStatusPayload)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowPayload) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTransitions

`func (o *WorkflowPayload) GetTransitions() []TransitionPayload`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowPayload) GetTransitionsOk() (*[]TransitionPayload, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowPayload) SetTransitions(v []TransitionPayload)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *WorkflowPayload) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


