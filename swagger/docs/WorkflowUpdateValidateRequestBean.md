# WorkflowUpdateValidateRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**WorkflowUpdateRequest**](WorkflowUpdateRequest.md) |  | 
**ValidationOptions** | Pointer to [**ValidationOptionsForUpdate**](ValidationOptionsForUpdate.md) |  | [optional] 

## Methods

### NewWorkflowUpdateValidateRequestBean

`func NewWorkflowUpdateValidateRequestBean(payload WorkflowUpdateRequest, ) *WorkflowUpdateValidateRequestBean`

NewWorkflowUpdateValidateRequestBean instantiates a new WorkflowUpdateValidateRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUpdateValidateRequestBeanWithDefaults

`func NewWorkflowUpdateValidateRequestBeanWithDefaults() *WorkflowUpdateValidateRequestBean`

NewWorkflowUpdateValidateRequestBeanWithDefaults instantiates a new WorkflowUpdateValidateRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *WorkflowUpdateValidateRequestBean) GetPayload() WorkflowUpdateRequest`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WorkflowUpdateValidateRequestBean) GetPayloadOk() (*WorkflowUpdateRequest, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WorkflowUpdateValidateRequestBean) SetPayload(v WorkflowUpdateRequest)`

SetPayload sets Payload field to given value.


### GetValidationOptions

`func (o *WorkflowUpdateValidateRequestBean) GetValidationOptions() ValidationOptionsForUpdate`

GetValidationOptions returns the ValidationOptions field if non-nil, zero value otherwise.

### GetValidationOptionsOk

`func (o *WorkflowUpdateValidateRequestBean) GetValidationOptionsOk() (*ValidationOptionsForUpdate, bool)`

GetValidationOptionsOk returns a tuple with the ValidationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOptions

`func (o *WorkflowUpdateValidateRequestBean) SetValidationOptions(v ValidationOptionsForUpdate)`

SetValidationOptions sets ValidationOptions field to given value.

### HasValidationOptions

`func (o *WorkflowUpdateValidateRequestBean) HasValidationOptions() bool`

HasValidationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


