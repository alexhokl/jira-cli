# ProjectFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | Pointer to **string** | The key of the feature. | [optional] 
**ImageUri** | Pointer to **string** | URI for the image representing the feature. | [optional] 
**LocalisedDescription** | Pointer to **string** | Localized display description for the feature. | [optional] 
**LocalisedName** | Pointer to **string** | Localized display name for the feature. | [optional] 
**Prerequisites** | Pointer to **[]string** | List of keys of the features required to enable the feature. | [optional] 
**ProjectId** | Pointer to **int64** | The ID of the project. | [optional] 
**State** | Pointer to **string** | The state of the feature. When updating the state of a feature, only ENABLED and DISABLED are supported. Responses can contain all values | [optional] 
**ToggleLocked** | Pointer to **bool** | Whether the state of the feature can be updated. | [optional] 

## Methods

### NewProjectFeature

`func NewProjectFeature() *ProjectFeature`

NewProjectFeature instantiates a new ProjectFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectFeatureWithDefaults

`func NewProjectFeatureWithDefaults() *ProjectFeature`

NewProjectFeatureWithDefaults instantiates a new ProjectFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeature

`func (o *ProjectFeature) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ProjectFeature) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ProjectFeature) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *ProjectFeature) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetImageUri

`func (o *ProjectFeature) GetImageUri() string`

GetImageUri returns the ImageUri field if non-nil, zero value otherwise.

### GetImageUriOk

`func (o *ProjectFeature) GetImageUriOk() (*string, bool)`

GetImageUriOk returns a tuple with the ImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUri

`func (o *ProjectFeature) SetImageUri(v string)`

SetImageUri sets ImageUri field to given value.

### HasImageUri

`func (o *ProjectFeature) HasImageUri() bool`

HasImageUri returns a boolean if a field has been set.

### GetLocalisedDescription

`func (o *ProjectFeature) GetLocalisedDescription() string`

GetLocalisedDescription returns the LocalisedDescription field if non-nil, zero value otherwise.

### GetLocalisedDescriptionOk

`func (o *ProjectFeature) GetLocalisedDescriptionOk() (*string, bool)`

GetLocalisedDescriptionOk returns a tuple with the LocalisedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisedDescription

`func (o *ProjectFeature) SetLocalisedDescription(v string)`

SetLocalisedDescription sets LocalisedDescription field to given value.

### HasLocalisedDescription

`func (o *ProjectFeature) HasLocalisedDescription() bool`

HasLocalisedDescription returns a boolean if a field has been set.

### GetLocalisedName

`func (o *ProjectFeature) GetLocalisedName() string`

GetLocalisedName returns the LocalisedName field if non-nil, zero value otherwise.

### GetLocalisedNameOk

`func (o *ProjectFeature) GetLocalisedNameOk() (*string, bool)`

GetLocalisedNameOk returns a tuple with the LocalisedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisedName

`func (o *ProjectFeature) SetLocalisedName(v string)`

SetLocalisedName sets LocalisedName field to given value.

### HasLocalisedName

`func (o *ProjectFeature) HasLocalisedName() bool`

HasLocalisedName returns a boolean if a field has been set.

### GetPrerequisites

`func (o *ProjectFeature) GetPrerequisites() []string`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *ProjectFeature) GetPrerequisitesOk() (*[]string, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *ProjectFeature) SetPrerequisites(v []string)`

SetPrerequisites sets Prerequisites field to given value.

### HasPrerequisites

`func (o *ProjectFeature) HasPrerequisites() bool`

HasPrerequisites returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectFeature) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectFeature) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectFeature) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectFeature) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetState

`func (o *ProjectFeature) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectFeature) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectFeature) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectFeature) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToggleLocked

`func (o *ProjectFeature) GetToggleLocked() bool`

GetToggleLocked returns the ToggleLocked field if non-nil, zero value otherwise.

### GetToggleLockedOk

`func (o *ProjectFeature) GetToggleLockedOk() (*bool, bool)`

GetToggleLockedOk returns a tuple with the ToggleLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToggleLocked

`func (o *ProjectFeature) SetToggleLocked(v bool)`

SetToggleLocked sets ToggleLocked field to given value.

### HasToggleLocked

`func (o *ProjectFeature) HasToggleLocked() bool`

HasToggleLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


