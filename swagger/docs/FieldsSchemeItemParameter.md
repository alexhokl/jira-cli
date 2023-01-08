# FieldsSchemeItemParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The custom description for the field, null to preserve current description | [optional] 
**IsRequired** | Pointer to **bool** | Whether the field is required, null to preserve current requirement setting | [optional] 

## Methods

### NewFieldsSchemeItemParameter

`func NewFieldsSchemeItemParameter() *FieldsSchemeItemParameter`

NewFieldsSchemeItemParameter instantiates a new FieldsSchemeItemParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldsSchemeItemParameterWithDefaults

`func NewFieldsSchemeItemParameterWithDefaults() *FieldsSchemeItemParameter`

NewFieldsSchemeItemParameterWithDefaults instantiates a new FieldsSchemeItemParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FieldsSchemeItemParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldsSchemeItemParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldsSchemeItemParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldsSchemeItemParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsRequired

`func (o *FieldsSchemeItemParameter) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FieldsSchemeItemParameter) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FieldsSchemeItemParameter) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FieldsSchemeItemParameter) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


