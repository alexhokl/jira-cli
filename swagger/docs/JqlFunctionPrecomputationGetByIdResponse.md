# JqlFunctionPrecomputationGetByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotFoundPrecomputationIDs** | Pointer to **[]string** | List of precomputations that were not found. | [optional] [readonly] 
**Precomputations** | Pointer to [**[]JqlFunctionPrecomputationBean**](JqlFunctionPrecomputationBean.md) | The list of precomputations. | [optional] [readonly] 

## Methods

### NewJqlFunctionPrecomputationGetByIdResponse

`func NewJqlFunctionPrecomputationGetByIdResponse() *JqlFunctionPrecomputationGetByIdResponse`

NewJqlFunctionPrecomputationGetByIdResponse instantiates a new JqlFunctionPrecomputationGetByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlFunctionPrecomputationGetByIdResponseWithDefaults

`func NewJqlFunctionPrecomputationGetByIdResponseWithDefaults() *JqlFunctionPrecomputationGetByIdResponse`

NewJqlFunctionPrecomputationGetByIdResponseWithDefaults instantiates a new JqlFunctionPrecomputationGetByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationGetByIdResponse) GetNotFoundPrecomputationIDs() []string`

GetNotFoundPrecomputationIDs returns the NotFoundPrecomputationIDs field if non-nil, zero value otherwise.

### GetNotFoundPrecomputationIDsOk

`func (o *JqlFunctionPrecomputationGetByIdResponse) GetNotFoundPrecomputationIDsOk() (*[]string, bool)`

GetNotFoundPrecomputationIDsOk returns a tuple with the NotFoundPrecomputationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationGetByIdResponse) SetNotFoundPrecomputationIDs(v []string)`

SetNotFoundPrecomputationIDs sets NotFoundPrecomputationIDs field to given value.

### HasNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationGetByIdResponse) HasNotFoundPrecomputationIDs() bool`

HasNotFoundPrecomputationIDs returns a boolean if a field has been set.

### GetPrecomputations

`func (o *JqlFunctionPrecomputationGetByIdResponse) GetPrecomputations() []JqlFunctionPrecomputationBean`

GetPrecomputations returns the Precomputations field if non-nil, zero value otherwise.

### GetPrecomputationsOk

`func (o *JqlFunctionPrecomputationGetByIdResponse) GetPrecomputationsOk() (*[]JqlFunctionPrecomputationBean, bool)`

GetPrecomputationsOk returns a tuple with the Precomputations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecomputations

`func (o *JqlFunctionPrecomputationGetByIdResponse) SetPrecomputations(v []JqlFunctionPrecomputationBean)`

SetPrecomputations sets Precomputations field to given value.

### HasPrecomputations

`func (o *JqlFunctionPrecomputationGetByIdResponse) HasPrecomputations() bool`

HasPrecomputations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


