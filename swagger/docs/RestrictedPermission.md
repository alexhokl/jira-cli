# RestrictedPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the permission. Either &#x60;id&#x60; or &#x60;key&#x60; must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions. | [optional] 
**Key** | Pointer to **string** | The key of the permission. Either &#x60;id&#x60; or &#x60;key&#x60; must be specified. Use [Get all permissions](#api-rest-api-3-permissions-get) to get the list of permissions. | [optional] 

## Methods

### NewRestrictedPermission

`func NewRestrictedPermission() *RestrictedPermission`

NewRestrictedPermission instantiates a new RestrictedPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedPermissionWithDefaults

`func NewRestrictedPermissionWithDefaults() *RestrictedPermission`

NewRestrictedPermissionWithDefaults instantiates a new RestrictedPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestrictedPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestrictedPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestrictedPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestrictedPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *RestrictedPermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RestrictedPermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RestrictedPermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RestrictedPermission) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


