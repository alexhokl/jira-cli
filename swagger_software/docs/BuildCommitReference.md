# BuildCommitReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the commit. E.g. for a Git repository this would be the SHA1 hash.  | 
**RepositoryUri** | **string** | An identifier for the repository containing the commit.  In most cases this should be the URL of the repository in the SCM provider.  For cases where the build was executed against a local repository etc. this should be some identifier that is unique to that repository.  | 

## Methods

### NewBuildCommitReference

`func NewBuildCommitReference(id string, repositoryUri string, ) *BuildCommitReference`

NewBuildCommitReference instantiates a new BuildCommitReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildCommitReferenceWithDefaults

`func NewBuildCommitReferenceWithDefaults() *BuildCommitReference`

NewBuildCommitReferenceWithDefaults instantiates a new BuildCommitReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BuildCommitReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuildCommitReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuildCommitReference) SetId(v string)`

SetId sets Id field to given value.


### GetRepositoryUri

`func (o *BuildCommitReference) GetRepositoryUri() string`

GetRepositoryUri returns the RepositoryUri field if non-nil, zero value otherwise.

### GetRepositoryUriOk

`func (o *BuildCommitReference) GetRepositoryUriOk() (*string, bool)`

GetRepositoryUriOk returns a tuple with the RepositoryUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUri

`func (o *BuildCommitReference) SetRepositoryUri(v string)`

SetRepositoryUri sets RepositoryUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


