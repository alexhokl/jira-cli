# ParameterRemovalDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **[]string** | Set of parameter names to remove | [optional] 
**SchemeId** | Pointer to **int64** | ID of the field scheme | [optional] 
**WorkTypeIds** | Pointer to **[]int64** | Set of work type (issue type) IDs | [optional] 

## Methods

### NewParameterRemovalDetails

`func NewParameterRemovalDetails() *ParameterRemovalDetails`

NewParameterRemovalDetails instantiates a new ParameterRemovalDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterRemovalDetailsWithDefaults

`func NewParameterRemovalDetailsWithDefaults() *ParameterRemovalDetails`

NewParameterRemovalDetailsWithDefaults instantiates a new ParameterRemovalDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *ParameterRemovalDetails) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ParameterRemovalDetails) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ParameterRemovalDetails) SetParameters(v []string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ParameterRemovalDetails) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSchemeId

`func (o *ParameterRemovalDetails) GetSchemeId() int64`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *ParameterRemovalDetails) GetSchemeIdOk() (*int64, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *ParameterRemovalDetails) SetSchemeId(v int64)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *ParameterRemovalDetails) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.

### GetWorkTypeIds

`func (o *ParameterRemovalDetails) GetWorkTypeIds() []int64`

GetWorkTypeIds returns the WorkTypeIds field if non-nil, zero value otherwise.

### GetWorkTypeIdsOk

`func (o *ParameterRemovalDetails) GetWorkTypeIdsOk() (*[]int64, bool)`

GetWorkTypeIdsOk returns a tuple with the WorkTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeIds

`func (o *ParameterRemovalDetails) SetWorkTypeIds(v []int64)`

SetWorkTypeIds sets WorkTypeIds field to given value.

### HasWorkTypeIds

`func (o *ParameterRemovalDetails) HasWorkTypeIds() bool`

HasWorkTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


