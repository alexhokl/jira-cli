# ProjectIssueTypeHierarchy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hierarchy** | Pointer to [**[]ProjectIssueTypesHierarchyLevel**](ProjectIssueTypesHierarchyLevel.md) | Details of an issue type hierarchy level. | [optional] [readonly] 
**ProjectId** | Pointer to **int64** | The ID of the project. | [optional] [readonly] 

## Methods

### NewProjectIssueTypeHierarchy

`func NewProjectIssueTypeHierarchy() *ProjectIssueTypeHierarchy`

NewProjectIssueTypeHierarchy instantiates a new ProjectIssueTypeHierarchy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueTypeHierarchyWithDefaults

`func NewProjectIssueTypeHierarchyWithDefaults() *ProjectIssueTypeHierarchy`

NewProjectIssueTypeHierarchyWithDefaults instantiates a new ProjectIssueTypeHierarchy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHierarchy

`func (o *ProjectIssueTypeHierarchy) GetHierarchy() []ProjectIssueTypesHierarchyLevel`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *ProjectIssueTypeHierarchy) GetHierarchyOk() (*[]ProjectIssueTypesHierarchyLevel, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *ProjectIssueTypeHierarchy) SetHierarchy(v []ProjectIssueTypesHierarchyLevel)`

SetHierarchy sets Hierarchy field to given value.

### HasHierarchy

`func (o *ProjectIssueTypeHierarchy) HasHierarchy() bool`

HasHierarchy returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectIssueTypeHierarchy) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectIssueTypeHierarchy) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectIssueTypeHierarchy) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectIssueTypeHierarchy) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


