# NotificationRecipientsRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIds** | Pointer to **[]string** | List of groupId memberships required to receive the notification. | [optional] 
**Groups** | Pointer to [**[]GroupName**](GroupName.md) | List of group memberships required to receive the notification. | [optional] 
**Permissions** | Pointer to [**[]RestrictedPermission**](RestrictedPermission.md) | List of permissions required to receive the notification. | [optional] 

## Methods

### NewNotificationRecipientsRestrictions

`func NewNotificationRecipientsRestrictions() *NotificationRecipientsRestrictions`

NewNotificationRecipientsRestrictions instantiates a new NotificationRecipientsRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRecipientsRestrictionsWithDefaults

`func NewNotificationRecipientsRestrictionsWithDefaults() *NotificationRecipientsRestrictions`

NewNotificationRecipientsRestrictionsWithDefaults instantiates a new NotificationRecipientsRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIds

`func (o *NotificationRecipientsRestrictions) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *NotificationRecipientsRestrictions) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *NotificationRecipientsRestrictions) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *NotificationRecipientsRestrictions) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetGroups

`func (o *NotificationRecipientsRestrictions) GetGroups() []GroupName`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *NotificationRecipientsRestrictions) GetGroupsOk() (*[]GroupName, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *NotificationRecipientsRestrictions) SetGroups(v []GroupName)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *NotificationRecipientsRestrictions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPermissions

`func (o *NotificationRecipientsRestrictions) GetPermissions() []RestrictedPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NotificationRecipientsRestrictions) GetPermissionsOk() (*[]RestrictedPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NotificationRecipientsRestrictions) SetPermissions(v []RestrictedPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NotificationRecipientsRestrictions) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


