# IssueTypeHierarchyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HierarchyLevel** | Pointer to **int32** | The hierarchy level of the issue type. 0, 1, 2, 3 .. n; Negative values for subtasks | [optional] 
**Name** | Pointer to **string** | The name of the issue type | [optional] 
**OnConflict** | Pointer to **string** | The conflict strategy to use when the issue type already exists. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewIssueTypeHierarchyPayload

`func NewIssueTypeHierarchyPayload() *IssueTypeHierarchyPayload`

NewIssueTypeHierarchyPayload instantiates a new IssueTypeHierarchyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeHierarchyPayloadWithDefaults

`func NewIssueTypeHierarchyPayloadWithDefaults() *IssueTypeHierarchyPayload`

NewIssueTypeHierarchyPayloadWithDefaults instantiates a new IssueTypeHierarchyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHierarchyLevel

`func (o *IssueTypeHierarchyPayload) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *IssueTypeHierarchyPayload) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *IssueTypeHierarchyPayload) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *IssueTypeHierarchyPayload) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeHierarchyPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeHierarchyPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeHierarchyPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeHierarchyPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *IssueTypeHierarchyPayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *IssueTypeHierarchyPayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *IssueTypeHierarchyPayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *IssueTypeHierarchyPayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *IssueTypeHierarchyPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *IssueTypeHierarchyPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *IssueTypeHierarchyPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *IssueTypeHierarchyPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


