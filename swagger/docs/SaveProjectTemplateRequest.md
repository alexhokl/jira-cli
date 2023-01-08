# SaveProjectTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int64** | The ID of the target project | [optional] 
**TemplateGenerationOptions** | Pointer to [**CustomTemplateOptions**](CustomTemplateOptions.md) |  | [optional] 
**TemplateType** | Pointer to **string** | The type of the template: LIVE | SNAPSHOT | [optional] 

## Methods

### NewSaveProjectTemplateRequest

`func NewSaveProjectTemplateRequest() *SaveProjectTemplateRequest`

NewSaveProjectTemplateRequest instantiates a new SaveProjectTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveProjectTemplateRequestWithDefaults

`func NewSaveProjectTemplateRequestWithDefaults() *SaveProjectTemplateRequest`

NewSaveProjectTemplateRequestWithDefaults instantiates a new SaveProjectTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SaveProjectTemplateRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SaveProjectTemplateRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SaveProjectTemplateRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SaveProjectTemplateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTemplateGenerationOptions

`func (o *SaveProjectTemplateRequest) GetTemplateGenerationOptions() CustomTemplateOptions`

GetTemplateGenerationOptions returns the TemplateGenerationOptions field if non-nil, zero value otherwise.

### GetTemplateGenerationOptionsOk

`func (o *SaveProjectTemplateRequest) GetTemplateGenerationOptionsOk() (*CustomTemplateOptions, bool)`

GetTemplateGenerationOptionsOk returns a tuple with the TemplateGenerationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateGenerationOptions

`func (o *SaveProjectTemplateRequest) SetTemplateGenerationOptions(v CustomTemplateOptions)`

SetTemplateGenerationOptions sets TemplateGenerationOptions field to given value.

### HasTemplateGenerationOptions

`func (o *SaveProjectTemplateRequest) HasTemplateGenerationOptions() bool`

HasTemplateGenerationOptions returns a boolean if a field has been set.

### GetTemplateType

`func (o *SaveProjectTemplateRequest) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *SaveProjectTemplateRequest) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *SaveProjectTemplateRequest) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *SaveProjectTemplateRequest) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


