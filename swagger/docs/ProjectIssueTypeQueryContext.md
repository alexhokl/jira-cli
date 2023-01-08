# ProjectIssueTypeQueryContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypes** | Pointer to **[]string** | The set of issue type IDs. | [optional] 
**Project** | Pointer to **string** | The ID of the project. | [optional] 

## Methods

### NewProjectIssueTypeQueryContext

`func NewProjectIssueTypeQueryContext() *ProjectIssueTypeQueryContext`

NewProjectIssueTypeQueryContext instantiates a new ProjectIssueTypeQueryContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueTypeQueryContextWithDefaults

`func NewProjectIssueTypeQueryContextWithDefaults() *ProjectIssueTypeQueryContext`

NewProjectIssueTypeQueryContextWithDefaults instantiates a new ProjectIssueTypeQueryContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypes

`func (o *ProjectIssueTypeQueryContext) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *ProjectIssueTypeQueryContext) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *ProjectIssueTypeQueryContext) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *ProjectIssueTypeQueryContext) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetProject

`func (o *ProjectIssueTypeQueryContext) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectIssueTypeQueryContext) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectIssueTypeQueryContext) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectIssueTypeQueryContext) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


