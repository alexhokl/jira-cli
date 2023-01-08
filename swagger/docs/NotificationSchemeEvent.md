# NotificationSchemeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**NotificationEvent**](NotificationEvent.md) |  | [optional] 
**Notifications** | Pointer to [**[]EventNotification**](EventNotification.md) |  | [optional] 

## Methods

### NewNotificationSchemeEvent

`func NewNotificationSchemeEvent() *NotificationSchemeEvent`

NewNotificationSchemeEvent instantiates a new NotificationSchemeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeEventWithDefaults

`func NewNotificationSchemeEventWithDefaults() *NotificationSchemeEvent`

NewNotificationSchemeEventWithDefaults instantiates a new NotificationSchemeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NotificationSchemeEvent) GetEvent() NotificationEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationSchemeEvent) GetEventOk() (*NotificationEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationSchemeEvent) SetEvent(v NotificationEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *NotificationSchemeEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetNotifications

`func (o *NotificationSchemeEvent) GetNotifications() []EventNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationSchemeEvent) GetNotificationsOk() (*[]EventNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationSchemeEvent) SetNotifications(v []EventNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationSchemeEvent) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


