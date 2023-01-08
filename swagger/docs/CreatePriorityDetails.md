# CreatePriorityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The ID for the avatar for the priority. Either the iconUrl or avatarId must be defined, but not both. This parameter is nullable and will become mandatory once the iconUrl parameter is deprecated. | [optional] 
**Description** | Pointer to **NullableString** | The description of the priority. | [optional] 
**IconUrl** | Pointer to **NullableString** | The URL of an icon for the priority. Accepted protocols are HTTP and HTTPS. Built in icons can also be used. Either the iconUrl or avatarId must be defined, but not both. | [optional] 
**Name** | **string** | The name of the priority. Must be unique. | 
**StatusColor** | **string** | The status color of the priority in 3-digit or 6-digit hexadecimal format. | 

## Methods

### NewCreatePriorityDetails

`func NewCreatePriorityDetails(name string, statusColor string, ) *CreatePriorityDetails`

NewCreatePriorityDetails instantiates a new CreatePriorityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePriorityDetailsWithDefaults

`func NewCreatePriorityDetailsWithDefaults() *CreatePriorityDetails`

NewCreatePriorityDetailsWithDefaults instantiates a new CreatePriorityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *CreatePriorityDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CreatePriorityDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CreatePriorityDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *CreatePriorityDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePriorityDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePriorityDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePriorityDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePriorityDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatePriorityDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatePriorityDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *CreatePriorityDetails) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreatePriorityDetails) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreatePriorityDetails) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreatePriorityDetails) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *CreatePriorityDetails) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *CreatePriorityDetails) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetName

`func (o *CreatePriorityDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePriorityDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePriorityDetails) SetName(v string)`

SetName sets Name field to given value.


### GetStatusColor

`func (o *CreatePriorityDetails) GetStatusColor() string`

GetStatusColor returns the StatusColor field if non-nil, zero value otherwise.

### GetStatusColorOk

`func (o *CreatePriorityDetails) GetStatusColorOk() (*string, bool)`

GetStatusColorOk returns a tuple with the StatusColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusColor

`func (o *CreatePriorityDetails) SetStatusColor(v string)`

SetStatusColor sets StatusColor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


