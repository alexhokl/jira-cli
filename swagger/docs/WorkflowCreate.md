# WorkflowCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the workflow to create. | [optional] 
**LoopedTransitionContainerLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Name** | **string** | The name of the workflow to create. | 
**StartPointLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Statuses** | [**[]StatusLayoutUpdate**](StatusLayoutUpdate.md) | The statuses associated with this workflow. | 
**Transitions** | [**[]TransitionUpdateDTO**](TransitionUpdateDTO.md) | The transitions of this workflow. | 

## Methods

### NewWorkflowCreate

`func NewWorkflowCreate(name string, statuses []StatusLayoutUpdate, transitions []TransitionUpdateDTO, ) *WorkflowCreate`

NewWorkflowCreate instantiates a new WorkflowCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCreateWithDefaults

`func NewWorkflowCreateWithDefaults() *WorkflowCreate`

NewWorkflowCreateWithDefaults instantiates a new WorkflowCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLoopedTransitionContainerLayout

`func (o *WorkflowCreate) GetLoopedTransitionContainerLayout() WorkflowLayout`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *WorkflowCreate) GetLoopedTransitionContainerLayoutOk() (*WorkflowLayout, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *WorkflowCreate) SetLoopedTransitionContainerLayout(v WorkflowLayout)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *WorkflowCreate) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### SetLoopedTransitionContainerLayoutNil

`func (o *WorkflowCreate) SetLoopedTransitionContainerLayoutNil(b bool)`

 SetLoopedTransitionContainerLayoutNil sets the value for LoopedTransitionContainerLayout to be an explicit nil

### UnsetLoopedTransitionContainerLayout
`func (o *WorkflowCreate) UnsetLoopedTransitionContainerLayout()`

UnsetLoopedTransitionContainerLayout ensures that no value is present for LoopedTransitionContainerLayout, not even an explicit nil
### GetName

`func (o *WorkflowCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCreate) SetName(v string)`

SetName sets Name field to given value.


### GetStartPointLayout

`func (o *WorkflowCreate) GetStartPointLayout() WorkflowLayout`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *WorkflowCreate) GetStartPointLayoutOk() (*WorkflowLayout, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *WorkflowCreate) SetStartPointLayout(v WorkflowLayout)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *WorkflowCreate) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### SetStartPointLayoutNil

`func (o *WorkflowCreate) SetStartPointLayoutNil(b bool)`

 SetStartPointLayoutNil sets the value for StartPointLayout to be an explicit nil

### UnsetStartPointLayout
`func (o *WorkflowCreate) UnsetStartPointLayout()`

UnsetStartPointLayout ensures that no value is present for StartPointLayout, not even an explicit nil
### GetStatuses

`func (o *WorkflowCreate) GetStatuses() []StatusLayoutUpdate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowCreate) GetStatusesOk() (*[]StatusLayoutUpdate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowCreate) SetStatuses(v []StatusLayoutUpdate)`

SetStatuses sets Statuses field to given value.


### GetTransitions

`func (o *WorkflowCreate) GetTransitions() []TransitionUpdateDTO`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowCreate) GetTransitionsOk() (*[]TransitionUpdateDTO, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowCreate) SetTransitions(v []TransitionUpdateDTO)`

SetTransitions sets Transitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


