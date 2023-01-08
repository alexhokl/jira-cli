# CustomFieldOptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether the option is disabled. | [optional] 
**Id** | **string** | The ID of the custom field option. | 
**Value** | Pointer to **string** | The value of the custom field option. | [optional] 

## Methods

### NewCustomFieldOptionUpdate

`func NewCustomFieldOptionUpdate(id string, ) *CustomFieldOptionUpdate`

NewCustomFieldOptionUpdate instantiates a new CustomFieldOptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionUpdateWithDefaults

`func NewCustomFieldOptionUpdateWithDefaults() *CustomFieldOptionUpdate`

NewCustomFieldOptionUpdateWithDefaults instantiates a new CustomFieldOptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *CustomFieldOptionUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CustomFieldOptionUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CustomFieldOptionUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CustomFieldOptionUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *CustomFieldOptionUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldOptionUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldOptionUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *CustomFieldOptionUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldOptionUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldOptionUpdate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldOptionUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


