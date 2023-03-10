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

// Details of the options for a select list issue field.
type IssueFieldOption struct {
	Config *IssueFieldOptionConfiguration `json:"config,omitempty"`
	// The unique identifier for the option. This is only unique within the select field's set of options.
	Id int64 `json:"id"`
	// The properties of the object, as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see [Issue Field Option Property Index](https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/)) are defined in the descriptor for the issue field module.
	Properties map[string]Object `json:"properties,omitempty"`
	// The option's name, which is displayed in Jira.
	Value string `json:"value"`
}
