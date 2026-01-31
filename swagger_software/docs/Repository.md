# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this repository. Max length is 255 characters. | 
**Description** | Pointer to **string** | Description of this repository. Max length is 1024 characters. | [optional] 
**ForkOf** | Pointer to **string** | The ID of the repository this repository was forked from, if it&#39;s a fork. Max length is 1024 characters. | [optional] 
**Url** | **string** | The URL of this repository. Max length is 2000 characters. | 
**Commits** | Pointer to [**[]Commit**](Commit.md) | List of commits to update in this repository. Must not contain duplicate entity IDs. Maximum number of commits is 400 | [optional] 
**Branches** | Pointer to [**[]Branch**](Branch.md) | List of branches to update in this repository. Must not contain duplicate entity IDs. Maximum number of branches is 400. | [optional] 
**PullRequests** | Pointer to [**[]PullRequest**](PullRequest.md) | List of pull requests to update in this repository. Must not contain duplicate entity IDs. Maximum number of pull requests is 400 | [optional] 
**Avatar** | Pointer to **string** | The URL of the avatar for this repository. Max length is 2000 characters. | [optional] 
**AvatarDescription** | Pointer to **string** | Description of the avatar for this repository. Max length is 1024 characters. | [optional] 
**Id** | **string** | The ID of this entity. Will be used for cross entity linking. Must be unique by entity type within a repository, i.e., only one commit can have ID &#39;X&#39; in repository &#39;Y&#39;. But adding, e.g., a branch with ID &#39;X&#39; to repository &#39;Y&#39; is acceptable. Only alphanumeric characters, and &#39;~.-_&#39;, are allowed. Max length is 1024 characters. | 
**UpdateSequenceId** | **int64** |  An ID used to apply an ordering to updates for this entity in the case of out-of-order receipt of update requests. This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the provider system, but other alternatives are valid (e.g. a provider could store a counter against each entity and increment that on each update to Jira). Updates for an entity that are received with an updateSqeuenceId lower than what is currently stored will be ignored. | 

## Methods

### NewRepository

`func NewRepository(name string, url string, id string, updateSequenceId int64, ) *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Repository) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Repository) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Repository) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Repository) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForkOf

`func (o *Repository) GetForkOf() string`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *Repository) GetForkOfOk() (*string, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *Repository) SetForkOf(v string)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *Repository) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.

### GetUrl

`func (o *Repository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCommits

`func (o *Repository) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *Repository) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *Repository) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *Repository) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetBranches

`func (o *Repository) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *Repository) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *Repository) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *Repository) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetPullRequests

`func (o *Repository) GetPullRequests() []PullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *Repository) GetPullRequestsOk() (*[]PullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *Repository) SetPullRequests(v []PullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *Repository) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetAvatar

`func (o *Repository) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Repository) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Repository) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Repository) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAvatarDescription

`func (o *Repository) GetAvatarDescription() string`

GetAvatarDescription returns the AvatarDescription field if non-nil, zero value otherwise.

### GetAvatarDescriptionOk

`func (o *Repository) GetAvatarDescriptionOk() (*string, bool)`

GetAvatarDescriptionOk returns a tuple with the AvatarDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarDescription

`func (o *Repository) SetAvatarDescription(v string)`

SetAvatarDescription sets AvatarDescription field to given value.

### HasAvatarDescription

`func (o *Repository) HasAvatarDescription() bool`

HasAvatarDescription returns a boolean if a field has been set.

### GetId

`func (o *Repository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateSequenceId

`func (o *Repository) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *Repository) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *Repository) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


