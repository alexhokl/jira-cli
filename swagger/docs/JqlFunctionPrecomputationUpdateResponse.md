# JqlFunctionPrecomputationUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotFoundPrecomputationIDs** | Pointer to **[]string** | List of precomputations that were not found and skipped. Only returned if the request passed skipNotFoundPrecomputations&#x3D;true. | [optional] [readonly] 

## Methods

### NewJqlFunctionPrecomputationUpdateResponse

`func NewJqlFunctionPrecomputationUpdateResponse() *JqlFunctionPrecomputationUpdateResponse`

NewJqlFunctionPrecomputationUpdateResponse instantiates a new JqlFunctionPrecomputationUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlFunctionPrecomputationUpdateResponseWithDefaults

`func NewJqlFunctionPrecomputationUpdateResponseWithDefaults() *JqlFunctionPrecomputationUpdateResponse`

NewJqlFunctionPrecomputationUpdateResponseWithDefaults instantiates a new JqlFunctionPrecomputationUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateResponse) GetNotFoundPrecomputationIDs() []string`

GetNotFoundPrecomputationIDs returns the NotFoundPrecomputationIDs field if non-nil, zero value otherwise.

### GetNotFoundPrecomputationIDsOk

`func (o *JqlFunctionPrecomputationUpdateResponse) GetNotFoundPrecomputationIDsOk() (*[]string, bool)`

GetNotFoundPrecomputationIDsOk returns a tuple with the NotFoundPrecomputationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateResponse) SetNotFoundPrecomputationIDs(v []string)`

SetNotFoundPrecomputationIDs sets NotFoundPrecomputationIDs field to given value.

### HasNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateResponse) HasNotFoundPrecomputationIDs() bool`

HasNotFoundPrecomputationIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


