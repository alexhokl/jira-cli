# FieldMetadataSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | If the field is a custom field, the configuration of the field. | [optional] [readonly] 
**Custom** | Pointer to **string** | If the field is a custom field, the URI of the field. | [optional] [readonly] 
**CustomId** | Pointer to **int64** | If the field is a custom field, the custom ID of the field. | [optional] [readonly] 
**Items** | Pointer to **string** | When the data type is an array, the name of the field items within the array. | [optional] [readonly] 
**System** | Pointer to **string** | If the field is a system field, the name of the field. | [optional] [readonly] 
**Type** | **string** | The data type of the field. | [readonly] 

## Methods

### NewFieldMetadataSchema

`func NewFieldMetadataSchema(type_ string, ) *FieldMetadataSchema`

NewFieldMetadataSchema instantiates a new FieldMetadataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMetadataSchemaWithDefaults

`func NewFieldMetadataSchemaWithDefaults() *FieldMetadataSchema`

NewFieldMetadataSchemaWithDefaults instantiates a new FieldMetadataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *FieldMetadataSchema) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FieldMetadataSchema) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FieldMetadataSchema) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FieldMetadataSchema) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCustom

`func (o *FieldMetadataSchema) GetCustom() string`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *FieldMetadataSchema) GetCustomOk() (*string, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *FieldMetadataSchema) SetCustom(v string)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *FieldMetadataSchema) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetCustomId

`func (o *FieldMetadataSchema) GetCustomId() int64`

GetCustomId returns the CustomId field if non-nil, zero value otherwise.

### GetCustomIdOk

`func (o *FieldMetadataSchema) GetCustomIdOk() (*int64, bool)`

GetCustomIdOk returns a tuple with the CustomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomId

`func (o *FieldMetadataSchema) SetCustomId(v int64)`

SetCustomId sets CustomId field to given value.

### HasCustomId

`func (o *FieldMetadataSchema) HasCustomId() bool`

HasCustomId returns a boolean if a field has been set.

### GetItems

`func (o *FieldMetadataSchema) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FieldMetadataSchema) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FieldMetadataSchema) SetItems(v string)`

SetItems sets Items field to given value.

### HasItems

`func (o *FieldMetadataSchema) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSystem

`func (o *FieldMetadataSchema) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *FieldMetadataSchema) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *FieldMetadataSchema) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *FieldMetadataSchema) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *FieldMetadataSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldMetadataSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldMetadataSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


