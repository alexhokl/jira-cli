# CompoundClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clauses** | [**[]JqlQueryClause**](JqlQueryClause.md) | The list of nested clauses. | 
**Operator** | **string** | The operator between the clauses. | 

## Methods

### NewCompoundClause

`func NewCompoundClause(clauses []JqlQueryClause, operator string, ) *CompoundClause`

NewCompoundClause instantiates a new CompoundClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompoundClauseWithDefaults

`func NewCompoundClauseWithDefaults() *CompoundClause`

NewCompoundClauseWithDefaults instantiates a new CompoundClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClauses

`func (o *CompoundClause) GetClauses() []JqlQueryClause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *CompoundClause) GetClausesOk() (*[]JqlQueryClause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *CompoundClause) SetClauses(v []JqlQueryClause)`

SetClauses sets Clauses field to given value.


### GetOperator

`func (o *CompoundClause) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CompoundClause) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CompoundClause) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


