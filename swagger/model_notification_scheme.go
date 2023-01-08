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

// Details about a notification scheme.
type NotificationScheme struct {
	// The description of the notification scheme.
	Description string `json:"description,omitempty"`
	// Expand options that include additional notification scheme details in the response.
	Expand string `json:"expand,omitempty"`
	// The ID of the notification scheme.
	Id int64 `json:"id,omitempty"`
	// The name of the notification scheme.
	Name string `json:"name,omitempty"`
	// The notification events and associated recipients.
	NotificationSchemeEvents []NotificationSchemeEvent `json:"notificationSchemeEvents,omitempty"`
	// The list of project IDs associated with the notification scheme.
	Projects []int64 `json:"projects,omitempty"`
	// The scope of the notification scheme.
	Scope *AllOfNotificationSchemeScope `json:"scope,omitempty"`
	Self string `json:"self,omitempty"`
}
