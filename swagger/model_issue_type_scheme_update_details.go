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

// Details of the name, description, and default issue type for an issue type scheme.
type IssueTypeSchemeUpdateDetails struct {
	// The ID of the default issue type of the issue type scheme.
	DefaultIssueTypeId string `json:"defaultIssueTypeId,omitempty"`
	// The description of the issue type scheme. The maximum length is 4000 characters.
	Description string `json:"description,omitempty"`
	// The name of the issue type scheme. The name must be unique. The maximum length is 255 characters.
	Name string `json:"name,omitempty"`
}
