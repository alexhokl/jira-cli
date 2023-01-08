# SharePermissionInputBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The user account ID that the filter is shared with. For a request, specify the &#x60;accountId&#x60; property for the user. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group, which uniquely identifies the group across all Atlassian products.For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*. Cannot be provided with &#x60;groupname&#x60;. | [optional] 
**Groupname** | Pointer to **string** | The name of the group to share the filter with. Set &#x60;type&#x60; to &#x60;group&#x60;. Please note that the name of a group is mutable, to reliably identify a group use &#x60;groupId&#x60;. | [optional] 
**ProjectId** | Pointer to **string** | The ID of the project to share the filter with. Set &#x60;type&#x60; to &#x60;project&#x60;. | [optional] 
**ProjectRoleId** | Pointer to **string** | The ID of the project role to share the filter with. Set &#x60;type&#x60; to &#x60;projectRole&#x60; and the &#x60;projectId&#x60; for the project that the role is in. | [optional] 
**Rights** | Pointer to **int32** | The rights for the share permission. | [optional] 
**Type** | **string** | The type of the share permission.Specify the type as follows:   *  &#x60;user&#x60; Share with a user.  *  &#x60;group&#x60; Share with a group. Specify &#x60;groupname&#x60; as well.  *  &#x60;project&#x60; Share with a project. Specify &#x60;projectId&#x60; as well.  *  &#x60;projectRole&#x60; Share with a project role in a project. Specify &#x60;projectId&#x60; and &#x60;projectRoleId&#x60; as well.  *  &#x60;global&#x60; Share globally, including anonymous users. If set, this type overrides all existing share permissions and must be deleted before any non-global share permissions is set.  *  &#x60;authenticated&#x60; Share with all logged-in users. This shows as &#x60;loggedin&#x60; in the response. If set, this type overrides all existing share permissions and must be deleted before any non-global share permissions is set. | 

## Methods

### NewSharePermissionInputBean

`func NewSharePermissionInputBean(type_ string, ) *SharePermissionInputBean`

NewSharePermissionInputBean instantiates a new SharePermissionInputBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePermissionInputBeanWithDefaults

`func NewSharePermissionInputBeanWithDefaults() *SharePermissionInputBean`

NewSharePermissionInputBeanWithDefaults instantiates a new SharePermissionInputBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SharePermissionInputBean) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SharePermissionInputBean) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SharePermissionInputBean) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SharePermissionInputBean) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroupId

`func (o *SharePermissionInputBean) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SharePermissionInputBean) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SharePermissionInputBean) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SharePermissionInputBean) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupname

`func (o *SharePermissionInputBean) GetGroupname() string`

GetGroupname returns the Groupname field if non-nil, zero value otherwise.

### GetGroupnameOk

`func (o *SharePermissionInputBean) GetGroupnameOk() (*string, bool)`

GetGroupnameOk returns a tuple with the Groupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupname

`func (o *SharePermissionInputBean) SetGroupname(v string)`

SetGroupname sets Groupname field to given value.

### HasGroupname

`func (o *SharePermissionInputBean) HasGroupname() bool`

HasGroupname returns a boolean if a field has been set.

### GetProjectId

`func (o *SharePermissionInputBean) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SharePermissionInputBean) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SharePermissionInputBean) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SharePermissionInputBean) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectRoleId

`func (o *SharePermissionInputBean) GetProjectRoleId() string`

GetProjectRoleId returns the ProjectRoleId field if non-nil, zero value otherwise.

### GetProjectRoleIdOk

`func (o *SharePermissionInputBean) GetProjectRoleIdOk() (*string, bool)`

GetProjectRoleIdOk returns a tuple with the ProjectRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoleId

`func (o *SharePermissionInputBean) SetProjectRoleId(v string)`

SetProjectRoleId sets ProjectRoleId field to given value.

### HasProjectRoleId

`func (o *SharePermissionInputBean) HasProjectRoleId() bool`

HasProjectRoleId returns a boolean if a field has been set.

### GetRights

`func (o *SharePermissionInputBean) GetRights() int32`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *SharePermissionInputBean) GetRightsOk() (*int32, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *SharePermissionInputBean) SetRights(v int32)`

SetRights sets Rights field to given value.

### HasRights

`func (o *SharePermissionInputBean) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetType

`func (o *SharePermissionInputBean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SharePermissionInputBean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SharePermissionInputBean) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


