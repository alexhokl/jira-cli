# RejectedBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**BuildKey**](BuildKey.md) |  | 
**Errors** | [**[]ErrorMessage**](ErrorMessage.md) | The error messages for the rejected build | 

## Methods

### NewRejectedBuild

`func NewRejectedBuild(key BuildKey, errors []ErrorMessage, ) *RejectedBuild`

NewRejectedBuild instantiates a new RejectedBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedBuildWithDefaults

`func NewRejectedBuildWithDefaults() *RejectedBuild`

NewRejectedBuildWithDefaults instantiates a new RejectedBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *RejectedBuild) GetKey() BuildKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RejectedBuild) GetKeyOk() (*BuildKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RejectedBuild) SetKey(v BuildKey)`

SetKey sets Key field to given value.


### GetErrors

`func (o *RejectedBuild) GetErrors() []ErrorMessage`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RejectedBuild) GetErrorsOk() (*[]ErrorMessage, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RejectedBuild) SetErrors(v []ErrorMessage)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


