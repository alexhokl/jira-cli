# FoundGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** | Avatar url for the group/team if present. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*. | [optional] 
**Html** | Pointer to **string** | The group name with the matched query string highlighted with the HTML bold tag. | [optional] 
**Labels** | Pointer to [**[]GroupLabel**](GroupLabel.md) |  | [optional] 
**ManagedBy** | Pointer to **string** | Describes who/how the team is managed. The possible values are   \\* external - when team is synced from an external directory like SCIM or HRIS, and team members cannot be modified.   \\* admins - when a team is managed by an admin (team members can only be modified by admins).   \\* team-members - managed by existing team members, new members need to be invited to join.   \\* open - anyone can join or modify this team. | [optional] 
**Name** | Pointer to **string** | The name of the group. The name of a group is mutable, to reliably identify a group use &#x60;&#x60;groupId&#x60;.&#x60; | [optional] 
**UsageType** | Pointer to **string** | Describes the type of group. The possible values are   \\* team-collaboration - A platform team managed in people directory.   \\* userbase-group - a group of users created in adminhub.   \\* admin-oversight - currently unused. | [optional] 

## Methods

### NewFoundGroup

`func NewFoundGroup() *FoundGroup`

NewFoundGroup instantiates a new FoundGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundGroupWithDefaults

`func NewFoundGroupWithDefaults() *FoundGroup`

NewFoundGroupWithDefaults instantiates a new FoundGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *FoundGroup) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *FoundGroup) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *FoundGroup) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *FoundGroup) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetGroupId

`func (o *FoundGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FoundGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FoundGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *FoundGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHtml

`func (o *FoundGroup) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *FoundGroup) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *FoundGroup) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *FoundGroup) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetLabels

`func (o *FoundGroup) GetLabels() []GroupLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FoundGroup) GetLabelsOk() (*[]GroupLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FoundGroup) SetLabels(v []GroupLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FoundGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetManagedBy

`func (o *FoundGroup) GetManagedBy() string`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *FoundGroup) GetManagedByOk() (*string, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *FoundGroup) SetManagedBy(v string)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *FoundGroup) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.

### GetName

`func (o *FoundGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoundGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoundGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FoundGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsageType

`func (o *FoundGroup) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *FoundGroup) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *FoundGroup) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *FoundGroup) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


