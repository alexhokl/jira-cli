# CreateCustomFieldContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the context. | [optional] 
**Id** | Pointer to **string** | The ID of the context. | [optional] [readonly] 
**IssueTypeIds** | Pointer to **[]string** | The list of issue types IDs for the context. If the list is empty, the context refers to all issue types. | [optional] 
**Name** | **string** | The name of the context. | 
**ProjectIds** | Pointer to **[]string** | The list of project IDs associated with the context. If the list is empty, the context is global. | [optional] 

## Methods

### NewCreateCustomFieldContext

`func NewCreateCustomFieldContext(name string, ) *CreateCustomFieldContext`

NewCreateCustomFieldContext instantiates a new CreateCustomFieldContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomFieldContextWithDefaults

`func NewCreateCustomFieldContextWithDefaults() *CreateCustomFieldContext`

NewCreateCustomFieldContextWithDefaults instantiates a new CreateCustomFieldContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateCustomFieldContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomFieldContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomFieldContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomFieldContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CreateCustomFieldContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCustomFieldContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCustomFieldContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCustomFieldContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *CreateCustomFieldContext) GetIssueTypeIds() []string`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *CreateCustomFieldContext) GetIssueTypeIdsOk() (*[]string, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *CreateCustomFieldContext) SetIssueTypeIds(v []string)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *CreateCustomFieldContext) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetName

`func (o *CreateCustomFieldContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomFieldContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomFieldContext) SetName(v string)`

SetName sets Name field to given value.


### GetProjectIds

`func (o *CreateCustomFieldContext) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateCustomFieldContext) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateCustomFieldContext) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateCustomFieldContext) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


