# CustomFieldContextDefaultValueSingleVersionPicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**VersionId** | **string** | The ID of the default version. | 
**VersionOrder** | Pointer to **string** | The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are &#x60;\&quot;releasedFirst\&quot;&#x60; and &#x60;\&quot;unreleasedFirst\&quot;&#x60;. | [optional] 

## Methods

### NewCustomFieldContextDefaultValueSingleVersionPicker

`func NewCustomFieldContextDefaultValueSingleVersionPicker(type_ string, versionId string, ) *CustomFieldContextDefaultValueSingleVersionPicker`

NewCustomFieldContextDefaultValueSingleVersionPicker instantiates a new CustomFieldContextDefaultValueSingleVersionPicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueSingleVersionPickerWithDefaults

`func NewCustomFieldContextDefaultValueSingleVersionPickerWithDefaults() *CustomFieldContextDefaultValueSingleVersionPicker`

NewCustomFieldContextDefaultValueSingleVersionPickerWithDefaults instantiates a new CustomFieldContextDefaultValueSingleVersionPicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) SetType(v string)`

SetType sets Type field to given value.


### GetVersionId

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionOrder

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetVersionOrder() string`

GetVersionOrder returns the VersionOrder field if non-nil, zero value otherwise.

### GetVersionOrderOk

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) GetVersionOrderOk() (*string, bool)`

GetVersionOrderOk returns a tuple with the VersionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrder

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) SetVersionOrder(v string)`

SetVersionOrder sets VersionOrder field to given value.

### HasVersionOrder

`func (o *CustomFieldContextDefaultValueSingleVersionPicker) HasVersionOrder() bool`

HasVersionOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


