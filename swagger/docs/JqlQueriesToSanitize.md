# JqlQueriesToSanitize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]JqlQueryToSanitize**](JqlQueryToSanitize.md) | The list of JQL queries to sanitize. Must contain unique values. Maximum of 20 queries. | 

## Methods

### NewJqlQueriesToSanitize

`func NewJqlQueriesToSanitize(queries []JqlQueryToSanitize, ) *JqlQueriesToSanitize`

NewJqlQueriesToSanitize instantiates a new JqlQueriesToSanitize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueriesToSanitizeWithDefaults

`func NewJqlQueriesToSanitizeWithDefaults() *JqlQueriesToSanitize`

NewJqlQueriesToSanitizeWithDefaults instantiates a new JqlQueriesToSanitize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *JqlQueriesToSanitize) GetQueries() []JqlQueryToSanitize`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *JqlQueriesToSanitize) GetQueriesOk() (*[]JqlQueryToSanitize, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *JqlQueriesToSanitize) SetQueries(v []JqlQueryToSanitize)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


