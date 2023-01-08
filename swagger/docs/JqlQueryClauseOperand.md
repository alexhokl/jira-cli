# JqlQueryClauseOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedOperand** | Pointer to **string** | Encoded operand, which can be used directly in a JQL query. | [optional] 
**Values** | [**[]JqlQueryUnitaryOperand**](JqlQueryUnitaryOperand.md) | The list of operand values. | 
**EncodedValue** | Pointer to **string** | Encoded value, which can be used directly in a JQL query. | [optional] 
**Value** | **string** | The operand value. | 
**Arguments** | **[]string** | The list of function arguments. | 
**Function** | **string** | The name of the function. | 
**Keyword** | **string** | The keyword that is the operand value. | 

## Methods

### NewJqlQueryClauseOperand

`func NewJqlQueryClauseOperand(values []JqlQueryUnitaryOperand, value string, arguments []string, function string, keyword string, ) *JqlQueryClauseOperand`

NewJqlQueryClauseOperand instantiates a new JqlQueryClauseOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryClauseOperandWithDefaults

`func NewJqlQueryClauseOperandWithDefaults() *JqlQueryClauseOperand`

NewJqlQueryClauseOperandWithDefaults instantiates a new JqlQueryClauseOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedOperand

`func (o *JqlQueryClauseOperand) GetEncodedOperand() string`

GetEncodedOperand returns the EncodedOperand field if non-nil, zero value otherwise.

### GetEncodedOperandOk

`func (o *JqlQueryClauseOperand) GetEncodedOperandOk() (*string, bool)`

GetEncodedOperandOk returns a tuple with the EncodedOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedOperand

`func (o *JqlQueryClauseOperand) SetEncodedOperand(v string)`

SetEncodedOperand sets EncodedOperand field to given value.

### HasEncodedOperand

`func (o *JqlQueryClauseOperand) HasEncodedOperand() bool`

HasEncodedOperand returns a boolean if a field has been set.

### GetValues

`func (o *JqlQueryClauseOperand) GetValues() []JqlQueryUnitaryOperand`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *JqlQueryClauseOperand) GetValuesOk() (*[]JqlQueryUnitaryOperand, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *JqlQueryClauseOperand) SetValues(v []JqlQueryUnitaryOperand)`

SetValues sets Values field to given value.


### GetEncodedValue

`func (o *JqlQueryClauseOperand) GetEncodedValue() string`

GetEncodedValue returns the EncodedValue field if non-nil, zero value otherwise.

### GetEncodedValueOk

`func (o *JqlQueryClauseOperand) GetEncodedValueOk() (*string, bool)`

GetEncodedValueOk returns a tuple with the EncodedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedValue

`func (o *JqlQueryClauseOperand) SetEncodedValue(v string)`

SetEncodedValue sets EncodedValue field to given value.

### HasEncodedValue

`func (o *JqlQueryClauseOperand) HasEncodedValue() bool`

HasEncodedValue returns a boolean if a field has been set.

### GetValue

`func (o *JqlQueryClauseOperand) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JqlQueryClauseOperand) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JqlQueryClauseOperand) SetValue(v string)`

SetValue sets Value field to given value.


### GetArguments

`func (o *JqlQueryClauseOperand) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JqlQueryClauseOperand) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JqlQueryClauseOperand) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetFunction

`func (o *JqlQueryClauseOperand) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *JqlQueryClauseOperand) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *JqlQueryClauseOperand) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetKeyword

`func (o *JqlQueryClauseOperand) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *JqlQueryClauseOperand) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *JqlQueryClauseOperand) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


