# ScreenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen. The maximum length is 255 characters. | [optional] 
**Name** | **string** | The name of the screen. The name must be unique. The maximum length is 255 characters. | 

## Methods

### NewScreenDetails

`func NewScreenDetails(name string, ) *ScreenDetails`

NewScreenDetails instantiates a new ScreenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenDetailsWithDefaults

`func NewScreenDetailsWithDefaults() *ScreenDetails`

NewScreenDetailsWithDefaults instantiates a new ScreenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScreenDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScreenDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


