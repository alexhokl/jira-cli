# JiraWorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Id** | Pointer to **string** | The ID of the status. | [optional] 
**Name** | Pointer to **string** | The name of the status. | [optional] 
**Scope** | Pointer to [**WorkflowScope**](WorkflowScope.md) |  | [optional] 
**StatusCategory** | Pointer to **string** | The category of the status. | [optional] 
**StatusReference** | Pointer to **string** | The reference of the status. | [optional] 

## Methods

### NewJiraWorkflowStatus

`func NewJiraWorkflowStatus() *JiraWorkflowStatus`

NewJiraWorkflowStatus instantiates a new JiraWorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraWorkflowStatusWithDefaults

`func NewJiraWorkflowStatusWithDefaults() *JiraWorkflowStatus`

NewJiraWorkflowStatusWithDefaults instantiates a new JiraWorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JiraWorkflowStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JiraWorkflowStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JiraWorkflowStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JiraWorkflowStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *JiraWorkflowStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JiraWorkflowStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JiraWorkflowStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JiraWorkflowStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JiraWorkflowStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JiraWorkflowStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JiraWorkflowStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JiraWorkflowStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *JiraWorkflowStatus) GetScope() WorkflowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *JiraWorkflowStatus) GetScopeOk() (*WorkflowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *JiraWorkflowStatus) SetScope(v WorkflowScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *JiraWorkflowStatus) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatusCategory

`func (o *JiraWorkflowStatus) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *JiraWorkflowStatus) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *JiraWorkflowStatus) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *JiraWorkflowStatus) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetStatusReference

`func (o *JiraWorkflowStatus) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *JiraWorkflowStatus) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *JiraWorkflowStatus) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.

### HasStatusReference

`func (o *JiraWorkflowStatus) HasStatusReference() bool`

HasStatusReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


