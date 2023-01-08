# NotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the event. | [optional] 
**Id** | Pointer to **int64** | The ID of the event. The event can be a [Jira system event](https://confluence.atlassian.com/x/8YdKLg#Creatinganotificationscheme-eventsEvents) or a [custom event](https://confluence.atlassian.com/x/AIlKLg). | [optional] 
**Name** | Pointer to **string** | The name of the event. | [optional] 
**TemplateEvent** | Pointer to [**NotificationEvent**](NotificationEvent.md) | The template of the event. Only custom events configured by Jira administrators have template. | [optional] 

## Methods

### NewNotificationEvent

`func NewNotificationEvent() *NotificationEvent`

NewNotificationEvent instantiates a new NotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventWithDefaults

`func NewNotificationEventWithDefaults() *NotificationEvent`

NewNotificationEventWithDefaults instantiates a new NotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NotificationEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *NotificationEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateEvent

`func (o *NotificationEvent) GetTemplateEvent() NotificationEvent`

GetTemplateEvent returns the TemplateEvent field if non-nil, zero value otherwise.

### GetTemplateEventOk

`func (o *NotificationEvent) GetTemplateEventOk() (*NotificationEvent, bool)`

GetTemplateEventOk returns a tuple with the TemplateEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEvent

`func (o *NotificationEvent) SetTemplateEvent(v NotificationEvent)`

SetTemplateEvent sets TemplateEvent field to given value.

### HasTemplateEvent

`func (o *NotificationEvent) HasTemplateEvent() bool`

HasTemplateEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


