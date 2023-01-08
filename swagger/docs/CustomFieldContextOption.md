# CustomFieldContextOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** | Whether the option is disabled. | 
**Id** | **string** | The ID of the custom field option. | 
**OptionId** | Pointer to **string** | For cascading options, the ID of the custom field option containing the cascading option. | [optional] 
**Value** | **string** | The value of the custom field option. | 

## Methods

### NewCustomFieldContextOption

`func NewCustomFieldContextOption(disabled bool, id string, value string, ) *CustomFieldContextOption`

NewCustomFieldContextOption instantiates a new CustomFieldContextOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextOptionWithDefaults

`func NewCustomFieldContextOptionWithDefaults() *CustomFieldContextOption`

NewCustomFieldContextOptionWithDefaults instantiates a new CustomFieldContextOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *CustomFieldContextOption) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CustomFieldContextOption) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CustomFieldContextOption) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetId

`func (o *CustomFieldContextOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldContextOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldContextOption) SetId(v string)`

SetId sets Id field to given value.


### GetOptionId

`func (o *CustomFieldContextOption) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *CustomFieldContextOption) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *CustomFieldContextOption) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *CustomFieldContextOption) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetValue

`func (o *CustomFieldContextOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldContextOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldContextOption) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


