# FieldWasClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | [**JqlQueryField**](JqlQueryField.md) |  | 
**Operand** | [**JqlQueryClauseOperand**](JqlQueryClauseOperand.md) |  | 
**Operator** | **string** | The operator between the field and operand. | 
**Predicates** | [**[]JqlQueryClauseTimePredicate**](JqlQueryClauseTimePredicate.md) | The list of time predicates. | 

## Methods

### NewFieldWasClause

`func NewFieldWasClause(field JqlQueryField, operand JqlQueryClauseOperand, operator string, predicates []JqlQueryClauseTimePredicate, ) *FieldWasClause`

NewFieldWasClause instantiates a new FieldWasClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWasClauseWithDefaults

`func NewFieldWasClauseWithDefaults() *FieldWasClause`

NewFieldWasClauseWithDefaults instantiates a new FieldWasClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldWasClause) GetField() JqlQueryField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldWasClause) GetFieldOk() (*JqlQueryField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldWasClause) SetField(v JqlQueryField)`

SetField sets Field field to given value.


### GetOperand

`func (o *FieldWasClause) GetOperand() JqlQueryClauseOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *FieldWasClause) GetOperandOk() (*JqlQueryClauseOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *FieldWasClause) SetOperand(v JqlQueryClauseOperand)`

SetOperand sets Operand field to given value.


### GetOperator

`func (o *FieldWasClause) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldWasClause) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldWasClause) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPredicates

`func (o *FieldWasClause) GetPredicates() []JqlQueryClauseTimePredicate`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *FieldWasClause) GetPredicatesOk() (*[]JqlQueryClauseTimePredicate, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *FieldWasClause) SetPredicates(v []JqlQueryClauseTimePredicate)`

SetPredicates sets Predicates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


