# CreateProjectDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeType** | **string** | The default assignee when creating issues for this project. | [optional] [default to null]
**AvatarId** | **int64** | An integer value for the project&#x27;s avatar. | [optional] [default to null]
**CategoryId** | **int64** | The ID of the project&#x27;s category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. | [optional] [default to null]
**Description** | **string** | A brief description of the project. | [optional] [default to null]
**FieldConfigurationScheme** | **int64** | The ID of the field configuration scheme for the project. Use the [Get all field configuration schemes](#api-rest-api-3-fieldconfigurationscheme-get) operation to get a list of field configuration scheme IDs. If you specify the field configuration scheme you cannot specify the project template key. | [optional] [default to null]
**IssueSecurityScheme** | **int64** | The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs. | [optional] [default to null]
**IssueTypeScheme** | **int64** | The ID of the issue type scheme for the project. Use the [Get all issue type schemes](#api-rest-api-3-issuetypescheme-get) operation to get a list of issue type scheme IDs. If you specify the issue type scheme you cannot specify the project template key. | [optional] [default to null]
**IssueTypeScreenScheme** | **int64** | The ID of the issue type screen scheme for the project. Use the [Get all issue type screen schemes](#api-rest-api-3-issuetypescreenscheme-get) operation to get a list of issue type screen scheme IDs. If you specify the issue type screen scheme you cannot specify the project template key. | [optional] [default to null]
**Key** | **string** | Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters. | [default to null]
**Lead** | **string** | This parameter is deprecated because of privacy changes. Use &#x60;leadAccountId&#x60; instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Either &#x60;lead&#x60; or &#x60;leadAccountId&#x60; must be set when creating a project. Cannot be provided with &#x60;leadAccountId&#x60;. | [optional] [default to null]
**LeadAccountId** | **string** | The account ID of the project lead. Either &#x60;lead&#x60; or &#x60;leadAccountId&#x60; must be set when creating a project. Cannot be provided with &#x60;lead&#x60;. | [optional] [default to null]
**Name** | **string** | The name of the project. | [default to null]
**NotificationScheme** | **int64** | The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs. | [optional] [default to null]
**PermissionScheme** | **int64** | The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs. | [optional] [default to null]
**ProjectTemplateKey** | **string** | A predefined configuration for a project. The type of the &#x60;projectTemplateKey&#x60; must match with the type of the &#x60;projectTypeKey&#x60;. | [optional] [default to null]
**ProjectTypeKey** | **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes), which defines the application-specific feature set. If you don&#x27;t specify the project template you have to specify the project type. | [optional] [default to null]
**Url** | **string** | A link to information about this project, such as project documentation | [optional] [default to null]
**WorkflowScheme** | **int64** | The ID of the workflow scheme for the project. Use the [Get all workflow schemes](#api-rest-api-3-workflowscheme-get) operation to get a list of workflow scheme IDs. If you specify the workflow scheme you cannot specify the project template key. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

