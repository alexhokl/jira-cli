# WorkflowTransitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]WorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) | The post-functions of the transition. | [optional] 
**Conditions** | Pointer to [**NullableConditionGroupConfiguration**](ConditionGroupConfiguration.md) |  | [optional] 
**CustomIssueEventId** | Pointer to **NullableString** | The custom event ID of the transition. | [optional] 
**Description** | Pointer to **string** | The description of the transition. | [optional] 
**Id** | Pointer to **string** | The ID of the transition. | [optional] 
**Links** | Pointer to [**[]WorkflowTransitionLinks**](WorkflowTransitionLinks.md) | The statuses the transition can start from, and the mapping of ports between the statuses. | [optional] 
**Name** | Pointer to **string** | The name of the transition. | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the transition. | [optional] 
**ToStatusReference** | Pointer to **string** | The status the transition goes to. | [optional] 
**TransitionScreen** | Pointer to [**NullableWorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) |  | [optional] 
**Triggers** | Pointer to [**[]WorkflowTrigger**](WorkflowTrigger.md) | The triggers of the transition. | [optional] 
**Type** | Pointer to **string** | The transition type. | [optional] 
**Validators** | Pointer to [**[]WorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) | The validators of the transition. | [optional] 

## Methods

### NewWorkflowTransitions

`func NewWorkflowTransitions() *WorkflowTransitions`

NewWorkflowTransitions instantiates a new WorkflowTransitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionsWithDefaults

`func NewWorkflowTransitionsWithDefaults() *WorkflowTransitions`

NewWorkflowTransitionsWithDefaults instantiates a new WorkflowTransitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *WorkflowTransitions) GetActions() []WorkflowRuleConfiguration`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WorkflowTransitions) GetActionsOk() (*[]WorkflowRuleConfiguration, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WorkflowTransitions) SetActions(v []WorkflowRuleConfiguration)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WorkflowTransitions) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *WorkflowTransitions) GetConditions() ConditionGroupConfiguration`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowTransitions) GetConditionsOk() (*ConditionGroupConfiguration, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowTransitions) SetConditions(v ConditionGroupConfiguration)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowTransitions) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *WorkflowTransitions) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *WorkflowTransitions) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetCustomIssueEventId

`func (o *WorkflowTransitions) GetCustomIssueEventId() string`

GetCustomIssueEventId returns the CustomIssueEventId field if non-nil, zero value otherwise.

### GetCustomIssueEventIdOk

`func (o *WorkflowTransitions) GetCustomIssueEventIdOk() (*string, bool)`

GetCustomIssueEventIdOk returns a tuple with the CustomIssueEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIssueEventId

`func (o *WorkflowTransitions) SetCustomIssueEventId(v string)`

SetCustomIssueEventId sets CustomIssueEventId field to given value.

### HasCustomIssueEventId

`func (o *WorkflowTransitions) HasCustomIssueEventId() bool`

HasCustomIssueEventId returns a boolean if a field has been set.

### SetCustomIssueEventIdNil

`func (o *WorkflowTransitions) SetCustomIssueEventIdNil(b bool)`

 SetCustomIssueEventIdNil sets the value for CustomIssueEventId to be an explicit nil

### UnsetCustomIssueEventId
`func (o *WorkflowTransitions) UnsetCustomIssueEventId()`

UnsetCustomIssueEventId ensures that no value is present for CustomIssueEventId, not even an explicit nil
### GetDescription

`func (o *WorkflowTransitions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTransitions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTransitions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTransitions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkflowTransitions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTransitions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTransitions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTransitions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *WorkflowTransitions) GetLinks() []WorkflowTransitionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkflowTransitions) GetLinksOk() (*[]WorkflowTransitionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkflowTransitions) SetLinks(v []WorkflowTransitionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WorkflowTransitions) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTransitions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTransitions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTransitions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTransitions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTransitions) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTransitions) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTransitions) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTransitions) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetToStatusReference

`func (o *WorkflowTransitions) GetToStatusReference() string`

GetToStatusReference returns the ToStatusReference field if non-nil, zero value otherwise.

### GetToStatusReferenceOk

`func (o *WorkflowTransitions) GetToStatusReferenceOk() (*string, bool)`

GetToStatusReferenceOk returns a tuple with the ToStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStatusReference

`func (o *WorkflowTransitions) SetToStatusReference(v string)`

SetToStatusReference sets ToStatusReference field to given value.

### HasToStatusReference

`func (o *WorkflowTransitions) HasToStatusReference() bool`

HasToStatusReference returns a boolean if a field has been set.

### GetTransitionScreen

`func (o *WorkflowTransitions) GetTransitionScreen() WorkflowRuleConfiguration`

GetTransitionScreen returns the TransitionScreen field if non-nil, zero value otherwise.

### GetTransitionScreenOk

`func (o *WorkflowTransitions) GetTransitionScreenOk() (*WorkflowRuleConfiguration, bool)`

GetTransitionScreenOk returns a tuple with the TransitionScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionScreen

`func (o *WorkflowTransitions) SetTransitionScreen(v WorkflowRuleConfiguration)`

SetTransitionScreen sets TransitionScreen field to given value.

### HasTransitionScreen

`func (o *WorkflowTransitions) HasTransitionScreen() bool`

HasTransitionScreen returns a boolean if a field has been set.

### SetTransitionScreenNil

`func (o *WorkflowTransitions) SetTransitionScreenNil(b bool)`

 SetTransitionScreenNil sets the value for TransitionScreen to be an explicit nil

### UnsetTransitionScreen
`func (o *WorkflowTransitions) UnsetTransitionScreen()`

UnsetTransitionScreen ensures that no value is present for TransitionScreen, not even an explicit nil
### GetTriggers

`func (o *WorkflowTransitions) GetTriggers() []WorkflowTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *WorkflowTransitions) GetTriggersOk() (*[]WorkflowTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *WorkflowTransitions) SetTriggers(v []WorkflowTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *WorkflowTransitions) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetType

`func (o *WorkflowTransitions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTransitions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTransitions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowTransitions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidators

`func (o *WorkflowTransitions) GetValidators() []WorkflowRuleConfiguration`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *WorkflowTransitions) GetValidatorsOk() (*[]WorkflowRuleConfiguration, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *WorkflowTransitions) SetValidators(v []WorkflowRuleConfiguration)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *WorkflowTransitions) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


