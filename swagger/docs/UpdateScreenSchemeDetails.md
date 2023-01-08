# UpdateScreenSchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen scheme. The maximum length is 255 characters. | [optional] 
**Name** | Pointer to **string** | The name of the screen scheme. The name must be unique. The maximum length is 255 characters. | [optional] 
**Screens** | Pointer to [**UpdateScreenTypes**](UpdateScreenTypes.md) | The IDs of the screens for the screen types of the screen scheme. Only screens used in classic projects are accepted. | [optional] 

## Methods

### NewUpdateScreenSchemeDetails

`func NewUpdateScreenSchemeDetails() *UpdateScreenSchemeDetails`

NewUpdateScreenSchemeDetails instantiates a new UpdateScreenSchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScreenSchemeDetailsWithDefaults

`func NewUpdateScreenSchemeDetailsWithDefaults() *UpdateScreenSchemeDetails`

NewUpdateScreenSchemeDetailsWithDefaults instantiates a new UpdateScreenSchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateScreenSchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateScreenSchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateScreenSchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateScreenSchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateScreenSchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateScreenSchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateScreenSchemeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateScreenSchemeDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScreens

`func (o *UpdateScreenSchemeDetails) GetScreens() UpdateScreenTypes`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *UpdateScreenSchemeDetails) GetScreensOk() (*UpdateScreenTypes, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *UpdateScreenSchemeDetails) SetScreens(v UpdateScreenTypes)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *UpdateScreenSchemeDetails) HasScreens() bool`

HasScreens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


