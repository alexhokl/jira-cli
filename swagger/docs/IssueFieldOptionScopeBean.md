# IssueFieldOptionScopeBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**GlobalScopeBean**](GlobalScopeBean.md) | Defines the behavior of the option within the global context. If this property is set, even if set to an empty object, then the option is available in all projects. | [optional] 
**Projects** | Pointer to **[]int64** | DEPRECATED | [optional] 
**Projects2** | Pointer to [**[]ProjectScopeBean**](ProjectScopeBean.md) | Defines the projects in which the option is available and the behavior of the option within each project. Specify one object per project. The behavior of the option in a project context overrides the behavior in the global context. | [optional] 

## Methods

### NewIssueFieldOptionScopeBean

`func NewIssueFieldOptionScopeBean() *IssueFieldOptionScopeBean`

NewIssueFieldOptionScopeBean instantiates a new IssueFieldOptionScopeBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldOptionScopeBeanWithDefaults

`func NewIssueFieldOptionScopeBeanWithDefaults() *IssueFieldOptionScopeBean`

NewIssueFieldOptionScopeBeanWithDefaults instantiates a new IssueFieldOptionScopeBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *IssueFieldOptionScopeBean) GetGlobal() GlobalScopeBean`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *IssueFieldOptionScopeBean) GetGlobalOk() (*GlobalScopeBean, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *IssueFieldOptionScopeBean) SetGlobal(v GlobalScopeBean)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *IssueFieldOptionScopeBean) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetProjects

`func (o *IssueFieldOptionScopeBean) GetProjects() []int64`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *IssueFieldOptionScopeBean) GetProjectsOk() (*[]int64, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *IssueFieldOptionScopeBean) SetProjects(v []int64)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *IssueFieldOptionScopeBean) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetProjects2

`func (o *IssueFieldOptionScopeBean) GetProjects2() []ProjectScopeBean`

GetProjects2 returns the Projects2 field if non-nil, zero value otherwise.

### GetProjects2Ok

`func (o *IssueFieldOptionScopeBean) GetProjects2Ok() (*[]ProjectScopeBean, bool)`

GetProjects2Ok returns a tuple with the Projects2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects2

`func (o *IssueFieldOptionScopeBean) SetProjects2(v []ProjectScopeBean)`

SetProjects2 sets Projects2 field to given value.

### HasProjects2

`func (o *IssueFieldOptionScopeBean) HasProjects2() bool`

HasProjects2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


