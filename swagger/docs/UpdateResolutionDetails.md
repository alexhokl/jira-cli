# UpdateResolutionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resolution. | [optional] 
**Name** | **string** | The name of the resolution. Must be unique. | 

## Methods

### NewUpdateResolutionDetails

`func NewUpdateResolutionDetails(name string, ) *UpdateResolutionDetails`

NewUpdateResolutionDetails instantiates a new UpdateResolutionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResolutionDetailsWithDefaults

`func NewUpdateResolutionDetailsWithDefaults() *UpdateResolutionDetails`

NewUpdateResolutionDetailsWithDefaults instantiates a new UpdateResolutionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateResolutionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateResolutionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateResolutionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateResolutionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateResolutionDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateResolutionDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateResolutionDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


