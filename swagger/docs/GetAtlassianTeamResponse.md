# GetAtlassianTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **float64** | The capacity for the Atlassian team. | [optional] 
**Id** | **string** | The Atlassian team ID. | 
**IssueSourceId** | Pointer to **int64** | The ID of the issue source for the Atlassian team. | [optional] 
**PlanningStyle** | **string** | The planning style for the Atlassian team. This is \&quot;Scrum\&quot; or \&quot;Kanban\&quot;. | 
**SprintLength** | Pointer to **int64** | The sprint length for the Atlassian team. | [optional] 

## Methods

### NewGetAtlassianTeamResponse

`func NewGetAtlassianTeamResponse(id string, planningStyle string, ) *GetAtlassianTeamResponse`

NewGetAtlassianTeamResponse instantiates a new GetAtlassianTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAtlassianTeamResponseWithDefaults

`func NewGetAtlassianTeamResponseWithDefaults() *GetAtlassianTeamResponse`

NewGetAtlassianTeamResponseWithDefaults instantiates a new GetAtlassianTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *GetAtlassianTeamResponse) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *GetAtlassianTeamResponse) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *GetAtlassianTeamResponse) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *GetAtlassianTeamResponse) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetId

`func (o *GetAtlassianTeamResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAtlassianTeamResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAtlassianTeamResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIssueSourceId

`func (o *GetAtlassianTeamResponse) GetIssueSourceId() int64`

GetIssueSourceId returns the IssueSourceId field if non-nil, zero value otherwise.

### GetIssueSourceIdOk

`func (o *GetAtlassianTeamResponse) GetIssueSourceIdOk() (*int64, bool)`

GetIssueSourceIdOk returns a tuple with the IssueSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSourceId

`func (o *GetAtlassianTeamResponse) SetIssueSourceId(v int64)`

SetIssueSourceId sets IssueSourceId field to given value.

### HasIssueSourceId

`func (o *GetAtlassianTeamResponse) HasIssueSourceId() bool`

HasIssueSourceId returns a boolean if a field has been set.

### GetPlanningStyle

`func (o *GetAtlassianTeamResponse) GetPlanningStyle() string`

GetPlanningStyle returns the PlanningStyle field if non-nil, zero value otherwise.

### GetPlanningStyleOk

`func (o *GetAtlassianTeamResponse) GetPlanningStyleOk() (*string, bool)`

GetPlanningStyleOk returns a tuple with the PlanningStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningStyle

`func (o *GetAtlassianTeamResponse) SetPlanningStyle(v string)`

SetPlanningStyle sets PlanningStyle field to given value.


### GetSprintLength

`func (o *GetAtlassianTeamResponse) GetSprintLength() int64`

GetSprintLength returns the SprintLength field if non-nil, zero value otherwise.

### GetSprintLengthOk

`func (o *GetAtlassianTeamResponse) GetSprintLengthOk() (*int64, bool)`

GetSprintLengthOk returns a tuple with the SprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprintLength

`func (o *GetAtlassianTeamResponse) SetSprintLength(v int64)`

SetSprintLength sets SprintLength field to given value.

### HasSprintLength

`func (o *GetAtlassianTeamResponse) HasSprintLength() bool`

HasSprintLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


