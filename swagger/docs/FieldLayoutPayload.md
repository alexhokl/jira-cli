# FieldLayoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**[]FieldLayoutConfiguration**](FieldLayoutConfiguration.md) | The field layout configuration. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/\\#api-rest-api-3-fieldconfiguration-post | [optional] 
**Description** | Pointer to **string** | The description of the field layout | [optional] 
**Name** | Pointer to **string** | The name of the field layout | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewFieldLayoutPayload

`func NewFieldLayoutPayload() *FieldLayoutPayload`

NewFieldLayoutPayload instantiates a new FieldLayoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldLayoutPayloadWithDefaults

`func NewFieldLayoutPayloadWithDefaults() *FieldLayoutPayload`

NewFieldLayoutPayloadWithDefaults instantiates a new FieldLayoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *FieldLayoutPayload) GetConfiguration() []FieldLayoutConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FieldLayoutPayload) GetConfigurationOk() (*[]FieldLayoutConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FieldLayoutPayload) SetConfiguration(v []FieldLayoutConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FieldLayoutPayload) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *FieldLayoutPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldLayoutPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldLayoutPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldLayoutPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FieldLayoutPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldLayoutPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldLayoutPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldLayoutPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *FieldLayoutPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *FieldLayoutPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *FieldLayoutPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *FieldLayoutPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


