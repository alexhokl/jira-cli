# NotificationScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the notification scheme. | [optional] 
**Expand** | Pointer to **string** | Expand options that include additional notification scheme details in the response. | [optional] 
**Id** | Pointer to **int64** | The ID of the notification scheme. | [optional] 
**Name** | Pointer to **string** | The name of the notification scheme. | [optional] 
**NotificationSchemeEvents** | Pointer to [**[]NotificationSchemeEvent**](NotificationSchemeEvent.md) | The notification events and associated recipients. | [optional] 
**Projects** | Pointer to **[]int64** | The list of project IDs associated with the notification scheme. | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the notification scheme. | [optional] 
**Self** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationScheme

`func NewNotificationScheme() *NotificationScheme`

NewNotificationScheme instantiates a new NotificationScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSchemeWithDefaults

`func NewNotificationSchemeWithDefaults() *NotificationScheme`

NewNotificationSchemeWithDefaults instantiates a new NotificationScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NotificationScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpand

`func (o *NotificationScheme) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *NotificationScheme) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *NotificationScheme) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *NotificationScheme) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetId

`func (o *NotificationScheme) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationScheme) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationScheme) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationSchemeEvents

`func (o *NotificationScheme) GetNotificationSchemeEvents() []NotificationSchemeEvent`

GetNotificationSchemeEvents returns the NotificationSchemeEvents field if non-nil, zero value otherwise.

### GetNotificationSchemeEventsOk

`func (o *NotificationScheme) GetNotificationSchemeEventsOk() (*[]NotificationSchemeEvent, bool)`

GetNotificationSchemeEventsOk returns a tuple with the NotificationSchemeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSchemeEvents

`func (o *NotificationScheme) SetNotificationSchemeEvents(v []NotificationSchemeEvent)`

SetNotificationSchemeEvents sets NotificationSchemeEvents field to given value.

### HasNotificationSchemeEvents

`func (o *NotificationScheme) HasNotificationSchemeEvents() bool`

HasNotificationSchemeEvents returns a boolean if a field has been set.

### GetProjects

`func (o *NotificationScheme) GetProjects() []int64`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *NotificationScheme) GetProjectsOk() (*[]int64, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *NotificationScheme) SetProjects(v []int64)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *NotificationScheme) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetScope

`func (o *NotificationScheme) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NotificationScheme) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NotificationScheme) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *NotificationScheme) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *NotificationScheme) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NotificationScheme) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NotificationScheme) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *NotificationScheme) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


