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

// The history metadata associated with the changed.
type AllOfChangelogHistoryMetadata struct {
	// The activity described in the history record.
	ActivityDescription string `json:"activityDescription,omitempty"`
	// The key of the activity described in the history record.
	ActivityDescriptionKey string `json:"activityDescriptionKey,omitempty"`
	// Details of the user whose action created the history record.
	Actor *Object `json:"actor,omitempty"`
	// Details of the cause that triggered the creation the history record.
	Cause *Object `json:"cause,omitempty"`
	// The description of the history record.
	Description string `json:"description,omitempty"`
	// The description key of the history record.
	DescriptionKey string `json:"descriptionKey,omitempty"`
	// The description of the email address associated the history record.
	EmailDescription string `json:"emailDescription,omitempty"`
	// The description key of the email address associated the history record.
	EmailDescriptionKey string `json:"emailDescriptionKey,omitempty"`
	// Additional arbitrary information about the history record.
	ExtraData map[string]string `json:"extraData,omitempty"`
	// Details of the system that generated the history record.
	Generator *Object `json:"generator,omitempty"`
	// The type of the history record.
	Type_ string `json:"type,omitempty"`
}
