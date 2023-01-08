# GetExclusionRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIds** | Pointer to **[]int64** | The IDs of the issues excluded from the plan. | [optional] 
**IssueTypeIds** | Pointer to **[]int64** | The IDs of the issue types excluded from the plan. | [optional] 
**NumberOfDaysToShowCompletedIssues** | **int32** | Issues completed this number of days ago are excluded from the plan. | 
**ReleaseIds** | Pointer to **[]int64** | The IDs of the releases excluded from the plan. | [optional] 
**WorkStatusCategoryIds** | Pointer to **[]int64** | The IDs of the work status categories excluded from the plan. | [optional] 
**WorkStatusIds** | Pointer to **[]int64** | The IDs of the work statuses excluded from the plan. | [optional] 

## Methods

### NewGetExclusionRulesResponse

`func NewGetExclusionRulesResponse(numberOfDaysToShowCompletedIssues int32, ) *GetExclusionRulesResponse`

NewGetExclusionRulesResponse instantiates a new GetExclusionRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExclusionRulesResponseWithDefaults

`func NewGetExclusionRulesResponseWithDefaults() *GetExclusionRulesResponse`

NewGetExclusionRulesResponseWithDefaults instantiates a new GetExclusionRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueIds

`func (o *GetExclusionRulesResponse) GetIssueIds() []int64`

GetIssueIds returns the IssueIds field if non-nil, zero value otherwise.

### GetIssueIdsOk

`func (o *GetExclusionRulesResponse) GetIssueIdsOk() (*[]int64, bool)`

GetIssueIdsOk returns a tuple with the IssueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIds

`func (o *GetExclusionRulesResponse) SetIssueIds(v []int64)`

SetIssueIds sets IssueIds field to given value.

### HasIssueIds

`func (o *GetExclusionRulesResponse) HasIssueIds() bool`

HasIssueIds returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *GetExclusionRulesResponse) GetIssueTypeIds() []int64`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *GetExclusionRulesResponse) GetIssueTypeIdsOk() (*[]int64, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *GetExclusionRulesResponse) SetIssueTypeIds(v []int64)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *GetExclusionRulesResponse) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetNumberOfDaysToShowCompletedIssues

`func (o *GetExclusionRulesResponse) GetNumberOfDaysToShowCompletedIssues() int32`

GetNumberOfDaysToShowCompletedIssues returns the NumberOfDaysToShowCompletedIssues field if non-nil, zero value otherwise.

### GetNumberOfDaysToShowCompletedIssuesOk

`func (o *GetExclusionRulesResponse) GetNumberOfDaysToShowCompletedIssuesOk() (*int32, bool)`

GetNumberOfDaysToShowCompletedIssuesOk returns a tuple with the NumberOfDaysToShowCompletedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysToShowCompletedIssues

`func (o *GetExclusionRulesResponse) SetNumberOfDaysToShowCompletedIssues(v int32)`

SetNumberOfDaysToShowCompletedIssues sets NumberOfDaysToShowCompletedIssues field to given value.


### GetReleaseIds

`func (o *GetExclusionRulesResponse) GetReleaseIds() []int64`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *GetExclusionRulesResponse) GetReleaseIdsOk() (*[]int64, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *GetExclusionRulesResponse) SetReleaseIds(v []int64)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *GetExclusionRulesResponse) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.

### GetWorkStatusCategoryIds

`func (o *GetExclusionRulesResponse) GetWorkStatusCategoryIds() []int64`

GetWorkStatusCategoryIds returns the WorkStatusCategoryIds field if non-nil, zero value otherwise.

### GetWorkStatusCategoryIdsOk

`func (o *GetExclusionRulesResponse) GetWorkStatusCategoryIdsOk() (*[]int64, bool)`

GetWorkStatusCategoryIdsOk returns a tuple with the WorkStatusCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkStatusCategoryIds

`func (o *GetExclusionRulesResponse) SetWorkStatusCategoryIds(v []int64)`

SetWorkStatusCategoryIds sets WorkStatusCategoryIds field to given value.

### HasWorkStatusCategoryIds

`func (o *GetExclusionRulesResponse) HasWorkStatusCategoryIds() bool`

HasWorkStatusCategoryIds returns a boolean if a field has been set.

### GetWorkStatusIds

`func (o *GetExclusionRulesResponse) GetWorkStatusIds() []int64`

GetWorkStatusIds returns the WorkStatusIds field if non-nil, zero value otherwise.

### GetWorkStatusIdsOk

`func (o *GetExclusionRulesResponse) GetWorkStatusIdsOk() (*[]int64, bool)`

GetWorkStatusIdsOk returns a tuple with the WorkStatusIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkStatusIds

`func (o *GetExclusionRulesResponse) SetWorkStatusIds(v []int64)`

SetWorkStatusIds sets WorkStatusIds field to given value.

### HasWorkStatusIds

`func (o *GetExclusionRulesResponse) HasWorkStatusIds() bool`

HasWorkStatusIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


