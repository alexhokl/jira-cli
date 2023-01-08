# Hierarchy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseLevelId** | Pointer to **int64** | The ID of the base level. This property is deprecated, see [Change notice: Removing hierarchy level IDs from next-gen APIs](https://developer.atlassian.com/cloud/jira/platform/change-notice-removing-hierarchy-level-ids-from-next-gen-apis/). | [optional] 
**Levels** | Pointer to [**[]SimplifiedHierarchyLevel**](SimplifiedHierarchyLevel.md) | Details about the hierarchy level. | [optional] [readonly] 

## Methods

### NewHierarchy

`func NewHierarchy() *Hierarchy`

NewHierarchy instantiates a new Hierarchy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHierarchyWithDefaults

`func NewHierarchyWithDefaults() *Hierarchy`

NewHierarchyWithDefaults instantiates a new Hierarchy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseLevelId

`func (o *Hierarchy) GetBaseLevelId() int64`

GetBaseLevelId returns the BaseLevelId field if non-nil, zero value otherwise.

### GetBaseLevelIdOk

`func (o *Hierarchy) GetBaseLevelIdOk() (*int64, bool)`

GetBaseLevelIdOk returns a tuple with the BaseLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLevelId

`func (o *Hierarchy) SetBaseLevelId(v int64)`

SetBaseLevelId sets BaseLevelId field to given value.

### HasBaseLevelId

`func (o *Hierarchy) HasBaseLevelId() bool`

HasBaseLevelId returns a boolean if a field has been set.

### GetLevels

`func (o *Hierarchy) GetLevels() []SimplifiedHierarchyLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *Hierarchy) GetLevelsOk() (*[]SimplifiedHierarchyLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *Hierarchy) SetLevels(v []SimplifiedHierarchyLevel)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *Hierarchy) HasLevels() bool`

HasLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


