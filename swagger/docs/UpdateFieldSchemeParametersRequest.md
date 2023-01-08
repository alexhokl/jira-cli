# UpdateFieldSchemeParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**FieldsSchemeItemParameter**](FieldsSchemeItemParameter.md) |  | [optional] 
**SchemeIds** | Pointer to **[]int64** | The list of field scheme IDs to update | [optional] 
**WorkTypeParameters** | Pointer to [**[]FieldsSchemeItemWorkTypeParameter**](FieldsSchemeItemWorkTypeParameter.md) | The list of work type-specific parameter overrides, may be empty if only default parameters are being updated | [optional] 

## Methods

### NewUpdateFieldSchemeParametersRequest

`func NewUpdateFieldSchemeParametersRequest() *UpdateFieldSchemeParametersRequest`

NewUpdateFieldSchemeParametersRequest instantiates a new UpdateFieldSchemeParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFieldSchemeParametersRequestWithDefaults

`func NewUpdateFieldSchemeParametersRequestWithDefaults() *UpdateFieldSchemeParametersRequest`

NewUpdateFieldSchemeParametersRequestWithDefaults instantiates a new UpdateFieldSchemeParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *UpdateFieldSchemeParametersRequest) GetParameters() FieldsSchemeItemParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateFieldSchemeParametersRequest) GetParametersOk() (*FieldsSchemeItemParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateFieldSchemeParametersRequest) SetParameters(v FieldsSchemeItemParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateFieldSchemeParametersRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSchemeIds

`func (o *UpdateFieldSchemeParametersRequest) GetSchemeIds() []int64`

GetSchemeIds returns the SchemeIds field if non-nil, zero value otherwise.

### GetSchemeIdsOk

`func (o *UpdateFieldSchemeParametersRequest) GetSchemeIdsOk() (*[]int64, bool)`

GetSchemeIdsOk returns a tuple with the SchemeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeIds

`func (o *UpdateFieldSchemeParametersRequest) SetSchemeIds(v []int64)`

SetSchemeIds sets SchemeIds field to given value.

### HasSchemeIds

`func (o *UpdateFieldSchemeParametersRequest) HasSchemeIds() bool`

HasSchemeIds returns a boolean if a field has been set.

### GetWorkTypeParameters

`func (o *UpdateFieldSchemeParametersRequest) GetWorkTypeParameters() []FieldsSchemeItemWorkTypeParameter`

GetWorkTypeParameters returns the WorkTypeParameters field if non-nil, zero value otherwise.

### GetWorkTypeParametersOk

`func (o *UpdateFieldSchemeParametersRequest) GetWorkTypeParametersOk() (*[]FieldsSchemeItemWorkTypeParameter, bool)`

GetWorkTypeParametersOk returns a tuple with the WorkTypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeParameters

`func (o *UpdateFieldSchemeParametersRequest) SetWorkTypeParameters(v []FieldsSchemeItemWorkTypeParameter)`

SetWorkTypeParameters sets WorkTypeParameters field to given value.

### HasWorkTypeParameters

`func (o *UpdateFieldSchemeParametersRequest) HasWorkTypeParameters() bool`

HasWorkTypeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


