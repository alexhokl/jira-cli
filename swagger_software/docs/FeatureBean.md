# FeatureBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardFeature** | Pointer to **string** |  | [optional] 
**BoardId** | Pointer to **int64** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**FeatureType** | Pointer to **string** |  | [optional] 
**ImageUri** | Pointer to **string** |  | [optional] 
**LearnMoreArticleId** | Pointer to **string** |  | [optional] 
**LearnMoreLink** | Pointer to **string** |  | [optional] 
**LocalisedDescription** | Pointer to **string** |  | [optional] 
**LocalisedGroup** | Pointer to **string** |  | [optional] 
**LocalisedName** | Pointer to **string** |  | [optional] 
**PermissibleEstimationTypes** | Pointer to [**[]GetFeaturesForBoard200ResponseFeaturesInnerPermissibleEstimationTypesInner**](GetFeaturesForBoard200ResponseFeaturesInnerPermissibleEstimationTypesInner.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ToggleLocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewFeatureBean

`func NewFeatureBean() *FeatureBean`

NewFeatureBean instantiates a new FeatureBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureBeanWithDefaults

`func NewFeatureBeanWithDefaults() *FeatureBean`

NewFeatureBeanWithDefaults instantiates a new FeatureBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardFeature

`func (o *FeatureBean) GetBoardFeature() string`

GetBoardFeature returns the BoardFeature field if non-nil, zero value otherwise.

### GetBoardFeatureOk

`func (o *FeatureBean) GetBoardFeatureOk() (*string, bool)`

GetBoardFeatureOk returns a tuple with the BoardFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardFeature

`func (o *FeatureBean) SetBoardFeature(v string)`

SetBoardFeature sets BoardFeature field to given value.

### HasBoardFeature

`func (o *FeatureBean) HasBoardFeature() bool`

HasBoardFeature returns a boolean if a field has been set.

### GetBoardId

`func (o *FeatureBean) GetBoardId() int64`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *FeatureBean) GetBoardIdOk() (*int64, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *FeatureBean) SetBoardId(v int64)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *FeatureBean) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetFeatureId

`func (o *FeatureBean) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *FeatureBean) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *FeatureBean) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *FeatureBean) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureType

`func (o *FeatureBean) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *FeatureBean) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *FeatureBean) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *FeatureBean) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetImageUri

`func (o *FeatureBean) GetImageUri() string`

GetImageUri returns the ImageUri field if non-nil, zero value otherwise.

### GetImageUriOk

`func (o *FeatureBean) GetImageUriOk() (*string, bool)`

GetImageUriOk returns a tuple with the ImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUri

`func (o *FeatureBean) SetImageUri(v string)`

SetImageUri sets ImageUri field to given value.

### HasImageUri

`func (o *FeatureBean) HasImageUri() bool`

HasImageUri returns a boolean if a field has been set.

### GetLearnMoreArticleId

`func (o *FeatureBean) GetLearnMoreArticleId() string`

GetLearnMoreArticleId returns the LearnMoreArticleId field if non-nil, zero value otherwise.

### GetLearnMoreArticleIdOk

`func (o *FeatureBean) GetLearnMoreArticleIdOk() (*string, bool)`

GetLearnMoreArticleIdOk returns a tuple with the LearnMoreArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMoreArticleId

`func (o *FeatureBean) SetLearnMoreArticleId(v string)`

SetLearnMoreArticleId sets LearnMoreArticleId field to given value.

### HasLearnMoreArticleId

`func (o *FeatureBean) HasLearnMoreArticleId() bool`

HasLearnMoreArticleId returns a boolean if a field has been set.

### GetLearnMoreLink

`func (o *FeatureBean) GetLearnMoreLink() string`

GetLearnMoreLink returns the LearnMoreLink field if non-nil, zero value otherwise.

### GetLearnMoreLinkOk

`func (o *FeatureBean) GetLearnMoreLinkOk() (*string, bool)`

GetLearnMoreLinkOk returns a tuple with the LearnMoreLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMoreLink

`func (o *FeatureBean) SetLearnMoreLink(v string)`

SetLearnMoreLink sets LearnMoreLink field to given value.

### HasLearnMoreLink

`func (o *FeatureBean) HasLearnMoreLink() bool`

HasLearnMoreLink returns a boolean if a field has been set.

### GetLocalisedDescription

`func (o *FeatureBean) GetLocalisedDescription() string`

GetLocalisedDescription returns the LocalisedDescription field if non-nil, zero value otherwise.

### GetLocalisedDescriptionOk

`func (o *FeatureBean) GetLocalisedDescriptionOk() (*string, bool)`

GetLocalisedDescriptionOk returns a tuple with the LocalisedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisedDescription

`func (o *FeatureBean) SetLocalisedDescription(v string)`

SetLocalisedDescription sets LocalisedDescription field to given value.

### HasLocalisedDescription

`func (o *FeatureBean) HasLocalisedDescription() bool`

HasLocalisedDescription returns a boolean if a field has been set.

### GetLocalisedGroup

`func (o *FeatureBean) GetLocalisedGroup() string`

GetLocalisedGroup returns the LocalisedGroup field if non-nil, zero value otherwise.

### GetLocalisedGroupOk

`func (o *FeatureBean) GetLocalisedGroupOk() (*string, bool)`

GetLocalisedGroupOk returns a tuple with the LocalisedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisedGroup

`func (o *FeatureBean) SetLocalisedGroup(v string)`

SetLocalisedGroup sets LocalisedGroup field to given value.

### HasLocalisedGroup

`func (o *FeatureBean) HasLocalisedGroup() bool`

HasLocalisedGroup returns a boolean if a field has been set.

### GetLocalisedName

`func (o *FeatureBean) GetLocalisedName() string`

GetLocalisedName returns the LocalisedName field if non-nil, zero value otherwise.

### GetLocalisedNameOk

`func (o *FeatureBean) GetLocalisedNameOk() (*string, bool)`

GetLocalisedNameOk returns a tuple with the LocalisedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisedName

`func (o *FeatureBean) SetLocalisedName(v string)`

SetLocalisedName sets LocalisedName field to given value.

### HasLocalisedName

`func (o *FeatureBean) HasLocalisedName() bool`

HasLocalisedName returns a boolean if a field has been set.

### GetPermissibleEstimationTypes

`func (o *FeatureBean) GetPermissibleEstimationTypes() []GetFeaturesForBoard200ResponseFeaturesInnerPermissibleEstimationTypesInner`

GetPermissibleEstimationTypes returns the PermissibleEstimationTypes field if non-nil, zero value otherwise.

### GetPermissibleEstimationTypesOk

`func (o *FeatureBean) GetPermissibleEstimationTypesOk() (*[]GetFeaturesForBoard200ResponseFeaturesInnerPermissibleEstimationTypesInner, bool)`

GetPermissibleEstimationTypesOk returns a tuple with the PermissibleEstimationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissibleEstimationTypes

`func (o *FeatureBean) SetPermissibleEstimationTypes(v []GetFeaturesForBoard200ResponseFeaturesInnerPermissibleEstimationTypesInner)`

SetPermissibleEstimationTypes sets PermissibleEstimationTypes field to given value.

### HasPermissibleEstimationTypes

`func (o *FeatureBean) HasPermissibleEstimationTypes() bool`

HasPermissibleEstimationTypes returns a boolean if a field has been set.

### GetState

`func (o *FeatureBean) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FeatureBean) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FeatureBean) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FeatureBean) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToggleLocked

`func (o *FeatureBean) GetToggleLocked() bool`

GetToggleLocked returns the ToggleLocked field if non-nil, zero value otherwise.

### GetToggleLockedOk

`func (o *FeatureBean) GetToggleLockedOk() (*bool, bool)`

GetToggleLockedOk returns a tuple with the ToggleLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggleLocked

`func (o *FeatureBean) SetToggleLocked(v bool)`

SetToggleLocked sets ToggleLocked field to given value.

### HasToggleLocked

`func (o *FeatureBean) HasToggleLocked() bool`

HasToggleLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


