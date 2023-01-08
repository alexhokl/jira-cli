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

// Details of an application property.
type ApplicationProperty struct {
	// The allowed values, if applicable.
	AllowedValues []string `json:"allowedValues,omitempty"`
	// The default value of the application property.
	DefaultValue string `json:"defaultValue,omitempty"`
	// The description of the application property.
	Desc string `json:"desc,omitempty"`
	Example string `json:"example,omitempty"`
	// The ID of the application property. The ID and key are the same.
	Id string `json:"id,omitempty"`
	// The key of the application property. The ID and key are the same.
	Key string `json:"key,omitempty"`
	// The name of the application property.
	Name string `json:"name,omitempty"`
	// The data type of the application property.
	Type_ string `json:"type,omitempty"`
	// The new value.
	Value string `json:"value,omitempty"`
}
