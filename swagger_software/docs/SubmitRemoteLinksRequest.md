# SubmitRemoteLinksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to Remote Link data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Properties are supplied as key/value pairs, a maximum of 5 properties can be supplied, and keys must not contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**RemoteLinks** | [**[]RemoteLinkData**](RemoteLinkData.md) | A list of Remote Links to submit to Jira.  Each Remote Link may be associated with one or more Jira issue keys, and will be associated with any properties included in this request.  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata4**](ProviderMetadata4.md) |  | [optional] 

## Methods

### NewSubmitRemoteLinksRequest

`func NewSubmitRemoteLinksRequest(remoteLinks []RemoteLinkData, ) *SubmitRemoteLinksRequest`

NewSubmitRemoteLinksRequest instantiates a new SubmitRemoteLinksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRemoteLinksRequestWithDefaults

`func NewSubmitRemoteLinksRequestWithDefaults() *SubmitRemoteLinksRequest`

NewSubmitRemoteLinksRequestWithDefaults instantiates a new SubmitRemoteLinksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitRemoteLinksRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitRemoteLinksRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitRemoteLinksRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitRemoteLinksRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRemoteLinks

`func (o *SubmitRemoteLinksRequest) GetRemoteLinks() []RemoteLinkData`

GetRemoteLinks returns the RemoteLinks field if non-nil, zero value otherwise.

### GetRemoteLinksOk

`func (o *SubmitRemoteLinksRequest) GetRemoteLinksOk() (*[]RemoteLinkData, bool)`

GetRemoteLinksOk returns a tuple with the RemoteLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLinks

`func (o *SubmitRemoteLinksRequest) SetRemoteLinks(v []RemoteLinkData)`

SetRemoteLinks sets RemoteLinks field to given value.


### GetProviderMetadata

`func (o *SubmitRemoteLinksRequest) GetProviderMetadata() ProviderMetadata4`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitRemoteLinksRequest) GetProviderMetadataOk() (*ProviderMetadata4, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitRemoteLinksRequest) SetProviderMetadata(v ProviderMetadata4)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitRemoteLinksRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


