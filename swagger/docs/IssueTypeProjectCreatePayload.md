# IssueTypeProjectCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeHierarchy** | Pointer to [**[]IssueTypeHierarchyPayload**](IssueTypeHierarchyPayload.md) | Defines the issue type hierarhy to be created and used during this project creation. This will only add new levels if there isn&#39;t an existing level | [optional] 
**IssueTypeScheme** | Pointer to [**IssueTypeSchemePayload**](IssueTypeSchemePayload.md) |  | [optional] 
**IssueTypes** | Pointer to [**[]IssueTypePayload**](IssueTypePayload.md) | Only needed if you want to create issue types, you can otherwise use the ids of issue types in the scheme configuration | [optional] 

## Methods

### NewIssueTypeProjectCreatePayload

`func NewIssueTypeProjectCreatePayload() *IssueTypeProjectCreatePayload`

NewIssueTypeProjectCreatePayload instantiates a new IssueTypeProjectCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeProjectCreatePayloadWithDefaults

`func NewIssueTypeProjectCreatePayloadWithDefaults() *IssueTypeProjectCreatePayload`

NewIssueTypeProjectCreatePayloadWithDefaults instantiates a new IssueTypeProjectCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeHierarchy

`func (o *IssueTypeProjectCreatePayload) GetIssueTypeHierarchy() []IssueTypeHierarchyPayload`

GetIssueTypeHierarchy returns the IssueTypeHierarchy field if non-nil, zero value otherwise.

### GetIssueTypeHierarchyOk

`func (o *IssueTypeProjectCreatePayload) GetIssueTypeHierarchyOk() (*[]IssueTypeHierarchyPayload, bool)`

GetIssueTypeHierarchyOk returns a tuple with the IssueTypeHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeHierarchy

`func (o *IssueTypeProjectCreatePayload) SetIssueTypeHierarchy(v []IssueTypeHierarchyPayload)`

SetIssueTypeHierarchy sets IssueTypeHierarchy field to given value.

### HasIssueTypeHierarchy

`func (o *IssueTypeProjectCreatePayload) HasIssueTypeHierarchy() bool`

HasIssueTypeHierarchy returns a boolean if a field has been set.

### SetIssueTypeHierarchyNil

`func (o *IssueTypeProjectCreatePayload) SetIssueTypeHierarchyNil(b bool)`

 SetIssueTypeHierarchyNil sets the value for IssueTypeHierarchy to be an explicit nil

### UnsetIssueTypeHierarchy
`func (o *IssueTypeProjectCreatePayload) UnsetIssueTypeHierarchy()`

UnsetIssueTypeHierarchy ensures that no value is present for IssueTypeHierarchy, not even an explicit nil
### GetIssueTypeScheme

`func (o *IssueTypeProjectCreatePayload) GetIssueTypeScheme() IssueTypeSchemePayload`

GetIssueTypeScheme returns the IssueTypeScheme field if non-nil, zero value otherwise.

### GetIssueTypeSchemeOk

`func (o *IssueTypeProjectCreatePayload) GetIssueTypeSchemeOk() (*IssueTypeSchemePayload, bool)`

GetIssueTypeSchemeOk returns a tuple with the IssueTypeScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScheme

`func (o *IssueTypeProjectCreatePayload) SetIssueTypeScheme(v IssueTypeSchemePayload)`

SetIssueTypeScheme sets IssueTypeScheme field to given value.

### HasIssueTypeScheme

`func (o *IssueTypeProjectCreatePayload) HasIssueTypeScheme() bool`

HasIssueTypeScheme returns a boolean if a field has been set.

### GetIssueTypes

`func (o *IssueTypeProjectCreatePayload) GetIssueTypes() []IssueTypePayload`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *IssueTypeProjectCreatePayload) GetIssueTypesOk() (*[]IssueTypePayload, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *IssueTypeProjectCreatePayload) SetIssueTypes(v []IssueTypePayload)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *IssueTypeProjectCreatePayload) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### SetIssueTypesNil

`func (o *IssueTypeProjectCreatePayload) SetIssueTypesNil(b bool)`

 SetIssueTypesNil sets the value for IssueTypes to be an explicit nil

### UnsetIssueTypes
`func (o *IssueTypeProjectCreatePayload) UnsetIssueTypes()`

UnsetIssueTypes ensures that no value is present for IssueTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


