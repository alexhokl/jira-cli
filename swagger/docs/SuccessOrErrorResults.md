# SuccessOrErrorResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**RemoveFieldParametersResultError**](RemoveFieldParametersResultError.md) |  | [optional] 
**FieldId** | Pointer to **string** |  | [optional] 
**SchemeId** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**WorkTypeIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewSuccessOrErrorResults

`func NewSuccessOrErrorResults() *SuccessOrErrorResults`

NewSuccessOrErrorResults instantiates a new SuccessOrErrorResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessOrErrorResultsWithDefaults

`func NewSuccessOrErrorResultsWithDefaults() *SuccessOrErrorResults`

NewSuccessOrErrorResultsWithDefaults instantiates a new SuccessOrErrorResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SuccessOrErrorResults) GetError() RemoveFieldParametersResultError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SuccessOrErrorResults) GetErrorOk() (*RemoveFieldParametersResultError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SuccessOrErrorResults) SetError(v RemoveFieldParametersResultError)`

SetError sets Error field to given value.

### HasError

`func (o *SuccessOrErrorResults) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFieldId

`func (o *SuccessOrErrorResults) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *SuccessOrErrorResults) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *SuccessOrErrorResults) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *SuccessOrErrorResults) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetSchemeId

`func (o *SuccessOrErrorResults) GetSchemeId() int64`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *SuccessOrErrorResults) GetSchemeIdOk() (*int64, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *SuccessOrErrorResults) SetSchemeId(v int64)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *SuccessOrErrorResults) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.

### GetSuccess

`func (o *SuccessOrErrorResults) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SuccessOrErrorResults) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SuccessOrErrorResults) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SuccessOrErrorResults) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWorkTypeIds

`func (o *SuccessOrErrorResults) GetWorkTypeIds() []int64`

GetWorkTypeIds returns the WorkTypeIds field if non-nil, zero value otherwise.

### GetWorkTypeIdsOk

`func (o *SuccessOrErrorResults) GetWorkTypeIdsOk() (*[]int64, bool)`

GetWorkTypeIdsOk returns a tuple with the WorkTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeIds

`func (o *SuccessOrErrorResults) SetWorkTypeIds(v []int64)`

SetWorkTypeIds sets WorkTypeIds field to given value.

### HasWorkTypeIds

`func (o *SuccessOrErrorResults) HasWorkTypeIds() bool`

HasWorkTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


