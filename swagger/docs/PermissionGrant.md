# PermissionGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | Pointer to [**PermissionHolder**](PermissionHolder.md) | The user or group being granted the permission. It consists of a &#x60;type&#x60;, a type-dependent &#x60;parameter&#x60; and a type-dependent &#x60;value&#x60;. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information. | [optional] 
**Id** | Pointer to **int64** | The ID of the permission granted details. | [optional] [readonly] 
**Permission** | Pointer to **string** | The permission to grant. This permission can be one of the built-in permissions or a custom permission added by an app. See [Built-in permissions](../api-group-permission-schemes/#built-in-permissions) in *Get all permission schemes* for more information about the built-in permissions. See the [project permission](https://developer.atlassian.com/cloud/jira/platform/modules/project-permission/) and [global permission](https://developer.atlassian.com/cloud/jira/platform/modules/global-permission/) module documentation for more information about custom permissions. | [optional] 
**Self** | Pointer to **string** | The URL of the permission granted details. | [optional] [readonly] 

## Methods

### NewPermissionGrant

`func NewPermissionGrant() *PermissionGrant`

NewPermissionGrant instantiates a new PermissionGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantWithDefaults

`func NewPermissionGrantWithDefaults() *PermissionGrant`

NewPermissionGrantWithDefaults instantiates a new PermissionGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *PermissionGrant) GetHolder() PermissionHolder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *PermissionGrant) GetHolderOk() (*PermissionHolder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *PermissionGrant) SetHolder(v PermissionHolder)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *PermissionGrant) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetId

`func (o *PermissionGrant) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionGrant) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionGrant) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *PermissionGrant) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionGrant) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionGrant) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionGrant) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetSelf

`func (o *PermissionGrant) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PermissionGrant) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PermissionGrant) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PermissionGrant) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


