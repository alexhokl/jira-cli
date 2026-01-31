# EntityIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to **[]string** | Commits IDs | [optional] 
**Branches** | Pointer to **[]string** | Branch IDs | [optional] 
**PullRequests** | Pointer to **[]string** | Pull request IDs | [optional] 

## Methods

### NewEntityIds

`func NewEntityIds() *EntityIds`

NewEntityIds instantiates a new EntityIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityIdsWithDefaults

`func NewEntityIdsWithDefaults() *EntityIds`

NewEntityIdsWithDefaults instantiates a new EntityIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *EntityIds) GetCommits() []string`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *EntityIds) GetCommitsOk() (*[]string, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *EntityIds) SetCommits(v []string)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *EntityIds) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetBranches

`func (o *EntityIds) GetBranches() []string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *EntityIds) GetBranchesOk() (*[]string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *EntityIds) SetBranches(v []string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *EntityIds) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetPullRequests

`func (o *EntityIds) GetPullRequests() []string`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *EntityIds) GetPullRequestsOk() (*[]string, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *EntityIds) SetPullRequests(v []string)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *EntityIds) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


