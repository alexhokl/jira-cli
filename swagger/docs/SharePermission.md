# SharePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**GroupName**](GroupName.md) | The group that the filter is shared with. For a request, specify the &#x60;groupId&#x60; or &#x60;name&#x60; property for the group. As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended. | [optional] 
**Id** | Pointer to **int64** | The unique identifier of the share permission. | [optional] [readonly] 
**Project** | Pointer to [**Project**](Project.md) | The project that the filter is shared with. This is similar to the project object returned by [Get project](#api-rest-api-3-project-projectIdOrKey-get) but it contains a subset of the properties, which are: &#x60;self&#x60;, &#x60;id&#x60;, &#x60;key&#x60;, &#x60;assigneeType&#x60;, &#x60;name&#x60;, &#x60;roles&#x60;, &#x60;avatarUrls&#x60;, &#x60;projectType&#x60;, &#x60;simplified&#x60;.   For a request, specify the &#x60;id&#x60; for the project. | [optional] 
**Role** | Pointer to [**ProjectRole**](ProjectRole.md) | The project role that the filter is shared with.   For a request, specify the &#x60;id&#x60; for the role. You must also specify the &#x60;project&#x60; object and &#x60;id&#x60; for the project that the role is in. | [optional] 
**Type** | **string** | The type of share permission:   *  &#x60;user&#x60; Shared with a user.  *  &#x60;group&#x60; Shared with a group. If set in a request, then specify &#x60;sharePermission.group&#x60; as well.  *  &#x60;project&#x60; Shared with a project. If set in a request, then specify &#x60;sharePermission.project&#x60; as well.  *  &#x60;projectRole&#x60; Share with a project role in a project. This value is not returned in responses. It is used in requests, where it needs to be specify with &#x60;projectId&#x60; and &#x60;projectRoleId&#x60;.  *  &#x60;global&#x60; Shared globally. If set in a request, no other &#x60;sharePermission&#x60; properties need to be specified.  *  &#x60;loggedin&#x60; Shared with all logged-in users. Note: This value is set in a request by specifying &#x60;authenticated&#x60; as the &#x60;type&#x60;.  *  &#x60;project-unknown&#x60; Shared with a project that the user does not have access to. Cannot be set in a request. | 
**User** | Pointer to [**UserBean**](UserBean.md) | The user account ID that the filter is shared with. For a request, specify the &#x60;accountId&#x60; property for the user. | [optional] 

## Methods

### NewSharePermission

`func NewSharePermission(type_ string, ) *SharePermission`

NewSharePermission instantiates a new SharePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePermissionWithDefaults

`func NewSharePermissionWithDefaults() *SharePermission`

NewSharePermissionWithDefaults instantiates a new SharePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SharePermission) GetGroup() GroupName`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SharePermission) GetGroupOk() (*GroupName, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SharePermission) SetGroup(v GroupName)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SharePermission) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *SharePermission) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharePermission) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharePermission) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SharePermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *SharePermission) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SharePermission) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SharePermission) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *SharePermission) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRole

`func (o *SharePermission) GetRole() ProjectRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SharePermission) GetRoleOk() (*ProjectRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SharePermission) SetRole(v ProjectRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *SharePermission) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *SharePermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharePermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharePermission) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *SharePermission) GetUser() UserBean`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SharePermission) GetUserOk() (*UserBean, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SharePermission) SetUser(v UserBean)`

SetUser sets User field to given value.

### HasUser

`func (o *SharePermission) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


