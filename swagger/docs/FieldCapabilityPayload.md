# FieldCapabilityPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldDefinitions** | Pointer to [**[]CustomFieldPayload**](CustomFieldPayload.md) | The custom field definitions. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/\\#api-rest-api-3-field-post | [optional] 
**FieldLayoutScheme** | Pointer to [**NullableFieldLayoutSchemePayload**](FieldLayoutSchemePayload.md) |  | [optional] 
**FieldLayouts** | Pointer to [**[]FieldLayoutPayload**](FieldLayoutPayload.md) | The field layouts configuration. | [optional] 
**IssueLayouts** | Pointer to [**[]IssueLayoutPayload**](IssueLayoutPayload.md) | The issue layouts configuration | [optional] 
**IssueTypeScreenScheme** | Pointer to [**NullableIssueTypeScreenSchemePayload**](IssueTypeScreenSchemePayload.md) |  | [optional] 
**ScreenScheme** | Pointer to [**[]ScreenSchemePayload**](ScreenSchemePayload.md) | The screen schemes See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-schemes/\\#api-rest-api-3-screenscheme-post | [optional] 
**Screens** | Pointer to [**[]ScreenPayload**](ScreenPayload.md) | The screens. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/\\#api-rest-api-3-screens-post | [optional] 

## Methods

### NewFieldCapabilityPayload

`func NewFieldCapabilityPayload() *FieldCapabilityPayload`

NewFieldCapabilityPayload instantiates a new FieldCapabilityPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldCapabilityPayloadWithDefaults

`func NewFieldCapabilityPayloadWithDefaults() *FieldCapabilityPayload`

NewFieldCapabilityPayloadWithDefaults instantiates a new FieldCapabilityPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldDefinitions

`func (o *FieldCapabilityPayload) GetCustomFieldDefinitions() []CustomFieldPayload`

GetCustomFieldDefinitions returns the CustomFieldDefinitions field if non-nil, zero value otherwise.

### GetCustomFieldDefinitionsOk

`func (o *FieldCapabilityPayload) GetCustomFieldDefinitionsOk() (*[]CustomFieldPayload, bool)`

GetCustomFieldDefinitionsOk returns a tuple with the CustomFieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldDefinitions

`func (o *FieldCapabilityPayload) SetCustomFieldDefinitions(v []CustomFieldPayload)`

SetCustomFieldDefinitions sets CustomFieldDefinitions field to given value.

### HasCustomFieldDefinitions

`func (o *FieldCapabilityPayload) HasCustomFieldDefinitions() bool`

HasCustomFieldDefinitions returns a boolean if a field has been set.

### SetCustomFieldDefinitionsNil

`func (o *FieldCapabilityPayload) SetCustomFieldDefinitionsNil(b bool)`

 SetCustomFieldDefinitionsNil sets the value for CustomFieldDefinitions to be an explicit nil

### UnsetCustomFieldDefinitions
`func (o *FieldCapabilityPayload) UnsetCustomFieldDefinitions()`

UnsetCustomFieldDefinitions ensures that no value is present for CustomFieldDefinitions, not even an explicit nil
### GetFieldLayoutScheme

`func (o *FieldCapabilityPayload) GetFieldLayoutScheme() FieldLayoutSchemePayload`

GetFieldLayoutScheme returns the FieldLayoutScheme field if non-nil, zero value otherwise.

### GetFieldLayoutSchemeOk

`func (o *FieldCapabilityPayload) GetFieldLayoutSchemeOk() (*FieldLayoutSchemePayload, bool)`

GetFieldLayoutSchemeOk returns a tuple with the FieldLayoutScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLayoutScheme

`func (o *FieldCapabilityPayload) SetFieldLayoutScheme(v FieldLayoutSchemePayload)`

SetFieldLayoutScheme sets FieldLayoutScheme field to given value.

### HasFieldLayoutScheme

`func (o *FieldCapabilityPayload) HasFieldLayoutScheme() bool`

HasFieldLayoutScheme returns a boolean if a field has been set.

### SetFieldLayoutSchemeNil

`func (o *FieldCapabilityPayload) SetFieldLayoutSchemeNil(b bool)`

 SetFieldLayoutSchemeNil sets the value for FieldLayoutScheme to be an explicit nil

### UnsetFieldLayoutScheme
`func (o *FieldCapabilityPayload) UnsetFieldLayoutScheme()`

UnsetFieldLayoutScheme ensures that no value is present for FieldLayoutScheme, not even an explicit nil
### GetFieldLayouts

`func (o *FieldCapabilityPayload) GetFieldLayouts() []FieldLayoutPayload`

GetFieldLayouts returns the FieldLayouts field if non-nil, zero value otherwise.

### GetFieldLayoutsOk

`func (o *FieldCapabilityPayload) GetFieldLayoutsOk() (*[]FieldLayoutPayload, bool)`

GetFieldLayoutsOk returns a tuple with the FieldLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLayouts

`func (o *FieldCapabilityPayload) SetFieldLayouts(v []FieldLayoutPayload)`

SetFieldLayouts sets FieldLayouts field to given value.

### HasFieldLayouts

`func (o *FieldCapabilityPayload) HasFieldLayouts() bool`

HasFieldLayouts returns a boolean if a field has been set.

### SetFieldLayoutsNil

`func (o *FieldCapabilityPayload) SetFieldLayoutsNil(b bool)`

 SetFieldLayoutsNil sets the value for FieldLayouts to be an explicit nil

### UnsetFieldLayouts
`func (o *FieldCapabilityPayload) UnsetFieldLayouts()`

UnsetFieldLayouts ensures that no value is present for FieldLayouts, not even an explicit nil
### GetIssueLayouts

`func (o *FieldCapabilityPayload) GetIssueLayouts() []IssueLayoutPayload`

GetIssueLayouts returns the IssueLayouts field if non-nil, zero value otherwise.

### GetIssueLayoutsOk

`func (o *FieldCapabilityPayload) GetIssueLayoutsOk() (*[]IssueLayoutPayload, bool)`

GetIssueLayoutsOk returns a tuple with the IssueLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLayouts

`func (o *FieldCapabilityPayload) SetIssueLayouts(v []IssueLayoutPayload)`

SetIssueLayouts sets IssueLayouts field to given value.

### HasIssueLayouts

`func (o *FieldCapabilityPayload) HasIssueLayouts() bool`

HasIssueLayouts returns a boolean if a field has been set.

### SetIssueLayoutsNil

`func (o *FieldCapabilityPayload) SetIssueLayoutsNil(b bool)`

 SetIssueLayoutsNil sets the value for IssueLayouts to be an explicit nil

### UnsetIssueLayouts
`func (o *FieldCapabilityPayload) UnsetIssueLayouts()`

UnsetIssueLayouts ensures that no value is present for IssueLayouts, not even an explicit nil
### GetIssueTypeScreenScheme

`func (o *FieldCapabilityPayload) GetIssueTypeScreenScheme() IssueTypeScreenSchemePayload`

GetIssueTypeScreenScheme returns the IssueTypeScreenScheme field if non-nil, zero value otherwise.

### GetIssueTypeScreenSchemeOk

`func (o *FieldCapabilityPayload) GetIssueTypeScreenSchemeOk() (*IssueTypeScreenSchemePayload, bool)`

GetIssueTypeScreenSchemeOk returns a tuple with the IssueTypeScreenScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScreenScheme

`func (o *FieldCapabilityPayload) SetIssueTypeScreenScheme(v IssueTypeScreenSchemePayload)`

SetIssueTypeScreenScheme sets IssueTypeScreenScheme field to given value.

### HasIssueTypeScreenScheme

`func (o *FieldCapabilityPayload) HasIssueTypeScreenScheme() bool`

HasIssueTypeScreenScheme returns a boolean if a field has been set.

### SetIssueTypeScreenSchemeNil

`func (o *FieldCapabilityPayload) SetIssueTypeScreenSchemeNil(b bool)`

 SetIssueTypeScreenSchemeNil sets the value for IssueTypeScreenScheme to be an explicit nil

### UnsetIssueTypeScreenScheme
`func (o *FieldCapabilityPayload) UnsetIssueTypeScreenScheme()`

UnsetIssueTypeScreenScheme ensures that no value is present for IssueTypeScreenScheme, not even an explicit nil
### GetScreenScheme

`func (o *FieldCapabilityPayload) GetScreenScheme() []ScreenSchemePayload`

GetScreenScheme returns the ScreenScheme field if non-nil, zero value otherwise.

### GetScreenSchemeOk

`func (o *FieldCapabilityPayload) GetScreenSchemeOk() (*[]ScreenSchemePayload, bool)`

GetScreenSchemeOk returns a tuple with the ScreenScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenScheme

`func (o *FieldCapabilityPayload) SetScreenScheme(v []ScreenSchemePayload)`

SetScreenScheme sets ScreenScheme field to given value.

### HasScreenScheme

`func (o *FieldCapabilityPayload) HasScreenScheme() bool`

HasScreenScheme returns a boolean if a field has been set.

### SetScreenSchemeNil

`func (o *FieldCapabilityPayload) SetScreenSchemeNil(b bool)`

 SetScreenSchemeNil sets the value for ScreenScheme to be an explicit nil

### UnsetScreenScheme
`func (o *FieldCapabilityPayload) UnsetScreenScheme()`

UnsetScreenScheme ensures that no value is present for ScreenScheme, not even an explicit nil
### GetScreens

`func (o *FieldCapabilityPayload) GetScreens() []ScreenPayload`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *FieldCapabilityPayload) GetScreensOk() (*[]ScreenPayload, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *FieldCapabilityPayload) SetScreens(v []ScreenPayload)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *FieldCapabilityPayload) HasScreens() bool`

HasScreens returns a boolean if a field has been set.

### SetScreensNil

`func (o *FieldCapabilityPayload) SetScreensNil(b bool)`

 SetScreensNil sets the value for Screens to be an explicit nil

### UnsetScreens
`func (o *FieldCapabilityPayload) UnsetScreens()`

UnsetScreens ensures that no value is present for Screens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


