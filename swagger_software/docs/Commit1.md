# Commit1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier or hash of the commit. Will be used for cross entity linking. Must be unique for all commits within a repository, i.e., only one commit can have ID &#39;X&#39; in repository &#39;Y&#39;. But adding, e.g., a branch with ID &#39;X&#39; to repository &#39;Y&#39; is acceptable. Only alphanumeric characters, and &#39;~.-_&#39;, are allowed. Max length is 1024 characters | 
**IssueKeys** | **[]string** | List of issues keys that this entity is associated with. They must be valid Jira issue keys. | 
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

### NewCommit1

`func NewCommit1(id string, issueKeys []string, updateSequenceId int64, message string, author Author, fileCount int32, url string, authorTimestamp string, displayId string, ) *Commit1`

NewCommit1 instantiates a new Commit1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommit1WithDefaults

`func NewCommit1WithDefaults() *Commit1`

NewCommit1WithDefaults instantiates a new Commit1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commit1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit1) SetId(v string)`

SetId sets Id field to given value.


### GetIssueKeys

`func (o *Commit1) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *Commit1) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *Commit1) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.


### GetUpdateSequenceId

`func (o *Commit1) GetUpdateSequenceId() int64`

GetUpdateSequenceId returns the UpdateSequenceId field if non-nil, zero value otherwise.

### GetUpdateSequenceIdOk

`func (o *Commit1) GetUpdateSequenceIdOk() (*int64, bool)`

GetUpdateSequenceIdOk returns a tuple with the UpdateSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceId

`func (o *Commit1) SetUpdateSequenceId(v int64)`

SetUpdateSequenceId sets UpdateSequenceId field to given value.


### GetHash

`func (o *Commit1) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Commit1) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Commit1) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Commit1) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetFlags

`func (o *Commit1) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Commit1) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Commit1) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Commit1) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetMessage

`func (o *Commit1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthor

`func (o *Commit1) GetAuthor() Author`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit1) GetAuthorOk() (*Author, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit1) SetAuthor(v Author)`

SetAuthor sets Author field to given value.


### GetFileCount

`func (o *Commit1) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *Commit1) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *Commit1) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.


### GetUrl

`func (o *Commit1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commit1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commit1) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFiles

`func (o *Commit1) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Commit1) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Commit1) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Commit1) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetAuthorTimestamp

`func (o *Commit1) GetAuthorTimestamp() string`

GetAuthorTimestamp returns the AuthorTimestamp field if non-nil, zero value otherwise.

### GetAuthorTimestampOk

`func (o *Commit1) GetAuthorTimestampOk() (*string, bool)`

GetAuthorTimestampOk returns a tuple with the AuthorTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTimestamp

`func (o *Commit1) SetAuthorTimestamp(v string)`

SetAuthorTimestamp sets AuthorTimestamp field to given value.


### GetDisplayId

`func (o *Commit1) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Commit1) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Commit1) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


