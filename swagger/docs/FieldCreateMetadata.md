# FieldCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedValues** | Pointer to **[]interface{}** | The list of values allowed in the field. | [optional] [readonly] 
**AutoCompleteUrl** | Pointer to **string** | The URL that can be used to automatically complete the field. | [optional] [readonly] 
**Configuration** | Pointer to **map[string]interface{}** | The configuration properties. | [optional] [readonly] 
**DefaultValue** | Pointer to **interface{}** | The default value of the field. | [optional] [readonly] 
**FieldId** | **string** | The field id. | [readonly] 
**HasDefaultValue** | Pointer to **bool** | Whether the field has a default value. | [optional] [readonly] 
**Key** | **string** | The key of the field. | [readonly] 
**Name** | **string** | The name of the field. | [readonly] 
**Operations** | **[]string** | The list of operations that can be performed on the field. | [readonly] 
**Required** | **bool** | Whether the field is required. | [readonly] 
**Schema** | [**JsonTypeBean**](JsonTypeBean.md) | The data type of the field. | [readonly] 

## Methods

### NewFieldCreateMetadata

`func NewFieldCreateMetadata(fieldId string, key string, name string, operations []string, required bool, schema JsonTypeBean, ) *FieldCreateMetadata`

NewFieldCreateMetadata instantiates a new FieldCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldCreateMetadataWithDefaults

`func NewFieldCreateMetadataWithDefaults() *FieldCreateMetadata`

NewFieldCreateMetadataWithDefaults instantiates a new FieldCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *FieldCreateMetadata) GetAllowedValues() []interface{}`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *FieldCreateMetadata) GetAllowedValuesOk() (*[]interface{}, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *FieldCreateMetadata) SetAllowedValues(v []interface{})`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *FieldCreateMetadata) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetAutoCompleteUrl

`func (o *FieldCreateMetadata) GetAutoCompleteUrl() string`

GetAutoCompleteUrl returns the AutoCompleteUrl field if non-nil, zero value otherwise.

### GetAutoCompleteUrlOk

`func (o *FieldCreateMetadata) GetAutoCompleteUrlOk() (*string, bool)`

GetAutoCompleteUrlOk returns a tuple with the AutoCompleteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteUrl

`func (o *FieldCreateMetadata) SetAutoCompleteUrl(v string)`

SetAutoCompleteUrl sets AutoCompleteUrl field to given value.

### HasAutoCompleteUrl

`func (o *FieldCreateMetadata) HasAutoCompleteUrl() bool`

HasAutoCompleteUrl returns a boolean if a field has been set.

### GetConfiguration

`func (o *FieldCreateMetadata) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FieldCreateMetadata) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FieldCreateMetadata) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FieldCreateMetadata) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDefaultValue

`func (o *FieldCreateMetadata) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FieldCreateMetadata) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FieldCreateMetadata) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FieldCreateMetadata) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *FieldCreateMetadata) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *FieldCreateMetadata) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetFieldId

`func (o *FieldCreateMetadata) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *FieldCreateMetadata) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *FieldCreateMetadata) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetHasDefaultValue

`func (o *FieldCreateMetadata) GetHasDefaultValue() bool`

GetHasDefaultValue returns the HasDefaultValue field if non-nil, zero value otherwise.

### GetHasDefaultValueOk

`func (o *FieldCreateMetadata) GetHasDefaultValueOk() (*bool, bool)`

GetHasDefaultValueOk returns a tuple with the HasDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultValue

`func (o *FieldCreateMetadata) SetHasDefaultValue(v bool)`

SetHasDefaultValue sets HasDefaultValue field to given value.

### HasHasDefaultValue

`func (o *FieldCreateMetadata) HasHasDefaultValue() bool`

HasHasDefaultValue returns a boolean if a field has been set.

### GetKey

`func (o *FieldCreateMetadata) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FieldCreateMetadata) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FieldCreateMetadata) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FieldCreateMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldCreateMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldCreateMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *FieldCreateMetadata) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *FieldCreateMetadata) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *FieldCreateMetadata) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### GetRequired

`func (o *FieldCreateMetadata) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldCreateMetadata) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldCreateMetadata) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSchema

`func (o *FieldCreateMetadata) GetSchema() JsonTypeBean`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *FieldCreateMetadata) GetSchemaOk() (*JsonTypeBean, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *FieldCreateMetadata) SetSchema(v JsonTypeBean)`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


