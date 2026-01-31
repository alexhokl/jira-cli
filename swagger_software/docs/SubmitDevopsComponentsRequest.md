# SubmitDevopsComponentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to incidents/components/review data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Properties are supplied as key/value pairs, and a maximum of 5 properties can be supplied, keys cannot contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**DevopsComponents** | [**[]Component**](Component.md) |  | 
**ProviderMetadata** | Pointer to [**ProviderMetadata6**](ProviderMetadata6.md) |  | [optional] 

## Methods

### NewSubmitDevopsComponentsRequest

`func NewSubmitDevopsComponentsRequest(devopsComponents []Component, ) *SubmitDevopsComponentsRequest`

NewSubmitDevopsComponentsRequest instantiates a new SubmitDevopsComponentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitDevopsComponentsRequestWithDefaults

`func NewSubmitDevopsComponentsRequestWithDefaults() *SubmitDevopsComponentsRequest`

NewSubmitDevopsComponentsRequestWithDefaults instantiates a new SubmitDevopsComponentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitDevopsComponentsRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitDevopsComponentsRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitDevopsComponentsRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitDevopsComponentsRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDevopsComponents

`func (o *SubmitDevopsComponentsRequest) GetDevopsComponents() []Component`

GetDevopsComponents returns the DevopsComponents field if non-nil, zero value otherwise.

### GetDevopsComponentsOk

`func (o *SubmitDevopsComponentsRequest) GetDevopsComponentsOk() (*[]Component, bool)`

GetDevopsComponentsOk returns a tuple with the DevopsComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevopsComponents

`func (o *SubmitDevopsComponentsRequest) SetDevopsComponents(v []Component)`

SetDevopsComponents sets DevopsComponents field to given value.


### GetProviderMetadata

`func (o *SubmitDevopsComponentsRequest) GetProviderMetadata() ProviderMetadata6`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitDevopsComponentsRequest) GetProviderMetadataOk() (*ProviderMetadata6, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitDevopsComponentsRequest) SetProviderMetadata(v ProviderMetadata6)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitDevopsComponentsRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


