# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HtmlBody** | Pointer to **string** | The HTML body of the email notification for the issue. | [optional] 
**Restrict** | Pointer to [**NotificationRecipientsRestrictions**](NotificationRecipientsRestrictions.md) | Restricts the notifications to users with the specified permissions. | [optional] 
**Subject** | Pointer to **string** | The subject of the email notification for the issue. If this is not specified, then the subject is set to the issue key and summary. | [optional] 
**TextBody** | Pointer to **string** | The plain text body of the email notification for the issue. | [optional] 
**To** | Pointer to [**NotificationRecipients**](NotificationRecipients.md) | The recipients of the email notification for the issue. | [optional] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtmlBody

`func (o *Notification) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *Notification) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *Notification) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *Notification) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetRestrict

`func (o *Notification) GetRestrict() NotificationRecipientsRestrictions`

GetRestrict returns the Restrict field if non-nil, zero value otherwise.

### GetRestrictOk

`func (o *Notification) GetRestrictOk() (*NotificationRecipientsRestrictions, bool)`

GetRestrictOk returns a tuple with the Restrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrict

`func (o *Notification) SetRestrict(v NotificationRecipientsRestrictions)`

SetRestrict sets Restrict field to given value.

### HasRestrict

`func (o *Notification) HasRestrict() bool`

HasRestrict returns a boolean if a field has been set.

### GetSubject

`func (o *Notification) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Notification) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Notification) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Notification) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTextBody

`func (o *Notification) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *Notification) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *Notification) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *Notification) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### GetTo

`func (o *Notification) GetTo() NotificationRecipients`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Notification) GetToOk() (*NotificationRecipients, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Notification) SetTo(v NotificationRecipients)`

SetTo sets To field to given value.

### HasTo

`func (o *Notification) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


