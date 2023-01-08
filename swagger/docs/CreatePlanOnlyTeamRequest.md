# CreatePlanOnlyTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **float64** | The capacity for the plan-only team. | [optional] 
**IssueSourceId** | Pointer to **int64** | The ID of the issue source for the plan-only team. | [optional] 
**MemberAccountIds** | Pointer to **[]string** | The account IDs of the plan-only team members. | [optional] 
**Name** | **string** | The plan-only team name. | 
**PlanningStyle** | **string** | The planning style for the plan-only team. This must be \&quot;Scrum\&quot; or \&quot;Kanban\&quot;. | 
**SprintLength** | Pointer to **int64** | The sprint length for the plan-only team. | [optional] 

## Methods

### NewCreatePlanOnlyTeamRequest

`func NewCreatePlanOnlyTeamRequest(name string, planningStyle string, ) *CreatePlanOnlyTeamRequest`

NewCreatePlanOnlyTeamRequest instantiates a new CreatePlanOnlyTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanOnlyTeamRequestWithDefaults

`func NewCreatePlanOnlyTeamRequestWithDefaults() *CreatePlanOnlyTeamRequest`

NewCreatePlanOnlyTeamRequestWithDefaults instantiates a new CreatePlanOnlyTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *CreatePlanOnlyTeamRequest) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CreatePlanOnlyTeamRequest) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CreatePlanOnlyTeamRequest) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CreatePlanOnlyTeamRequest) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetIssueSourceId

`func (o *CreatePlanOnlyTeamRequest) GetIssueSourceId() int64`

GetIssueSourceId returns the IssueSourceId field if non-nil, zero value otherwise.

### GetIssueSourceIdOk

`func (o *CreatePlanOnlyTeamRequest) GetIssueSourceIdOk() (*int64, bool)`

GetIssueSourceIdOk returns a tuple with the IssueSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSourceId

`func (o *CreatePlanOnlyTeamRequest) SetIssueSourceId(v int64)`

SetIssueSourceId sets IssueSourceId field to given value.

### HasIssueSourceId

`func (o *CreatePlanOnlyTeamRequest) HasIssueSourceId() bool`

HasIssueSourceId returns a boolean if a field has been set.

### GetMemberAccountIds

`func (o *CreatePlanOnlyTeamRequest) GetMemberAccountIds() []string`

GetMemberAccountIds returns the MemberAccountIds field if non-nil, zero value otherwise.

### GetMemberAccountIdsOk

`func (o *CreatePlanOnlyTeamRequest) GetMemberAccountIdsOk() (*[]string, bool)`

GetMemberAccountIdsOk returns a tuple with the MemberAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAccountIds

`func (o *CreatePlanOnlyTeamRequest) SetMemberAccountIds(v []string)`

SetMemberAccountIds sets MemberAccountIds field to given value.

### HasMemberAccountIds

`func (o *CreatePlanOnlyTeamRequest) HasMemberAccountIds() bool`

HasMemberAccountIds returns a boolean if a field has been set.

### GetName

`func (o *CreatePlanOnlyTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlanOnlyTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlanOnlyTeamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlanningStyle

`func (o *CreatePlanOnlyTeamRequest) GetPlanningStyle() string`

GetPlanningStyle returns the PlanningStyle field if non-nil, zero value otherwise.

### GetPlanningStyleOk

`func (o *CreatePlanOnlyTeamRequest) GetPlanningStyleOk() (*string, bool)`

GetPlanningStyleOk returns a tuple with the PlanningStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningStyle

`func (o *CreatePlanOnlyTeamRequest) SetPlanningStyle(v string)`

SetPlanningStyle sets PlanningStyle field to given value.


### GetSprintLength

`func (o *CreatePlanOnlyTeamRequest) GetSprintLength() int64`

GetSprintLength returns the SprintLength field if non-nil, zero value otherwise.

### GetSprintLengthOk

`func (o *CreatePlanOnlyTeamRequest) GetSprintLengthOk() (*int64, bool)`

GetSprintLengthOk returns a tuple with the SprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprintLength

`func (o *CreatePlanOnlyTeamRequest) SetSprintLength(v int64)`

SetSprintLength sets SprintLength field to given value.

### HasSprintLength

`func (o *CreatePlanOnlyTeamRequest) HasSprintLength() bool`

HasSprintLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


