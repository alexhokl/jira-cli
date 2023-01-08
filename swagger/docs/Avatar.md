# Avatar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | The file name of the avatar icon. Returned for system avatars. | [optional] [readonly] 
**Id** | **string** | The ID of the avatar. | 
**IsDeletable** | Pointer to **bool** | Whether the avatar can be deleted. | [optional] [readonly] 
**IsSelected** | Pointer to **bool** | Whether the avatar is used in Jira. For example, shown as a project&#39;s avatar. | [optional] [readonly] 
**IsSystemAvatar** | Pointer to **bool** | Whether the avatar is a system avatar. | [optional] [readonly] 
**Owner** | Pointer to **string** | The owner of the avatar. For a system avatar the owner is null (and nothing is returned). For non-system avatars this is the appropriate identifier, such as the ID for a project or the account ID for a user. | [optional] [readonly] 
**Urls** | Pointer to **map[string]string** | The list of avatar icon URLs. | [optional] [readonly] 

## Methods

### NewAvatar

`func NewAvatar(id string, ) *Avatar`

NewAvatar instantiates a new Avatar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvatarWithDefaults

`func NewAvatarWithDefaults() *Avatar`

NewAvatarWithDefaults instantiates a new Avatar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *Avatar) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Avatar) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Avatar) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Avatar) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *Avatar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Avatar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Avatar) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeletable

`func (o *Avatar) GetIsDeletable() bool`

GetIsDeletable returns the IsDeletable field if non-nil, zero value otherwise.

### GetIsDeletableOk

`func (o *Avatar) GetIsDeletableOk() (*bool, bool)`

GetIsDeletableOk returns a tuple with the IsDeletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletable

`func (o *Avatar) SetIsDeletable(v bool)`

SetIsDeletable sets IsDeletable field to given value.

### HasIsDeletable

`func (o *Avatar) HasIsDeletable() bool`

HasIsDeletable returns a boolean if a field has been set.

### GetIsSelected

`func (o *Avatar) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *Avatar) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *Avatar) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.

### HasIsSelected

`func (o *Avatar) HasIsSelected() bool`

HasIsSelected returns a boolean if a field has been set.

### GetIsSystemAvatar

`func (o *Avatar) GetIsSystemAvatar() bool`

GetIsSystemAvatar returns the IsSystemAvatar field if non-nil, zero value otherwise.

### GetIsSystemAvatarOk

`func (o *Avatar) GetIsSystemAvatarOk() (*bool, bool)`

GetIsSystemAvatarOk returns a tuple with the IsSystemAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAvatar

`func (o *Avatar) SetIsSystemAvatar(v bool)`

SetIsSystemAvatar sets IsSystemAvatar field to given value.

### HasIsSystemAvatar

`func (o *Avatar) HasIsSystemAvatar() bool`

HasIsSystemAvatar returns a boolean if a field has been set.

### GetOwner

`func (o *Avatar) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Avatar) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Avatar) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Avatar) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUrls

`func (o *Avatar) GetUrls() map[string]string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Avatar) GetUrlsOk() (*map[string]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Avatar) SetUrls(v map[string]string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *Avatar) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


