# FieldChangedClause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | [**JqlQueryField**](JqlQueryField.md) |  | 
**Operator** | **string** | The operator applied to the field. | 
**Predicates** | [**[]JqlQueryClauseTimePredicate**](JqlQueryClauseTimePredicate.md) | The list of time predicates. | 

## Methods

### NewFieldChangedClause

`func NewFieldChangedClause(field JqlQueryField, operator string, predicates []JqlQueryClauseTimePredicate, ) *FieldChangedClause`

NewFieldChangedClause instantiates a new FieldChangedClause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldChangedClauseWithDefaults

`func NewFieldChangedClauseWithDefaults() *FieldChangedClause`

NewFieldChangedClauseWithDefaults instantiates a new FieldChangedClause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldChangedClause) GetField() JqlQueryField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldChangedClause) GetFieldOk() (*JqlQueryField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldChangedClause) SetField(v JqlQueryField)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldChangedClause) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldChangedClause) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldChangedClause) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPredicates

`func (o *FieldChangedClause) GetPredicates() []JqlQueryClauseTimePredicate`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *FieldChangedClause) GetPredicatesOk() (*[]JqlQueryClauseTimePredicate, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *FieldChangedClause) SetPredicates(v []JqlQueryClauseTimePredicate)`

SetPredicates sets Predicates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


