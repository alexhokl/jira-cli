# CreateExclusionRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIds** | Pointer to **[]int64** | The IDs of the issues to exclude from the plan. | [optional] 
**IssueTypeIds** | Pointer to **[]int64** | The IDs of the issue types to exclude from the plan. | [optional] 
**NumberOfDaysToShowCompletedIssues** | Pointer to **int32** | Issues completed this number of days ago will be excluded from the plan. | [optional] 
**ReleaseIds** | Pointer to **[]int64** | The IDs of the releases to exclude from the plan. | [optional] 
**WorkStatusCategoryIds** | Pointer to **[]int64** | The IDs of the work status categories to exclude from the plan. | [optional] 
**WorkStatusIds** | Pointer to **[]int64** | The IDs of the work statuses to exclude from the plan. | [optional] 

## Methods

### NewCreateExclusionRulesRequest

`func NewCreateExclusionRulesRequest() *CreateExclusionRulesRequest`

NewCreateExclusionRulesRequest instantiates a new CreateExclusionRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExclusionRulesRequestWithDefaults

`func NewCreateExclusionRulesRequestWithDefaults() *CreateExclusionRulesRequest`

NewCreateExclusionRulesRequestWithDefaults instantiates a new CreateExclusionRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueIds

`func (o *CreateExclusionRulesRequest) GetIssueIds() []int64`

GetIssueIds returns the IssueIds field if non-nil, zero value otherwise.

### GetIssueIdsOk

`func (o *CreateExclusionRulesRequest) GetIssueIdsOk() (*[]int64, bool)`

GetIssueIdsOk returns a tuple with the IssueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIds

`func (o *CreateExclusionRulesRequest) SetIssueIds(v []int64)`

SetIssueIds sets IssueIds field to given value.

### HasIssueIds

`func (o *CreateExclusionRulesRequest) HasIssueIds() bool`

HasIssueIds returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *CreateExclusionRulesRequest) GetIssueTypeIds() []int64`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *CreateExclusionRulesRequest) GetIssueTypeIdsOk() (*[]int64, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *CreateExclusionRulesRequest) SetIssueTypeIds(v []int64)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *CreateExclusionRulesRequest) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetNumberOfDaysToShowCompletedIssues

`func (o *CreateExclusionRulesRequest) GetNumberOfDaysToShowCompletedIssues() int32`

GetNumberOfDaysToShowCompletedIssues returns the NumberOfDaysToShowCompletedIssues field if non-nil, zero value otherwise.

### GetNumberOfDaysToShowCompletedIssuesOk

`func (o *CreateExclusionRulesRequest) GetNumberOfDaysToShowCompletedIssuesOk() (*int32, bool)`

GetNumberOfDaysToShowCompletedIssuesOk returns a tuple with the NumberOfDaysToShowCompletedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysToShowCompletedIssues

`func (o *CreateExclusionRulesRequest) SetNumberOfDaysToShowCompletedIssues(v int32)`

SetNumberOfDaysToShowCompletedIssues sets NumberOfDaysToShowCompletedIssues field to given value.

### HasNumberOfDaysToShowCompletedIssues

`func (o *CreateExclusionRulesRequest) HasNumberOfDaysToShowCompletedIssues() bool`

HasNumberOfDaysToShowCompletedIssues returns a boolean if a field has been set.

### GetReleaseIds

`func (o *CreateExclusionRulesRequest) GetReleaseIds() []int64`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *CreateExclusionRulesRequest) GetReleaseIdsOk() (*[]int64, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *CreateExclusionRulesRequest) SetReleaseIds(v []int64)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *CreateExclusionRulesRequest) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.

### GetWorkStatusCategoryIds

`func (o *CreateExclusionRulesRequest) GetWorkStatusCategoryIds() []int64`

GetWorkStatusCategoryIds returns the WorkStatusCategoryIds field if non-nil, zero value otherwise.

### GetWorkStatusCategoryIdsOk

`func (o *CreateExclusionRulesRequest) GetWorkStatusCategoryIdsOk() (*[]int64, bool)`

GetWorkStatusCategoryIdsOk returns a tuple with the WorkStatusCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkStatusCategoryIds

`func (o *CreateExclusionRulesRequest) SetWorkStatusCategoryIds(v []int64)`

SetWorkStatusCategoryIds sets WorkStatusCategoryIds field to given value.

### HasWorkStatusCategoryIds

`func (o *CreateExclusionRulesRequest) HasWorkStatusCategoryIds() bool`

HasWorkStatusCategoryIds returns a boolean if a field has been set.

### GetWorkStatusIds

`func (o *CreateExclusionRulesRequest) GetWorkStatusIds() []int64`

GetWorkStatusIds returns the WorkStatusIds field if non-nil, zero value otherwise.

### GetWorkStatusIdsOk

`func (o *CreateExclusionRulesRequest) GetWorkStatusIdsOk() (*[]int64, bool)`

GetWorkStatusIdsOk returns a tuple with the WorkStatusIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkStatusIds

`func (o *CreateExclusionRulesRequest) SetWorkStatusIds(v []int64)`

SetWorkStatusIds sets WorkStatusIds field to given value.

### HasWorkStatusIds

`func (o *CreateExclusionRulesRequest) HasWorkStatusIds() bool`

HasWorkStatusIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


