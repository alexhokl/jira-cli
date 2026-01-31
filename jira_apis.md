# Jira Cloud REST API v3 - Available Endpoints

This document summarizes the API endpoints available in the auto-generated
Swagger clients.

> **Note:** This document covers two API clients:
> - **[Jira Cloud Platform REST API v3](#jira-cloud-rest-api-v3---available-endpoints)** (`swagger/` directory)
> - **[Jira Software REST API](#jira-software-agile-rest-api---available-endpoints)** (`swagger_software/` directory)

---

## Table of Contents (Platform API)

- [Issues](#issues)
- [Issue Search](#issue-search)
- [Issue Comments](#issue-comments)
- [Issue Attachments](#issue-attachments)
- [Issue Links](#issue-links)
- [Issue Worklogs](#issue-worklogs)
- [Issue Watchers & Votes](#issue-watchers--votes)
- [Issue Properties](#issue-properties)
- [Issue Types](#issue-types)
- [Issue Fields](#issue-fields)
- [Issue Priorities](#issue-priorities)
- [Issue Resolutions](#issue-resolutions)
- [Issue Security](#issue-security)
- [Projects](#projects)
- [Project Components](#project-components)
- [Project Versions](#project-versions)
- [Project Roles](#project-roles)
- [Users](#users)
- [User Search](#user-search)
- [Groups](#groups)
- [Filters](#filters)
- [Dashboards](#dashboards)
- [Workflows](#workflows)
- [Workflow Schemes](#workflow-schemes)
- [Screens](#screens)
- [Permissions](#permissions)
- [JQL](#jql)
- [Labels](#labels)
- [Server & Settings](#server--settings)
- [Other APIs](#other-apis)

---

## Issues

**Service:** `IssuesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssue` | Get issue details |
| `CreateIssue` | Create a new issue |
| `CreateIssues` | Bulk create issues |
| `EditIssue` | Edit an existing issue |
| `DeleteIssue` | Delete an issue |
| `AssignIssue` | Assign an issue to a user |
| `DoTransition` | Transition an issue to a new status |
| `GetTransitions` | Get available transitions for an issue |
| `GetChangeLogs` | Get issue changelog |
| `GetChangeLogsByIds` | Get changelogs by IDs |
| `GetBulkChangelogs` | Bulk get changelogs |
| `BulkFetchIssues` | Bulk fetch multiple issues |
| `GetCreateIssueMeta` | Get metadata for creating issues |
| `GetCreateIssueMetaIssueTypes` | Get issue types for create metadata |
| `GetCreateIssueMetaIssueTypeId` | Get metadata for specific issue type |
| `GetEditIssueMeta` | Get metadata for editing an issue |
| `GetEvents` | Get issue events |
| `GetIssueLimitReport` | Get issue limit report |
| `Notify` | Send notification for an issue |
| `ArchiveIssues` | Archive issues |
| `ArchiveIssuesAsync` | Archive issues asynchronously |
| `UnarchiveIssues` | Unarchive issues |
| `ExportArchivedIssues` | Export archived issues |

**Service:** `IssueBulkOperationsAPIService`

| Method | Description |
|--------|-------------|
| `GetAvailableTransitions` | Get available transitions for bulk operation |
| `GetBulkEditableFields` | Get fields editable in bulk |
| `GetBulkOperationProgress` | Get bulk operation progress |
| `SubmitBulkDelete` | Bulk delete issues |
| `SubmitBulkEdit` | Bulk edit issues |
| `SubmitBulkMove` | Bulk move issues |
| `SubmitBulkTransition` | Bulk transition issues |
| `SubmitBulkWatch` | Bulk watch issues |
| `SubmitBulkUnwatch` | Bulk unwatch issues |

---

## Issue Search

**Service:** `IssueSearchAPIService`

| Method | Description |
|--------|-------------|
| `SearchForIssuesUsingJql` | Search issues using JQL (GET) |
| `SearchForIssuesUsingJqlPost` | Search issues using JQL (POST) |
| `SearchAndReconsileIssuesUsingJql` | Search and reconcile issues (GET) |
| `SearchAndReconsileIssuesUsingJqlPost` | Search and reconcile issues (POST) |
| `CountIssues` | Count issues matching JQL |
| `GetIssuePickerResource` | Get issue picker suggestions |
| `MatchIssues` | Match issues against JQL |

---

## Issue Comments

**Service:** `IssueCommentsAPIService`

| Method | Description |
|--------|-------------|
| `GetComments` | Get comments for an issue |
| `GetComment` | Get a specific comment |
| `GetCommentsByIds` | Get comments by IDs |
| `AddComment` | Add a comment to an issue |
| `UpdateComment` | Update a comment |
| `DeleteComment` | Delete a comment |

**Service:** `IssueCommentPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetCommentPropertyKeys` | Get comment property keys |
| `GetCommentProperty` | Get a comment property |
| `SetCommentProperty` | Set a comment property |
| `DeleteCommentProperty` | Delete a comment property |

---

## Issue Attachments

**Service:** `IssueAttachmentsAPIService`

| Method | Description |
|--------|-------------|
| `AddAttachment` | Add attachment to an issue |
| `GetAttachment` | Get attachment metadata |
| `GetAttachmentContent` | Get attachment content |
| `GetAttachmentThumbnail` | Get attachment thumbnail |
| `GetAttachmentMeta` | Get attachment settings |
| `RemoveAttachment` | Remove an attachment |
| `ExpandAttachmentForHumans` | Expand attachment for humans |
| `ExpandAttachmentForMachines` | Expand attachment for machines |

---

## Issue Links

**Service:** `IssueLinksAPIService`

| Method | Description |
|--------|-------------|
| `LinkIssues` | Create an issue link |
| `GetIssueLink` | Get an issue link |
| `DeleteIssueLink` | Delete an issue link |

**Service:** `IssueLinkTypesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueLinkTypes` | Get all issue link types |
| `GetIssueLinkType` | Get an issue link type |
| `CreateIssueLinkType` | Create an issue link type |
| `UpdateIssueLinkType` | Update an issue link type |
| `DeleteIssueLinkType` | Delete an issue link type |

**Service:** `IssueRemoteLinksAPIService`

| Method | Description |
|--------|-------------|
| `GetRemoteIssueLinks` | Get remote issue links |
| `GetRemoteIssueLinkById` | Get a remote issue link |
| `CreateOrUpdateRemoteIssueLink` | Create or update remote link |
| `UpdateRemoteIssueLink` | Update a remote issue link |
| `DeleteRemoteIssueLinkById` | Delete remote link by ID |
| `DeleteRemoteIssueLinkByGlobalId` | Delete remote link by global ID |

---

## Issue Worklogs

**Service:** `IssueWorklogsAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueWorklog` | Get worklogs for an issue |
| `GetWorklog` | Get a specific worklog |
| `GetWorklogsForIds` | Get worklogs by IDs |
| `AddWorklog` | Add a worklog to an issue |
| `UpdateWorklog` | Update a worklog |
| `DeleteWorklog` | Delete a worklog |
| `BulkDeleteWorklogs` | Bulk delete worklogs |
| `BulkMoveWorklogs` | Bulk move worklogs |
| `GetIdsOfWorklogsModifiedSince` | Get modified worklog IDs |
| `GetIdsOfWorklogsDeletedSince` | Get deleted worklog IDs |

**Service:** `IssueWorklogPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetWorklogPropertyKeys` | Get worklog property keys |
| `GetWorklogProperty` | Get a worklog property |
| `SetWorklogProperty` | Set a worklog property |
| `DeleteWorklogProperty` | Delete a worklog property |

---

## Issue Watchers & Votes

**Service:** `IssueWatchersAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueWatchers` | Get issue watchers |
| `AddWatcher` | Add a watcher |
| `RemoveWatcher` | Remove a watcher |
| `GetIsWatchingIssueBulk` | Bulk check if watching |

**Service:** `IssueVotesAPIService`

| Method | Description |
|--------|-------------|
| `GetVotes` | Get votes for an issue |
| `AddVote` | Add a vote |
| `RemoveVote` | Remove a vote |

---

## Issue Properties

**Service:** `IssuePropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssuePropertyKeys` | Get issue property keys |
| `GetIssueProperty` | Get an issue property |
| `SetIssueProperty` | Set an issue property |
| `DeleteIssueProperty` | Delete an issue property |
| `BulkSetIssueProperty` | Bulk set issue property |
| `BulkSetIssuePropertiesByIssue` | Bulk set properties by issue |
| `BulkSetIssuesPropertiesList` | Bulk set issues properties list |
| `BulkDeleteIssueProperty` | Bulk delete issue property |

---

## Issue Types

**Service:** `IssueTypesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueAllTypes` | Get all issue types |
| `GetIssueType` | Get an issue type |
| `GetIssueTypesForProject` | Get issue types for a project |
| `GetAlternativeIssueTypes` | Get alternative issue types |
| `CreateIssueType` | Create an issue type |
| `UpdateIssueType` | Update an issue type |
| `DeleteIssueType` | Delete an issue type |
| `CreateIssueTypeAvatar` | Create issue type avatar |

**Service:** `IssueTypeSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllIssueTypeSchemes` | Get all issue type schemes |
| `GetIssueTypeSchemeForProjects` | Get schemes for projects |
| `GetIssueTypeSchemesMapping` | Get scheme mappings |
| `CreateIssueTypeScheme` | Create issue type scheme |
| `UpdateIssueTypeScheme` | Update issue type scheme |
| `DeleteIssueTypeScheme` | Delete issue type scheme |
| `AddIssueTypesToIssueTypeScheme` | Add issue types to scheme |
| `RemoveIssueTypeFromIssueTypeScheme` | Remove issue type from scheme |
| `ReorderIssueTypesInIssueTypeScheme` | Reorder issue types in scheme |
| `AssignIssueTypeSchemeToProject` | Assign scheme to project |

**Service:** `IssueTypePropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueTypePropertyKeys` | Get issue type property keys |
| `GetIssueTypeProperty` | Get issue type property |
| `SetIssueTypeProperty` | Set issue type property |
| `DeleteIssueTypeProperty` | Delete issue type property |

---

## Issue Fields

**Service:** `IssueFieldsAPIService`

| Method | Description |
|--------|-------------|
| `GetFields` | Get all fields |
| `GetFieldsPaginated` | Get fields paginated |
| `GetProjectFields` | Get fields for a project |
| `GetTrashedFieldsPaginated` | Get trashed fields |
| `CreateCustomField` | Create a custom field |
| `UpdateCustomField` | Update a custom field |
| `DeleteCustomField` | Delete a custom field |
| `TrashCustomField` | Trash a custom field |
| `RestoreCustomField` | Restore a custom field |
| `GetContextsForFieldDeprecated` | Get contexts for field (deprecated) |

**Service:** `IssueCustomFieldContextsAPIService`

| Method | Description |
|--------|-------------|
| `GetContextsForField` | Get contexts for a field |
| `CreateCustomFieldContext` | Create custom field context |
| `UpdateCustomFieldContext` | Update custom field context |
| `DeleteCustomFieldContext` | Delete custom field context |
| `GetDefaultValues` | Get default values |
| `SetDefaultValues` | Set default values |
| `GetIssueTypeMappingsForContexts` | Get issue type mappings |
| `AddIssueTypesToContext` | Add issue types to context |
| `RemoveIssueTypesFromContext` | Remove issue types from context |
| `GetProjectContextMapping` | Get project context mapping |
| `AssignProjectsToCustomFieldContext` | Assign projects to context |
| `RemoveCustomFieldContextFromProjects` | Remove context from projects |
| `GetCustomFieldContextsForProjectsAndIssueTypes` | Get contexts for projects/issue types |

**Service:** `IssueCustomFieldOptionsAPIService`

| Method | Description |
|--------|-------------|
| `GetOptionsForContext` | Get options for context |
| `GetCustomFieldOption` | Get a custom field option |
| `CreateCustomFieldOption` | Create custom field option |
| `UpdateCustomFieldOption` | Update custom field option |
| `DeleteCustomFieldOption` | Delete custom field option |
| `ReorderCustomFieldOptions` | Reorder options |
| `ReplaceCustomFieldOption` | Replace an option |

**Service:** `IssueFieldConfigurationsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllFieldConfigurations` | Get all field configurations |
| `GetAllFieldConfigurationSchemes` | Get all field config schemes |
| `CreateFieldConfiguration` | Create field configuration |
| `UpdateFieldConfiguration` | Update field configuration |
| `DeleteFieldConfiguration` | Delete field configuration |
| `GetFieldConfigurationItems` | Get field configuration items |
| `UpdateFieldConfigurationItems` | Update field config items |
| `CreateFieldConfigurationScheme` | Create field config scheme |
| `UpdateFieldConfigurationScheme` | Update field config scheme |
| `DeleteFieldConfigurationScheme` | Delete field config scheme |
| `GetFieldConfigurationSchemeMappings` | Get scheme mappings |
| `SetFieldConfigurationSchemeMapping` | Set scheme mapping |
| `GetFieldConfigurationSchemeProjectMapping` | Get project mapping |
| `AssignFieldConfigurationSchemeToProject` | Assign scheme to project |
| `RemoveIssueTypesFromGlobalFieldConfigurationScheme` | Remove issue types from global scheme |

---

## Issue Priorities

**Service:** `IssuePrioritiesAPIService`

| Method | Description |
|--------|-------------|
| `GetPriorities` | Get all priorities |
| `GetPriority` | Get a priority |
| `SearchPriorities` | Search priorities |
| `CreatePriority` | Create a priority |
| `UpdatePriority` | Update a priority |
| `DeletePriority` | Delete a priority |
| `MovePriorities` | Move priorities |
| `SetDefaultPriority` | Set default priority |

**Service:** `PrioritySchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetPrioritySchemes` | Get priority schemes |
| `CreatePriorityScheme` | Create priority scheme |
| `UpdatePriorityScheme` | Update priority scheme |
| `DeletePriorityScheme` | Delete priority scheme |
| `GetPrioritiesByPriorityScheme` | Get priorities by scheme |
| `GetAvailablePrioritiesByPriorityScheme` | Get available priorities |
| `GetProjectsByPriorityScheme` | Get projects by scheme |
| `SuggestedPrioritiesForMappings` | Get suggested priorities |

---

## Issue Resolutions

**Service:** `IssueResolutionsAPIService`

| Method | Description |
|--------|-------------|
| `GetResolutions` | Get all resolutions |
| `GetResolution` | Get a resolution |
| `SearchResolutions` | Search resolutions |
| `CreateResolution` | Create a resolution |
| `UpdateResolution` | Update a resolution |
| `DeleteResolution` | Delete a resolution |
| `MoveResolutions` | Move resolutions |
| `SetDefaultResolution` | Set default resolution |

---

## Issue Security

**Service:** `IssueSecurityLevelAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueSecurityLevel` | Get security level |
| `GetIssueSecurityLevelMembers` | Get security level members |

**Service:** `IssueSecuritySchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueSecuritySchemes` | Get all security schemes |
| `GetIssueSecurityScheme` | Get a security scheme |
| `CreateIssueSecurityScheme` | Create security scheme |
| `UpdateIssueSecurityScheme` | Update security scheme |
| `DeleteSecurityScheme` | Delete security scheme |
| `GetSecurityLevels` | Get security levels |
| `AddSecurityLevel` | Add security level |
| `UpdateSecurityLevel` | Update security level |
| `RemoveLevel` | Remove security level |
| `GetSecurityLevelMembers` | Get level members |
| `AddSecurityLevelMembers` | Add level members |
| `RemoveMemberFromSecurityLevel` | Remove level member |
| `SetDefaultLevels` | Set default levels |
| `SearchSecuritySchemes` | Search security schemes |
| `SearchProjectsUsingSecuritySchemes` | Search projects using schemes |
| `AssociateSchemesToProjects` | Associate schemes to projects |

---

## Projects

**Service:** `ProjectsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllProjects` | Get all projects |
| `GetProject` | Get a project |
| `SearchProjects` | Search projects |
| `GetRecent` | Get recent projects |
| `CreateProject` | Create a project |
| `UpdateProject` | Update a project |
| `DeleteProject` | Delete a project |
| `DeleteProjectAsynchronously` | Delete project async |
| `ArchiveProject` | Archive a project |
| `Restore` | Restore a project |
| `GetAllStatuses` | Get all statuses for project |
| `GetHierarchy` | Get project hierarchy |
| `GetNotificationSchemeForProject` | Get notification scheme |

**Service:** `ProjectCategoriesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllProjectCategories` | Get all project categories |
| `GetProjectCategoryById` | Get a project category |
| `CreateProjectCategory` | Create project category |
| `UpdateProjectCategory` | Update project category |
| `RemoveProjectCategory` | Remove project category |

**Service:** `ProjectTypesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllProjectTypes` | Get all project types |
| `GetAllAccessibleProjectTypes` | Get accessible project types |
| `GetProjectTypeByKey` | Get project type by key |
| `GetAccessibleProjectTypeByKey` | Get accessible type by key |

**Service:** `ProjectAvatarsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllProjectAvatars` | Get project avatars |
| `CreateProjectAvatar` | Create project avatar |
| `UpdateProjectAvatar` | Update project avatar |
| `DeleteProjectAvatar` | Delete project avatar |

**Service:** `ProjectPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetProjectPropertyKeys` | Get property keys |
| `GetProjectProperty` | Get a property |
| `SetProjectProperty` | Set a property |
| `DeleteProjectProperty` | Delete a property |

**Service:** `ProjectFeaturesAPIService`

| Method | Description |
|--------|-------------|
| `GetFeaturesForProject` | Get project features |
| `ToggleFeatureForProject` | Toggle a feature |

**Service:** `ProjectEmailAPIService`

| Method | Description |
|--------|-------------|
| `GetProjectEmail` | Get project email |
| `UpdateProjectEmail` | Update project email |

**Service:** `ProjectKeyAndNameValidationAPIService`

| Method | Description |
|--------|-------------|
| `ValidateProjectKey` | Validate project key |
| `GetValidProjectKey` | Get valid project key |
| `GetValidProjectName` | Get valid project name |

**Service:** `ProjectPermissionSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetAssignedPermissionScheme` | Get assigned permission scheme |
| `AssignPermissionScheme` | Assign permission scheme |
| `GetProjectIssueSecurityScheme` | Get issue security scheme |
| `GetSecurityLevelsForProject` | Get security levels |

**Service:** `ProjectClassificationLevelsAPIService`

| Method | Description |
|--------|-------------|
| `GetDefaultProjectClassification` | Get default classification |
| `UpdateDefaultProjectClassification` | Update default classification |
| `RemoveDefaultProjectClassification` | Remove default classification |

**Service:** `ProjectTemplatesAPIService`

| Method | Description |
|--------|-------------|
| `CreateProjectWithCustomTemplate` | Create project with template |
| `SaveTemplate` | Save template |
| `EditTemplate` | Edit template |
| `RemoveTemplate` | Remove template |
| `LiveTemplate` | Live template |

---

## Project Components

**Service:** `ProjectComponentsAPIService`

| Method | Description |
|--------|-------------|
| `GetProjectComponents` | Get project components |
| `GetProjectComponentsPaginated` | Get components paginated |
| `GetComponent` | Get a component |
| `CreateComponent` | Create a component |
| `UpdateComponent` | Update a component |
| `DeleteComponent` | Delete a component |
| `GetComponentRelatedIssues` | Get related issues |
| `FindComponentsForProjects` | Find components for projects |

---

## Project Versions

**Service:** `ProjectVersionsAPIService`

| Method | Description |
|--------|-------------|
| `GetProjectVersions` | Get project versions |
| `GetProjectVersionsPaginated` | Get versions paginated |
| `GetVersion` | Get a version |
| `CreateVersion` | Create a version |
| `UpdateVersion` | Update a version |
| `DeleteVersion` | Delete a version |
| `MoveVersion` | Move a version |
| `MergeVersions` | Merge versions |
| `DeleteAndReplaceVersion` | Delete and replace version |
| `GetVersionRelatedIssues` | Get related issues |
| `GetVersionUnresolvedIssues` | Get unresolved issues |
| `GetRelatedWork` | Get related work |
| `CreateRelatedWork` | Create related work |
| `UpdateRelatedWork` | Update related work |
| `DeleteRelatedWork` | Delete related work |

---

## Project Roles

**Service:** `ProjectRolesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllProjectRoles` | Get all project roles |
| `GetProjectRoleById` | Get project role by ID |
| `GetProjectRoles` | Get roles for a project |
| `GetProjectRole` | Get a project role |
| `GetProjectRoleDetails` | Get role details |
| `CreateProjectRole` | Create a project role |
| `FullyUpdateProjectRole` | Fully update role |
| `PartialUpdateProjectRole` | Partially update role |
| `DeleteProjectRole` | Delete a project role |

**Service:** `ProjectRoleActorsAPIService`

| Method | Description |
|--------|-------------|
| `GetProjectRoleActorsForRole` | Get actors for role |
| `AddActorUsers` | Add actor users |
| `DeleteActor` | Delete an actor |
| `SetActors` | Set actors |
| `AddProjectRoleActorsToRole` | Add actors to role |
| `DeleteProjectRoleActorsFromRole` | Delete actors from role |

---

## Users

**Service:** `UsersAPIService`

| Method | Description |
|--------|-------------|
| `GetUser` | Get a user |
| `GetAllUsers` | Get all users |
| `GetAllUsersDefault` | Get all users (default) |
| `BulkGetUsers` | Bulk get users |
| `BulkGetUsersMigration` | Bulk get users migration |
| `CreateUser` | Create a user |
| `RemoveUser` | Remove a user |
| `GetUserGroups` | Get user groups |
| `GetUserEmail` | Get user email |
| `GetUserEmailBulk` | Bulk get user emails |
| `GetUserDefaultColumns` | Get default columns |
| `SetUserColumns` | Set user columns |
| `ResetUserColumns` | Reset user columns |

**Service:** `UserPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetUserPropertyKeys` | Get property keys |
| `GetUserProperty` | Get a property |
| `SetUserProperty` | Set a property |
| `DeleteUserProperty` | Delete a property |

**Service:** `MyselfAPIService`

| Method | Description |
|--------|-------------|
| `GetCurrentUser` | Get current user |
| `GetLocale` | Get locale |
| `SetLocale` | Set locale |
| `GetPreference` | Get preference |
| `SetPreference` | Set preference |
| `RemovePreference` | Remove preference |

---

## User Search

**Service:** `UserSearchAPIService`

| Method | Description |
|--------|-------------|
| `FindUsers` | Find users |
| `FindUsersForPicker` | Find users for picker |
| `FindUsersByQuery` | Find users by query |
| `FindUserKeysByQuery` | Find user keys by query |
| `FindAssignableUsers` | Find assignable users |
| `FindBulkAssignableUsers` | Bulk find assignable users |
| `FindUsersWithAllPermissions` | Find users with permissions |
| `FindUsersWithBrowsePermission` | Find users with browse permission |

**Service:** `GroupAndUserPickerAPIService`

| Method | Description |
|--------|-------------|
| `FindUsersAndGroups` | Find users and groups |

---

## Groups

**Service:** `GroupsAPIService`

| Method | Description |
|--------|-------------|
| `GetGroup` | Get a group |
| `FindGroups` | Find groups |
| `BulkGetGroups` | Bulk get groups |
| `CreateGroup` | Create a group |
| `RemoveGroup` | Remove a group |
| `GetUsersFromGroup` | Get users from group |
| `AddUserToGroup` | Add user to group |
| `RemoveUserFromGroup` | Remove user from group |

---

## Filters

**Service:** `FiltersAPIService`

| Method | Description |
|--------|-------------|
| `GetFilter` | Get a filter |
| `GetFiltersPaginated` | Get filters paginated |
| `GetMyFilters` | Get my filters |
| `GetFavouriteFilters` | Get favourite filters |
| `CreateFilter` | Create a filter |
| `UpdateFilter` | Update a filter |
| `DeleteFilter` | Delete a filter |
| `SetFavouriteForFilter` | Set as favourite |
| `DeleteFavouriteForFilter` | Remove from favourites |
| `ChangeFilterOwner` | Change filter owner |
| `GetColumns` | Get filter columns |
| `SetColumns` | Set filter columns |
| `ResetColumns` | Reset filter columns |

**Service:** `FilterSharingAPIService`

| Method | Description |
|--------|-------------|
| `GetSharePermissions` | Get share permissions |
| `GetSharePermission` | Get a share permission |
| `AddSharePermission` | Add share permission |
| `DeleteSharePermission` | Delete share permission |
| `GetDefaultShareScope` | Get default share scope |
| `SetDefaultShareScope` | Set default share scope |

---

## Dashboards

**Service:** `DashboardsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllDashboards` | Get all dashboards |
| `GetDashboardsPaginated` | Get dashboards paginated |
| `GetDashboard` | Get a dashboard |
| `CreateDashboard` | Create a dashboard |
| `UpdateDashboard` | Update a dashboard |
| `DeleteDashboard` | Delete a dashboard |
| `CopyDashboard` | Copy a dashboard |
| `BulkEditDashboards` | Bulk edit dashboards |
| `GetAllGadgets` | Get all gadgets |
| `GetAllAvailableDashboardGadgets` | Get available gadgets |
| `AddGadget` | Add a gadget |
| `UpdateGadget` | Update a gadget |
| `RemoveGadget` | Remove a gadget |
| `GetDashboardItemPropertyKeys` | Get item property keys |
| `GetDashboardItemProperty` | Get item property |
| `SetDashboardItemProperty` | Set item property |
| `DeleteDashboardItemProperty` | Delete item property |

---

## Workflows

**Service:** `WorkflowsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllWorkflows` | Get all workflows |
| `GetWorkflowsPaginated` | Get workflows paginated |
| `SearchWorkflows` | Search workflows |
| `ReadWorkflows` | Read workflows |
| `CreateWorkflow` | Create a workflow |
| `CreateWorkflows` | Create multiple workflows |
| `UpdateWorkflows` | Update workflows |
| `DeleteInactiveWorkflow` | Delete inactive workflow |
| `ValidateCreateWorkflows` | Validate create workflows |
| `ValidateUpdateWorkflows` | Validate update workflows |
| `GetDefaultEditor` | Get default editor |
| `WorkflowCapabilities` | Get workflow capabilities |
| `ListWorkflowHistory` | List workflow history |
| `ReadWorkflowFromHistory` | Read workflow from history |
| `ReadWorkflowPreviews` | Read workflow previews |
| `GetProjectUsagesForWorkflow` | Get project usages |
| `GetWorkflowSchemeUsagesForWorkflow` | Get scheme usages |
| `GetWorkflowProjectIssueTypeUsages` | Get issue type usages |

**Service:** `WorkflowStatusesAPIService`

| Method | Description |
|--------|-------------|
| `GetStatuses` | Get all statuses |
| `GetStatus` | Get a status |

**Service:** `WorkflowStatusCategoriesAPIService`

| Method | Description |
|--------|-------------|
| `GetStatusCategories` | Get status categories |
| `GetStatusCategory` | Get a status category |

**Service:** `WorkflowTransitionPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `GetWorkflowTransitionProperties` | Get transition properties |
| `CreateWorkflowTransitionProperty` | Create transition property |
| `UpdateWorkflowTransitionProperty` | Update transition property |
| `DeleteWorkflowTransitionProperty` | Delete transition property |

**Service:** `WorkflowTransitionRulesAPIService`

| Method | Description |
|--------|-------------|
| `GetWorkflowTransitionRuleConfigurations` | Get rule configurations |
| `UpdateWorkflowTransitionRuleConfigurations` | Update rule configs |
| `DeleteWorkflowTransitionRuleConfigurations` | Delete rule configs |

---

## Workflow Schemes

**Service:** `WorkflowSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllWorkflowSchemes` | Get all workflow schemes |
| `GetWorkflowScheme` | Get a workflow scheme |
| `ReadWorkflowSchemes` | Read workflow schemes |
| `CreateWorkflowScheme` | Create workflow scheme |
| `UpdateWorkflowScheme` | Update workflow scheme |
| `DeleteWorkflowScheme` | Delete workflow scheme |
| `UpdateSchemes` | Update schemes |
| `GetDefaultWorkflow` | Get default workflow |
| `UpdateDefaultWorkflow` | Update default workflow |
| `DeleteDefaultWorkflow` | Delete default workflow |
| `GetWorkflow` | Get workflow |
| `UpdateWorkflowMapping` | Update workflow mapping |
| `DeleteWorkflowMapping` | Delete workflow mapping |
| `GetWorkflowSchemeIssueType` | Get scheme issue type |
| `SetWorkflowSchemeIssueType` | Set scheme issue type |
| `DeleteWorkflowSchemeIssueType` | Delete scheme issue type |
| `GetRequiredWorkflowSchemeMappings` | Get required mappings |
| `GetProjectUsagesForWorkflowScheme` | Get project usages |
| `SwitchWorkflowSchemeForProject` | Switch scheme for project |

**Service:** `WorkflowSchemeDraftsAPIService`

| Method | Description |
|--------|-------------|
| `GetWorkflowSchemeDraft` | Get scheme draft |
| `CreateWorkflowSchemeDraftFromParent` | Create draft from parent |
| `UpdateWorkflowSchemeDraft` | Update scheme draft |
| `DeleteWorkflowSchemeDraft` | Delete scheme draft |
| `PublishDraftWorkflowScheme` | Publish draft |
| `GetDraftDefaultWorkflow` | Get draft default workflow |
| `UpdateDraftDefaultWorkflow` | Update draft default |
| `DeleteDraftDefaultWorkflow` | Delete draft default |
| `GetDraftWorkflow` | Get draft workflow |
| `UpdateDraftWorkflowMapping` | Update draft mapping |
| `DeleteDraftWorkflowMapping` | Delete draft mapping |
| `GetWorkflowSchemeDraftIssueType` | Get draft issue type |
| `SetWorkflowSchemeDraftIssueType` | Set draft issue type |
| `DeleteWorkflowSchemeDraftIssueType` | Delete draft issue type |

**Service:** `WorkflowSchemeProjectAssociationsAPIService`

| Method | Description |
|--------|-------------|
| `GetWorkflowSchemeProjectAssociations` | Get project associations |
| `AssignSchemeToProject` | Assign scheme to project |

---

## Screens

**Service:** `ScreensAPIService`

| Method | Description |
|--------|-------------|
| `GetScreens` | Get screens |
| `GetScreensForField` | Get screens for field |
| `GetAvailableScreenFields` | Get available fields |
| `CreateScreen` | Create a screen |
| `UpdateScreen` | Update a screen |
| `DeleteScreen` | Delete a screen |
| `AddFieldToDefaultScreen` | Add field to default screen |

**Service:** `ScreenSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetScreenSchemes` | Get screen schemes |
| `CreateScreenScheme` | Create screen scheme |
| `UpdateScreenScheme` | Update screen scheme |
| `DeleteScreenScheme` | Delete screen scheme |

**Service:** `ScreenTabsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllScreenTabs` | Get all screen tabs |
| `GetBulkScreenTabs` | Bulk get screen tabs |
| `AddScreenTab` | Add a screen tab |
| `RenameScreenTab` | Rename a screen tab |
| `DeleteScreenTab` | Delete a screen tab |
| `MoveScreenTab` | Move a screen tab |

**Service:** `ScreenTabFieldsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllScreenTabFields` | Get all tab fields |
| `AddScreenTabField` | Add tab field |
| `RemoveScreenTabField` | Remove tab field |
| `MoveScreenTabField` | Move tab field |

**Service:** `IssueTypeScreenSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueTypeScreenSchemes` | Get issue type screen schemes |
| `GetIssueTypeScreenSchemeMappings` | Get scheme mappings |
| `GetIssueTypeScreenSchemeProjectAssociations` | Get project associations |
| `GetProjectsForIssueTypeScreenScheme` | Get projects for scheme |
| `CreateIssueTypeScreenScheme` | Create scheme |
| `UpdateIssueTypeScreenScheme` | Update scheme |
| `DeleteIssueTypeScreenScheme` | Delete scheme |
| `AppendMappingsForIssueTypeScreenScheme` | Append mappings |
| `RemoveMappingsFromIssueTypeScreenScheme` | Remove mappings |
| `UpdateDefaultScreenScheme` | Update default scheme |
| `AssignIssueTypeScreenSchemeToProject` | Assign to project |

---

## Permissions

**Service:** `PermissionsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllPermissions` | Get all permissions |
| `GetMyPermissions` | Get my permissions |
| `GetBulkPermissions` | Bulk get permissions |
| `GetPermittedProjects` | Get permitted projects |

**Service:** `PermissionSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllPermissionSchemes` | Get all permission schemes |
| `GetPermissionScheme` | Get a permission scheme |
| `CreatePermissionScheme` | Create permission scheme |
| `UpdatePermissionScheme` | Update permission scheme |
| `DeletePermissionScheme` | Delete permission scheme |
| `GetPermissionSchemeGrants` | Get scheme grants |
| `GetPermissionSchemeGrant` | Get a scheme grant |
| `CreatePermissionGrant` | Create permission grant |
| `DeletePermissionSchemeEntity` | Delete scheme entity |

---

## JQL

**Service:** `JQLAPIService`

| Method | Description |
|--------|-------------|
| `GetAutoComplete` | Get JQL autocomplete suggestions |
| `GetAutoCompletePost` | Get autocomplete (POST) |
| `GetFieldAutoCompleteForQueryString` | Get field autocomplete |
| `ParseJqlQueries` | Parse JQL queries |
| `SanitiseJqlQueries` | Sanitise JQL queries |
| `MigrateQueries` | Migrate queries |

**Service:** `JiraExpressionsAPIService`

| Method | Description |
|--------|-------------|
| `AnalyseExpression` | Analyse expression |
| `EvaluateJiraExpression` | Evaluate Jira expression |
| `EvaluateJSISJiraExpression` | Evaluate JSIS expression |

---

## Labels

**Service:** `LabelsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllLabels` | Get all labels |

---

## Server & Settings

**Service:** `ServerInfoAPIService`

| Method | Description |
|--------|-------------|
| `GetServerInfo` | Get server information |

**Service:** `JiraSettingsAPIService`

| Method | Description |
|--------|-------------|
| `GetConfiguration` | Get configuration |
| `GetAdvancedSettings` | Get advanced settings |
| `GetApplicationProperty` | Get application property |
| `SetApplicationProperty` | Set application property |

**Service:** `StatusAPIService`

| Method | Description |
|--------|-------------|
| `Search` | Search statuses |
| `GetStatusesById` | Get statuses by ID |
| `GetStatusesByName` | Get statuses by name |
| `CreateStatuses` | Create statuses |
| `UpdateStatuses` | Update statuses |
| `DeleteStatusesById` | Delete statuses by ID |
| `GetProjectUsagesForStatus` | Get project usages |
| `GetProjectIssueTypeUsagesForStatus` | Get issue type usages |
| `GetWorkflowUsagesForStatus` | Get workflow usages |

---

## Other APIs

### Announcement Banner

**Service:** `AnnouncementBannerAPIService`

| Method | Description |
|--------|-------------|
| `GetBanner` | Get announcement banner |
| `SetBanner` | Set announcement banner |

### Application Roles

**Service:** `ApplicationRolesAPIService`

| Method | Description |
|--------|-------------|
| `GetAllApplicationRoles` | Get all application roles |
| `GetApplicationRole` | Get an application role |

### Audit Records

**Service:** `AuditRecordsAPIService`

| Method | Description |
|--------|-------------|
| `GetAuditRecords` | Get audit records |

### Avatars

**Service:** `AvatarsAPIService`

| Method | Description |
|--------|-------------|
| `GetAvatars` | Get avatars |
| `GetAllSystemAvatars` | Get system avatars |
| `GetAvatarImageByID` | Get avatar image by ID |
| `GetAvatarImageByOwner` | Get avatar by owner |
| `GetAvatarImageByType` | Get avatar by type |
| `StoreAvatar` | Store an avatar |
| `DeleteAvatar` | Delete an avatar |

### Classification Levels

**Service:** `ClassificationLevelsAPIService`

| Method | Description |
|--------|-------------|
| `GetAllUserDataClassificationLevels` | Get classification levels |

### Issue Notification Schemes

**Service:** `IssueNotificationSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetNotificationSchemes` | Get notification schemes |
| `GetNotificationScheme` | Get a notification scheme |
| `CreateNotificationScheme` | Create notification scheme |
| `UpdateNotificationScheme` | Update notification scheme |
| `DeleteNotificationScheme` | Delete notification scheme |
| `AddNotifications` | Add notifications |
| `RemoveNotificationFromNotificationScheme` | Remove notification |
| `GetNotificationSchemeToProjectMappings` | Get project mappings |

### Issue Navigator Settings

**Service:** `IssueNavigatorSettingsAPIService`

| Method | Description |
|--------|-------------|
| `GetIssueNavigatorDefaultColumns` | Get default columns |
| `SetIssueNavigatorDefaultColumns` | Set default columns |

### License Metrics

**Service:** `LicenseMetricsAPIService`

| Method | Description |
|--------|-------------|
| `GetLicense` | Get license |
| `GetApproximateLicenseCount` | Get approximate license count |
| `GetApproximateApplicationLicenseCount` | Get app license count |

### Plans

**Service:** `PlansAPIService`

| Method | Description |
|--------|-------------|
| `GetPlans` | Get plans |
| `GetPlan` | Get a plan |
| `CreatePlan` | Create a plan |
| `UpdatePlan` | Update a plan |
| `DuplicatePlan` | Duplicate a plan |
| `ArchivePlan` | Archive a plan |
| `TrashPlan` | Trash a plan |

### Teams in Plan

**Service:** `TeamsInPlanAPIService`

| Method | Description |
|--------|-------------|
| `GetTeams` | Get teams |
| `GetAtlassianTeam` | Get Atlassian team |
| `AddAtlassianTeam` | Add Atlassian team |
| `UpdateAtlassianTeam` | Update Atlassian team |
| `RemoveAtlassianTeam` | Remove Atlassian team |
| `GetPlanOnlyTeam` | Get plan-only team |
| `CreatePlanOnlyTeam` | Create plan-only team |
| `UpdatePlanOnlyTeam` | Update plan-only team |
| `DeletePlanOnlyTeam` | Delete plan-only team |

### Tasks

**Service:** `TasksAPIService`

| Method | Description |
|--------|-------------|
| `GetTask` | Get a task |
| `CancelTask` | Cancel a task |

### Time Tracking

**Service:** `TimeTrackingAPIService`

| Method | Description |
|--------|-------------|
| `GetAvailableTimeTrackingImplementations` | Get implementations |
| `GetSelectedTimeTrackingImplementation` | Get selected implementation |
| `SelectTimeTrackingImplementation` | Select implementation |
| `GetSharedTimeTrackingConfiguration` | Get shared config |
| `SetSharedTimeTrackingConfiguration` | Set shared config |

### Webhooks

**Service:** `WebhooksAPIService`

| Method | Description |
|--------|-------------|
| `GetDynamicWebhooksForApp` | Get webhooks for app |
| `RegisterDynamicWebhooks` | Register webhooks |
| `DeleteWebhookById` | Delete webhook |
| `GetFailedWebhooks` | Get failed webhooks |
| `RefreshWebhooks` | Refresh webhooks |

### App Properties

**Service:** `AppPropertiesAPIService`

| Method | Description |
|--------|-------------|
| `AddonPropertiesResourceGetAddonPropertiesGet` | Get addon properties |
| `AddonPropertiesResourceGetAddonPropertyGet` | Get addon property |
| `AddonPropertiesResourcePutAddonPropertyPut` | Put addon property |
| `AddonPropertiesResourceDeleteAddonPropertyDelete` | Delete addon property |
| `GetForgeAppPropertyKeys` | Get Forge app keys |
| `GetForgeAppProperty` | Get Forge app property |
| `PutForgeAppProperty` | Put Forge app property |
| `DeleteForgeAppProperty` | Delete Forge app property |

### Field Schemes

**Service:** `FieldSchemesAPIService`

| Method | Description |
|--------|-------------|
| `GetFieldAssociationSchemes` | Get field association schemes |
| `GetFieldAssociationSchemeById` | Get scheme by ID |
| `CreateFieldAssociationScheme` | Create scheme |
| `UpdateFieldAssociationScheme` | Update scheme |
| `DeleteFieldAssociationScheme` | Delete scheme |
| `SearchFieldAssociationSchemeFields` | Search scheme fields |
| `SearchFieldAssociationSchemeProjects` | Search scheme projects |
| `GetProjectsWithFieldSchemes` | Get projects with schemes |
| `AssociateProjectsToFieldAssociationSchemes` | Associate projects |
| `RemoveFieldsAssociatedWithSchemes` | Remove fields |
| `UpdateFieldsAssociatedWithSchemes` | Update fields |
| `GetFieldAssociationSchemeItemParameters` | Get item parameters |
| `UpdateFieldAssociationSchemeItemParameters` | Update item parameters |
| `RemoveFieldAssociationSchemeItemParameters` | Remove item parameters |

### Issue Redaction

**Service:** `IssueRedactionAPIService`

| Method | Description |
|--------|-------------|
| `Redact` | Redact issue content |
| `GetRedactionStatus` | Get redaction status |

### App Data Policies

**Service:** `AppDataPoliciesAPIService`

| Method | Description |
|--------|-------------|
| `GetPolicies` | Get policies |
| `GetPolicy` | Get a policy |

---

# Jira Software (Agile) REST API - Available Endpoints

This section summarizes the API endpoints available in the auto-generated
Swagger client (`swagger_software/` directory).

> **Note:** This client is generated from the **Jira Software REST API**
> (Agile API). It provides access to boards, sprints, backlogs, epics, and
> DevOps-related features.

---

## Table of Contents (Software API)

- [Backlog](#backlog)
- [Board](#board)
- [Sprint](#sprint)
- [Epic](#epic)
- [Issue (Agile)](#issue-agile)
- [Builds](#builds)
- [Deployments](#deployments)
- [Development Information](#development-information)
- [Feature Flags](#feature-flags)
- [Remote Links](#remote-links)
- [Operations](#operations)
- [Security Information](#security-information)
- [DevOps Components](#devops-components)

---

## Backlog

**Service:** `BacklogAPIService`

| Method | Description |
|--------|-------------|
| `MoveIssuesToBacklog` | Move issues to the backlog |
| `MoveIssuesToBacklogForBoard` | Move issues to backlog for a specific board |

---

## Board

**Service:** `BoardAPIService`

| Method | Description |
|--------|-------------|
| `CreateBoard` | Create a new board |
| `DeleteBoard` | Delete a board |
| `GetAllBoards` | Get all boards |
| `GetBoard` | Get a specific board |
| `GetBoardByFilterId` | Get board by filter ID |
| `GetConfiguration` | Get board configuration |
| `GetFeaturesForBoard` | Get features for a board |
| `ToggleFeatures` | Toggle board features |
| `GetAllSprints` | Get all sprints for a board |
| `GetAllVersions` | Get all versions for a board |
| `GetEpics` | Get epics for a board |
| `GetIssuesForBoard` | Get issues for a board |
| `GetIssuesForBacklog` | Get backlog issues for a board |
| `GetIssuesWithoutEpicForBoard` | Get issues without epic for a board |
| `GetBoardIssuesForEpic` | Get board issues for an epic |
| `GetBoardIssuesForSprint` | Get board issues for a sprint |
| `MoveIssuesToBoard` | Move issues to a board |
| `GetProjects` | Get projects for a board |
| `GetProjectsFull` | Get full project details for a board |
| `GetAllQuickFilters` | Get all quick filters for a board |
| `GetQuickFilter` | Get a specific quick filter |
| `GetReportsForBoard` | Get reports for a board |
| `GetBoardPropertyKeys` | Get board property keys |
| `GetBoardProperty` | Get a board property |
| `SetBoardProperty` | Set a board property |
| `DeleteBoardProperty` | Delete a board property |

---

## Sprint

**Service:** `SprintAPIService`

| Method | Description |
|--------|-------------|
| `CreateSprint` | Create a new sprint |
| `GetSprint` | Get a sprint |
| `UpdateSprint` | Fully update a sprint |
| `PartiallyUpdateSprint` | Partially update a sprint |
| `DeleteSprint` | Delete a sprint |
| `GetIssuesForSprint` | Get issues for a sprint |
| `MoveIssuesToSprintAndRank` | Move issues to sprint and rank them |
| `SwapSprint` | Swap sprint with another |
| `GetPropertiesKeys` | Get sprint property keys |
| `GetProperty` | Get a sprint property |
| `SetProperty` | Set a sprint property |
| `DeleteProperty` | Delete a sprint property |

---

## Epic

**Service:** `EpicAPIService`

| Method | Description |
|--------|-------------|
| `GetEpic` | Get an epic |
| `PartiallyUpdateEpic` | Partially update an epic |
| `GetIssuesForEpic` | Get issues for an epic |
| `GetIssuesWithoutEpic` | Get issues without an epic |
| `MoveIssuesToEpic` | Move issues to an epic |
| `RemoveIssuesFromEpic` | Remove issues from an epic |
| `RankEpics` | Rank epics |

---

## Issue (Agile)

**Service:** `IssueAPIService`

| Method | Description |
|--------|-------------|
| `GetIssue` | Get agile issue details |
| `RankIssues` | Rank issues |
| `GetIssueEstimationForBoard` | Get issue estimation for a board |
| `EstimateIssueForBoard` | Set issue estimation for a board |

---

## Builds

**Service:** `BuildsAPIService`

| Method | Description |
|--------|-------------|
| `SubmitBuilds` | Submit build information |
| `GetBuildByKey` | Get build by key |
| `DeleteBuildByKey` | Delete build by key |
| `DeleteBuildsByProperty` | Delete builds by property |

---

## Deployments

**Service:** `DeploymentsAPIService`

| Method | Description |
|--------|-------------|
| `SubmitDeployments` | Submit deployment information |
| `GetDeploymentByKey` | Get deployment by key |
| `GetDeploymentGatingStatusByKey` | Get deployment gating status |
| `DeleteDeploymentByKey` | Delete deployment by key |
| `DeleteDeploymentsByProperty` | Delete deployments by property |

---

## Development Information

**Service:** `DevelopmentInformationAPIService`

| Method | Description |
|--------|-------------|
| `StoreDevelopmentInformation` | Store development information |
| `GetRepository` | Get repository information |
| `DeleteRepository` | Delete repository |
| `DeleteEntity` | Delete a development entity |
| `DeleteByProperties` | Delete by properties |
| `ExistsByProperties` | Check if data exists by properties |

---

## Feature Flags

**Service:** `FeatureFlagsAPIService`

| Method | Description |
|--------|-------------|
| `SubmitFeatureFlags` | Submit feature flag information |
| `GetFeatureFlagById` | Get feature flag by ID |
| `DeleteFeatureFlagById` | Delete feature flag by ID |
| `DeleteFeatureFlagsByProperty` | Delete feature flags by property |

---

## Remote Links

**Service:** `RemoteLinksAPIService`

| Method | Description |
|--------|-------------|
| `SubmitRemoteLinks` | Submit remote links |
| `GetRemoteLinkById` | Get remote link by ID |
| `DeleteRemoteLinkById` | Delete remote link by ID |
| `DeleteRemoteLinksByProperty` | Delete remote links by property |

---

## Operations

**Service:** `OperationsAPIService`

| Method | Description |
|--------|-------------|
| `SubmitEntity` | Submit incidents/reviews |
| `GetIncidentById` | Get incident by ID |
| `DeleteIncidentById` | Delete incident by ID |
| `GetReviewById` | Get review by ID |
| `DeleteReviewById` | Delete review by ID |
| `DeleteEntityByProperty` | Delete entities by property |
| `SubmitOperationsWorkspaces` | Submit operations workspaces |
| `GetWorkspaces` | Get operations workspaces |
| `DeleteWorkspaces` | Delete operations workspaces |

---

## Security Information

**Service:** `SecurityInformationAPIService`

| Method | Description |
|--------|-------------|
| `SubmitVulnerabilities` | Submit vulnerability information |
| `GetVulnerabilityById` | Get vulnerability by ID |
| `DeleteVulnerabilityById` | Delete vulnerability by ID |
| `DeleteVulnerabilitiesByProperty` | Delete vulnerabilities by property |
| `SubmitWorkspaces` | Submit security workspaces |
| `GetLinkedWorkspaces` | Get linked workspaces |
| `GetLinkedWorkspaceById` | Get linked workspace by ID |
| `DeleteLinkedWorkspaces` | Delete linked workspaces |

---

## DevOps Components

**Service:** `DevOpsComponentsAPIService`

| Method | Description |
|--------|-------------|
| `SubmitComponents` | Submit DevOps components |
| `GetComponentById` | Get component by ID |
| `DeleteComponentById` | Delete component by ID |
| `DeleteComponentsByProperty` | Delete components by property |

---

## Notes

1. **Authentication**: All API calls require authentication via the
   `getAuthContext()` helper function in `cmd/helper.go`.

2. **Client Creation**: Use `newClient()` from `cmd/helper.go` to create the
   Platform API client. For Software API, create a separate client using the
   `swagger_software` package.

3. **Available APIs**:
   - **Jira Cloud Platform REST API v3** (`swagger/`): Issues, projects,
     workflows, users, permissions, and more.
   - **Jira Software REST API** (`swagger_software/`): Boards, sprints,
     backlogs, epics, builds, deployments, and DevOps features.

4. **Not Included**:
   - **Jira Service Management API**: Service desk, queues, SLAs

5. **Usage Example (Platform API)**:
   ```go
   client := newClient()
   ctx := getAuthContext()
   issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJECT-123").Execute()
   ```

6. **Usage Example (Software API)**:
   ```go
   import "github.com/alexhokl/jira-cli/swagger_software"

   cfg := swagger_software.NewConfiguration()
   client := swagger_software.NewAPIClient(cfg)
   boards, _, err := client.BoardAPI.GetAllBoards(ctx).Execute()
   ```
