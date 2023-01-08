# IssueTypePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **NullableInt64** | The avatar ID of the issue type. Go to https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-avatars/\\#api-rest-api-3-avatar-type-system-get to choose an avatarId existing in Jira | [optional] 
**Description** | Pointer to **NullableString** | The description of the issue type | [optional] 
**HierarchyLevel** | Pointer to **int32** | The hierarchy level of the issue type. 0, 1, 2, 3 .. n; Negative values for subtasks | [optional] 
**Name** | Pointer to **string** | The name of the issue type | [optional] 
**OnConflict** | Pointer to **string** | The conflict strategy to use when the issue type already exists. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewIssueTypePayload

`func NewIssueTypePayload() *IssueTypePayload`

NewIssueTypePayload instantiates a new IssueTypePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypePayloadWithDefaults

`func NewIssueTypePayloadWithDefaults() *IssueTypePayload`

NewIssueTypePayloadWithDefaults instantiates a new IssueTypePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *IssueTypePayload) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *IssueTypePayload) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *IssueTypePayload) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *IssueTypePayload) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### SetAvatarIdNil

`func (o *IssueTypePayload) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *IssueTypePayload) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil
### GetDescription

`func (o *IssueTypePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IssueTypePayload) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IssueTypePayload) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHierarchyLevel

`func (o *IssueTypePayload) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *IssueTypePayload) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *IssueTypePayload) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *IssueTypePayload) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetName

`func (o *IssueTypePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *IssueTypePayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *IssueTypePayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *IssueTypePayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *IssueTypePayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *IssueTypePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *IssueTypePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *IssueTypePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *IssueTypePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


