# SaveTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateDescription** | Pointer to **string** | The description of the template | [optional] 
**TemplateFromProjectRequest** | Pointer to [**SaveProjectTemplateRequest**](SaveProjectTemplateRequest.md) |  | [optional] 
**TemplateName** | Pointer to **string** | The name of the template | [optional] 

## Methods

### NewSaveTemplateRequest

`func NewSaveTemplateRequest() *SaveTemplateRequest`

NewSaveTemplateRequest instantiates a new SaveTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveTemplateRequestWithDefaults

`func NewSaveTemplateRequestWithDefaults() *SaveTemplateRequest`

NewSaveTemplateRequestWithDefaults instantiates a new SaveTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateDescription

`func (o *SaveTemplateRequest) GetTemplateDescription() string`

GetTemplateDescription returns the TemplateDescription field if non-nil, zero value otherwise.

### GetTemplateDescriptionOk

`func (o *SaveTemplateRequest) GetTemplateDescriptionOk() (*string, bool)`

GetTemplateDescriptionOk returns a tuple with the TemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDescription

`func (o *SaveTemplateRequest) SetTemplateDescription(v string)`

SetTemplateDescription sets TemplateDescription field to given value.

### HasTemplateDescription

`func (o *SaveTemplateRequest) HasTemplateDescription() bool`

HasTemplateDescription returns a boolean if a field has been set.

### GetTemplateFromProjectRequest

`func (o *SaveTemplateRequest) GetTemplateFromProjectRequest() SaveProjectTemplateRequest`

GetTemplateFromProjectRequest returns the TemplateFromProjectRequest field if non-nil, zero value otherwise.

### GetTemplateFromProjectRequestOk

`func (o *SaveTemplateRequest) GetTemplateFromProjectRequestOk() (*SaveProjectTemplateRequest, bool)`

GetTemplateFromProjectRequestOk returns a tuple with the TemplateFromProjectRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFromProjectRequest

`func (o *SaveTemplateRequest) SetTemplateFromProjectRequest(v SaveProjectTemplateRequest)`

SetTemplateFromProjectRequest sets TemplateFromProjectRequest field to given value.

### HasTemplateFromProjectRequest

`func (o *SaveTemplateRequest) HasTemplateFromProjectRequest() bool`

HasTemplateFromProjectRequest returns a boolean if a field has been set.

### GetTemplateName

`func (o *SaveTemplateRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *SaveTemplateRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *SaveTemplateRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *SaveTemplateRequest) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


