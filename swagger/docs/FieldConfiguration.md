# FieldConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the field configuration. | 
**Id** | **int64** | The ID of the field configuration. | 
**IsDefault** | Pointer to **bool** | Whether the field configuration is the default. | [optional] 
**Name** | **string** | The name of the field configuration. | 

## Methods

### NewFieldConfiguration

`func NewFieldConfiguration(description string, id int64, name string, ) *FieldConfiguration`

NewFieldConfiguration instantiates a new FieldConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigurationWithDefaults

`func NewFieldConfigurationWithDefaults() *FieldConfiguration`

NewFieldConfigurationWithDefaults instantiates a new FieldConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FieldConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *FieldConfiguration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldConfiguration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldConfiguration) SetId(v int64)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *FieldConfiguration) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *FieldConfiguration) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *FieldConfiguration) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *FieldConfiguration) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *FieldConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldConfiguration) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


