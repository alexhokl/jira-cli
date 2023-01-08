# TaskProgressBeanObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the task. | [optional] 
**ElapsedRuntime** | **int64** | The execution time of the task, in milliseconds. | 
**Finished** | Pointer to **int64** | A timestamp recording when the task was finished. | [optional] 
**Id** | **string** | The ID of the task. | 
**LastUpdate** | **int64** | A timestamp recording when the task progress was last updated. | 
**Message** | Pointer to **string** | Information about the progress of the task. | [optional] 
**Progress** | **int64** | The progress of the task, as a percentage complete. | 
**Result** | Pointer to **interface{}** | The result of the task execution. | [optional] 
**Self** | **string** | The URL of the task. | 
**Started** | Pointer to **int64** | A timestamp recording when the task was started. | [optional] 
**Status** | **string** | The status of the task. | 
**Submitted** | **int64** | A timestamp recording when the task was submitted. | 
**SubmittedBy** | **int64** | The ID of the user who submitted the task. | 

## Methods

### NewTaskProgressBeanObject

`func NewTaskProgressBeanObject(elapsedRuntime int64, id string, lastUpdate int64, progress int64, self string, status string, submitted int64, submittedBy int64, ) *TaskProgressBeanObject`

NewTaskProgressBeanObject instantiates a new TaskProgressBeanObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskProgressBeanObjectWithDefaults

`func NewTaskProgressBeanObjectWithDefaults() *TaskProgressBeanObject`

NewTaskProgressBeanObjectWithDefaults instantiates a new TaskProgressBeanObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TaskProgressBeanObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskProgressBeanObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskProgressBeanObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskProgressBeanObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElapsedRuntime

`func (o *TaskProgressBeanObject) GetElapsedRuntime() int64`

GetElapsedRuntime returns the ElapsedRuntime field if non-nil, zero value otherwise.

### GetElapsedRuntimeOk

`func (o *TaskProgressBeanObject) GetElapsedRuntimeOk() (*int64, bool)`

GetElapsedRuntimeOk returns a tuple with the ElapsedRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedRuntime

`func (o *TaskProgressBeanObject) SetElapsedRuntime(v int64)`

SetElapsedRuntime sets ElapsedRuntime field to given value.


### GetFinished

`func (o *TaskProgressBeanObject) GetFinished() int64`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *TaskProgressBeanObject) GetFinishedOk() (*int64, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *TaskProgressBeanObject) SetFinished(v int64)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *TaskProgressBeanObject) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *TaskProgressBeanObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskProgressBeanObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskProgressBeanObject) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdate

`func (o *TaskProgressBeanObject) GetLastUpdate() int64`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *TaskProgressBeanObject) GetLastUpdateOk() (*int64, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *TaskProgressBeanObject) SetLastUpdate(v int64)`

SetLastUpdate sets LastUpdate field to given value.


### GetMessage

`func (o *TaskProgressBeanObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskProgressBeanObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskProgressBeanObject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskProgressBeanObject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProgress

`func (o *TaskProgressBeanObject) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TaskProgressBeanObject) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TaskProgressBeanObject) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetResult

`func (o *TaskProgressBeanObject) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaskProgressBeanObject) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaskProgressBeanObject) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TaskProgressBeanObject) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *TaskProgressBeanObject) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TaskProgressBeanObject) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetSelf

`func (o *TaskProgressBeanObject) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *TaskProgressBeanObject) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *TaskProgressBeanObject) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetStarted

`func (o *TaskProgressBeanObject) GetStarted() int64`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *TaskProgressBeanObject) GetStartedOk() (*int64, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *TaskProgressBeanObject) SetStarted(v int64)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *TaskProgressBeanObject) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *TaskProgressBeanObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskProgressBeanObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskProgressBeanObject) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmitted

`func (o *TaskProgressBeanObject) GetSubmitted() int64`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *TaskProgressBeanObject) GetSubmittedOk() (*int64, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *TaskProgressBeanObject) SetSubmitted(v int64)`

SetSubmitted sets Submitted field to given value.


### GetSubmittedBy

`func (o *TaskProgressBeanObject) GetSubmittedBy() int64`

GetSubmittedBy returns the SubmittedBy field if non-nil, zero value otherwise.

### GetSubmittedByOk

`func (o *TaskProgressBeanObject) GetSubmittedByOk() (*int64, bool)`

GetSubmittedByOk returns a tuple with the SubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedBy

`func (o *TaskProgressBeanObject) SetSubmittedBy(v int64)`

SetSubmittedBy sets SubmittedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


