# CreateUiModificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contexts** | Pointer to [**[]UiModificationContextDetails**](UiModificationContextDetails.md) | List of contexts of the UI modification. The maximum number of contexts is 1000. | [optional] 
**Data** | Pointer to **string** | The data of the UI modification. The maximum size of the data is 50000 characters. | [optional] 
**Description** | Pointer to **string** | The description of the UI modification. The maximum length is 255 characters. | [optional] 
**Name** | **string** | The name of the UI modification. The maximum length is 255 characters. | 

## Methods

### NewCreateUiModificationDetails

`func NewCreateUiModificationDetails(name string, ) *CreateUiModificationDetails`

NewCreateUiModificationDetails instantiates a new CreateUiModificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUiModificationDetailsWithDefaults

`func NewCreateUiModificationDetailsWithDefaults() *CreateUiModificationDetails`

NewCreateUiModificationDetailsWithDefaults instantiates a new CreateUiModificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContexts

`func (o *CreateUiModificationDetails) GetContexts() []UiModificationContextDetails`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *CreateUiModificationDetails) GetContextsOk() (*[]UiModificationContextDetails, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *CreateUiModificationDetails) SetContexts(v []UiModificationContextDetails)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *CreateUiModificationDetails) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetData

`func (o *CreateUiModificationDetails) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateUiModificationDetails) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateUiModificationDetails) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CreateUiModificationDetails) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUiModificationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUiModificationDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUiModificationDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUiModificationDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateUiModificationDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUiModificationDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUiModificationDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


