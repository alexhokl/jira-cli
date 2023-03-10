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

// The details of a UI modification's context, which define where to activate the UI modification.
type UiModificationContextDetails struct {
	// The ID of the UI modification context.
	Id string `json:"id,omitempty"`
	// Whether a context is available. For example, when a project is deleted the context becomes unavailable.
	IsAvailable bool `json:"isAvailable,omitempty"`
	// The issue type ID of the context.
	IssueTypeId string `json:"issueTypeId"`
	// The project ID of the context.
	ProjectId string `json:"projectId"`
	// The view type of the context. Only `GIC` (Global Issue Create) is supported.
	ViewType string `json:"viewType"`
}
