# JQLReferenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JqlReservedWords** | Pointer to **[]string** | List of JQL query reserved words. | [optional] 
**VisibleFieldNames** | Pointer to [**[]FieldReferenceData**](FieldReferenceData.md) | List of fields usable in JQL queries. | [optional] 
**VisibleFunctionNames** | Pointer to [**[]FunctionReferenceData**](FunctionReferenceData.md) | List of functions usable in JQL queries. | [optional] 

## Methods

### NewJQLReferenceData

`func NewJQLReferenceData() *JQLReferenceData`

NewJQLReferenceData instantiates a new JQLReferenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJQLReferenceDataWithDefaults

`func NewJQLReferenceDataWithDefaults() *JQLReferenceData`

NewJQLReferenceDataWithDefaults instantiates a new JQLReferenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJqlReservedWords

`func (o *JQLReferenceData) GetJqlReservedWords() []string`

GetJqlReservedWords returns the JqlReservedWords field if non-nil, zero value otherwise.

### GetJqlReservedWordsOk

`func (o *JQLReferenceData) GetJqlReservedWordsOk() (*[]string, bool)`

GetJqlReservedWordsOk returns a tuple with the JqlReservedWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqlReservedWords

`func (o *JQLReferenceData) SetJqlReservedWords(v []string)`

SetJqlReservedWords sets JqlReservedWords field to given value.

### HasJqlReservedWords

`func (o *JQLReferenceData) HasJqlReservedWords() bool`

HasJqlReservedWords returns a boolean if a field has been set.

### GetVisibleFieldNames

`func (o *JQLReferenceData) GetVisibleFieldNames() []FieldReferenceData`

GetVisibleFieldNames returns the VisibleFieldNames field if non-nil, zero value otherwise.

### GetVisibleFieldNamesOk

`func (o *JQLReferenceData) GetVisibleFieldNamesOk() (*[]FieldReferenceData, bool)`

GetVisibleFieldNamesOk returns a tuple with the VisibleFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleFieldNames

`func (o *JQLReferenceData) SetVisibleFieldNames(v []FieldReferenceData)`

SetVisibleFieldNames sets VisibleFieldNames field to given value.

### HasVisibleFieldNames

`func (o *JQLReferenceData) HasVisibleFieldNames() bool`

HasVisibleFieldNames returns a boolean if a field has been set.

### GetVisibleFunctionNames

`func (o *JQLReferenceData) GetVisibleFunctionNames() []FunctionReferenceData`

GetVisibleFunctionNames returns the VisibleFunctionNames field if non-nil, zero value otherwise.

### GetVisibleFunctionNamesOk

`func (o *JQLReferenceData) GetVisibleFunctionNamesOk() (*[]FunctionReferenceData, bool)`

GetVisibleFunctionNamesOk returns a tuple with the VisibleFunctionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleFunctionNames

`func (o *JQLReferenceData) SetVisibleFunctionNames(v []FunctionReferenceData)`

SetVisibleFunctionNames sets VisibleFunctionNames field to given value.

### HasVisibleFunctionNames

`func (o *JQLReferenceData) HasVisibleFunctionNames() bool`

HasVisibleFunctionNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


