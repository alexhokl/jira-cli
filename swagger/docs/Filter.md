# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproximateLastUsed** | Pointer to **time.Time** | \\[Experimental\\] Approximate last used time. Returns the date and time when the filter was last used. Returns &#x60;null&#x60; if the filter hasn&#39;t been used after tracking was enabled. For performance reasons, timestamps aren&#39;t updated in real time and therefore may not be exactly accurate. | [optional] [readonly] 
**Description** | Pointer to **string** | A description of the filter. | [optional] 
**EditPermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The groups and projects that can edit the filter. | [optional] 
**Favourite** | Pointer to **bool** | Whether the filter is selected as a favorite. | [optional] 
**FavouritedCount** | Pointer to **int64** | The count of how many users have selected this filter as a favorite, including the filter owner. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the filter. | [optional] [readonly] 
**Jql** | Pointer to **string** | The JQL query for the filter. For example, *project &#x3D; SSP AND issuetype &#x3D; Bug*. | [optional] 
**Name** | **string** | The name of the filter. Must be unique. | 
**Owner** | Pointer to [**User**](User.md) | The user who owns the filter. This is defaulted to the creator of the filter, however Jira administrators can change the owner of a shared filter in the admin settings. | [optional] [readonly] 
**SearchUrl** | Pointer to **string** | A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter&#39;s JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql&#x3D;project+%3D+SSP+AND+issuetype+%3D+Bug*. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the filter. | [optional] [readonly] 
**SharePermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The groups and projects that the filter is shared with. | [optional] 
**SharedUsers** | Pointer to [**UserList**](UserList.md) | A paginated list of the users that the filter is shared with. This includes users that are members of the groups or can browse the projects that the filter is shared with. | [optional] [readonly] 
**Subscriptions** | Pointer to [**FilterSubscriptionsList**](FilterSubscriptionsList.md) | A paginated list of the users that are subscribed to the filter. | [optional] [readonly] 
**ViewUrl** | Pointer to **string** | A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter&#x3D;10100*. | [optional] [readonly] 

## Methods

### NewFilter

`func NewFilter(name string, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximateLastUsed

`func (o *Filter) GetApproximateLastUsed() time.Time`

GetApproximateLastUsed returns the ApproximateLastUsed field if non-nil, zero value otherwise.

### GetApproximateLastUsedOk

`func (o *Filter) GetApproximateLastUsedOk() (*time.Time, bool)`

GetApproximateLastUsedOk returns a tuple with the ApproximateLastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateLastUsed

`func (o *Filter) SetApproximateLastUsed(v time.Time)`

SetApproximateLastUsed sets ApproximateLastUsed field to given value.

### HasApproximateLastUsed

`func (o *Filter) HasApproximateLastUsed() bool`

HasApproximateLastUsed returns a boolean if a field has been set.

### GetDescription

`func (o *Filter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Filter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Filter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Filter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditPermissions

`func (o *Filter) GetEditPermissions() []SharePermission`

GetEditPermissions returns the EditPermissions field if non-nil, zero value otherwise.

### GetEditPermissionsOk

`func (o *Filter) GetEditPermissionsOk() (*[]SharePermission, bool)`

GetEditPermissionsOk returns a tuple with the EditPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermissions

`func (o *Filter) SetEditPermissions(v []SharePermission)`

SetEditPermissions sets EditPermissions field to given value.

### HasEditPermissions

`func (o *Filter) HasEditPermissions() bool`

HasEditPermissions returns a boolean if a field has been set.

### GetFavourite

`func (o *Filter) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *Filter) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *Filter) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *Filter) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetFavouritedCount

`func (o *Filter) GetFavouritedCount() int64`

GetFavouritedCount returns the FavouritedCount field if non-nil, zero value otherwise.

### GetFavouritedCountOk

`func (o *Filter) GetFavouritedCountOk() (*int64, bool)`

GetFavouritedCountOk returns a tuple with the FavouritedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavouritedCount

`func (o *Filter) SetFavouritedCount(v int64)`

SetFavouritedCount sets FavouritedCount field to given value.

### HasFavouritedCount

`func (o *Filter) HasFavouritedCount() bool`

HasFavouritedCount returns a boolean if a field has been set.

### GetId

`func (o *Filter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Filter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Filter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Filter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJql

`func (o *Filter) GetJql() string`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *Filter) GetJqlOk() (*string, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *Filter) SetJql(v string)`

SetJql sets Jql field to given value.

### HasJql

`func (o *Filter) HasJql() bool`

HasJql returns a boolean if a field has been set.

### GetName

`func (o *Filter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filter) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *Filter) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Filter) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Filter) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Filter) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSearchUrl

`func (o *Filter) GetSearchUrl() string`

GetSearchUrl returns the SearchUrl field if non-nil, zero value otherwise.

### GetSearchUrlOk

`func (o *Filter) GetSearchUrlOk() (*string, bool)`

GetSearchUrlOk returns a tuple with the SearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchUrl

`func (o *Filter) SetSearchUrl(v string)`

SetSearchUrl sets SearchUrl field to given value.

### HasSearchUrl

`func (o *Filter) HasSearchUrl() bool`

HasSearchUrl returns a boolean if a field has been set.

### GetSelf

`func (o *Filter) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Filter) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Filter) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Filter) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSharePermissions

`func (o *Filter) GetSharePermissions() []SharePermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *Filter) GetSharePermissionsOk() (*[]SharePermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *Filter) SetSharePermissions(v []SharePermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *Filter) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetSharedUsers

`func (o *Filter) GetSharedUsers() UserList`

GetSharedUsers returns the SharedUsers field if non-nil, zero value otherwise.

### GetSharedUsersOk

`func (o *Filter) GetSharedUsersOk() (*UserList, bool)`

GetSharedUsersOk returns a tuple with the SharedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedUsers

`func (o *Filter) SetSharedUsers(v UserList)`

SetSharedUsers sets SharedUsers field to given value.

### HasSharedUsers

`func (o *Filter) HasSharedUsers() bool`

HasSharedUsers returns a boolean if a field has been set.

### GetSubscriptions

`func (o *Filter) GetSubscriptions() FilterSubscriptionsList`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *Filter) GetSubscriptionsOk() (*FilterSubscriptionsList, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *Filter) SetSubscriptions(v FilterSubscriptionsList)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *Filter) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetViewUrl

`func (o *Filter) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *Filter) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *Filter) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *Filter) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


