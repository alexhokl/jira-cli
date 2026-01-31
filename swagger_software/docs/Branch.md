# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this entity. Will be used for cross entity linking. Must be unique by entity type within a repository, i.e., only one commit can have ID &#39;X&#39; in repository &#39;Y&#39;. But adding, e.g., a branch with ID &#39;X&#39; to repository &#39;Y&#39; is acceptable. Only alphanumeric characters, and &#39;~.-_&#39;, are allowed. Max length is 1024 characters. | 
**IssueKeys** | Pointer to **[]string** | List of issues keys that this entity is associated with. They must be valid Jira issue keys. | [optional] 
**Associations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | The Jira issue keys or IDs to associate the branch with. | [optional] 
**UpdateSequenceId** | **int64** | An ID used to apply an ordering to updates for this entity in the case of out-of-order receipt of update requests. This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the provider system, but other alternatives are valid (e.g. a provider could store a counter against each entity and increment that on each update to Jira). Updates for an entity that are received with an updateSqeuenceId lower than what is currently stored will be ignored. | 
**Name** | **string** | The name of the branch. Max length is 512 characters. | 
**LastCommit** | [**Commit1**](Commit1.md) |  | 
**CreatePullRequestUrl** | Pointer to **string** | The URL of the page for creating a pull request from this branch. Max length is 2000 characters. | [optional] 
**Url** | **string** | The URL of the branch. Max length is 2000 characters. | 

## Methods

### NewBranch

`func NewBranch(id string, updateSequenceId int64, name string, lastCommit Commit1, url string, ) *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Branch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v string)`

SetId sets Id field to given value.


### GetIssueKeys

`func (o *Branch) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *Branch) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *Branch) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *Branch) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *Branch) GetAssociations() []IssueIdOrKeysAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Branch) GetAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Branch) SetAssociations(v []IssueIdOrKeysAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Branch) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *Branch) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *Branch) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *Branch) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.


### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.


### GetLastCommit

`func (o *Branch) GetLastCommit() Commit1`

GetLastCommit returns the LastCommit field if non-nil, zero value otherwise.

### GetLastCommitOk

`func (o *Branch) GetLastCommitOk() (*Commit1, bool)`

GetLastCommitOk returns a tuple with the LastCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommit

`func (o *Branch) SetLastCommit(v Commit1)`

SetLastCommit sets LastCommit field to given value.


### GetCreatePullRequestUrl

`func (o *Branch) GetCreatePullRequestUrl() string`

GetCreatePullRequestUrl returns the CreatePullRequestUrl field if non-nil, zero value otherwise.

### GetCreatePullRequestUrlOk

`func (o *Branch) GetCreatePullRequestUrlOk() (*string, bool)`

GetCreatePullRequestUrlOk returns a tuple with the CreatePullRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePullRequestUrl

`func (o *Branch) SetCreatePullRequestUrl(v string)`

SetCreatePullRequestUrl sets CreatePullRequestUrl field to given value.

### HasCreatePullRequestUrl

`func (o *Branch) HasCreatePullRequestUrl() bool`

HasCreatePullRequestUrl returns a boolean if a field has been set.

### GetUrl

`func (o *Branch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Branch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Branch) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


