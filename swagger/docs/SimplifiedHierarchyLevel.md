# SimplifiedHierarchyLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AboveLevelId** | Pointer to **int64** | The ID of the level above this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 
**BelowLevelId** | Pointer to **int64** | The ID of the level below this one in the hierarchy. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 
**ExternalUuid** | Pointer to **string** | The external UUID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 
**HierarchyLevelNumber** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int64** | The ID of the hierarchy level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 
**IssueTypeIds** | Pointer to **[]int64** | The issue types available in this hierarchy level. | [optional] 
**Level** | Pointer to **int32** | The level of this item in the hierarchy. | [optional] 
**Name** | Pointer to **string** | The name of this hierarchy level. | [optional] 
**ProjectConfigurationId** | Pointer to **int64** | The ID of the project configuration. This property is deprecated, see [Change oticen: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 

## Methods

### NewSimplifiedHierarchyLevel

`func NewSimplifiedHierarchyLevel() *SimplifiedHierarchyLevel`

NewSimplifiedHierarchyLevel instantiates a new SimplifiedHierarchyLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedHierarchyLevelWithDefaults

`func NewSimplifiedHierarchyLevelWithDefaults() *SimplifiedHierarchyLevel`

NewSimplifiedHierarchyLevelWithDefaults instantiates a new SimplifiedHierarchyLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAboveLevelId

`func (o *SimplifiedHierarchyLevel) GetAboveLevelId() int64`

GetAboveLevelId returns the AboveLevelId field if non-nil, zero value otherwise.

### GetAboveLevelIdOk

`func (o *SimplifiedHierarchyLevel) GetAboveLevelIdOk() (*int64, bool)`

GetAboveLevelIdOk returns a tuple with the AboveLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAboveLevelId

`func (o *SimplifiedHierarchyLevel) SetAboveLevelId(v int64)`

SetAboveLevelId sets AboveLevelId field to given value.

### HasAboveLevelId

`func (o *SimplifiedHierarchyLevel) HasAboveLevelId() bool`

HasAboveLevelId returns a boolean if a field has been set.

### GetBelowLevelId

`func (o *SimplifiedHierarchyLevel) GetBelowLevelId() int64`

GetBelowLevelId returns the BelowLevelId field if non-nil, zero value otherwise.

### GetBelowLevelIdOk

`func (o *SimplifiedHierarchyLevel) GetBelowLevelIdOk() (*int64, bool)`

GetBelowLevelIdOk returns a tuple with the BelowLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowLevelId

`func (o *SimplifiedHierarchyLevel) SetBelowLevelId(v int64)`

SetBelowLevelId sets BelowLevelId field to given value.

### HasBelowLevelId

`func (o *SimplifiedHierarchyLevel) HasBelowLevelId() bool`

HasBelowLevelId returns a boolean if a field has been set.

### GetExternalUuid

`func (o *SimplifiedHierarchyLevel) GetExternalUuid() string`

GetExternalUuid returns the ExternalUuid field if non-nil, zero value otherwise.

### GetExternalUuidOk

`func (o *SimplifiedHierarchyLevel) GetExternalUuidOk() (*string, bool)`

GetExternalUuidOk returns a tuple with the ExternalUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUuid

`func (o *SimplifiedHierarchyLevel) SetExternalUuid(v string)`

SetExternalUuid sets ExternalUuid field to given value.

### HasExternalUuid

`func (o *SimplifiedHierarchyLevel) HasExternalUuid() bool`

HasExternalUuid returns a boolean if a field has been set.

### GetHierarchyLevelNumber

`func (o *SimplifiedHierarchyLevel) GetHierarchyLevelNumber() int32`

GetHierarchyLevelNumber returns the HierarchyLevelNumber field if non-nil, zero value otherwise.

### GetHierarchyLevelNumberOk

`func (o *SimplifiedHierarchyLevel) GetHierarchyLevelNumberOk() (*int32, bool)`

GetHierarchyLevelNumberOk returns a tuple with the HierarchyLevelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevelNumber

`func (o *SimplifiedHierarchyLevel) SetHierarchyLevelNumber(v int32)`

SetHierarchyLevelNumber sets HierarchyLevelNumber field to given value.

### HasHierarchyLevelNumber

`func (o *SimplifiedHierarchyLevel) HasHierarchyLevelNumber() bool`

HasHierarchyLevelNumber returns a boolean if a field has been set.

### GetId

`func (o *SimplifiedHierarchyLevel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedHierarchyLevel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedHierarchyLevel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedHierarchyLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *SimplifiedHierarchyLevel) GetIssueTypeIds() []int64`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *SimplifiedHierarchyLevel) GetIssueTypeIdsOk() (*[]int64, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *SimplifiedHierarchyLevel) SetIssueTypeIds(v []int64)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *SimplifiedHierarchyLevel) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetLevel

`func (o *SimplifiedHierarchyLevel) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SimplifiedHierarchyLevel) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SimplifiedHierarchyLevel) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SimplifiedHierarchyLevel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedHierarchyLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedHierarchyLevel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedHierarchyLevel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedHierarchyLevel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectConfigurationId

`func (o *SimplifiedHierarchyLevel) GetProjectConfigurationId() int64`

GetProjectConfigurationId returns the ProjectConfigurationId field if non-nil, zero value otherwise.

### GetProjectConfigurationIdOk

`func (o *SimplifiedHierarchyLevel) GetProjectConfigurationIdOk() (*int64, bool)`

GetProjectConfigurationIdOk returns a tuple with the ProjectConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectConfigurationId

`func (o *SimplifiedHierarchyLevel) SetProjectConfigurationId(v int64)`

SetProjectConfigurationId sets ProjectConfigurationId field to given value.

### HasProjectConfigurationId

`func (o *SimplifiedHierarchyLevel) HasProjectConfigurationId() bool`

HasProjectConfigurationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


