# FeatureToggleRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardId** | Pointer to **int64** |  | [optional] 
**Enabling** | Pointer to **bool** |  | [optional] 
**Feature** | Pointer to **string** |  | [optional] 

## Methods

### NewFeatureToggleRequestBean

`func NewFeatureToggleRequestBean() *FeatureToggleRequestBean`

NewFeatureToggleRequestBean instantiates a new FeatureToggleRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureToggleRequestBeanWithDefaults

`func NewFeatureToggleRequestBeanWithDefaults() *FeatureToggleRequestBean`

NewFeatureToggleRequestBeanWithDefaults instantiates a new FeatureToggleRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardId

`func (o *FeatureToggleRequestBean) GetBoardId() int64`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *FeatureToggleRequestBean) GetBoardIdOk() (*int64, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *FeatureToggleRequestBean) SetBoardId(v int64)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *FeatureToggleRequestBean) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetEnabling

`func (o *FeatureToggleRequestBean) GetEnabling() bool`

GetEnabling returns the Enabling field if non-nil, zero value otherwise.

### GetEnablingOk

`func (o *FeatureToggleRequestBean) GetEnablingOk() (*bool, bool)`

GetEnablingOk returns a tuple with the Enabling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabling

`func (o *FeatureToggleRequestBean) SetEnabling(v bool)`

SetEnabling sets Enabling field to given value.

### HasEnabling

`func (o *FeatureToggleRequestBean) HasEnabling() bool`

HasEnabling returns a boolean if a field has been set.

### GetFeature

`func (o *FeatureToggleRequestBean) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureToggleRequestBean) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureToggleRequestBean) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureToggleRequestBean) HasFeature() bool`

HasFeature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


