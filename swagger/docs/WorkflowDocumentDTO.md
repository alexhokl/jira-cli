# WorkflowDocumentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdateAuthorAAID** | Pointer to **string** |  | [optional] 
**LoopedTransitionContainerLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**WorkflowScope**](WorkflowScope.md) |  | [optional] 
**StartPointLayout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Statuses** | Pointer to [**[]WorkflowReferenceStatus**](WorkflowReferenceStatus.md) |  | [optional] 
**Transitions** | Pointer to [**[]WorkflowTransitions**](WorkflowTransitions.md) |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**DocumentVersion**](DocumentVersion.md) |  | [optional] 

## Methods

### NewWorkflowDocumentDTO

`func NewWorkflowDocumentDTO() *WorkflowDocumentDTO`

NewWorkflowDocumentDTO instantiates a new WorkflowDocumentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDocumentDTOWithDefaults

`func NewWorkflowDocumentDTOWithDefaults() *WorkflowDocumentDTO`

NewWorkflowDocumentDTOWithDefaults instantiates a new WorkflowDocumentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *WorkflowDocumentDTO) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WorkflowDocumentDTO) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WorkflowDocumentDTO) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WorkflowDocumentDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowDocumentDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowDocumentDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowDocumentDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowDocumentDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkflowDocumentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowDocumentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowDocumentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowDocumentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdateAuthorAAID

`func (o *WorkflowDocumentDTO) GetLastUpdateAuthorAAID() string`

GetLastUpdateAuthorAAID returns the LastUpdateAuthorAAID field if non-nil, zero value otherwise.

### GetLastUpdateAuthorAAIDOk

`func (o *WorkflowDocumentDTO) GetLastUpdateAuthorAAIDOk() (*string, bool)`

GetLastUpdateAuthorAAIDOk returns a tuple with the LastUpdateAuthorAAID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAuthorAAID

`func (o *WorkflowDocumentDTO) SetLastUpdateAuthorAAID(v string)`

SetLastUpdateAuthorAAID sets LastUpdateAuthorAAID field to given value.

### HasLastUpdateAuthorAAID

`func (o *WorkflowDocumentDTO) HasLastUpdateAuthorAAID() bool`

HasLastUpdateAuthorAAID returns a boolean if a field has been set.

### GetLoopedTransitionContainerLayout

`func (o *WorkflowDocumentDTO) GetLoopedTransitionContainerLayout() WorkflowLayout`

GetLoopedTransitionContainerLayout returns the LoopedTransitionContainerLayout field if non-nil, zero value otherwise.

### GetLoopedTransitionContainerLayoutOk

`func (o *WorkflowDocumentDTO) GetLoopedTransitionContainerLayoutOk() (*WorkflowLayout, bool)`

GetLoopedTransitionContainerLayoutOk returns a tuple with the LoopedTransitionContainerLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopedTransitionContainerLayout

`func (o *WorkflowDocumentDTO) SetLoopedTransitionContainerLayout(v WorkflowLayout)`

SetLoopedTransitionContainerLayout sets LoopedTransitionContainerLayout field to given value.

### HasLoopedTransitionContainerLayout

`func (o *WorkflowDocumentDTO) HasLoopedTransitionContainerLayout() bool`

HasLoopedTransitionContainerLayout returns a boolean if a field has been set.

### SetLoopedTransitionContainerLayoutNil

`func (o *WorkflowDocumentDTO) SetLoopedTransitionContainerLayoutNil(b bool)`

 SetLoopedTransitionContainerLayoutNil sets the value for LoopedTransitionContainerLayout to be an explicit nil

### UnsetLoopedTransitionContainerLayout
`func (o *WorkflowDocumentDTO) UnsetLoopedTransitionContainerLayout()`

UnsetLoopedTransitionContainerLayout ensures that no value is present for LoopedTransitionContainerLayout, not even an explicit nil
### GetName

`func (o *WorkflowDocumentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowDocumentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowDocumentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowDocumentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *WorkflowDocumentDTO) GetScope() WorkflowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowDocumentDTO) GetScopeOk() (*WorkflowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowDocumentDTO) SetScope(v WorkflowScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WorkflowDocumentDTO) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStartPointLayout

`func (o *WorkflowDocumentDTO) GetStartPointLayout() WorkflowLayout`

GetStartPointLayout returns the StartPointLayout field if non-nil, zero value otherwise.

### GetStartPointLayoutOk

`func (o *WorkflowDocumentDTO) GetStartPointLayoutOk() (*WorkflowLayout, bool)`

GetStartPointLayoutOk returns a tuple with the StartPointLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPointLayout

`func (o *WorkflowDocumentDTO) SetStartPointLayout(v WorkflowLayout)`

SetStartPointLayout sets StartPointLayout field to given value.

### HasStartPointLayout

`func (o *WorkflowDocumentDTO) HasStartPointLayout() bool`

HasStartPointLayout returns a boolean if a field has been set.

### SetStartPointLayoutNil

`func (o *WorkflowDocumentDTO) SetStartPointLayoutNil(b bool)`

 SetStartPointLayoutNil sets the value for StartPointLayout to be an explicit nil

### UnsetStartPointLayout
`func (o *WorkflowDocumentDTO) UnsetStartPointLayout()`

UnsetStartPointLayout ensures that no value is present for StartPointLayout, not even an explicit nil
### GetStatuses

`func (o *WorkflowDocumentDTO) GetStatuses() []WorkflowReferenceStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowDocumentDTO) GetStatusesOk() (*[]WorkflowReferenceStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowDocumentDTO) SetStatuses(v []WorkflowReferenceStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowDocumentDTO) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTransitions

`func (o *WorkflowDocumentDTO) GetTransitions() []WorkflowTransitions`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *WorkflowDocumentDTO) GetTransitionsOk() (*[]WorkflowTransitions, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *WorkflowDocumentDTO) SetTransitions(v []WorkflowTransitions)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *WorkflowDocumentDTO) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.

### GetUpdated

`func (o *WorkflowDocumentDTO) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *WorkflowDocumentDTO) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *WorkflowDocumentDTO) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *WorkflowDocumentDTO) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowDocumentDTO) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowDocumentDTO) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowDocumentDTO) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowDocumentDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


