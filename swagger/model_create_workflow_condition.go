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

// A workflow transition condition.
type CreateWorkflowCondition struct {
	// The list of workflow conditions.
	Conditions []CreateWorkflowCondition `json:"conditions,omitempty"`
	// EXPERIMENTAL. The configuration of the transition rule.
	Configuration map[string]Object `json:"configuration,omitempty"`
	// The compound condition operator.
	Operator string `json:"operator,omitempty"`
	// The type of the transition rule.
	Type_ string `json:"type,omitempty"`
}