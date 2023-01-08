# WorkflowUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStatusMappings** | Pointer to [**[]StatusMigration**](StatusMigration.md) | The mapping of old to new status ID. | [optional] 
**Description** | Pointer to **string** | The new description for this workflow. | [optional] 
**Id** | **string** | The ID of this workflow. | 
**LoopedTransitionContainerLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**StartPointLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**StatusMappings** | Pointer to [**[]StatusMappingDTO**](StatusMappingDTO.md) | The mapping of old to new status ID for a specific project and issue type. | [optional] 
**Statuses** | [**[]StatusLayoutUpdate**](StatusLayoutUpdate.md) | The statuses associated with this workflow. | 
**Transitions** | [**[]TransitionUpdateDTO**](TransitionUpdateDTO.md) | The transitions of this workflow. | 
**Version** | [**DocumentVersion**](DocumentVersion.md) |  | 

## Methods

### NewWorkflowUpdate

`func NewWorkflowUpdate(id string, statuses []StatusLayoutUpdate, transitions []TransitionUpdateDTO, version DocumentVersion, ) *WorkflowUpdate`

NewWorkflowUpdate instantiates a new WorkflowUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUpdateWithDefaults

`func NewWorkflowUpdateWithDefaults() *WorkflowUpdate`

NewWorkflowUpdateWithDefaults instantiates a new WorkflowUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStatusMappings

`func (o *WorkflowUpdate) GetDefaultStatusMappings() []StatusMigration`

GetDefaultStatusMappings returns the DefaultStatusMappings field if non-nil, zero value otherwise.

### GetDefaultStatusMappingsOk

`func (o *WorkflowUpdate) GetDefaultStatusMappingsOk() (*[]StatusMigration, bool)`

GetDefaultStatusMappingsOk returns a tuple with the DefaultStatusMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatusMappings

`func (o *WorkflowUpdate) SetDefaultStatusMappings(v []StatusMigration)`

SetDefaultStatusMappings sets DefaultStatusMappings field to given value.

### HasDefaultStatusMappings

`func (o *WorkflowUpdate) HasDefaultStatusMappings() bool`

HasDefaultStatusMappings returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkflowUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetLoopedTransitionContainerLayout

`func (o *WorkflowUpdate) GetLoopedTransitionContainerLayout() WorkflowLayout`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *WorkflowUpdate) GetLoopedTransitionContainerLayoutOk() (*WorkflowLayout, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *WorkflowUpdate) SetLoopedTransitionContainerLayout(v WorkflowLayout)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *WorkflowUpdate) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### SetLoopedTransitionContainerLayoutNil

`func (o *WorkflowUpdate) SetLoopedTransitionContainerLayoutNil(b bool)`

 SetLoopedTransitionContainerLayoutNil sets the value for LoopedTransitionContainerLayout to be an explicit nil

### UnsetLoopedTransitionContainerLayout
`func (o *WorkflowUpdate) UnsetLoopedTransitionContainerLayout()`

UnsetLoopedTransitionContainerLayout ensures that no value is present for LoopedTransitionContainerLayout, not even an explicit nil
### GetStartPointLayout

`func (o *WorkflowUpdate) GetStartPointLayout() WorkflowLayout`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *WorkflowUpdate) GetStartPointLayoutOk() (*WorkflowLayout, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *WorkflowUpdate) SetStartPointLayout(v WorkflowLayout)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *WorkflowUpdate) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### SetStartPointLayoutNil

`func (o *WorkflowUpdate) SetStartPointLayoutNil(b bool)`

 SetStartPointLayoutNil sets the value for StartPointLayout to be an explicit nil

### UnsetStartPointLayout
`func (o *WorkflowUpdate) UnsetStartPointLayout()`

UnsetStartPointLayout ensures that no value is present for StartPointLayout, not even an explicit nil
### GetStatusMappings

`func (o *WorkflowUpdate) GetStatusMappings() []StatusMappingDTO`

GetStatusMappings returns the StatusMappings field if non-nil, zero value otherwise.

### GetStatusMappingsOk

`func (o *WorkflowUpdate) GetStatusMappingsOk() (*[]StatusMappingDTO, bool)`

GetStatusMappingsOk returns a tuple with the StatusMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappings

`func (o *WorkflowUpdate) SetStatusMappings(v []StatusMappingDTO)`

SetStatusMappings sets StatusMappings field to given value.

### HasStatusMappings

`func (o *WorkflowUpdate) HasStatusMappings() bool`

HasStatusMappings returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowUpdate) GetStatuses() []StatusLayoutUpdate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowUpdate) GetStatusesOk() (*[]StatusLayoutUpdate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowUpdate) SetStatuses(v []StatusLayoutUpdate)`

SetStatuses sets Statuses field to given value.


### GetTransitions

`func (o *WorkflowUpdate) GetTransitions() []TransitionUpdateDTO`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowUpdate) GetTransitionsOk() (*[]TransitionUpdateDTO, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowUpdate) SetTransitions(v []TransitionUpdateDTO)`

SetTransitions sets Transitions field to given value.


### GetVersion

`func (o *WorkflowUpdate) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowUpdate) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowUpdate) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


