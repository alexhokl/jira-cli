# JiraStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] 
**Id** | Pointer to **string** | The ID of the status. | [optional] 
**Name** | Pointer to **string** | The name of the status. | [optional] 
**Scope** | Pointer to [**StatusScope**](StatusScope.md) |  | [optional] 
**StatusCategory** | Pointer to **string** | The category of the status. | [optional] 

## Methods

### NewJiraStatus

`func NewJiraStatus() *JiraStatus`

NewJiraStatus instantiates a new JiraStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraStatusWithDefaults

`func NewJiraStatusWithDefaults() *JiraStatus`

NewJiraStatusWithDefaults instantiates a new JiraStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *JiraStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JiraStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JiraStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JiraStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *JiraStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JiraStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JiraStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JiraStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *JiraStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JiraStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JiraStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JiraStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *JiraStatus) GetScope() StatusScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *JiraStatus) GetScopeOk() (*StatusScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *JiraStatus) SetScope(v StatusScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *JiraStatus) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatusCategory

`func (o *JiraStatus) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *JiraStatus) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *JiraStatus) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *JiraStatus) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


