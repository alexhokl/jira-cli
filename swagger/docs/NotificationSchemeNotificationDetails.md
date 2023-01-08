# NotificationSchemeNotificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | **string** | The notification type, e.g &#x60;CurrentAssignee&#x60;, &#x60;Group&#x60;, &#x60;EmailAddress&#x60;. | 
**Parameter** | Pointer to **string** | The value corresponding to the specified notification type. | [optional] 

## Methods

### NewNotificationSchemeNotificationDetails

`func NewNotificationSchemeNotificationDetails(notificationType string, ) *NotificationSchemeNotificationDetails`

NewNotificationSchemeNotificationDetails instantiates a new NotificationSchemeNotificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeNotificationDetailsWithDefaults

`func NewNotificationSchemeNotificationDetailsWithDefaults() *NotificationSchemeNotificationDetails`

NewNotificationSchemeNotificationDetailsWithDefaults instantiates a new NotificationSchemeNotificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *NotificationSchemeNotificationDetails) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationSchemeNotificationDetails) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationSchemeNotificationDetails) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetParameter

`func (o *NotificationSchemeNotificationDetails) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *NotificationSchemeNotificationDetails) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *NotificationSchemeNotificationDetails) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *NotificationSchemeNotificationDetails) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


