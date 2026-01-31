# RepositoryErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) | Repository errors | [optional] 
**Commits** | Pointer to [**[]EntityError**](EntityError.md) | Commits errors | [optional] 
**Branches** | Pointer to [**[]EntityError**](EntityError.md) | Branches errors | [optional] 
**PullRequests** | Pointer to [**[]EntityError**](EntityError.md) | Pull requests errors | [optional] 

## Methods

### NewRepositoryErrors

`func NewRepositoryErrors() *RepositoryErrors`

NewRepositoryErrors instantiates a new RepositoryErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryErrorsWithDefaults

`func NewRepositoryErrorsWithDefaults() *RepositoryErrors`

NewRepositoryErrorsWithDefaults instantiates a new RepositoryErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *RepositoryErrors) GetErrorMessages() []ErrorMessage`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *RepositoryErrors) GetErrorMessagesOk() (*[]ErrorMessage, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *RepositoryErrors) SetErrorMessages(v []ErrorMessage)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *RepositoryErrors) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetCommits

`func (o *RepositoryErrors) GetCommits() []EntityError`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *RepositoryErrors) GetCommitsOk() (*[]EntityError, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *RepositoryErrors) SetCommits(v []EntityError)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *RepositoryErrors) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetBranches

`func (o *RepositoryErrors) GetBranches() []EntityError`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *RepositoryErrors) GetBranchesOk() (*[]EntityError, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *RepositoryErrors) SetBranches(v []EntityError)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *RepositoryErrors) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetPullRequests

`func (o *RepositoryErrors) GetPullRequests() []EntityError`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *RepositoryErrors) GetPullRequestsOk() (*[]EntityError, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *RepositoryErrors) SetPullRequests(v []EntityError)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *RepositoryErrors) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


