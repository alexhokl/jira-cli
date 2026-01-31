# SubmitBuildsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedBuilds** | Pointer to [**[]BuildKey**](BuildKey.md) | The keys of builds that have been accepted for submission. A build key is a composite key that consists of &#x60;pipelineId&#x60; and &#x60;buildNumber&#x60;.  A build may be rejected if it was only associated with unknown issue keys, or if the submitted data for that build does not match the required schema.  Note that a build that isn&#39;t updated due to it&#39;s &#x60;updateSequenceNumber&#x60; being out of order is not considered a failed submission.  | [optional] 
**RejectedBuilds** | Pointer to [**[]RejectedBuild**](RejectedBuild.md) | Details of builds that have not been accepted for submission.  A build may be rejected if it was only associated with unknown issue keys, or if the submitted data for the build does not match the required schema.  | [optional] 
**UnknownIssueKeys** | Pointer to **[]string** | Issue keys that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a build has been associated with issue keys other than those in this array it will still be stored against those valid keys. If a build was only associated with issue keys deemed to be invalid it won&#39;t be persisted.  | [optional] 
**UnknownAssociations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | Associations that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a build has been associated with any other association other than those in this array it will still be stored against those valid associations. If a build was only associated with the associations in this array, it is deemed to be invalid and it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitBuildsResponse

`func NewSubmitBuildsResponse() *SubmitBuildsResponse`

NewSubmitBuildsResponse instantiates a new SubmitBuildsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitBuildsResponseWithDefaults

`func NewSubmitBuildsResponseWithDefaults() *SubmitBuildsResponse`

NewSubmitBuildsResponseWithDefaults instantiates a new SubmitBuildsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedBuilds

`func (o *SubmitBuildsResponse) GetAcceptedBuilds() []BuildKey`

GetAcceptedBuilds returns the AcceptedBuilds field if non-nil, zero value otherwise.

### GetAcceptedBuildsOk

`func (o *SubmitBuildsResponse) GetAcceptedBuildsOk() (*[]BuildKey, bool)`

GetAcceptedBuildsOk returns a tuple with the AcceptedBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBuilds

`func (o *SubmitBuildsResponse) SetAcceptedBuilds(v []BuildKey)`

SetAcceptedBuilds sets AcceptedBuilds field to given value.

### HasAcceptedBuilds

`func (o *SubmitBuildsResponse) HasAcceptedBuilds() bool`

HasAcceptedBuilds returns a boolean if a field has been set.

### GetRejectedBuilds

`func (o *SubmitBuildsResponse) GetRejectedBuilds() []RejectedBuild`

GetRejectedBuilds returns the RejectedBuilds field if non-nil, zero value otherwise.

### GetRejectedBuildsOk

`func (o *SubmitBuildsResponse) GetRejectedBuildsOk() (*[]RejectedBuild, bool)`

GetRejectedBuildsOk returns a tuple with the RejectedBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBuilds

`func (o *SubmitBuildsResponse) SetRejectedBuilds(v []RejectedBuild)`

SetRejectedBuilds sets RejectedBuilds field to given value.

### HasRejectedBuilds

`func (o *SubmitBuildsResponse) HasRejectedBuilds() bool`

HasRejectedBuilds returns a boolean if a field has been set.

### GetUnknownIssueKeys

`func (o *SubmitBuildsResponse) GetUnknownIssueKeys() []string`

GetUnknownIssueKeys returns the UnknownIssueKeys field if non-nil, zero value otherwise.

### GetUnknownIssueKeysOk

`func (o *SubmitBuildsResponse) GetUnknownIssueKeysOk() (*[]string, bool)`

GetUnknownIssueKeysOk returns a tuple with the UnknownIssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownIssueKeys

`func (o *SubmitBuildsResponse) SetUnknownIssueKeys(v []string)`

SetUnknownIssueKeys sets UnknownIssueKeys field to given value.

### HasUnknownIssueKeys

`func (o *SubmitBuildsResponse) HasUnknownIssueKeys() bool`

HasUnknownIssueKeys returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *SubmitBuildsResponse) GetUnknownAssociations() []IssueIdOrKeysAssociation`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *SubmitBuildsResponse) GetUnknownAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *SubmitBuildsResponse) SetUnknownAssociations(v []IssueIdOrKeysAssociation)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *SubmitBuildsResponse) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


