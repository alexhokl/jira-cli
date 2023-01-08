# FieldConfigurationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the field within the field configuration. | [optional] 
**Id** | **string** | The ID of the field within the field configuration. | 
**IsHidden** | Pointer to **bool** | Whether the field is hidden in the field configuration. | [optional] 
**IsRequired** | Pointer to **bool** | Whether the field is required in the field configuration. | [optional] 
**Renderer** | Pointer to **string** | The renderer type for the field within the field configuration. | [optional] 

## Methods

### NewFieldConfigurationItem

`func NewFieldConfigurationItem(id string, ) *FieldConfigurationItem`

NewFieldConfigurationItem instantiates a new FieldConfigurationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigurationItemWithDefaults

`func NewFieldConfigurationItemWithDefaults() *FieldConfigurationItem`

NewFieldConfigurationItemWithDefaults instantiates a new FieldConfigurationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FieldConfigurationItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldConfigurationItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldConfigurationItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldConfigurationItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FieldConfigurationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldConfigurationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldConfigurationItem) SetId(v string)`

SetId sets Id field to given value.


### GetIsHidden

`func (o *FieldConfigurationItem) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *FieldConfigurationItem) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *FieldConfigurationItem) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *FieldConfigurationItem) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsRequired

`func (o *FieldConfigurationItem) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FieldConfigurationItem) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FieldConfigurationItem) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FieldConfigurationItem) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetRenderer

`func (o *FieldConfigurationItem) GetRenderer() string`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *FieldConfigurationItem) GetRendererOk() (*string, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *FieldConfigurationItem) SetRenderer(v string)`

SetRenderer sets Renderer field to given value.

### HasRenderer

`func (o *FieldConfigurationItem) HasRenderer() bool`

HasRenderer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


