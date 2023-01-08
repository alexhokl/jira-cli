# DeprecatedWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** | The description of the workflow. | [optional] [readonly] 
**LastModifiedDate** | Pointer to **string** | The datetime the workflow was last modified. | [optional] [readonly] 
**LastModifiedUser** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [readonly] 
**LastModifiedUserAccountId** | Pointer to **string** | The account ID of the user that last modified the workflow. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the workflow. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope where this workflow applies | [optional] [readonly] 
**Steps** | Pointer to **int32** | The number of steps included in the workflow. | [optional] [readonly] 

## Methods

### NewDeprecatedWorkflow

`func NewDeprecatedWorkflow() *DeprecatedWorkflow`

NewDeprecatedWorkflow instantiates a new DeprecatedWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedWorkflowWithDefaults

`func NewDeprecatedWorkflowWithDefaults() *DeprecatedWorkflow`

NewDeprecatedWorkflowWithDefaults instantiates a new DeprecatedWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *DeprecatedWorkflow) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeprecatedWorkflow) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeprecatedWorkflow) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DeprecatedWorkflow) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *DeprecatedWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeprecatedWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeprecatedWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeprecatedWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DeprecatedWorkflow) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DeprecatedWorkflow) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DeprecatedWorkflow) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DeprecatedWorkflow) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedUser

`func (o *DeprecatedWorkflow) GetLastModifiedUser() string`

GetLastModifiedUser returns the LastModifiedUser field if non-nil, zero value otherwise.

### GetLastModifiedUserOk

`func (o *DeprecatedWorkflow) GetLastModifiedUserOk() (*string, bool)`

GetLastModifiedUserOk returns a tuple with the LastModifiedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedUser

`func (o *DeprecatedWorkflow) SetLastModifiedUser(v string)`

SetLastModifiedUser sets LastModifiedUser field to given value.

### HasLastModifiedUser

`func (o *DeprecatedWorkflow) HasLastModifiedUser() bool`

HasLastModifiedUser returns a boolean if a field has been set.

### GetLastModifiedUserAccountId

`func (o *DeprecatedWorkflow) GetLastModifiedUserAccountId() string`

GetLastModifiedUserAccountId returns the LastModifiedUserAccountId field if non-nil, zero value otherwise.

### GetLastModifiedUserAccountIdOk

`func (o *DeprecatedWorkflow) GetLastModifiedUserAccountIdOk() (*string, bool)`

GetLastModifiedUserAccountIdOk returns a tuple with the LastModifiedUserAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedUserAccountId

`func (o *DeprecatedWorkflow) SetLastModifiedUserAccountId(v string)`

SetLastModifiedUserAccountId sets LastModifiedUserAccountId field to given value.

### HasLastModifiedUserAccountId

`func (o *DeprecatedWorkflow) HasLastModifiedUserAccountId() bool`

HasLastModifiedUserAccountId returns a boolean if a field has been set.

### GetName

`func (o *DeprecatedWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeprecatedWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeprecatedWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeprecatedWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *DeprecatedWorkflow) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DeprecatedWorkflow) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DeprecatedWorkflow) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DeprecatedWorkflow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSteps

`func (o *DeprecatedWorkflow) GetSteps() int32`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *DeprecatedWorkflow) GetStepsOk() (*int32, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *DeprecatedWorkflow) SetSteps(v int32)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *DeprecatedWorkflow) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


