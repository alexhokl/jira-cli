# JiraConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentsEnabled** | Pointer to **bool** | Whether the ability to add attachments to issues is enabled. | [optional] [readonly] 
**IssueLinkingEnabled** | Pointer to **bool** | Whether the ability to link issues is enabled. | [optional] [readonly] 
**SubTasksEnabled** | Pointer to **bool** | Whether the ability to create subtasks for issues is enabled. | [optional] [readonly] 
**TimeTrackingConfiguration** | Pointer to [**TimeTrackingConfiguration**](TimeTrackingConfiguration.md) | The configuration of time tracking. | [optional] [readonly] 
**TimeTrackingEnabled** | Pointer to **bool** | Whether the ability to track time is enabled. This property is deprecated. | [optional] [readonly] 
**UnassignedIssuesAllowed** | Pointer to **bool** | Whether the ability to create unassigned issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details. | [optional] [readonly] 
**VotingEnabled** | Pointer to **bool** | Whether the ability for users to vote on issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details. | [optional] [readonly] 
**WatchingEnabled** | Pointer to **bool** | Whether the ability for users to watch issues is enabled. See [Configuring Jira application options](https://confluence.atlassian.com/x/uYXKM) for details. | [optional] [readonly] 

## Methods

### NewJiraConfiguration

`func NewJiraConfiguration() *JiraConfiguration`

NewJiraConfiguration instantiates a new JiraConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraConfigurationWithDefaults

`func NewJiraConfigurationWithDefaults() *JiraConfiguration`

NewJiraConfigurationWithDefaults instantiates a new JiraConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentsEnabled

`func (o *JiraConfiguration) GetAttachmentsEnabled() bool`

GetAttachmentsEnabled returns the AttachmentsEnabled field if non-nil, zero value otherwise.

### GetAttachmentsEnabledOk

`func (o *JiraConfiguration) GetAttachmentsEnabledOk() (*bool, bool)`

GetAttachmentsEnabledOk returns a tuple with the AttachmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsEnabled

`func (o *JiraConfiguration) SetAttachmentsEnabled(v bool)`

SetAttachmentsEnabled sets AttachmentsEnabled field to given value.

### HasAttachmentsEnabled

`func (o *JiraConfiguration) HasAttachmentsEnabled() bool`

HasAttachmentsEnabled returns a boolean if a field has been set.

### GetIssueLinkingEnabled

`func (o *JiraConfiguration) GetIssueLinkingEnabled() bool`

GetIssueLinkingEnabled returns the IssueLinkingEnabled field if non-nil, zero value otherwise.

### GetIssueLinkingEnabledOk

`func (o *JiraConfiguration) GetIssueLinkingEnabledOk() (*bool, bool)`

GetIssueLinkingEnabledOk returns a tuple with the IssueLinkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLinkingEnabled

`func (o *JiraConfiguration) SetIssueLinkingEnabled(v bool)`

SetIssueLinkingEnabled sets IssueLinkingEnabled field to given value.

### HasIssueLinkingEnabled

`func (o *JiraConfiguration) HasIssueLinkingEnabled() bool`

HasIssueLinkingEnabled returns a boolean if a field has been set.

### GetSubTasksEnabled

`func (o *JiraConfiguration) GetSubTasksEnabled() bool`

GetSubTasksEnabled returns the SubTasksEnabled field if non-nil, zero value otherwise.

### GetSubTasksEnabledOk

`func (o *JiraConfiguration) GetSubTasksEnabledOk() (*bool, bool)`

GetSubTasksEnabledOk returns a tuple with the SubTasksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTasksEnabled

`func (o *JiraConfiguration) SetSubTasksEnabled(v bool)`

SetSubTasksEnabled sets SubTasksEnabled field to given value.

### HasSubTasksEnabled

`func (o *JiraConfiguration) HasSubTasksEnabled() bool`

HasSubTasksEnabled returns a boolean if a field has been set.

### GetTimeTrackingConfiguration

`func (o *JiraConfiguration) GetTimeTrackingConfiguration() TimeTrackingConfiguration`

GetTimeTrackingConfiguration returns the TimeTrackingConfiguration field if non-nil, zero value otherwise.

### GetTimeTrackingConfigurationOk

`func (o *JiraConfiguration) GetTimeTrackingConfigurationOk() (*TimeTrackingConfiguration, bool)`

GetTimeTrackingConfigurationOk returns a tuple with the TimeTrackingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTrackingConfiguration

`func (o *JiraConfiguration) SetTimeTrackingConfiguration(v TimeTrackingConfiguration)`

SetTimeTrackingConfiguration sets TimeTrackingConfiguration field to given value.

### HasTimeTrackingConfiguration

`func (o *JiraConfiguration) HasTimeTrackingConfiguration() bool`

HasTimeTrackingConfiguration returns a boolean if a field has been set.

### GetTimeTrackingEnabled

`func (o *JiraConfiguration) GetTimeTrackingEnabled() bool`

GetTimeTrackingEnabled returns the TimeTrackingEnabled field if non-nil, zero value otherwise.

### GetTimeTrackingEnabledOk

`func (o *JiraConfiguration) GetTimeTrackingEnabledOk() (*bool, bool)`

GetTimeTrackingEnabledOk returns a tuple with the TimeTrackingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTrackingEnabled

`func (o *JiraConfiguration) SetTimeTrackingEnabled(v bool)`

SetTimeTrackingEnabled sets TimeTrackingEnabled field to given value.

### HasTimeTrackingEnabled

`func (o *JiraConfiguration) HasTimeTrackingEnabled() bool`

HasTimeTrackingEnabled returns a boolean if a field has been set.

### GetUnassignedIssuesAllowed

`func (o *JiraConfiguration) GetUnassignedIssuesAllowed() bool`

GetUnassignedIssuesAllowed returns the UnassignedIssuesAllowed field if non-nil, zero value otherwise.

### GetUnassignedIssuesAllowedOk

`func (o *JiraConfiguration) GetUnassignedIssuesAllowedOk() (*bool, bool)`

GetUnassignedIssuesAllowedOk returns a tuple with the UnassignedIssuesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignedIssuesAllowed

`func (o *JiraConfiguration) SetUnassignedIssuesAllowed(v bool)`

SetUnassignedIssuesAllowed sets UnassignedIssuesAllowed field to given value.

### HasUnassignedIssuesAllowed

`func (o *JiraConfiguration) HasUnassignedIssuesAllowed() bool`

HasUnassignedIssuesAllowed returns a boolean if a field has been set.

### GetVotingEnabled

`func (o *JiraConfiguration) GetVotingEnabled() bool`

GetVotingEnabled returns the VotingEnabled field if non-nil, zero value otherwise.

### GetVotingEnabledOk

`func (o *JiraConfiguration) GetVotingEnabledOk() (*bool, bool)`

GetVotingEnabledOk returns a tuple with the VotingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingEnabled

`func (o *JiraConfiguration) SetVotingEnabled(v bool)`

SetVotingEnabled sets VotingEnabled field to given value.

### HasVotingEnabled

`func (o *JiraConfiguration) HasVotingEnabled() bool`

HasVotingEnabled returns a boolean if a field has been set.

### GetWatchingEnabled

`func (o *JiraConfiguration) GetWatchingEnabled() bool`

GetWatchingEnabled returns the WatchingEnabled field if non-nil, zero value otherwise.

### GetWatchingEnabledOk

`func (o *JiraConfiguration) GetWatchingEnabledOk() (*bool, bool)`

GetWatchingEnabledOk returns a tuple with the WatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchingEnabled

`func (o *JiraConfiguration) SetWatchingEnabled(v bool)`

SetWatchingEnabled sets WatchingEnabled field to given value.

### HasWatchingEnabled

`func (o *JiraConfiguration) HasWatchingEnabled() bool`

HasWatchingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


