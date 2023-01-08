# ProjectScopeBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **[]string** | Defines the behavior of the option in the project.If notSelectable is set, the option cannot be set as the field&#39;s value. This is useful for archiving an option that has previously been selected but shouldn&#39;t be used anymore.If defaultValue is set, the option is selected by default. | [optional] 
**Id** | Pointer to **int64** | The ID of the project that the option&#39;s behavior applies to. | [optional] 

## Methods

### NewProjectScopeBean

`func NewProjectScopeBean() *ProjectScopeBean`

NewProjectScopeBean instantiates a new ProjectScopeBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectScopeBeanWithDefaults

`func NewProjectScopeBeanWithDefaults() *ProjectScopeBean`

NewProjectScopeBeanWithDefaults instantiates a new ProjectScopeBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ProjectScopeBean) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProjectScopeBean) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProjectScopeBean) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProjectScopeBean) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *ProjectScopeBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectScopeBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectScopeBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectScopeBean) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


