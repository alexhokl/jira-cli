# BulkContextualConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **interface{}** | The field configuration. | [optional] 
**CustomFieldId** | **string** | The ID of the custom field. | 
**FieldContextId** | **string** | The ID of the field context the configuration is associated with. | [readonly] 
**Id** | **string** | The ID of the configuration. | 
**Schema** | Pointer to **interface{}** | The field value schema. | [optional] 

## Methods

### NewBulkContextualConfiguration

`func NewBulkContextualConfiguration(customFieldId string, fieldContextId string, id string, ) *BulkContextualConfiguration`

NewBulkContextualConfiguration instantiates a new BulkContextualConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkContextualConfigurationWithDefaults

`func NewBulkContextualConfigurationWithDefaults() *BulkContextualConfiguration`

NewBulkContextualConfigurationWithDefaults instantiates a new BulkContextualConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *BulkContextualConfiguration) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BulkContextualConfiguration) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BulkContextualConfiguration) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BulkContextualConfiguration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *BulkContextualConfiguration) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *BulkContextualConfiguration) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetCustomFieldId

`func (o *BulkContextualConfiguration) GetCustomFieldId() string`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *BulkContextualConfiguration) GetCustomFieldIdOk() (*string, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *BulkContextualConfiguration) SetCustomFieldId(v string)`

SetCustomFieldId sets CustomFieldId field to given value.


### GetFieldContextId

`func (o *BulkContextualConfiguration) GetFieldContextId() string`

GetFieldContextId returns the FieldContextId field if non-nil, zero value otherwise.

### GetFieldContextIdOk

`func (o *BulkContextualConfiguration) GetFieldContextIdOk() (*string, bool)`

GetFieldContextIdOk returns a tuple with the FieldContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContextId

`func (o *BulkContextualConfiguration) SetFieldContextId(v string)`

SetFieldContextId sets FieldContextId field to given value.


### GetId

`func (o *BulkContextualConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkContextualConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkContextualConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetSchema

`func (o *BulkContextualConfiguration) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BulkContextualConfiguration) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BulkContextualConfiguration) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *BulkContextualConfiguration) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *BulkContextualConfiguration) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *BulkContextualConfiguration) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


