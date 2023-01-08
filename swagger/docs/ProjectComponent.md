# ProjectComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ari** | Pointer to **string** | Compass component&#39;s ID. Can&#39;t be updated. Not required for creating a Project Component. | [optional] [readonly] 
**Assignee** | Pointer to [**User**](User.md) | The details of the user associated with &#x60;assigneeType&#x60;, if any. See &#x60;realAssignee&#x60; for details of the user assigned to issues created with this component. | [optional] [readonly] 
**AssigneeType** | Pointer to **string** | The nominal user type used to determine the assignee for issues created with this component. See &#x60;realAssigneeType&#x60; for details on how the type of the user, and hence the user, assigned to issues is determined. Can take the following values:   *  &#x60;PROJECT_LEAD&#x60; the assignee to any issues created with this component is nominally the lead for the project the component is in.  *  &#x60;COMPONENT_LEAD&#x60; the assignee to any issues created with this component is nominally the lead for the component.  *  &#x60;UNASSIGNED&#x60; an assignee is not set for issues created with this component.  *  &#x60;PROJECT_DEFAULT&#x60; the assignee to any issues created with this component is nominally the default assignee for the project that the component is in.  Default value: &#x60;PROJECT_DEFAULT&#x60;.   Optional when creating or updating a component. | [optional] 
**Description** | Pointer to **string** | The description for the component. Optional when creating or updating a component. | [optional] 
**Id** | Pointer to **string** | The unique identifier for the component. | [optional] [readonly] 
**IsAssigneeTypeValid** | Pointer to **bool** | Whether a user is associated with &#x60;assigneeType&#x60;. For example, if the &#x60;assigneeType&#x60; is set to &#x60;COMPONENT_LEAD&#x60; but the component lead is not set, then &#x60;false&#x60; is returned. | [optional] [readonly] 
**Lead** | Pointer to [**User**](User.md) | The user details for the component&#39;s lead user. | [optional] [readonly] 
**LeadAccountId** | Pointer to **string** | The accountId of the component&#39;s lead user. The accountId uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**LeadUserName** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Metadata** | Pointer to **map[string]string** | Compass component&#39;s metadata. Can&#39;t be updated. Not required for creating a Project Component. | [optional] [readonly] 
**Name** | Pointer to **string** | The unique name for the component in the project. Required when creating a component. Optional when updating a component. The maximum length is 255 characters. | [optional] 
**Project** | Pointer to **string** | The key of the project the component is assigned to. Required when creating a component. Can&#39;t be updated. | [optional] 
**ProjectId** | Pointer to **int64** | The ID of the project the component is assigned to. | [optional] [readonly] 
**RealAssignee** | Pointer to [**User**](User.md) | The user assigned to issues created with this component, when &#x60;assigneeType&#x60; does not identify a valid assignee. | [optional] [readonly] 
**RealAssigneeType** | Pointer to **string** | The type of the assignee that is assigned to issues created with this component, when an assignee cannot be set from the &#x60;assigneeType&#x60;. For example, &#x60;assigneeType&#x60; is set to &#x60;COMPONENT_LEAD&#x60; but no component lead is set. This property is set to one of the following values:   *  &#x60;PROJECT_LEAD&#x60; when &#x60;assigneeType&#x60; is &#x60;PROJECT_LEAD&#x60; and the project lead has permission to be assigned issues in the project that the component is in.  *  &#x60;COMPONENT_LEAD&#x60; when &#x60;assignee&#x60;Type is &#x60;COMPONENT_LEAD&#x60; and the component lead has permission to be assigned issues in the project that the component is in.  *  &#x60;UNASSIGNED&#x60; when &#x60;assigneeType&#x60; is &#x60;UNASSIGNED&#x60; and Jira is configured to allow unassigned issues.  *  &#x60;PROJECT_DEFAULT&#x60; when none of the preceding cases are true. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the component. | [optional] [readonly] 

## Methods

### NewProjectComponent

`func NewProjectComponent() *ProjectComponent`

NewProjectComponent instantiates a new ProjectComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectComponentWithDefaults

`func NewProjectComponentWithDefaults() *ProjectComponent`

NewProjectComponentWithDefaults instantiates a new ProjectComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAri

`func (o *ProjectComponent) GetAri() string`

GetAri returns the Ari field if non-nil, zero value otherwise.

### GetAriOk

`func (o *ProjectComponent) GetAriOk() (*string, bool)`

GetAriOk returns a tuple with the Ari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAri

`func (o *ProjectComponent) SetAri(v string)`

SetAri sets Ari field to given value.

### HasAri

`func (o *ProjectComponent) HasAri() bool`

HasAri returns a boolean if a field has been set.

### GetAssignee

`func (o *ProjectComponent) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ProjectComponent) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ProjectComponent) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ProjectComponent) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAssigneeType

`func (o *ProjectComponent) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *ProjectComponent) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *ProjectComponent) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *ProjectComponent) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectComponent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectComponent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectComponent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectComponent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ProjectComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectComponent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAssigneeTypeValid

`func (o *ProjectComponent) GetIsAssigneeTypeValid() bool`

GetIsAssigneeTypeValid returns the IsAssigneeTypeValid field if non-nil, zero value otherwise.

### GetIsAssigneeTypeValidOk

`func (o *ProjectComponent) GetIsAssigneeTypeValidOk() (*bool, bool)`

GetIsAssigneeTypeValidOk returns a tuple with the IsAssigneeTypeValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigneeTypeValid

`func (o *ProjectComponent) SetIsAssigneeTypeValid(v bool)`

SetIsAssigneeTypeValid sets IsAssigneeTypeValid field to given value.

### HasIsAssigneeTypeValid

`func (o *ProjectComponent) HasIsAssigneeTypeValid() bool`

HasIsAssigneeTypeValid returns a boolean if a field has been set.

### GetLead

`func (o *ProjectComponent) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ProjectComponent) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ProjectComponent) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ProjectComponent) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetLeadAccountId

`func (o *ProjectComponent) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *ProjectComponent) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *ProjectComponent) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *ProjectComponent) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetLeadUserName

`func (o *ProjectComponent) GetLeadUserName() string`

GetLeadUserName returns the LeadUserName field if non-nil, zero value otherwise.

### GetLeadUserNameOk

`func (o *ProjectComponent) GetLeadUserNameOk() (*string, bool)`

GetLeadUserNameOk returns a tuple with the LeadUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadUserName

`func (o *ProjectComponent) SetLeadUserName(v string)`

SetLeadUserName sets LeadUserName field to given value.

### HasLeadUserName

`func (o *ProjectComponent) HasLeadUserName() bool`

HasLeadUserName returns a boolean if a field has been set.

### GetMetadata

`func (o *ProjectComponent) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectComponent) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectComponent) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProjectComponent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ProjectComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *ProjectComponent) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectComponent) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectComponent) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectComponent) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectComponent) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectComponent) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectComponent) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectComponent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRealAssignee

`func (o *ProjectComponent) GetRealAssignee() User`

GetRealAssignee returns the RealAssignee field if non-nil, zero value otherwise.

### GetRealAssigneeOk

`func (o *ProjectComponent) GetRealAssigneeOk() (*User, bool)`

GetRealAssigneeOk returns a tuple with the RealAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealAssignee

`func (o *ProjectComponent) SetRealAssignee(v User)`

SetRealAssignee sets RealAssignee field to given value.

### HasRealAssignee

`func (o *ProjectComponent) HasRealAssignee() bool`

HasRealAssignee returns a boolean if a field has been set.

### GetRealAssigneeType

`func (o *ProjectComponent) GetRealAssigneeType() string`

GetRealAssigneeType returns the RealAssigneeType field if non-nil, zero value otherwise.

### GetRealAssigneeTypeOk

`func (o *ProjectComponent) GetRealAssigneeTypeOk() (*string, bool)`

GetRealAssigneeTypeOk returns a tuple with the RealAssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealAssigneeType

`func (o *ProjectComponent) SetRealAssigneeType(v string)`

SetRealAssigneeType sets RealAssigneeType field to given value.

### HasRealAssigneeType

`func (o *ProjectComponent) HasRealAssigneeType() bool`

HasRealAssigneeType returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectComponent) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectComponent) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectComponent) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ProjectComponent) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


