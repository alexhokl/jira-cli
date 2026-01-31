# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this entity. Will be used for cross entity linking. Must be unique by entity type within a repository, i.e., only one commit can have ID &#39;X&#39; in repository &#39;Y&#39;. But adding, e.g., a branch with ID &#39;X&#39; to repository &#39;Y&#39; is acceptable. Only alphanumeric characters, and &#39;~.-_&#39;, are allowed. Max length is 1024 characters | 
**IssueKeys** | Pointer to **[]string** | List of issues keys that this entity is associated with. They must be valid Jira issue keys. | [optional] 
**Associations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | The Jira issue keys or IDs to associate the pull request with. | [optional] 
**UpdateSequenceId** | **int64** | An ID used to apply an ordering to updates for this entity in the case of out-of-order receipt of update requests. This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the provider system, but other alternatives are valid (e.g. a provider could store a counter against each entity and increment that on each update to Jira). Updates for an entity that are received with an updateSqeuenceId lower than what is currently stored will be ignored. | 
**Status** | **string** | The status of the pull request. In the case of concurrent updates, priority is given in the order OPEN, MERGED, DECLINED, UNKNOWN | 
**Title** | **string** | Title of the pull request. Max length is 1024 characters. | 
**Author** | [**Author**](Author.md) |  | 
**CommentCount** | **int32** | The number of comments on the pull request | 
**SourceBranch** | **string** | The name of the source branch of this PR. Max length is 255 characters. | 
**SourceBranchUrl** | Pointer to **string** | The url of the source branch of this PR. This is used to match this PR against the branch. Max length is 2000 characters. | [optional] 
**LastUpdate** | **string** | The most recent update to this PR. Formatted as a UTC ISO 8601 date time format. | 
**DestinationBranch** | Pointer to **string** | The name of destination branch of this PR. Max length is 255 characters. | [optional] 
**DestinationBranchUrl** | Pointer to **string** | The url of the destination branch of this PR. Max length is 2000 characters. | [optional] 
**Reviewers** | Pointer to [**[]Reviewer**](Reviewer.md) | The list of reviewers of this pull request | [optional] 
**Url** | **string** | The URL of this pull request. Max length is 2000 characters. | 
**DisplayId** | **string** | Shortened identifier for this pull request, used for display. Max length is 255 characters. | 

## Methods

### NewPullRequest

`func NewPullRequest(id string, updateSequenceId int64, status string, title string, author Author, commentCount int32, sourceBranch string, lastUpdate string, url string, displayId string, ) *PullRequest`

NewPullRequest instantiates a new PullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestWithDefaults

`func NewPullRequestWithDefaults() *PullRequest`

NewPullRequestWithDefaults instantiates a new PullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIssueKeys

`func (o *PullRequest) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *PullRequest) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *PullRequest) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *PullRequest) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *PullRequest) GetAssociations() []IssueIdOrKeysAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *PullRequest) GetAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *PullRequest) SetAssociations(v []IssueIdOrKeysAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *PullRequest) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *PullRequest) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *PullRequest) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *PullRequest) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.


### GetStatus

`func (o *PullRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PullRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PullRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *PullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAuthor

`func (o *PullRequest) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PullRequest) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PullRequest) SetAuthor(v Author)`

SetAuthor sets Author field to given value.


### GetCommentCount

`func (o *PullRequest) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *PullRequest) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *PullRequest) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetSourceBranch

`func (o *PullRequest) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *PullRequest) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *PullRequest) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.


### GetSourceBranchUrl

`func (o *PullRequest) GetSourceBranchUrl() string`

GetSourceBranchUrl returns the SourceBranchUrl field if non-nil, zero value otherwise.

### GetSourceBranchUrlOk

`func (o *PullRequest) GetSourceBranchUrlOk() (*string, bool)`

GetSourceBranchUrlOk returns a tuple with the SourceBranchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranchUrl

`func (o *PullRequest) SetSourceBranchUrl(v string)`

SetSourceBranchUrl sets SourceBranchUrl field to given value.

### HasSourceBranchUrl

`func (o *PullRequest) HasSourceBranchUrl() bool`

HasSourceBranchUrl returns a boolean if a field has been set.

### GetLastUpdate

`func (o *PullRequest) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *PullRequest) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *PullRequest) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.


### GetDestinationBranch

`func (o *PullRequest) GetDestinationBranch() string`

GetDestinationBranch returns the DestinationBranch field if non-nil, zero value otherwise.

### GetDestinationBranchOk

`func (o *PullRequest) GetDestinationBranchOk() (*string, bool)`

GetDestinationBranchOk returns a tuple with the DestinationBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationBranch

`func (o *PullRequest) SetDestinationBranch(v string)`

SetDestinationBranch sets DestinationBranch field to given value.

### HasDestinationBranch

`func (o *PullRequest) HasDestinationBranch() bool`

HasDestinationBranch returns a boolean if a field has been set.

### GetDestinationBranchUrl

`func (o *PullRequest) GetDestinationBranchUrl() string`

GetDestinationBranchUrl returns the DestinationBranchUrl field if non-nil, zero value otherwise.

### GetDestinationBranchUrlOk

`func (o *PullRequest) GetDestinationBranchUrlOk() (*string, bool)`

GetDestinationBranchUrlOk returns a tuple with the DestinationBranchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationBranchUrl

`func (o *PullRequest) SetDestinationBranchUrl(v string)`

SetDestinationBranchUrl sets DestinationBranchUrl field to given value.

### HasDestinationBranchUrl

`func (o *PullRequest) HasDestinationBranchUrl() bool`

HasDestinationBranchUrl returns a boolean if a field has been set.

### GetReviewers

`func (o *PullRequest) GetReviewers() []Reviewer`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *PullRequest) GetReviewersOk() (*[]Reviewer, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *PullRequest) SetReviewers(v []Reviewer)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *PullRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetUrl

`func (o *PullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayId

`func (o *PullRequest) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *PullRequest) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *PullRequest) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


