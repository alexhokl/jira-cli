# WorkflowStatusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Id** | Pointer to **string** | The ID of the status. | [optional] 
**Name** | **string** | The name of the status. | 
**StatusCategory** | **string** | The category of the status. | 
**StatusReference** | **string** | The reference of the status. | 

## Methods

### NewWorkflowStatusUpdate

`func NewWorkflowStatusUpdate(name string, statusCategory string, statusReference string, ) *WorkflowStatusUpdate`

NewWorkflowStatusUpdate instantiates a new WorkflowStatusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusUpdateWithDefaults

`func NewWorkflowStatusUpdateWithDefaults() *WorkflowStatusUpdate`

NewWorkflowStatusUpdateWithDefaults instantiates a new WorkflowStatusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowStatusUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowStatusUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowStatusUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowStatusUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkflowStatusUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatusUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatusUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowStatusUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowStatusUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStatusUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStatusUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetStatusCategory

`func (o *WorkflowStatusUpdate) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *WorkflowStatusUpdate) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *WorkflowStatusUpdate) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.


### GetStatusReference

`func (o *WorkflowStatusUpdate) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *WorkflowStatusUpdate) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *WorkflowStatusUpdate) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


