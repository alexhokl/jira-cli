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

// The field configuration to issue type mapping.
type FieldConfigurationToIssueTypeMapping struct {
	// The ID of the field configuration.
	FieldConfigurationId string `json:"fieldConfigurationId"`
	// The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration. An issue type can be included only once in a request.
	IssueTypeId string `json:"issueTypeId"`
}
