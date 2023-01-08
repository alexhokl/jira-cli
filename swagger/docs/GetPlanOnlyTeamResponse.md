# GetPlanOnlyTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **float64** | The capacity for the plan-only team. | [optional] 
**Id** | **int64** | The plan-only team ID. | 
**IssueSourceId** | Pointer to **int64** | The ID of the issue source for the plan-only team. | [optional] 
**MemberAccountIds** | Pointer to **[]string** | The account IDs of the plan-only team members. | [optional] 
**Name** | **string** | The plan-only team name. | 
**PlanningStyle** | **string** | The planning style for the plan-only team. This is \&quot;Scrum\&quot; or \&quot;Kanban\&quot;. | 
**SprintLength** | Pointer to **int64** | The sprint length for the plan-only team. | [optional] 

## Methods

### NewGetPlanOnlyTeamResponse

`func NewGetPlanOnlyTeamResponse(id int64, name string, planningStyle string, ) *GetPlanOnlyTeamResponse`

NewGetPlanOnlyTeamResponse instantiates a new GetPlanOnlyTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlanOnlyTeamResponseWithDefaults

`func NewGetPlanOnlyTeamResponseWithDefaults() *GetPlanOnlyTeamResponse`

NewGetPlanOnlyTeamResponseWithDefaults instantiates a new GetPlanOnlyTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *GetPlanOnlyTeamResponse) GetCapacity() float64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *GetPlanOnlyTeamResponse) GetCapacityOk() (*float64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *GetPlanOnlyTeamResponse) SetCapacity(v float64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *GetPlanOnlyTeamResponse) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetId

`func (o *GetPlanOnlyTeamResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlanOnlyTeamResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlanOnlyTeamResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueSourceId

`func (o *GetPlanOnlyTeamResponse) GetIssueSourceId() int64`

GetIssueSourceId returns the IssueSourceId field if non-nil, zero value otherwise.

### GetIssueSourceIdOk

`func (o *GetPlanOnlyTeamResponse) GetIssueSourceIdOk() (*int64, bool)`

GetIssueSourceIdOk returns a tuple with the IssueSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSourceId

`func (o *GetPlanOnlyTeamResponse) SetIssueSourceId(v int64)`

SetIssueSourceId sets IssueSourceId field to given value.

### HasIssueSourceId

`func (o *GetPlanOnlyTeamResponse) HasIssueSourceId() bool`

HasIssueSourceId returns a boolean if a field has been set.

### GetMemberAccountIds

`func (o *GetPlanOnlyTeamResponse) GetMemberAccountIds() []string`

GetMemberAccountIds returns the MemberAccountIds field if non-nil, zero value otherwise.

### GetMemberAccountIdsOk

`func (o *GetPlanOnlyTeamResponse) GetMemberAccountIdsOk() (*[]string, bool)`

GetMemberAccountIdsOk returns a tuple with the MemberAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAccountIds

`func (o *GetPlanOnlyTeamResponse) SetMemberAccountIds(v []string)`

SetMemberAccountIds sets MemberAccountIds field to given value.

### HasMemberAccountIds

`func (o *GetPlanOnlyTeamResponse) HasMemberAccountIds() bool`

HasMemberAccountIds returns a boolean if a field has been set.

### GetName

`func (o *GetPlanOnlyTeamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPlanOnlyTeamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPlanOnlyTeamResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPlanningStyle

`func (o *GetPlanOnlyTeamResponse) GetPlanningStyle() string`

GetPlanningStyle returns the PlanningStyle field if non-nil, zero value otherwise.

### GetPlanningStyleOk

`func (o *GetPlanOnlyTeamResponse) GetPlanningStyleOk() (*string, bool)`

GetPlanningStyleOk returns a tuple with the PlanningStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningStyle

`func (o *GetPlanOnlyTeamResponse) SetPlanningStyle(v string)`

SetPlanningStyle sets PlanningStyle field to given value.


### GetSprintLength

`func (o *GetPlanOnlyTeamResponse) GetSprintLength() int64`

GetSprintLength returns the SprintLength field if non-nil, zero value otherwise.

### GetSprintLengthOk

`func (o *GetPlanOnlyTeamResponse) GetSprintLengthOk() (*int64, bool)`

GetSprintLengthOk returns a tuple with the SprintLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprintLength

`func (o *GetPlanOnlyTeamResponse) SetSprintLength(v int64)`

SetSprintLength sets SprintLength field to given value.

### HasSprintLength

`func (o *GetPlanOnlyTeamResponse) HasSprintLength() bool`

HasSprintLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


