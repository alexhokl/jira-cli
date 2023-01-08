# NotificationSchemeNotificationDetailsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | Pointer to **string** | The type of notification. | [optional] 
**Parameter** | Pointer to **string** | The parameter of the notification, should be eiither null if not required, or PCRI. | [optional] 

## Methods

### NewNotificationSchemeNotificationDetailsPayload

`func NewNotificationSchemeNotificationDetailsPayload() *NotificationSchemeNotificationDetailsPayload`

NewNotificationSchemeNotificationDetailsPayload instantiates a new NotificationSchemeNotificationDetailsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeNotificationDetailsPayloadWithDefaults

`func NewNotificationSchemeNotificationDetailsPayloadWithDefaults() *NotificationSchemeNotificationDetailsPayload`

NewNotificationSchemeNotificationDetailsPayloadWithDefaults instantiates a new NotificationSchemeNotificationDetailsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *NotificationSchemeNotificationDetailsPayload) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotificationSchemeNotificationDetailsPayload) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotificationSchemeNotificationDetailsPayload) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *NotificationSchemeNotificationDetailsPayload) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetParameter

`func (o *NotificationSchemeNotificationDetailsPayload) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *NotificationSchemeNotificationDetailsPayload) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *NotificationSchemeNotificationDetailsPayload) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *NotificationSchemeNotificationDetailsPayload) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


