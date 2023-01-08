# WorkflowSchemeProjectSwitchBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingsByIssueTypeOverride** | Pointer to [**[]MappingsByIssueTypeOverride**](MappingsByIssueTypeOverride.md) | The mappings for migrating issues from old statuses to new statuses when switching from one workflow scheme to another. This field is required if any statuses in the current project&#39;s workflows would no longer exist in the target workflow scheme. Each mapping defines how to update issues from an old status to the corresponding new status in the issue’s new workflow. | [optional] 
**ProjectId** | Pointer to **string** | The ID of the project to switch the workflow scheme for | [optional] 
**TargetSchemeId** | Pointer to **string** | The ID of the target workflow scheme to switch to | [optional] 

## Methods

### NewWorkflowSchemeProjectSwitchBean

`func NewWorkflowSchemeProjectSwitchBean() *WorkflowSchemeProjectSwitchBean`

NewWorkflowSchemeProjectSwitchBean instantiates a new WorkflowSchemeProjectSwitchBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeProjectSwitchBeanWithDefaults

`func NewWorkflowSchemeProjectSwitchBeanWithDefaults() *WorkflowSchemeProjectSwitchBean`

NewWorkflowSchemeProjectSwitchBeanWithDefaults instantiates a new WorkflowSchemeProjectSwitchBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappingsByIssueTypeOverride

`func (o *WorkflowSchemeProjectSwitchBean) GetMappingsByIssueTypeOverride() []MappingsByIssueTypeOverride`

GetMappingsByIssueTypeOverride returns the MappingsByIssueTypeOverride field if non-nil, zero value otherwise.

### GetMappingsByIssueTypeOverrideOk

`func (o *WorkflowSchemeProjectSwitchBean) GetMappingsByIssueTypeOverrideOk() (*[]MappingsByIssueTypeOverride, bool)`

GetMappingsByIssueTypeOverrideOk returns a tuple with the MappingsByIssueTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingsByIssueTypeOverride

`func (o *WorkflowSchemeProjectSwitchBean) SetMappingsByIssueTypeOverride(v []MappingsByIssueTypeOverride)`

SetMappingsByIssueTypeOverride sets MappingsByIssueTypeOverride field to given value.

### HasMappingsByIssueTypeOverride

`func (o *WorkflowSchemeProjectSwitchBean) HasMappingsByIssueTypeOverride() bool`

HasMappingsByIssueTypeOverride returns a boolean if a field has been set.

### GetProjectId

`func (o *WorkflowSchemeProjectSwitchBean) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkflowSchemeProjectSwitchBean) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkflowSchemeProjectSwitchBean) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WorkflowSchemeProjectSwitchBean) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTargetSchemeId

`func (o *WorkflowSchemeProjectSwitchBean) GetTargetSchemeId() string`

GetTargetSchemeId returns the TargetSchemeId field if non-nil, zero value otherwise.

### GetTargetSchemeIdOk

`func (o *WorkflowSchemeProjectSwitchBean) GetTargetSchemeIdOk() (*string, bool)`

GetTargetSchemeIdOk returns a tuple with the TargetSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSchemeId

`func (o *WorkflowSchemeProjectSwitchBean) SetTargetSchemeId(v string)`

SetTargetSchemeId sets TargetSchemeId field to given value.

### HasTargetSchemeId

`func (o *WorkflowSchemeProjectSwitchBean) HasTargetSchemeId() bool`

HasTargetSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


