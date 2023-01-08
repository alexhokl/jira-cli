# ProjectInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastIssueUpdateTime** | Pointer to **time.Time** | The last issue update time. | [optional] [readonly] 
**TotalIssueCount** | Pointer to **int64** | Total issue count. | [optional] [readonly] 

## Methods

### NewProjectInsight

`func NewProjectInsight() *ProjectInsight`

NewProjectInsight instantiates a new ProjectInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInsightWithDefaults

`func NewProjectInsightWithDefaults() *ProjectInsight`

NewProjectInsightWithDefaults instantiates a new ProjectInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastIssueUpdateTime

`func (o *ProjectInsight) GetLastIssueUpdateTime() time.Time`

GetLastIssueUpdateTime returns the LastIssueUpdateTime field if non-nil, zero value otherwise.

### GetLastIssueUpdateTimeOk

`func (o *ProjectInsight) GetLastIssueUpdateTimeOk() (*time.Time, bool)`

GetLastIssueUpdateTimeOk returns a tuple with the LastIssueUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIssueUpdateTime

`func (o *ProjectInsight) SetLastIssueUpdateTime(v time.Time)`

SetLastIssueUpdateTime sets LastIssueUpdateTime field to given value.

### HasLastIssueUpdateTime

`func (o *ProjectInsight) HasLastIssueUpdateTime() bool`

HasLastIssueUpdateTime returns a boolean if a field has been set.

### GetTotalIssueCount

`func (o *ProjectInsight) GetTotalIssueCount() int64`

GetTotalIssueCount returns the TotalIssueCount field if non-nil, zero value otherwise.

### GetTotalIssueCountOk

`func (o *ProjectInsight) GetTotalIssueCountOk() (*int64, bool)`

GetTotalIssueCountOk returns a tuple with the TotalIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIssueCount

`func (o *ProjectInsight) SetTotalIssueCount(v int64)`

SetTotalIssueCount sets TotalIssueCount field to given value.

### HasTotalIssueCount

`func (o *ProjectInsight) HasTotalIssueCount() bool`

HasTotalIssueCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


