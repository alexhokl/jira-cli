# Priority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The avatarId of the avatar for the issue priority. This parameter is nullable and when set, this avatar references the universal avatar APIs. | [optional] 
**Description** | Pointer to **string** | The description of the issue priority. | [optional] 
**IconUrl** | Pointer to **string** | The URL of the icon for the issue priority. | [optional] 
**Id** | Pointer to **string** | The ID of the issue priority. | [optional] 
**IsDefault** | Pointer to **bool** | Whether this priority is the default. | [optional] 
**Name** | Pointer to **string** | The name of the issue priority. | [optional] 
**Schemes** | Pointer to [**ExpandPrioritySchemePage**](ExpandPrioritySchemePage.md) | Priority schemes associated with the issue priority. | [optional] 
**Self** | Pointer to **string** | The URL of the issue priority. | [optional] 
**StatusColor** | Pointer to **string** | The color used to indicate the issue priority. | [optional] 

## Methods

### NewPriority

`func NewPriority() *Priority`

NewPriority instantiates a new Priority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityWithDefaults

`func NewPriorityWithDefaults() *Priority`

NewPriorityWithDefaults instantiates a new Priority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *Priority) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *Priority) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *Priority) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *Priority) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *Priority) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Priority) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Priority) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Priority) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *Priority) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Priority) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Priority) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Priority) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *Priority) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Priority) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Priority) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Priority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *Priority) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Priority) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Priority) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Priority) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *Priority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Priority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Priority) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Priority) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemes

`func (o *Priority) GetSchemes() ExpandPrioritySchemePage`

GetSchemes returns the Schemes field if non-nil, zero value otherwise.

### GetSchemesOk

`func (o *Priority) GetSchemesOk() (*ExpandPrioritySchemePage, bool)`

GetSchemesOk returns a tuple with the Schemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemes

`func (o *Priority) SetSchemes(v ExpandPrioritySchemePage)`

SetSchemes sets Schemes field to given value.

### HasSchemes

`func (o *Priority) HasSchemes() bool`

HasSchemes returns a boolean if a field has been set.

### GetSelf

`func (o *Priority) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Priority) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Priority) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Priority) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStatusColor

`func (o *Priority) GetStatusColor() string`

GetStatusColor returns the StatusColor field if non-nil, zero value otherwise.

### GetStatusColorOk

`func (o *Priority) GetStatusColorOk() (*string, bool)`

GetStatusColorOk returns a tuple with the StatusColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusColor

`func (o *Priority) SetStatusColor(v string)`

SetStatusColor sets StatusColor field to given value.

### HasStatusColor

`func (o *Priority) HasStatusColor() bool`

HasStatusColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


