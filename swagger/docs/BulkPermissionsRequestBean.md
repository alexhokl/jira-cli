# BulkPermissionsRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of a user. | [optional] 
**GlobalPermissions** | Pointer to **[]string** | Global permissions to look up. | [optional] 
**ProjectPermissions** | Pointer to [**[]BulkProjectPermissions**](BulkProjectPermissions.md) | Project permissions with associated projects and issues to look up. | [optional] 

## Methods

### NewBulkPermissionsRequestBean

`func NewBulkPermissionsRequestBean() *BulkPermissionsRequestBean`

NewBulkPermissionsRequestBean instantiates a new BulkPermissionsRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPermissionsRequestBeanWithDefaults

`func NewBulkPermissionsRequestBeanWithDefaults() *BulkPermissionsRequestBean`

NewBulkPermissionsRequestBeanWithDefaults instantiates a new BulkPermissionsRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BulkPermissionsRequestBean) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BulkPermissionsRequestBean) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BulkPermissionsRequestBean) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BulkPermissionsRequestBean) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *BulkPermissionsRequestBean) GetGlobalPermissions() []string`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BulkPermissionsRequestBean) GetGlobalPermissionsOk() (*[]string, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BulkPermissionsRequestBean) SetGlobalPermissions(v []string)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *BulkPermissionsRequestBean) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetProjectPermissions

`func (o *BulkPermissionsRequestBean) GetProjectPermissions() []BulkProjectPermissions`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *BulkPermissionsRequestBean) GetProjectPermissionsOk() (*[]BulkProjectPermissions, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *BulkPermissionsRequestBean) SetProjectPermissions(v []BulkProjectPermissions)`

SetProjectPermissions sets ProjectPermissions field to given value.

### HasProjectPermissions

`func (o *BulkPermissionsRequestBean) HasProjectPermissions() bool`

HasProjectPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


