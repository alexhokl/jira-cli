# FieldIdentifierObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewFieldIdentifierObject

`func NewFieldIdentifierObject(type_ string, ) *FieldIdentifierObject`

NewFieldIdentifierObject instantiates a new FieldIdentifierObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldIdentifierObjectWithDefaults

`func NewFieldIdentifierObjectWithDefaults() *FieldIdentifierObject`

NewFieldIdentifierObjectWithDefaults instantiates a new FieldIdentifierObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *FieldIdentifierObject) GetIdentifier() map[string]interface{}`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *FieldIdentifierObject) GetIdentifierOk() (*map[string]interface{}, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *FieldIdentifierObject) SetIdentifier(v map[string]interface{})`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *FieldIdentifierObject) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetType

`func (o *FieldIdentifierObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldIdentifierObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldIdentifierObject) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


