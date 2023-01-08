# FunctionOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | **[]string** | The list of function arguments. | 
**EncodedOperand** | Pointer to **string** | Encoded operand, which can be used directly in a JQL query. | [optional] 
**Function** | **string** | The name of the function. | 

## Methods

### NewFunctionOperand

`func NewFunctionOperand(arguments []string, function string, ) *FunctionOperand`

NewFunctionOperand instantiates a new FunctionOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionOperandWithDefaults

`func NewFunctionOperandWithDefaults() *FunctionOperand`

NewFunctionOperandWithDefaults instantiates a new FunctionOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *FunctionOperand) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *FunctionOperand) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *FunctionOperand) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetEncodedOperand

`func (o *FunctionOperand) GetEncodedOperand() string`

GetEncodedOperand returns the EncodedOperand field if non-nil, zero value otherwise.

### GetEncodedOperandOk

`func (o *FunctionOperand) GetEncodedOperandOk() (*string, bool)`

GetEncodedOperandOk returns a tuple with the EncodedOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedOperand

`func (o *FunctionOperand) SetEncodedOperand(v string)`

SetEncodedOperand sets EncodedOperand field to given value.

### HasEncodedOperand

`func (o *FunctionOperand) HasEncodedOperand() bool`

HasEncodedOperand returns a boolean if a field has been set.

### GetFunction

`func (o *FunctionOperand) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *FunctionOperand) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *FunctionOperand) SetFunction(v string)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


