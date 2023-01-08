# CustomFieldOptionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether the option is disabled. | [optional] 
**OptionId** | Pointer to **string** | For cascading options, the ID of a parent option. | [optional] 
**Value** | **string** | The value of the custom field option. | 

## Methods

### NewCustomFieldOptionCreate

`func NewCustomFieldOptionCreate(value string, ) *CustomFieldOptionCreate`

NewCustomFieldOptionCreate instantiates a new CustomFieldOptionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldOptionCreateWithDefaults

`func NewCustomFieldOptionCreateWithDefaults() *CustomFieldOptionCreate`

NewCustomFieldOptionCreateWithDefaults instantiates a new CustomFieldOptionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *CustomFieldOptionCreate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CustomFieldOptionCreate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CustomFieldOptionCreate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CustomFieldOptionCreate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOptionId

`func (o *CustomFieldOptionCreate) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *CustomFieldOptionCreate) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *CustomFieldOptionCreate) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *CustomFieldOptionCreate) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetValue

`func (o *CustomFieldOptionCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldOptionCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldOptionCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


