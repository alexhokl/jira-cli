# BuildReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**BuildCommitReference**](BuildCommitReference.md) |  | [optional] 
**Ref** | Pointer to [**BuildRefReference**](BuildRefReference.md) |  | [optional] 

## Methods

### NewBuildReferences

`func NewBuildReferences() *BuildReferences`

NewBuildReferences instantiates a new BuildReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildReferencesWithDefaults

`func NewBuildReferencesWithDefaults() *BuildReferences`

NewBuildReferencesWithDefaults instantiates a new BuildReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *BuildReferences) GetCommit() BuildCommitReference`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BuildReferences) GetCommitOk() (*BuildCommitReference, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BuildReferences) SetCommit(v BuildCommitReference)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *BuildReferences) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetRef

`func (o *BuildReferences) GetRef() BuildRefReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *BuildReferences) GetRefOk() (*BuildRefReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *BuildReferences) SetRef(v BuildRefReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *BuildReferences) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


