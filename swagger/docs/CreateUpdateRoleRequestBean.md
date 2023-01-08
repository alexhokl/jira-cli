# CreateUpdateRoleRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the project role. Required when fully updating a project role. Optional when creating or partially updating a project role. | [optional] 
**Name** | Pointer to **string** | The name of the project role. Must be unique. Cannot begin or end with whitespace. The maximum length is 255 characters. Required when creating a project role. Optional when partially updating a project role. | [optional] 

## Methods

### NewCreateUpdateRoleRequestBean

`func NewCreateUpdateRoleRequestBean() *CreateUpdateRoleRequestBean`

NewCreateUpdateRoleRequestBean instantiates a new CreateUpdateRoleRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateRoleRequestBeanWithDefaults

`func NewCreateUpdateRoleRequestBeanWithDefaults() *CreateUpdateRoleRequestBean`

NewCreateUpdateRoleRequestBeanWithDefaults instantiates a new CreateUpdateRoleRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateUpdateRoleRequestBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdateRoleRequestBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdateRoleRequestBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdateRoleRequestBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateUpdateRoleRequestBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdateRoleRequestBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdateRoleRequestBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUpdateRoleRequestBean) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


