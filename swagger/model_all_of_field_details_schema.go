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

// The data schema for the field.
type AllOfFieldDetailsSchema struct {
	// If the field is a custom field, the configuration of the field.
	Configuration map[string]Object `json:"configuration,omitempty"`
	// If the field is a custom field, the URI of the field.
	Custom string `json:"custom,omitempty"`
	// If the field is a custom field, the custom ID of the field.
	CustomId int64 `json:"customId,omitempty"`
	// When the data type is an array, the name of the field items within the array.
	Items string `json:"items,omitempty"`
	// If the field is a system field, the name of the field.
	System string `json:"system,omitempty"`
	// The data type of the field.
	Type_ string `json:"type"`
}
