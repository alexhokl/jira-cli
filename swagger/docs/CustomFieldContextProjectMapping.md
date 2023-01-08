# CustomFieldContextProjectMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextId** | **string** | The ID of the context. | [readonly] 
**IsGlobalContext** | Pointer to **bool** | Whether context is global. | [optional] [readonly] 
**ProjectId** | Pointer to **string** | The ID of the project. | [optional] [readonly] 

## Methods

### NewCustomFieldContextProjectMapping

`func NewCustomFieldContextProjectMapping(contextId string, ) *CustomFieldContextProjectMapping`

NewCustomFieldContextProjectMapping instantiates a new CustomFieldContextProjectMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldContextProjectMappingWithDefaults

`func NewCustomFieldContextProjectMappingWithDefaults() *CustomFieldContextProjectMapping`

NewCustomFieldContextProjectMappingWithDefaults instantiates a new CustomFieldContextProjectMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextId

`func (o *CustomFieldContextProjectMapping) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *CustomFieldContextProjectMapping) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *CustomFieldContextProjectMapping) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetIsGlobalContext

`func (o *CustomFieldContextProjectMapping) GetIsGlobalContext() bool`

GetIsGlobalContext returns the IsGlobalContext field if non-nil, zero value otherwise.

### GetIsGlobalContextOk

`func (o *CustomFieldContextProjectMapping) GetIsGlobalContextOk() (*bool, bool)`

GetIsGlobalContextOk returns a tuple with the IsGlobalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalContext

`func (o *CustomFieldContextProjectMapping) SetIsGlobalContext(v bool)`

SetIsGlobalContext sets IsGlobalContext field to given value.

### HasIsGlobalContext

`func (o *CustomFieldContextProjectMapping) HasIsGlobalContext() bool`

HasIsGlobalContext returns a boolean if a field has been set.

### GetProjectId

`func (o *CustomFieldContextProjectMapping) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CustomFieldContextProjectMapping) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CustomFieldContextProjectMapping) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CustomFieldContextProjectMapping) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


