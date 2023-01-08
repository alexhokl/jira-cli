# NotificationSchemeEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**NotificationSchemeEventIDPayload**](NotificationSchemeEventIDPayload.md) |  | [optional] 
**Notifications** | Pointer to [**[]NotificationSchemeNotificationDetailsPayload**](NotificationSchemeNotificationDetailsPayload.md) | The configuration for notification recipents | [optional] 

## Methods

### NewNotificationSchemeEventPayload

`func NewNotificationSchemeEventPayload() *NotificationSchemeEventPayload`

NewNotificationSchemeEventPayload instantiates a new NotificationSchemeEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeEventPayloadWithDefaults

`func NewNotificationSchemeEventPayloadWithDefaults() *NotificationSchemeEventPayload`

NewNotificationSchemeEventPayloadWithDefaults instantiates a new NotificationSchemeEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NotificationSchemeEventPayload) GetEvent() NotificationSchemeEventIDPayload`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationSchemeEventPayload) GetEventOk() (*NotificationSchemeEventIDPayload, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationSchemeEventPayload) SetEvent(v NotificationSchemeEventIDPayload)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *NotificationSchemeEventPayload) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetNotifications

`func (o *NotificationSchemeEventPayload) GetNotifications() []NotificationSchemeNotificationDetailsPayload`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationSchemeEventPayload) GetNotificationsOk() (*[]NotificationSchemeNotificationDetailsPayload, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationSchemeEventPayload) SetNotifications(v []NotificationSchemeNotificationDetailsPayload)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationSchemeEventPayload) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


