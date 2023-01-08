# CreateWorkflowDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the workflow. The maximum length is 1000 characters. | [optional] 
**Name** | **string** | The name of the workflow. The name must be unique. The maximum length is 255 characters. Characters can be separated by a whitespace but the name cannot start or end with a whitespace. | 
**Statuses** | [**[]CreateWorkflowStatusDetails**](CreateWorkflowStatusDetails.md) | The statuses of the workflow. Any status that does not include a transition is added to the workflow without a transition. | 
**Transitions** | [**[]CreateWorkflowTransitionDetails**](CreateWorkflowTransitionDetails.md) | The transitions of the workflow. For the request to be valid, these transitions must:   *  include one *initial* transition.  *  not use the same name for a *global* and *directed* transition.  *  have a unique name for each *global* transition.  *  have a unique &#39;to&#39; status for each *global* transition.  *  have unique names for each transition from a status.  *  not have a &#39;from&#39; status on *initial* and *global* transitions.  *  have a &#39;from&#39; status on *directed* transitions.  All the transition statuses must be included in &#x60;statuses&#x60;. | 

## Methods

### NewCreateWorkflowDetails

`func NewCreateWorkflowDetails(name string, statuses []CreateWorkflowStatusDetails, transitions []CreateWorkflowTransitionDetails, ) *CreateWorkflowDetails`

NewCreateWorkflowDetails instantiates a new CreateWorkflowDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowDetailsWithDefaults

`func NewCreateWorkflowDetailsWithDefaults() *CreateWorkflowDetails`

NewCreateWorkflowDetailsWithDefaults instantiates a new CreateWorkflowDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateWorkflowDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkflowDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkflowDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkflowDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateWorkflowDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowDetails) SetName(v string)`

SetName sets Name field to given value.


### GetStatuses

`func (o *CreateWorkflowDetails) GetStatuses() []CreateWorkflowStatusDetails`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *CreateWorkflowDetails) GetStatusesOk() (*[]CreateWorkflowStatusDetails, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *CreateWorkflowDetails) SetStatuses(v []CreateWorkflowStatusDetails)`

SetStatuses sets Statuses field to given value.


### GetTransitions

`func (o *CreateWorkflowDetails) GetTransitions() []CreateWorkflowTransitionDetails`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *CreateWorkflowDetails) GetTransitionsOk() (*[]CreateWorkflowTransitionDetails, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *CreateWorkflowDetails) SetTransitions(v []CreateWorkflowTransitionDetails)`

SetTransitions sets Transitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


