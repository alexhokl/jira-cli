# Project

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Whether the project is archived. | [optional] [default to null]
**ArchivedBy** | [***AllOfProjectArchivedBy**](AllOfProjectArchivedBy.md) | The user who archived the project. | [optional] [default to null]
**ArchivedDate** | [**time.Time**](time.Time.md) | The date when the project was archived. | [optional] [default to null]
**AssigneeType** | **string** | The default assignee when creating issues for this project. | [optional] [default to null]
**AvatarUrls** | [***AllOfProjectAvatarUrls**](AllOfProjectAvatarUrls.md) | The URLs of the project&#x27;s avatars. | [optional] [default to null]
**Components** | [**[]ProjectComponent**](ProjectComponent.md) | List of the components contained in the project. | [optional] [default to null]
**Deleted** | **bool** | Whether the project is marked as deleted. | [optional] [default to null]
**DeletedBy** | [***AllOfProjectDeletedBy**](AllOfProjectDeletedBy.md) | The user who marked the project as deleted. | [optional] [default to null]
**DeletedDate** | [**time.Time**](time.Time.md) | The date when the project was marked as deleted. | [optional] [default to null]
**Description** | **string** | A brief description of the project. | [optional] [default to null]
**Email** | **string** | An email address associated with the project. | [optional] [default to null]
**Expand** | **string** | Expand options that include additional project details in the response. | [optional] [default to null]
**Favourite** | **bool** | Whether the project is selected as a favorite. | [optional] [default to null]
**Id** | **string** | The ID of the project. | [optional] [default to null]
**Insight** | [***AllOfProjectInsight**](AllOfProjectInsight.md) | Insights about the project. | [optional] [default to null]
**IsPrivate** | **bool** | Whether the project is private. | [optional] [default to null]
**IssueTypeHierarchy** | [***AllOfProjectIssueTypeHierarchy**](AllOfProjectIssueTypeHierarchy.md) | The issue type hierarchy for the project. | [optional] [default to null]
**IssueTypes** | [**[]IssueTypeDetails**](IssueTypeDetails.md) | List of the issue types available in the project. | [optional] [default to null]
**Key** | **string** | The key of the project. | [optional] [default to null]
**LandingPageInfo** | [***AllOfProjectLandingPageInfo**](AllOfProjectLandingPageInfo.md) | The project landing page info. | [optional] [default to null]
**Lead** | [***AllOfProjectLead**](AllOfProjectLead.md) | The username of the project lead. | [optional] [default to null]
**Name** | **string** | The name of the project. | [optional] [default to null]
**Permissions** | [***AllOfProjectPermissions**](AllOfProjectPermissions.md) | User permissions on the project | [optional] [default to null]
**ProjectCategory** | [***AllOfProjectProjectCategory**](AllOfProjectProjectCategory.md) | The category the project belongs to. | [optional] [default to null]
**ProjectTypeKey** | **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes) of the project. | [optional] [default to null]
**Properties** | [**map[string]Object**](.md) | Map of project properties | [optional] [default to null]
**RetentionTillDate** | [**time.Time**](time.Time.md) | The date when the project is deleted permanently. | [optional] [default to null]
**Roles** | **map[string]string** | The name and self URL for each role defined in the project. For more information, see [Create project role](#api-rest-api-3-role-post). | [optional] [default to null]
**Self** | **string** | The URL of the project details. | [optional] [default to null]
**Simplified** | **bool** | Whether the project is simplified. | [optional] [default to null]
**Style** | **string** | The type of the project. | [optional] [default to null]
**Url** | **string** | A link to information about this project, such as project documentation. | [optional] [default to null]
**Uuid** | **string** | Unique ID for next-gen projects. | [optional] [default to null]
**Versions** | [**[]Version**](Version.md) | The versions defined in the project. For more information, see [Create version](#api-rest-api-3-version-post). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

