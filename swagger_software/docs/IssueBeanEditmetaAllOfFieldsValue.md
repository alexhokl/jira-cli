# IssueBeanEditmetaAllOfFieldsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedValues** | Pointer to **[]interface{}** | The list of values allowed in the field. | [optional] [readonly] 
**AutoCompleteUrl** | Pointer to **string** | The URL that can be used to automatically complete the field. | [optional] [readonly] 
**Configuration** | Pointer to **map[string]interface{}** | The configuration properties. | [optional] [readonly] 
**DefaultValue** | Pointer to **interface{}** | The default value of the field. | [optional] [readonly] 
**HasDefaultValue** | Pointer to **bool** | Whether the field has a default value. | [optional] [readonly] 
**Key** | **string** | The key of the field. | [readonly] 
**Name** | **string** | The name of the field. | [readonly] 
**Operations** | **[]string** | The list of operations that can be performed on the field. | [readonly] 
**Required** | **bool** | Whether the field is required. | [readonly] 
**Schema** | [**FieldMetadataSchema**](FieldMetadataSchema.md) |  | 

## Methods

### NewIssueBeanEditmetaAllOfFieldsValue

`func NewIssueBeanEditmetaAllOfFieldsValue(key string, name string, operations []string, required bool, schema FieldMetadataSchema, ) *IssueBeanEditmetaAllOfFieldsValue`

NewIssueBeanEditmetaAllOfFieldsValue instantiates a new IssueBeanEditmetaAllOfFieldsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBeanEditmetaAllOfFieldsValueWithDefaults

`func NewIssueBeanEditmetaAllOfFieldsValueWithDefaults() *IssueBeanEditmetaAllOfFieldsValue`

NewIssueBeanEditmetaAllOfFieldsValueWithDefaults instantiates a new IssueBeanEditmetaAllOfFieldsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedValues

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetAllowedValues() []interface{}`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetAllowedValuesOk() (*[]interface{}, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetAllowedValues(v []interface{})`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *IssueBeanEditmetaAllOfFieldsValue) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetAutoCompleteUrl

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetAutoCompleteUrl() string`

GetAutoCompleteUrl returns the AutoCompleteUrl field if non-nil, zero value otherwise.

### GetAutoCompleteUrlOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetAutoCompleteUrlOk() (*string, bool)`

GetAutoCompleteUrlOk returns a tuple with the AutoCompleteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteUrl

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetAutoCompleteUrl(v string)`

SetAutoCompleteUrl sets AutoCompleteUrl field to given value.

### HasAutoCompleteUrl

`func (o *IssueBeanEditmetaAllOfFieldsValue) HasAutoCompleteUrl() bool`

HasAutoCompleteUrl returns a boolean if a field has been set.

### GetConfiguration

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IssueBeanEditmetaAllOfFieldsValue) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *IssueBeanEditmetaAllOfFieldsValue) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetHasDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetHasDefaultValue() bool`

GetHasDefaultValue returns the HasDefaultValue field if non-nil, zero value otherwise.

### GetHasDefaultValueOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetHasDefaultValueOk() (*bool, bool)`

GetHasDefaultValueOk returns a tuple with the HasDefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetHasDefaultValue(v bool)`

SetHasDefaultValue sets HasDefaultValue field to given value.

### HasHasDefaultValue

`func (o *IssueBeanEditmetaAllOfFieldsValue) HasHasDefaultValue() bool`

HasHasDefaultValue returns a boolean if a field has been set.

### GetKey

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### GetRequired

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetSchema

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetSchema() FieldMetadataSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *IssueBeanEditmetaAllOfFieldsValue) GetSchemaOk() (*FieldMetadataSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *IssueBeanEditmetaAllOfFieldsValue) SetSchema(v FieldMetadataSchema)`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


