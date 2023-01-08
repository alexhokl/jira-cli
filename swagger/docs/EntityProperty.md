# EntityProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the property. Required on create and update. | [optional] 
**Value** | Pointer to **interface{}** | The value of the property. Required on create and update. | [optional] 

## Methods

### NewEntityProperty

`func NewEntityProperty() *EntityProperty`

NewEntityProperty instantiates a new EntityProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertyWithDefaults

`func NewEntityPropertyWithDefaults() *EntityProperty`

NewEntityPropertyWithDefaults instantiates a new EntityProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EntityProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EntityProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *EntityProperty) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityProperty) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityProperty) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *EntityProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EntityProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EntityProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


