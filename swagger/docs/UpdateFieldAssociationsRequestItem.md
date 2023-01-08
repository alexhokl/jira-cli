# UpdateFieldAssociationsRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictedToWorkTypes** | Pointer to **[]int64** | (optional) Work types to restrict field to. Replaces any existing work type associations for the field. If not provided, the field is associated to any work types. | [optional] 
**SchemeIds** | **[]int64** | Scheme IDs to associate field with | 

## Methods

### NewUpdateFieldAssociationsRequestItem

`func NewUpdateFieldAssociationsRequestItem(schemeIds []int64, ) *UpdateFieldAssociationsRequestItem`

NewUpdateFieldAssociationsRequestItem instantiates a new UpdateFieldAssociationsRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFieldAssociationsRequestItemWithDefaults

`func NewUpdateFieldAssociationsRequestItemWithDefaults() *UpdateFieldAssociationsRequestItem`

NewUpdateFieldAssociationsRequestItemWithDefaults instantiates a new UpdateFieldAssociationsRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictedToWorkTypes

`func (o *UpdateFieldAssociationsRequestItem) GetRestrictedToWorkTypes() []int64`

GetRestrictedToWorkTypes returns the RestrictedToWorkTypes field if non-nil, zero value otherwise.

### GetRestrictedToWorkTypesOk

`func (o *UpdateFieldAssociationsRequestItem) GetRestrictedToWorkTypesOk() (*[]int64, bool)`

GetRestrictedToWorkTypesOk returns a tuple with the RestrictedToWorkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedToWorkTypes

`func (o *UpdateFieldAssociationsRequestItem) SetRestrictedToWorkTypes(v []int64)`

SetRestrictedToWorkTypes sets RestrictedToWorkTypes field to given value.

### HasRestrictedToWorkTypes

`func (o *UpdateFieldAssociationsRequestItem) HasRestrictedToWorkTypes() bool`

HasRestrictedToWorkTypes returns a boolean if a field has been set.

### GetSchemeIds

`func (o *UpdateFieldAssociationsRequestItem) GetSchemeIds() []int64`

GetSchemeIds returns the SchemeIds field if non-nil, zero value otherwise.

### GetSchemeIdsOk

`func (o *UpdateFieldAssociationsRequestItem) GetSchemeIdsOk() (*[]int64, bool)`

GetSchemeIdsOk returns a tuple with the SchemeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeIds

`func (o *UpdateFieldAssociationsRequestItem) SetSchemeIds(v []int64)`

SetSchemeIds sets SchemeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


