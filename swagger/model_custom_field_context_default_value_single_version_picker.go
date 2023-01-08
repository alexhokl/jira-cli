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

// The default value for a version picker custom field.
type CustomFieldContextDefaultValueSingleVersionPicker struct {
	Type_ string `json:"type"`
	// The ID of the default version.
	VersionId string `json:"versionId"`
	// The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are `\"releasedFirst\"` and `\"unreleasedFirst\"`.
	VersionOrder string `json:"versionOrder,omitempty"`
}
