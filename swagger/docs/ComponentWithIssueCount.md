# ComponentWithIssueCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**User**](User.md) | The details of the user associated with &#x60;assigneeType&#x60;, if any. See &#x60;realAssignee&#x60; for details of the user assigned to issues created with this component. | [optional] 
**AssigneeType** | Pointer to **string** | The nominal user type used to determine the assignee for issues created with this component. See &#x60;realAssigneeType&#x60; for details on how the type of the user, and hence the user, assigned to issues is determined. Takes the following values:   *  &#x60;PROJECT_LEAD&#x60; the assignee to any issues created with this component is nominally the lead for the project the component is in.  *  &#x60;COMPONENT_LEAD&#x60; the assignee to any issues created with this component is nominally the lead for the component.  *  &#x60;UNASSIGNED&#x60; an assignee is not set for issues created with this component.  *  &#x60;PROJECT_DEFAULT&#x60; the assignee to any issues created with this component is nominally the default assignee for the project that the component is in. | [optional] [readonly] 
**Description** | Pointer to **string** | The description for the component. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the component. | [optional] [readonly] 
**IsAssigneeTypeValid** | Pointer to **bool** | Whether a user is associated with &#x60;assigneeType&#x60;. For example, if the &#x60;assigneeType&#x60; is set to &#x60;COMPONENT_LEAD&#x60; but the component lead is not set, then &#x60;false&#x60; is returned. | [optional] [readonly] 
**IssueCount** | Pointer to **int64** | Count of issues for the component. | [optional] [readonly] 
**Lead** | Pointer to [**User**](User.md) | The user details for the component&#39;s lead user. | [optional] 
**Name** | Pointer to **string** | The name for the component. | [optional] [readonly] 
**Project** | Pointer to **string** | The key of the project to which the component is assigned. | [optional] [readonly] 
**ProjectId** | Pointer to **int64** | Not used. | [optional] [readonly] 
**RealAssignee** | Pointer to [**User**](User.md) | The user assigned to issues created with this component, when &#x60;assigneeType&#x60; does not identify a valid assignee. | [optional] 
**RealAssigneeType** | Pointer to **string** | The type of the assignee that is assigned to issues created with this component, when an assignee cannot be set from the &#x60;assigneeType&#x60;. For example, &#x60;assigneeType&#x60; is set to &#x60;COMPONENT_LEAD&#x60; but no component lead is set. This property is set to one of the following values:   *  &#x60;PROJECT_LEAD&#x60; when &#x60;assigneeType&#x60; is &#x60;PROJECT_LEAD&#x60; and the project lead has permission to be assigned issues in the project that the component is in.  *  &#x60;COMPONENT_LEAD&#x60; when &#x60;assignee&#x60;Type is &#x60;COMPONENT_LEAD&#x60; and the component lead has permission to be assigned issues in the project that the component is in.  *  &#x60;UNASSIGNED&#x60; when &#x60;assigneeType&#x60; is &#x60;UNASSIGNED&#x60; and Jira is configured to allow unassigned issues.  *  &#x60;PROJECT_DEFAULT&#x60; when none of the preceding cases are true. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL for this count of the issues contained in the component. | [optional] [readonly] 

## Methods

### NewComponentWithIssueCount

`func NewComponentWithIssueCount() *ComponentWithIssueCount`

NewComponentWithIssueCount instantiates a new ComponentWithIssueCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithIssueCountWithDefaults

`func NewComponentWithIssueCountWithDefaults() *ComponentWithIssueCount`

NewComponentWithIssueCountWithDefaults instantiates a new ComponentWithIssueCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ComponentWithIssueCount) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ComponentWithIssueCount) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ComponentWithIssueCount) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ComponentWithIssueCount) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetAssigneeType

`func (o *ComponentWithIssueCount) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *ComponentWithIssueCount) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *ComponentWithIssueCount) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *ComponentWithIssueCount) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetDescription

`func (o *ComponentWithIssueCount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentWithIssueCount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentWithIssueCount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentWithIssueCount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ComponentWithIssueCount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentWithIssueCount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentWithIssueCount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentWithIssueCount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAssigneeTypeValid

`func (o *ComponentWithIssueCount) GetIsAssigneeTypeValid() bool`

GetIsAssigneeTypeValid returns the IsAssigneeTypeValid field if non-nil, zero value otherwise.

### GetIsAssigneeTypeValidOk

`func (o *ComponentWithIssueCount) GetIsAssigneeTypeValidOk() (*bool, bool)`

GetIsAssigneeTypeValidOk returns a tuple with the IsAssigneeTypeValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssigneeTypeValid

`func (o *ComponentWithIssueCount) SetIsAssigneeTypeValid(v bool)`

SetIsAssigneeTypeValid sets IsAssigneeTypeValid field to given value.

### HasIsAssigneeTypeValid

`func (o *ComponentWithIssueCount) HasIsAssigneeTypeValid() bool`

HasIsAssigneeTypeValid returns a boolean if a field has been set.

### GetIssueCount

`func (o *ComponentWithIssueCount) GetIssueCount() int64`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *ComponentWithIssueCount) GetIssueCountOk() (*int64, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *ComponentWithIssueCount) SetIssueCount(v int64)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *ComponentWithIssueCount) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetLead

`func (o *ComponentWithIssueCount) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ComponentWithIssueCount) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ComponentWithIssueCount) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ComponentWithIssueCount) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetName

`func (o *ComponentWithIssueCount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentWithIssueCount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentWithIssueCount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentWithIssueCount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *ComponentWithIssueCount) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ComponentWithIssueCount) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ComponentWithIssueCount) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ComponentWithIssueCount) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *ComponentWithIssueCount) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ComponentWithIssueCount) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ComponentWithIssueCount) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ComponentWithIssueCount) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRealAssignee

`func (o *ComponentWithIssueCount) GetRealAssignee() User`

GetRealAssignee returns the RealAssignee field if non-nil, zero value otherwise.

### GetRealAssigneeOk

`func (o *ComponentWithIssueCount) GetRealAssigneeOk() (*User, bool)`

GetRealAssigneeOk returns a tuple with the RealAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealAssignee

`func (o *ComponentWithIssueCount) SetRealAssignee(v User)`

SetRealAssignee sets RealAssignee field to given value.

### HasRealAssignee

`func (o *ComponentWithIssueCount) HasRealAssignee() bool`

HasRealAssignee returns a boolean if a field has been set.

### GetRealAssigneeType

`func (o *ComponentWithIssueCount) GetRealAssigneeType() string`

GetRealAssigneeType returns the RealAssigneeType field if non-nil, zero value otherwise.

### GetRealAssigneeTypeOk

`func (o *ComponentWithIssueCount) GetRealAssigneeTypeOk() (*string, bool)`

GetRealAssigneeTypeOk returns a tuple with the RealAssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealAssigneeType

`func (o *ComponentWithIssueCount) SetRealAssigneeType(v string)`

SetRealAssigneeType sets RealAssigneeType field to given value.

### HasRealAssigneeType

`func (o *ComponentWithIssueCount) HasRealAssigneeType() bool`

HasRealAssigneeType returns a boolean if a field has been set.

### GetSelf

`func (o *ComponentWithIssueCount) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ComponentWithIssueCount) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ComponentWithIssueCount) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ComponentWithIssueCount) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


