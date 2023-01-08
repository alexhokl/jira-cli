# ListOperand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedOperand** | Pointer to **string** | Encoded operand, which can be used directly in a JQL query. | [optional] 
**Values** | [**[]JqlQueryUnitaryOperand**](JqlQueryUnitaryOperand.md) | The list of operand values. | 

## Methods

### NewListOperand

`func NewListOperand(values []JqlQueryUnitaryOperand, ) *ListOperand`

NewListOperand instantiates a new ListOperand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperandWithDefaults

`func NewListOperandWithDefaults() *ListOperand`

NewListOperandWithDefaults instantiates a new ListOperand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodedOperand

`func (o *ListOperand) GetEncodedOperand() string`

GetEncodedOperand returns the EncodedOperand field if non-nil, zero value otherwise.

### GetEncodedOperandOk

`func (o *ListOperand) GetEncodedOperandOk() (*string, bool)`

GetEncodedOperandOk returns a tuple with the EncodedOperand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedOperand

`func (o *ListOperand) SetEncodedOperand(v string)`

SetEncodedOperand sets EncodedOperand field to given value.

### HasEncodedOperand

`func (o *ListOperand) HasEncodedOperand() bool`

HasEncodedOperand returns a boolean if a field has been set.

### GetValues

`func (o *ListOperand) GetValues() []JqlQueryUnitaryOperand`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListOperand) GetValuesOk() (*[]JqlQueryUnitaryOperand, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListOperand) SetValues(v []JqlQueryUnitaryOperand)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


