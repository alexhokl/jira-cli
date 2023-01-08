# WorkflowId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draft** | **bool** | Whether the workflow is in the draft state. | 
**Name** | **string** | The name of the workflow. | 

## Methods

### NewWorkflowId

`func NewWorkflowId(draft bool, name string, ) *WorkflowId`

NewWorkflowId instantiates a new WorkflowId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowIdWithDefaults

`func NewWorkflowIdWithDefaults() *WorkflowId`

NewWorkflowIdWithDefaults instantiates a new WorkflowId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraft

`func (o *WorkflowId) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *WorkflowId) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *WorkflowId) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetName

`func (o *WorkflowId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowId) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


