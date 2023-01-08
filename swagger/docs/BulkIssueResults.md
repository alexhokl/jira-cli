# BulkIssueResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueErrors** | Pointer to [**[]IssueError**](IssueError.md) | When Jira can&#39;t return an issue enumerated in a request due to a retriable error or payload constraint, we&#39;ll return the respective issue ID with a corresponding error message. This list is empty when there are no errors Issues which aren&#39;t found or that the user doesn&#39;t have permission to view won&#39;t be returned in this list. | [optional] [readonly] 
**Issues** | Pointer to [**[]IssueBean**](IssueBean.md) | The list of issues. | [optional] [readonly] 

## Methods

### NewBulkIssueResults

`func NewBulkIssueResults() *BulkIssueResults`

NewBulkIssueResults instantiates a new BulkIssueResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssueResultsWithDefaults

`func NewBulkIssueResultsWithDefaults() *BulkIssueResults`

NewBulkIssueResultsWithDefaults instantiates a new BulkIssueResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueErrors

`func (o *BulkIssueResults) GetIssueErrors() []IssueError`

GetIssueErrors returns the IssueErrors field if non-nil, zero value otherwise.

### GetIssueErrorsOk

`func (o *BulkIssueResults) GetIssueErrorsOk() (*[]IssueError, bool)`

GetIssueErrorsOk returns a tuple with the IssueErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueErrors

`func (o *BulkIssueResults) SetIssueErrors(v []IssueError)`

SetIssueErrors sets IssueErrors field to given value.

### HasIssueErrors

`func (o *BulkIssueResults) HasIssueErrors() bool`

HasIssueErrors returns a boolean if a field has been set.

### GetIssues

`func (o *BulkIssueResults) GetIssues() []IssueBean`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *BulkIssueResults) GetIssuesOk() (*[]IssueBean, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *BulkIssueResults) SetIssues(v []IssueBean)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *BulkIssueResults) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


