# WorkflowPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the workflow. | [optional] 
**Id** | Pointer to **string** | The ID of the workflow. | [optional] 
**LoopedTransitionContainerLayout** | Pointer to [**WorkflowPreviewLayout**](WorkflowPreviewLayout.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the workflow. | [optional] 
**QueryContext** | Pointer to [**[]ProjectIssueTypeQueryContext**](ProjectIssueTypeQueryContext.md) | The project and issue type context for this workflow query. | [optional] 
**Scope** | Pointer to [**WorkflowPreviewScope**](WorkflowPreviewScope.md) |  | [optional] 
**StartPointLayout** | Pointer to [**WorkflowPreviewLayout**](WorkflowPreviewLayout.md) |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowPreviewStatus**](WorkflowPreviewStatus.md) | The statuses referenced in this workflow. | [optional] 
**Transitions** | Pointer to [**[]TransitionPreview**](TransitionPreview.md) | The transitions of the workflow. | [optional] 
**Version** | Pointer to [**WorkflowDocumentVersionBean**](WorkflowDocumentVersionBean.md) |  | [optional] 

## Methods

### NewWorkflowPreview

`func NewWorkflowPreview() *WorkflowPreview`

NewWorkflowPreview instantiates a new WorkflowPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPreviewWithDefaults

`func NewWorkflowPreviewWithDefaults() *WorkflowPreview`

NewWorkflowPreviewWithDefaults instantiates a new WorkflowPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowPreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowPreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowPreview) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowPreview) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkflowPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowPreview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowPreview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoopedTransitionContainerLayout

`func (o *WorkflowPreview) GetLoopedTransitionContainerLayout() WorkflowPreviewLayout`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *WorkflowPreview) GetLoopedTransitionContainerLayoutOk() (*WorkflowPreviewLayout, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *WorkflowPreview) SetLoopedTransitionContainerLayout(v WorkflowPreviewLayout)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *WorkflowPreview) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### GetName

`func (o *WorkflowPreview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPreview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPreview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPreview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryContext

`func (o *WorkflowPreview) GetQueryContext() []ProjectIssueTypeQueryContext`

GetQueryContext returns the QueryContext field if non-nil, zero value otherwise.

### GetQueryContextOk

`func (o *WorkflowPreview) GetQueryContextOk() (*[]ProjectIssueTypeQueryContext, bool)`

GetQueryContextOk returns a tuple with the QueryContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryContext

`func (o *WorkflowPreview) SetQueryContext(v []ProjectIssueTypeQueryContext)`

SetQueryContext sets QueryContext field to given value.

### HasQueryContext

`func (o *WorkflowPreview) HasQueryContext() bool`

HasQueryContext returns a boolean if a field has been set.

### GetScope

`func (o *WorkflowPreview) GetScope() WorkflowPreviewScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowPreview) GetScopeOk() (*WorkflowPreviewScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowPreview) SetScope(v WorkflowPreviewScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WorkflowPreview) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStartPointLayout

`func (o *WorkflowPreview) GetStartPointLayout() WorkflowPreviewLayout`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *WorkflowPreview) GetStartPointLayoutOk() (*WorkflowPreviewLayout, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *WorkflowPreview) SetStartPointLayout(v WorkflowPreviewLayout)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *WorkflowPreview) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowPreview) GetStatuses() []WorkflowPreviewStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowPreview) GetStatusesOk() (*[]WorkflowPreviewStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowPreview) SetStatuses(v []WorkflowPreviewStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowPreview) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTransitions

`func (o *WorkflowPreview) GetTransitions() []TransitionPreview`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowPreview) GetTransitionsOk() (*[]TransitionPreview, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowPreview) SetTransitions(v []TransitionPreview)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *WorkflowPreview) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowPreview) GetVersion() WorkflowDocumentVersionBean`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowPreview) GetVersionOk() (*WorkflowDocumentVersionBean, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowPreview) SetVersion(v WorkflowDocumentVersionBean)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowPreview) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


