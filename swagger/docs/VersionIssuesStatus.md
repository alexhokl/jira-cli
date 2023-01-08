# VersionIssuesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **int64** | Count of issues with status *done*. | [optional] [readonly] 
**InProgress** | Pointer to **int64** | Count of issues with status *in progress*. | [optional] [readonly] 
**ToDo** | Pointer to **int64** | Count of issues with status *to do*. | [optional] [readonly] 
**Unmapped** | Pointer to **int64** | Count of issues with a status other than *to do*, *in progress*, and *done*. | [optional] [readonly] 

## Methods

### NewVersionIssuesStatus

`func NewVersionIssuesStatus() *VersionIssuesStatus`

NewVersionIssuesStatus instantiates a new VersionIssuesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionIssuesStatusWithDefaults

`func NewVersionIssuesStatusWithDefaults() *VersionIssuesStatus`

NewVersionIssuesStatusWithDefaults instantiates a new VersionIssuesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *VersionIssuesStatus) GetDone() int64`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *VersionIssuesStatus) GetDoneOk() (*int64, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *VersionIssuesStatus) SetDone(v int64)`

SetDone sets Done field to given value.

### HasDone

`func (o *VersionIssuesStatus) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetInProgress

`func (o *VersionIssuesStatus) GetInProgress() int64`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *VersionIssuesStatus) GetInProgressOk() (*int64, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *VersionIssuesStatus) SetInProgress(v int64)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *VersionIssuesStatus) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetToDo

`func (o *VersionIssuesStatus) GetToDo() int64`

GetToDo returns the ToDo field if non-nil, zero value otherwise.

### GetToDoOk

`func (o *VersionIssuesStatus) GetToDoOk() (*int64, bool)`

GetToDoOk returns a tuple with the ToDo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDo

`func (o *VersionIssuesStatus) SetToDo(v int64)`

SetToDo sets ToDo field to given value.

### HasToDo

`func (o *VersionIssuesStatus) HasToDo() bool`

HasToDo returns a boolean if a field has been set.

### GetUnmapped

`func (o *VersionIssuesStatus) GetUnmapped() int64`

GetUnmapped returns the Unmapped field if non-nil, zero value otherwise.

### GetUnmappedOk

`func (o *VersionIssuesStatus) GetUnmappedOk() (*int64, bool)`

GetUnmappedOk returns a tuple with the Unmapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapped

`func (o *VersionIssuesStatus) SetUnmapped(v int64)`

SetUnmapped sets Unmapped field to given value.

### HasUnmapped

`func (o *VersionIssuesStatus) HasUnmapped() bool`

HasUnmapped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


