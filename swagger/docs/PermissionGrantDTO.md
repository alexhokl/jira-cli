# PermissionGrantDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationAccess** | Pointer to **[]string** |  | [optional] 
**GroupCustomFields** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Groups** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**PermissionKeys** | Pointer to **[]string** |  | [optional] 
**ProjectRoles** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**SpecialGrants** | Pointer to **[]string** |  | [optional] 
**UserCustomFields** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Users** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewPermissionGrantDTO

`func NewPermissionGrantDTO() *PermissionGrantDTO`

NewPermissionGrantDTO instantiates a new PermissionGrantDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantDTOWithDefaults

`func NewPermissionGrantDTOWithDefaults() *PermissionGrantDTO`

NewPermissionGrantDTOWithDefaults instantiates a new PermissionGrantDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationAccess

`func (o *PermissionGrantDTO) GetApplicationAccess() []string`

GetApplicationAccess returns the ApplicationAccess field if non-nil, zero value otherwise.

### GetApplicationAccessOk

`func (o *PermissionGrantDTO) GetApplicationAccessOk() (*[]string, bool)`

GetApplicationAccessOk returns a tuple with the ApplicationAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAccess

`func (o *PermissionGrantDTO) SetApplicationAccess(v []string)`

SetApplicationAccess sets ApplicationAccess field to given value.

### HasApplicationAccess

`func (o *PermissionGrantDTO) HasApplicationAccess() bool`

HasApplicationAccess returns a boolean if a field has been set.

### GetGroupCustomFields

`func (o *PermissionGrantDTO) GetGroupCustomFields() []ProjectCreateResourceIdentifier`

GetGroupCustomFields returns the GroupCustomFields field if non-nil, zero value otherwise.

### GetGroupCustomFieldsOk

`func (o *PermissionGrantDTO) GetGroupCustomFieldsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetGroupCustomFieldsOk returns a tuple with the GroupCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCustomFields

`func (o *PermissionGrantDTO) SetGroupCustomFields(v []ProjectCreateResourceIdentifier)`

SetGroupCustomFields sets GroupCustomFields field to given value.

### HasGroupCustomFields

`func (o *PermissionGrantDTO) HasGroupCustomFields() bool`

HasGroupCustomFields returns a boolean if a field has been set.

### GetGroups

`func (o *PermissionGrantDTO) GetGroups() []ProjectCreateResourceIdentifier`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PermissionGrantDTO) GetGroupsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PermissionGrantDTO) SetGroups(v []ProjectCreateResourceIdentifier)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PermissionGrantDTO) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPermissionKeys

`func (o *PermissionGrantDTO) GetPermissionKeys() []string`

GetPermissionKeys returns the PermissionKeys field if non-nil, zero value otherwise.

### GetPermissionKeysOk

`func (o *PermissionGrantDTO) GetPermissionKeysOk() (*[]string, bool)`

GetPermissionKeysOk returns a tuple with the PermissionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionKeys

`func (o *PermissionGrantDTO) SetPermissionKeys(v []string)`

SetPermissionKeys sets PermissionKeys field to given value.

### HasPermissionKeys

`func (o *PermissionGrantDTO) HasPermissionKeys() bool`

HasPermissionKeys returns a boolean if a field has been set.

### GetProjectRoles

`func (o *PermissionGrantDTO) GetProjectRoles() []ProjectCreateResourceIdentifier`

GetProjectRoles returns the ProjectRoles field if non-nil, zero value otherwise.

### GetProjectRolesOk

`func (o *PermissionGrantDTO) GetProjectRolesOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetProjectRolesOk returns a tuple with the ProjectRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoles

`func (o *PermissionGrantDTO) SetProjectRoles(v []ProjectCreateResourceIdentifier)`

SetProjectRoles sets ProjectRoles field to given value.

### HasProjectRoles

`func (o *PermissionGrantDTO) HasProjectRoles() bool`

HasProjectRoles returns a boolean if a field has been set.

### GetSpecialGrants

`func (o *PermissionGrantDTO) GetSpecialGrants() []string`

GetSpecialGrants returns the SpecialGrants field if non-nil, zero value otherwise.

### GetSpecialGrantsOk

`func (o *PermissionGrantDTO) GetSpecialGrantsOk() (*[]string, bool)`

GetSpecialGrantsOk returns a tuple with the SpecialGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialGrants

`func (o *PermissionGrantDTO) SetSpecialGrants(v []string)`

SetSpecialGrants sets SpecialGrants field to given value.

### HasSpecialGrants

`func (o *PermissionGrantDTO) HasSpecialGrants() bool`

HasSpecialGrants returns a boolean if a field has been set.

### GetUserCustomFields

`func (o *PermissionGrantDTO) GetUserCustomFields() []ProjectCreateResourceIdentifier`

GetUserCustomFields returns the UserCustomFields field if non-nil, zero value otherwise.

### GetUserCustomFieldsOk

`func (o *PermissionGrantDTO) GetUserCustomFieldsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetUserCustomFieldsOk returns a tuple with the UserCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCustomFields

`func (o *PermissionGrantDTO) SetUserCustomFields(v []ProjectCreateResourceIdentifier)`

SetUserCustomFields sets UserCustomFields field to given value.

### HasUserCustomFields

`func (o *PermissionGrantDTO) HasUserCustomFields() bool`

HasUserCustomFields returns a boolean if a field has been set.

### GetUsers

`func (o *PermissionGrantDTO) GetUsers() []ProjectCreateResourceIdentifier`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PermissionGrantDTO) GetUsersOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PermissionGrantDTO) SetUsers(v []ProjectCreateResourceIdentifier)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PermissionGrantDTO) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


