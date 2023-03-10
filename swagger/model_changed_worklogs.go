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

// List of changed worklogs.
type ChangedWorklogs struct {
	LastPage bool `json:"lastPage,omitempty"`
	// The URL of the next list of changed worklogs.
	NextPage string `json:"nextPage,omitempty"`
	// The URL of this changed worklogs list.
	Self string `json:"self,omitempty"`
	// The datetime of the first worklog item in the list.
	Since int64 `json:"since,omitempty"`
	// The datetime of the last worklog item in the list.
	Until int64 `json:"until,omitempty"`
	// Changed worklog list.
	Values []ChangedWorklog `json:"values,omitempty"`
}
