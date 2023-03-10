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

// Details of an issue resolution.
type Resolution struct {
	// The description of the issue resolution.
	Description string `json:"description,omitempty"`
	// The ID of the issue resolution.
	Id string `json:"id,omitempty"`
	// The name of the issue resolution.
	Name string `json:"name,omitempty"`
	// The URL of the issue resolution.
	Self string `json:"self,omitempty"`
}
