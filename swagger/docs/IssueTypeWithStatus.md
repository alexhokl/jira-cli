# IssueTypeWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the issue type. | [readonly] 
**Name** | **string** | The name of the issue type. | [readonly] 
**Self** | **string** | The URL of the issue type&#39;s status details. | [readonly] 
**Statuses** | [**[]StatusDetails**](StatusDetails.md) | List of status details for the issue type. | [readonly] 
**Subtask** | **bool** | Whether this issue type represents subtasks. | [readonly] 

## Methods

### NewIssueTypeWithStatus

`func NewIssueTypeWithStatus(id string, name string, self string, statuses []StatusDetails, subtask bool, ) *IssueTypeWithStatus`

NewIssueTypeWithStatus instantiates a new IssueTypeWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeWithStatusWithDefaults

`func NewIssueTypeWithStatusWithDefaults() *IssueTypeWithStatus`

NewIssueTypeWithStatusWithDefaults instantiates a new IssueTypeWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueTypeWithStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeWithStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeWithStatus) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IssueTypeWithStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeWithStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeWithStatus) SetName(v string)`

SetName sets Name field to given value.


### GetSelf

`func (o *IssueTypeWithStatus) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueTypeWithStatus) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueTypeWithStatus) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetStatuses

`func (o *IssueTypeWithStatus) GetStatuses() []StatusDetails`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *IssueTypeWithStatus) GetStatusesOk() (*[]StatusDetails, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *IssueTypeWithStatus) SetStatuses(v []StatusDetails)`

SetStatuses sets Statuses field to given value.


### GetSubtask

`func (o *IssueTypeWithStatus) GetSubtask() bool`

GetSubtask returns the Subtask field if non-nil, zero value otherwise.

### GetSubtaskOk

`func (o *IssueTypeWithStatus) GetSubtaskOk() (*bool, bool)`

GetSubtaskOk returns a tuple with the Subtask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtask

`func (o *IssueTypeWithStatus) SetSubtask(v bool)`

SetSubtask sets Subtask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


