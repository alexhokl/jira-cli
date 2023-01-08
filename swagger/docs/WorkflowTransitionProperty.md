# WorkflowTransitionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the transition property. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the transition property. Also known as the name of the transition property. | [optional] [readonly] 
**Value** | **string** | The value of the transition property. | 

## Methods

### NewWorkflowTransitionProperty

`func NewWorkflowTransitionProperty(value string, ) *WorkflowTransitionProperty`

NewWorkflowTransitionProperty instantiates a new WorkflowTransitionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionPropertyWithDefaults

`func NewWorkflowTransitionPropertyWithDefaults() *WorkflowTransitionProperty`

NewWorkflowTransitionPropertyWithDefaults instantiates a new WorkflowTransitionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTransitionProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTransitionProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTransitionProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTransitionProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *WorkflowTransitionProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowTransitionProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowTransitionProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowTransitionProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowTransitionProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowTransitionProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowTransitionProperty) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


