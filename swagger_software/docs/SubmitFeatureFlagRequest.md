# SubmitFeatureFlagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to Feature Flag data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Note that these properties will never be returned with Feature Flag data. They are not intended for use as metadata to associate with a Feature Flag. Internally they are stored as a hash so that personal information etc. is never stored within Jira.  Properties are supplied as key/value pairs, a maximum of 5 properties can be supplied, and keys must not contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**Flags** | [**[]FeatureFlagData**](FeatureFlagData.md) | A list of Feature Flags to submit to Jira.  Each Feature Flag may be associated with 1 or more Jira issue keys, and will be associated with any properties included in this request.  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata1**](ProviderMetadata1.md) |  | [optional] 

## Methods

### NewSubmitFeatureFlagRequest

`func NewSubmitFeatureFlagRequest(flags []FeatureFlagData, ) *SubmitFeatureFlagRequest`

NewSubmitFeatureFlagRequest instantiates a new SubmitFeatureFlagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitFeatureFlagRequestWithDefaults

`func NewSubmitFeatureFlagRequestWithDefaults() *SubmitFeatureFlagRequest`

NewSubmitFeatureFlagRequestWithDefaults instantiates a new SubmitFeatureFlagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitFeatureFlagRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitFeatureFlagRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitFeatureFlagRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitFeatureFlagRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetFlags

`func (o *SubmitFeatureFlagRequest) GetFlags() []FeatureFlagData`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SubmitFeatureFlagRequest) GetFlagsOk() (*[]FeatureFlagData, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SubmitFeatureFlagRequest) SetFlags(v []FeatureFlagData)`

SetFlags sets Flags field to given value.


### GetProviderMetadata

`func (o *SubmitFeatureFlagRequest) GetProviderMetadata() ProviderMetadata1`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitFeatureFlagRequest) GetProviderMetadataOk() (*ProviderMetadata1, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitFeatureFlagRequest) SetProviderMetadata(v ProviderMetadata1)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitFeatureFlagRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


