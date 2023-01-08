# JexpEvaluateCtxJqlIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** | The maximum number of issues to return from the JQL query. max results value considered may be lower than the number specific here. | [optional] 
**NextPageToken** | Pointer to **string** | The token for a page to fetch that is not the first page. The first page has a &#x60;nextPageToken&#x60; of &#x60;null&#x60;. Use the &#x60;nextPageToken&#x60; to fetch the next page of issues. | [optional] 
**Query** | Pointer to **string** | The JQL query, required to be bounded. Additionally, &#x60;orderBy&#x60; clause can contain a maximum of 7 fields | [optional] 

## Methods

### NewJexpEvaluateCtxJqlIssues

`func NewJexpEvaluateCtxJqlIssues() *JexpEvaluateCtxJqlIssues`

NewJexpEvaluateCtxJqlIssues instantiates a new JexpEvaluateCtxJqlIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJexpEvaluateCtxJqlIssuesWithDefaults

`func NewJexpEvaluateCtxJqlIssuesWithDefaults() *JexpEvaluateCtxJqlIssues`

NewJexpEvaluateCtxJqlIssuesWithDefaults instantiates a new JexpEvaluateCtxJqlIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *JexpEvaluateCtxJqlIssues) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *JexpEvaluateCtxJqlIssues) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *JexpEvaluateCtxJqlIssues) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *JexpEvaluateCtxJqlIssues) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPageToken

`func (o *JexpEvaluateCtxJqlIssues) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *JexpEvaluateCtxJqlIssues) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *JexpEvaluateCtxJqlIssues) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *JexpEvaluateCtxJqlIssues) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetQuery

`func (o *JexpEvaluateCtxJqlIssues) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *JexpEvaluateCtxJqlIssues) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *JexpEvaluateCtxJqlIssues) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *JexpEvaluateCtxJqlIssues) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


