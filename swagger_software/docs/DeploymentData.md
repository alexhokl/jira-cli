# DeploymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentSequenceNumber** | **int64** | This is the identifier for the deployment. It must be unique for the specified pipeline and environment. It must be a monotonically increasing number, as this is used to sequence the deployments.  | 
**UpdateSequenceNumber** | **int64** | A number used to apply an order to the updates to the deployment, as identified by the deploymentSequenceNumber, in the case of out-of-order receipt of update requests. It must be a monotonically increasing number. For example, epoch time could be one way to generate the updateSequenceNumber.  | 
**IssueKeys** | Pointer to **[]string** | Deprecated. The Jira issue keys to associate the Deployment information with. Should replace this field with the \&quot;associations\&quot; field to associate Deployment information with issueKeys or other types of associations.  | [optional] 
**Associations** | Pointer to [**[]DeploymentDataAssociationsInner**](DeploymentDataAssociationsInner.md) | The entities to associate the Deployment information with. | [optional] 
**DisplayName** | **string** | The human-readable name for the deployment. Will be shown in the UI.  | 
**Url** | **string** | A URL users can use to link to this deployment, in this environment.  | 
**Description** | **string** | A short description of the deployment  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user as a summary of the state of the deployment.  | 
**Label** | Pointer to **string** | An (optional) additional label that may be displayed with deployment information. Can be used to display version information etc. for the deployment.  | [optional] 
**Duration** | Pointer to **int64** | The duration of the deployment (in seconds).  | [optional] 
**State** | **string** | The state of the deployment  | 
**Pipeline** | [**Pipeline**](Pipeline.md) |  | 
**Environment** | [**Environment**](Environment.md) |  | 
**Commands** | Pointer to [**[]Command**](Command.md) | A list of commands to be actioned for this Deployment  | [optional] 
**SchemaVersion** | Pointer to **string** | The DeploymentData schema version used for this deployment data.  Placeholder to support potential schema changes in the future.  | [optional] [default to "1.0"]

## Methods

### NewDeploymentData

`func NewDeploymentData(deploymentSequenceNumber int64, updateSequenceNumber int64, displayName string, url string, description string, lastUpdated time.Time, state string, pipeline Pipeline, environment Environment, ) *DeploymentData`

NewDeploymentData instantiates a new DeploymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDataWithDefaults

`func NewDeploymentDataWithDefaults() *DeploymentData`

NewDeploymentDataWithDefaults instantiates a new DeploymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentSequenceNumber

`func (o *DeploymentData) GetDeploymentSequenceNumber() int64`

GetDeploymentSequenceNumber returns the DeploymentSequenceNumber field if non-nil, zero value otherwise.

### GetDeploymentSequenceNumberOk

`func (o *DeploymentData) GetDeploymentSequenceNumberOk() (*int64, bool)`

GetDeploymentSequenceNumberOk returns a tuple with the DeploymentSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSequenceNumber

`func (o *DeploymentData) SetDeploymentSequenceNumber(v int64)`

SetDeploymentSequenceNumber sets DeploymentSequenceNumber field to given value.


### GetUpdateSequenceNumber

`func (o *DeploymentData) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *DeploymentData) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *DeploymentData) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetIssueKeys

`func (o *DeploymentData) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *DeploymentData) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *DeploymentData) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *DeploymentData) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *DeploymentData) GetAssociations() []DeploymentDataAssociationsInner`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *DeploymentData) GetAssociationsOk() (*[]DeploymentDataAssociationsInner, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *DeploymentData) SetAssociations(v []DeploymentDataAssociationsInner)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *DeploymentData) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetDisplayName

`func (o *DeploymentData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeploymentData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeploymentData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUrl

`func (o *DeploymentData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeploymentData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeploymentData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *DeploymentData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeploymentData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeploymentData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastUpdated

`func (o *DeploymentData) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeploymentData) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeploymentData) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLabel

`func (o *DeploymentData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeploymentData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeploymentData) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeploymentData) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDuration

`func (o *DeploymentData) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeploymentData) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeploymentData) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DeploymentData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetState

`func (o *DeploymentData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeploymentData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeploymentData) SetState(v string)`

SetState sets State field to given value.


### GetPipeline

`func (o *DeploymentData) GetPipeline() Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DeploymentData) GetPipelineOk() (*Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DeploymentData) SetPipeline(v Pipeline)`

SetPipeline sets Pipeline field to given value.


### GetEnvironment

`func (o *DeploymentData) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeploymentData) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeploymentData) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetCommands

`func (o *DeploymentData) GetCommands() []Command`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *DeploymentData) GetCommandsOk() (*[]Command, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *DeploymentData) SetCommands(v []Command)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *DeploymentData) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DeploymentData) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeploymentData) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeploymentData) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DeploymentData) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


