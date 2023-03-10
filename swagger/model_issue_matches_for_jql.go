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

// A list of the issues matched to a JQL query or details of errors encountered during matching.
type IssueMatchesForJql struct {
	// A list of errors.
	Errors []string `json:"errors"`
	// A list of issue IDs.
	MatchedIssues []int64 `json:"matchedIssues"`
}
