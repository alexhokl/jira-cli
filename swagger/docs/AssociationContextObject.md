# AssociationContextObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAssociationContextObject

`func NewAssociationContextObject(type_ string, ) *AssociationContextObject`

NewAssociationContextObject instantiates a new AssociationContextObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationContextObjectWithDefaults

`func NewAssociationContextObjectWithDefaults() *AssociationContextObject`

NewAssociationContextObjectWithDefaults instantiates a new AssociationContextObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *AssociationContextObject) GetIdentifier() map[string]interface{}`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AssociationContextObject) GetIdentifierOk() (*map[string]interface{}, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AssociationContextObject) SetIdentifier(v map[string]interface{})`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AssociationContextObject) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetType

`func (o *AssociationContextObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssociationContextObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssociationContextObject) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


