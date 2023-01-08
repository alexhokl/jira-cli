# CustomFieldContextDefaultValueMultipleVersionPicker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**VersionIds** | **[]string** | The IDs of the default versions. | 
**VersionOrder** | Pointer to **string** | The order the pickable versions are displayed in. If not provided, the released-first order is used. Available version orders are &#x60;\&quot;releasedFirst\&quot;&#x60; and &#x60;\&quot;unreleasedFirst\&quot;&#x60;. | [optional] 

## Methods

### NewCustomFieldContextDefaultValueMultipleVersionPicker

`func NewCustomFieldContextDefaultValueMultipleVersionPicker(type_ string, versionIds []string, ) *CustomFieldContextDefaultValueMultipleVersionPicker`

NewCustomFieldContextDefaultValueMultipleVersionPicker instantiates a new CustomFieldContextDefaultValueMultipleVersionPicker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueMultipleVersionPickerWithDefaults

`func NewCustomFieldContextDefaultValueMultipleVersionPickerWithDefaults() *CustomFieldContextDefaultValueMultipleVersionPicker`

NewCustomFieldContextDefaultValueMultipleVersionPickerWithDefaults instantiates a new CustomFieldContextDefaultValueMultipleVersionPicker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) SetType(v string)`

SetType sets Type field to given value.


### GetVersionIds

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetVersionIds() []string`

GetVersionIds returns the VersionIds field if non-nil, zero value otherwise.

### GetVersionIdsOk

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetVersionIdsOk() (*[]string, bool)`

GetVersionIdsOk returns a tuple with the VersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIds

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) SetVersionIds(v []string)`

SetVersionIds sets VersionIds field to given value.


### GetVersionOrder

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetVersionOrder() string`

GetVersionOrder returns the VersionOrder field if non-nil, zero value otherwise.

### GetVersionOrderOk

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) GetVersionOrderOk() (*string, bool)`

GetVersionOrderOk returns a tuple with the VersionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionOrder

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) SetVersionOrder(v string)`

SetVersionOrder sets VersionOrder field to given value.

### HasVersionOrder

`func (o *CustomFieldContextDefaultValueMultipleVersionPicker) HasVersionOrder() bool`

HasVersionOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


