# ProjectTemplateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archetype** | Pointer to [**ProjectArchetype**](ProjectArchetype.md) |  | [optional] 
**DefaultBoardView** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LiveTemplateProjectIdReference** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectTemplateKey** | Pointer to [**ProjectTemplateKey**](ProjectTemplateKey.md) |  | [optional] 
**SnapshotTemplate** | Pointer to **map[string]interface{}** |  | [optional] 
**TemplateGenerationOptions** | Pointer to [**CustomTemplateOptions**](CustomTemplateOptions.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectTemplateModel

`func NewProjectTemplateModel() *ProjectTemplateModel`

NewProjectTemplateModel instantiates a new ProjectTemplateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateModelWithDefaults

`func NewProjectTemplateModelWithDefaults() *ProjectTemplateModel`

NewProjectTemplateModelWithDefaults instantiates a new ProjectTemplateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchetype

`func (o *ProjectTemplateModel) GetArchetype() ProjectArchetype`

GetArchetype returns the Archetype field if non-nil, zero value otherwise.

### GetArchetypeOk

`func (o *ProjectTemplateModel) GetArchetypeOk() (*ProjectArchetype, bool)`

GetArchetypeOk returns a tuple with the Archetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchetype

`func (o *ProjectTemplateModel) SetArchetype(v ProjectArchetype)`

SetArchetype sets Archetype field to given value.

### HasArchetype

`func (o *ProjectTemplateModel) HasArchetype() bool`

HasArchetype returns a boolean if a field has been set.

### GetDefaultBoardView

`func (o *ProjectTemplateModel) GetDefaultBoardView() string`

GetDefaultBoardView returns the DefaultBoardView field if non-nil, zero value otherwise.

### GetDefaultBoardViewOk

`func (o *ProjectTemplateModel) GetDefaultBoardViewOk() (*string, bool)`

GetDefaultBoardViewOk returns a tuple with the DefaultBoardView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBoardView

`func (o *ProjectTemplateModel) SetDefaultBoardView(v string)`

SetDefaultBoardView sets DefaultBoardView field to given value.

### HasDefaultBoardView

`func (o *ProjectTemplateModel) HasDefaultBoardView() bool`

HasDefaultBoardView returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectTemplateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTemplateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTemplateModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTemplateModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLiveTemplateProjectIdReference

`func (o *ProjectTemplateModel) GetLiveTemplateProjectIdReference() int64`

GetLiveTemplateProjectIdReference returns the LiveTemplateProjectIdReference field if non-nil, zero value otherwise.

### GetLiveTemplateProjectIdReferenceOk

`func (o *ProjectTemplateModel) GetLiveTemplateProjectIdReferenceOk() (*int64, bool)`

GetLiveTemplateProjectIdReferenceOk returns a tuple with the LiveTemplateProjectIdReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTemplateProjectIdReference

`func (o *ProjectTemplateModel) SetLiveTemplateProjectIdReference(v int64)`

SetLiveTemplateProjectIdReference sets LiveTemplateProjectIdReference field to given value.

### HasLiveTemplateProjectIdReference

`func (o *ProjectTemplateModel) HasLiveTemplateProjectIdReference() bool`

HasLiveTemplateProjectIdReference returns a boolean if a field has been set.

### GetName

`func (o *ProjectTemplateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectTemplateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectTemplateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectTemplateModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectTemplateKey

`func (o *ProjectTemplateModel) GetProjectTemplateKey() ProjectTemplateKey`

GetProjectTemplateKey returns the ProjectTemplateKey field if non-nil, zero value otherwise.

### GetProjectTemplateKeyOk

`func (o *ProjectTemplateModel) GetProjectTemplateKeyOk() (*ProjectTemplateKey, bool)`

GetProjectTemplateKeyOk returns a tuple with the ProjectTemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplateKey

`func (o *ProjectTemplateModel) SetProjectTemplateKey(v ProjectTemplateKey)`

SetProjectTemplateKey sets ProjectTemplateKey field to given value.

### HasProjectTemplateKey

`func (o *ProjectTemplateModel) HasProjectTemplateKey() bool`

HasProjectTemplateKey returns a boolean if a field has been set.

### GetSnapshotTemplate

`func (o *ProjectTemplateModel) GetSnapshotTemplate() map[string]interface{}`

GetSnapshotTemplate returns the SnapshotTemplate field if non-nil, zero value otherwise.

### GetSnapshotTemplateOk

`func (o *ProjectTemplateModel) GetSnapshotTemplateOk() (*map[string]interface{}, bool)`

GetSnapshotTemplateOk returns a tuple with the SnapshotTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTemplate

`func (o *ProjectTemplateModel) SetSnapshotTemplate(v map[string]interface{})`

SetSnapshotTemplate sets SnapshotTemplate field to given value.

### HasSnapshotTemplate

`func (o *ProjectTemplateModel) HasSnapshotTemplate() bool`

HasSnapshotTemplate returns a boolean if a field has been set.

### GetTemplateGenerationOptions

`func (o *ProjectTemplateModel) GetTemplateGenerationOptions() CustomTemplateOptions`

GetTemplateGenerationOptions returns the TemplateGenerationOptions field if non-nil, zero value otherwise.

### GetTemplateGenerationOptionsOk

`func (o *ProjectTemplateModel) GetTemplateGenerationOptionsOk() (*CustomTemplateOptions, bool)`

GetTemplateGenerationOptionsOk returns a tuple with the TemplateGenerationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateGenerationOptions

`func (o *ProjectTemplateModel) SetTemplateGenerationOptions(v CustomTemplateOptions)`

SetTemplateGenerationOptions sets TemplateGenerationOptions field to given value.

### HasTemplateGenerationOptions

`func (o *ProjectTemplateModel) HasTemplateGenerationOptions() bool`

HasTemplateGenerationOptions returns a boolean if a field has been set.

### GetType

`func (o *ProjectTemplateModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectTemplateModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectTemplateModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectTemplateModel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


