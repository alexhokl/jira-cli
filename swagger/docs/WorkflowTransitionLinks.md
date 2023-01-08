# WorkflowTransitionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | Pointer to **NullableInt32** | The port that the transition starts from. | [optional] 
**FromStatusReference** | Pointer to **NullableString** | The status that the transition starts from. | [optional] 
**ToPort** | Pointer to **NullableInt32** | The port that the transition goes to. | [optional] 

## Methods

### NewWorkflowTransitionLinks

`func NewWorkflowTransitionLinks() *WorkflowTransitionLinks`

NewWorkflowTransitionLinks instantiates a new WorkflowTransitionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionLinksWithDefaults

`func NewWorkflowTransitionLinksWithDefaults() *WorkflowTransitionLinks`

NewWorkflowTransitionLinksWithDefaults instantiates a new WorkflowTransitionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *WorkflowTransitionLinks) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *WorkflowTransitionLinks) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *WorkflowTransitionLinks) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *WorkflowTransitionLinks) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### SetFromPortNil

`func (o *WorkflowTransitionLinks) SetFromPortNil(b bool)`

 SetFromPortNil sets the value for FromPort to be an explicit nil

### UnsetFromPort
`func (o *WorkflowTransitionLinks) UnsetFromPort()`

UnsetFromPort ensures that no value is present for FromPort, not even an explicit nil
### GetFromStatusReference

`func (o *WorkflowTransitionLinks) GetFromStatusReference() string`

GetFromStatusReference returns the FromStatusReference field if non-nil, zero value otherwise.

### GetFromStatusReferenceOk

`func (o *WorkflowTransitionLinks) GetFromStatusReferenceOk() (*string, bool)`

GetFromStatusReferenceOk returns a tuple with the FromStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStatusReference

`func (o *WorkflowTransitionLinks) SetFromStatusReference(v string)`

SetFromStatusReference sets FromStatusReference field to given value.

### HasFromStatusReference

`func (o *WorkflowTransitionLinks) HasFromStatusReference() bool`

HasFromStatusReference returns a boolean if a field has been set.

### SetFromStatusReferenceNil

`func (o *WorkflowTransitionLinks) SetFromStatusReferenceNil(b bool)`

 SetFromStatusReferenceNil sets the value for FromStatusReference to be an explicit nil

### UnsetFromStatusReference
`func (o *WorkflowTransitionLinks) UnsetFromStatusReference()`

UnsetFromStatusReference ensures that no value is present for FromStatusReference, not even an explicit nil
### GetToPort

`func (o *WorkflowTransitionLinks) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *WorkflowTransitionLinks) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *WorkflowTransitionLinks) SetToPort(v int32)`

SetToPort sets ToPort field to given value.

### HasToPort

`func (o *WorkflowTransitionLinks) HasToPort() bool`

HasToPort returns a boolean if a field has been set.

### SetToPortNil

`func (o *WorkflowTransitionLinks) SetToPortNil(b bool)`

 SetToPortNil sets the value for ToPort to be an explicit nil

### UnsetToPort
`func (o *WorkflowTransitionLinks) UnsetToPort()`

UnsetToPort ensures that no value is present for ToPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


