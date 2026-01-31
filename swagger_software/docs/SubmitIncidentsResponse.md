# SubmitIncidentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedIncidents** | Pointer to **[]string** | The IDs of Incidents that have been accepted for submission.  A Incident may be rejected if it was only associated with unknown project keys.  Note that a Incident that isn&#39;t updated due to it&#39;s updateSequenceNumber being out of order is not considered a failed submission.  | [optional] 
**FailedIncidents** | Pointer to [**map[string][]ErrorMessage1**](array.md) | Details of Incidents that have not been accepted for submission, usually due to a problem with the request data.  The object (if present) will be keyed by Incident ID and include any errors associated with that Incident that have prevented it being submitted.  | [optional] 
**UnknownProjectKeys** | Pointer to **[]string** | Project keys that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF&#x60; is sometimes incorrectly identified as a Jira project key), or they may be for projects that no longer exist.  If a Incident has been associated with project keys other than those in this array it will still be stored against those valid keys. If a Incident was only associated with project keys deemed to be invalid it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitIncidentsResponse

`func NewSubmitIncidentsResponse() *SubmitIncidentsResponse`

NewSubmitIncidentsResponse instantiates a new SubmitIncidentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitIncidentsResponseWithDefaults

`func NewSubmitIncidentsResponseWithDefaults() *SubmitIncidentsResponse`

NewSubmitIncidentsResponseWithDefaults instantiates a new SubmitIncidentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedIncidents

`func (o *SubmitIncidentsResponse) GetAcceptedIncidents() []string`

GetAcceptedIncidents returns the AcceptedIncidents field if non-nil, zero value otherwise.

### GetAcceptedIncidentsOk

`func (o *SubmitIncidentsResponse) GetAcceptedIncidentsOk() (*[]string, bool)`

GetAcceptedIncidentsOk returns a tuple with the AcceptedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedIncidents

`func (o *SubmitIncidentsResponse) SetAcceptedIncidents(v []string)`

SetAcceptedIncidents sets AcceptedIncidents field to given value.

### HasAcceptedIncidents

`func (o *SubmitIncidentsResponse) HasAcceptedIncidents() bool`

HasAcceptedIncidents returns a boolean if a field has been set.

### GetFailedIncidents

`func (o *SubmitIncidentsResponse) GetFailedIncidents() map[string][]ErrorMessage1`

GetFailedIncidents returns the FailedIncidents field if non-nil, zero value otherwise.

### GetFailedIncidentsOk

`func (o *SubmitIncidentsResponse) GetFailedIncidentsOk() (*map[string][]ErrorMessage1, bool)`

GetFailedIncidentsOk returns a tuple with the FailedIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedIncidents

`func (o *SubmitIncidentsResponse) SetFailedIncidents(v map[string][]ErrorMessage1)`

SetFailedIncidents sets FailedIncidents field to given value.

### HasFailedIncidents

`func (o *SubmitIncidentsResponse) HasFailedIncidents() bool`

HasFailedIncidents returns a boolean if a field has been set.

### GetUnknownProjectKeys

`func (o *SubmitIncidentsResponse) GetUnknownProjectKeys() []string`

GetUnknownProjectKeys returns the UnknownProjectKeys field if non-nil, zero value otherwise.

### GetUnknownProjectKeysOk

`func (o *SubmitIncidentsResponse) GetUnknownProjectKeysOk() (*[]string, bool)`

GetUnknownProjectKeysOk returns a tuple with the UnknownProjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownProjectKeys

`func (o *SubmitIncidentsResponse) SetUnknownProjectKeys(v []string)`

SetUnknownProjectKeys sets UnknownProjectKeys field to given value.

### HasUnknownProjectKeys

`func (o *SubmitIncidentsResponse) HasUnknownProjectKeys() bool`

HasUnknownProjectKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


