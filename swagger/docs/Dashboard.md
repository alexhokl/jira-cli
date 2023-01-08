# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticRefreshMs** | Pointer to **int32** | The automatic refresh interval for the dashboard in milliseconds. | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**EditPermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The details of any edit share permissions for the dashboard. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the dashboard. | [optional] [readonly] 
**IsFavourite** | Pointer to **bool** | Whether the dashboard is selected as a favorite by the user. | [optional] [readonly] 
**IsWritable** | Pointer to **bool** | Whether the current user has permission to edit the dashboard. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the dashboard. | [optional] [readonly] 
**Owner** | Pointer to [**UserBean**](UserBean.md) | The owner of the dashboard. | [optional] [readonly] 
**Popularity** | Pointer to **int64** | The number of users who have this dashboard as a favorite. | [optional] [readonly] 
**Rank** | Pointer to **int32** | The rank of this dashboard. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these dashboard details. | [optional] [readonly] 
**SharePermissions** | Pointer to [**[]SharePermission**](SharePermission.md) | The details of any view share permissions for the dashboard. | [optional] [readonly] 
**SystemDashboard** | Pointer to **bool** | Whether the current dashboard is system dashboard. | [optional] [readonly] 
**View** | Pointer to **string** | The URL of the dashboard. | [optional] [readonly] 

## Methods

### NewDashboard

`func NewDashboard() *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticRefreshMs

`func (o *Dashboard) GetAutomaticRefreshMs() int32`

GetAutomaticRefreshMs returns the AutomaticRefreshMs field if non-nil, zero value otherwise.

### GetAutomaticRefreshMsOk

`func (o *Dashboard) GetAutomaticRefreshMsOk() (*int32, bool)`

GetAutomaticRefreshMsOk returns a tuple with the AutomaticRefreshMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshMs

`func (o *Dashboard) SetAutomaticRefreshMs(v int32)`

SetAutomaticRefreshMs sets AutomaticRefreshMs field to given value.

### HasAutomaticRefreshMs

`func (o *Dashboard) HasAutomaticRefreshMs() bool`

HasAutomaticRefreshMs returns a boolean if a field has been set.

### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditPermissions

`func (o *Dashboard) GetEditPermissions() []SharePermission`

GetEditPermissions returns the EditPermissions field if non-nil, zero value otherwise.

### GetEditPermissionsOk

`func (o *Dashboard) GetEditPermissionsOk() (*[]SharePermission, bool)`

GetEditPermissionsOk returns a tuple with the EditPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermissions

`func (o *Dashboard) SetEditPermissions(v []SharePermission)`

SetEditPermissions sets EditPermissions field to given value.

### HasEditPermissions

`func (o *Dashboard) HasEditPermissions() bool`

HasEditPermissions returns a boolean if a field has been set.

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFavourite

`func (o *Dashboard) GetIsFavourite() bool`

GetIsFavourite returns the IsFavourite field if non-nil, zero value otherwise.

### GetIsFavouriteOk

`func (o *Dashboard) GetIsFavouriteOk() (*bool, bool)`

GetIsFavouriteOk returns a tuple with the IsFavourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavourite

`func (o *Dashboard) SetIsFavourite(v bool)`

SetIsFavourite sets IsFavourite field to given value.

### HasIsFavourite

`func (o *Dashboard) HasIsFavourite() bool`

HasIsFavourite returns a boolean if a field has been set.

### GetIsWritable

`func (o *Dashboard) GetIsWritable() bool`

GetIsWritable returns the IsWritable field if non-nil, zero value otherwise.

### GetIsWritableOk

`func (o *Dashboard) GetIsWritableOk() (*bool, bool)`

GetIsWritableOk returns a tuple with the IsWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWritable

`func (o *Dashboard) SetIsWritable(v bool)`

SetIsWritable sets IsWritable field to given value.

### HasIsWritable

`func (o *Dashboard) HasIsWritable() bool`

HasIsWritable returns a boolean if a field has been set.

### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Dashboard) GetOwner() UserBean`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Dashboard) GetOwnerOk() (*UserBean, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Dashboard) SetOwner(v UserBean)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Dashboard) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPopularity

`func (o *Dashboard) GetPopularity() int64`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *Dashboard) GetPopularityOk() (*int64, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *Dashboard) SetPopularity(v int64)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *Dashboard) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetRank

`func (o *Dashboard) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Dashboard) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Dashboard) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Dashboard) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSelf

`func (o *Dashboard) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Dashboard) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Dashboard) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Dashboard) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSharePermissions

`func (o *Dashboard) GetSharePermissions() []SharePermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *Dashboard) GetSharePermissionsOk() (*[]SharePermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *Dashboard) SetSharePermissions(v []SharePermission)`

SetSharePermissions sets SharePermissions field to given value.

### HasSharePermissions

`func (o *Dashboard) HasSharePermissions() bool`

HasSharePermissions returns a boolean if a field has been set.

### GetSystemDashboard

`func (o *Dashboard) GetSystemDashboard() bool`

GetSystemDashboard returns the SystemDashboard field if non-nil, zero value otherwise.

### GetSystemDashboardOk

`func (o *Dashboard) GetSystemDashboardOk() (*bool, bool)`

GetSystemDashboardOk returns a tuple with the SystemDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDashboard

`func (o *Dashboard) SetSystemDashboard(v bool)`

SetSystemDashboard sets SystemDashboard field to given value.

### HasSystemDashboard

`func (o *Dashboard) HasSystemDashboard() bool`

HasSystemDashboard returns a boolean if a field has been set.

### GetView

`func (o *Dashboard) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Dashboard) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Dashboard) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Dashboard) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


