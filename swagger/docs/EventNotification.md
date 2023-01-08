# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address. | [optional] 
**Expand** | Pointer to **string** | Expand options that include additional event notification details in the response. | [optional] 
**Field** | Pointer to [**FieldDetails**](FieldDetails.md) | The custom user or group field. | [optional] 
**Group** | Pointer to [**GroupName**](GroupName.md) | The specified group. | [optional] 
**Id** | Pointer to **int64** | The ID of the notification. | [optional] 
**NotificationType** | Pointer to **string** | Identifies the recipients of the notification. | [optional] 
**Parameter** | Pointer to **string** | As a group&#39;s name can change, use of &#x60;recipient&#x60; is recommended. The identifier associated with the &#x60;notificationType&#x60; value that defines the receiver of the notification, where the receiver isn&#39;t implied by &#x60;notificationType&#x60; value. So, when &#x60;notificationType&#x60; is:   *  &#x60;User&#x60; The &#x60;parameter&#x60; is the user account ID.  *  &#x60;Group&#x60; The &#x60;parameter&#x60; is the group name.  *  &#x60;ProjectRole&#x60; The &#x60;parameter&#x60; is the project role ID.  *  &#x60;UserCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field.  *  &#x60;GroupCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field. | [optional] 
**ProjectRole** | Pointer to [**ProjectRole**](ProjectRole.md) | The specified project role. | [optional] 
**Recipient** | Pointer to **string** | The identifier associated with the &#x60;notificationType&#x60; value that defines the receiver of the notification, where the receiver isn&#39;t implied by the &#x60;notificationType&#x60; value. So, when &#x60;notificationType&#x60; is:   *  &#x60;User&#x60;, &#x60;recipient&#x60; is the user account ID.  *  &#x60;Group&#x60;, &#x60;recipient&#x60; is the group ID.  *  &#x60;ProjectRole&#x60;, &#x60;recipient&#x60; is the project role ID.  *  &#x60;UserCustomField&#x60;, &#x60;recipient&#x60; is the ID of the custom field.  *  &#x60;GroupCustomField&#x60;, &#x60;recipient&#x60; is the ID of the custom field. | [optional] 
**User** | Pointer to [**UserDetails**](UserDetails.md) | The specified user. | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification() *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *EventNotification) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EventNotification) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EventNotification) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *EventNotification) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetExpand

`func (o *EventNotification) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *EventNotification) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *EventNotification) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *EventNotification) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetField

`func (o *EventNotification) GetField() FieldDetails`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *EventNotification) GetFieldOk() (*FieldDetails, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *EventNotification) SetField(v FieldDetails)`

SetField sets Field field to given value.

### HasField

`func (o *EventNotification) HasField() bool`

HasField returns a boolean if a field has been set.

### GetGroup

`func (o *EventNotification) GetGroup() GroupName`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EventNotification) GetGroupOk() (*GroupName, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EventNotification) SetGroup(v GroupName)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EventNotification) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *EventNotification) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventNotification) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventNotification) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventNotification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationType

`func (o *EventNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *EventNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *EventNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *EventNotification) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetParameter

`func (o *EventNotification) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *EventNotification) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *EventNotification) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *EventNotification) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetProjectRole

`func (o *EventNotification) GetProjectRole() ProjectRole`

GetProjectRole returns the ProjectRole field if non-nil, zero value otherwise.

### GetProjectRoleOk

`func (o *EventNotification) GetProjectRoleOk() (*ProjectRole, bool)`

GetProjectRoleOk returns a tuple with the ProjectRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRole

`func (o *EventNotification) SetProjectRole(v ProjectRole)`

SetProjectRole sets ProjectRole field to given value.

### HasProjectRole

`func (o *EventNotification) HasProjectRole() bool`

HasProjectRole returns a boolean if a field has been set.

### GetRecipient

`func (o *EventNotification) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EventNotification) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EventNotification) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *EventNotification) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetUser

`func (o *EventNotification) GetUser() UserDetails`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventNotification) GetUserOk() (*UserDetails, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventNotification) SetUser(v UserDetails)`

SetUser sets User field to given value.

### HasUser

`func (o *EventNotification) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


