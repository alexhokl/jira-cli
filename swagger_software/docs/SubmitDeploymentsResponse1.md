# SubmitDeploymentsResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentSequenceNumber** | Pointer to **int64** | This is the identifier for the Deployment.  | [optional] 
**PipelineId** | Pointer to **string** | The ID of the Deployment&#39;s pipeline.  | [optional] 
**EnvironmentId** | Pointer to **string** | The ID of the Deployment&#39;s environment.  | [optional] 
**UpdatedTimestamp** | Pointer to **time.Time** | Time the deployment gating status was updated.  | [optional] 
**GatingStatus** | Pointer to **string** | The gating status  | [optional] 
**Details** | Pointer to [**[]SubmitDeploymentsResponse1DetailsInner**](SubmitDeploymentsResponse1DetailsInner.md) |  | [optional] 

## Methods

### NewSubmitDeploymentsResponse1

`func NewSubmitDeploymentsResponse1() *SubmitDeploymentsResponse1`

NewSubmitDeploymentsResponse1 instantiates a new SubmitDeploymentsResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDeploymentsResponse1WithDefaults

`func NewSubmitDeploymentsResponse1WithDefaults() *SubmitDeploymentsResponse1`

NewSubmitDeploymentsResponse1WithDefaults instantiates a new SubmitDeploymentsResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentSequenceNumber

`func (o *SubmitDeploymentsResponse1) GetDeploymentSequenceNumber() int64`

GetDeploymentSequenceNumber returns the DeploymentSequenceNumber field if non-nil, zero value otherwise.

### GetDeploymentSequenceNumberOk

`func (o *SubmitDeploymentsResponse1) GetDeploymentSequenceNumberOk() (*int64, bool)`

GetDeploymentSequenceNumberOk returns a tuple with the DeploymentSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSequenceNumber

`func (o *SubmitDeploymentsResponse1) SetDeploymentSequenceNumber(v int64)`

SetDeploymentSequenceNumber sets DeploymentSequenceNumber field to given value.

### HasDeploymentSequenceNumber

`func (o *SubmitDeploymentsResponse1) HasDeploymentSequenceNumber() bool`

HasDeploymentSequenceNumber returns a boolean if a field has been set.

### GetPipelineId

`func (o *SubmitDeploymentsResponse1) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *SubmitDeploymentsResponse1) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *SubmitDeploymentsResponse1) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *SubmitDeploymentsResponse1) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SubmitDeploymentsResponse1) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SubmitDeploymentsResponse1) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SubmitDeploymentsResponse1) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SubmitDeploymentsResponse1) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SubmitDeploymentsResponse1) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SubmitDeploymentsResponse1) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SubmitDeploymentsResponse1) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *SubmitDeploymentsResponse1) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetGatingStatus

`func (o *SubmitDeploymentsResponse1) GetGatingStatus() string`

GetGatingStatus returns the GatingStatus field if non-nil, zero value otherwise.

### GetGatingStatusOk

`func (o *SubmitDeploymentsResponse1) GetGatingStatusOk() (*string, bool)`

GetGatingStatusOk returns a tuple with the GatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatingStatus

`func (o *SubmitDeploymentsResponse1) SetGatingStatus(v string)`

SetGatingStatus sets GatingStatus field to given value.

### HasGatingStatus

`func (o *SubmitDeploymentsResponse1) HasGatingStatus() bool`

HasGatingStatus returns a boolean if a field has been set.

### GetDetails

`func (o *SubmitDeploymentsResponse1) GetDetails() []SubmitDeploymentsResponse1DetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SubmitDeploymentsResponse1) GetDetailsOk() (*[]SubmitDeploymentsResponse1DetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SubmitDeploymentsResponse1) SetDetails(v []SubmitDeploymentsResponse1DetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SubmitDeploymentsResponse1) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


