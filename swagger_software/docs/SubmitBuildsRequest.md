# SubmitBuildsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to build data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Note that these properties will never be returned with build data. They are not intended for use as metadata to associate with a build. Internally they are stored as a hash so that personal information etc. is never stored within Jira.  Properties are supplied as key/value pairs, a maximum of 5 properties can be supplied, and keys must not contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**Builds** | [**[]BuildData**](BuildData.md) | A list of builds to submit to Jira.  Each build may be associated with one or more Jira issue keys, and will be associated with any properties included in this request.  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata3**](ProviderMetadata3.md) |  | [optional] 

## Methods

### NewSubmitBuildsRequest

`func NewSubmitBuildsRequest(builds []BuildData, ) *SubmitBuildsRequest`

NewSubmitBuildsRequest instantiates a new SubmitBuildsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitBuildsRequestWithDefaults

`func NewSubmitBuildsRequestWithDefaults() *SubmitBuildsRequest`

NewSubmitBuildsRequestWithDefaults instantiates a new SubmitBuildsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitBuildsRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitBuildsRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitBuildsRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitBuildsRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetBuilds

`func (o *SubmitBuildsRequest) GetBuilds() []BuildData`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *SubmitBuildsRequest) GetBuildsOk() (*[]BuildData, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *SubmitBuildsRequest) SetBuilds(v []BuildData)`

SetBuilds sets Builds field to given value.


### GetProviderMetadata

`func (o *SubmitBuildsRequest) GetProviderMetadata() ProviderMetadata3`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitBuildsRequest) GetProviderMetadataOk() (*ProviderMetadata3, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitBuildsRequest) SetProviderMetadata(v ProviderMetadata3)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitBuildsRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


