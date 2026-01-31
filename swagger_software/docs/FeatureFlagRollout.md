# FeatureFlagRollout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | Pointer to **float32** | If the Feature Flag rollout is a simple percentage rollout  | [optional] 
**Text** | Pointer to **string** | A text status to display that represents the rollout. This could be e.g. a named cohort.  | [optional] 
**Rules** | Pointer to **int32** | A count of the number of rules active for this Feature Flag in an environment.  | [optional] 

## Methods

### NewFeatureFlagRollout

`func NewFeatureFlagRollout() *FeatureFlagRollout`

NewFeatureFlagRollout instantiates a new FeatureFlagRollout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagRolloutWithDefaults

`func NewFeatureFlagRolloutWithDefaults() *FeatureFlagRollout`

NewFeatureFlagRolloutWithDefaults instantiates a new FeatureFlagRollout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *FeatureFlagRollout) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *FeatureFlagRollout) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *FeatureFlagRollout) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *FeatureFlagRollout) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetText

`func (o *FeatureFlagRollout) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FeatureFlagRollout) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FeatureFlagRollout) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *FeatureFlagRollout) HasText() bool`

HasText returns a boolean if a field has been set.

### GetRules

`func (o *FeatureFlagRollout) GetRules() int32`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FeatureFlagRollout) GetRulesOk() (*int32, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FeatureFlagRollout) SetRules(v int32)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FeatureFlagRollout) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


