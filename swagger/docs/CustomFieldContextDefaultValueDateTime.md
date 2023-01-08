# CustomFieldContextDefaultValueDateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | Pointer to **string** | The default date-time in ISO format. Ignored if &#x60;useCurrent&#x60; is true. | [optional] 
**Type** | **string** |  | 
**UseCurrent** | Pointer to **bool** | Whether to use the current date. | [optional] [default to false]

## Methods

### NewCustomFieldContextDefaultValueDateTime

`func NewCustomFieldContextDefaultValueDateTime(type_ string, ) *CustomFieldContextDefaultValueDateTime`

NewCustomFieldContextDefaultValueDateTime instantiates a new CustomFieldContextDefaultValueDateTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueDateTimeWithDefaults

`func NewCustomFieldContextDefaultValueDateTimeWithDefaults() *CustomFieldContextDefaultValueDateTime`

NewCustomFieldContextDefaultValueDateTimeWithDefaults instantiates a new CustomFieldContextDefaultValueDateTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTime

`func (o *CustomFieldContextDefaultValueDateTime) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *CustomFieldContextDefaultValueDateTime) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *CustomFieldContextDefaultValueDateTime) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *CustomFieldContextDefaultValueDateTime) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldContextDefaultValueDateTime) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValueDateTime) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValueDateTime) SetType(v string)`

SetType sets Type field to given value.


### GetUseCurrent

`func (o *CustomFieldContextDefaultValueDateTime) GetUseCurrent() bool`

GetUseCurrent returns the UseCurrent field if non-nil, zero value otherwise.

### GetUseCurrentOk

`func (o *CustomFieldContextDefaultValueDateTime) GetUseCurrentOk() (*bool, bool)`

GetUseCurrentOk returns a tuple with the UseCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrent

`func (o *CustomFieldContextDefaultValueDateTime) SetUseCurrent(v bool)`

SetUseCurrent sets UseCurrent field to given value.

### HasUseCurrent

`func (o *CustomFieldContextDefaultValueDateTime) HasUseCurrent() bool`

HasUseCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


