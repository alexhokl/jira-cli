# NotificationRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **bool** | Whether the notification should be sent to the issue&#39;s assignees. | [optional] 
**GroupIds** | Pointer to **[]string** | List of groupIds to receive the notification. | [optional] 
**Groups** | Pointer to [**[]GroupName**](GroupName.md) | List of groups to receive the notification. | [optional] 
**Reporter** | Pointer to **bool** | Whether the notification should be sent to the issue&#39;s reporter. | [optional] 
**Users** | Pointer to [**[]UserDetails**](UserDetails.md) | List of users to receive the notification. | [optional] 
**Voters** | Pointer to **bool** | Whether the notification should be sent to the issue&#39;s voters. | [optional] 
**Watchers** | Pointer to **bool** | Whether the notification should be sent to the issue&#39;s watchers. | [optional] 

## Methods

### NewNotificationRecipients

`func NewNotificationRecipients() *NotificationRecipients`

NewNotificationRecipients instantiates a new NotificationRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRecipientsWithDefaults

`func NewNotificationRecipientsWithDefaults() *NotificationRecipients`

NewNotificationRecipientsWithDefaults instantiates a new NotificationRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *NotificationRecipients) GetAssignee() bool`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *NotificationRecipients) GetAssigneeOk() (*bool, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *NotificationRecipients) SetAssignee(v bool)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *NotificationRecipients) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetGroupIds

`func (o *NotificationRecipients) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *NotificationRecipients) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *NotificationRecipients) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *NotificationRecipients) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetGroups

`func (o *NotificationRecipients) GetGroups() []GroupName`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *NotificationRecipients) GetGroupsOk() (*[]GroupName, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *NotificationRecipients) SetGroups(v []GroupName)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *NotificationRecipients) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetReporter

`func (o *NotificationRecipients) GetReporter() bool`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *NotificationRecipients) GetReporterOk() (*bool, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *NotificationRecipients) SetReporter(v bool)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *NotificationRecipients) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetUsers

`func (o *NotificationRecipients) GetUsers() []UserDetails`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NotificationRecipients) GetUsersOk() (*[]UserDetails, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NotificationRecipients) SetUsers(v []UserDetails)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NotificationRecipients) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetVoters

`func (o *NotificationRecipients) GetVoters() bool`

GetVoters returns the Voters field if non-nil, zero value otherwise.

### GetVotersOk

`func (o *NotificationRecipients) GetVotersOk() (*bool, bool)`

GetVotersOk returns a tuple with the Voters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoters

`func (o *NotificationRecipients) SetVoters(v bool)`

SetVoters sets Voters field to given value.

### HasVoters

`func (o *NotificationRecipients) HasVoters() bool`

HasVoters returns a boolean if a field has been set.

### GetWatchers

`func (o *NotificationRecipients) GetWatchers() bool`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *NotificationRecipients) GetWatchersOk() (*bool, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *NotificationRecipients) SetWatchers(v bool)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *NotificationRecipients) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


