# WorkflowScheme

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflow** | **string** | The name of the default workflow for the workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. If &#x60;defaultWorkflow&#x60; is not specified when creating a workflow scheme, it is set to *Jira Workflow (jira)*. | [optional] [default to null]
**Description** | **string** | The description of the workflow scheme. | [optional] [default to null]
**Draft** | **bool** | Whether the workflow scheme is a draft or not. | [optional] [default to null]
**Id** | **int64** | The ID of the workflow scheme. | [optional] [default to null]
**IssueTypeMappings** | **map[string]string** | The issue type to workflow mappings, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme. | [optional] [default to null]
**IssueTypes** | [**map[string]IssueTypeDetails**](IssueTypeDetails.md) | The issue types available in Jira. | [optional] [default to null]
**LastModified** | **string** | The date-time that the draft workflow scheme was last modified. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows. | [optional] [default to null]
**LastModifiedUser** | [***AllOfWorkflowSchemeLastModifiedUser**](AllOfWorkflowSchemeLastModifiedUser.md) | The user that last modified the draft workflow scheme. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows. | [optional] [default to null]
**Name** | **string** | The name of the workflow scheme. The name must be unique. The maximum length is 255 characters. Required when creating a workflow scheme. | [optional] [default to null]
**OriginalDefaultWorkflow** | **string** | For draft workflow schemes, this property is the name of the default workflow for the original workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. | [optional] [default to null]
**OriginalIssueTypeMappings** | **map[string]string** | For draft workflow schemes, this property is the issue type to workflow mappings for the original workflow scheme, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme. | [optional] [default to null]
**Self** | **string** |  | [optional] [default to null]
**UpdateDraftIfNeeded** | **bool** | Whether to create or update a draft workflow scheme when updating an active workflow scheme. An active workflow scheme is a workflow scheme that is used by at least one project. The following examples show how this property works:   *  Update an active workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;true&#x60;: If a draft workflow scheme exists, it is updated. Otherwise, a draft workflow scheme is created.  *  Update an active workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;false&#x60;: An error is returned, as active workflow schemes cannot be updated.  *  Update an inactive workflow scheme with &#x60;updateDraftIfNeeded&#x60; set to &#x60;true&#x60;: The workflow scheme is updated, as inactive workflow schemes do not require drafts to update.  Defaults to &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

