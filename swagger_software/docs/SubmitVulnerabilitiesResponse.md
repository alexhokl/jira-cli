# SubmitVulnerabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedVulnerabilities** | Pointer to **[]string** | The IDs of Vulnerabilities that have been accepted for submission.  A Vulnerability may be rejected if it was only associated with unknown project keys.  Note that a Vulnerability that isn&#39;t updated due to it&#39;s updateSequenceNumber being out of order is not considered a failed submission.  | [optional] 
**FailedVulnerabilities** | Pointer to [**map[string][]ErrorMessage1**](array.md) | Details of Vulnerabilities that have not been accepted for submission, usually due to a problem with the request data.  The object (if present) will be keyed by Vulnerability ID and include any errors associated with that Vulnerability that have prevented it being submitted.  | [optional] 
**UnknownAssociations** | Pointer to [**[]VulnerabilityDetailsAddAssociationsInner**](VulnerabilityDetailsAddAssociationsInner.md) | Associations (e.g. Service IDs) that are not known on this Jira instance (if any).  If a Vulnerability has been associated with any other association other than those in this array it will still be stored against those valid associations. If a Vulnerability was only associated with the associations in this array, it is deemed to be invalid and it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitVulnerabilitiesResponse

`func NewSubmitVulnerabilitiesResponse() *SubmitVulnerabilitiesResponse`

NewSubmitVulnerabilitiesResponse instantiates a new SubmitVulnerabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitVulnerabilitiesResponseWithDefaults

`func NewSubmitVulnerabilitiesResponseWithDefaults() *SubmitVulnerabilitiesResponse`

NewSubmitVulnerabilitiesResponseWithDefaults instantiates a new SubmitVulnerabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) GetAcceptedVulnerabilities() []string`

GetAcceptedVulnerabilities returns the AcceptedVulnerabilities field if non-nil, zero value otherwise.

### GetAcceptedVulnerabilitiesOk

`func (o *SubmitVulnerabilitiesResponse) GetAcceptedVulnerabilitiesOk() (*[]string, bool)`

GetAcceptedVulnerabilitiesOk returns a tuple with the AcceptedVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) SetAcceptedVulnerabilities(v []string)`

SetAcceptedVulnerabilities sets AcceptedVulnerabilities field to given value.

### HasAcceptedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) HasAcceptedVulnerabilities() bool`

HasAcceptedVulnerabilities returns a boolean if a field has been set.

### GetFailedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) GetFailedVulnerabilities() map[string][]ErrorMessage1`

GetFailedVulnerabilities returns the FailedVulnerabilities field if non-nil, zero value otherwise.

### GetFailedVulnerabilitiesOk

`func (o *SubmitVulnerabilitiesResponse) GetFailedVulnerabilitiesOk() (*map[string][]ErrorMessage1, bool)`

GetFailedVulnerabilitiesOk returns a tuple with the FailedVulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) SetFailedVulnerabilities(v map[string][]ErrorMessage1)`

SetFailedVulnerabilities sets FailedVulnerabilities field to given value.

### HasFailedVulnerabilities

`func (o *SubmitVulnerabilitiesResponse) HasFailedVulnerabilities() bool`

HasFailedVulnerabilities returns a boolean if a field has been set.

### GetUnknownAssociations

`func (o *SubmitVulnerabilitiesResponse) GetUnknownAssociations() []VulnerabilityDetailsAddAssociationsInner`

GetUnknownAssociations returns the UnknownAssociations field if non-nil, zero value otherwise.

### GetUnknownAssociationsOk

`func (o *SubmitVulnerabilitiesResponse) GetUnknownAssociationsOk() (*[]VulnerabilityDetailsAddAssociationsInner, bool)`

GetUnknownAssociationsOk returns a tuple with the UnknownAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownAssociations

`func (o *SubmitVulnerabilitiesResponse) SetUnknownAssociations(v []VulnerabilityDetailsAddAssociationsInner)`

SetUnknownAssociations sets UnknownAssociations field to given value.

### HasUnknownAssociations

`func (o *SubmitVulnerabilitiesResponse) HasUnknownAssociations() bool`

HasUnknownAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


