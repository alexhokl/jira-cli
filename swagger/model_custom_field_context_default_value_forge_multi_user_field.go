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

// Defaults for a Forge collection of users custom field.
type CustomFieldContextDefaultValueForgeMultiUserField struct {
	// The IDs of the default users.
	AccountIds []string `json:"accountIds"`
	// The ID of the context.
	ContextId string `json:"contextId"`
	Type_ string `json:"type"`
}
