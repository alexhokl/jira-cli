# FilterDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproximateLastUsed** | Pointer to **time.Time** | \\[Experimental\\] Approximate last used time. Returns the date and time when the filter was last used. Returns &#x60;null&#x60; if the filter hasn&#39;t been used after tracking was enabled. For performance reasons, timestamps aren&#39;t updated in real time and therefore may not be exactly accurate. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the filter. | [optional] 
**EditPermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The groups and projects that can edit the filter. This can be specified when updating a filter, but not when creating a filter. | [optional] 
**Expand** | Pointer to **string** | Expand options that include additional filter details in the response. | [optional] [readonly] 
**Favourite** | Pointer to **bool** | Whether the filter is selected as a favorite by any users, not including the filter owner. | [optional] [readonly] 
**FavouritedCount** | Pointer to **int64** | The count of how many users have selected this filter as a favorite, including the filter owner. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the filter. | [optional] [readonly] 
**Jql** | Pointer to **string** | The JQL query for the filter. For example, *project &#x3D; SSP AND issuetype &#x3D; Bug*. | [optional] [readonly] 
**Name** | **string** | The name of the filter. | 
**Owner** | Pointer to [**User**](User.md) | The user who owns the filter. Defaults to the creator of the filter, however, Jira administrators can change the owner of a shared filter in the admin settings. | [optional] [readonly] 
**SearchUrl** | Pointer to **string** | A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter&#39;s JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql&#x3D;project+%3D+SSP+AND+issuetype+%3D+Bug*. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the filter. | [optional] [readonly] 
**SharePermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The groups and projects that the filter is shared with. This can be specified when updating a filter, but not when creating a filter. | [optional] 
**Subscriptions** | Pointer to [**[]FilterSubscription**](FilterSubscription.md) | The users that are subscribed to the filter. | [optional] [readonly] 
**ViewUrl** | Pointer to **string** | A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter&#x3D;10100*. | [optional] [readonly] 

## Methods

### NewFilterDetails

`func NewFilterDetails(name string, ) *FilterDetails`

NewFilterDetails instantiates a new FilterDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDetailsWithDefaults

`func NewFilterDetailsWithDefaults() *FilterDetails`

NewFilterDetailsWithDefaults instantiates a new FilterDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproximateLastUsed

`func (o *FilterDetails) GetApproximateLastUsed() time.Time`

GetApproximateLastUsed returns the ApproximateLastUsed field if non-nil, zero value otherwise.

### GetApproximateLastUsedOk

`func (o *FilterDetails) GetApproximateLastUsedOk() (*time.Time, bool)`

GetApproximateLastUsedOk returns a tuple with the ApproximateLastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateLastUsed

`func (o *FilterDetails) SetApproximateLastUsed(v time.Time)`

SetApproximateLastUsed sets ApproximateLastUsed field to given value.

### HasApproximateLastUsed

`func (o *FilterDetails) HasApproximateLastUsed() bool`

HasApproximateLastUsed returns a boolean if a field has been set.

### GetDescription

`func (o *FilterDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FilterDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditPermissions

`func (o *FilterDetails) GetEditPermissions() []SharePermission`

GetEditPermissions returns the EditPermissions field if non-nil, zero value otherwise.

### GetEditPermissionsOk

`func (o *FilterDetails) GetEditPermissionsOk() (*[]SharePermission, bool)`

GetEditPermissionsOk returns a tuple with the EditPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermissions

`func (o *FilterDetails) SetEditPermissions(v []SharePermission)`

SetEditPermissions sets EditPermissions field to given value.

### HasEditPermissions

`func (o *FilterDetails) HasEditPermissions() bool`

HasEditPermissions returns a boolean if a field has been set.

### GetExpand

`func (o *FilterDetails) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *FilterDetails) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *FilterDetails) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *FilterDetails) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFavourite

`func (o *FilterDetails) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *FilterDetails) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *FilterDetails) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *FilterDetails) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetFavouritedCount

`func (o *FilterDetails) GetFavouritedCount() int64`

GetFavouritedCount returns the FavouritedCount field if non-nil, zero value otherwise.

### GetFavouritedCountOk

`func (o *FilterDetails) GetFavouritedCountOk() (*int64, bool)`

GetFavouritedCountOk returns a tuple with the FavouritedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavouritedCount

`func (o *FilterDetails) SetFavouritedCount(v int64)`

SetFavouritedCount sets FavouritedCount field to given value.

### HasFavouritedCount

`func (o *FilterDetails) HasFavouritedCount() bool`

HasFavouritedCount returns a boolean if a field has been set.

### GetId

`func (o *FilterDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FilterDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJql

`func (o *FilterDetails) GetJql() string`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *FilterDetails) GetJqlOk() (*string, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *FilterDetails) SetJql(v string)`

SetJql sets Jql field to given value.

### HasJql

`func (o *FilterDetails) HasJql() bool`

HasJql returns a boolean if a field has been set.

### GetName

`func (o *FilterDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterDetails) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *FilterDetails) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilterDetails) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilterDetails) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FilterDetails) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSearchUrl

`func (o *FilterDetails) GetSearchUrl() string`

GetSearchUrl returns the SearchUrl field if non-nil, zero value otherwise.

### GetSearchUrlOk

`func (o *FilterDetails) GetSearchUrlOk() (*string, bool)`

GetSearchUrlOk returns a tuple with the SearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchUrl

`func (o *FilterDetails) SetSearchUrl(v string)`

SetSearchUrl sets SearchUrl field to given value.

### HasSearchUrl

`func (o *FilterDetails) HasSearchUrl() bool`

HasSearchUrl returns a boolean if a field has been set.

### GetSelf

`func (o *FilterDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *FilterDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *FilterDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *FilterDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSharePermissions

`func (o *FilterDetails) GetSharePermissions() []SharePermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *FilterDetails) GetSharePermissionsOk() (*[]SharePermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *FilterDetails) SetSharePermissions(v []SharePermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *FilterDetails) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetSubscriptions

`func (o *FilterDetails) GetSubscriptions() []FilterSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *FilterDetails) GetSubscriptionsOk() (*[]FilterSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *FilterDetails) SetSubscriptions(v []FilterSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *FilterDetails) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetViewUrl

`func (o *FilterDetails) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *FilterDetails) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *FilterDetails) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *FilterDetails) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


