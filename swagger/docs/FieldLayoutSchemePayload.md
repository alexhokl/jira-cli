# FieldLayoutSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFieldLayout** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the field layout scheme | [optional] 
**ExplicitMappings** | Pointer to [**map[string]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | There is a default configuration \&quot;fieldlayout\&quot; that is applied to all issue types using this scheme that don&#39;t have an explicit mapping users can create (or re-use existing) configurations for other issue types and map them to this scheme | [optional] 
**Name** | Pointer to **string** | The name of the field layout scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewFieldLayoutSchemePayload

`func NewFieldLayoutSchemePayload() *FieldLayoutSchemePayload`

NewFieldLayoutSchemePayload instantiates a new FieldLayoutSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldLayoutSchemePayloadWithDefaults

`func NewFieldLayoutSchemePayloadWithDefaults() *FieldLayoutSchemePayload`

NewFieldLayoutSchemePayloadWithDefaults instantiates a new FieldLayoutSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFieldLayout

`func (o *FieldLayoutSchemePayload) GetDefaultFieldLayout() ProjectCreateResourceIdentifier`

GetDefaultFieldLayout returns the DefaultFieldLayout field if non-nil, zero value otherwise.

### GetDefaultFieldLayoutOk

`func (o *FieldLayoutSchemePayload) GetDefaultFieldLayoutOk() (*ProjectCreateResourceIdentifier, bool)`

GetDefaultFieldLayoutOk returns a tuple with the DefaultFieldLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFieldLayout

`func (o *FieldLayoutSchemePayload) SetDefaultFieldLayout(v ProjectCreateResourceIdentifier)`

SetDefaultFieldLayout sets DefaultFieldLayout field to given value.

### HasDefaultFieldLayout

`func (o *FieldLayoutSchemePayload) HasDefaultFieldLayout() bool`

HasDefaultFieldLayout returns a boolean if a field has been set.

### GetDescription

`func (o *FieldLayoutSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldLayoutSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldLayoutSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldLayoutSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExplicitMappings

`func (o *FieldLayoutSchemePayload) GetExplicitMappings() map[string]ProjectCreateResourceIdentifier`

GetExplicitMappings returns the ExplicitMappings field if non-nil, zero value otherwise.

### GetExplicitMappingsOk

`func (o *FieldLayoutSchemePayload) GetExplicitMappingsOk() (*map[string]ProjectCreateResourceIdentifier, bool)`

GetExplicitMappingsOk returns a tuple with the ExplicitMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMappings

`func (o *FieldLayoutSchemePayload) SetExplicitMappings(v map[string]ProjectCreateResourceIdentifier)`

SetExplicitMappings sets ExplicitMappings field to given value.

### HasExplicitMappings

`func (o *FieldLayoutSchemePayload) HasExplicitMappings() bool`

HasExplicitMappings returns a boolean if a field has been set.

### GetName

`func (o *FieldLayoutSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldLayoutSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldLayoutSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldLayoutSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *FieldLayoutSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *FieldLayoutSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *FieldLayoutSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *FieldLayoutSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


