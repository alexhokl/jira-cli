# FieldSchemeToFieldsPartialFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**FieldId** | **string** |  | 
**SchemeId** | **int64** |  | 
**Success** | **bool** |  | 
**WorkTypeIds** | **[]int64** |  | 

## Methods

### NewFieldSchemeToFieldsPartialFailure

`func NewFieldSchemeToFieldsPartialFailure(fieldId string, schemeId int64, success bool, workTypeIds []int64, ) *FieldSchemeToFieldsPartialFailure`

NewFieldSchemeToFieldsPartialFailure instantiates a new FieldSchemeToFieldsPartialFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSchemeToFieldsPartialFailureWithDefaults

`func NewFieldSchemeToFieldsPartialFailureWithDefaults() *FieldSchemeToFieldsPartialFailure`

NewFieldSchemeToFieldsPartialFailureWithDefaults instantiates a new FieldSchemeToFieldsPartialFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FieldSchemeToFieldsPartialFailure) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FieldSchemeToFieldsPartialFailure) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FieldSchemeToFieldsPartialFailure) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FieldSchemeToFieldsPartialFailure) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFieldId

`func (o *FieldSchemeToFieldsPartialFailure) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *FieldSchemeToFieldsPartialFailure) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *FieldSchemeToFieldsPartialFailure) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetSchemeId

`func (o *FieldSchemeToFieldsPartialFailure) GetSchemeId() int64`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *FieldSchemeToFieldsPartialFailure) GetSchemeIdOk() (*int64, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *FieldSchemeToFieldsPartialFailure) SetSchemeId(v int64)`

SetSchemeId sets SchemeId field to given value.


### GetSuccess

`func (o *FieldSchemeToFieldsPartialFailure) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FieldSchemeToFieldsPartialFailure) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FieldSchemeToFieldsPartialFailure) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetWorkTypeIds

`func (o *FieldSchemeToFieldsPartialFailure) GetWorkTypeIds() []int64`

GetWorkTypeIds returns the WorkTypeIds field if non-nil, zero value otherwise.

### GetWorkTypeIdsOk

`func (o *FieldSchemeToFieldsPartialFailure) GetWorkTypeIdsOk() (*[]int64, bool)`

GetWorkTypeIdsOk returns a tuple with the WorkTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeIds

`func (o *FieldSchemeToFieldsPartialFailure) SetWorkTypeIds(v []int64)`

SetWorkTypeIds sets WorkTypeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


