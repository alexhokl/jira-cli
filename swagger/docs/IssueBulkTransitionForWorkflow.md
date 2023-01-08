# IssueBulkTransitionForWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTransitionsFiltered** | Pointer to **bool** | Indicates whether all the transitions of this workflow are available in the transitions list or not. | [optional] [readonly] 
**Issues** | Pointer to **[]string** | List of issue keys from the request which are associated with this workflow. | [optional] [readonly] 
**Transitions** | Pointer to [**[]SimplifiedIssueTransition**](SimplifiedIssueTransition.md) | List of transitions available for issues from the request which are associated with this workflow.   **This list includes only those transitions that are common across the issues in this workflow and do not involve any additional field updates.**  | [optional] [readonly] 

## Methods

### NewIssueBulkTransitionForWorkflow

`func NewIssueBulkTransitionForWorkflow() *IssueBulkTransitionForWorkflow`

NewIssueBulkTransitionForWorkflow instantiates a new IssueBulkTransitionForWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkTransitionForWorkflowWithDefaults

`func NewIssueBulkTransitionForWorkflowWithDefaults() *IssueBulkTransitionForWorkflow`

NewIssueBulkTransitionForWorkflowWithDefaults instantiates a new IssueBulkTransitionForWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTransitionsFiltered

`func (o *IssueBulkTransitionForWorkflow) GetIsTransitionsFiltered() bool`

GetIsTransitionsFiltered returns the IsTransitionsFiltered field if non-nil, zero value otherwise.

### GetIsTransitionsFilteredOk

`func (o *IssueBulkTransitionForWorkflow) GetIsTransitionsFilteredOk() (*bool, bool)`

GetIsTransitionsFilteredOk returns a tuple with the IsTransitionsFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransitionsFiltered

`func (o *IssueBulkTransitionForWorkflow) SetIsTransitionsFiltered(v bool)`

SetIsTransitionsFiltered sets IsTransitionsFiltered field to given value.

### HasIsTransitionsFiltered

`func (o *IssueBulkTransitionForWorkflow) HasIsTransitionsFiltered() bool`

HasIsTransitionsFiltered returns a boolean if a field has been set.

### GetIssues

`func (o *IssueBulkTransitionForWorkflow) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *IssueBulkTransitionForWorkflow) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *IssueBulkTransitionForWorkflow) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *IssueBulkTransitionForWorkflow) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetTransitions

`func (o *IssueBulkTransitionForWorkflow) GetTransitions() []SimplifiedIssueTransition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *IssueBulkTransitionForWorkflow) GetTransitionsOk() (*[]SimplifiedIssueTransition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *IssueBulkTransitionForWorkflow) SetTransitions(v []SimplifiedIssueTransition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *IssueBulkTransitionForWorkflow) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


