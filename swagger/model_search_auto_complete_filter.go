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

// Details of how to filter and list search auto complete information.
type SearchAutoCompleteFilter struct {
	// Include collapsed fields for fields that have non-unique names.
	IncludeCollapsedFields bool `json:"includeCollapsedFields,omitempty"`
	// List of project IDs used to filter the visible field details returned.
	ProjectIds []int64 `json:"projectIds,omitempty"`
}
