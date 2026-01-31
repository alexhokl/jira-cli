# SubmitDeploymentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedDeployments** | Pointer to [**[]DeploymentKey**](DeploymentKey.md) | The keys of deployments that have been accepted for submission. A deployment key is a composite key that consists of &#x60;pipelineId&#x60;, &#x60;environmentId&#x60; and &#x60;deploymentSequenceNumber&#x60;.  A deployment may be rejected if it was only associated with unknown issue keys.  Note that a deployment that isn&#39;t updated due to it&#39;s updateSequenceNumber being out of order is not considered a failed submission.  | [optional] 
**RejectedDeployments** | Pointer to [**[]RejectedDeployment**](RejectedDeployment.md) | Details of deployments that have not been accepted for submission, usually due to a problem with the request data.  The object will contain the deployment key and any errors associated with that deployment that have prevented it being submitted.  | [optional] 
**UnknownIssueKeys** | Pointer to **[]string** | Issue keys that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a deployment has been associated with issue keys other than those in this array it will still be stored against those valid keys. If a deployment was only associated with issue keys deemed to be invalid it won&#39;t be persisted.  | [optional] 
**UnknownAssociations** | Pointer to [**[]DeploymentDataAssociationsInner**](DeploymentDataAssociationsInner.md) | Associations (e.g. Issue Keys or Service IDs) that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a deployment has been associated with any other association other than those in this array it will still be stored against those valid associations. If a deployment was only associated with the associations in this array, it is deemed to be invalid and it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitDeploymentsResponse

`func NewSubmitDeploymentsResponse() *SubmitDeploymentsResponse`

NewSubmitDeploymentsResponse instantiates a new SubmitDeploymentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDeploymentsResponseWithDefaults

`func NewSubmitDeploymentsResponseWithDefaults() *SubmitDeploymentsResponse`

NewSubmitDeploymentsResponseWithDefaults instantiates a new SubmitDeploymentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedDeployments

`func (o *SubmitDeploymentsResponse) GetAcceptedDeployments() []DeploymentKey`

GetAcceptedDeployments returns the AcceptedDeployments field if non-nil, zero value otherwise.

### GetAcceptedDeploymentsOk

`func (o *SubmitDeploymentsResponse) GetAcceptedDeploymentsOk() (*[]DeploymentKey, bool)`

GetAcceptedDeploymentsOk returns a tuple with the AcceptedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDeployments

`func (o *SubmitDeploymentsResponse) SetAcceptedDeployments(v []DeploymentKey)`

SetAcceptedDeployments sets AcceptedDeployments field to given value.

### HasAcceptedDeployments

`func (o *SubmitDeploymentsResponse) HasAcceptedDeployments() bool`

HasAcceptedDeployments returns a boolean if a field has been set.

### GetRejectedDeployments

`func (o *SubmitDeploymentsResponse) GetRejectedDeployments() []RejectedDeployment`

GetRejectedDeployments returns the RejectedDeployments field if non-nil, zero value otherwise.

### GetRejectedDeploymentsOk

`func (o *SubmitDeploymentsResponse) GetRejectedDeploymentsOk() (*[]RejectedDeployment, bool)`

GetRejectedDeploymentsOk returns a tuple with the RejectedDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedDeployments

`func (o *SubmitDeploymentsResponse) SetRejectedDeployments(v []RejectedDeployment)`

SetRejectedDeployments sets RejectedDeployments field to given value.

### HasRejectedDeployments

`func (o *SubmitDeploymentsResponse) HasRejectedDeployments() bool`

HasRejectedDeployments returns a boolean if a field has been set.

### GetUnknownIssueKeys

`func (o *SubmitDeploymentsResponse) GetUnknownIssueKeys() []string`

GetUnknownIssueKeys returns the UnknownIssueKeys field if non-nil, zero value otherwise.

### GetUnknownIssueKeysOk

`func (o *SubmitDeploymentsResponse) GetUnknownIssueKeysOk() (*[]string, bool)`

GetUnknownIssueKeysOk returns a tuple with the UnknownIssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownIssueKeys

`func (o *SubmitDeploymentsResponse) SetUnknownIssueKeys(v []string)`

SetUnknownIssueKeys sets UnknownIssueKeys field to given value.

### HasUnknownIssueKeys

`func (o *SubmitDeploymentsResponse) HasUnknownIssueKeys() bool`

HasUnknownIssueKeys returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *SubmitDeploymentsResponse) GetUnknownAssociations() []DeploymentDataAssociationsInner`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *SubmitDeploymentsResponse) GetUnknownAssociationsOk() (*[]DeploymentDataAssociationsInner, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *SubmitDeploymentsResponse) SetUnknownAssociations(v []DeploymentDataAssociationsInner)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *SubmitDeploymentsResponse) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


