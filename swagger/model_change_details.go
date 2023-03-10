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

// A change item.
type ChangeDetails struct {
	// The name of the field changed.
	Field string `json:"field,omitempty"`
	// The ID of the field changed.
	FieldId string `json:"fieldId,omitempty"`
	// The type of the field changed.
	Fieldtype string `json:"fieldtype,omitempty"`
	// The details of the original value.
	From string `json:"from,omitempty"`
	// The details of the original value as a string.
	FromString string `json:"fromString,omitempty"`
	// The details of the new value.
	To string `json:"to,omitempty"`
	// The details of the new value as a string.
	ToString string `json:"toString,omitempty"`
}
