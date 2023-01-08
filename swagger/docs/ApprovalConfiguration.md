# ApprovalConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **string** | Whether the approval configuration is active. | 
**ConditionType** | **string** | How the required approval count is calculated. It may be configured to require a specific number of approvals, or approval by a percentage of approvers. If the approvers source field is Approver groups, you can configure how many approvals per group are required for the request to be approved. The number will be the same across all groups. | 
**ConditionValue** | **string** | The number or percentage of approvals required for a request to be approved. If &#x60;conditionType&#x60; is &#x60;number&#x60;, the value must be 20 or less. If &#x60;conditionType&#x60; is &#x60;percent&#x60;, the value must be 100 or less. | 
**Exclude** | Pointer to **[]string** | A list of roles that should be excluded as possible approvers. | [optional] 
**FieldId** | **string** | The custom field ID of the \&quot;Approvers\&quot; or \&quot;Approver Groups\&quot; field. | 
**PrePopulatedFieldId** | Pointer to **NullableString** | The custom field ID of the field used to pre-populate the Approver field. Only supports the \&quot;Affected Services\&quot; field. | [optional] 
**TransitionApproved** | **string** | The numeric ID of the transition to be executed if the request is approved. | 
**TransitionRejected** | **string** | The numeric ID of the transition to be executed if the request is declined. | 

## Methods

### NewApprovalConfiguration

`func NewApprovalConfiguration(active string, conditionType string, conditionValue string, fieldId string, transitionApproved string, transitionRejected string, ) *ApprovalConfiguration`

NewApprovalConfiguration instantiates a new ApprovalConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConfigurationWithDefaults

`func NewApprovalConfigurationWithDefaults() *ApprovalConfiguration`

NewApprovalConfigurationWithDefaults instantiates a new ApprovalConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApprovalConfiguration) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApprovalConfiguration) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApprovalConfiguration) SetActive(v string)`

SetActive sets Active field to given value.


### GetConditionType

`func (o *ApprovalConfiguration) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ApprovalConfiguration) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ApprovalConfiguration) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetConditionValue

`func (o *ApprovalConfiguration) GetConditionValue() string`

GetConditionValue returns the ConditionValue field if non-nil, zero value otherwise.

### GetConditionValueOk

`func (o *ApprovalConfiguration) GetConditionValueOk() (*string, bool)`

GetConditionValueOk returns a tuple with the ConditionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionValue

`func (o *ApprovalConfiguration) SetConditionValue(v string)`

SetConditionValue sets ConditionValue field to given value.


### GetExclude

`func (o *ApprovalConfiguration) GetExclude() []*string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ApprovalConfiguration) GetExcludeOk() (*[]*string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ApprovalConfiguration) SetExclude(v []*string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ApprovalConfiguration) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### SetExcludeNil

`func (o *ApprovalConfiguration) SetExcludeNil(b bool)`

 SetExcludeNil sets the value for Exclude to be an explicit nil

### UnsetExclude
`func (o *ApprovalConfiguration) UnsetExclude()`

UnsetExclude ensures that no value is present for Exclude, not even an explicit nil
### GetFieldId

`func (o *ApprovalConfiguration) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ApprovalConfiguration) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ApprovalConfiguration) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetPrePopulatedFieldId

`func (o *ApprovalConfiguration) GetPrePopulatedFieldId() string`

GetPrePopulatedFieldId returns the PrePopulatedFieldId field if non-nil, zero value otherwise.

### GetPrePopulatedFieldIdOk

`func (o *ApprovalConfiguration) GetPrePopulatedFieldIdOk() (*string, bool)`

GetPrePopulatedFieldIdOk returns a tuple with the PrePopulatedFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePopulatedFieldId

`func (o *ApprovalConfiguration) SetPrePopulatedFieldId(v string)`

SetPrePopulatedFieldId sets PrePopulatedFieldId field to given value.

### HasPrePopulatedFieldId

`func (o *ApprovalConfiguration) HasPrePopulatedFieldId() bool`

HasPrePopulatedFieldId returns a boolean if a field has been set.

### SetPrePopulatedFieldIdNil

`func (o *ApprovalConfiguration) SetPrePopulatedFieldIdNil(b bool)`

 SetPrePopulatedFieldIdNil sets the value for PrePopulatedFieldId to be an explicit nil

### UnsetPrePopulatedFieldId
`func (o *ApprovalConfiguration) UnsetPrePopulatedFieldId()`

UnsetPrePopulatedFieldId ensures that no value is present for PrePopulatedFieldId, not even an explicit nil
### GetTransitionApproved

`func (o *ApprovalConfiguration) GetTransitionApproved() string`

GetTransitionApproved returns the TransitionApproved field if non-nil, zero value otherwise.

### GetTransitionApprovedOk

`func (o *ApprovalConfiguration) GetTransitionApprovedOk() (*string, bool)`

GetTransitionApprovedOk returns a tuple with the TransitionApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionApproved

`func (o *ApprovalConfiguration) SetTransitionApproved(v string)`

SetTransitionApproved sets TransitionApproved field to given value.


### GetTransitionRejected

`func (o *ApprovalConfiguration) GetTransitionRejected() string`

GetTransitionRejected returns the TransitionRejected field if non-nil, zero value otherwise.

### GetTransitionRejectedOk

`func (o *ApprovalConfiguration) GetTransitionRejectedOk() (*string, bool)`

GetTransitionRejectedOk returns a tuple with the TransitionRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionRejected

`func (o *ApprovalConfiguration) SetTransitionRejected(v string)`

SetTransitionRejected sets TransitionRejected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


