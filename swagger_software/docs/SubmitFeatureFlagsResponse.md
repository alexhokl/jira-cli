# SubmitFeatureFlagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedFeatureFlags** | Pointer to **[]string** | The IDs of Feature Flags that have been accepted for submission.  A Feature Flag may be rejected if it was only associated with unknown issue keys.  Note that a Feature Flag that isn&#39;t updated due to it&#39;s updateSequenceId being out of order is not considered a failed submission.  | [optional] 
**FailedFeatureFlags** | Pointer to [**map[string][]ErrorMessage1**](array.md) | Details of Feature Flags that have not been accepted for submission, usually due to a problem with the request data.  The object (if present) will be keyed by Feature Flag ID and include any errors associated with that Feature Flag that have prevented it being submitted.  | [optional] 
**UnknownIssueKeys** | Pointer to **[]string** | Issue keys that are not known on this Jira instance (if any).   These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a Feature Flag has been associated with issue keys other than those in this array it will still be stored against those valid keys. If a Feature Flag was only associated with issue keys deemed to be invalid it won&#39;t be persisted.  | [optional] 
**UnknownAssociations** | Pointer to [**[]IssueIdOrKeysAssociation**](IssueIdOrKeysAssociation.md) | Associations that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF-8&#x60; is sometimes incorrectly identified as a Jira issue key), or they may be for projects that no longer exist.  If a feature flag has been associated with any other association other than those in this array it will still be stored against those valid associations. If a feature flag was only associated with the associations in this array, it is deemed to be invalid and it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitFeatureFlagsResponse

`func NewSubmitFeatureFlagsResponse() *SubmitFeatureFlagsResponse`

NewSubmitFeatureFlagsResponse instantiates a new SubmitFeatureFlagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitFeatureFlagsResponseWithDefaults

`func NewSubmitFeatureFlagsResponseWithDefaults() *SubmitFeatureFlagsResponse`

NewSubmitFeatureFlagsResponseWithDefaults instantiates a new SubmitFeatureFlagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) GetAcceptedFeatureFlags() []string`

GetAcceptedFeatureFlags returns the AcceptedFeatureFlags field if non-nil, zero value otherwise.

### GetAcceptedFeatureFlagsOk

`func (o *SubmitFeatureFlagsResponse) GetAcceptedFeatureFlagsOk() (*[]string, bool)`

GetAcceptedFeatureFlagsOk returns a tuple with the AcceptedFeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) SetAcceptedFeatureFlags(v []string)`

SetAcceptedFeatureFlags sets AcceptedFeatureFlags field to given value.

### HasAcceptedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) HasAcceptedFeatureFlags() bool`

HasAcceptedFeatureFlags returns a boolean if a field has been set.

### GetFailedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) GetFailedFeatureFlags() map[string][]ErrorMessage1`

GetFailedFeatureFlags returns the FailedFeatureFlags field if non-nil, zero value otherwise.

### GetFailedFeatureFlagsOk

`func (o *SubmitFeatureFlagsResponse) GetFailedFeatureFlagsOk() (*map[string][]ErrorMessage1, bool)`

GetFailedFeatureFlagsOk returns a tuple with the FailedFeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) SetFailedFeatureFlags(v map[string][]ErrorMessage1)`

SetFailedFeatureFlags sets FailedFeatureFlags field to given value.

### HasFailedFeatureFlags

`func (o *SubmitFeatureFlagsResponse) HasFailedFeatureFlags() bool`

HasFailedFeatureFlags returns a boolean if a field has been set.

### GetUnknownIssueKeys

`func (o *SubmitFeatureFlagsResponse) GetUnknownIssueKeys() []string`

GetUnknownIssueKeys returns the UnknownIssueKeys field if non-nil, zero value otherwise.

### GetUnknownIssueKeysOk

`func (o *SubmitFeatureFlagsResponse) GetUnknownIssueKeysOk() (*[]string, bool)`

GetUnknownIssueKeysOk returns a tuple with the UnknownIssueKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownIssueKeys

`func (o *SubmitFeatureFlagsResponse) SetUnknownIssueKeys(v []string)`

SetUnknownIssueKeys sets UnknownIssueKeys field to given value.

### HasUnknownIssueKeys

`func (o *SubmitFeatureFlagsResponse) HasUnknownIssueKeys() bool`

HasUnknownIssueKeys returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *SubmitFeatureFlagsResponse) GetUnknownAssociations() []IssueIdOrKeysAssociation`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *SubmitFeatureFlagsResponse) GetUnknownAssociationsOk() (*[]IssueIdOrKeysAssociation, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *SubmitFeatureFlagsResponse) SetUnknownAssociations(v []IssueIdOrKeysAssociation)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *SubmitFeatureFlagsResponse) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


