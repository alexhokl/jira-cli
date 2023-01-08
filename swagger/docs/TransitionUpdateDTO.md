# TransitionUpdateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]WorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) | The post-functions of the transition. | [optional] 
**Conditions** | Pointer to [**NullableConditionGroupUpdate**](ConditionGroupUpdate.md) |  | [optional] 
**CustomIssueEventId** | Pointer to **string** | The custom event ID of the transition. | [optional] 
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

### NewTransitionUpdateDTO

`func NewTransitionUpdateDTO() *TransitionUpdateDTO`

NewTransitionUpdateDTO instantiates a new TransitionUpdateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionUpdateDTOWithDefaults

`func NewTransitionUpdateDTOWithDefaults() *TransitionUpdateDTO`

NewTransitionUpdateDTOWithDefaults instantiates a new TransitionUpdateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TransitionUpdateDTO) GetActions() []WorkflowRuleConfiguration`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TransitionUpdateDTO) GetActionsOk() (*[]WorkflowRuleConfiguration, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TransitionUpdateDTO) SetActions(v []WorkflowRuleConfiguration)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TransitionUpdateDTO) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *TransitionUpdateDTO) GetConditions() ConditionGroupUpdate`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TransitionUpdateDTO) GetConditionsOk() (*ConditionGroupUpdate, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TransitionUpdateDTO) SetConditions(v ConditionGroupUpdate)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TransitionUpdateDTO) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *TransitionUpdateDTO) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *TransitionUpdateDTO) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetCustomIssueEventId

`func (o *TransitionUpdateDTO) GetCustomIssueEventId() string`

GetCustomIssueEventId returns the CustomIssueEventId field if non-nil, zero value otherwise.

### GetCustomIssueEventIdOk

`func (o *TransitionUpdateDTO) GetCustomIssueEventIdOk() (*string, bool)`

GetCustomIssueEventIdOk returns a tuple with the CustomIssueEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIssueEventId

`func (o *TransitionUpdateDTO) SetCustomIssueEventId(v string)`

SetCustomIssueEventId sets CustomIssueEventId field to given value.

### HasCustomIssueEventId

`func (o *TransitionUpdateDTO) HasCustomIssueEventId() bool`

HasCustomIssueEventId returns a boolean if a field has been set.

### GetDescription

`func (o *TransitionUpdateDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitionUpdateDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitionUpdateDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransitionUpdateDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TransitionUpdateDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitionUpdateDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitionUpdateDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransitionUpdateDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *TransitionUpdateDTO) GetLinks() []WorkflowTransitionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransitionUpdateDTO) GetLinksOk() (*[]WorkflowTransitionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransitionUpdateDTO) SetLinks(v []WorkflowTransitionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransitionUpdateDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TransitionUpdateDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitionUpdateDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitionUpdateDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitionUpdateDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *TransitionUpdateDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TransitionUpdateDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TransitionUpdateDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TransitionUpdateDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetToStatusReference

`func (o *TransitionUpdateDTO) GetToStatusReference() string`

GetToStatusReference returns the ToStatusReference field if non-nil, zero value otherwise.

### GetToStatusReferenceOk

`func (o *TransitionUpdateDTO) GetToStatusReferenceOk() (*string, bool)`

GetToStatusReferenceOk returns a tuple with the ToStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStatusReference

`func (o *TransitionUpdateDTO) SetToStatusReference(v string)`

SetToStatusReference sets ToStatusReference field to given value.

### HasToStatusReference

`func (o *TransitionUpdateDTO) HasToStatusReference() bool`

HasToStatusReference returns a boolean if a field has been set.

### GetTransitionScreen

`func (o *TransitionUpdateDTO) GetTransitionScreen() WorkflowRuleConfiguration`

GetTransitionScreen returns the TransitionScreen field if non-nil, zero value otherwise.

### GetTransitionScreenOk

`func (o *TransitionUpdateDTO) GetTransitionScreenOk() (*WorkflowRuleConfiguration, bool)`

GetTransitionScreenOk returns a tuple with the TransitionScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionScreen

`func (o *TransitionUpdateDTO) SetTransitionScreen(v WorkflowRuleConfiguration)`

SetTransitionScreen sets TransitionScreen field to given value.

### HasTransitionScreen

`func (o *TransitionUpdateDTO) HasTransitionScreen() bool`

HasTransitionScreen returns a boolean if a field has been set.

### SetTransitionScreenNil

`func (o *TransitionUpdateDTO) SetTransitionScreenNil(b bool)`

 SetTransitionScreenNil sets the value for TransitionScreen to be an explicit nil

### UnsetTransitionScreen
`func (o *TransitionUpdateDTO) UnsetTransitionScreen()`

UnsetTransitionScreen ensures that no value is present for TransitionScreen, not even an explicit nil
### GetTriggers

`func (o *TransitionUpdateDTO) GetTriggers() []WorkflowTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TransitionUpdateDTO) GetTriggersOk() (*[]WorkflowTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TransitionUpdateDTO) SetTriggers(v []WorkflowTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TransitionUpdateDTO) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetType

`func (o *TransitionUpdateDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitionUpdateDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitionUpdateDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransitionUpdateDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidators

`func (o *TransitionUpdateDTO) GetValidators() []WorkflowRuleConfiguration`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *TransitionUpdateDTO) GetValidatorsOk() (*[]WorkflowRuleConfiguration, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *TransitionUpdateDTO) SetValidators(v []WorkflowRuleConfiguration)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *TransitionUpdateDTO) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


