# IssueTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The ID of the issue type&#39;s avatar. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the issue type. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Unique ID for next-gen projects. | [optional] [readonly] 
**HierarchyLevel** | Pointer to **int32** | Hierarchy level of the issue type. | [optional] [readonly] 
**IconUrl** | Pointer to **string** | The URL of the issue type&#39;s avatar. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the issue type. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the issue type. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | Details of the next-gen projects the issue type is available in. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these issue type details. | [optional] [readonly] 
**Subtask** | Pointer to **bool** | Whether this issue type is used to create subtasks. | [optional] [readonly] 

## Methods

### NewIssueTypeDetails

`func NewIssueTypeDetails() *IssueTypeDetails`

NewIssueTypeDetails instantiates a new IssueTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeDetailsWithDefaults

`func NewIssueTypeDetailsWithDefaults() *IssueTypeDetails`

NewIssueTypeDetailsWithDefaults instantiates a new IssueTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *IssueTypeDetails) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *IssueTypeDetails) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *IssueTypeDetails) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *IssueTypeDetails) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityId

`func (o *IssueTypeDetails) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IssueTypeDetails) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IssueTypeDetails) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IssueTypeDetails) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *IssueTypeDetails) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *IssueTypeDetails) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *IssueTypeDetails) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *IssueTypeDetails) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetIconUrl

`func (o *IssueTypeDetails) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *IssueTypeDetails) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *IssueTypeDetails) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *IssueTypeDetails) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *IssueTypeDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueTypeDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *IssueTypeDetails) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IssueTypeDetails) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IssueTypeDetails) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IssueTypeDetails) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *IssueTypeDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueTypeDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueTypeDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IssueTypeDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubtask

`func (o *IssueTypeDetails) GetSubtask() bool`

GetSubtask returns the Subtask field if non-nil, zero value otherwise.

### GetSubtaskOk

`func (o *IssueTypeDetails) GetSubtaskOk() (*bool, bool)`

GetSubtaskOk returns a tuple with the Subtask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtask

`func (o *IssueTypeDetails) SetSubtask(v bool)`

SetSubtask sets Subtask field to given value.

### HasSubtask

`func (o *IssueTypeDetails) HasSubtask() bool`

HasSubtask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


