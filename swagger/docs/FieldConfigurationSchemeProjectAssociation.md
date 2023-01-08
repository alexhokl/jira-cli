# FieldConfigurationSchemeProjectAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldConfigurationSchemeId** | Pointer to **string** | The ID of the field configuration scheme. If the field configuration scheme ID is &#x60;null&#x60;, the operation assigns the default field configuration scheme. | [optional] 
**ProjectId** | **string** | The ID of the project. | 

## Methods

### NewFieldConfigurationSchemeProjectAssociation

`func NewFieldConfigurationSchemeProjectAssociation(projectId string, ) *FieldConfigurationSchemeProjectAssociation`

NewFieldConfigurationSchemeProjectAssociation instantiates a new FieldConfigurationSchemeProjectAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigurationSchemeProjectAssociationWithDefaults

`func NewFieldConfigurationSchemeProjectAssociationWithDefaults() *FieldConfigurationSchemeProjectAssociation`

NewFieldConfigurationSchemeProjectAssociationWithDefaults instantiates a new FieldConfigurationSchemeProjectAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldConfigurationSchemeId

`func (o *FieldConfigurationSchemeProjectAssociation) GetFieldConfigurationSchemeId() string`

GetFieldConfigurationSchemeId returns the FieldConfigurationSchemeId field if non-nil, zero value otherwise.

### GetFieldConfigurationSchemeIdOk

`func (o *FieldConfigurationSchemeProjectAssociation) GetFieldConfigurationSchemeIdOk() (*string, bool)`

GetFieldConfigurationSchemeIdOk returns a tuple with the FieldConfigurationSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfigurationSchemeId

`func (o *FieldConfigurationSchemeProjectAssociation) SetFieldConfigurationSchemeId(v string)`

SetFieldConfigurationSchemeId sets FieldConfigurationSchemeId field to given value.

### HasFieldConfigurationSchemeId

`func (o *FieldConfigurationSchemeProjectAssociation) HasFieldConfigurationSchemeId() bool`

HasFieldConfigurationSchemeId returns a boolean if a field has been set.

### GetProjectId

`func (o *FieldConfigurationSchemeProjectAssociation) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FieldConfigurationSchemeProjectAssociation) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FieldConfigurationSchemeProjectAssociation) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


