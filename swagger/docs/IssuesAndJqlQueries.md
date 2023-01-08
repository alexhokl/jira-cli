# IssuesAndJQLQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIds** | **[]int64** | A list of issue IDs. | 
**Jqls** | **[]string** | A list of JQL queries. | 

## Methods

### NewIssuesAndJQLQueries

`func NewIssuesAndJQLQueries(issueIds []int64, jqls []string, ) *IssuesAndJQLQueries`

NewIssuesAndJQLQueries instantiates a new IssuesAndJQLQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesAndJQLQueriesWithDefaults

`func NewIssuesAndJQLQueriesWithDefaults() *IssuesAndJQLQueries`

NewIssuesAndJQLQueriesWithDefaults instantiates a new IssuesAndJQLQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueIds

`func (o *IssuesAndJQLQueries) GetIssueIds() []int64`

GetIssueIds returns the IssueIds field if non-nil, zero value otherwise.

### GetIssueIdsOk

`func (o *IssuesAndJQLQueries) GetIssueIdsOk() (*[]int64, bool)`

GetIssueIdsOk returns a tuple with the IssueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIds

`func (o *IssuesAndJQLQueries) SetIssueIds(v []int64)`

SetIssueIds sets IssueIds field to given value.


### GetJqls

`func (o *IssuesAndJQLQueries) GetJqls() []string`

GetJqls returns the Jqls field if non-nil, zero value otherwise.

### GetJqlsOk

`func (o *IssuesAndJQLQueries) GetJqlsOk() (*[]string, bool)`

GetJqlsOk returns a tuple with the Jqls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqls

`func (o *IssuesAndJQLQueries) SetJqls(v []string)`

SetJqls sets Jqls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


