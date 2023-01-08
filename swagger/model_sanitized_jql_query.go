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

// Details of the sanitized JQL query.
type SanitizedJqlQuery struct {
	// The account ID of the user for whom sanitization was performed.
	AccountId string `json:"accountId,omitempty"`
	// The list of errors.
	Errors *AllOfSanitizedJqlQueryErrors `json:"errors,omitempty"`
	// The initial query.
	InitialQuery string `json:"initialQuery,omitempty"`
	// The sanitized query, if there were no errors.
	SanitizedQuery string `json:"sanitizedQuery,omitempty"`
}