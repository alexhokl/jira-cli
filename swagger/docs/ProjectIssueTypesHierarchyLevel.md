# ProjectIssueTypesHierarchyLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The ID of the issue type hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] [readonly] 
**IssueTypes** | Pointer to [**[]IssueTypeInfo**](IssueTypeInfo.md) | The list of issue types in the hierarchy level. | [optional] [readonly] 
**Level** | Pointer to **int32** | The level of the issue type hierarchy level. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the issue type hierarchy level. | [optional] [readonly] 

## Methods

### NewProjectIssueTypesHierarchyLevel

`func NewProjectIssueTypesHierarchyLevel() *ProjectIssueTypesHierarchyLevel`

NewProjectIssueTypesHierarchyLevel instantiates a new ProjectIssueTypesHierarchyLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueTypesHierarchyLevelWithDefaults

`func NewProjectIssueTypesHierarchyLevelWithDefaults() *ProjectIssueTypesHierarchyLevel`

NewProjectIssueTypesHierarchyLevelWithDefaults instantiates a new ProjectIssueTypesHierarchyLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *ProjectIssueTypesHierarchyLevel) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ProjectIssueTypesHierarchyLevel) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ProjectIssueTypesHierarchyLevel) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ProjectIssueTypesHierarchyLevel) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetIssueTypes

`func (o *ProjectIssueTypesHierarchyLevel) GetIssueTypes() []IssueTypeInfo`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *ProjectIssueTypesHierarchyLevel) GetIssueTypesOk() (*[]IssueTypeInfo, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *ProjectIssueTypesHierarchyLevel) SetIssueTypes(v []IssueTypeInfo)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *ProjectIssueTypesHierarchyLevel) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetLevel

`func (o *ProjectIssueTypesHierarchyLevel) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ProjectIssueTypesHierarchyLevel) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ProjectIssueTypesHierarchyLevel) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ProjectIssueTypesHierarchyLevel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *ProjectIssueTypesHierarchyLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectIssueTypesHierarchyLevel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectIssueTypesHierarchyLevel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectIssueTypesHierarchyLevel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


