# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | Pointer to [**[]VersionApprover**](VersionApprover.md) | If the expand option &#x60;approvers&#x60; is used, returns a list containing the approvers for this version. | [optional] [readonly] 
**Archived** | Pointer to **bool** | Indicates that the version is archived. Optional when creating or updating a version. | [optional] 
**Description** | Pointer to **string** | The description of the version. Optional when creating or updating a version. The maximum size is 16,384 bytes. | [optional] 
**Driver** | Pointer to **string** | If the expand option &#x60;driver&#x60; is used, returns the Atlassian account ID of the driver. | [optional] [readonly] 
**Expand** | Pointer to **string** | Use [expand](em&gt;#expansion) to include additional information about version in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;operations&#x60; Returns the list of operations available for this version.  *  &#x60;issuesstatus&#x60; Returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*.  *  &#x60;driver&#x60; Returns the Atlassian account ID of the version driver.  *  &#x60;approvers&#x60; Returns a list containing approvers for this version.  Optional for create and update. | [optional] 
**Id** | Pointer to **string** | The ID of the version. | [optional] [readonly] 
**IssuesStatusForFixVersion** | Pointer to [**VersionIssuesStatus**](VersionIssuesStatus.md) | If the expand option &#x60;issuesstatus&#x60; is used, returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*. | [optional] [readonly] 
**MoveUnfixedIssuesTo** | Pointer to **string** | The URL of the self link to the version to which all unfixed issues are moved when a version is released. Not applicable when creating a version. Optional when updating a version. | [optional] 
**Name** | Pointer to **string** | The unique name of the version. Required when creating a version. Optional when updating a version. The maximum length is 255 characters. | [optional] 
**Operations** | Pointer to [**[]SimpleLink**](SimpleLink.md) | If the expand option &#x60;operations&#x60; is used, returns the list of operations available for this version. | [optional] [readonly] 
**Overdue** | Pointer to **bool** | Indicates that the version is overdue. | [optional] [readonly] 
**Project** | Pointer to **string** | Deprecated. Use &#x60;projectId&#x60;. | [optional] 
**ProjectId** | Pointer to **int64** | The ID of the project to which this version is attached. Required when creating a version. Not applicable when updating a version. | [optional] 
**ReleaseDate** | Pointer to **string** | The release date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version. | [optional] 
**Released** | Pointer to **bool** | Indicates that the version is released. If the version is released a request to release again is ignored. Not applicable when creating a version. Optional when updating a version. | [optional] 
**Self** | Pointer to **string** | The URL of the version. | [optional] [readonly] 
**StartDate** | Pointer to **string** | The start date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version. | [optional] 
**UserReleaseDate** | Pointer to **string** | The date on which work on this version is expected to finish, expressed in the instance&#39;s *Day/Month/Year Format* date format. | [optional] [readonly] 
**UserStartDate** | Pointer to **string** | The date on which work on this version is expected to start, expressed in the instance&#39;s *Day/Month/Year Format* date format. | [optional] [readonly] 

## Methods

### NewVersion

`func NewVersion() *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *Version) GetApprovers() []VersionApprover`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *Version) GetApproversOk() (*[]VersionApprover, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *Version) SetApprovers(v []VersionApprover)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *Version) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetArchived

`func (o *Version) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Version) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Version) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Version) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDescription

`func (o *Version) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Version) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Version) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Version) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDriver

`func (o *Version) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *Version) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *Version) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *Version) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetExpand

`func (o *Version) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *Version) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *Version) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *Version) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetId

`func (o *Version) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Version) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Version) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Version) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuesStatusForFixVersion

`func (o *Version) GetIssuesStatusForFixVersion() VersionIssuesStatus`

GetIssuesStatusForFixVersion returns the IssuesStatusForFixVersion field if non-nil, zero value otherwise.

### GetIssuesStatusForFixVersionOk

`func (o *Version) GetIssuesStatusForFixVersionOk() (*VersionIssuesStatus, bool)`

GetIssuesStatusForFixVersionOk returns a tuple with the IssuesStatusForFixVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesStatusForFixVersion

`func (o *Version) SetIssuesStatusForFixVersion(v VersionIssuesStatus)`

SetIssuesStatusForFixVersion sets IssuesStatusForFixVersion field to given value.

### HasIssuesStatusForFixVersion

`func (o *Version) HasIssuesStatusForFixVersion() bool`

HasIssuesStatusForFixVersion returns a boolean if a field has been set.

### GetMoveUnfixedIssuesTo

`func (o *Version) GetMoveUnfixedIssuesTo() string`

GetMoveUnfixedIssuesTo returns the MoveUnfixedIssuesTo field if non-nil, zero value otherwise.

### GetMoveUnfixedIssuesToOk

`func (o *Version) GetMoveUnfixedIssuesToOk() (*string, bool)`

GetMoveUnfixedIssuesToOk returns a tuple with the MoveUnfixedIssuesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveUnfixedIssuesTo

`func (o *Version) SetMoveUnfixedIssuesTo(v string)`

SetMoveUnfixedIssuesTo sets MoveUnfixedIssuesTo field to given value.

### HasMoveUnfixedIssuesTo

`func (o *Version) HasMoveUnfixedIssuesTo() bool`

HasMoveUnfixedIssuesTo returns a boolean if a field has been set.

### GetName

`func (o *Version) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Version) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Version) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Version) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperations

`func (o *Version) GetOperations() []SimpleLink`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Version) GetOperationsOk() (*[]SimpleLink, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Version) SetOperations(v []SimpleLink)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Version) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetOverdue

`func (o *Version) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *Version) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *Version) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *Version) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetProject

`func (o *Version) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Version) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Version) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Version) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *Version) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Version) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Version) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Version) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReleaseDate

`func (o *Version) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Version) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Version) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *Version) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleased

`func (o *Version) GetReleased() bool`

GetReleased returns the Released field if non-nil, zero value otherwise.

### GetReleasedOk

`func (o *Version) GetReleasedOk() (*bool, bool)`

GetReleasedOk returns a tuple with the Released field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleased

`func (o *Version) SetReleased(v bool)`

SetReleased sets Released field to given value.

### HasReleased

`func (o *Version) HasReleased() bool`

HasReleased returns a boolean if a field has been set.

### GetSelf

`func (o *Version) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Version) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Version) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Version) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartDate

`func (o *Version) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Version) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Version) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Version) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUserReleaseDate

`func (o *Version) GetUserReleaseDate() string`

GetUserReleaseDate returns the UserReleaseDate field if non-nil, zero value otherwise.

### GetUserReleaseDateOk

`func (o *Version) GetUserReleaseDateOk() (*string, bool)`

GetUserReleaseDateOk returns a tuple with the UserReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReleaseDate

`func (o *Version) SetUserReleaseDate(v string)`

SetUserReleaseDate sets UserReleaseDate field to given value.

### HasUserReleaseDate

`func (o *Version) HasUserReleaseDate() bool`

HasUserReleaseDate returns a boolean if a field has been set.

### GetUserStartDate

`func (o *Version) GetUserStartDate() string`

GetUserStartDate returns the UserStartDate field if non-nil, zero value otherwise.

### GetUserStartDateOk

`func (o *Version) GetUserStartDateOk() (*string, bool)`

GetUserStartDateOk returns a tuple with the UserStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStartDate

`func (o *Version) SetUserStartDate(v string)`

SetUserStartDate sets UserStartDate field to given value.

### HasUserStartDate

`func (o *Version) HasUserStartDate() bool`

HasUserStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


