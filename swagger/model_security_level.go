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

// Details of an issue level security item.
type SecurityLevel struct {
	// The description of the issue level security item.
	Description string `json:"description,omitempty"`
	// The ID of the issue level security item.
	Id string `json:"id,omitempty"`
	// Whether the issue level security item is the default.
	IsDefault bool `json:"isDefault,omitempty"`
	// The ID of the issue level security scheme.
	IssueSecuritySchemeId string `json:"issueSecuritySchemeId,omitempty"`
	// The name of the issue level security item.
	Name string `json:"name,omitempty"`
	// The URL of the issue level security item.
	Self string `json:"self,omitempty"`
}
