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

// Additional details about a project.
type ProjectInsight struct {
	// The last issue update time.
	LastIssueUpdateTime time.Time `json:"lastIssueUpdateTime,omitempty"`
	// Total issue count.
	TotalIssueCount int64 `json:"totalIssueCount,omitempty"`
}
