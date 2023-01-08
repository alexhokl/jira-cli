# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | Whether the project is archived. | [optional] [readonly] 
**ArchivedBy** | Pointer to [**User**](User.md) | The user who archived the project. | [optional] [readonly] 
**ArchivedDate** | Pointer to **time.Time** | The date when the project was archived. | [optional] [readonly] 
**AssigneeType** | Pointer to **string** | The default assignee when creating issues for this project. | [optional] [readonly] 
**AvatarUrls** | Pointer to [**AvatarUrlsBean**](AvatarUrlsBean.md) | The URLs of the project&#39;s avatars. | [optional] [readonly] 
**Components** | Pointer to [**[]ProjectComponent**](ProjectComponent.md) | List of the components contained in the project. | [optional] [readonly] 
**Deleted** | Pointer to **bool** | Whether the project is marked as deleted. | [optional] [readonly] 
**DeletedBy** | Pointer to [**User**](User.md) | The user who marked the project as deleted. | [optional] [readonly] 
**DeletedDate** | Pointer to **time.Time** | The date when the project was marked as deleted. | [optional] [readonly] 
**Description** | Pointer to **string** | A brief description of the project. | [optional] [readonly] 
**Email** | Pointer to **string** | An email address associated with the project. | [optional] 
**Expand** | Pointer to **string** | Expand options that include additional project details in the response. | [optional] [readonly] 
**Favourite** | Pointer to **bool** | Whether the project is selected as a favorite. | [optional] 
**Id** | Pointer to **string** | The ID of the project. | [optional] 
**Insight** | Pointer to [**ProjectInsight**](ProjectInsight.md) | Insights about the project. | [optional] [readonly] 
**IsPrivate** | Pointer to **bool** | Whether the project is private from the user&#39;s perspective. This means the user can&#39;t see the project or any associated issues. | [optional] [readonly] 
**IssueTypeHierarchy** | Pointer to [**Hierarchy**](Hierarchy.md) | The issue type hierarchy for the project. | [optional] [readonly] 
**IssueTypes** | Pointer to [**[]IssueTypeDetails**](IssueTypeDetails.md) | List of the issue types available in the project. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the project. | [optional] [readonly] 
**LandingPageInfo** | Pointer to [**ProjectLandingPageInfo**](ProjectLandingPageInfo.md) | The project landing page info. | [optional] [readonly] 
**Lead** | Pointer to [**User**](User.md) | The username of the project lead. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the project. | [optional] [readonly] 
**Permissions** | Pointer to [**ProjectPermissions**](ProjectPermissions.md) | User permissions on the project | [optional] [readonly] 
**ProjectCategory** | Pointer to [**ProjectCategory**](ProjectCategory.md) | The category the project belongs to. | [optional] [readonly] 
**ProjectTypeKey** | Pointer to **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes) of the project. | [optional] [readonly] 
**Properties** | Pointer to **map[string]interface{}** | Map of project properties | [optional] [readonly] 
**RetentionTillDate** | Pointer to **time.Time** | The date when the project is deleted permanently. | [optional] [readonly] 
**Roles** | Pointer to **map[string]string** | The name and self URL for each role defined in the project. For more information, see [Create project role](#api-rest-api-3-role-post). | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the project details. | [optional] [readonly] 
**Simplified** | Pointer to **bool** | Whether the project is simplified. | [optional] [readonly] 
**Style** | Pointer to **string** | The type of the project. | [optional] [readonly] 
**Url** | Pointer to **string** | A link to information about this project, such as project documentation. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Unique ID for next-gen projects. | [optional] [readonly] 
**Versions** | Pointer to [**[]Version**](Version.md) | The versions defined in the project. For more information, see [Create version](#api-rest-api-3-version-post). | [optional] [readonly] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *Project) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Project) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Project) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Project) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedBy

`func (o *Project) GetArchivedBy() User`

GetArchivedBy returns the ArchivedBy field if non-nil, zero value otherwise.

### GetArchivedByOk

`func (o *Project) GetArchivedByOk() (*User, bool)`

GetArchivedByOk returns a tuple with the ArchivedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedBy

`func (o *Project) SetArchivedBy(v User)`

SetArchivedBy sets ArchivedBy field to given value.

### HasArchivedBy

`func (o *Project) HasArchivedBy() bool`

HasArchivedBy returns a boolean if a field has been set.

### GetArchivedDate

`func (o *Project) GetArchivedDate() time.Time`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *Project) GetArchivedDateOk() (*time.Time, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *Project) SetArchivedDate(v time.Time)`

SetArchivedDate sets ArchivedDate field to given value.

### HasArchivedDate

`func (o *Project) HasArchivedDate() bool`

HasArchivedDate returns a boolean if a field has been set.

### GetAssigneeType

`func (o *Project) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *Project) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *Project) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *Project) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetAvatarUrls

`func (o *Project) GetAvatarUrls() AvatarUrlsBean`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *Project) GetAvatarUrlsOk() (*AvatarUrlsBean, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *Project) SetAvatarUrls(v AvatarUrlsBean)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *Project) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetComponents

`func (o *Project) GetComponents() []ProjectComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Project) GetComponentsOk() (*[]ProjectComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Project) SetComponents(v []ProjectComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Project) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDeleted

`func (o *Project) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Project) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Project) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Project) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Project) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Project) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Project) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Project) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedDate

`func (o *Project) GetDeletedDate() time.Time`

GetDeletedDate returns the DeletedDate field if non-nil, zero value otherwise.

### GetDeletedDateOk

`func (o *Project) GetDeletedDateOk() (*time.Time, bool)`

GetDeletedDateOk returns a tuple with the DeletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDate

`func (o *Project) SetDeletedDate(v time.Time)`

SetDeletedDate sets DeletedDate field to given value.

### HasDeletedDate

`func (o *Project) HasDeletedDate() bool`

HasDeletedDate returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *Project) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Project) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Project) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Project) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpand

`func (o *Project) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *Project) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *Project) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *Project) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFavourite

`func (o *Project) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *Project) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *Project) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *Project) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsight

`func (o *Project) GetInsight() ProjectInsight`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *Project) GetInsightOk() (*ProjectInsight, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *Project) SetInsight(v ProjectInsight)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *Project) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Project) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Project) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Project) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Project) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetIssueTypeHierarchy

`func (o *Project) GetIssueTypeHierarchy() Hierarchy`

GetIssueTypeHierarchy returns the IssueTypeHierarchy field if non-nil, zero value otherwise.

### GetIssueTypeHierarchyOk

`func (o *Project) GetIssueTypeHierarchyOk() (*Hierarchy, bool)`

GetIssueTypeHierarchyOk returns a tuple with the IssueTypeHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeHierarchy

`func (o *Project) SetIssueTypeHierarchy(v Hierarchy)`

SetIssueTypeHierarchy sets IssueTypeHierarchy field to given value.

### HasIssueTypeHierarchy

`func (o *Project) HasIssueTypeHierarchy() bool`

HasIssueTypeHierarchy returns a boolean if a field has been set.

### GetIssueTypes

`func (o *Project) GetIssueTypes() []IssueTypeDetails`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *Project) GetIssueTypesOk() (*[]IssueTypeDetails, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *Project) SetIssueTypes(v []IssueTypeDetails)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *Project) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetKey

`func (o *Project) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Project) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Project) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Project) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLandingPageInfo

`func (o *Project) GetLandingPageInfo() ProjectLandingPageInfo`

GetLandingPageInfo returns the LandingPageInfo field if non-nil, zero value otherwise.

### GetLandingPageInfoOk

`func (o *Project) GetLandingPageInfoOk() (*ProjectLandingPageInfo, bool)`

GetLandingPageInfoOk returns a tuple with the LandingPageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageInfo

`func (o *Project) SetLandingPageInfo(v ProjectLandingPageInfo)`

SetLandingPageInfo sets LandingPageInfo field to given value.

### HasLandingPageInfo

`func (o *Project) HasLandingPageInfo() bool`

HasLandingPageInfo returns a boolean if a field has been set.

### GetLead

`func (o *Project) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *Project) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *Project) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *Project) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *Project) GetPermissions() ProjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Project) GetPermissionsOk() (*ProjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Project) SetPermissions(v ProjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Project) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetProjectCategory

`func (o *Project) GetProjectCategory() ProjectCategory`

GetProjectCategory returns the ProjectCategory field if non-nil, zero value otherwise.

### GetProjectCategoryOk

`func (o *Project) GetProjectCategoryOk() (*ProjectCategory, bool)`

GetProjectCategoryOk returns a tuple with the ProjectCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCategory

`func (o *Project) SetProjectCategory(v ProjectCategory)`

SetProjectCategory sets ProjectCategory field to given value.

### HasProjectCategory

`func (o *Project) HasProjectCategory() bool`

HasProjectCategory returns a boolean if a field has been set.

### GetProjectTypeKey

`func (o *Project) GetProjectTypeKey() string`

GetProjectTypeKey returns the ProjectTypeKey field if non-nil, zero value otherwise.

### GetProjectTypeKeyOk

`func (o *Project) GetProjectTypeKeyOk() (*string, bool)`

GetProjectTypeKeyOk returns a tuple with the ProjectTypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeKey

`func (o *Project) SetProjectTypeKey(v string)`

SetProjectTypeKey sets ProjectTypeKey field to given value.

### HasProjectTypeKey

`func (o *Project) HasProjectTypeKey() bool`

HasProjectTypeKey returns a boolean if a field has been set.

### GetProperties

`func (o *Project) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Project) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Project) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Project) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRetentionTillDate

`func (o *Project) GetRetentionTillDate() time.Time`

GetRetentionTillDate returns the RetentionTillDate field if non-nil, zero value otherwise.

### GetRetentionTillDateOk

`func (o *Project) GetRetentionTillDateOk() (*time.Time, bool)`

GetRetentionTillDateOk returns a tuple with the RetentionTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTillDate

`func (o *Project) SetRetentionTillDate(v time.Time)`

SetRetentionTillDate sets RetentionTillDate field to given value.

### HasRetentionTillDate

`func (o *Project) HasRetentionTillDate() bool`

HasRetentionTillDate returns a boolean if a field has been set.

### GetRoles

`func (o *Project) GetRoles() map[string]string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Project) GetRolesOk() (*map[string]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Project) SetRoles(v map[string]string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Project) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSelf

`func (o *Project) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Project) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Project) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Project) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSimplified

`func (o *Project) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *Project) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *Project) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *Project) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.

### GetStyle

`func (o *Project) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *Project) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *Project) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *Project) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetUrl

`func (o *Project) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Project) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Project) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Project) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *Project) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Project) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Project) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Project) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersions

`func (o *Project) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Project) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Project) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Project) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


