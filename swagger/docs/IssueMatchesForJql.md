# IssueMatchesForJQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | **[]string** | A list of errors. | 
**MatchedIssues** | **[]int64** | A list of issue IDs. | 

## Methods

### NewIssueMatchesForJQL

`func NewIssueMatchesForJQL(errors []string, matchedIssues []int64, ) *IssueMatchesForJQL`

NewIssueMatchesForJQL instantiates a new IssueMatchesForJQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueMatchesForJQLWithDefaults

`func NewIssueMatchesForJQLWithDefaults() *IssueMatchesForJQL`

NewIssueMatchesForJQLWithDefaults instantiates a new IssueMatchesForJQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *IssueMatchesForJQL) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IssueMatchesForJQL) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IssueMatchesForJQL) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetMatchedIssues

`func (o *IssueMatchesForJQL) GetMatchedIssues() []int64`

GetMatchedIssues returns the MatchedIssues field if non-nil, zero value otherwise.

### GetMatchedIssuesOk

`func (o *IssueMatchesForJQL) GetMatchedIssuesOk() (*[]int64, bool)`

GetMatchedIssuesOk returns a tuple with the MatchedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedIssues

`func (o *IssueMatchesForJQL) SetMatchedIssues(v []int64)`

SetMatchedIssues sets MatchedIssues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


