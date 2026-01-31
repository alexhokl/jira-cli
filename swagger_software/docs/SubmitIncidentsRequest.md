# SubmitIncidentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** | Properties assigned to incidents/components/review data that can then be used for delete / query operations.  Examples might be an account or user ID that can then be used to clean up data if an account is removed from the Provider system.  Properties are supplied as key/value pairs, and a maximum of 5 properties can be supplied, keys cannot contain &#39;:&#39; or start with &#39;_&#39;.  | [optional] 
**ProviderMetadata** | Pointer to [**ProviderMetadata6**](ProviderMetadata6.md) |  | [optional] 
**Incidents** | Pointer to [**[]Incident**](Incident.md) |  | [optional] 
**Reviews** | Pointer to [**[]Review**](Review.md) |  | [optional] 

## Methods

### NewSubmitIncidentsRequest

`func NewSubmitIncidentsRequest() *SubmitIncidentsRequest`

NewSubmitIncidentsRequest instantiates a new SubmitIncidentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitIncidentsRequestWithDefaults

`func NewSubmitIncidentsRequestWithDefaults() *SubmitIncidentsRequest`

NewSubmitIncidentsRequestWithDefaults instantiates a new SubmitIncidentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SubmitIncidentsRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SubmitIncidentsRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SubmitIncidentsRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SubmitIncidentsRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetProviderMetadata

`func (o *SubmitIncidentsRequest) GetProviderMetadata() ProviderMetadata6`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *SubmitIncidentsRequest) GetProviderMetadataOk() (*ProviderMetadata6, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *SubmitIncidentsRequest) SetProviderMetadata(v ProviderMetadata6)`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *SubmitIncidentsRequest) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetIncidents

`func (o *SubmitIncidentsRequest) GetIncidents() []Incident`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *SubmitIncidentsRequest) GetIncidentsOk() (*[]Incident, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *SubmitIncidentsRequest) SetIncidents(v []Incident)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *SubmitIncidentsRequest) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetReviews

`func (o *SubmitIncidentsRequest) GetReviews() []Review`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *SubmitIncidentsRequest) GetReviewsOk() (*[]Review, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *SubmitIncidentsRequest) SetReviews(v []Review)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *SubmitIncidentsRequest) HasReviews() bool`

HasReviews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


