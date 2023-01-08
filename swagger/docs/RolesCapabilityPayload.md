# RolesCapabilityPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleToProjectActors** | Pointer to [**map[string][]ProjectCreateResourceIdentifier**](array.md) | A map of role PCRI (can be ID or REF) to a list of user or group PCRI IDs to associate with the role and project. | [optional] 
**Roles** | Pointer to [**[]RolePayload**](RolePayload.md) | The list of roles to create. | [optional] 

## Methods

### NewRolesCapabilityPayload

`func NewRolesCapabilityPayload() *RolesCapabilityPayload`

NewRolesCapabilityPayload instantiates a new RolesCapabilityPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesCapabilityPayloadWithDefaults

`func NewRolesCapabilityPayloadWithDefaults() *RolesCapabilityPayload`

NewRolesCapabilityPayloadWithDefaults instantiates a new RolesCapabilityPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleToProjectActors

`func (o *RolesCapabilityPayload) GetRoleToProjectActors() map[string][]ProjectCreateResourceIdentifier`

GetRoleToProjectActors returns the RoleToProjectActors field if non-nil, zero value otherwise.

### GetRoleToProjectActorsOk

`func (o *RolesCapabilityPayload) GetRoleToProjectActorsOk() (*map[string][]ProjectCreateResourceIdentifier, bool)`

GetRoleToProjectActorsOk returns a tuple with the RoleToProjectActors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleToProjectActors

`func (o *RolesCapabilityPayload) SetRoleToProjectActors(v map[string][]ProjectCreateResourceIdentifier)`

SetRoleToProjectActors sets RoleToProjectActors field to given value.

### HasRoleToProjectActors

`func (o *RolesCapabilityPayload) HasRoleToProjectActors() bool`

HasRoleToProjectActors returns a boolean if a field has been set.

### GetRoles

`func (o *RolesCapabilityPayload) GetRoles() []RolePayload`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RolesCapabilityPayload) GetRolesOk() (*[]RolePayload, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RolesCapabilityPayload) SetRoles(v []RolePayload)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RolesCapabilityPayload) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


