# ProjectCustomTemplateCreateRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**CustomTemplatesProjectDetails**](CustomTemplatesProjectDetails.md) |  | [optional] 
**Template** | Pointer to [**CustomTemplateRequestDTO**](CustomTemplateRequestDTO.md) |  | [optional] 

## Methods

### NewProjectCustomTemplateCreateRequestDTO

`func NewProjectCustomTemplateCreateRequestDTO() *ProjectCustomTemplateCreateRequestDTO`

NewProjectCustomTemplateCreateRequestDTO instantiates a new ProjectCustomTemplateCreateRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCustomTemplateCreateRequestDTOWithDefaults

`func NewProjectCustomTemplateCreateRequestDTOWithDefaults() *ProjectCustomTemplateCreateRequestDTO`

NewProjectCustomTemplateCreateRequestDTOWithDefaults instantiates a new ProjectCustomTemplateCreateRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ProjectCustomTemplateCreateRequestDTO) GetDetails() CustomTemplatesProjectDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProjectCustomTemplateCreateRequestDTO) GetDetailsOk() (*CustomTemplatesProjectDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProjectCustomTemplateCreateRequestDTO) SetDetails(v CustomTemplatesProjectDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProjectCustomTemplateCreateRequestDTO) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTemplate

`func (o *ProjectCustomTemplateCreateRequestDTO) GetTemplate() CustomTemplateRequestDTO`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ProjectCustomTemplateCreateRequestDTO) GetTemplateOk() (*CustomTemplateRequestDTO, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ProjectCustomTemplateCreateRequestDTO) SetTemplate(v CustomTemplateRequestDTO)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ProjectCustomTemplateCreateRequestDTO) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


