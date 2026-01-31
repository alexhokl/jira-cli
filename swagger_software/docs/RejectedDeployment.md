# RejectedDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**DeploymentKey**](DeploymentKey.md) |  | 
**Errors** | [**[]ErrorMessage1**](ErrorMessage1.md) | The error messages for the rejected deployment | 

## Methods

### NewRejectedDeployment

`func NewRejectedDeployment(key DeploymentKey, errors []ErrorMessage1, ) *RejectedDeployment`

NewRejectedDeployment instantiates a new RejectedDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedDeploymentWithDefaults

`func NewRejectedDeploymentWithDefaults() *RejectedDeployment`

NewRejectedDeploymentWithDefaults instantiates a new RejectedDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *RejectedDeployment) GetKey() DeploymentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RejectedDeployment) GetKeyOk() (*DeploymentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RejectedDeployment) SetKey(v DeploymentKey)`

SetKey sets Key field to given value.


### GetErrors

`func (o *RejectedDeployment) GetErrors() []ErrorMessage1`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RejectedDeployment) GetErrorsOk() (*[]ErrorMessage1, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RejectedDeployment) SetErrors(v []ErrorMessage1)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


