# CustomFieldContextDefaultValueDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | The default date in ISO format. Ignored if &#x60;useCurrent&#x60; is true. | [optional] 
**Type** | **string** |  | 
**UseCurrent** | Pointer to **bool** | Whether to use the current date. | [optional] [default to false]

## Methods

### NewCustomFieldContextDefaultValueDate

`func NewCustomFieldContextDefaultValueDate(type_ string, ) *CustomFieldContextDefaultValueDate`

NewCustomFieldContextDefaultValueDate instantiates a new CustomFieldContextDefaultValueDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueDateWithDefaults

`func NewCustomFieldContextDefaultValueDateWithDefaults() *CustomFieldContextDefaultValueDate`

NewCustomFieldContextDefaultValueDateWithDefaults instantiates a new CustomFieldContextDefaultValueDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CustomFieldContextDefaultValueDate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CustomFieldContextDefaultValueDate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CustomFieldContextDefaultValueDate) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CustomFieldContextDefaultValueDate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldContextDefaultValueDate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValueDate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValueDate) SetType(v string)`

SetType sets Type field to given value.


### GetUseCurrent

`func (o *CustomFieldContextDefaultValueDate) GetUseCurrent() bool`

GetUseCurrent returns the UseCurrent field if non-nil, zero value otherwise.

### GetUseCurrentOk

`func (o *CustomFieldContextDefaultValueDate) GetUseCurrentOk() (*bool, bool)`

GetUseCurrentOk returns a tuple with the UseCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrent

`func (o *CustomFieldContextDefaultValueDate) SetUseCurrent(v bool)`

SetUseCurrent sets UseCurrent field to given value.

### HasUseCurrent

`func (o *CustomFieldContextDefaultValueDate) HasUseCurrent() bool`

HasUseCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


