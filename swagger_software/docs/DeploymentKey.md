# DeploymentKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineId** | **string** | The identifier of a pipeline, must be unique for the provider.  | 
**EnvironmentId** | **string** | The identifier of an environment, must be unique for the provider so that it can be shared across pipelines.  | 
**DeploymentSequenceNumber** | **int64** | This is the identifier for the deployment. It must be unique for the specified pipeline and environment. It must be a monotonically increasing number, as this is used to sequence the deployments.  | 

## Methods

### NewDeploymentKey

`func NewDeploymentKey(pipelineId string, environmentId string, deploymentSequenceNumber int64, ) *DeploymentKey`

NewDeploymentKey instantiates a new DeploymentKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentKeyWithDefaults

`func NewDeploymentKeyWithDefaults() *DeploymentKey`

NewDeploymentKeyWithDefaults instantiates a new DeploymentKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineId

`func (o *DeploymentKey) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *DeploymentKey) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *DeploymentKey) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.


### GetEnvironmentId

`func (o *DeploymentKey) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DeploymentKey) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DeploymentKey) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDeploymentSequenceNumber

`func (o *DeploymentKey) GetDeploymentSequenceNumber() int64`

GetDeploymentSequenceNumber returns the DeploymentSequenceNumber field if non-nil, zero value otherwise.

### GetDeploymentSequenceNumberOk

`func (o *DeploymentKey) GetDeploymentSequenceNumberOk() (*int64, bool)`

GetDeploymentSequenceNumberOk returns a tuple with the DeploymentSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSequenceNumber

`func (o *DeploymentKey) SetDeploymentSequenceNumber(v int64)`

SetDeploymentSequenceNumber sets DeploymentSequenceNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


