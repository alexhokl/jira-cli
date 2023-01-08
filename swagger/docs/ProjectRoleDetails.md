# ProjectRoleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Whether this role is the admin role for the project. | [optional] [readonly] 
**Default** | Pointer to **bool** | Whether this role is the default role for the project. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the project role. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the project role. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the project role. | [optional] 
**RoleConfigurable** | Pointer to **bool** | Whether the roles are configurable for this project. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the role. Indicated for roles associated with [next-gen projects](https://confluence.atlassian.com/x/loMyO). | [optional] [readonly] 
**Self** | Pointer to **string** | The URL the project role details. | [optional] [readonly] 
**TranslatedName** | Pointer to **string** | The translated name of the project role. | [optional] 
**Type** | Pointer to **string** | The type of the project role. This is \&quot;DEFAULT\&quot; or \&quot;GUEST\\_ROLE\&quot;. | [optional] [readonly] 

## Methods

### NewProjectRoleDetails

`func NewProjectRoleDetails() *ProjectRoleDetails`

NewProjectRoleDetails instantiates a new ProjectRoleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleDetailsWithDefaults

`func NewProjectRoleDetailsWithDefaults() *ProjectRoleDetails`

NewProjectRoleDetailsWithDefaults instantiates a new ProjectRoleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *ProjectRoleDetails) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ProjectRoleDetails) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ProjectRoleDetails) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ProjectRoleDetails) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetDefault

`func (o *ProjectRoleDetails) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ProjectRoleDetails) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ProjectRoleDetails) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ProjectRoleDetails) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectRoleDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectRoleDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectRoleDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectRoleDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ProjectRoleDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectRoleDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectRoleDetails) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectRoleDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectRoleDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRoleDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRoleDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRoleDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleConfigurable

`func (o *ProjectRoleDetails) GetRoleConfigurable() bool`

GetRoleConfigurable returns the RoleConfigurable field if non-nil, zero value otherwise.

### GetRoleConfigurableOk

`func (o *ProjectRoleDetails) GetRoleConfigurableOk() (*bool, bool)`

GetRoleConfigurableOk returns a tuple with the RoleConfigurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleConfigurable

`func (o *ProjectRoleDetails) SetRoleConfigurable(v bool)`

SetRoleConfigurable sets RoleConfigurable field to given value.

### HasRoleConfigurable

`func (o *ProjectRoleDetails) HasRoleConfigurable() bool`

HasRoleConfigurable returns a boolean if a field has been set.

### GetScope

`func (o *ProjectRoleDetails) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProjectRoleDetails) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProjectRoleDetails) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ProjectRoleDetails) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *ProjectRoleDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectRoleDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectRoleDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ProjectRoleDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTranslatedName

`func (o *ProjectRoleDetails) GetTranslatedName() string`

GetTranslatedName returns the TranslatedName field if non-nil, zero value otherwise.

### GetTranslatedNameOk

`func (o *ProjectRoleDetails) GetTranslatedNameOk() (*string, bool)`

GetTranslatedNameOk returns a tuple with the TranslatedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedName

`func (o *ProjectRoleDetails) SetTranslatedName(v string)`

SetTranslatedName sets TranslatedName field to given value.

### HasTranslatedName

`func (o *ProjectRoleDetails) HasTranslatedName() bool`

HasTranslatedName returns a boolean if a field has been set.

### GetType

`func (o *ProjectRoleDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectRoleDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectRoleDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectRoleDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


