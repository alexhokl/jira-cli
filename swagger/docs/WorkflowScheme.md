# WorkflowScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflow** | Pointer to **string** | The name of the default workflow for the workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. If &#x60;defaultWorkflow&#x60; is not specified when creating a workflow scheme, it is set to *Jira Workflow (jira)*. | [optional] 
**Description** | Pointer to **string** | The description of the workflow scheme. | [optional] 
**Draft** | Pointer to **bool** | Whether the workflow scheme is a draft or not. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the workflow scheme. | [optional] [readonly] 
**IssueTypeMappings** | Pointer to **map[string]string** | The issue type to workflow mappings, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme. | [optional] 
**IssueTypes** | Pointer to [**map[string]IssueTypeDetails**](IssueTypeDetails.md) | The issue types available in Jira. | [optional] [readonly] 
**LastModified** | Pointer to **string** | The date-time that the draft workflow scheme was last modified. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows. | [optional] [readonly] 
**LastModifiedUser** | Pointer to [**User**](User.md) | The user that last modified the draft workflow scheme. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the workflow scheme. The name must be unique. The maximum length is 255 characters. Required when creating a workflow scheme. | [optional] 
**OriginalDefaultWorkflow** | Pointer to **string** | For draft workflow schemes, this property is the name of the default workflow for the original workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. | [optional] [readonly] 
**OriginalIssueTypeMappings** | Pointer to **map[string]string** | For draft workflow schemes, this property is the issue type to workflow mappings for the original workflow scheme, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme. | [optional] [readonly] 
**Self** | Pointer to **string** |  | [optional] [readonly] 
**UpdateDraftIfNeeded** | Pointer to **bool** | Whether to create or update a draft workflow scheme when updating an active workflow scheme. An active workflow scheme is a workflow scheme that is used by at least one project. The following examples show how this property works:   *  Update an active workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;true&#x60;: If a draft workflow scheme exists, it is updated. Otherwise, a draft workflow scheme is created.  *  Update an active workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;false&#x60;: An error is returned, as active workflow schemes cannot be updated.  *  Update an inactive workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;true&#x60;: The workflow scheme is updated, as inactive workflow schemes do not require drafts to update.  Defaults to &#x60;false&#x60;. | [optional] 

## Methods

### NewWorkflowScheme

`func NewWorkflowScheme() *WorkflowScheme`

NewWorkflowScheme instantiates a new WorkflowScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeWithDefaults

`func NewWorkflowSchemeWithDefaults() *WorkflowScheme`

NewWorkflowSchemeWithDefaults instantiates a new WorkflowScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflow

`func (o *WorkflowScheme) GetDefaultWorkflow() string`

GetDefaultWorkflow returns the DefaultWorkflow field if non-nil, zero value otherwise.

### GetDefaultWorkflowOk

`func (o *WorkflowScheme) GetDefaultWorkflowOk() (*string, bool)`

GetDefaultWorkflowOk returns a tuple with the DefaultWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflow

`func (o *WorkflowScheme) SetDefaultWorkflow(v string)`

SetDefaultWorkflow sets DefaultWorkflow field to given value.

### HasDefaultWorkflow

`func (o *WorkflowScheme) HasDefaultWorkflow() bool`

HasDefaultWorkflow returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDraft

`func (o *WorkflowScheme) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *WorkflowScheme) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *WorkflowScheme) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *WorkflowScheme) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetId

`func (o *WorkflowScheme) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowScheme) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowScheme) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTypeMappings

`func (o *WorkflowScheme) GetIssueTypeMappings() map[string]string`

GetIssueTypeMappings returns the IssueTypeMappings field if non-nil, zero value otherwise.

### GetIssueTypeMappingsOk

`func (o *WorkflowScheme) GetIssueTypeMappingsOk() (*map[string]string, bool)`

GetIssueTypeMappingsOk returns a tuple with the IssueTypeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeMappings

`func (o *WorkflowScheme) SetIssueTypeMappings(v map[string]string)`

SetIssueTypeMappings sets IssueTypeMappings field to given value.

### HasIssueTypeMappings

`func (o *WorkflowScheme) HasIssueTypeMappings() bool`

HasIssueTypeMappings returns a boolean if a field has been set.

### GetIssueTypes

`func (o *WorkflowScheme) GetIssueTypes() map[string]IssueTypeDetails`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *WorkflowScheme) GetIssueTypesOk() (*map[string]IssueTypeDetails, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *WorkflowScheme) SetIssueTypes(v map[string]IssueTypeDetails)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *WorkflowScheme) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetLastModified

`func (o *WorkflowScheme) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *WorkflowScheme) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *WorkflowScheme) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *WorkflowScheme) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastModifiedUser

`func (o *WorkflowScheme) GetLastModifiedUser() User`

GetLastModifiedUser returns the LastModifiedUser field if non-nil, zero value otherwise.

### GetLastModifiedUserOk

`func (o *WorkflowScheme) GetLastModifiedUserOk() (*User, bool)`

GetLastModifiedUserOk returns a tuple with the LastModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedUser

`func (o *WorkflowScheme) SetLastModifiedUser(v User)`

SetLastModifiedUser sets LastModifiedUser field to given value.

### HasLastModifiedUser

`func (o *WorkflowScheme) HasLastModifiedUser() bool`

HasLastModifiedUser returns a boolean if a field has been set.

### GetName

`func (o *WorkflowScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalDefaultWorkflow

`func (o *WorkflowScheme) GetOriginalDefaultWorkflow() string`

GetOriginalDefaultWorkflow returns the OriginalDefaultWorkflow field if non-nil, zero value otherwise.

### GetOriginalDefaultWorkflowOk

`func (o *WorkflowScheme) GetOriginalDefaultWorkflowOk() (*string, bool)`

GetOriginalDefaultWorkflowOk returns a tuple with the OriginalDefaultWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDefaultWorkflow

`func (o *WorkflowScheme) SetOriginalDefaultWorkflow(v string)`

SetOriginalDefaultWorkflow sets OriginalDefaultWorkflow field to given value.

### HasOriginalDefaultWorkflow

`func (o *WorkflowScheme) HasOriginalDefaultWorkflow() bool`

HasOriginalDefaultWorkflow returns a boolean if a field has been set.

### GetOriginalIssueTypeMappings

`func (o *WorkflowScheme) GetOriginalIssueTypeMappings() map[string]string`

GetOriginalIssueTypeMappings returns the OriginalIssueTypeMappings field if non-nil, zero value otherwise.

### GetOriginalIssueTypeMappingsOk

`func (o *WorkflowScheme) GetOriginalIssueTypeMappingsOk() (*map[string]string, bool)`

GetOriginalIssueTypeMappingsOk returns a tuple with the OriginalIssueTypeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalIssueTypeMappings

`func (o *WorkflowScheme) SetOriginalIssueTypeMappings(v map[string]string)`

SetOriginalIssueTypeMappings sets OriginalIssueTypeMappings field to given value.

### HasOriginalIssueTypeMappings

`func (o *WorkflowScheme) HasOriginalIssueTypeMappings() bool`

HasOriginalIssueTypeMappings returns a boolean if a field has been set.

### GetSelf

`func (o *WorkflowScheme) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *WorkflowScheme) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *WorkflowScheme) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *WorkflowScheme) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetUpdateDraftIfNeeded

`func (o *WorkflowScheme) GetUpdateDraftIfNeeded() bool`

GetUpdateDraftIfNeeded returns the UpdateDraftIfNeeded field if non-nil, zero value otherwise.

### GetUpdateDraftIfNeededOk

`func (o *WorkflowScheme) GetUpdateDraftIfNeededOk() (*bool, bool)`

GetUpdateDraftIfNeededOk returns a tuple with the UpdateDraftIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDraftIfNeeded

`func (o *WorkflowScheme) SetUpdateDraftIfNeeded(v bool)`

SetUpdateDraftIfNeeded sets UpdateDraftIfNeeded field to given value.

### HasUpdateDraftIfNeeded

`func (o *WorkflowScheme) HasUpdateDraftIfNeeded() bool`

HasUpdateDraftIfNeeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


