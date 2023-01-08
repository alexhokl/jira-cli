# FunctionReferenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the function. | [optional] 
**IsList** | Pointer to **string** | Whether the function can take a list of arguments. | [optional] 
**SupportsListAndSingleValueOperators** | Pointer to **string** | Whether the function supports both single and list value operators. | [optional] 
**Types** | Pointer to **[]string** | The data types returned by the function. | [optional] 
**Value** | Pointer to **string** | The function identifier. | [optional] 

## Methods

### NewFunctionReferenceData

`func NewFunctionReferenceData() *FunctionReferenceData`

NewFunctionReferenceData instantiates a new FunctionReferenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionReferenceDataWithDefaults

`func NewFunctionReferenceDataWithDefaults() *FunctionReferenceData`

NewFunctionReferenceDataWithDefaults instantiates a new FunctionReferenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FunctionReferenceData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionReferenceData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionReferenceData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionReferenceData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsList

`func (o *FunctionReferenceData) GetIsList() string`

GetIsList returns the IsList field if non-nil, zero value otherwise.

### GetIsListOk

`func (o *FunctionReferenceData) GetIsListOk() (*string, bool)`

GetIsListOk returns a tuple with the IsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsList

`func (o *FunctionReferenceData) SetIsList(v string)`

SetIsList sets IsList field to given value.

### HasIsList

`func (o *FunctionReferenceData) HasIsList() bool`

HasIsList returns a boolean if a field has been set.

### GetSupportsListAndSingleValueOperators

`func (o *FunctionReferenceData) GetSupportsListAndSingleValueOperators() string`

GetSupportsListAndSingleValueOperators returns the SupportsListAndSingleValueOperators field if non-nil, zero value otherwise.

### GetSupportsListAndSingleValueOperatorsOk

`func (o *FunctionReferenceData) GetSupportsListAndSingleValueOperatorsOk() (*string, bool)`

GetSupportsListAndSingleValueOperatorsOk returns a tuple with the SupportsListAndSingleValueOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsListAndSingleValueOperators

`func (o *FunctionReferenceData) SetSupportsListAndSingleValueOperators(v string)`

SetSupportsListAndSingleValueOperators sets SupportsListAndSingleValueOperators field to given value.

### HasSupportsListAndSingleValueOperators

`func (o *FunctionReferenceData) HasSupportsListAndSingleValueOperators() bool`

HasSupportsListAndSingleValueOperators returns a boolean if a field has been set.

### GetTypes

`func (o *FunctionReferenceData) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *FunctionReferenceData) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *FunctionReferenceData) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *FunctionReferenceData) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetValue

`func (o *FunctionReferenceData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FunctionReferenceData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FunctionReferenceData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FunctionReferenceData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


