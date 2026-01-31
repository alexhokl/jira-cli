# FeatureFlagSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | A URL users can use to link to a summary view of this flag, if appropriate.  This could be any location that makes sense in the Provider system (e.g. if the summary information comes from a specific environment, it might make sense to link the user to the flag in that environment).  | [optional] 
**Status** | [**FeatureFlagStatus**](FeatureFlagStatus.md) |  | 
**LastUpdated** | **time.Time** | The last-updated timestamp to present to the user as a summary of the state of the Feature Flag.  Providers may choose to supply the last-updated timestamp from a specific environment, or the &#39;most recent&#39; last-updated timestamp across all environments - whatever makes sense in the Provider system.  Expected format is an RFC3339 formatted string.  | 

## Methods

### NewFeatureFlagSummary

`func NewFeatureFlagSummary(status FeatureFlagStatus, lastUpdated time.Time, ) *FeatureFlagSummary`

NewFeatureFlagSummary instantiates a new FeatureFlagSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagSummaryWithDefaults

`func NewFeatureFlagSummaryWithDefaults() *FeatureFlagSummary`

NewFeatureFlagSummaryWithDefaults instantiates a new FeatureFlagSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FeatureFlagSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FeatureFlagSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FeatureFlagSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FeatureFlagSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FeatureFlagSummary) GetStatus() FeatureFlagStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeatureFlagSummary) GetStatusOk() (*FeatureFlagStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeatureFlagSummary) SetStatus(v FeatureFlagStatus)`

SetStatus sets Status field to given value.


### GetLastUpdated

`func (o *FeatureFlagSummary) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FeatureFlagSummary) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FeatureFlagSummary) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


