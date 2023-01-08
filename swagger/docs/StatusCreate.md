# StatusCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Name** | **string** | The name of the status. | 
**StatusCategory** | **string** | The category of the status. | 

## Methods

### NewStatusCreate

`func NewStatusCreate(name string, statusCategory string, ) *StatusCreate`

NewStatusCreate instantiates a new StatusCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCreateWithDefaults

`func NewStatusCreateWithDefaults() *StatusCreate`

NewStatusCreateWithDefaults instantiates a new StatusCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StatusCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StatusCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusCreate) SetName(v string)`

SetName sets Name field to given value.


### GetStatusCategory

`func (o *StatusCreate) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *StatusCreate) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *StatusCreate) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


