# AddAtlassianTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **float64** | The capacity for the Atlassian team. | [optional] 
**Id** | **string** | The Atlassian team ID. | 
**IssueSourceId** | Pointer to **int64** | The ID of the issue source for the Atlassian team. | [optional] 
**PlanningStyle** | **string** | The planning style for the Atlassian team. This must be \&quot;Scrum\&quot; or \&quot;Kanban\&quot;. | 
**SprintLength** | Pointer to **int64** | The sprint length for the Atlassian team. | [optional] 

## Methods

### NewAddAtlassianTeamRequest

`func NewAddAtlassianTeamRequest(id string, planningStyle string, ) *AddAtlassianTeamRequest`

NewAddAtlassianTeamRequest instantiates a new AddAtlassianTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAtlassianTeamRequestWithDefaults

`func NewAddAtlassianTeamRequestWithDefaults() *AddAtlassianTeamRequest`

NewAddAtlassianTeamRequestWithDefaults instantiates a new AddAtlassianTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *AddAtlassianTeamRequest) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AddAtlassianTeamRequest) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AddAtlassianTeamRequest) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AddAtlassianTeamRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetId

`func (o *AddAtlassianTeamRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddAtlassianTeamRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddAtlassianTeamRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIssueSourceId

`func (o *AddAtlassianTeamRequest) GetIssueSourceId() int64`

GetIssueSourceId returns the IssueSourceId field if non-nil, zero value otherwise.

### GetIssueSourceIdOk

`func (o *AddAtlassianTeamRequest) GetIssueSourceIdOk() (*int64, bool)`

GetIssueSourceIdOk returns a tuple with the IssueSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSourceId

`func (o *AddAtlassianTeamRequest) SetIssueSourceId(v int64)`

SetIssueSourceId sets IssueSourceId field to given value.

### HasIssueSourceId

`func (o *AddAtlassianTeamRequest) HasIssueSourceId() bool`

HasIssueSourceId returns a boolean if a field has been set.

### GetPlanningStyle

`func (o *AddAtlassianTeamRequest) GetPlanningStyle() string`

GetPlanningStyle returns the PlanningStyle field if non-nil, zero value otherwise.

### GetPlanningStyleOk

`func (o *AddAtlassianTeamRequest) GetPlanningStyleOk() (*string, bool)`

GetPlanningStyleOk returns a tuple with the PlanningStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningStyle

`func (o *AddAtlassianTeamRequest) SetPlanningStyle(v string)`

SetPlanningStyle sets PlanningStyle field to given value.


### GetSprintLength

`func (o *AddAtlassianTeamRequest) GetSprintLength() int64`

GetSprintLength returns the SprintLength field if non-nil, zero value otherwise.

### GetSprintLengthOk

`func (o *AddAtlassianTeamRequest) GetSprintLengthOk() (*int64, bool)`

GetSprintLengthOk returns a tuple with the SprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprintLength

`func (o *AddAtlassianTeamRequest) SetSprintLength(v int64)`

SetSprintLength sets SprintLength field to given value.

### HasSprintLength

`func (o *AddAtlassianTeamRequest) HasSprintLength() bool`

HasSprintLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


