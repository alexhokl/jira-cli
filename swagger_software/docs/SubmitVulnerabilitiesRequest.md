# SubmitVulnerabilitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to **string** | Indicates the operation being performed by the provider system when sending this data. \&quot;NORMAL\&quot; - Data received during real-time, user-triggered actions (e.g. user closed or updated a vulnerability). \&quot;SCAN\&quot; - Data sent through some automated process (e.g. some periodically scheduled repository scan). \&quot;BACKFILL\&quot; - Data received while backfilling existing data (e.g. pushing historical vulnerabilities when re-connect a workspace). Default is \&quot;NORMAL\&quot;. \&quot;NORMAL\&quot; traffic has higher priority but tighter rate limits, \&quot;SCAN\&quot; traffic has medium priority and looser limits, \&quot;BACKFILL\&quot; has lower priority and much looser limits  | [optional] 
**Properties** | Pointer to **map[string]string** | Properties assigned to vulnerability data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Properties are supplied as key/value pairs, and a maximum of 5 properties can be supplied, keys cannot contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**Vulnerabilities** | [**[]VulnerabilityDetails**](VulnerabilityDetails.md) |  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata5**](ProviderMetadata5.md) |  | [optional] 

## Methods

### NewSubmitVulnerabilitiesRequest

`func NewSubmitVulnerabilitiesRequest(vulnerabilities []VulnerabilityDetails, ) *SubmitVulnerabilitiesRequest`

NewSubmitVulnerabilitiesRequest instantiates a new SubmitVulnerabilitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitVulnerabilitiesRequestWithDefaults

`func NewSubmitVulnerabilitiesRequestWithDefaults() *SubmitVulnerabilitiesRequest`

NewSubmitVulnerabilitiesRequestWithDefaults instantiates a new SubmitVulnerabilitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *SubmitVulnerabilitiesRequest) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SubmitVulnerabilitiesRequest) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SubmitVulnerabilitiesRequest) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *SubmitVulnerabilitiesRequest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetProperties

`func (o *SubmitVulnerabilitiesRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitVulnerabilitiesRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitVulnerabilitiesRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitVulnerabilitiesRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *SubmitVulnerabilitiesRequest) GetVulnerabilities() []VulnerabilityDetails`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *SubmitVulnerabilitiesRequest) GetVulnerabilitiesOk() (*[]VulnerabilityDetails, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *SubmitVulnerabilitiesRequest) SetVulnerabilities(v []VulnerabilityDetails)`

SetVulnerabilities sets Vulnerabilities field to given value.


### GetProviderMetadata

`func (o *SubmitVulnerabilitiesRequest) GetProviderMetadata() ProviderMetadata5`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitVulnerabilitiesRequest) GetProviderMetadataOk() (*ProviderMetadata5, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitVulnerabilitiesRequest) SetProviderMetadata(v ProviderMetadata5)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitVulnerabilitiesRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


