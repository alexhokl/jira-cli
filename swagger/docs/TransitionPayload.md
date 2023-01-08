# TransitionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]RulePayload**](RulePayload.md) | The actions that are performed when the transition is made | [optional] 
**Conditions** | Pointer to [**ConditionGroupPayload**](ConditionGroupPayload.md) |  | [optional] 
**CustomIssueEventId** | Pointer to **string** | Mechanism in Jira for triggering certain actions, like notifications, automations, etc. Unless a custom notification scheme is configure, it&#39;s better not to provide any value here | [optional] 
**Description** | Pointer to **string** | The description of the transition | [optional] 
**From** | Pointer to [**[]FromLayoutPayload**](FromLayoutPayload.md) | The statuses that the transition can be made from | [optional] 
**Id** | Pointer to **int32** | The id of the transition | [optional] 
**Name** | Pointer to **string** | The name of the transition | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the transition | [optional] 
**To** | Pointer to [**ToLayoutPayload**](ToLayoutPayload.md) |  | [optional] 
**TransitionScreen** | Pointer to [**RulePayload**](RulePayload.md) |  | [optional] 
**Triggers** | Pointer to [**[]RulePayload**](RulePayload.md) | The triggers that are performed when the transition is made | [optional] 
**Type** | Pointer to **string** | The type of the transition | [optional] 
**Validators** | Pointer to [**[]RulePayload**](RulePayload.md) | The validators that are performed when the transition is made | [optional] 

## Methods

### NewTransitionPayload

`func NewTransitionPayload() *TransitionPayload`

NewTransitionPayload instantiates a new TransitionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionPayloadWithDefaults

`func NewTransitionPayloadWithDefaults() *TransitionPayload`

NewTransitionPayloadWithDefaults instantiates a new TransitionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TransitionPayload) GetActions() []RulePayload`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TransitionPayload) GetActionsOk() (*[]RulePayload, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TransitionPayload) SetActions(v []RulePayload)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TransitionPayload) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *TransitionPayload) GetConditions() ConditionGroupPayload`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TransitionPayload) GetConditionsOk() (*ConditionGroupPayload, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TransitionPayload) SetConditions(v ConditionGroupPayload)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TransitionPayload) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCustomIssueEventId

`func (o *TransitionPayload) GetCustomIssueEventId() string`

GetCustomIssueEventId returns the CustomIssueEventId field if non-nil, zero value otherwise.

### GetCustomIssueEventIdOk

`func (o *TransitionPayload) GetCustomIssueEventIdOk() (*string, bool)`

GetCustomIssueEventIdOk returns a tuple with the CustomIssueEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIssueEventId

`func (o *TransitionPayload) SetCustomIssueEventId(v string)`

SetCustomIssueEventId sets CustomIssueEventId field to given value.

### HasCustomIssueEventId

`func (o *TransitionPayload) HasCustomIssueEventId() bool`

HasCustomIssueEventId returns a boolean if a field has been set.

### GetDescription

`func (o *TransitionPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitionPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitionPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransitionPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *TransitionPayload) GetFrom() []FromLayoutPayload`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransitionPayload) GetFromOk() (*[]FromLayoutPayload, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransitionPayload) SetFrom(v []FromLayoutPayload)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TransitionPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *TransitionPayload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitionPayload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitionPayload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TransitionPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TransitionPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitionPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitionPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitionPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *TransitionPayload) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TransitionPayload) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TransitionPayload) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TransitionPayload) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTo

`func (o *TransitionPayload) GetTo() ToLayoutPayload`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransitionPayload) GetToOk() (*ToLayoutPayload, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransitionPayload) SetTo(v ToLayoutPayload)`

SetTo sets To field to given value.

### HasTo

`func (o *TransitionPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransitionScreen

`func (o *TransitionPayload) GetTransitionScreen() RulePayload`

GetTransitionScreen returns the TransitionScreen field if non-nil, zero value otherwise.

### GetTransitionScreenOk

`func (o *TransitionPayload) GetTransitionScreenOk() (*RulePayload, bool)`

GetTransitionScreenOk returns a tuple with the TransitionScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionScreen

`func (o *TransitionPayload) SetTransitionScreen(v RulePayload)`

SetTransitionScreen sets TransitionScreen field to given value.

### HasTransitionScreen

`func (o *TransitionPayload) HasTransitionScreen() bool`

HasTransitionScreen returns a boolean if a field has been set.

### GetTriggers

`func (o *TransitionPayload) GetTriggers() []RulePayload`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TransitionPayload) GetTriggersOk() (*[]RulePayload, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TransitionPayload) SetTriggers(v []RulePayload)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TransitionPayload) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetType

`func (o *TransitionPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitionPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitionPayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransitionPayload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidators

`func (o *TransitionPayload) GetValidators() []RulePayload`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *TransitionPayload) GetValidatorsOk() (*[]RulePayload, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *TransitionPayload) SetValidators(v []RulePayload)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *TransitionPayload) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


