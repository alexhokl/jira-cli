# JexpJqlIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** | The maximum number of issues to return from the JQL query. Inspect &#x60;meta.issues.jql.maxResults&#x60; in the response to ensure the maximum value has not been exceeded. | [optional] 
**Query** | Pointer to **string** | The JQL query. | [optional] 
**StartAt** | Pointer to **int64** | The index of the first issue to return from the JQL query. | [optional] 
**Validation** | Pointer to **string** | Determines how to validate the JQL query and treat the validation results. | [optional] [default to "strict"]

## Methods

### NewJexpJqlIssues

`func NewJexpJqlIssues() *JexpJqlIssues`

NewJexpJqlIssues instantiates a new JexpJqlIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJexpJqlIssuesWithDefaults

`func NewJexpJqlIssuesWithDefaults() *JexpJqlIssues`

NewJexpJqlIssuesWithDefaults instantiates a new JexpJqlIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *JexpJqlIssues) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *JexpJqlIssues) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *JexpJqlIssues) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *JexpJqlIssues) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetQuery

`func (o *JexpJqlIssues) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *JexpJqlIssues) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *JexpJqlIssues) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *JexpJqlIssues) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetStartAt

`func (o *JexpJqlIssues) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *JexpJqlIssues) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *JexpJqlIssues) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *JexpJqlIssues) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetValidation

`func (o *JexpJqlIssues) GetValidation() string`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *JexpJqlIssues) GetValidationOk() (*string, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *JexpJqlIssues) SetValidation(v string)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *JexpJqlIssues) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


