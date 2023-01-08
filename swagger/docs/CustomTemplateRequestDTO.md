# CustomTemplateRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boards** | Pointer to [**NullableBoardsPayload**](BoardsPayload.md) |  | [optional] 
**Field** | Pointer to [**NullableFieldCapabilityPayload**](FieldCapabilityPayload.md) |  | [optional] 
**IssueType** | Pointer to [**NullableIssueTypeProjectCreatePayload**](IssueTypeProjectCreatePayload.md) |  | [optional] 
**Notification** | Pointer to [**NullableNotificationSchemePayload**](NotificationSchemePayload.md) |  | [optional] 
**PermissionScheme** | Pointer to [**NullablePermissionPayloadDTO**](PermissionPayloadDTO.md) |  | [optional] 
**Project** | Pointer to [**ProjectPayload**](ProjectPayload.md) |  | [optional] 
**Role** | Pointer to [**NullableRolesCapabilityPayload**](RolesCapabilityPayload.md) |  | [optional] 
**Scope** | Pointer to [**NullableScopePayload**](ScopePayload.md) |  | [optional] 
**Security** | Pointer to [**NullableSecuritySchemePayload**](SecuritySchemePayload.md) |  | [optional] 
**Workflow** | Pointer to [**NullableWorkflowCapabilityPayload**](WorkflowCapabilityPayload.md) |  | [optional] 

## Methods

### NewCustomTemplateRequestDTO

`func NewCustomTemplateRequestDTO() *CustomTemplateRequestDTO`

NewCustomTemplateRequestDTO instantiates a new CustomTemplateRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTemplateRequestDTOWithDefaults

`func NewCustomTemplateRequestDTOWithDefaults() *CustomTemplateRequestDTO`

NewCustomTemplateRequestDTOWithDefaults instantiates a new CustomTemplateRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoards

`func (o *CustomTemplateRequestDTO) GetBoards() BoardsPayload`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *CustomTemplateRequestDTO) GetBoardsOk() (*BoardsPayload, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *CustomTemplateRequestDTO) SetBoards(v BoardsPayload)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *CustomTemplateRequestDTO) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### SetBoardsNil

`func (o *CustomTemplateRequestDTO) SetBoardsNil(b bool)`

 SetBoardsNil sets the value for Boards to be an explicit nil

### UnsetBoards
`func (o *CustomTemplateRequestDTO) UnsetBoards()`

UnsetBoards ensures that no value is present for Boards, not even an explicit nil
### GetField

`func (o *CustomTemplateRequestDTO) GetField() FieldCapabilityPayload`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CustomTemplateRequestDTO) GetFieldOk() (*FieldCapabilityPayload, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CustomTemplateRequestDTO) SetField(v FieldCapabilityPayload)`

SetField sets Field field to given value.

### HasField

`func (o *CustomTemplateRequestDTO) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *CustomTemplateRequestDTO) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *CustomTemplateRequestDTO) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetIssueType

`func (o *CustomTemplateRequestDTO) GetIssueType() IssueTypeProjectCreatePayload`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *CustomTemplateRequestDTO) GetIssueTypeOk() (*IssueTypeProjectCreatePayload, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *CustomTemplateRequestDTO) SetIssueType(v IssueTypeProjectCreatePayload)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *CustomTemplateRequestDTO) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### SetIssueTypeNil

`func (o *CustomTemplateRequestDTO) SetIssueTypeNil(b bool)`

 SetIssueTypeNil sets the value for IssueType to be an explicit nil

### UnsetIssueType
`func (o *CustomTemplateRequestDTO) UnsetIssueType()`

UnsetIssueType ensures that no value is present for IssueType, not even an explicit nil
### GetNotification

`func (o *CustomTemplateRequestDTO) GetNotification() NotificationSchemePayload`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *CustomTemplateRequestDTO) GetNotificationOk() (*NotificationSchemePayload, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *CustomTemplateRequestDTO) SetNotification(v NotificationSchemePayload)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *CustomTemplateRequestDTO) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### SetNotificationNil

`func (o *CustomTemplateRequestDTO) SetNotificationNil(b bool)`

 SetNotificationNil sets the value for Notification to be an explicit nil

### UnsetNotification
`func (o *CustomTemplateRequestDTO) UnsetNotification()`

UnsetNotification ensures that no value is present for Notification, not even an explicit nil
### GetPermissionScheme

`func (o *CustomTemplateRequestDTO) GetPermissionScheme() PermissionPayloadDTO`

GetPermissionScheme returns the PermissionScheme field if non-nil, zero value otherwise.

### GetPermissionSchemeOk

`func (o *CustomTemplateRequestDTO) GetPermissionSchemeOk() (*PermissionPayloadDTO, bool)`

GetPermissionSchemeOk returns a tuple with the PermissionScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionScheme

`func (o *CustomTemplateRequestDTO) SetPermissionScheme(v PermissionPayloadDTO)`

SetPermissionScheme sets PermissionScheme field to given value.

### HasPermissionScheme

`func (o *CustomTemplateRequestDTO) HasPermissionScheme() bool`

HasPermissionScheme returns a boolean if a field has been set.

### SetPermissionSchemeNil

`func (o *CustomTemplateRequestDTO) SetPermissionSchemeNil(b bool)`

 SetPermissionSchemeNil sets the value for PermissionScheme to be an explicit nil

### UnsetPermissionScheme
`func (o *CustomTemplateRequestDTO) UnsetPermissionScheme()`

UnsetPermissionScheme ensures that no value is present for PermissionScheme, not even an explicit nil
### GetProject

`func (o *CustomTemplateRequestDTO) GetProject() ProjectPayload`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CustomTemplateRequestDTO) GetProjectOk() (*ProjectPayload, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CustomTemplateRequestDTO) SetProject(v ProjectPayload)`

SetProject sets Project field to given value.

### HasProject

`func (o *CustomTemplateRequestDTO) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRole

`func (o *CustomTemplateRequestDTO) GetRole() RolesCapabilityPayload`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CustomTemplateRequestDTO) GetRoleOk() (*RolesCapabilityPayload, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CustomTemplateRequestDTO) SetRole(v RolesCapabilityPayload)`

SetRole sets Role field to given value.

### HasRole

`func (o *CustomTemplateRequestDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *CustomTemplateRequestDTO) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *CustomTemplateRequestDTO) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetScope

`func (o *CustomTemplateRequestDTO) GetScope() ScopePayload`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CustomTemplateRequestDTO) GetScopeOk() (*ScopePayload, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CustomTemplateRequestDTO) SetScope(v ScopePayload)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CustomTemplateRequestDTO) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CustomTemplateRequestDTO) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CustomTemplateRequestDTO) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSecurity

`func (o *CustomTemplateRequestDTO) GetSecurity() SecuritySchemePayload`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CustomTemplateRequestDTO) GetSecurityOk() (*SecuritySchemePayload, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CustomTemplateRequestDTO) SetSecurity(v SecuritySchemePayload)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *CustomTemplateRequestDTO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### SetSecurityNil

`func (o *CustomTemplateRequestDTO) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *CustomTemplateRequestDTO) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetWorkflow

`func (o *CustomTemplateRequestDTO) GetWorkflow() WorkflowCapabilityPayload`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CustomTemplateRequestDTO) GetWorkflowOk() (*WorkflowCapabilityPayload, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CustomTemplateRequestDTO) SetWorkflow(v WorkflowCapabilityPayload)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CustomTemplateRequestDTO) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *CustomTemplateRequestDTO) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *CustomTemplateRequestDTO) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


