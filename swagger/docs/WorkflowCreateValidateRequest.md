# WorkflowCreateValidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**WorkflowCreateRequest**](WorkflowCreateRequest.md) |  | 
**ValidationOptions** | Pointer to [**ValidationOptionsForCreate**](ValidationOptionsForCreate.md) |  | [optional] 

## Methods

### NewWorkflowCreateValidateRequest

`func NewWorkflowCreateValidateRequest(payload WorkflowCreateRequest, ) *WorkflowCreateValidateRequest`

NewWorkflowCreateValidateRequest instantiates a new WorkflowCreateValidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCreateValidateRequestWithDefaults

`func NewWorkflowCreateValidateRequestWithDefaults() *WorkflowCreateValidateRequest`

NewWorkflowCreateValidateRequestWithDefaults instantiates a new WorkflowCreateValidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *WorkflowCreateValidateRequest) GetPayload() WorkflowCreateRequest`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WorkflowCreateValidateRequest) GetPayloadOk() (*WorkflowCreateRequest, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WorkflowCreateValidateRequest) SetPayload(v WorkflowCreateRequest)`

SetPayload sets Payload field to given value.


### GetValidationOptions

`func (o *WorkflowCreateValidateRequest) GetValidationOptions() ValidationOptionsForCreate`

GetValidationOptions returns the ValidationOptions field if non-nil, zero value otherwise.

### GetValidationOptionsOk

`func (o *WorkflowCreateValidateRequest) GetValidationOptionsOk() (*ValidationOptionsForCreate, bool)`

GetValidationOptionsOk returns a tuple with the ValidationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOptions

`func (o *WorkflowCreateValidateRequest) SetValidationOptions(v ValidationOptionsForCreate)`

SetValidationOptions sets ValidationOptions field to given value.

### HasValidationOptions

`func (o *WorkflowCreateValidateRequest) HasValidationOptions() bool`

HasValidationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


