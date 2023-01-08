# NotificationSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the notification scheme | [optional] 
**Name** | Pointer to **string** | The name of the notification scheme | [optional] 
**NotificationSchemeEvents** | Pointer to [**[]NotificationSchemeEventPayload**](NotificationSchemeEventPayload.md) | The events and notifications for the notification scheme | [optional] 
**OnConflict** | Pointer to **string** | The strategy to use when there is a conflict with an existing entity | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewNotificationSchemePayload

`func NewNotificationSchemePayload() *NotificationSchemePayload`

NewNotificationSchemePayload instantiates a new NotificationSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemePayloadWithDefaults

`func NewNotificationSchemePayloadWithDefaults() *NotificationSchemePayload`

NewNotificationSchemePayloadWithDefaults instantiates a new NotificationSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NotificationSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *NotificationSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationSchemeEvents

`func (o *NotificationSchemePayload) GetNotificationSchemeEvents() []NotificationSchemeEventPayload`

GetNotificationSchemeEvents returns the NotificationSchemeEvents field if non-nil, zero value otherwise.

### GetNotificationSchemeEventsOk

`func (o *NotificationSchemePayload) GetNotificationSchemeEventsOk() (*[]NotificationSchemeEventPayload, bool)`

GetNotificationSchemeEventsOk returns a tuple with the NotificationSchemeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSchemeEvents

`func (o *NotificationSchemePayload) SetNotificationSchemeEvents(v []NotificationSchemeEventPayload)`

SetNotificationSchemeEvents sets NotificationSchemeEvents field to given value.

### HasNotificationSchemeEvents

`func (o *NotificationSchemePayload) HasNotificationSchemeEvents() bool`

HasNotificationSchemeEvents returns a boolean if a field has been set.

### GetOnConflict

`func (o *NotificationSchemePayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *NotificationSchemePayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *NotificationSchemePayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *NotificationSchemePayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *NotificationSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *NotificationSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *NotificationSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *NotificationSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


