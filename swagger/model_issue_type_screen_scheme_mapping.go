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

// The IDs of the screen schemes for the issue type IDs.
type IssueTypeScreenSchemeMapping struct {
	// The ID of the issue type or *default*. Only issue types used in classic projects are accepted. An entry for *default* must be provided and defines the mapping for all issue types without a screen scheme.
	IssueTypeId string `json:"issueTypeId"`
	// The ID of the screen scheme. Only screen schemes used in classic projects are accepted.
	ScreenSchemeId string `json:"screenSchemeId"`
}
