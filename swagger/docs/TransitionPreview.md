# TransitionPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]PreviewRuleConfiguration**](PreviewRuleConfiguration.md) | The post-functions of the transition. | [optional] 
**Conditions** | Pointer to [**NullablePreviewConditionGroupConfiguration**](PreviewConditionGroupConfiguration.md) |  | [optional] 
**CustomIssueEventId** | Pointer to **string** | The custom issue event ID for the transition. | [optional] 
**Description** | Pointer to **string** | The description of the transition. | [optional] 
**Id** | Pointer to **string** | The ID of the transition. | [optional] 
**Links** | Pointer to [**[]TransitionLink**](TransitionLink.md) | The statuses the transition can start from, and the mapping of ports between the statuses. | [optional] 
**Name** | Pointer to **string** | The name of the transition. | [optional] 
**ToStatusReference** | Pointer to **string** | The status the transition goes to. | [optional] 
**TransitionScreen** | Pointer to [**NullablePreviewRuleConfiguration**](PreviewRuleConfiguration.md) |  | [optional] 
**Triggers** | Pointer to [**[]PreviewTrigger**](PreviewTrigger.md) | The triggers of the transition. | [optional] 
**Type** | Pointer to **string** | The transition type. | [optional] 
**Validators** | Pointer to [**[]PreviewRuleConfiguration**](PreviewRuleConfiguration.md) | The validators of the transition. | [optional] 

## Methods

### NewTransitionPreview

`func NewTransitionPreview() *TransitionPreview`

NewTransitionPreview instantiates a new TransitionPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionPreviewWithDefaults

`func NewTransitionPreviewWithDefaults() *TransitionPreview`

NewTransitionPreviewWithDefaults instantiates a new TransitionPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TransitionPreview) GetActions() []PreviewRuleConfiguration`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TransitionPreview) GetActionsOk() (*[]PreviewRuleConfiguration, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TransitionPreview) SetActions(v []PreviewRuleConfiguration)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TransitionPreview) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *TransitionPreview) GetConditions() PreviewConditionGroupConfiguration`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TransitionPreview) GetConditionsOk() (*PreviewConditionGroupConfiguration, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TransitionPreview) SetConditions(v PreviewConditionGroupConfiguration)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TransitionPreview) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *TransitionPreview) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *TransitionPreview) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetCustomIssueEventId

`func (o *TransitionPreview) GetCustomIssueEventId() string`

GetCustomIssueEventId returns the CustomIssueEventId field if non-nil, zero value otherwise.

### GetCustomIssueEventIdOk

`func (o *TransitionPreview) GetCustomIssueEventIdOk() (*string, bool)`

GetCustomIssueEventIdOk returns a tuple with the CustomIssueEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIssueEventId

`func (o *TransitionPreview) SetCustomIssueEventId(v string)`

SetCustomIssueEventId sets CustomIssueEventId field to given value.

### HasCustomIssueEventId

`func (o *TransitionPreview) HasCustomIssueEventId() bool`

HasCustomIssueEventId returns a boolean if a field has been set.

### GetDescription

`func (o *TransitionPreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitionPreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitionPreview) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransitionPreview) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *TransitionPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitionPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitionPreview) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransitionPreview) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *TransitionPreview) GetLinks() []TransitionLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransitionPreview) GetLinksOk() (*[]TransitionLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransitionPreview) SetLinks(v []TransitionLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransitionPreview) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TransitionPreview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransitionPreview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransitionPreview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransitionPreview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToStatusReference

`func (o *TransitionPreview) GetToStatusReference() string`

GetToStatusReference returns the ToStatusReference field if non-nil, zero value otherwise.

### GetToStatusReferenceOk

`func (o *TransitionPreview) GetToStatusReferenceOk() (*string, bool)`

GetToStatusReferenceOk returns a tuple with the ToStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStatusReference

`func (o *TransitionPreview) SetToStatusReference(v string)`

SetToStatusReference sets ToStatusReference field to given value.

### HasToStatusReference

`func (o *TransitionPreview) HasToStatusReference() bool`

HasToStatusReference returns a boolean if a field has been set.

### GetTransitionScreen

`func (o *TransitionPreview) GetTransitionScreen() PreviewRuleConfiguration`

GetTransitionScreen returns the TransitionScreen field if non-nil, zero value otherwise.

### GetTransitionScreenOk

`func (o *TransitionPreview) GetTransitionScreenOk() (*PreviewRuleConfiguration, bool)`

GetTransitionScreenOk returns a tuple with the TransitionScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionScreen

`func (o *TransitionPreview) SetTransitionScreen(v PreviewRuleConfiguration)`

SetTransitionScreen sets TransitionScreen field to given value.

### HasTransitionScreen

`func (o *TransitionPreview) HasTransitionScreen() bool`

HasTransitionScreen returns a boolean if a field has been set.

### SetTransitionScreenNil

`func (o *TransitionPreview) SetTransitionScreenNil(b bool)`

 SetTransitionScreenNil sets the value for TransitionScreen to be an explicit nil

### UnsetTransitionScreen
`func (o *TransitionPreview) UnsetTransitionScreen()`

UnsetTransitionScreen ensures that no value is present for TransitionScreen, not even an explicit nil
### GetTriggers

`func (o *TransitionPreview) GetTriggers() []PreviewTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TransitionPreview) GetTriggersOk() (*[]PreviewTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TransitionPreview) SetTriggers(v []PreviewTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TransitionPreview) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetType

`func (o *TransitionPreview) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransitionPreview) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransitionPreview) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransitionPreview) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidators

`func (o *TransitionPreview) GetValidators() []PreviewRuleConfiguration`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *TransitionPreview) GetValidatorsOk() (*[]PreviewRuleConfiguration, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *TransitionPreview) SetValidators(v []PreviewRuleConfiguration)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *TransitionPreview) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


