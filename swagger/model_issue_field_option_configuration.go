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

// Details of the projects the option is available in.
type IssueFieldOptionConfiguration struct {
	// DEPRECATED
	Attributes []string `json:"attributes,omitempty"`
	// Defines the projects that the option is available in. If the scope is not defined, then the option is available in all projects.
	Scope *AllOfIssueFieldOptionConfigurationScope `json:"scope,omitempty"`
}
