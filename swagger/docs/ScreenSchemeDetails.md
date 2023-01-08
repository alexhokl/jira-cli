# ScreenSchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen scheme. The maximum length is 255 characters. | [optional] 
**Name** | **string** | The name of the screen scheme. The name must be unique. The maximum length is 255 characters. | 
**Screens** | [**ScreenTypes**](ScreenTypes.md) | The IDs of the screens for the screen types of the screen scheme. Only screens used in classic projects are accepted. | 

## Methods

### NewScreenSchemeDetails

`func NewScreenSchemeDetails(name string, screens ScreenTypes, ) *ScreenSchemeDetails`

NewScreenSchemeDetails instantiates a new ScreenSchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenSchemeDetailsWithDefaults

`func NewScreenSchemeDetailsWithDefaults() *ScreenSchemeDetails`

NewScreenSchemeDetailsWithDefaults instantiates a new ScreenSchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScreenSchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenSchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenSchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenSchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScreenSchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenSchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenSchemeDetails) SetName(v string)`

SetName sets Name field to given value.


### GetScreens

`func (o *ScreenSchemeDetails) GetScreens() ScreenTypes`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *ScreenSchemeDetails) GetScreensOk() (*ScreenTypes, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *ScreenSchemeDetails) SetScreens(v ScreenTypes)`

SetScreens sets Screens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


