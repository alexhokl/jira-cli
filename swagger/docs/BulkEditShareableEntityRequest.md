# BulkEditShareableEntityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Allowed action for bulk edit shareable entity | 
**ChangeOwnerDetails** | Pointer to [**BulkChangeOwnerDetails**](BulkChangeOwnerDetails.md) | The details of change owner action. | [optional] 
**EntityIds** | **[]int64** | The id list of shareable entities to be changed. | 
**ExtendAdminPermissions** | Pointer to **bool** | Whether the actions are executed by users with Administer Jira global permission. | [optional] 
**PermissionDetails** | Pointer to [**PermissionDetails**](PermissionDetails.md) | The permission details to be changed. | [optional] 

## Methods

### NewBulkEditShareableEntityRequest

`func NewBulkEditShareableEntityRequest(action string, entityIds []int64, ) *BulkEditShareableEntityRequest`

NewBulkEditShareableEntityRequest instantiates a new BulkEditShareableEntityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditShareableEntityRequestWithDefaults

`func NewBulkEditShareableEntityRequestWithDefaults() *BulkEditShareableEntityRequest`

NewBulkEditShareableEntityRequestWithDefaults instantiates a new BulkEditShareableEntityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkEditShareableEntityRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkEditShareableEntityRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkEditShareableEntityRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetChangeOwnerDetails

`func (o *BulkEditShareableEntityRequest) GetChangeOwnerDetails() BulkChangeOwnerDetails`

GetChangeOwnerDetails returns the ChangeOwnerDetails field if non-nil, zero value otherwise.

### GetChangeOwnerDetailsOk

`func (o *BulkEditShareableEntityRequest) GetChangeOwnerDetailsOk() (*BulkChangeOwnerDetails, bool)`

GetChangeOwnerDetailsOk returns a tuple with the ChangeOwnerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeOwnerDetails

`func (o *BulkEditShareableEntityRequest) SetChangeOwnerDetails(v BulkChangeOwnerDetails)`

SetChangeOwnerDetails sets ChangeOwnerDetails field to given value.

### HasChangeOwnerDetails

`func (o *BulkEditShareableEntityRequest) HasChangeOwnerDetails() bool`

HasChangeOwnerDetails returns a boolean if a field has been set.

### GetEntityIds

`func (o *BulkEditShareableEntityRequest) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *BulkEditShareableEntityRequest) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *BulkEditShareableEntityRequest) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.


### GetExtendAdminPermissions

`func (o *BulkEditShareableEntityRequest) GetExtendAdminPermissions() bool`

GetExtendAdminPermissions returns the ExtendAdminPermissions field if non-nil, zero value otherwise.

### GetExtendAdminPermissionsOk

`func (o *BulkEditShareableEntityRequest) GetExtendAdminPermissionsOk() (*bool, bool)`

GetExtendAdminPermissionsOk returns a tuple with the ExtendAdminPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendAdminPermissions

`func (o *BulkEditShareableEntityRequest) SetExtendAdminPermissions(v bool)`

SetExtendAdminPermissions sets ExtendAdminPermissions field to given value.

### HasExtendAdminPermissions

`func (o *BulkEditShareableEntityRequest) HasExtendAdminPermissions() bool`

HasExtendAdminPermissions returns a boolean if a field has been set.

### GetPermissionDetails

`func (o *BulkEditShareableEntityRequest) GetPermissionDetails() PermissionDetails`

GetPermissionDetails returns the PermissionDetails field if non-nil, zero value otherwise.

### GetPermissionDetailsOk

`func (o *BulkEditShareableEntityRequest) GetPermissionDetailsOk() (*PermissionDetails, bool)`

GetPermissionDetailsOk returns a tuple with the PermissionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionDetails

`func (o *BulkEditShareableEntityRequest) SetPermissionDetails(v PermissionDetails)`

SetPermissionDetails sets PermissionDetails field to given value.

### HasPermissionDetails

`func (o *BulkEditShareableEntityRequest) HasPermissionDetails() bool`

HasPermissionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


