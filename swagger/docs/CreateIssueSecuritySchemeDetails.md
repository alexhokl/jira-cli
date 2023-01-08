# CreateIssueSecuritySchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the issue security scheme. | [optional] 
**Levels** | Pointer to [**[]SecuritySchemeLevelBean**](SecuritySchemeLevelBean.md) | The list of scheme levels which should be added to the security scheme. | [optional] 
**Name** | **string** | The name of the issue security scheme. Must be unique (case-insensitive). | 

## Methods

### NewCreateIssueSecuritySchemeDetails

`func NewCreateIssueSecuritySchemeDetails(name string, ) *CreateIssueSecuritySchemeDetails`

NewCreateIssueSecuritySchemeDetails instantiates a new CreateIssueSecuritySchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueSecuritySchemeDetailsWithDefaults

`func NewCreateIssueSecuritySchemeDetailsWithDefaults() *CreateIssueSecuritySchemeDetails`

NewCreateIssueSecuritySchemeDetailsWithDefaults instantiates a new CreateIssueSecuritySchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateIssueSecuritySchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIssueSecuritySchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIssueSecuritySchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIssueSecuritySchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLevels

`func (o *CreateIssueSecuritySchemeDetails) GetLevels() []SecuritySchemeLevelBean`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *CreateIssueSecuritySchemeDetails) GetLevelsOk() (*[]SecuritySchemeLevelBean, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *CreateIssueSecuritySchemeDetails) SetLevels(v []SecuritySchemeLevelBean)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *CreateIssueSecuritySchemeDetails) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetName

`func (o *CreateIssueSecuritySchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIssueSecuritySchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIssueSecuritySchemeDetails) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


