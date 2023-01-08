# DashboardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the dashboard. | [optional] 
**EditPermissions** | [**[]SharePermission**](SharePermission.md) | The edit permissions for the dashboard. | 
**Name** | **string** | The name of the dashboard. | 
**SharePermissions** | [**[]SharePermission**](SharePermission.md) | The share permissions for the dashboard. | 

## Methods

### NewDashboardDetails

`func NewDashboardDetails(editPermissions []SharePermission, name string, sharePermissions []SharePermission, ) *DashboardDetails`

NewDashboardDetails instantiates a new DashboardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardDetailsWithDefaults

`func NewDashboardDetailsWithDefaults() *DashboardDetails`

NewDashboardDetailsWithDefaults instantiates a new DashboardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DashboardDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditPermissions

`func (o *DashboardDetails) GetEditPermissions() []SharePermission`

GetEditPermissions returns the EditPermissions field if non-nil, zero value otherwise.

### GetEditPermissionsOk

`func (o *DashboardDetails) GetEditPermissionsOk() (*[]SharePermission, bool)`

GetEditPermissionsOk returns a tuple with the EditPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermissions

`func (o *DashboardDetails) SetEditPermissions(v []SharePermission)`

SetEditPermissions sets EditPermissions field to given value.


### GetName

`func (o *DashboardDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardDetails) SetName(v string)`

SetName sets Name field to given value.


### GetSharePermissions

`func (o *DashboardDetails) GetSharePermissions() []SharePermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *DashboardDetails) GetSharePermissionsOk() (*[]SharePermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *DashboardDetails) SetSharePermissions(v []SharePermission)`

SetSharePermissions sets SharePermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


