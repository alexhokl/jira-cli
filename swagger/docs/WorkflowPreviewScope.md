# WorkflowPreviewScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**NullableWorkflowProjectIdScope**](WorkflowProjectIdScope.md) |  | [optional] 
**Type** | Pointer to **string** | The scope of the workflow. &#x60;GLOBAL&#x60; for company-managed projects and &#x60;PROJECT&#x60; for team-managed projects. | [optional] 

## Methods

### NewWorkflowPreviewScope

`func NewWorkflowPreviewScope() *WorkflowPreviewScope`

NewWorkflowPreviewScope instantiates a new WorkflowPreviewScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPreviewScopeWithDefaults

`func NewWorkflowPreviewScopeWithDefaults() *WorkflowPreviewScope`

NewWorkflowPreviewScopeWithDefaults instantiates a new WorkflowPreviewScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *WorkflowPreviewScope) GetProject() WorkflowProjectIdScope`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *WorkflowPreviewScope) GetProjectOk() (*WorkflowProjectIdScope, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *WorkflowPreviewScope) SetProject(v WorkflowProjectIdScope)`

SetProject sets Project field to given value.

### HasProject

`func (o *WorkflowPreviewScope) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *WorkflowPreviewScope) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *WorkflowPreviewScope) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetType

`func (o *WorkflowPreviewScope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowPreviewScope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowPreviewScope) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowPreviewScope) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


