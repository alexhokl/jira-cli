# FeatureFlagDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A URL users can use to link to this Feature Flag, in this environment.  | 
**LastUpdated** | **time.Time** | The last-updated timestamp for this Feature Flag, in this environment.  Expected format is an RFC3339 formatted string.  | 
**Environment** | [**EnvironmentDetails**](EnvironmentDetails.md) |  | 
**Status** | [**FeatureFlagStatus**](FeatureFlagStatus.md) |  | 

## Methods

### NewFeatureFlagDetails

`func NewFeatureFlagDetails(url string, lastUpdated time.Time, environment EnvironmentDetails, status FeatureFlagStatus, ) *FeatureFlagDetails`

NewFeatureFlagDetails instantiates a new FeatureFlagDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagDetailsWithDefaults

`func NewFeatureFlagDetailsWithDefaults() *FeatureFlagDetails`

NewFeatureFlagDetailsWithDefaults instantiates a new FeatureFlagDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FeatureFlagDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FeatureFlagDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FeatureFlagDetails) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLastUpdated

`func (o *FeatureFlagDetails) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FeatureFlagDetails) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FeatureFlagDetails) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetEnvironment

`func (o *FeatureFlagDetails) GetEnvironment() EnvironmentDetails`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FeatureFlagDetails) GetEnvironmentOk() (*EnvironmentDetails, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FeatureFlagDetails) SetEnvironment(v EnvironmentDetails)`

SetEnvironment sets Environment field to given value.


### GetStatus

`func (o *FeatureFlagDetails) GetStatus() FeatureFlagStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeatureFlagDetails) GetStatusOk() (*FeatureFlagStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeatureFlagDetails) SetStatus(v FeatureFlagStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


