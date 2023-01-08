# IssueTypeCreateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the issue type. | [optional] 
**HierarchyLevel** | Pointer to **int32** | The hierarchy level of the issue type. Use:   *  &#x60;-1&#x60; for Subtask.  *  &#x60;0&#x60; for Base.  Defaults to &#x60;0&#x60;. | [optional] 
**Name** | **string** | The unique name for the issue type. The maximum length is 60 characters. | 
**Type** | Pointer to **string** | Deprecated. Use &#x60;hierarchyLevel&#x60; instead. See the [deprecation notice](https://community.developer.atlassian.com/t/deprecation-of-the-epic-link-parent-link-and-other-related-fields-in-rest-apis-and-webhooks/54048) for details.  Whether the issue type is &#x60;subtype&#x60; or &#x60;standard&#x60;. Defaults to &#x60;standard&#x60;. | [optional] 

## Methods

### NewIssueTypeCreateBean

`func NewIssueTypeCreateBean(name string, ) *IssueTypeCreateBean`

NewIssueTypeCreateBean instantiates a new IssueTypeCreateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeCreateBeanWithDefaults

`func NewIssueTypeCreateBeanWithDefaults() *IssueTypeCreateBean`

NewIssueTypeCreateBeanWithDefaults instantiates a new IssueTypeCreateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueTypeCreateBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeCreateBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeCreateBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeCreateBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *IssueTypeCreateBean) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *IssueTypeCreateBean) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *IssueTypeCreateBean) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *IssueTypeCreateBean) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeCreateBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeCreateBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeCreateBean) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IssueTypeCreateBean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueTypeCreateBean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueTypeCreateBean) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueTypeCreateBean) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


