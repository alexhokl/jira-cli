# SubmitDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to deployment data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Properties are supplied as key/value pairs, and a maximum of 5 properties can be supplied, keys cannot contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**Deployments** | [**[]DeploymentData**](DeploymentData.md) | A list of deployments to submit to Jira.  Each deployment may be associated with one or more Jira issue keys, and will be associated with any properties included in this request.  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata2**](ProviderMetadata2.md) |  | [optional] 

## Methods

### NewSubmitDeploymentRequest

`func NewSubmitDeploymentRequest(deployments []DeploymentData, ) *SubmitDeploymentRequest`

NewSubmitDeploymentRequest instantiates a new SubmitDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDeploymentRequestWithDefaults

`func NewSubmitDeploymentRequestWithDefaults() *SubmitDeploymentRequest`

NewSubmitDeploymentRequestWithDefaults instantiates a new SubmitDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitDeploymentRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitDeploymentRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitDeploymentRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitDeploymentRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDeployments

`func (o *SubmitDeploymentRequest) GetDeployments() []DeploymentData`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *SubmitDeploymentRequest) GetDeploymentsOk() (*[]DeploymentData, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *SubmitDeploymentRequest) SetDeployments(v []DeploymentData)`

SetDeployments sets Deployments field to given value.


### GetProviderMetadata

`func (o *SubmitDeploymentRequest) GetProviderMetadata() ProviderMetadata2`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitDeploymentRequest) GetProviderMetadataOk() (*ProviderMetadata2, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitDeploymentRequest) SetProviderMetadata(v ProviderMetadata2)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitDeploymentRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


