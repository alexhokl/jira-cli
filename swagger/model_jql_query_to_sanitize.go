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

// The JQL query to sanitize for the account ID. If the account ID is null, sanitizing is performed for an anonymous user.
type JqlQueryToSanitize struct {
	// The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	AccountId string `json:"accountId,omitempty"`
	// The query to sanitize.
	Query string `json:"query"`
}
