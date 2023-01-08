# CustomTemplatesProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **string** | The access level of the project. Only used by team-managed project | [optional] 
**AdditionalPropertiesField** | Pointer to **map[string]string** | Additional properties of the project | [optional] 
**AssigneeType** | Pointer to **string** | The default assignee when creating issues in the project | [optional] 
**AvatarId** | Pointer to **int64** | The ID of the project&#39;s avatar. Use the \\[Get project avatars\\](\\#api-rest-api-3-project-projectIdOrKey-avatar-get) operation to list the available avatars in a project. | [optional] 
**CategoryId** | Pointer to **int64** | The ID of the project&#39;s category. A complete list of category IDs is found using the [Get all project categories](#api-rest-api-3-projectCategory-get) operation. | [optional] 
**Description** | Pointer to **string** | Brief description of the project | [optional] 
**EnableComponents** | Pointer to **bool** | Whether components are enabled for the project. Only used by company-managed project | [optional] [default to false]
**Key** | Pointer to **string** | Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. The maximum length is 10 characters. | [optional] 
**Language** | Pointer to **string** | The default language for the project | [optional] 
**LeadAccountId** | Pointer to **string** | The account ID of the project lead. Either &#x60;lead&#x60; or &#x60;leadAccountId&#x60; must be set when creating a project. Cannot be provided with &#x60;lead&#x60;. | [optional] 
**Name** | Pointer to **string** | Name of the project | [optional] 
**Url** | Pointer to **string** | A link to information about this project, such as project documentation | [optional] 

## Methods

### NewCustomTemplatesProjectDetails

`func NewCustomTemplatesProjectDetails() *CustomTemplatesProjectDetails`

NewCustomTemplatesProjectDetails instantiates a new CustomTemplatesProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTemplatesProjectDetailsWithDefaults

`func NewCustomTemplatesProjectDetailsWithDefaults() *CustomTemplatesProjectDetails`

NewCustomTemplatesProjectDetailsWithDefaults instantiates a new CustomTemplatesProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *CustomTemplatesProjectDetails) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CustomTemplatesProjectDetails) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CustomTemplatesProjectDetails) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *CustomTemplatesProjectDetails) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetAdditionalPropertiesField

`func (o *CustomTemplatesProjectDetails) GetAdditionalPropertiesField() map[string]string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *CustomTemplatesProjectDetails) GetAdditionalPropertiesFieldOk() (*map[string]string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *CustomTemplatesProjectDetails) SetAdditionalPropertiesField(v map[string]string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.

### HasAdditionalPropertiesField

`func (o *CustomTemplatesProjectDetails) HasAdditionalPropertiesField() bool`

HasAdditionalPropertiesField returns a boolean if a field has been set.

### GetAssigneeType

`func (o *CustomTemplatesProjectDetails) GetAssigneeType() string`

GetAssigneeType returns the AssigneeType field if non-nil, zero value otherwise.

### GetAssigneeTypeOk

`func (o *CustomTemplatesProjectDetails) GetAssigneeTypeOk() (*string, bool)`

GetAssigneeTypeOk returns a tuple with the AssigneeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeType

`func (o *CustomTemplatesProjectDetails) SetAssigneeType(v string)`

SetAssigneeType sets AssigneeType field to given value.

### HasAssigneeType

`func (o *CustomTemplatesProjectDetails) HasAssigneeType() bool`

HasAssigneeType returns a boolean if a field has been set.

### GetAvatarId

`func (o *CustomTemplatesProjectDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CustomTemplatesProjectDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CustomTemplatesProjectDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *CustomTemplatesProjectDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetCategoryId

`func (o *CustomTemplatesProjectDetails) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CustomTemplatesProjectDetails) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CustomTemplatesProjectDetails) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CustomTemplatesProjectDetails) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetDescription

`func (o *CustomTemplatesProjectDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomTemplatesProjectDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomTemplatesProjectDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomTemplatesProjectDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableComponents

`func (o *CustomTemplatesProjectDetails) GetEnableComponents() bool`

GetEnableComponents returns the EnableComponents field if non-nil, zero value otherwise.

### GetEnableComponentsOk

`func (o *CustomTemplatesProjectDetails) GetEnableComponentsOk() (*bool, bool)`

GetEnableComponentsOk returns a tuple with the EnableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableComponents

`func (o *CustomTemplatesProjectDetails) SetEnableComponents(v bool)`

SetEnableComponents sets EnableComponents field to given value.

### HasEnableComponents

`func (o *CustomTemplatesProjectDetails) HasEnableComponents() bool`

HasEnableComponents returns a boolean if a field has been set.

### GetKey

`func (o *CustomTemplatesProjectDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomTemplatesProjectDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomTemplatesProjectDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomTemplatesProjectDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLanguage

`func (o *CustomTemplatesProjectDetails) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CustomTemplatesProjectDetails) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CustomTemplatesProjectDetails) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CustomTemplatesProjectDetails) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLeadAccountId

`func (o *CustomTemplatesProjectDetails) GetLeadAccountId() string`

GetLeadAccountId returns the LeadAccountId field if non-nil, zero value otherwise.

### GetLeadAccountIdOk

`func (o *CustomTemplatesProjectDetails) GetLeadAccountIdOk() (*string, bool)`

GetLeadAccountIdOk returns a tuple with the LeadAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadAccountId

`func (o *CustomTemplatesProjectDetails) SetLeadAccountId(v string)`

SetLeadAccountId sets LeadAccountId field to given value.

### HasLeadAccountId

`func (o *CustomTemplatesProjectDetails) HasLeadAccountId() bool`

HasLeadAccountId returns a boolean if a field has been set.

### GetName

`func (o *CustomTemplatesProjectDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomTemplatesProjectDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomTemplatesProjectDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomTemplatesProjectDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *CustomTemplatesProjectDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomTemplatesProjectDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomTemplatesProjectDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CustomTemplatesProjectDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


