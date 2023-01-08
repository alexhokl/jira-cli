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

// Details of updates for a custom field.
type ConnectCustomFieldValues struct {
	// The list of custom field update details.
	UpdateValueList []ConnectCustomFieldValue `json:"updateValueList,omitempty"`
}
