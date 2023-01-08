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

// A project category.
type ProjectCategory struct {
	// The description of the project category.
	Description string `json:"description,omitempty"`
	// The ID of the project category.
	Id string `json:"id,omitempty"`
	// The name of the project category. Required on create, optional on update.
	Name string `json:"name,omitempty"`
	// The URL of the project category.
	Self string `json:"self,omitempty"`
}
