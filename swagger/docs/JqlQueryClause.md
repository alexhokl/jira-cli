# JqlQueryClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clauses** | [**[]JqlQueryClause**](JqlQueryClause.md) | The list of nested clauses. | 
**Operator** | **string** | The operator applied to the field. | 
**Field** | [**JqlQueryField**](JqlQueryField.md) |  | 
**Operand** | [**JqlQueryClauseOperand**](JqlQueryClauseOperand.md) |  | 
**Predicates** | [**[]JqlQueryClauseTimePredicate**](JqlQueryClauseTimePredicate.md) | The list of time predicates. | 

## Methods

### NewJqlQueryClause

`func NewJqlQueryClause(clauses []JqlQueryClause, operator string, field JqlQueryField, operand JqlQueryClauseOperand, predicates []JqlQueryClauseTimePredicate, ) *JqlQueryClause`

NewJqlQueryClause instantiates a new JqlQueryClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryClauseWithDefaults

`func NewJqlQueryClauseWithDefaults() *JqlQueryClause`

NewJqlQueryClauseWithDefaults instantiates a new JqlQueryClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClauses

`func (o *JqlQueryClause) GetClauses() []JqlQueryClause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *JqlQueryClause) GetClausesOk() (*[]JqlQueryClause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *JqlQueryClause) SetClauses(v []JqlQueryClause)`

SetClauses sets Clauses field to given value.


### GetOperator

`func (o *JqlQueryClause) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JqlQueryClause) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JqlQueryClause) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetField

`func (o *JqlQueryClause) GetField() JqlQueryField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *JqlQueryClause) GetFieldOk() (*JqlQueryField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *JqlQueryClause) SetField(v JqlQueryField)`

SetField sets Field field to given value.


### GetOperand

`func (o *JqlQueryClause) GetOperand() JqlQueryClauseOperand`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *JqlQueryClause) GetOperandOk() (*JqlQueryClauseOperand, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *JqlQueryClause) SetOperand(v JqlQueryClauseOperand)`

SetOperand sets Operand field to given value.


### GetPredicates

`func (o *JqlQueryClause) GetPredicates() []JqlQueryClauseTimePredicate`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *JqlQueryClause) GetPredicatesOk() (*[]JqlQueryClauseTimePredicate, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *JqlQueryClause) SetPredicates(v []JqlQueryClauseTimePredicate)`

SetPredicates sets Predicates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


