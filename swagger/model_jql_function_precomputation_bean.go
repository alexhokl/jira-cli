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
import (
	"time"
)

// Jql function precomputation.
type JqlFunctionPrecomputationBean struct {
	Arguments []string `json:"arguments,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Field string `json:"field,omitempty"`
	FunctionKey string `json:"functionKey,omitempty"`
	FunctionName string `json:"functionName,omitempty"`
	Id string `json:"id,omitempty"`
	Operator string `json:"operator,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Used time.Time `json:"used,omitempty"`
	Value string `json:"value,omitempty"`
}
