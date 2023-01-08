# WorkflowProjectIssueTypeUsagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token for the next page of issue type usages. | [optional] 
**Values** | Pointer to [**[]WorkflowProjectIssueTypeUsage**](WorkflowProjectIssueTypeUsage.md) | The list of issue types. | [optional] 

## Methods

### NewWorkflowProjectIssueTypeUsagePage

`func NewWorkflowProjectIssueTypeUsagePage() *WorkflowProjectIssueTypeUsagePage`

NewWorkflowProjectIssueTypeUsagePage instantiates a new WorkflowProjectIssueTypeUsagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowProjectIssueTypeUsagePageWithDefaults

`func NewWorkflowProjectIssueTypeUsagePageWithDefaults() *WorkflowProjectIssueTypeUsagePage`

NewWorkflowProjectIssueTypeUsagePageWithDefaults instantiates a new WorkflowProjectIssueTypeUsagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *WorkflowProjectIssueTypeUsagePage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WorkflowProjectIssueTypeUsagePage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WorkflowProjectIssueTypeUsagePage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *WorkflowProjectIssueTypeUsagePage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetValues

`func (o *WorkflowProjectIssueTypeUsagePage) GetValues() []WorkflowProjectIssueTypeUsage`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkflowProjectIssueTypeUsagePage) GetValuesOk() (*[]WorkflowProjectIssueTypeUsage, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkflowProjectIssueTypeUsagePage) SetValues(v []WorkflowProjectIssueTypeUsage)`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkflowProjectIssueTypeUsagePage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


