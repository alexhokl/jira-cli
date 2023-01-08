# CreateProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeType** | Pointer to **string** | The default assignee when creating issues for this project. | [optional] 
**AvatarId** | Pointer to **int64** | An integer value for the project&#39;s avatar. | [optional] 
**CategoryId** | Pointer to **int64** | The ID of the project&#39;s category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. | [optional] 
**Description** | Pointer to **string** | A brief description of the project. | [optional] 
**FieldConfigurationScheme** | Pointer to **int64** | The ID of the field configuration scheme for the project. Use the [Get all field configuration schemes](#api-rest-api-3-fieldconfigurationscheme-get) operation to get a list of field configuration scheme IDs. If you specify the field configuration scheme you cannot specify the project template key. | [optional] 
**IssueSecurityScheme** | Pointer to **int64** | The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs. | [optional] 
**IssueTypeScheme** | Pointer to **int64** | The ID of the issue type scheme for the project. Use the [Get all issue type schemes](#api-rest-api-3-issuetypescheme-get) operation to get a list of issue type scheme IDs. If you specify the issue type scheme you cannot specify the project template key. | [optional] 
**IssueTypeScreenScheme** | Pointer to **int64** | The ID of the issue type screen scheme for the project. Use the [Get all issue type screen schemes](#api-rest-api-3-issuetypescreenscheme-get) operation to get a list of issue type screen scheme IDs. If you specify the issue type screen scheme you cannot specify the project template key. | [optional] 
**Key** | **string** | Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters. | 
**Lead** | Pointer to **string** | This parameter is deprecated because of privacy changes. Use &#x60;leadAccountId&#x60; instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Either &#x60;lead&#x60; or &#x60;leadAccountId&#x60; must be set when creating a project. Cannot be provided with &#x60;leadAccountId&#x60;. | [optional] 
**LeadAccountId** | Pointer to **string** | The account ID of the project lead. Either &#x60;lead&#x60; or &#x60;leadAccountId&#x60; must be set when creating a project. Cannot be provided with &#x60;lead&#x60;. | [optional] 
**Name** | **string** | The name of the project. | 
**NotificationScheme** | Pointer to **int64** | The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs. | [optional] 
**PermissionScheme** | Pointer to **int64** | The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs. | [optional] 
**ProjectTemplateKey** | Pointer to **string** | A predefined configuration for a project. The type of the &#x60;projectTemplateKey&#x60; must match with the type of the &#x60;projectTypeKey&#x60;. | [optional] 
**ProjectTypeKey** | Pointer to **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes), which defines the application-specific feature set. If you don&#39;t specify the project template you have to specify the project type. | [optional] 
**Url** | Pointer to **string** | A link to information about this project, such as project documentation | [optional] 
**WorkflowScheme** | Pointer to **int64** | The ID of the workflow scheme for the project. Use the [Get all workflow schemes](#api-rest-api-3-workflowscheme-get) operation to get a list of workflow scheme IDs. If you specify the workflow scheme you cannot specify the project template key. | [optional] 

## Methods

### NewCreateProjectDetails

`func NewCreateProjectDetails(key string, name string, ) *CreateProjectDetails`

NewCreateProjectDetails instantiates a new CreateProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectDetailsWithDefaults

`func NewCreateProjectDetailsWithDefaults() *CreateProjectDetails`

NewCreateProjectDetailsWithDefaults instantiates a new CreateProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeType

`func (o *CreateProjectDetails) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *CreateProjectDetails) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *CreateProjectDetails) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *CreateProjectDetails) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetAvatarId

`func (o *CreateProjectDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CreateProjectDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CreateProjectDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *CreateProjectDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetCategoryId

`func (o *CreateProjectDetails) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateProjectDetails) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateProjectDetails) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CreateProjectDetails) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateProjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldConfigurationScheme

`func (o *CreateProjectDetails) GetFieldConfigurationScheme() int64`

GetFieldConfigurationScheme returns the FieldConfigurationScheme field if non-nil, zero value otherwise.

### GetFieldConfigurationSchemeOk

`func (o *CreateProjectDetails) GetFieldConfigurationSchemeOk() (*int64, bool)`

GetFieldConfigurationSchemeOk returns a tuple with the FieldConfigurationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigurationScheme

`func (o *CreateProjectDetails) SetFieldConfigurationScheme(v int64)`

SetFieldConfigurationScheme sets FieldConfigurationScheme field to given value.

### HasFieldConfigurationScheme

`func (o *CreateProjectDetails) HasFieldConfigurationScheme() bool`

HasFieldConfigurationScheme returns a boolean if a field has been set.

### GetIssueSecurityScheme

`func (o *CreateProjectDetails) GetIssueSecurityScheme() int64`

GetIssueSecurityScheme returns the IssueSecurityScheme field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeOk

`func (o *CreateProjectDetails) GetIssueSecuritySchemeOk() (*int64, bool)`

GetIssueSecuritySchemeOk returns a tuple with the IssueSecurityScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecurityScheme

`func (o *CreateProjectDetails) SetIssueSecurityScheme(v int64)`

SetIssueSecurityScheme sets IssueSecurityScheme field to given value.

### HasIssueSecurityScheme

`func (o *CreateProjectDetails) HasIssueSecurityScheme() bool`

HasIssueSecurityScheme returns a boolean if a field has been set.

### GetIssueTypeScheme

`func (o *CreateProjectDetails) GetIssueTypeScheme() int64`

GetIssueTypeScheme returns the IssueTypeScheme field if non-nil, zero value otherwise.

### GetIssueTypeSchemeOk

`func (o *CreateProjectDetails) GetIssueTypeSchemeOk() (*int64, bool)`

GetIssueTypeSchemeOk returns a tuple with the IssueTypeScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScheme

`func (o *CreateProjectDetails) SetIssueTypeScheme(v int64)`

SetIssueTypeScheme sets IssueTypeScheme field to given value.

### HasIssueTypeScheme

`func (o *CreateProjectDetails) HasIssueTypeScheme() bool`

HasIssueTypeScheme returns a boolean if a field has been set.

### GetIssueTypeScreenScheme

`func (o *CreateProjectDetails) GetIssueTypeScreenScheme() int64`

GetIssueTypeScreenScheme returns the IssueTypeScreenScheme field if non-nil, zero value otherwise.

### GetIssueTypeScreenSchemeOk

`func (o *CreateProjectDetails) GetIssueTypeScreenSchemeOk() (*int64, bool)`

GetIssueTypeScreenSchemeOk returns a tuple with the IssueTypeScreenScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScreenScheme

`func (o *CreateProjectDetails) SetIssueTypeScreenScheme(v int64)`

SetIssueTypeScreenScheme sets IssueTypeScreenScheme field to given value.

### HasIssueTypeScreenScheme

`func (o *CreateProjectDetails) HasIssueTypeScreenScheme() bool`

HasIssueTypeScreenScheme returns a boolean if a field has been set.

### GetKey

`func (o *CreateProjectDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateProjectDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateProjectDetails) SetKey(v string)`

SetKey sets Key field to given value.


### GetLead

`func (o *CreateProjectDetails) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *CreateProjectDetails) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *CreateProjectDetails) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *CreateProjectDetails) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetLeadAccountId

`func (o *CreateProjectDetails) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *CreateProjectDetails) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *CreateProjectDetails) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *CreateProjectDetails) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetName

`func (o *CreateProjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationScheme

`func (o *CreateProjectDetails) GetNotificationScheme() int64`

GetNotificationScheme returns the NotificationScheme field if non-nil, zero value otherwise.

### GetNotificationSchemeOk

`func (o *CreateProjectDetails) GetNotificationSchemeOk() (*int64, bool)`

GetNotificationSchemeOk returns a tuple with the NotificationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationScheme

`func (o *CreateProjectDetails) SetNotificationScheme(v int64)`

SetNotificationScheme sets NotificationScheme field to given value.

### HasNotificationScheme

`func (o *CreateProjectDetails) HasNotificationScheme() bool`

HasNotificationScheme returns a boolean if a field has been set.

### GetPermissionScheme

`func (o *CreateProjectDetails) GetPermissionScheme() int64`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *CreateProjectDetails) GetPermissionSchemeOk() (*int64, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *CreateProjectDetails) SetPermissionScheme(v int64)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *CreateProjectDetails) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### GetProjectTemplateKey

`func (o *CreateProjectDetails) GetProjectTemplateKey() string`

GetProjectTemplateKey returns the ProjectTemplateKey field if non-nil, zero value otherwise.

### GetProjectTemplateKeyOk

`func (o *CreateProjectDetails) GetProjectTemplateKeyOk() (*string, bool)`

GetProjectTemplateKeyOk returns a tuple with the ProjectTemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplateKey

`func (o *CreateProjectDetails) SetProjectTemplateKey(v string)`

SetProjectTemplateKey sets ProjectTemplateKey field to given value.

### HasProjectTemplateKey

`func (o *CreateProjectDetails) HasProjectTemplateKey() bool`

HasProjectTemplateKey returns a boolean if a field has been set.

### GetProjectTypeKey

`func (o *CreateProjectDetails) GetProjectTypeKey() string`

GetProjectTypeKey returns the ProjectTypeKey field if non-nil, zero value otherwise.

### GetProjectTypeKeyOk

`func (o *CreateProjectDetails) GetProjectTypeKeyOk() (*string, bool)`

GetProjectTypeKeyOk returns a tuple with the ProjectTypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeKey

`func (o *CreateProjectDetails) SetProjectTypeKey(v string)`

SetProjectTypeKey sets ProjectTypeKey field to given value.

### HasProjectTypeKey

`func (o *CreateProjectDetails) HasProjectTypeKey() bool`

HasProjectTypeKey returns a boolean if a field has been set.

### GetUrl

`func (o *CreateProjectDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateProjectDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateProjectDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateProjectDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWorkflowScheme

`func (o *CreateProjectDetails) GetWorkflowScheme() int64`

GetWorkflowScheme returns the WorkflowScheme field if non-nil, zero value otherwise.

### GetWorkflowSchemeOk

`func (o *CreateProjectDetails) GetWorkflowSchemeOk() (*int64, bool)`

GetWorkflowSchemeOk returns a tuple with the WorkflowScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowScheme

`func (o *CreateProjectDetails) SetWorkflowScheme(v int64)`

SetWorkflowScheme sets WorkflowScheme field to given value.

### HasWorkflowScheme

`func (o *CreateProjectDetails) HasWorkflowScheme() bool`

HasWorkflowScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


