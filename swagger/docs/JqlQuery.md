# JqlQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderBy** | Pointer to [**JqlQueryOrderByClause**](JqlQueryOrderByClause.md) |  | [optional] 
**Where** | Pointer to [**JqlQueryClause**](JqlQueryClause.md) |  | [optional] 

## Methods

### NewJqlQuery

`func NewJqlQuery() *JqlQuery`

NewJqlQuery instantiates a new JqlQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryWithDefaults

`func NewJqlQueryWithDefaults() *JqlQuery`

NewJqlQueryWithDefaults instantiates a new JqlQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderBy

`func (o *JqlQuery) GetOrderBy() JqlQueryOrderByClause`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *JqlQuery) GetOrderByOk() (*JqlQueryOrderByClause, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *JqlQuery) SetOrderBy(v JqlQueryOrderByClause)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *JqlQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetWhere

`func (o *JqlQuery) GetWhere() JqlQueryClause`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *JqlQuery) GetWhereOk() (*JqlQueryClause, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *JqlQuery) SetWhere(v JqlQueryClause)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *JqlQuery) HasWhere() bool`

HasWhere returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


