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

// The project and issue type mapping with a matching custom field context.
type ContextForProjectAndIssueType struct {
	// The ID of the custom field context.
	ContextId string `json:"contextId"`
	// The ID of the issue type.
	IssueTypeId string `json:"issueTypeId"`
	// The ID of the project.
	ProjectId string `json:"projectId"`
}
