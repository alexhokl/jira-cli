# EditTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateDescription** | Pointer to **string** | The description of the template | [optional] 
**TemplateGenerationOptions** | Pointer to [**CustomTemplateOptions**](CustomTemplateOptions.md) |  | [optional] 
**TemplateKey** | Pointer to **string** | The unique identifier of the template | [optional] 
**TemplateName** | Pointer to **string** | The name of the template | [optional] 

## Methods

### NewEditTemplateRequest

`func NewEditTemplateRequest() *EditTemplateRequest`

NewEditTemplateRequest instantiates a new EditTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTemplateRequestWithDefaults

`func NewEditTemplateRequestWithDefaults() *EditTemplateRequest`

NewEditTemplateRequestWithDefaults instantiates a new EditTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateDescription

`func (o *EditTemplateRequest) GetTemplateDescription() string`

GetTemplateDescription returns the TemplateDescription field if non-nil, zero value otherwise.

### GetTemplateDescriptionOk

`func (o *EditTemplateRequest) GetTemplateDescriptionOk() (*string, bool)`

GetTemplateDescriptionOk returns a tuple with the TemplateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDescription

`func (o *EditTemplateRequest) SetTemplateDescription(v string)`

SetTemplateDescription sets TemplateDescription field to given value.

### HasTemplateDescription

`func (o *EditTemplateRequest) HasTemplateDescription() bool`

HasTemplateDescription returns a boolean if a field has been set.

### GetTemplateGenerationOptions

`func (o *EditTemplateRequest) GetTemplateGenerationOptions() CustomTemplateOptions`

GetTemplateGenerationOptions returns the TemplateGenerationOptions field if non-nil, zero value otherwise.

### GetTemplateGenerationOptionsOk

`func (o *EditTemplateRequest) GetTemplateGenerationOptionsOk() (*CustomTemplateOptions, bool)`

GetTemplateGenerationOptionsOk returns a tuple with the TemplateGenerationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateGenerationOptions

`func (o *EditTemplateRequest) SetTemplateGenerationOptions(v CustomTemplateOptions)`

SetTemplateGenerationOptions sets TemplateGenerationOptions field to given value.

### HasTemplateGenerationOptions

`func (o *EditTemplateRequest) HasTemplateGenerationOptions() bool`

HasTemplateGenerationOptions returns a boolean if a field has been set.

### GetTemplateKey

`func (o *EditTemplateRequest) GetTemplateKey() string`

GetTemplateKey returns the TemplateKey field if non-nil, zero value otherwise.

### GetTemplateKeyOk

`func (o *EditTemplateRequest) GetTemplateKeyOk() (*string, bool)`

GetTemplateKeyOk returns a tuple with the TemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateKey

`func (o *EditTemplateRequest) SetTemplateKey(v string)`

SetTemplateKey sets TemplateKey field to given value.

### HasTemplateKey

`func (o *EditTemplateRequest) HasTemplateKey() bool`

HasTemplateKey returns a boolean if a field has been set.

### GetTemplateName

`func (o *EditTemplateRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *EditTemplateRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *EditTemplateRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *EditTemplateRequest) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


