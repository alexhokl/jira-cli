# UpdatePriorityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The ID for the avatar for the priority. This parameter is nullable and both iconUrl and avatarId cannot be defined. | [optional] 
**Description** | Pointer to **NullableString** | The description of the priority. | [optional] 
**IconUrl** | Pointer to **NullableString** | The URL of an icon for the priority. Accepted protocols are HTTP and HTTPS. Built in icons can also be used. Both iconUrl and avatarId cannot be defined. | [optional] 
**Name** | Pointer to **NullableString** | The name of the priority. Must be unique. | [optional] 
**StatusColor** | Pointer to **NullableString** | The status color of the priority in 3-digit or 6-digit hexadecimal format. | [optional] 

## Methods

### NewUpdatePriorityDetails

`func NewUpdatePriorityDetails() *UpdatePriorityDetails`

NewUpdatePriorityDetails instantiates a new UpdatePriorityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriorityDetailsWithDefaults

`func NewUpdatePriorityDetailsWithDefaults() *UpdatePriorityDetails`

NewUpdatePriorityDetailsWithDefaults instantiates a new UpdatePriorityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *UpdatePriorityDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *UpdatePriorityDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *UpdatePriorityDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *UpdatePriorityDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePriorityDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePriorityDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePriorityDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePriorityDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePriorityDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePriorityDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *UpdatePriorityDetails) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *UpdatePriorityDetails) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *UpdatePriorityDetails) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *UpdatePriorityDetails) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *UpdatePriorityDetails) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *UpdatePriorityDetails) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetName

`func (o *UpdatePriorityDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePriorityDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePriorityDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePriorityDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatePriorityDetails) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatePriorityDetails) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatusColor

`func (o *UpdatePriorityDetails) GetStatusColor() string`

GetStatusColor returns the StatusColor field if non-nil, zero value otherwise.

### GetStatusColorOk

`func (o *UpdatePriorityDetails) GetStatusColorOk() (*string, bool)`

GetStatusColorOk returns a tuple with the StatusColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusColor

`func (o *UpdatePriorityDetails) SetStatusColor(v string)`

SetStatusColor sets StatusColor field to given value.

### HasStatusColor

`func (o *UpdatePriorityDetails) HasStatusColor() bool`

HasStatusColor returns a boolean if a field has been set.

### SetStatusColorNil

`func (o *UpdatePriorityDetails) SetStatusColorNil(b bool)`

 SetStatusColorNil sets the value for StatusColor to be an explicit nil

### UnsetStatusColor
`func (o *UpdatePriorityDetails) UnsetStatusColor()`

UnsetStatusColor ensures that no value is present for StatusColor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


