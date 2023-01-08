# JqlFunctionPrecomputationUpdateErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | Pointer to **[]string** | The list of error messages produced by this operation. | [optional] [readonly] 
**NotFoundPrecomputationIDs** | Pointer to **[]string** | List of precomputations that were not found. | [optional] [readonly] 

## Methods

### NewJqlFunctionPrecomputationUpdateErrorResponse

`func NewJqlFunctionPrecomputationUpdateErrorResponse() *JqlFunctionPrecomputationUpdateErrorResponse`

NewJqlFunctionPrecomputationUpdateErrorResponse instantiates a new JqlFunctionPrecomputationUpdateErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlFunctionPrecomputationUpdateErrorResponseWithDefaults

`func NewJqlFunctionPrecomputationUpdateErrorResponseWithDefaults() *JqlFunctionPrecomputationUpdateErrorResponse`

NewJqlFunctionPrecomputationUpdateErrorResponseWithDefaults instantiates a new JqlFunctionPrecomputationUpdateErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) GetNotFoundPrecomputationIDs() []string`

GetNotFoundPrecomputationIDs returns the NotFoundPrecomputationIDs field if non-nil, zero value otherwise.

### GetNotFoundPrecomputationIDsOk

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) GetNotFoundPrecomputationIDsOk() (*[]string, bool)`

GetNotFoundPrecomputationIDsOk returns a tuple with the NotFoundPrecomputationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) SetNotFoundPrecomputationIDs(v []string)`

SetNotFoundPrecomputationIDs sets NotFoundPrecomputationIDs field to given value.

### HasNotFoundPrecomputationIDs

`func (o *JqlFunctionPrecomputationUpdateErrorResponse) HasNotFoundPrecomputationIDs() bool`

HasNotFoundPrecomputationIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


