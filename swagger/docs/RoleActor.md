# RoleActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorGroup** | Pointer to [**ProjectRoleGroup**](ProjectRoleGroup.md) |  | [optional] [readonly] 
**ActorUser** | Pointer to [**ProjectRoleUser**](ProjectRoleUser.md) |  | [optional] [readonly] 
**AvatarUrl** | Pointer to **string** | The avatar of the role actor. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the role actor. For users, depending on the user’s privacy setting, this may return an alternative value for the user&#39;s name. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the role actor. | [optional] [readonly] 
**Name** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of role actor. | [optional] [readonly] 

## Methods

### NewRoleActor

`func NewRoleActor() *RoleActor`

NewRoleActor instantiates a new RoleActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleActorWithDefaults

`func NewRoleActorWithDefaults() *RoleActor`

NewRoleActorWithDefaults instantiates a new RoleActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorGroup

`func (o *RoleActor) GetActorGroup() ProjectRoleGroup`

GetActorGroup returns the ActorGroup field if non-nil, zero value otherwise.

### GetActorGroupOk

`func (o *RoleActor) GetActorGroupOk() (*ProjectRoleGroup, bool)`

GetActorGroupOk returns a tuple with the ActorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorGroup

`func (o *RoleActor) SetActorGroup(v ProjectRoleGroup)`

SetActorGroup sets ActorGroup field to given value.

### HasActorGroup

`func (o *RoleActor) HasActorGroup() bool`

HasActorGroup returns a boolean if a field has been set.

### GetActorUser

`func (o *RoleActor) GetActorUser() ProjectRoleUser`

GetActorUser returns the ActorUser field if non-nil, zero value otherwise.

### GetActorUserOk

`func (o *RoleActor) GetActorUserOk() (*ProjectRoleUser, bool)`

GetActorUserOk returns a tuple with the ActorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUser

`func (o *RoleActor) SetActorUser(v ProjectRoleUser)`

SetActorUser sets ActorUser field to given value.

### HasActorUser

`func (o *RoleActor) HasActorUser() bool`

HasActorUser returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *RoleActor) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *RoleActor) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *RoleActor) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *RoleActor) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *RoleActor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleActor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RoleActor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RoleActor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *RoleActor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleActor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleActor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RoleActor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleActor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleActor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RoleActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleActor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleActor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


