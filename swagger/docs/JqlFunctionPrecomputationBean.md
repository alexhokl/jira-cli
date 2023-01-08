# JqlFunctionPrecomputationBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** | The list of arguments function was invoked with. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The timestamp of the precomputation creation. | [optional] [readonly] 
**Error** | Pointer to **string** | The error message to be displayed to the user. | [optional] [readonly] 
**Field** | Pointer to **string** | The field the function was executed against. | [optional] [readonly] 
**FunctionKey** | Pointer to **string** | The function key. | [optional] [readonly] 
**FunctionName** | Pointer to **string** | The name of the function. | [optional] [readonly] 
**Id** | Pointer to **string** | The id of the precomputation. | [optional] [readonly] 
**Operator** | Pointer to **string** | The operator in context of which function was executed. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The timestamp of the precomputation last update. | [optional] [readonly] 
**Used** | Pointer to **time.Time** | The timestamp of the precomputation last usage. | [optional] [readonly] 
**Value** | Pointer to **string** | The JQL fragment stored as the precomputation. | [optional] [readonly] 

## Methods

### NewJqlFunctionPrecomputationBean

`func NewJqlFunctionPrecomputationBean() *JqlFunctionPrecomputationBean`

NewJqlFunctionPrecomputationBean instantiates a new JqlFunctionPrecomputationBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlFunctionPrecomputationBeanWithDefaults

`func NewJqlFunctionPrecomputationBeanWithDefaults() *JqlFunctionPrecomputationBean`

NewJqlFunctionPrecomputationBeanWithDefaults instantiates a new JqlFunctionPrecomputationBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *JqlFunctionPrecomputationBean) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *JqlFunctionPrecomputationBean) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *JqlFunctionPrecomputationBean) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *JqlFunctionPrecomputationBean) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCreated

`func (o *JqlFunctionPrecomputationBean) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JqlFunctionPrecomputationBean) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JqlFunctionPrecomputationBean) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JqlFunctionPrecomputationBean) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetError

`func (o *JqlFunctionPrecomputationBean) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *JqlFunctionPrecomputationBean) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *JqlFunctionPrecomputationBean) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *JqlFunctionPrecomputationBean) HasError() bool`

HasError returns a boolean if a field has been set.

### GetField

`func (o *JqlFunctionPrecomputationBean) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *JqlFunctionPrecomputationBean) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *JqlFunctionPrecomputationBean) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *JqlFunctionPrecomputationBean) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFunctionKey

`func (o *JqlFunctionPrecomputationBean) GetFunctionKey() string`

GetFunctionKey returns the FunctionKey field if non-nil, zero value otherwise.

### GetFunctionKeyOk

`func (o *JqlFunctionPrecomputationBean) GetFunctionKeyOk() (*string, bool)`

GetFunctionKeyOk returns a tuple with the FunctionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionKey

`func (o *JqlFunctionPrecomputationBean) SetFunctionKey(v string)`

SetFunctionKey sets FunctionKey field to given value.

### HasFunctionKey

`func (o *JqlFunctionPrecomputationBean) HasFunctionKey() bool`

HasFunctionKey returns a boolean if a field has been set.

### GetFunctionName

`func (o *JqlFunctionPrecomputationBean) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *JqlFunctionPrecomputationBean) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *JqlFunctionPrecomputationBean) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *JqlFunctionPrecomputationBean) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetId

`func (o *JqlFunctionPrecomputationBean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JqlFunctionPrecomputationBean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JqlFunctionPrecomputationBean) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JqlFunctionPrecomputationBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperator

`func (o *JqlFunctionPrecomputationBean) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JqlFunctionPrecomputationBean) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JqlFunctionPrecomputationBean) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *JqlFunctionPrecomputationBean) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetUpdated

`func (o *JqlFunctionPrecomputationBean) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *JqlFunctionPrecomputationBean) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *JqlFunctionPrecomputationBean) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *JqlFunctionPrecomputationBean) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUsed

`func (o *JqlFunctionPrecomputationBean) GetUsed() time.Time`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *JqlFunctionPrecomputationBean) GetUsedOk() (*time.Time, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *JqlFunctionPrecomputationBean) SetUsed(v time.Time)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *JqlFunctionPrecomputationBean) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetValue

`func (o *JqlFunctionPrecomputationBean) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JqlFunctionPrecomputationBean) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JqlFunctionPrecomputationBean) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *JqlFunctionPrecomputationBean) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


