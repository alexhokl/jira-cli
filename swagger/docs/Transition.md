# Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the transition. | 
**From** | **[]string** | The statuses the transition can start from. | 
**Id** | **string** | The ID of the transition. | 
**Name** | **string** | The name of the transition. | 
**Properties** | Pointer to **map[string]interface{}** | The properties of the transition. | [optional] 
**Rules** | Pointer to [**WorkflowRules**](WorkflowRules.md) |  | [optional] 
**Screen** | Pointer to [**TransitionScreenDetails**](TransitionScreenDetails.md) |  | [optional] 
**To** | **string** | The status the transition goes to. | 
**Type** | **string** | The type of the transition. | 

## Methods

### NewTransition

`func NewTransition(description string, from []string, id string, name string, to string, type_ string, ) *Transition`

NewTransition instantiates a new Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionWithDefaults

`func NewTransitionWithDefaults() *Transition`

NewTransitionWithDefaults instantiates a new Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Transition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFrom

`func (o *Transition) GetFrom() []string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Transition) GetFromOk() (*[]string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Transition) SetFrom(v []string)`

SetFrom sets From field to given value.


### GetId

`func (o *Transition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Transition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transition) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *Transition) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Transition) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Transition) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Transition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRules

`func (o *Transition) GetRules() WorkflowRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Transition) GetRulesOk() (*WorkflowRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Transition) SetRules(v WorkflowRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Transition) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetScreen

`func (o *Transition) GetScreen() TransitionScreenDetails`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *Transition) GetScreenOk() (*TransitionScreenDetails, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *Transition) SetScreen(v TransitionScreenDetails)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *Transition) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetTo

`func (o *Transition) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Transition) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Transition) SetTo(v string)`

SetTo sets To field to given value.


### GetType

`func (o *Transition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


