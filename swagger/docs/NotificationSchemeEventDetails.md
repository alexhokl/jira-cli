# NotificationSchemeEventDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NotificationSchemeEventTypeId**](NotificationSchemeEventTypeId.md) | The ID of the event. | 
**Notifications** | [**[]NotificationSchemeNotificationDetails**](NotificationSchemeNotificationDetails.md) | The list of notifications mapped to a specified event. | 

## Methods

### NewNotificationSchemeEventDetails

`func NewNotificationSchemeEventDetails(event NotificationSchemeEventTypeId, notifications []NotificationSchemeNotificationDetails, ) *NotificationSchemeEventDetails`

NewNotificationSchemeEventDetails instantiates a new NotificationSchemeEventDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeEventDetailsWithDefaults

`func NewNotificationSchemeEventDetailsWithDefaults() *NotificationSchemeEventDetails`

NewNotificationSchemeEventDetailsWithDefaults instantiates a new NotificationSchemeEventDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NotificationSchemeEventDetails) GetEvent() NotificationSchemeEventTypeId`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NotificationSchemeEventDetails) GetEventOk() (*NotificationSchemeEventTypeId, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NotificationSchemeEventDetails) SetEvent(v NotificationSchemeEventTypeId)`

SetEvent sets Event field to given value.


### GetNotifications

`func (o *NotificationSchemeEventDetails) GetNotifications() []NotificationSchemeNotificationDetails`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationSchemeEventDetails) GetNotificationsOk() (*[]NotificationSchemeNotificationDetails, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationSchemeEventDetails) SetNotifications(v []NotificationSchemeNotificationDetails)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


