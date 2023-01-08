# JqlFunctionPrecomputationUpdateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | The error message to be displayed to the user if the given function clause is no longer valid during recalculation of the precomputation. | [optional] 
**Id** | **string** | The id of the precomputation to update. | 
**Value** | Pointer to **string** | The new value of the precomputation. | [optional] 

## Methods

### NewJqlFunctionPrecomputationUpdateBean

`func NewJqlFunctionPrecomputationUpdateBean(id string, ) *JqlFunctionPrecomputationUpdateBean`

NewJqlFunctionPrecomputationUpdateBean instantiates a new JqlFunctionPrecomputationUpdateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlFunctionPrecomputationUpdateBeanWithDefaults

`func NewJqlFunctionPrecomputationUpdateBeanWithDefaults() *JqlFunctionPrecomputationUpdateBean`

NewJqlFunctionPrecomputationUpdateBeanWithDefaults instantiates a new JqlFunctionPrecomputationUpdateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *JqlFunctionPrecomputationUpdateBean) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JqlFunctionPrecomputationUpdateBean) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JqlFunctionPrecomputationUpdateBean) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JqlFunctionPrecomputationUpdateBean) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *JqlFunctionPrecomputationUpdateBean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JqlFunctionPrecomputationUpdateBean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JqlFunctionPrecomputationUpdateBean) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *JqlFunctionPrecomputationUpdateBean) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JqlFunctionPrecomputationUpdateBean) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JqlFunctionPrecomputationUpdateBean) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *JqlFunctionPrecomputationUpdateBean) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


