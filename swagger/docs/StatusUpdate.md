# StatusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Id** | **string** | The ID of the status. | 
**Name** | **string** | The name of the status. | 
**StatusCategory** | **string** | The category of the status. | 

## Methods

### NewStatusUpdate

`func NewStatusUpdate(id string, name string, statusCategory string, ) *StatusUpdate`

NewStatusUpdate instantiates a new StatusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusUpdateWithDefaults

`func NewStatusUpdateWithDefaults() *StatusUpdate`

NewStatusUpdateWithDefaults instantiates a new StatusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StatusUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *StatusUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatusUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatusUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StatusUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetStatusCategory

`func (o *StatusUpdate) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *StatusUpdate) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *StatusUpdate) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


