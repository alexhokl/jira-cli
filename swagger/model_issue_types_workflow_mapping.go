/*
 * The Jira Cloud platform REST API
 *
 * Jira Cloud platform REST API documentation
 *
 * API version: 1001.0.0-SNAPSHOT
 * Contact: ecosystem@atlassian.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details about the mapping between issue types and a workflow.
type IssueTypesWorkflowMapping struct {
	// Whether the workflow is the default workflow for the workflow scheme.
	DefaultMapping bool `json:"defaultMapping,omitempty"`
	// The list of issue type IDs.
	IssueTypes []string `json:"issueTypes,omitempty"`
	// Whether a draft workflow scheme is created or updated when updating an active workflow scheme. The draft is updated with the new workflow-issue types mapping. Defaults to `false`.
	UpdateDraftIfNeeded bool `json:"updateDraftIfNeeded,omitempty"`
	// The name of the workflow. Optional if updating the workflow-issue types mapping.
	Workflow string `json:"workflow,omitempty"`
}
