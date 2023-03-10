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

// The date the refreshed webhooks expire.
type WebhooksExpirationDate struct {
	// The expiration date of all the refreshed webhooks.
	ExpirationDate int64 `json:"expirationDate"`
}
