# PermissionScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the permission scheme. | [optional] 
**Expand** | Pointer to **string** | The expand options available for the permission scheme. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the permission scheme. | [optional] [readonly] 
**Name** | **string** | The name of the permission scheme. Must be unique. | 
**Permissions** | Pointer to [**[]PermissionGrant**](PermissionGrant.md) | The permission scheme to create or update. See [About permission schemes and grants](../api-group-permission-schemes/#about-permission-schemes-and-grants) for more information. | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the permission scheme. | [optional] 
**Self** | Pointer to **string** | The URL of the permission scheme. | [optional] [readonly] 

## Methods

### NewPermissionScheme

`func NewPermissionScheme(name string, ) *PermissionScheme`

NewPermissionScheme instantiates a new PermissionScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSchemeWithDefaults

`func NewPermissionSchemeWithDefaults() *PermissionScheme`

NewPermissionSchemeWithDefaults instantiates a new PermissionScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PermissionScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpand

`func (o *PermissionScheme) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *PermissionScheme) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *PermissionScheme) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *PermissionScheme) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetId

`func (o *PermissionScheme) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionScheme) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionScheme) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PermissionScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionScheme) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *PermissionScheme) GetPermissions() []PermissionGrant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionScheme) GetPermissionsOk() (*[]PermissionGrant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionScheme) SetPermissions(v []PermissionGrant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PermissionScheme) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetScope

`func (o *PermissionScheme) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PermissionScheme) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PermissionScheme) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PermissionScheme) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *PermissionScheme) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PermissionScheme) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PermissionScheme) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PermissionScheme) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


