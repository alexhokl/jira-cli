# BulkOperationProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | A timestamp of when the task was submitted. | [optional] 
**FailedAccessibleIssues** | Pointer to **map[string][]string** | Map of issue IDs for which the operation failed and that the user has permission to view, to their one or more reasons for failure. These reasons are open-ended text descriptions of the error and are not selected from a predefined list of standard reasons. | [optional] 
**InvalidOrInaccessibleIssueCount** | Pointer to **int32** | The number of issues that are either invalid or issues that the user doesn&#39;t have permission to view, regardless of the success or failure of the operation. | [optional] 
**ProcessedAccessibleIssues** | Pointer to **[]int64** | List of issue IDs for which the operation was successful and that the user has permission to view. | [optional] 
**ProgressPercent** | Pointer to **int64** | Progress of the task as a percentage. | [optional] 
**Started** | Pointer to **time.Time** | A timestamp of when the task was started. | [optional] 
**Status** | Pointer to **string** | The status of the task. | [optional] 
**SubmittedBy** | Pointer to [**User**](User.md) |  | [optional] 
**TaskId** | Pointer to **string** | The ID of the task. | [optional] [readonly] 
**TotalIssueCount** | Pointer to **int32** | The number of issues that the bulk operation was attempted on. | [optional] 
**Updated** | Pointer to **time.Time** | A timestamp of when the task progress was last updated. | [optional] 

## Methods

### NewBulkOperationProgress

`func NewBulkOperationProgress() *BulkOperationProgress`

NewBulkOperationProgress instantiates a new BulkOperationProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationProgressWithDefaults

`func NewBulkOperationProgressWithDefaults() *BulkOperationProgress`

NewBulkOperationProgressWithDefaults instantiates a new BulkOperationProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BulkOperationProgress) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BulkOperationProgress) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BulkOperationProgress) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BulkOperationProgress) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFailedAccessibleIssues

`func (o *BulkOperationProgress) GetFailedAccessibleIssues() map[string][]string`

GetFailedAccessibleIssues returns the FailedAccessibleIssues field if non-nil, zero value otherwise.

### GetFailedAccessibleIssuesOk

`func (o *BulkOperationProgress) GetFailedAccessibleIssuesOk() (*map[string][]string, bool)`

GetFailedAccessibleIssuesOk returns a tuple with the FailedAccessibleIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccessibleIssues

`func (o *BulkOperationProgress) SetFailedAccessibleIssues(v map[string][]string)`

SetFailedAccessibleIssues sets FailedAccessibleIssues field to given value.

### HasFailedAccessibleIssues

`func (o *BulkOperationProgress) HasFailedAccessibleIssues() bool`

HasFailedAccessibleIssues returns a boolean if a field has been set.

### GetInvalidOrInaccessibleIssueCount

`func (o *BulkOperationProgress) GetInvalidOrInaccessibleIssueCount() int32`

GetInvalidOrInaccessibleIssueCount returns the InvalidOrInaccessibleIssueCount field if non-nil, zero value otherwise.

### GetInvalidOrInaccessibleIssueCountOk

`func (o *BulkOperationProgress) GetInvalidOrInaccessibleIssueCountOk() (*int32, bool)`

GetInvalidOrInaccessibleIssueCountOk returns a tuple with the InvalidOrInaccessibleIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidOrInaccessibleIssueCount

`func (o *BulkOperationProgress) SetInvalidOrInaccessibleIssueCount(v int32)`

SetInvalidOrInaccessibleIssueCount sets InvalidOrInaccessibleIssueCount field to given value.

### HasInvalidOrInaccessibleIssueCount

`func (o *BulkOperationProgress) HasInvalidOrInaccessibleIssueCount() bool`

HasInvalidOrInaccessibleIssueCount returns a boolean if a field has been set.

### GetProcessedAccessibleIssues

`func (o *BulkOperationProgress) GetProcessedAccessibleIssues() []int64`

GetProcessedAccessibleIssues returns the ProcessedAccessibleIssues field if non-nil, zero value otherwise.

### GetProcessedAccessibleIssuesOk

`func (o *BulkOperationProgress) GetProcessedAccessibleIssuesOk() (*[]int64, bool)`

GetProcessedAccessibleIssuesOk returns a tuple with the ProcessedAccessibleIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAccessibleIssues

`func (o *BulkOperationProgress) SetProcessedAccessibleIssues(v []int64)`

SetProcessedAccessibleIssues sets ProcessedAccessibleIssues field to given value.

### HasProcessedAccessibleIssues

`func (o *BulkOperationProgress) HasProcessedAccessibleIssues() bool`

HasProcessedAccessibleIssues returns a boolean if a field has been set.

### GetProgressPercent

`func (o *BulkOperationProgress) GetProgressPercent() int64`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *BulkOperationProgress) GetProgressPercentOk() (*int64, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *BulkOperationProgress) SetProgressPercent(v int64)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *BulkOperationProgress) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetStarted

`func (o *BulkOperationProgress) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *BulkOperationProgress) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *BulkOperationProgress) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *BulkOperationProgress) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *BulkOperationProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkOperationProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkOperationProgress) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkOperationProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmittedBy

`func (o *BulkOperationProgress) GetSubmittedBy() User`

GetSubmittedBy returns the SubmittedBy field if non-nil, zero value otherwise.

### GetSubmittedByOk

`func (o *BulkOperationProgress) GetSubmittedByOk() (*User, bool)`

GetSubmittedByOk returns a tuple with the SubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedBy

`func (o *BulkOperationProgress) SetSubmittedBy(v User)`

SetSubmittedBy sets SubmittedBy field to given value.

### HasSubmittedBy

`func (o *BulkOperationProgress) HasSubmittedBy() bool`

HasSubmittedBy returns a boolean if a field has been set.

### GetTaskId

`func (o *BulkOperationProgress) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BulkOperationProgress) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BulkOperationProgress) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *BulkOperationProgress) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTotalIssueCount

`func (o *BulkOperationProgress) GetTotalIssueCount() int32`

GetTotalIssueCount returns the TotalIssueCount field if non-nil, zero value otherwise.

### GetTotalIssueCountOk

`func (o *BulkOperationProgress) GetTotalIssueCountOk() (*int32, bool)`

GetTotalIssueCountOk returns a tuple with the TotalIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIssueCount

`func (o *BulkOperationProgress) SetTotalIssueCount(v int32)`

SetTotalIssueCount sets TotalIssueCount field to given value.

### HasTotalIssueCount

`func (o *BulkOperationProgress) HasTotalIssueCount() bool`

HasTotalIssueCount returns a boolean if a field has been set.

### GetUpdated

`func (o *BulkOperationProgress) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BulkOperationProgress) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BulkOperationProgress) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BulkOperationProgress) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


