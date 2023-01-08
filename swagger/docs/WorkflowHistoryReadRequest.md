# WorkflowHistoryReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int64** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkflowHistoryReadRequest

`func NewWorkflowHistoryReadRequest() *WorkflowHistoryReadRequest`

NewWorkflowHistoryReadRequest instantiates a new WorkflowHistoryReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowHistoryReadRequestWithDefaults

`func NewWorkflowHistoryReadRequestWithDefaults() *WorkflowHistoryReadRequest`

NewWorkflowHistoryReadRequestWithDefaults instantiates a new WorkflowHistoryReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *WorkflowHistoryReadRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowHistoryReadRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowHistoryReadRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowHistoryReadRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowHistoryReadRequest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowHistoryReadRequest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowHistoryReadRequest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowHistoryReadRequest) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


