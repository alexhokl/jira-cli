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

// Defaults for a Forge user custom field.
type CustomFieldContextDefaultValueForgeUserField struct {
	// The ID of the default user.
	AccountId string `json:"accountId"`
	// The ID of the context.
	ContextId string `json:"contextId"`
	Type_ string `json:"type"`
	UserFilter *UserFilter `json:"userFilter"`
}
