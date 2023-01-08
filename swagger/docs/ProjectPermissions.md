# ProjectPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanEdit** | Pointer to **bool** | Whether the logged user can edit the project. | [optional] [readonly] 

## Methods

### NewProjectPermissions

`func NewProjectPermissions() *ProjectPermissions`

NewProjectPermissions instantiates a new ProjectPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectPermissionsWithDefaults

`func NewProjectPermissionsWithDefaults() *ProjectPermissions`

NewProjectPermissionsWithDefaults instantiates a new ProjectPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanEdit

`func (o *ProjectPermissions) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *ProjectPermissions) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *ProjectPermissions) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *ProjectPermissions) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


