# JqlQueryField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedName** | Pointer to **string** | The encoded name of the field, which can be used directly in a JQL query. | [optional] 
**Name** | **string** | The name of the field. | 
**Property** | Pointer to [**[]JqlQueryFieldEntityProperty**](JqlQueryFieldEntityProperty.md) | When the field refers to a value in an entity property, details of the entity property value. | [optional] 

## Methods

### NewJqlQueryField

`func NewJqlQueryField(name string, ) *JqlQueryField`

NewJqlQueryField instantiates a new JqlQueryField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryFieldWithDefaults

`func NewJqlQueryFieldWithDefaults() *JqlQueryField`

NewJqlQueryFieldWithDefaults instantiates a new JqlQueryField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedName

`func (o *JqlQueryField) GetEncodedName() string`

GetEncodedName returns the EncodedName field if non-nil, zero value otherwise.

### GetEncodedNameOk

`func (o *JqlQueryField) GetEncodedNameOk() (*string, bool)`

GetEncodedNameOk returns a tuple with the EncodedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedName

`func (o *JqlQueryField) SetEncodedName(v string)`

SetEncodedName sets EncodedName field to given value.

### HasEncodedName

`func (o *JqlQueryField) HasEncodedName() bool`

HasEncodedName returns a boolean if a field has been set.

### GetName

`func (o *JqlQueryField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JqlQueryField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JqlQueryField) SetName(v string)`

SetName sets Name field to given value.


### GetProperty

`func (o *JqlQueryField) GetProperty() []JqlQueryFieldEntityProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *JqlQueryField) GetPropertyOk() (*[]JqlQueryFieldEntityProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *JqlQueryField) SetProperty(v []JqlQueryFieldEntityProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *JqlQueryField) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


