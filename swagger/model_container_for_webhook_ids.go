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

// Container for a list of webhook IDs.
type ContainerForWebhookIds struct {
	// A list of webhook IDs.
	WebhookIds []int64 `json:"webhookIds"`
}
