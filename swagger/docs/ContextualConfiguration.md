# ContextualConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **interface{}** | The field configuration. | [optional] 
**FieldContextId** | **string** | The ID of the field context the configuration is associated with. | [readonly] 
**Id** | **string** | The ID of the configuration. | 
**Schema** | Pointer to **interface{}** | The field value schema. | [optional] 

## Methods

### NewContextualConfiguration

`func NewContextualConfiguration(fieldContextId string, id string, ) *ContextualConfiguration`

NewContextualConfiguration instantiates a new ContextualConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextualConfigurationWithDefaults

`func NewContextualConfigurationWithDefaults() *ContextualConfiguration`

NewContextualConfigurationWithDefaults instantiates a new ContextualConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ContextualConfiguration) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ContextualConfiguration) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ContextualConfiguration) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ContextualConfiguration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *ContextualConfiguration) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *ContextualConfiguration) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetFieldContextId

`func (o *ContextualConfiguration) GetFieldContextId() string`

GetFieldContextId returns the FieldContextId field if non-nil, zero value otherwise.

### GetFieldContextIdOk

`func (o *ContextualConfiguration) GetFieldContextIdOk() (*string, bool)`

GetFieldContextIdOk returns a tuple with the FieldContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContextId

`func (o *ContextualConfiguration) SetFieldContextId(v string)`

SetFieldContextId sets FieldContextId field to given value.


### GetId

`func (o *ContextualConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContextualConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContextualConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetSchema

`func (o *ContextualConfiguration) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ContextualConfiguration) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ContextualConfiguration) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ContextualConfiguration) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ContextualConfiguration) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ContextualConfiguration) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


