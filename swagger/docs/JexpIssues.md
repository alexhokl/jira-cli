# JexpIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jql** | Pointer to [**JexpJqlIssues**](JexpJqlIssues.md) | The JQL query that specifies the set of issues available in the Jira expression. | [optional] 

## Methods

### NewJexpIssues

`func NewJexpIssues() *JexpIssues`

NewJexpIssues instantiates a new JexpIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJexpIssuesWithDefaults

`func NewJexpIssuesWithDefaults() *JexpIssues`

NewJexpIssuesWithDefaults instantiates a new JexpIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJql

`func (o *JexpIssues) GetJql() JexpJqlIssues`

GetJql returns the Jql field if non-nil, zero value otherwise.

### GetJqlOk

`func (o *JexpIssues) GetJqlOk() (*JexpJqlIssues, bool)`

GetJqlOk returns a tuple with the Jql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJql

`func (o *JexpIssues) SetJql(v JexpJqlIssues)`

SetJql sets Jql field to given value.

### HasJql

`func (o *JexpIssues) HasJql() bool`

HasJql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


