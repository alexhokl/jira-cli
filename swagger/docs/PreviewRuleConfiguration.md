# PreviewRuleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A transient identifier for this element, unique within this response but not guaranteed to stable across requests. | [optional] 
**Parameters** | Pointer to **map[string]string** | The parameters of the rule. | [optional] 
**RuleKey** | Pointer to **string** | The rule key of the rule. | [optional] 

## Methods

### NewPreviewRuleConfiguration

`func NewPreviewRuleConfiguration() *PreviewRuleConfiguration`

NewPreviewRuleConfiguration instantiates a new PreviewRuleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewRuleConfigurationWithDefaults

`func NewPreviewRuleConfigurationWithDefaults() *PreviewRuleConfiguration`

NewPreviewRuleConfigurationWithDefaults instantiates a new PreviewRuleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PreviewRuleConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreviewRuleConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreviewRuleConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PreviewRuleConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameters

`func (o *PreviewRuleConfiguration) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PreviewRuleConfiguration) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PreviewRuleConfiguration) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PreviewRuleConfiguration) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRuleKey

`func (o *PreviewRuleConfiguration) GetRuleKey() string`

GetRuleKey returns the RuleKey field if non-nil, zero value otherwise.

### GetRuleKeyOk

`func (o *PreviewRuleConfiguration) GetRuleKeyOk() (*string, bool)`

GetRuleKeyOk returns a tuple with the RuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleKey

`func (o *PreviewRuleConfiguration) SetRuleKey(v string)`

SetRuleKey sets RuleKey field to given value.

### HasRuleKey

`func (o *PreviewRuleConfiguration) HasRuleKey() bool`

HasRuleKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


