# CreateWorkflowTransitionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the transition. The maximum length is 1000 characters. | [optional] 
**From** | Pointer to **[]string** | The statuses the transition can start from. | [optional] 
**Name** | **string** | The name of the transition. The maximum length is 60 characters. | 
**Properties** | Pointer to **map[string]string** | The properties of the transition. | [optional] 
**Rules** | Pointer to [**CreateWorkflowTransitionRulesDetails**](CreateWorkflowTransitionRulesDetails.md) | The rules of the transition. | [optional] 
**Screen** | Pointer to [**CreateWorkflowTransitionScreenDetails**](CreateWorkflowTransitionScreenDetails.md) | The screen of the transition. | [optional] 
**To** | **string** | The status the transition goes to. | 
**Type** | **string** | The type of the transition. | 

## Methods

### NewCreateWorkflowTransitionDetails

`func NewCreateWorkflowTransitionDetails(name string, to string, type_ string, ) *CreateWorkflowTransitionDetails`

NewCreateWorkflowTransitionDetails instantiates a new CreateWorkflowTransitionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowTransitionDetailsWithDefaults

`func NewCreateWorkflowTransitionDetailsWithDefaults() *CreateWorkflowTransitionDetails`

NewCreateWorkflowTransitionDetailsWithDefaults instantiates a new CreateWorkflowTransitionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateWorkflowTransitionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkflowTransitionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkflowTransitionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkflowTransitionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *CreateWorkflowTransitionDetails) GetFrom() []string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateWorkflowTransitionDetails) GetFromOk() (*[]string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateWorkflowTransitionDetails) SetFrom(v []string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreateWorkflowTransitionDetails) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetName

`func (o *CreateWorkflowTransitionDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowTransitionDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowTransitionDetails) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *CreateWorkflowTransitionDetails) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateWorkflowTransitionDetails) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateWorkflowTransitionDetails) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateWorkflowTransitionDetails) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRules

`func (o *CreateWorkflowTransitionDetails) GetRules() CreateWorkflowTransitionRulesDetails`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateWorkflowTransitionDetails) GetRulesOk() (*CreateWorkflowTransitionRulesDetails, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateWorkflowTransitionDetails) SetRules(v CreateWorkflowTransitionRulesDetails)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateWorkflowTransitionDetails) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetScreen

`func (o *CreateWorkflowTransitionDetails) GetScreen() CreateWorkflowTransitionScreenDetails`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *CreateWorkflowTransitionDetails) GetScreenOk() (*CreateWorkflowTransitionScreenDetails, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *CreateWorkflowTransitionDetails) SetScreen(v CreateWorkflowTransitionScreenDetails)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *CreateWorkflowTransitionDetails) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetTo

`func (o *CreateWorkflowTransitionDetails) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateWorkflowTransitionDetails) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateWorkflowTransitionDetails) SetTo(v string)`

SetTo sets To field to given value.


### GetType

`func (o *CreateWorkflowTransitionDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWorkflowTransitionDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWorkflowTransitionDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


