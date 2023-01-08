# IssueTypeSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIssueTypeId** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Description** | Pointer to **NullableString** | The description of the issue type scheme | [optional] 
**IssueTypeIds** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | The issue type IDs for the issue type scheme | [optional] 
**Name** | Pointer to **string** | The name of the issue type scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewIssueTypeSchemePayload

`func NewIssueTypeSchemePayload() *IssueTypeSchemePayload`

NewIssueTypeSchemePayload instantiates a new IssueTypeSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeSchemePayloadWithDefaults

`func NewIssueTypeSchemePayloadWithDefaults() *IssueTypeSchemePayload`

NewIssueTypeSchemePayloadWithDefaults instantiates a new IssueTypeSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultIssueTypeId

`func (o *IssueTypeSchemePayload) GetDefaultIssueTypeId() ProjectCreateResourceIdentifier`

GetDefaultIssueTypeId returns the DefaultIssueTypeId field if non-nil, zero value otherwise.

### GetDefaultIssueTypeIdOk

`func (o *IssueTypeSchemePayload) GetDefaultIssueTypeIdOk() (*ProjectCreateResourceIdentifier, bool)`

GetDefaultIssueTypeIdOk returns a tuple with the DefaultIssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueTypeId

`func (o *IssueTypeSchemePayload) SetDefaultIssueTypeId(v ProjectCreateResourceIdentifier)`

SetDefaultIssueTypeId sets DefaultIssueTypeId field to given value.

### HasDefaultIssueTypeId

`func (o *IssueTypeSchemePayload) HasDefaultIssueTypeId() bool`

HasDefaultIssueTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IssueTypeSchemePayload) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IssueTypeSchemePayload) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIssueTypeIds

`func (o *IssueTypeSchemePayload) GetIssueTypeIds() []ProjectCreateResourceIdentifier`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *IssueTypeSchemePayload) GetIssueTypeIdsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *IssueTypeSchemePayload) SetIssueTypeIds(v []ProjectCreateResourceIdentifier)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *IssueTypeSchemePayload) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *IssueTypeSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *IssueTypeSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *IssueTypeSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *IssueTypeSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


