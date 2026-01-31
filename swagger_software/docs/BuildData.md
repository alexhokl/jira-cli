# BuildData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** | The schema version used for this data.  Placeholder to support potential schema changes in the future.  | [optional] [default to "1.0"]
**PipelineId** | **string** | An ID that relates a sequence of builds. Depending on your use case this might be a project ID, pipeline ID, plan key etc. - whatever logical unit you use to group a sequence of builds.  The combination of &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60; must uniquely identify a build you have provided.  | 
**BuildNumber** | **int64** | Identifies a build within the sequence of builds identified by the build &#x60;pipelineId&#x60;.  Used to identify the &#39;most recent&#39; build in that sequence of builds.  The combination of &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60; must uniquely identify a build you have provided.  | 
**UpdateSequenceNumber** | **int64** | A number used to apply an order to the updates to the build, as identified by &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60;, in the case of out-of-order receipt of update requests.  It must be a monotonically increasing number. For example, epoch time could be one way to generate the &#x60;updateSequenceNumber&#x60;.  Updates for a build that is received with an &#x60;updateSqeuenceNumber&#x60; less than or equal to what is currently stored will be ignored.  | 
**DisplayName** | **string** | The human-readable name for the build.  Will be shown in the UI.  | 
**Description** | Pointer to **string** | An optional description to attach to this build.  This may be anything that makes sense in your system.  | [optional] 
**Label** | Pointer to **string** | A human-readable string that to provide information about the build.  | [optional] 
**Url** | **string** | The URL to this build in your system.  | 
**State** | **string** | The state of a build.  * &#x60;pending&#x60; - The build is queued, or some manual action is required. * &#x60;in_progress&#x60; - The build is currently running. * &#x60;successful&#x60; - The build completed successfully. * &#x60;failed&#x60; - The build failed. * &#x60;cancelled&#x60; - The build has been cancelled or stopped. * &#x60;unknown&#x60; - The build is in an unknown state.  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user as a summary of the state of the build.  | 
**IssueKeys** | Pointer to **[]string** | The Jira issue keys to associate the build information with.  You are free to associate issue keys in any way you like. However, we recommend that you use the name of the branch the build was executed on, and extract issue keys from that name using a simple regex. This has the advantage that it provides an intuitive association of builds to issue keys.  | [optional] 
**Associations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | The Jira issue keys or IDs to associate the build with. | [optional] 
**TestInfo** | Pointer to [**TestInfo**](TestInfo.md) |  | [optional] 
**References** | Pointer to [**[]BuildReferences**](BuildReferences.md) | Optional information that links a build to a commit, branch etc.  | [optional] 

## Methods

### NewBuildData

`func NewBuildData(pipelineId string, buildNumber int64, updateSequenceNumber int64, displayName string, url string, state string, lastUpdated time.Time, ) *BuildData`

NewBuildData instantiates a new BuildData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildDataWithDefaults

`func NewBuildDataWithDefaults() *BuildData`

NewBuildDataWithDefaults instantiates a new BuildData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *BuildData) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *BuildData) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *BuildData) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *BuildData) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetPipelineId

`func (o *BuildData) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *BuildData) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *BuildData) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.


### GetBuildNumber

`func (o *BuildData) GetBuildNumber() int64`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *BuildData) GetBuildNumberOk() (*int64, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *BuildData) SetBuildNumber(v int64)`

SetBuildNumber sets BuildNumber field to given value.


### GetUpdateSequenceNumber

`func (o *BuildData) GetUpdateSequenceNumber() int64`

GetUpdateSequenceNumber returns the UpdateSequenceNumber field if non-nil, zero value otherwise.

### GetUpdateSequenceNumberOk

`func (o *BuildData) GetUpdateSequenceNumberOk() (*int64, bool)`

GetUpdateSequenceNumberOk returns a tuple with the UpdateSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSequenceNumber

`func (o *BuildData) SetUpdateSequenceNumber(v int64)`

SetUpdateSequenceNumber sets UpdateSequenceNumber field to given value.


### GetDisplayName

`func (o *BuildData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BuildData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BuildData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *BuildData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *BuildData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BuildData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BuildData) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BuildData) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUrl

`func (o *BuildData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BuildData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BuildData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetState

`func (o *BuildData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BuildData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BuildData) SetState(v string)`

SetState sets State field to given value.


### GetLastUpdated

`func (o *BuildData) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BuildData) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BuildData) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetIssueKeys

`func (o *BuildData) GetIssueKeys() []string`

GetIssueKeys returns the IssueKeys field if non-nil, zero value otherwise.

### GetIssueKeysOk

`func (o *BuildData) GetIssueKeysOk() (*[]string, bool)`

GetIssueKeysOk returns a tuple with the IssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKeys

`func (o *BuildData) SetIssueKeys(v []string)`

SetIssueKeys sets IssueKeys field to given value.

### HasIssueKeys

`func (o *BuildData) HasIssueKeys() bool`

HasIssueKeys returns a boolean if a field has been set.

### GetAssociations

`func (o *BuildData) GetAssociations() []IssueIdOrKeysAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *BuildData) GetAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *BuildData) SetAssociations(v []IssueIdOrKeysAssociation)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *BuildData) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetTestInfo

`func (o *BuildData) GetTestInfo() TestInfo`

GetTestInfo returns the TestInfo field if non-nil, zero value otherwise.

### GetTestInfoOk

`func (o *BuildData) GetTestInfoOk() (*TestInfo, bool)`

GetTestInfoOk returns a tuple with the TestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestInfo

`func (o *BuildData) SetTestInfo(v TestInfo)`

SetTestInfo sets TestInfo field to given value.

### HasTestInfo

`func (o *BuildData) HasTestInfo() bool`

HasTestInfo returns a boolean if a field has been set.

### GetReferences

`func (o *BuildData) GetReferences() []BuildReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BuildData) GetReferencesOk() (*[]BuildReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BuildData) SetReferences(v []BuildReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BuildData) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


