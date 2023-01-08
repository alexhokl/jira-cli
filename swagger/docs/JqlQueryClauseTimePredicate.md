# JqlQueryClauseTimePredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operand** | [**JqlQueryClauseOperand**](JqlQueryClauseOperand.md) |  | 
**Operator** | **string** | The operator between the field and the operand. | 

## Methods

### NewJqlQueryClauseTimePredicate

`func NewJqlQueryClauseTimePredicate(operand JqlQueryClauseOperand, operator string, ) *JqlQueryClauseTimePredicate`

NewJqlQueryClauseTimePredicate instantiates a new JqlQueryClauseTimePredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryClauseTimePredicateWithDefaults

`func NewJqlQueryClauseTimePredicateWithDefaults() *JqlQueryClauseTimePredicate`

NewJqlQueryClauseTimePredicateWithDefaults instantiates a new JqlQueryClauseTimePredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperand

`func (o *JqlQueryClauseTimePredicate) GetOperand() JqlQueryClauseOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *JqlQueryClauseTimePredicate) GetOperandOk() (*JqlQueryClauseOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *JqlQueryClauseTimePredicate) SetOperand(v JqlQueryClauseOperand)`

SetOperand sets Operand field to given value.


### GetOperator

`func (o *JqlQueryClauseTimePredicate) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JqlQueryClauseTimePredicate) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JqlQueryClauseTimePredicate) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


