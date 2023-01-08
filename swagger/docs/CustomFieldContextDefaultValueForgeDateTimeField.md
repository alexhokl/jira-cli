# CustomFieldContextDefaultValueForgeDateTimeField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextId** | **string** | The ID of the context. | 
**DateTime** | Pointer to **string** | The default date-time in ISO format. Ignored if &#x60;useCurrent&#x60; is true. | [optional] 
**Type** | **string** |  | 
**UseCurrent** | Pointer to **bool** | Whether to use the current date. | [optional] [default to false]

## Methods

### NewCustomFieldContextDefaultValueForgeDateTimeField

`func NewCustomFieldContextDefaultValueForgeDateTimeField(contextId string, type_ string, ) *CustomFieldContextDefaultValueForgeDateTimeField`

NewCustomFieldContextDefaultValueForgeDateTimeField instantiates a new CustomFieldContextDefaultValueForgeDateTimeField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextDefaultValueForgeDateTimeFieldWithDefaults

`func NewCustomFieldContextDefaultValueForgeDateTimeFieldWithDefaults() *CustomFieldContextDefaultValueForgeDateTimeField`

NewCustomFieldContextDefaultValueForgeDateTimeFieldWithDefaults instantiates a new CustomFieldContextDefaultValueForgeDateTimeField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextId

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetDateTime

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) SetType(v string)`

SetType sets Type field to given value.


### GetUseCurrent

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetUseCurrent() bool`

GetUseCurrent returns the UseCurrent field if non-nil, zero value otherwise.

### GetUseCurrentOk

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) GetUseCurrentOk() (*bool, bool)`

GetUseCurrentOk returns a tuple with the UseCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCurrent

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) SetUseCurrent(v bool)`

SetUseCurrent sets UseCurrent field to given value.

### HasUseCurrent

`func (o *CustomFieldContextDefaultValueForgeDateTimeField) HasUseCurrent() bool`

HasUseCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


