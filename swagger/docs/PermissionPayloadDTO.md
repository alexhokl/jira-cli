# PermissionPayloadDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddAddonRole** | Pointer to **bool** | Configuration to generate addon role. Default is false if null. Only applies to GLOBAL-scoped permission scheme | [optional] 
**Description** | Pointer to **string** | The description of the permission scheme | [optional] 
**Grants** | Pointer to [**[]PermissionGrantDTO**](PermissionGrantDTO.md) | List of permission grants | [optional] 
**Name** | Pointer to **string** | The name of the permission scheme | [optional] 
**OnConflict** | Pointer to **string** | The strategy to use when there is a conflict with an existing permission scheme. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters; NEW - If the entity exist, try and create a new one with a different name | [optional] [default to "FAIL"]
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewPermissionPayloadDTO

`func NewPermissionPayloadDTO() *PermissionPayloadDTO`

NewPermissionPayloadDTO instantiates a new PermissionPayloadDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionPayloadDTOWithDefaults

`func NewPermissionPayloadDTOWithDefaults() *PermissionPayloadDTO`

NewPermissionPayloadDTOWithDefaults instantiates a new PermissionPayloadDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddAddonRole

`func (o *PermissionPayloadDTO) GetAddAddonRole() bool`

GetAddAddonRole returns the AddAddonRole field if non-nil, zero value otherwise.

### GetAddAddonRoleOk

`func (o *PermissionPayloadDTO) GetAddAddonRoleOk() (*bool, bool)`

GetAddAddonRoleOk returns a tuple with the AddAddonRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAddonRole

`func (o *PermissionPayloadDTO) SetAddAddonRole(v bool)`

SetAddAddonRole sets AddAddonRole field to given value.

### HasAddAddonRole

`func (o *PermissionPayloadDTO) HasAddAddonRole() bool`

HasAddAddonRole returns a boolean if a field has been set.

### GetDescription

`func (o *PermissionPayloadDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionPayloadDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionPayloadDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionPayloadDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrants

`func (o *PermissionPayloadDTO) GetGrants() []PermissionGrantDTO`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *PermissionPayloadDTO) GetGrantsOk() (*[]PermissionGrantDTO, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *PermissionPayloadDTO) SetGrants(v []PermissionGrantDTO)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *PermissionPayloadDTO) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetName

`func (o *PermissionPayloadDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionPayloadDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionPayloadDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionPayloadDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *PermissionPayloadDTO) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *PermissionPayloadDTO) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *PermissionPayloadDTO) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *PermissionPayloadDTO) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *PermissionPayloadDTO) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *PermissionPayloadDTO) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *PermissionPayloadDTO) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *PermissionPayloadDTO) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


