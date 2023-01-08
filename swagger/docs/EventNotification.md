# EventNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address. | [optional] [default to null]
**Expand** | **string** | Expand options that include additional event notification details in the response. | [optional] [default to null]
**Field** | [***AllOfEventNotificationField**](AllOfEventNotificationField.md) | The custom user or group field. | [optional] [default to null]
**Group** | [***AllOfEventNotificationGroup**](AllOfEventNotificationGroup.md) | The specified group. | [optional] [default to null]
**Id** | **int64** | The ID of the notification. | [optional] [default to null]
**NotificationType** | **string** | Identifies the recipients of the notification. | [optional] [default to null]
**Parameter** | **string** | As a group&#x27;s name can change, use of &#x60;recipient&#x60; is recommended. The identifier associated with the &#x60;notificationType&#x60; value that defines the receiver of the notification, where the receiver isn&#x27;t implied by &#x60;notificationType&#x60; value. So, when &#x60;notificationType&#x60; is:   *  &#x60;User&#x60; The &#x60;parameter&#x60; is the user account ID.  *  &#x60;Group&#x60; The &#x60;parameter&#x60; is the group name.  *  &#x60;ProjectRole&#x60; The &#x60;parameter&#x60; is the project role ID.  *  &#x60;UserCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field.  *  &#x60;GroupCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field. | [optional] [default to null]
**ProjectRole** | [***AllOfEventNotificationProjectRole**](AllOfEventNotificationProjectRole.md) | The specified project role. | [optional] [default to null]
**Recipient** | **string** | The identifier associated with the &#x60;notificationType&#x60; value that defines the receiver of the notification, where the receiver isn&#x27;t implied by the &#x60;notificationType&#x60; value. So, when &#x60;notificationType&#x60; is:   *  &#x60;User&#x60;, &#x60;recipient&#x60; is the user account ID.  *  &#x60;Group&#x60;, &#x60;recipient&#x60; is the group ID.  *  &#x60;ProjectRole&#x60;, &#x60;recipient&#x60; is the project role ID.  *  &#x60;UserCustomField&#x60;, &#x60;recipient&#x60; is the ID of the custom field.  *  &#x60;GroupCustomField&#x60;, &#x60;recipient&#x60; is the ID of the custom field. | [optional] [default to null]
**User** | [***AllOfEventNotificationUser**](AllOfEventNotificationUser.md) | The specified user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

