# FeatureFlagStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the Feature Flag is enabled in the given environment (or in summary).  Enabled may imply a partial rollout, which can be described using the &#39;rollout&#39; field.  | 
**DefaultValue** | Pointer to **string** | The value served by this Feature Flag when it is disabled. This could be the actual value or an alias, as appropriate.  This value may be presented to the user in the UI.  | [optional] 
**Rollout** | Pointer to [**FeatureFlagRollout**](FeatureFlagRollout.md) |  | [optional] 

## Methods

### NewFeatureFlagStatus

`func NewFeatureFlagStatus(enabled bool, ) *FeatureFlagStatus`

NewFeatureFlagStatus instantiates a new FeatureFlagStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagStatusWithDefaults

`func NewFeatureFlagStatusWithDefaults() *FeatureFlagStatus`

NewFeatureFlagStatusWithDefaults instantiates a new FeatureFlagStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FeatureFlagStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FeatureFlagStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FeatureFlagStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefaultValue

`func (o *FeatureFlagStatus) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FeatureFlagStatus) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FeatureFlagStatus) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FeatureFlagStatus) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRollout

`func (o *FeatureFlagStatus) GetRollout() FeatureFlagRollout`

GetRollout returns the Rollout field if non-nil, zero value otherwise.

### GetRolloutOk

`func (o *FeatureFlagStatus) GetRolloutOk() (*FeatureFlagRollout, bool)`

GetRolloutOk returns a tuple with the Rollout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollout

`func (o *FeatureFlagStatus) SetRollout(v FeatureFlagRollout)`

SetRollout sets Rollout field to given value.

### HasRollout

`func (o *FeatureFlagStatus) HasRollout() bool`

HasRollout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


