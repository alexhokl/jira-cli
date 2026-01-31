# SubmitDevopsComponentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedComponents** | Pointer to **[]string** | The IDs of Components that have been accepted for submission.  A Component may be rejected if it was only associated with unknown project keys.  Note that a Component that isn&#39;t updated due to it&#39;s updateSequenceNumber being out of order is not considered a failed submission.  | [optional] 
**FailedComponents** | Pointer to [**map[string][]ErrorMessage1**](array.md) | Details of Components that have not been accepted for submission, usually due to a problem with the request data.  The object (if present) will be keyed by Component ID and include any errors associated with that Component that have prevented it being submitted.  | [optional] 
**UnknownProjectKeys** | Pointer to **[]string** | Project keys that are not known on this Jira instance (if any).  These may be invalid keys (e.g. &#x60;UTF&#x60; is sometimes incorrectly identified as a Jira project key), or they may be for projects that no longer exist.  If a Component has been associated with project keys other than those in this array it will still be stored against those valid keys. If a Component was only associated with project keys deemed to be invalid it won&#39;t be persisted.  | [optional] 

## Methods

### NewSubmitDevopsComponentsResponse

`func NewSubmitDevopsComponentsResponse() *SubmitDevopsComponentsResponse`

NewSubmitDevopsComponentsResponse instantiates a new SubmitDevopsComponentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDevopsComponentsResponseWithDefaults

`func NewSubmitDevopsComponentsResponseWithDefaults() *SubmitDevopsComponentsResponse`

NewSubmitDevopsComponentsResponseWithDefaults instantiates a new SubmitDevopsComponentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedComponents

`func (o *SubmitDevopsComponentsResponse) GetAcceptedComponents() []string`

GetAcceptedComponents returns the AcceptedComponents field if non-nil, zero value otherwise.

### GetAcceptedComponentsOk

`func (o *SubmitDevopsComponentsResponse) GetAcceptedComponentsOk() (*[]string, bool)`

GetAcceptedComponentsOk returns a tuple with the AcceptedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedComponents

`func (o *SubmitDevopsComponentsResponse) SetAcceptedComponents(v []string)`

SetAcceptedComponents sets AcceptedComponents field to given value.

### HasAcceptedComponents

`func (o *SubmitDevopsComponentsResponse) HasAcceptedComponents() bool`

HasAcceptedComponents returns a boolean if a field has been set.

### GetFailedComponents

`func (o *SubmitDevopsComponentsResponse) GetFailedComponents() map[string][]ErrorMessage1`

GetFailedComponents returns the FailedComponents field if non-nil, zero value otherwise.

### GetFailedComponentsOk

`func (o *SubmitDevopsComponentsResponse) GetFailedComponentsOk() (*map[string][]ErrorMessage1, bool)`

GetFailedComponentsOk returns a tuple with the FailedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedComponents

`func (o *SubmitDevopsComponentsResponse) SetFailedComponents(v map[string][]ErrorMessage1)`

SetFailedComponents sets FailedComponents field to given value.

### HasFailedComponents

`func (o *SubmitDevopsComponentsResponse) HasFailedComponents() bool`

HasFailedComponents returns a boolean if a field has been set.

### GetUnknownProjectKeys

`func (o *SubmitDevopsComponentsResponse) GetUnknownProjectKeys() []string`

GetUnknownProjectKeys returns the UnknownProjectKeys field if non-nil, zero value otherwise.

### GetUnknownProjectKeysOk

`func (o *SubmitDevopsComponentsResponse) GetUnknownProjectKeysOk() (*[]string, bool)`

GetUnknownProjectKeysOk returns a tuple with the UnknownProjectKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownProjectKeys

`func (o *SubmitDevopsComponentsResponse) SetUnknownProjectKeys(v []string)`

SetUnknownProjectKeys sets UnknownProjectKeys field to given value.

### HasUnknownProjectKeys

`func (o *SubmitDevopsComponentsResponse) HasUnknownProjectKeys() bool`

HasUnknownProjectKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


