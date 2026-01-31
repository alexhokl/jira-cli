# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier or hash of the commit. Will be used for cross entity linking. Must be unique for all commits within a repository, i.e., only one commit can have ID &#39;X&#39; in repository &#39;Y&#39;. But adding, e.g., a branch with ID &#39;X&#39; to repository &#39;Y&#39; is acceptable. Only alphanumeric characters, and &#39;~.-_&#39;, are allowed. Max length is 1024 characters | 
**IssueKeys** | Pointer to **[]string** | List of issues keys that this entity is associated with. They must be valid Jira issue keys. | [optional] 
**Associations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | The Jira issue keys or IDs to associate the commit with. | [optional] 
**UpdateSequenceId** | **int64** | An ID used to apply an ordering to updates for this entity in the case of out-of-order receipt of update requests. This can be any monotonically increasing number. A suggested implementation is to use epoch millis from the provider system, but other alternatives are valid (e.g. a provider could store a counter against each entity and increment that on each update to Jira). Updates for an entity that are received with an updateSqeuenceId lower than what is currently stored will be ignored. | 
**Hash** | Pointer to **string** | Deprecated. Use the id field instead. | [optional] 
**Flags** | Pointer to **[]string** | The set of flags for this commit | [optional] 
**Message** | **string** | The commit message. Max length is 1024 characters. If anything longer is supplied, it will be truncated down to 1024 characters. | 
**Author** | [**Author**](Author.md) |  | 
**FileCount** | **int32** | The total number of files added, removed, or modified by this commit | 
**Url** | **string** | The URL of this commit. Max length is 2000 characters. | 
**Files** | Pointer to [**[]File**](File.md) | List of file changes. Max number of files is 10. Currently, only the first 5 files are shown (sorted by path) in the UI. This UI behavior may change without notice. | [optional] 
**AuthorTimestamp** | **string** | The author timestamp of this commit. Formatted as a UTC ISO 8601 date time format. | 
**DisplayId** | **string** | Shortened identifier for this commit, used for display. Max length is 255 characters. | 

## Methods

### NewCommit

`func NewCommit(id string, updateSequenceId int64, message string, author Author, fileCount int32, url string, authorTimestamp string, displayId string, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit) SetId(v string)`

SetId sets Id field to given value.


### GetIssueKeys

`func (o *Commit) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *Commit) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *Commit) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *Commit) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *Commit) GetAssociations() []IssueIdOrKeysAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *Commit) GetAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *Commit) SetAssociations(v []IssueIdOrKeysAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *Commit) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetUpdateSequenceId

`func (o *Commit) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *Commit) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *Commit) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.


### GetHash

`func (o *Commit) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Commit) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Commit) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Commit) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetFlags

`func (o *Commit) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Commit) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Commit) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Commit) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMessage

`func (o *Commit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthor

`func (o *Commit) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit) SetAuthor(v Author)`

SetAuthor sets Author field to given value.


### GetFileCount

`func (o *Commit) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *Commit) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *Commit) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.


### GetUrl

`func (o *Commit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFiles

`func (o *Commit) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Commit) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Commit) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Commit) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetAuthorTimestamp

`func (o *Commit) GetAuthorTimestamp() string`

GetAuthorTimestamp returns the AuthorTimestamp field if non-nil, zero value otherwise.

### GetAuthorTimestampOk

`func (o *Commit) GetAuthorTimestampOk() (*string, bool)`

GetAuthorTimestampOk returns a tuple with the AuthorTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTimestamp

`func (o *Commit) SetAuthorTimestamp(v string)`

SetAuthorTimestamp sets AuthorTimestamp field to given value.


### GetDisplayId

`func (o *Commit) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Commit) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Commit) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


