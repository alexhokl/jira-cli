# UpdateProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeType** | Pointer to **string** | The default assignee when creating issues for this project. | [optional] 
**AvatarId** | Pointer to **int64** | An integer value for the project&#39;s avatar. | [optional] 
**CategoryId** | Pointer to **int64** | The ID of the project&#39;s category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. To remove the project category from the project, set the value to &#x60;-1.&#x60; | [optional] 
**Description** | Pointer to **string** | A brief description of the project. | [optional] 
**IssueSecurityScheme** | Pointer to **int64** | The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) resource to get all issue security scheme IDs. | [optional] 
**Key** | Pointer to **string** | Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters. | [optional] 
**Lead** | Pointer to **string** | This parameter is deprecated because of privacy changes. Use &#x60;leadAccountId&#x60; instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. The user name of the project lead. Cannot be provided with &#x60;leadAccountId&#x60;. | [optional] 
**LeadAccountId** | Pointer to **string** | The account ID of the project lead. Cannot be provided with &#x60;lead&#x60;. | [optional] 
**Name** | Pointer to **string** | The name of the project. | [optional] 
**NotificationScheme** | Pointer to **int64** | The ID of the notification scheme for the project. Use the [Get notification schemes](#api-rest-api-3-notificationscheme-get) resource to get a list of notification scheme IDs. | [optional] 
**PermissionScheme** | Pointer to **int64** | The ID of the permission scheme for the project. Use the [Get all permission schemes](#api-rest-api-3-permissionscheme-get) resource to see a list of all permission scheme IDs. | [optional] 
**ReleasedProjectKeys** | Pointer to **[]string** | Previous project keys to be released from the current project. Released keys must belong to the current project and not contain the current project key | [optional] 
**Url** | Pointer to **string** | A link to information about this project, such as project documentation | [optional] 

## Methods

### NewUpdateProjectDetails

`func NewUpdateProjectDetails() *UpdateProjectDetails`

NewUpdateProjectDetails instantiates a new UpdateProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectDetailsWithDefaults

`func NewUpdateProjectDetailsWithDefaults() *UpdateProjectDetails`

NewUpdateProjectDetailsWithDefaults instantiates a new UpdateProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeType

`func (o *UpdateProjectDetails) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *UpdateProjectDetails) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *UpdateProjectDetails) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *UpdateProjectDetails) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetAvatarId

`func (o *UpdateProjectDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *UpdateProjectDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *UpdateProjectDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *UpdateProjectDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetCategoryId

`func (o *UpdateProjectDetails) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdateProjectDetails) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdateProjectDetails) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdateProjectDetails) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssueSecurityScheme

`func (o *UpdateProjectDetails) GetIssueSecurityScheme() int64`

GetIssueSecurityScheme returns the IssueSecurityScheme field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeOk

`func (o *UpdateProjectDetails) GetIssueSecuritySchemeOk() (*int64, bool)`

GetIssueSecuritySchemeOk returns a tuple with the IssueSecurityScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecurityScheme

`func (o *UpdateProjectDetails) SetIssueSecurityScheme(v int64)`

SetIssueSecurityScheme sets IssueSecurityScheme field to given value.

### HasIssueSecurityScheme

`func (o *UpdateProjectDetails) HasIssueSecurityScheme() bool`

HasIssueSecurityScheme returns a boolean if a field has been set.

### GetKey

`func (o *UpdateProjectDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateProjectDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateProjectDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateProjectDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLead

`func (o *UpdateProjectDetails) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *UpdateProjectDetails) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *UpdateProjectDetails) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *UpdateProjectDetails) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetLeadAccountId

`func (o *UpdateProjectDetails) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *UpdateProjectDetails) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *UpdateProjectDetails) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *UpdateProjectDetails) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateProjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProjectDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationScheme

`func (o *UpdateProjectDetails) GetNotificationScheme() int64`

GetNotificationScheme returns the NotificationScheme field if non-nil, zero value otherwise.

### GetNotificationSchemeOk

`func (o *UpdateProjectDetails) GetNotificationSchemeOk() (*int64, bool)`

GetNotificationSchemeOk returns a tuple with the NotificationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationScheme

`func (o *UpdateProjectDetails) SetNotificationScheme(v int64)`

SetNotificationScheme sets NotificationScheme field to given value.

### HasNotificationScheme

`func (o *UpdateProjectDetails) HasNotificationScheme() bool`

HasNotificationScheme returns a boolean if a field has been set.

### GetPermissionScheme

`func (o *UpdateProjectDetails) GetPermissionScheme() int64`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *UpdateProjectDetails) GetPermissionSchemeOk() (*int64, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *UpdateProjectDetails) SetPermissionScheme(v int64)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *UpdateProjectDetails) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### GetReleasedProjectKeys

`func (o *UpdateProjectDetails) GetReleasedProjectKeys() []string`

GetReleasedProjectKeys returns the ReleasedProjectKeys field if non-nil, zero value otherwise.

### GetReleasedProjectKeysOk

`func (o *UpdateProjectDetails) GetReleasedProjectKeysOk() (*[]string, bool)`

GetReleasedProjectKeysOk returns a tuple with the ReleasedProjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedProjectKeys

`func (o *UpdateProjectDetails) SetReleasedProjectKeys(v []string)`

SetReleasedProjectKeys sets ReleasedProjectKeys field to given value.

### HasReleasedProjectKeys

`func (o *UpdateProjectDetails) HasReleasedProjectKeys() bool`

HasReleasedProjectKeys returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateProjectDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateProjectDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateProjectDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateProjectDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


