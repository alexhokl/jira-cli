# ProjectRoleActorsUpdateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorisedActors** | Pointer to **map[string][]string** | The actors to add to the project role.  Add groups using:   *  &#x60;atlassian-group-role-actor&#x60; and a list of group names.  *  &#x60;atlassian-group-role-actor-id&#x60; and a list of group IDs.  As a group&#39;s name can change, use of &#x60;atlassian-group-role-actor-id&#x60; is recommended. For example, &#x60;\&quot;atlassian-group-role-actor-id\&quot;:[\&quot;eef79f81-0b89-4fca-a736-4be531a10869\&quot;,\&quot;77f6ab39-e755-4570-a6ae-2d7a8df0bcb8\&quot;]&#x60;.  Add users using &#x60;atlassian-user-role-actor&#x60; and a list of account IDs. For example, &#x60;\&quot;atlassian-user-role-actor\&quot;:[\&quot;12345678-9abc-def1-2345-6789abcdef12\&quot;, \&quot;abcdef12-3456-789a-bcde-f123456789ab\&quot;]&#x60;. | [optional] 
**Id** | Pointer to **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | [optional] [readonly] 

## Methods

### NewProjectRoleActorsUpdateBean

`func NewProjectRoleActorsUpdateBean() *ProjectRoleActorsUpdateBean`

NewProjectRoleActorsUpdateBean instantiates a new ProjectRoleActorsUpdateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleActorsUpdateBeanWithDefaults

`func NewProjectRoleActorsUpdateBeanWithDefaults() *ProjectRoleActorsUpdateBean`

NewProjectRoleActorsUpdateBeanWithDefaults instantiates a new ProjectRoleActorsUpdateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorisedActors

`func (o *ProjectRoleActorsUpdateBean) GetCategorisedActors() map[string][]string`

GetCategorisedActors returns the CategorisedActors field if non-nil, zero value otherwise.

### GetCategorisedActorsOk

`func (o *ProjectRoleActorsUpdateBean) GetCategorisedActorsOk() (*map[string][]string, bool)`

GetCategorisedActorsOk returns a tuple with the CategorisedActors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorisedActors

`func (o *ProjectRoleActorsUpdateBean) SetCategorisedActors(v map[string][]string)`

SetCategorisedActors sets CategorisedActors field to given value.

### HasCategorisedActors

`func (o *ProjectRoleActorsUpdateBean) HasCategorisedActors() bool`

HasCategorisedActors returns a boolean if a field has been set.

### GetId

`func (o *ProjectRoleActorsUpdateBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRoleActorsUpdateBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRoleActorsUpdateBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRoleActorsUpdateBean) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


