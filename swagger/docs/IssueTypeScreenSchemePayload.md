# IssueTypeScreenSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultScreenScheme** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the issue type screen scheme | [optional] 
**ExplicitMappings** | Pointer to [**map[string]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | The IDs of the screen schemes for the issue type IDs and default. A default entry is required to create an issue type screen scheme, it defines the mapping for all issue types without a screen scheme. | [optional] 
**Name** | Pointer to **string** | The name of the issue type screen scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewIssueTypeScreenSchemePayload

`func NewIssueTypeScreenSchemePayload() *IssueTypeScreenSchemePayload`

NewIssueTypeScreenSchemePayload instantiates a new IssueTypeScreenSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeScreenSchemePayloadWithDefaults

`func NewIssueTypeScreenSchemePayloadWithDefaults() *IssueTypeScreenSchemePayload`

NewIssueTypeScreenSchemePayloadWithDefaults instantiates a new IssueTypeScreenSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultScreenScheme

`func (o *IssueTypeScreenSchemePayload) GetDefaultScreenScheme() ProjectCreateResourceIdentifier`

GetDefaultScreenScheme returns the DefaultScreenScheme field if non-nil, zero value otherwise.

### GetDefaultScreenSchemeOk

`func (o *IssueTypeScreenSchemePayload) GetDefaultScreenSchemeOk() (*ProjectCreateResourceIdentifier, bool)`

GetDefaultScreenSchemeOk returns a tuple with the DefaultScreenScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScreenScheme

`func (o *IssueTypeScreenSchemePayload) SetDefaultScreenScheme(v ProjectCreateResourceIdentifier)`

SetDefaultScreenScheme sets DefaultScreenScheme field to given value.

### HasDefaultScreenScheme

`func (o *IssueTypeScreenSchemePayload) HasDefaultScreenScheme() bool`

HasDefaultScreenScheme returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeScreenSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeScreenSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeScreenSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeScreenSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExplicitMappings

`func (o *IssueTypeScreenSchemePayload) GetExplicitMappings() map[string]ProjectCreateResourceIdentifier`

GetExplicitMappings returns the ExplicitMappings field if non-nil, zero value otherwise.

### GetExplicitMappingsOk

`func (o *IssueTypeScreenSchemePayload) GetExplicitMappingsOk() (*map[string]ProjectCreateResourceIdentifier, bool)`

GetExplicitMappingsOk returns a tuple with the ExplicitMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMappings

`func (o *IssueTypeScreenSchemePayload) SetExplicitMappings(v map[string]ProjectCreateResourceIdentifier)`

SetExplicitMappings sets ExplicitMappings field to given value.

### HasExplicitMappings

`func (o *IssueTypeScreenSchemePayload) HasExplicitMappings() bool`

HasExplicitMappings returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeScreenSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeScreenSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeScreenSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeScreenSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *IssueTypeScreenSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *IssueTypeScreenSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *IssueTypeScreenSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *IssueTypeScreenSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


