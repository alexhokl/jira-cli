# WorkflowSchemeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflowId** | Pointer to **string** | The ID of the workflow for issue types without having a mapping defined in this workflow scheme. Only used in global-scoped workflow schemes. If the &#x60;defaultWorkflowId&#x60; isn&#39;t specified, this is set to *Jira Workflow (jira)*. | [optional] 
**Description** | **string** | The new description for this workflow scheme. | 
**Id** | **string** | The ID of this workflow scheme. | 
**Name** | **string** | The new name for this workflow scheme. | 
**StatusMappingsByIssueTypeOverride** | Pointer to [**[]MappingsByIssueTypeOverride**](MappingsByIssueTypeOverride.md) | Overrides, for the selected issue types, any status mappings provided in &#x60;statusMappingsByWorkflows&#x60;. Status mappings are required when the new workflow for an issue type doesn&#39;t contain all statuses that the old workflow has. Status mappings can be provided by a combination of &#x60;statusMappingsByWorkflows&#x60; and &#x60;statusMappingsByIssueTypeOverride&#x60;. | [optional] 
**StatusMappingsByWorkflows** | Pointer to [**[]MappingsByWorkflow**](MappingsByWorkflow.md) | The status mappings by workflows. Status mappings are required when the new workflow for an issue type doesn&#39;t contain all statuses that the old workflow has. Status mappings can be provided by a combination of &#x60;statusMappingsByWorkflows&#x60; and &#x60;statusMappingsByIssueTypeOverride&#x60;. | [optional] 
**Version** | [**DocumentVersion**](DocumentVersion.md) |  | 
**WorkflowsForIssueTypes** | Pointer to [**[]WorkflowSchemeAssociation**](WorkflowSchemeAssociation.md) | Mappings from workflows to issue types. | [optional] 

## Methods

### NewWorkflowSchemeUpdateRequest

`func NewWorkflowSchemeUpdateRequest(description string, id string, name string, version DocumentVersion, ) *WorkflowSchemeUpdateRequest`

NewWorkflowSchemeUpdateRequest instantiates a new WorkflowSchemeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeUpdateRequestWithDefaults

`func NewWorkflowSchemeUpdateRequestWithDefaults() *WorkflowSchemeUpdateRequest`

NewWorkflowSchemeUpdateRequestWithDefaults instantiates a new WorkflowSchemeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequest) GetDefaultWorkflowId() string`

GetDefaultWorkflowId returns the DefaultWorkflowId field if non-nil, zero value otherwise.

### GetDefaultWorkflowIdOk

`func (o *WorkflowSchemeUpdateRequest) GetDefaultWorkflowIdOk() (*string, bool)`

GetDefaultWorkflowIdOk returns a tuple with the DefaultWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequest) SetDefaultWorkflowId(v string)`

SetDefaultWorkflowId sets DefaultWorkflowId field to given value.

### HasDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequest) HasDefaultWorkflowId() bool`

HasDefaultWorkflowId returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchemeUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchemeUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchemeUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *WorkflowSchemeUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSchemeUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSchemeUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowSchemeUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchemeUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchemeUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatusMappingsByIssueTypeOverride

`func (o *WorkflowSchemeUpdateRequest) GetStatusMappingsByIssueTypeOverride() []MappingsByIssueTypeOverride`

GetStatusMappingsByIssueTypeOverride returns the StatusMappingsByIssueTypeOverride field if non-nil, zero value otherwise.

### GetStatusMappingsByIssueTypeOverrideOk

`func (o *WorkflowSchemeUpdateRequest) GetStatusMappingsByIssueTypeOverrideOk() (*[]MappingsByIssueTypeOverride, bool)`

GetStatusMappingsByIssueTypeOverrideOk returns a tuple with the StatusMappingsByIssueTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappingsByIssueTypeOverride

`func (o *WorkflowSchemeUpdateRequest) SetStatusMappingsByIssueTypeOverride(v []MappingsByIssueTypeOverride)`

SetStatusMappingsByIssueTypeOverride sets StatusMappingsByIssueTypeOverride field to given value.

### HasStatusMappingsByIssueTypeOverride

`func (o *WorkflowSchemeUpdateRequest) HasStatusMappingsByIssueTypeOverride() bool`

HasStatusMappingsByIssueTypeOverride returns a boolean if a field has been set.

### GetStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequest) GetStatusMappingsByWorkflows() []MappingsByWorkflow`

GetStatusMappingsByWorkflows returns the StatusMappingsByWorkflows field if non-nil, zero value otherwise.

### GetStatusMappingsByWorkflowsOk

`func (o *WorkflowSchemeUpdateRequest) GetStatusMappingsByWorkflowsOk() (*[]MappingsByWorkflow, bool)`

GetStatusMappingsByWorkflowsOk returns a tuple with the StatusMappingsByWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequest) SetStatusMappingsByWorkflows(v []MappingsByWorkflow)`

SetStatusMappingsByWorkflows sets StatusMappingsByWorkflows field to given value.

### HasStatusMappingsByWorkflows

`func (o *WorkflowSchemeUpdateRequest) HasStatusMappingsByWorkflows() bool`

HasStatusMappingsByWorkflows returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSchemeUpdateRequest) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSchemeUpdateRequest) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSchemeUpdateRequest) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.


### GetWorkflowsForIssueTypes

`func (o *WorkflowSchemeUpdateRequest) GetWorkflowsForIssueTypes() []WorkflowSchemeAssociation`

GetWorkflowsForIssueTypes returns the WorkflowsForIssueTypes field if non-nil, zero value otherwise.

### GetWorkflowsForIssueTypesOk

`func (o *WorkflowSchemeUpdateRequest) GetWorkflowsForIssueTypesOk() (*[]WorkflowSchemeAssociation, bool)`

GetWorkflowsForIssueTypesOk returns a tuple with the WorkflowsForIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsForIssueTypes

`func (o *WorkflowSchemeUpdateRequest) SetWorkflowsForIssueTypes(v []WorkflowSchemeAssociation)`

SetWorkflowsForIssueTypes sets WorkflowsForIssueTypes field to given value.

### HasWorkflowsForIssueTypes

`func (o *WorkflowSchemeUpdateRequest) HasWorkflowsForIssueTypes() bool`

HasWorkflowsForIssueTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


