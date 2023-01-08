# UserPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeprecatedKey** | Pointer to **bool** | Indicate whether the permission key is deprecated. Note that deprecated keys cannot be used in the &#x60;permissions parameter of Get my permissions. Deprecated keys are not returned by Get all permissions.&#x60; | [optional] 
**Description** | Pointer to **string** | The description of the permission. | [optional] 
**HavePermission** | Pointer to **bool** | Whether the permission is available to the user in the queried context. | [optional] 
**Id** | Pointer to **string** | The ID of the permission. Either &#x60;id&#x60; or &#x60;key&#x60; must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions. | [optional] 
**Key** | Pointer to **string** | The key of the permission. Either &#x60;id&#x60; or &#x60;key&#x60; must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions. | [optional] 
**Name** | Pointer to **string** | The name of the permission. | [optional] 
**Type** | Pointer to **string** | The type of the permission. | [optional] 

## Methods

### NewUserPermission

`func NewUserPermission() *UserPermission`

NewUserPermission instantiates a new UserPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionWithDefaults

`func NewUserPermissionWithDefaults() *UserPermission`

NewUserPermissionWithDefaults instantiates a new UserPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeprecatedKey

`func (o *UserPermission) GetDeprecatedKey() bool`

GetDeprecatedKey returns the DeprecatedKey field if non-nil, zero value otherwise.

### GetDeprecatedKeyOk

`func (o *UserPermission) GetDeprecatedKeyOk() (*bool, bool)`

GetDeprecatedKeyOk returns a tuple with the DeprecatedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedKey

`func (o *UserPermission) SetDeprecatedKey(v bool)`

SetDeprecatedKey sets DeprecatedKey field to given value.

### HasDeprecatedKey

`func (o *UserPermission) HasDeprecatedKey() bool`

HasDeprecatedKey returns a boolean if a field has been set.

### GetDescription

`func (o *UserPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHavePermission

`func (o *UserPermission) GetHavePermission() bool`

GetHavePermission returns the HavePermission field if non-nil, zero value otherwise.

### GetHavePermissionOk

`func (o *UserPermission) GetHavePermissionOk() (*bool, bool)`

GetHavePermissionOk returns a tuple with the HavePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHavePermission

`func (o *UserPermission) SetHavePermission(v bool)`

SetHavePermission sets HavePermission field to given value.

### HasHavePermission

`func (o *UserPermission) HasHavePermission() bool`

HasHavePermission returns a boolean if a field has been set.

### GetId

`func (o *UserPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *UserPermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserPermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserPermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserPermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UserPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UserPermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserPermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserPermission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserPermission) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


