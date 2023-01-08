# IssueTypeIssueCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The ID of the issue type&#39;s avatar. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the issue type. | [optional] [readonly] 
**EntityId** | Pointer to **string** | Unique ID for next-gen projects. | [optional] [readonly] 
**Expand** | Pointer to **string** | Expand options that include additional issue type metadata details in the response. | [optional] [readonly] 
**Fields** | Pointer to [**map[string]FieldMetadata**](FieldMetadata.md) | List of the fields available when creating an issue for the issue type. | [optional] [readonly] 
**HierarchyLevel** | Pointer to **int32** | Hierarchy level of the issue type. | [optional] [readonly] 
**IconUrl** | Pointer to **string** | The URL of the issue type&#39;s avatar. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the issue type. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the issue type. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | Details of the next-gen projects the issue type is available in. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these issue type details. | [optional] [readonly] 
**Subtask** | Pointer to **bool** | Whether this issue type is used to create subtasks. | [optional] [readonly] 

## Methods

### NewIssueTypeIssueCreateMetadata

`func NewIssueTypeIssueCreateMetadata() *IssueTypeIssueCreateMetadata`

NewIssueTypeIssueCreateMetadata instantiates a new IssueTypeIssueCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeIssueCreateMetadataWithDefaults

`func NewIssueTypeIssueCreateMetadataWithDefaults() *IssueTypeIssueCreateMetadata`

NewIssueTypeIssueCreateMetadataWithDefaults instantiates a new IssueTypeIssueCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *IssueTypeIssueCreateMetadata) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *IssueTypeIssueCreateMetadata) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *IssueTypeIssueCreateMetadata) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *IssueTypeIssueCreateMetadata) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeIssueCreateMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeIssueCreateMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeIssueCreateMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeIssueCreateMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityId

`func (o *IssueTypeIssueCreateMetadata) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IssueTypeIssueCreateMetadata) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IssueTypeIssueCreateMetadata) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *IssueTypeIssueCreateMetadata) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetExpand

`func (o *IssueTypeIssueCreateMetadata) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *IssueTypeIssueCreateMetadata) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *IssueTypeIssueCreateMetadata) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *IssueTypeIssueCreateMetadata) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFields

`func (o *IssueTypeIssueCreateMetadata) GetFields() map[string]FieldMetadata`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IssueTypeIssueCreateMetadata) GetFieldsOk() (*map[string]FieldMetadata, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IssueTypeIssueCreateMetadata) SetFields(v map[string]FieldMetadata)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IssueTypeIssueCreateMetadata) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *IssueTypeIssueCreateMetadata) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *IssueTypeIssueCreateMetadata) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *IssueTypeIssueCreateMetadata) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *IssueTypeIssueCreateMetadata) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetIconUrl

`func (o *IssueTypeIssueCreateMetadata) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *IssueTypeIssueCreateMetadata) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *IssueTypeIssueCreateMetadata) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *IssueTypeIssueCreateMetadata) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *IssueTypeIssueCreateMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeIssueCreateMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeIssueCreateMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueTypeIssueCreateMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeIssueCreateMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeIssueCreateMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeIssueCreateMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeIssueCreateMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *IssueTypeIssueCreateMetadata) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IssueTypeIssueCreateMetadata) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IssueTypeIssueCreateMetadata) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IssueTypeIssueCreateMetadata) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *IssueTypeIssueCreateMetadata) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IssueTypeIssueCreateMetadata) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IssueTypeIssueCreateMetadata) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IssueTypeIssueCreateMetadata) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubtask

`func (o *IssueTypeIssueCreateMetadata) GetSubtask() bool`

GetSubtask returns the Subtask field if non-nil, zero value otherwise.

### GetSubtaskOk

`func (o *IssueTypeIssueCreateMetadata) GetSubtaskOk() (*bool, bool)`

GetSubtaskOk returns a tuple with the Subtask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtask

`func (o *IssueTypeIssueCreateMetadata) SetSubtask(v bool)`

SetSubtask sets Subtask field to given value.

### HasSubtask

`func (o *IssueTypeIssueCreateMetadata) HasSubtask() bool`

HasSubtask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


