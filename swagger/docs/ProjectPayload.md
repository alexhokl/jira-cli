# ProjectPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldLayoutSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**IssueSecuritySchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**IssueTypeSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**IssueTypeScreenSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**NotificationSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**PermissionSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**ProjectTypeKey** | Pointer to **string** | The [project type](https://confluence.atlassian.com/x/GwiiLQ#Jiraapplicationsoverview-Productfeaturesandprojecttypes), which defines the application-specific feature set. If you don&#39;t specify the project template you have to specify the project type. | [optional] 
**WorkflowSchemeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewProjectPayload

`func NewProjectPayload() *ProjectPayload`

NewProjectPayload instantiates a new ProjectPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPayloadWithDefaults

`func NewProjectPayloadWithDefaults() *ProjectPayload`

NewProjectPayloadWithDefaults instantiates a new ProjectPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldLayoutSchemeId

`func (o *ProjectPayload) GetFieldLayoutSchemeId() ProjectCreateResourceIdentifier`

GetFieldLayoutSchemeId returns the FieldLayoutSchemeId field if non-nil, zero value otherwise.

### GetFieldLayoutSchemeIdOk

`func (o *ProjectPayload) GetFieldLayoutSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetFieldLayoutSchemeIdOk returns a tuple with the FieldLayoutSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLayoutSchemeId

`func (o *ProjectPayload) SetFieldLayoutSchemeId(v ProjectCreateResourceIdentifier)`

SetFieldLayoutSchemeId sets FieldLayoutSchemeId field to given value.

### HasFieldLayoutSchemeId

`func (o *ProjectPayload) HasFieldLayoutSchemeId() bool`

HasFieldLayoutSchemeId returns a boolean if a field has been set.

### GetIssueSecuritySchemeId

`func (o *ProjectPayload) GetIssueSecuritySchemeId() ProjectCreateResourceIdentifier`

GetIssueSecuritySchemeId returns the IssueSecuritySchemeId field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeIdOk

`func (o *ProjectPayload) GetIssueSecuritySchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetIssueSecuritySchemeIdOk returns a tuple with the IssueSecuritySchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecuritySchemeId

`func (o *ProjectPayload) SetIssueSecuritySchemeId(v ProjectCreateResourceIdentifier)`

SetIssueSecuritySchemeId sets IssueSecuritySchemeId field to given value.

### HasIssueSecuritySchemeId

`func (o *ProjectPayload) HasIssueSecuritySchemeId() bool`

HasIssueSecuritySchemeId returns a boolean if a field has been set.

### GetIssueTypeSchemeId

`func (o *ProjectPayload) GetIssueTypeSchemeId() ProjectCreateResourceIdentifier`

GetIssueTypeSchemeId returns the IssueTypeSchemeId field if non-nil, zero value otherwise.

### GetIssueTypeSchemeIdOk

`func (o *ProjectPayload) GetIssueTypeSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetIssueTypeSchemeIdOk returns a tuple with the IssueTypeSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeSchemeId

`func (o *ProjectPayload) SetIssueTypeSchemeId(v ProjectCreateResourceIdentifier)`

SetIssueTypeSchemeId sets IssueTypeSchemeId field to given value.

### HasIssueTypeSchemeId

`func (o *ProjectPayload) HasIssueTypeSchemeId() bool`

HasIssueTypeSchemeId returns a boolean if a field has been set.

### GetIssueTypeScreenSchemeId

`func (o *ProjectPayload) GetIssueTypeScreenSchemeId() ProjectCreateResourceIdentifier`

GetIssueTypeScreenSchemeId returns the IssueTypeScreenSchemeId field if non-nil, zero value otherwise.

### GetIssueTypeScreenSchemeIdOk

`func (o *ProjectPayload) GetIssueTypeScreenSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetIssueTypeScreenSchemeIdOk returns a tuple with the IssueTypeScreenSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScreenSchemeId

`func (o *ProjectPayload) SetIssueTypeScreenSchemeId(v ProjectCreateResourceIdentifier)`

SetIssueTypeScreenSchemeId sets IssueTypeScreenSchemeId field to given value.

### HasIssueTypeScreenSchemeId

`func (o *ProjectPayload) HasIssueTypeScreenSchemeId() bool`

HasIssueTypeScreenSchemeId returns a boolean if a field has been set.

### GetNotificationSchemeId

`func (o *ProjectPayload) GetNotificationSchemeId() ProjectCreateResourceIdentifier`

GetNotificationSchemeId returns the NotificationSchemeId field if non-nil, zero value otherwise.

### GetNotificationSchemeIdOk

`func (o *ProjectPayload) GetNotificationSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetNotificationSchemeIdOk returns a tuple with the NotificationSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSchemeId

`func (o *ProjectPayload) SetNotificationSchemeId(v ProjectCreateResourceIdentifier)`

SetNotificationSchemeId sets NotificationSchemeId field to given value.

### HasNotificationSchemeId

`func (o *ProjectPayload) HasNotificationSchemeId() bool`

HasNotificationSchemeId returns a boolean if a field has been set.

### GetPcri

`func (o *ProjectPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *ProjectPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *ProjectPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *ProjectPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetPermissionSchemeId

`func (o *ProjectPayload) GetPermissionSchemeId() ProjectCreateResourceIdentifier`

GetPermissionSchemeId returns the PermissionSchemeId field if non-nil, zero value otherwise.

### GetPermissionSchemeIdOk

`func (o *ProjectPayload) GetPermissionSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetPermissionSchemeIdOk returns a tuple with the PermissionSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSchemeId

`func (o *ProjectPayload) SetPermissionSchemeId(v ProjectCreateResourceIdentifier)`

SetPermissionSchemeId sets PermissionSchemeId field to given value.

### HasPermissionSchemeId

`func (o *ProjectPayload) HasPermissionSchemeId() bool`

HasPermissionSchemeId returns a boolean if a field has been set.

### GetProjectTypeKey

`func (o *ProjectPayload) GetProjectTypeKey() string`

GetProjectTypeKey returns the ProjectTypeKey field if non-nil, zero value otherwise.

### GetProjectTypeKeyOk

`func (o *ProjectPayload) GetProjectTypeKeyOk() (*string, bool)`

GetProjectTypeKeyOk returns a tuple with the ProjectTypeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeKey

`func (o *ProjectPayload) SetProjectTypeKey(v string)`

SetProjectTypeKey sets ProjectTypeKey field to given value.

### HasProjectTypeKey

`func (o *ProjectPayload) HasProjectTypeKey() bool`

HasProjectTypeKey returns a boolean if a field has been set.

### GetWorkflowSchemeId

`func (o *ProjectPayload) GetWorkflowSchemeId() ProjectCreateResourceIdentifier`

GetWorkflowSchemeId returns the WorkflowSchemeId field if non-nil, zero value otherwise.

### GetWorkflowSchemeIdOk

`func (o *ProjectPayload) GetWorkflowSchemeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetWorkflowSchemeIdOk returns a tuple with the WorkflowSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSchemeId

`func (o *ProjectPayload) SetWorkflowSchemeId(v ProjectCreateResourceIdentifier)`

SetWorkflowSchemeId sets WorkflowSchemeId field to given value.

### HasWorkflowSchemeId

`func (o *ProjectPayload) HasWorkflowSchemeId() bool`

HasWorkflowSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


