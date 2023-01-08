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

// Details of an issue remote link.
type RemoteIssueLink struct {
	// Details of the remote application the linked item is in.
	Application *AllOfRemoteIssueLinkApplication `json:"application,omitempty"`
	// The global ID of the link, such as the ID of the item on the remote system.
	GlobalId string `json:"globalId,omitempty"`
	// The ID of the link.
	Id int64 `json:"id,omitempty"`
	// Details of the item linked to.
	Object *AllOfRemoteIssueLinkObject `json:"object,omitempty"`
	// Description of the relationship between the issue and the linked item.
	Relationship string `json:"relationship,omitempty"`
	// The URL of the link.
	Self string `json:"self,omitempty"`
}
