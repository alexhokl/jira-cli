# JiraWorkflowPreviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Id** | Pointer to **string** | The ID of the status. | [optional] 
**Name** | Pointer to **string** | The name of the status. | [optional] 
**RawName** | Pointer to **string** | The raw name of the status. | [optional] 
**Scope** | Pointer to [**WorkflowPreviewScope**](WorkflowPreviewScope.md) |  | [optional] 
**StatusCategory** | Pointer to **string** | The category of the status. | [optional] 
**StatusReference** | Pointer to **string** | The reference of the status. Unique within this response but not guaranteed to be stable across requests. | [optional] 

## Methods

### NewJiraWorkflowPreviewStatus

`func NewJiraWorkflowPreviewStatus() *JiraWorkflowPreviewStatus`

NewJiraWorkflowPreviewStatus instantiates a new JiraWorkflowPreviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraWorkflowPreviewStatusWithDefaults

`func NewJiraWorkflowPreviewStatusWithDefaults() *JiraWorkflowPreviewStatus`

NewJiraWorkflowPreviewStatusWithDefaults instantiates a new JiraWorkflowPreviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JiraWorkflowPreviewStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JiraWorkflowPreviewStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JiraWorkflowPreviewStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JiraWorkflowPreviewStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *JiraWorkflowPreviewStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JiraWorkflowPreviewStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JiraWorkflowPreviewStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JiraWorkflowPreviewStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JiraWorkflowPreviewStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JiraWorkflowPreviewStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JiraWorkflowPreviewStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JiraWorkflowPreviewStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRawName

`func (o *JiraWorkflowPreviewStatus) GetRawName() string`

GetRawName returns the RawName field if non-nil, zero value otherwise.

### GetRawNameOk

`func (o *JiraWorkflowPreviewStatus) GetRawNameOk() (*string, bool)`

GetRawNameOk returns a tuple with the RawName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawName

`func (o *JiraWorkflowPreviewStatus) SetRawName(v string)`

SetRawName sets RawName field to given value.

### HasRawName

`func (o *JiraWorkflowPreviewStatus) HasRawName() bool`

HasRawName returns a boolean if a field has been set.

### GetScope

`func (o *JiraWorkflowPreviewStatus) GetScope() WorkflowPreviewScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *JiraWorkflowPreviewStatus) GetScopeOk() (*WorkflowPreviewScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *JiraWorkflowPreviewStatus) SetScope(v WorkflowPreviewScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *JiraWorkflowPreviewStatus) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatusCategory

`func (o *JiraWorkflowPreviewStatus) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *JiraWorkflowPreviewStatus) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *JiraWorkflowPreviewStatus) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *JiraWorkflowPreviewStatus) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStatusReference

`func (o *JiraWorkflowPreviewStatus) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *JiraWorkflowPreviewStatus) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *JiraWorkflowPreviewStatus) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.

### HasStatusReference

`func (o *JiraWorkflowPreviewStatus) HasStatusReference() bool`

HasStatusReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


