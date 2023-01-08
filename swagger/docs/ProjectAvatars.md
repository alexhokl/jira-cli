# ProjectAvatars

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**[]Avatar**](Avatar.md) | List of avatars added to Jira. These avatars may be deleted. | [optional] [readonly] 
**System** | Pointer to [**[]Avatar**](Avatar.md) | List of avatars included with Jira. These avatars cannot be deleted. | [optional] [readonly] 

## Methods

### NewProjectAvatars

`func NewProjectAvatars() *ProjectAvatars`

NewProjectAvatars instantiates a new ProjectAvatars object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAvatarsWithDefaults

`func NewProjectAvatarsWithDefaults() *ProjectAvatars`

NewProjectAvatarsWithDefaults instantiates a new ProjectAvatars object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *ProjectAvatars) GetCustom() []Avatar`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ProjectAvatars) GetCustomOk() (*[]Avatar, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ProjectAvatars) SetCustom(v []Avatar)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ProjectAvatars) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetSystem

`func (o *ProjectAvatars) GetSystem() []Avatar`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ProjectAvatars) GetSystemOk() (*[]Avatar, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ProjectAvatars) SetSystem(v []Avatar)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ProjectAvatars) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


