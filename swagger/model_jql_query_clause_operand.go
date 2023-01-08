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

// Details of an operand in a JQL clause.
type JqlQueryClauseOperand struct {
    ListOperand
    ValueOperand
    FunctionOperand
    KeywordOperand
}
