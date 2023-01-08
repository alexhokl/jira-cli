# Fields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retain** | Pointer to **NullableBool** | If &#x60;true&#x60;, will try to retain original non-null issue field values on move. | [optional] [default to true]
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFields

`func NewFields() *Fields`

NewFields instantiates a new Fields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldsWithDefaults

`func NewFieldsWithDefaults() *Fields`

NewFieldsWithDefaults instantiates a new Fields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetain

`func (o *Fields) GetRetain() bool`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *Fields) GetRetainOk() (*bool, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *Fields) SetRetain(v bool)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *Fields) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### SetRetainNil

`func (o *Fields) SetRetainNil(b bool)`

 SetRetainNil sets the value for Retain to be an explicit nil

### UnsetRetain
`func (o *Fields) UnsetRetain()`

UnsetRetain ensures that no value is present for Retain, not even an explicit nil
### GetType

`func (o *Fields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Fields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Fields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Fields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Fields) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Fields) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Fields) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Fields) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


