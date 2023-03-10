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

// Details about a notification event.
type NotificationEvent struct {
	// The description of the event.
	Description string `json:"description,omitempty"`
	// The ID of the event. The event can be a [Jira system event](https://confluence.atlassian.com/x/8YdKLg#Creatinganotificationscheme-eventsEvents) or a [custom event](https://confluence.atlassian.com/x/AIlKLg).
	Id int64 `json:"id,omitempty"`
	// The name of the event.
	Name string `json:"name,omitempty"`
	// The template of the event. Only custom events configured by Jira administrators have template.
	TemplateEvent *AllOfNotificationEventTemplateEvent `json:"templateEvent,omitempty"`
}
