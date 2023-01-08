# RolePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultActors** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | The default actors for the role. By adding default actors, the role will be added to any future projects created | [optional] 
**Description** | Pointer to **string** | The description of the role | [optional] 
**Name** | Pointer to **string** | The name of the role | [optional] 
**OnConflict** | Pointer to **string** | The strategy to use when there is a conflict with an existing project role. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters | [optional] [default to "USE"]
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the role. Only used by project-scoped project | [optional] 

## Methods

### NewRolePayload

`func NewRolePayload() *RolePayload`

NewRolePayload instantiates a new RolePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePayloadWithDefaults

`func NewRolePayloadWithDefaults() *RolePayload`

NewRolePayloadWithDefaults instantiates a new RolePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultActors

`func (o *RolePayload) GetDefaultActors() []ProjectCreateResourceIdentifier`

GetDefaultActors returns the DefaultActors field if non-nil, zero value otherwise.

### GetDefaultActorsOk

`func (o *RolePayload) GetDefaultActorsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetDefaultActorsOk returns a tuple with the DefaultActors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActors

`func (o *RolePayload) SetDefaultActors(v []ProjectCreateResourceIdentifier)`

SetDefaultActors sets DefaultActors field to given value.

### HasDefaultActors

`func (o *RolePayload) HasDefaultActors() bool`

HasDefaultActors returns a boolean if a field has been set.

### GetDescription

`func (o *RolePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *RolePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *RolePayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *RolePayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *RolePayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *RolePayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *RolePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *RolePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *RolePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *RolePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetType

`func (o *RolePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RolePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RolePayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RolePayload) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


