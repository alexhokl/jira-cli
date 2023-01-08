# AddNotificationsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationSchemeEvents** | [**[]NotificationSchemeEventDetails**](NotificationSchemeEventDetails.md) | The list of notifications which should be added to the notification scheme. | 

## Methods

### NewAddNotificationsDetails

`func NewAddNotificationsDetails(notificationSchemeEvents []NotificationSchemeEventDetails, ) *AddNotificationsDetails`

NewAddNotificationsDetails instantiates a new AddNotificationsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNotificationsDetailsWithDefaults

`func NewAddNotificationsDetailsWithDefaults() *AddNotificationsDetails`

NewAddNotificationsDetailsWithDefaults instantiates a new AddNotificationsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationSchemeEvents

`func (o *AddNotificationsDetails) GetNotificationSchemeEvents() []NotificationSchemeEventDetails`

GetNotificationSchemeEvents returns the NotificationSchemeEvents field if non-nil, zero value otherwise.

### GetNotificationSchemeEventsOk

`func (o *AddNotificationsDetails) GetNotificationSchemeEventsOk() (*[]NotificationSchemeEventDetails, bool)`

GetNotificationSchemeEventsOk returns a tuple with the NotificationSchemeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSchemeEvents

`func (o *AddNotificationsDetails) SetNotificationSchemeEvents(v []NotificationSchemeEventDetails)`

SetNotificationSchemeEvents sets NotificationSchemeEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


