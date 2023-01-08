# ApplicationRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGroups** | Pointer to **[]string** | The groups that are granted default access for this application role. As a group&#39;s name can change, use of &#x60;defaultGroupsDetails&#x60; is recommended to identify a groups. | [optional] 
**DefaultGroupsDetails** | Pointer to [**[]GroupName**](GroupName.md) | The groups that are granted default access for this application role. | [optional] 
**Defined** | Pointer to **bool** | Deprecated. | [optional] 
**GroupDetails** | Pointer to [**[]GroupName**](GroupName.md) | The groups associated with the application role. | [optional] 
**Groups** | Pointer to **[]string** | The groups associated with the application role. As a group&#39;s name can change, use of &#x60;groupDetails&#x60; is recommended to identify a groups. | [optional] 
**HasUnlimitedSeats** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** | The key of the application role. | [optional] 
**Name** | Pointer to **string** | The display name of the application role. | [optional] 
**NumberOfSeats** | Pointer to **int32** | The maximum count of users on your license. | [optional] 
**Platform** | Pointer to **bool** | Indicates if the application role belongs to Jira platform (&#x60;jira-core&#x60;). | [optional] 
**RemainingSeats** | Pointer to **int32** | The count of users remaining on your license. | [optional] 
**SelectedByDefault** | Pointer to **bool** | Determines whether this application role should be selected by default on user creation. | [optional] 
**UserCount** | Pointer to **int32** | The number of users counting against your license. | [optional] 
**UserCountDescription** | Pointer to **string** | The [type of users](https://confluence.atlassian.com/x/lRW3Ng) being counted against your license. | [optional] 

## Methods

### NewApplicationRole

`func NewApplicationRole() *ApplicationRole`

NewApplicationRole instantiates a new ApplicationRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRoleWithDefaults

`func NewApplicationRoleWithDefaults() *ApplicationRole`

NewApplicationRoleWithDefaults instantiates a new ApplicationRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGroups

`func (o *ApplicationRole) GetDefaultGroups() []string`

GetDefaultGroups returns the DefaultGroups field if non-nil, zero value otherwise.

### GetDefaultGroupsOk

`func (o *ApplicationRole) GetDefaultGroupsOk() (*[]string, bool)`

GetDefaultGroupsOk returns a tuple with the DefaultGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroups

`func (o *ApplicationRole) SetDefaultGroups(v []string)`

SetDefaultGroups sets DefaultGroups field to given value.

### HasDefaultGroups

`func (o *ApplicationRole) HasDefaultGroups() bool`

HasDefaultGroups returns a boolean if a field has been set.

### GetDefaultGroupsDetails

`func (o *ApplicationRole) GetDefaultGroupsDetails() []GroupName`

GetDefaultGroupsDetails returns the DefaultGroupsDetails field if non-nil, zero value otherwise.

### GetDefaultGroupsDetailsOk

`func (o *ApplicationRole) GetDefaultGroupsDetailsOk() (*[]GroupName, bool)`

GetDefaultGroupsDetailsOk returns a tuple with the DefaultGroupsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupsDetails

`func (o *ApplicationRole) SetDefaultGroupsDetails(v []GroupName)`

SetDefaultGroupsDetails sets DefaultGroupsDetails field to given value.

### HasDefaultGroupsDetails

`func (o *ApplicationRole) HasDefaultGroupsDetails() bool`

HasDefaultGroupsDetails returns a boolean if a field has been set.

### GetDefined

`func (o *ApplicationRole) GetDefined() bool`

GetDefined returns the Defined field if non-nil, zero value otherwise.

### GetDefinedOk

`func (o *ApplicationRole) GetDefinedOk() (*bool, bool)`

GetDefinedOk returns a tuple with the Defined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefined

`func (o *ApplicationRole) SetDefined(v bool)`

SetDefined sets Defined field to given value.

### HasDefined

`func (o *ApplicationRole) HasDefined() bool`

HasDefined returns a boolean if a field has been set.

### GetGroupDetails

`func (o *ApplicationRole) GetGroupDetails() []GroupName`

GetGroupDetails returns the GroupDetails field if non-nil, zero value otherwise.

### GetGroupDetailsOk

`func (o *ApplicationRole) GetGroupDetailsOk() (*[]GroupName, bool)`

GetGroupDetailsOk returns a tuple with the GroupDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDetails

`func (o *ApplicationRole) SetGroupDetails(v []GroupName)`

SetGroupDetails sets GroupDetails field to given value.

### HasGroupDetails

`func (o *ApplicationRole) HasGroupDetails() bool`

HasGroupDetails returns a boolean if a field has been set.

### GetGroups

`func (o *ApplicationRole) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ApplicationRole) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ApplicationRole) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ApplicationRole) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHasUnlimitedSeats

`func (o *ApplicationRole) GetHasUnlimitedSeats() bool`

GetHasUnlimitedSeats returns the HasUnlimitedSeats field if non-nil, zero value otherwise.

### GetHasUnlimitedSeatsOk

`func (o *ApplicationRole) GetHasUnlimitedSeatsOk() (*bool, bool)`

GetHasUnlimitedSeatsOk returns a tuple with the HasUnlimitedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnlimitedSeats

`func (o *ApplicationRole) SetHasUnlimitedSeats(v bool)`

SetHasUnlimitedSeats sets HasUnlimitedSeats field to given value.

### HasHasUnlimitedSeats

`func (o *ApplicationRole) HasHasUnlimitedSeats() bool`

HasHasUnlimitedSeats returns a boolean if a field has been set.

### GetKey

`func (o *ApplicationRole) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationRole) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationRole) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApplicationRole) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfSeats

`func (o *ApplicationRole) GetNumberOfSeats() int32`

GetNumberOfSeats returns the NumberOfSeats field if non-nil, zero value otherwise.

### GetNumberOfSeatsOk

`func (o *ApplicationRole) GetNumberOfSeatsOk() (*int32, bool)`

GetNumberOfSeatsOk returns a tuple with the NumberOfSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSeats

`func (o *ApplicationRole) SetNumberOfSeats(v int32)`

SetNumberOfSeats sets NumberOfSeats field to given value.

### HasNumberOfSeats

`func (o *ApplicationRole) HasNumberOfSeats() bool`

HasNumberOfSeats returns a boolean if a field has been set.

### GetPlatform

`func (o *ApplicationRole) GetPlatform() bool`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApplicationRole) GetPlatformOk() (*bool, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ApplicationRole) SetPlatform(v bool)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ApplicationRole) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRemainingSeats

`func (o *ApplicationRole) GetRemainingSeats() int32`

GetRemainingSeats returns the RemainingSeats field if non-nil, zero value otherwise.

### GetRemainingSeatsOk

`func (o *ApplicationRole) GetRemainingSeatsOk() (*int32, bool)`

GetRemainingSeatsOk returns a tuple with the RemainingSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingSeats

`func (o *ApplicationRole) SetRemainingSeats(v int32)`

SetRemainingSeats sets RemainingSeats field to given value.

### HasRemainingSeats

`func (o *ApplicationRole) HasRemainingSeats() bool`

HasRemainingSeats returns a boolean if a field has been set.

### GetSelectedByDefault

`func (o *ApplicationRole) GetSelectedByDefault() bool`

GetSelectedByDefault returns the SelectedByDefault field if non-nil, zero value otherwise.

### GetSelectedByDefaultOk

`func (o *ApplicationRole) GetSelectedByDefaultOk() (*bool, bool)`

GetSelectedByDefaultOk returns a tuple with the SelectedByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedByDefault

`func (o *ApplicationRole) SetSelectedByDefault(v bool)`

SetSelectedByDefault sets SelectedByDefault field to given value.

### HasSelectedByDefault

`func (o *ApplicationRole) HasSelectedByDefault() bool`

HasSelectedByDefault returns a boolean if a field has been set.

### GetUserCount

`func (o *ApplicationRole) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ApplicationRole) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ApplicationRole) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ApplicationRole) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetUserCountDescription

`func (o *ApplicationRole) GetUserCountDescription() string`

GetUserCountDescription returns the UserCountDescription field if non-nil, zero value otherwise.

### GetUserCountDescriptionOk

`func (o *ApplicationRole) GetUserCountDescriptionOk() (*string, bool)`

GetUserCountDescriptionOk returns a tuple with the UserCountDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCountDescription

`func (o *ApplicationRole) SetUserCountDescription(v string)`

SetUserCountDescription sets UserCountDescription field to given value.

### HasUserCountDescription

`func (o *ApplicationRole) HasUserCountDescription() bool`

HasUserCountDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


