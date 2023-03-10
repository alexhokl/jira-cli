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

// A changelog.
type Changelog struct {
	// The user who made the change.
	Author *AllOfChangelogAuthor `json:"author,omitempty"`
	// The date on which the change took place.
	Created time.Time `json:"created,omitempty"`
	// The history metadata associated with the changed.
	HistoryMetadata *AllOfChangelogHistoryMetadata `json:"historyMetadata,omitempty"`
	// The ID of the changelog.
	Id string `json:"id,omitempty"`
	// The list of items changed.
	Items []ChangeDetails `json:"items,omitempty"`
}
