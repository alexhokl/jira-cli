# CreateResolutionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resolution. | [optional] 
**Name** | **string** | The name of the resolution. Must be unique (case-insensitive). | 

## Methods

### NewCreateResolutionDetails

`func NewCreateResolutionDetails(name string, ) *CreateResolutionDetails`

NewCreateResolutionDetails instantiates a new CreateResolutionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResolutionDetailsWithDefaults

`func NewCreateResolutionDetailsWithDefaults() *CreateResolutionDetails`

NewCreateResolutionDetailsWithDefaults instantiates a new CreateResolutionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateResolutionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResolutionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResolutionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateResolutionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateResolutionDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateResolutionDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateResolutionDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


