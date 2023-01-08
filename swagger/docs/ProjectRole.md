# ProjectRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actors** | Pointer to [**[]RoleActor**](RoleActor.md) | The list of users who act in this role. | [optional] [readonly] 
**Admin** | Pointer to **bool** | Whether this role is the admin role for the project. | [optional] [readonly] 
**CurrentUserRole** | Pointer to **bool** | Whether the calling user is part of this role. | [optional] 
**Default** | Pointer to **bool** | Whether this role is the default role for the project | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the project role. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the project role. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the project role. | [optional] 
**RoleConfigurable** | Pointer to **bool** | Whether the roles are configurable for this project. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the role. Indicated for roles associated with [next-gen projects](https://confluence.atlassian.com/x/loMyO). | [optional] [readonly] 
**Self** | Pointer to **string** | The URL the project role details. | [optional] [readonly] 
**TranslatedName** | Pointer to **string** | The translated name of the project role. | [optional] 

## Methods

### NewProjectRole

`func NewProjectRole() *ProjectRole`

NewProjectRole instantiates a new ProjectRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleWithDefaults

`func NewProjectRoleWithDefaults() *ProjectRole`

NewProjectRoleWithDefaults instantiates a new ProjectRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActors

`func (o *ProjectRole) GetActors() []RoleActor`

GetActors returns the Actors field if non-nil, zero value otherwise.

### GetActorsOk

`func (o *ProjectRole) GetActorsOk() (*[]RoleActor, bool)`

GetActorsOk returns a tuple with the Actors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActors

`func (o *ProjectRole) SetActors(v []RoleActor)`

SetActors sets Actors field to given value.

### HasActors

`func (o *ProjectRole) HasActors() bool`

HasActors returns a boolean if a field has been set.

### GetAdmin

`func (o *ProjectRole) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ProjectRole) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ProjectRole) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ProjectRole) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCurrentUserRole

`func (o *ProjectRole) GetCurrentUserRole() bool`

GetCurrentUserRole returns the CurrentUserRole field if non-nil, zero value otherwise.

### GetCurrentUserRoleOk

`func (o *ProjectRole) GetCurrentUserRoleOk() (*bool, bool)`

GetCurrentUserRoleOk returns a tuple with the CurrentUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserRole

`func (o *ProjectRole) SetCurrentUserRole(v bool)`

SetCurrentUserRole sets CurrentUserRole field to given value.

### HasCurrentUserRole

`func (o *ProjectRole) HasCurrentUserRole() bool`

HasCurrentUserRole returns a boolean if a field has been set.

### GetDefault

`func (o *ProjectRole) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ProjectRole) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ProjectRole) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ProjectRole) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ProjectRole) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRole) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRole) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleConfigurable

`func (o *ProjectRole) GetRoleConfigurable() bool`

GetRoleConfigurable returns the RoleConfigurable field if non-nil, zero value otherwise.

### GetRoleConfigurableOk

`func (o *ProjectRole) GetRoleConfigurableOk() (*bool, bool)`

GetRoleConfigurableOk returns a tuple with the RoleConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleConfigurable

`func (o *ProjectRole) SetRoleConfigurable(v bool)`

SetRoleConfigurable sets RoleConfigurable field to given value.

### HasRoleConfigurable

`func (o *ProjectRole) HasRoleConfigurable() bool`

HasRoleConfigurable returns a boolean if a field has been set.

### GetScope

`func (o *ProjectRole) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProjectRole) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProjectRole) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ProjectRole) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectRole) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectRole) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectRole) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ProjectRole) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTranslatedName

`func (o *ProjectRole) GetTranslatedName() string`

GetTranslatedName returns the TranslatedName field if non-nil, zero value otherwise.

### GetTranslatedNameOk

`func (o *ProjectRole) GetTranslatedNameOk() (*string, bool)`

GetTranslatedNameOk returns a tuple with the TranslatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedName

`func (o *ProjectRole) SetTranslatedName(v string)`

SetTranslatedName sets TranslatedName field to given value.

### HasTranslatedName

`func (o *ProjectRole) HasTranslatedName() bool`

HasTranslatedName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


