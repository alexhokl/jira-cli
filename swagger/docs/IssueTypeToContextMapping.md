# IssueTypeToContextMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextId** | **string** | The ID of the context. | 
**IsAnyIssueType** | Pointer to **bool** | Whether the context is mapped to any issue type. | [optional] 
**IssueTypeId** | Pointer to **string** | The ID of the issue type. | [optional] 

## Methods

### NewIssueTypeToContextMapping

`func NewIssueTypeToContextMapping(contextId string, ) *IssueTypeToContextMapping`

NewIssueTypeToContextMapping instantiates a new IssueTypeToContextMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeToContextMappingWithDefaults

`func NewIssueTypeToContextMappingWithDefaults() *IssueTypeToContextMapping`

NewIssueTypeToContextMappingWithDefaults instantiates a new IssueTypeToContextMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextId

`func (o *IssueTypeToContextMapping) GetContextId() string`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *IssueTypeToContextMapping) GetContextIdOk() (*string, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *IssueTypeToContextMapping) SetContextId(v string)`

SetContextId sets ContextId field to given value.


### GetIsAnyIssueType

`func (o *IssueTypeToContextMapping) GetIsAnyIssueType() bool`

GetIsAnyIssueType returns the IsAnyIssueType field if non-nil, zero value otherwise.

### GetIsAnyIssueTypeOk

`func (o *IssueTypeToContextMapping) GetIsAnyIssueTypeOk() (*bool, bool)`

GetIsAnyIssueTypeOk returns a tuple with the IsAnyIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyIssueType

`func (o *IssueTypeToContextMapping) SetIsAnyIssueType(v bool)`

SetIsAnyIssueType sets IsAnyIssueType field to given value.

### HasIsAnyIssueType

`func (o *IssueTypeToContextMapping) HasIsAnyIssueType() bool`

HasIsAnyIssueType returns a boolean if a field has been set.

### GetIssueTypeId

`func (o *IssueTypeToContextMapping) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *IssueTypeToContextMapping) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *IssueTypeToContextMapping) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *IssueTypeToContextMapping) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


