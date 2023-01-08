# TaskProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the task. | [optional] 
**ElapsedRuntime** | **int64** | The execution time of the task, in milliseconds. | 
**Finished** | Pointer to **time.Time** | A timestamp recording when the task was finished. | [optional] 
**Id** | **string** | The ID of the task. | 
**LastUpdate** | **time.Time** | A timestamp recording when the task progress was last updated. | 
**Message** | Pointer to **string** | Information about the progress of the task. | [optional] 
**Progress** | **int64** | The progress of the task, as a percentage complete. | 
**Result** | Pointer to **interface{}** | The result of the task execution. | [optional] 
**Self** | **string** | The URL of the task. | 
**Started** | Pointer to **time.Time** | A timestamp recording when the task was started. | [optional] 
**Status** | **string** | The status of the task. | 
**Submitted** | Pointer to **time.Time** | A timestamp recording when the task was submitted. | [optional] 
**SubmittedBy** | **int64** | The ID of the user who submitted the task. | 

## Methods

### NewTaskProgress

`func NewTaskProgress(elapsedRuntime int64, id string, lastUpdate time.Time, progress int64, self string, status string, submittedBy int64, ) *TaskProgress`

NewTaskProgress instantiates a new TaskProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskProgressWithDefaults

`func NewTaskProgressWithDefaults() *TaskProgress`

NewTaskProgressWithDefaults instantiates a new TaskProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TaskProgress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskProgress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskProgress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskProgress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElapsedRuntime

`func (o *TaskProgress) GetElapsedRuntime() int64`

GetElapsedRuntime returns the ElapsedRuntime field if non-nil, zero value otherwise.

### GetElapsedRuntimeOk

`func (o *TaskProgress) GetElapsedRuntimeOk() (*int64, bool)`

GetElapsedRuntimeOk returns a tuple with the ElapsedRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedRuntime

`func (o *TaskProgress) SetElapsedRuntime(v int64)`

SetElapsedRuntime sets ElapsedRuntime field to given value.


### GetFinished

`func (o *TaskProgress) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *TaskProgress) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *TaskProgress) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *TaskProgress) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *TaskProgress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskProgress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskProgress) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdate

`func (o *TaskProgress) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *TaskProgress) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *TaskProgress) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.


### GetMessage

`func (o *TaskProgress) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskProgress) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskProgress) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskProgress) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProgress

`func (o *TaskProgress) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TaskProgress) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TaskProgress) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetResult

`func (o *TaskProgress) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaskProgress) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaskProgress) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TaskProgress) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TaskProgress) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TaskProgress) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetSelf

`func (o *TaskProgress) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *TaskProgress) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *TaskProgress) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetStarted

`func (o *TaskProgress) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *TaskProgress) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *TaskProgress) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *TaskProgress) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *TaskProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskProgress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmitted

`func (o *TaskProgress) GetSubmitted() time.Time`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *TaskProgress) GetSubmittedOk() (*time.Time, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *TaskProgress) SetSubmitted(v time.Time)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *TaskProgress) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetSubmittedBy

`func (o *TaskProgress) GetSubmittedBy() int64`

GetSubmittedBy returns the SubmittedBy field if non-nil, zero value otherwise.

### GetSubmittedByOk

`func (o *TaskProgress) GetSubmittedByOk() (*int64, bool)`

GetSubmittedByOk returns a tuple with the SubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedBy

`func (o *TaskProgress) SetSubmittedBy(v int64)`

SetSubmittedBy sets SubmittedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


