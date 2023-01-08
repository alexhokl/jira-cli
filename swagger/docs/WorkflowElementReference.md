# WorkflowElementReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyKey** | Pointer to **string** | A property key. | [optional] 
**RuleId** | Pointer to **string** | A rule ID. | [optional] 
**StatusMappingReference** | Pointer to [**ProjectAndIssueTypePair**](ProjectAndIssueTypePair.md) |  | [optional] 
**StatusReference** | Pointer to **string** | A status reference. | [optional] 
**TransitionId** | Pointer to **string** | A transition ID. | [optional] 

## Methods

### NewWorkflowElementReference

`func NewWorkflowElementReference() *WorkflowElementReference`

NewWorkflowElementReference instantiates a new WorkflowElementReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowElementReferenceWithDefaults

`func NewWorkflowElementReferenceWithDefaults() *WorkflowElementReference`

NewWorkflowElementReferenceWithDefaults instantiates a new WorkflowElementReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyKey

`func (o *WorkflowElementReference) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *WorkflowElementReference) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *WorkflowElementReference) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.

### HasPropertyKey

`func (o *WorkflowElementReference) HasPropertyKey() bool`

HasPropertyKey returns a boolean if a field has been set.

### GetRuleId

`func (o *WorkflowElementReference) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *WorkflowElementReference) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *WorkflowElementReference) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *WorkflowElementReference) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetStatusMappingReference

`func (o *WorkflowElementReference) GetStatusMappingReference() ProjectAndIssueTypePair`

GetStatusMappingReference returns the StatusMappingReference field if non-nil, zero value otherwise.

### GetStatusMappingReferenceOk

`func (o *WorkflowElementReference) GetStatusMappingReferenceOk() (*ProjectAndIssueTypePair, bool)`

GetStatusMappingReferenceOk returns a tuple with the StatusMappingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMappingReference

`func (o *WorkflowElementReference) SetStatusMappingReference(v ProjectAndIssueTypePair)`

SetStatusMappingReference sets StatusMappingReference field to given value.

### HasStatusMappingReference

`func (o *WorkflowElementReference) HasStatusMappingReference() bool`

HasStatusMappingReference returns a boolean if a field has been set.

### GetStatusReference

`func (o *WorkflowElementReference) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *WorkflowElementReference) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *WorkflowElementReference) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.

### HasStatusReference

`func (o *WorkflowElementReference) HasStatusReference() bool`

HasStatusReference returns a boolean if a field has been set.

### GetTransitionId

`func (o *WorkflowElementReference) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *WorkflowElementReference) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *WorkflowElementReference) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.

### HasTransitionId

`func (o *WorkflowElementReference) HasTransitionId() bool`

HasTransitionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


