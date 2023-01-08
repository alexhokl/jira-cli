# WorkflowRulesSearchDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidRules** | Pointer to **[]string** | List of workflow rule IDs that do not belong to the workflow or can not be found. | [optional] 
**ValidRules** | Pointer to [**[]WorkflowTransitionRules**](WorkflowTransitionRules.md) | List of valid workflow transition rules. | [optional] 
**WorkflowEntityId** | Pointer to **string** | The workflow ID. | [optional] 

## Methods

### NewWorkflowRulesSearchDetails

`func NewWorkflowRulesSearchDetails() *WorkflowRulesSearchDetails`

NewWorkflowRulesSearchDetails instantiates a new WorkflowRulesSearchDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRulesSearchDetailsWithDefaults

`func NewWorkflowRulesSearchDetailsWithDefaults() *WorkflowRulesSearchDetails`

NewWorkflowRulesSearchDetailsWithDefaults instantiates a new WorkflowRulesSearchDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidRules

`func (o *WorkflowRulesSearchDetails) GetInvalidRules() []string`

GetInvalidRules returns the InvalidRules field if non-nil, zero value otherwise.

### GetInvalidRulesOk

`func (o *WorkflowRulesSearchDetails) GetInvalidRulesOk() (*[]string, bool)`

GetInvalidRulesOk returns a tuple with the InvalidRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRules

`func (o *WorkflowRulesSearchDetails) SetInvalidRules(v []string)`

SetInvalidRules sets InvalidRules field to given value.

### HasInvalidRules

`func (o *WorkflowRulesSearchDetails) HasInvalidRules() bool`

HasInvalidRules returns a boolean if a field has been set.

### GetValidRules

`func (o *WorkflowRulesSearchDetails) GetValidRules() []WorkflowTransitionRules`

GetValidRules returns the ValidRules field if non-nil, zero value otherwise.

### GetValidRulesOk

`func (o *WorkflowRulesSearchDetails) GetValidRulesOk() (*[]WorkflowTransitionRules, bool)`

GetValidRulesOk returns a tuple with the ValidRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRules

`func (o *WorkflowRulesSearchDetails) SetValidRules(v []WorkflowTransitionRules)`

SetValidRules sets ValidRules field to given value.

### HasValidRules

`func (o *WorkflowRulesSearchDetails) HasValidRules() bool`

HasValidRules returns a boolean if a field has been set.

### GetWorkflowEntityId

`func (o *WorkflowRulesSearchDetails) GetWorkflowEntityId() string`

GetWorkflowEntityId returns the WorkflowEntityId field if non-nil, zero value otherwise.

### GetWorkflowEntityIdOk

`func (o *WorkflowRulesSearchDetails) GetWorkflowEntityIdOk() (*string, bool)`

GetWorkflowEntityIdOk returns a tuple with the WorkflowEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEntityId

`func (o *WorkflowRulesSearchDetails) SetWorkflowEntityId(v string)`

SetWorkflowEntityId sets WorkflowEntityId field to given value.

### HasWorkflowEntityId

`func (o *WorkflowRulesSearchDetails) HasWorkflowEntityId() bool`

HasWorkflowEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


