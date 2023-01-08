# IssueFieldOptionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **[]string** | DEPRECATED | [optional] 
**Scope** | Pointer to [**IssueFieldOptionScopeBean**](IssueFieldOptionScopeBean.md) | Defines the projects that the option is available in. If the scope is not defined, then the option is available in all projects. | [optional] 

## Methods

### NewIssueFieldOptionConfiguration

`func NewIssueFieldOptionConfiguration() *IssueFieldOptionConfiguration`

NewIssueFieldOptionConfiguration instantiates a new IssueFieldOptionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldOptionConfigurationWithDefaults

`func NewIssueFieldOptionConfigurationWithDefaults() *IssueFieldOptionConfiguration`

NewIssueFieldOptionConfigurationWithDefaults instantiates a new IssueFieldOptionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IssueFieldOptionConfiguration) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IssueFieldOptionConfiguration) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IssueFieldOptionConfiguration) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IssueFieldOptionConfiguration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetScope

`func (o *IssueFieldOptionConfiguration) GetScope() IssueFieldOptionScopeBean`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IssueFieldOptionConfiguration) GetScopeOk() (*IssueFieldOptionScopeBean, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IssueFieldOptionConfiguration) SetScope(v IssueFieldOptionScopeBean)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IssueFieldOptionConfiguration) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


