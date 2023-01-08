# Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIsSubtask** | Pointer to [**Error**](Error.md) |  | [optional] 
**IssuesInArchivedProjects** | Pointer to [**Error**](Error.md) |  | [optional] 
**IssuesInUnlicensedProjects** | Pointer to [**Error**](Error.md) |  | [optional] 
**IssuesNotFound** | Pointer to [**Error**](Error.md) |  | [optional] 
**UserDoesNotHavePermission** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewErrors

`func NewErrors() *Errors`

NewErrors instantiates a new Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsWithDefaults

`func NewErrorsWithDefaults() *Errors`

NewErrorsWithDefaults instantiates a new Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueIsSubtask

`func (o *Errors) GetIssueIsSubtask() Error`

GetIssueIsSubtask returns the IssueIsSubtask field if non-nil, zero value otherwise.

### GetIssueIsSubtaskOk

`func (o *Errors) GetIssueIsSubtaskOk() (*Error, bool)`

GetIssueIsSubtaskOk returns a tuple with the IssueIsSubtask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIsSubtask

`func (o *Errors) SetIssueIsSubtask(v Error)`

SetIssueIsSubtask sets IssueIsSubtask field to given value.

### HasIssueIsSubtask

`func (o *Errors) HasIssueIsSubtask() bool`

HasIssueIsSubtask returns a boolean if a field has been set.

### GetIssuesInArchivedProjects

`func (o *Errors) GetIssuesInArchivedProjects() Error`

GetIssuesInArchivedProjects returns the IssuesInArchivedProjects field if non-nil, zero value otherwise.

### GetIssuesInArchivedProjectsOk

`func (o *Errors) GetIssuesInArchivedProjectsOk() (*Error, bool)`

GetIssuesInArchivedProjectsOk returns a tuple with the IssuesInArchivedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesInArchivedProjects

`func (o *Errors) SetIssuesInArchivedProjects(v Error)`

SetIssuesInArchivedProjects sets IssuesInArchivedProjects field to given value.

### HasIssuesInArchivedProjects

`func (o *Errors) HasIssuesInArchivedProjects() bool`

HasIssuesInArchivedProjects returns a boolean if a field has been set.

### GetIssuesInUnlicensedProjects

`func (o *Errors) GetIssuesInUnlicensedProjects() Error`

GetIssuesInUnlicensedProjects returns the IssuesInUnlicensedProjects field if non-nil, zero value otherwise.

### GetIssuesInUnlicensedProjectsOk

`func (o *Errors) GetIssuesInUnlicensedProjectsOk() (*Error, bool)`

GetIssuesInUnlicensedProjectsOk returns a tuple with the IssuesInUnlicensedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesInUnlicensedProjects

`func (o *Errors) SetIssuesInUnlicensedProjects(v Error)`

SetIssuesInUnlicensedProjects sets IssuesInUnlicensedProjects field to given value.

### HasIssuesInUnlicensedProjects

`func (o *Errors) HasIssuesInUnlicensedProjects() bool`

HasIssuesInUnlicensedProjects returns a boolean if a field has been set.

### GetIssuesNotFound

`func (o *Errors) GetIssuesNotFound() Error`

GetIssuesNotFound returns the IssuesNotFound field if non-nil, zero value otherwise.

### GetIssuesNotFoundOk

`func (o *Errors) GetIssuesNotFoundOk() (*Error, bool)`

GetIssuesNotFoundOk returns a tuple with the IssuesNotFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesNotFound

`func (o *Errors) SetIssuesNotFound(v Error)`

SetIssuesNotFound sets IssuesNotFound field to given value.

### HasIssuesNotFound

`func (o *Errors) HasIssuesNotFound() bool`

HasIssuesNotFound returns a boolean if a field has been set.

### GetUserDoesNotHavePermission

`func (o *Errors) GetUserDoesNotHavePermission() Error`

GetUserDoesNotHavePermission returns the UserDoesNotHavePermission field if non-nil, zero value otherwise.

### GetUserDoesNotHavePermissionOk

`func (o *Errors) GetUserDoesNotHavePermissionOk() (*Error, bool)`

GetUserDoesNotHavePermissionOk returns a tuple with the UserDoesNotHavePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDoesNotHavePermission

`func (o *Errors) SetUserDoesNotHavePermission(v Error)`

SetUserDoesNotHavePermission sets UserDoesNotHavePermission field to given value.

### HasUserDoesNotHavePermission

`func (o *Errors) HasUserDoesNotHavePermission() bool`

HasUserDoesNotHavePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


