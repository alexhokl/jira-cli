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

type SimplifiedHierarchyLevel struct {
	// The ID of the level above this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	AboveLevelId int64 `json:"aboveLevelId,omitempty"`
	// The ID of the level below this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	BelowLevelId int64 `json:"belowLevelId,omitempty"`
	// The external UUID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	ExternalUuid string `json:"externalUuid,omitempty"`
	HierarchyLevelNumber int32 `json:"hierarchyLevelNumber,omitempty"`
	// The ID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	Id int64 `json:"id,omitempty"`
	// The issue types available in this hierarchy level.
	IssueTypeIds []int64 `json:"issueTypeIds,omitempty"`
	// The level of this item in the hierarchy.
	Level int32 `json:"level,omitempty"`
	// The name of this hierarchy level.
	Name string `json:"name,omitempty"`
	// The ID of the project configuration. This property is deprecated, see [Change oticen: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/).
	ProjectConfigurationId int64 `json:"projectConfigurationId,omitempty"`
}
