# ValueOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedValue** | Pointer to **string** | Encoded value, which can be used directly in a JQL query. | [optional] 
**Value** | **string** | The operand value. | 

## Methods

### NewValueOperand

`func NewValueOperand(value string, ) *ValueOperand`

NewValueOperand instantiates a new ValueOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueOperandWithDefaults

`func NewValueOperandWithDefaults() *ValueOperand`

NewValueOperandWithDefaults instantiates a new ValueOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedValue

`func (o *ValueOperand) GetEncodedValue() string`

GetEncodedValue returns the EncodedValue field if non-nil, zero value otherwise.

### GetEncodedValueOk

`func (o *ValueOperand) GetEncodedValueOk() (*string, bool)`

GetEncodedValueOk returns a tuple with the EncodedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedValue

`func (o *ValueOperand) SetEncodedValue(v string)`

SetEncodedValue sets EncodedValue field to given value.

### HasEncodedValue

`func (o *ValueOperand) HasEncodedValue() bool`

HasEncodedValue returns a boolean if a field has been set.

### GetValue

`func (o *ValueOperand) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueOperand) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueOperand) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


