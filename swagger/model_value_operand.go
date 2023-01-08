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

// An operand that is a user-provided value.
type ValueOperand struct {
	// Encoded value, which can be used directly in a JQL query.
	EncodedValue string `json:"encodedValue,omitempty"`
	// The operand value.
	Value string `json:"value"`
}
