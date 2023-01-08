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

type IssueFieldOptionScopeBean struct {
	// Defines the behavior of the option within the global context. If this property is set, even if set to an empty object, then the option is available in all projects.
	Global *AllOfIssueFieldOptionScopeBeanGlobal `json:"global,omitempty"`
	// DEPRECATED
	Projects []int64 `json:"projects,omitempty"`
	// Defines the projects in which the option is available and the behavior of the option within each project. Specify one object per project. The behavior of the option in a project context overrides the behavior in the global context.
	Projects2 []ProjectScopeBean `json:"projects2,omitempty"`
}
