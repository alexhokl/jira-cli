# FieldsSchemeItemWorkTypeParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The custom description for the field for this work type, null to use default or preserve current | [optional] 
**IsRequired** | Pointer to **bool** | Whether the field is required for this work type, null to use default or preserve current | [optional] 
**WorkTypeId** | Pointer to **int64** | The ID of the work type (issue type) for which these parameters apply | [optional] 

## Methods

### NewFieldsSchemeItemWorkTypeParameter

`func NewFieldsSchemeItemWorkTypeParameter() *FieldsSchemeItemWorkTypeParameter`

NewFieldsSchemeItemWorkTypeParameter instantiates a new FieldsSchemeItemWorkTypeParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldsSchemeItemWorkTypeParameterWithDefaults

`func NewFieldsSchemeItemWorkTypeParameterWithDefaults() *FieldsSchemeItemWorkTypeParameter`

NewFieldsSchemeItemWorkTypeParameterWithDefaults instantiates a new FieldsSchemeItemWorkTypeParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FieldsSchemeItemWorkTypeParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldsSchemeItemWorkTypeParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldsSchemeItemWorkTypeParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldsSchemeItemWorkTypeParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsRequired

`func (o *FieldsSchemeItemWorkTypeParameter) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *FieldsSchemeItemWorkTypeParameter) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *FieldsSchemeItemWorkTypeParameter) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *FieldsSchemeItemWorkTypeParameter) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetWorkTypeId

`func (o *FieldsSchemeItemWorkTypeParameter) GetWorkTypeId() int64`

GetWorkTypeId returns the WorkTypeId field if non-nil, zero value otherwise.

### GetWorkTypeIdOk

`func (o *FieldsSchemeItemWorkTypeParameter) GetWorkTypeIdOk() (*int64, bool)`

GetWorkTypeIdOk returns a tuple with the WorkTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeId

`func (o *FieldsSchemeItemWorkTypeParameter) SetWorkTypeId(v int64)`

SetWorkTypeId sets WorkTypeId field to given value.

### HasWorkTypeId

`func (o *FieldsSchemeItemWorkTypeParameter) HasWorkTypeId() bool`

HasWorkTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


