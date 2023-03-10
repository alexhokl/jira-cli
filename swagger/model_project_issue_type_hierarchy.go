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

// The hierarchy of issue types within a project.
type ProjectIssueTypeHierarchy struct {
	// Details of an issue type hierarchy level.
	Hierarchy []ProjectIssueTypesHierarchyLevel `json:"hierarchy,omitempty"`
	// The ID of the project.
	ProjectId int64 `json:"projectId,omitempty"`
}
