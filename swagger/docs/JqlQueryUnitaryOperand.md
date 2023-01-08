# JqlQueryUnitaryOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedValue** | Pointer to **string** | Encoded value, which can be used directly in a JQL query. | [optional] 
**Value** | **string** | The operand value. | 
**Arguments** | **[]string** | The list of function arguments. | 
**EncodedOperand** | Pointer to **string** | Encoded operand, which can be used directly in a JQL query. | [optional] 
**Function** | **string** | The name of the function. | 
**Keyword** | **string** | The keyword that is the operand value. | 

## Methods

### NewJqlQueryUnitaryOperand

`func NewJqlQueryUnitaryOperand(value string, arguments []string, function string, keyword string, ) *JqlQueryUnitaryOperand`

NewJqlQueryUnitaryOperand instantiates a new JqlQueryUnitaryOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryUnitaryOperandWithDefaults

`func NewJqlQueryUnitaryOperandWithDefaults() *JqlQueryUnitaryOperand`

NewJqlQueryUnitaryOperandWithDefaults instantiates a new JqlQueryUnitaryOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedValue

`func (o *JqlQueryUnitaryOperand) GetEncodedValue() string`

GetEncodedValue returns the EncodedValue field if non-nil, zero value otherwise.

### GetEncodedValueOk

`func (o *JqlQueryUnitaryOperand) GetEncodedValueOk() (*string, bool)`

GetEncodedValueOk returns a tuple with the EncodedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedValue

`func (o *JqlQueryUnitaryOperand) SetEncodedValue(v string)`

SetEncodedValue sets EncodedValue field to given value.

### HasEncodedValue

`func (o *JqlQueryUnitaryOperand) HasEncodedValue() bool`

HasEncodedValue returns a boolean if a field has been set.

### GetValue

`func (o *JqlQueryUnitaryOperand) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JqlQueryUnitaryOperand) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JqlQueryUnitaryOperand) SetValue(v string)`

SetValue sets Value field to given value.


### GetArguments

`func (o *JqlQueryUnitaryOperand) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JqlQueryUnitaryOperand) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JqlQueryUnitaryOperand) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetEncodedOperand

`func (o *JqlQueryUnitaryOperand) GetEncodedOperand() string`

GetEncodedOperand returns the EncodedOperand field if non-nil, zero value otherwise.

### GetEncodedOperandOk

`func (o *JqlQueryUnitaryOperand) GetEncodedOperandOk() (*string, bool)`

GetEncodedOperandOk returns a tuple with the EncodedOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedOperand

`func (o *JqlQueryUnitaryOperand) SetEncodedOperand(v string)`

SetEncodedOperand sets EncodedOperand field to given value.

### HasEncodedOperand

`func (o *JqlQueryUnitaryOperand) HasEncodedOperand() bool`

HasEncodedOperand returns a boolean if a field has been set.

### GetFunction

`func (o *JqlQueryUnitaryOperand) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *JqlQueryUnitaryOperand) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *JqlQueryUnitaryOperand) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetKeyword

`func (o *JqlQueryUnitaryOperand) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *JqlQueryUnitaryOperand) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *JqlQueryUnitaryOperand) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


