# UiModificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contexts** | Pointer to [**[]UiModificationContextDetails**](UiModificationContextDetails.md) | List of contexts of the UI modification. The maximum number of contexts is 1000. | [optional] [readonly] 
**Data** | Pointer to **string** | The data of the UI modification. The maximum size of the data is 50000 characters. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the UI modification. The maximum length is 255 characters. | [optional] [readonly] 
**Id** | **string** | The ID of the UI modification. | [readonly] 
**Name** | **string** | The name of the UI modification. The maximum length is 255 characters. | [readonly] 
**Self** | **string** | The URL of the UI modification. | [readonly] 

## Methods

### NewUiModificationDetails

`func NewUiModificationDetails(id string, name string, self string, ) *UiModificationDetails`

NewUiModificationDetails instantiates a new UiModificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiModificationDetailsWithDefaults

`func NewUiModificationDetailsWithDefaults() *UiModificationDetails`

NewUiModificationDetailsWithDefaults instantiates a new UiModificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContexts

`func (o *UiModificationDetails) GetContexts() []UiModificationContextDetails`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *UiModificationDetails) GetContextsOk() (*[]UiModificationContextDetails, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *UiModificationDetails) SetContexts(v []UiModificationContextDetails)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *UiModificationDetails) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetData

`func (o *UiModificationDetails) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UiModificationDetails) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UiModificationDetails) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UiModificationDetails) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *UiModificationDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UiModificationDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UiModificationDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UiModificationDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UiModificationDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiModificationDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiModificationDetails) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UiModificationDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UiModificationDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UiModificationDetails) SetName(v string)`

SetName sets Name field to given value.


### GetSelf

`func (o *UiModificationDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UiModificationDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UiModificationDetails) SetSelf(v string)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


