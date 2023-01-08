# CustomFieldContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the context. | 
**Id** | **string** | The ID of the context. | 
**IsAnyIssueType** | **bool** | Whether the context apply to all issue types. | 
**IsGlobalContext** | **bool** | Whether the context is global. | 
**Name** | **string** | The name of the context. | 

## Methods

### NewCustomFieldContext

`func NewCustomFieldContext(description string, id string, isAnyIssueType bool, isGlobalContext bool, name string, ) *CustomFieldContext`

NewCustomFieldContext instantiates a new CustomFieldContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextWithDefaults

`func NewCustomFieldContextWithDefaults() *CustomFieldContext`

NewCustomFieldContextWithDefaults instantiates a new CustomFieldContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomFieldContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomFieldContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomFieldContext) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *CustomFieldContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldContext) SetId(v string)`

SetId sets Id field to given value.


### GetIsAnyIssueType

`func (o *CustomFieldContext) GetIsAnyIssueType() bool`

GetIsAnyIssueType returns the IsAnyIssueType field if non-nil, zero value otherwise.

### GetIsAnyIssueTypeOk

`func (o *CustomFieldContext) GetIsAnyIssueTypeOk() (*bool, bool)`

GetIsAnyIssueTypeOk returns a tuple with the IsAnyIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyIssueType

`func (o *CustomFieldContext) SetIsAnyIssueType(v bool)`

SetIsAnyIssueType sets IsAnyIssueType field to given value.


### GetIsGlobalContext

`func (o *CustomFieldContext) GetIsGlobalContext() bool`

GetIsGlobalContext returns the IsGlobalContext field if non-nil, zero value otherwise.

### GetIsGlobalContextOk

`func (o *CustomFieldContext) GetIsGlobalContextOk() (*bool, bool)`

GetIsGlobalContextOk returns a tuple with the IsGlobalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalContext

`func (o *CustomFieldContext) SetIsGlobalContext(v bool)`

SetIsGlobalContext sets IsGlobalContext field to given value.


### GetName

`func (o *CustomFieldContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldContext) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


