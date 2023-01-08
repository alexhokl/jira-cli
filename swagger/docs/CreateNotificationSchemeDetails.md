# CreateNotificationSchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the notification scheme. | [optional] 
**Name** | **string** | The name of the notification scheme. Must be unique (case-insensitive). | 
**NotificationSchemeEvents** | Pointer to [**[]NotificationSchemeEventDetails**](NotificationSchemeEventDetails.md) | The list of notifications which should be added to the notification scheme. | [optional] 

## Methods

### NewCreateNotificationSchemeDetails

`func NewCreateNotificationSchemeDetails(name string, ) *CreateNotificationSchemeDetails`

NewCreateNotificationSchemeDetails instantiates a new CreateNotificationSchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationSchemeDetailsWithDefaults

`func NewCreateNotificationSchemeDetailsWithDefaults() *CreateNotificationSchemeDetails`

NewCreateNotificationSchemeDetailsWithDefaults instantiates a new CreateNotificationSchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateNotificationSchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNotificationSchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNotificationSchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNotificationSchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateNotificationSchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationSchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationSchemeDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationSchemeEvents

`func (o *CreateNotificationSchemeDetails) GetNotificationSchemeEvents() []NotificationSchemeEventDetails`

GetNotificationSchemeEvents returns the NotificationSchemeEvents field if non-nil, zero value otherwise.

### GetNotificationSchemeEventsOk

`func (o *CreateNotificationSchemeDetails) GetNotificationSchemeEventsOk() (*[]NotificationSchemeEventDetails, bool)`

GetNotificationSchemeEventsOk returns a tuple with the NotificationSchemeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSchemeEvents

`func (o *CreateNotificationSchemeDetails) SetNotificationSchemeEvents(v []NotificationSchemeEventDetails)`

SetNotificationSchemeEvents sets NotificationSchemeEvents field to given value.

### HasNotificationSchemeEvents

`func (o *CreateNotificationSchemeDetails) HasNotificationSchemeEvents() bool`

HasNotificationSchemeEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


