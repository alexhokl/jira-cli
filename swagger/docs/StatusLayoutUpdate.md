# StatusLayoutUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalConfiguration** | Pointer to [**NullableApprovalConfiguration**](ApprovalConfiguration.md) |  | [optional] 
**Layout** | Pointer to [**NullableWorkflowLayout**](WorkflowLayout.md) |  | [optional] 
**Properties** | **map[string]string** | The properties for this status layout. | 
**StatusReference** | **string** | A unique ID which the status will use to refer to this layout configuration. | 

## Methods

### NewStatusLayoutUpdate

`func NewStatusLayoutUpdate(properties map[string]string, statusReference string, ) *StatusLayoutUpdate`

NewStatusLayoutUpdate instantiates a new StatusLayoutUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusLayoutUpdateWithDefaults

`func NewStatusLayoutUpdateWithDefaults() *StatusLayoutUpdate`

NewStatusLayoutUpdateWithDefaults instantiates a new StatusLayoutUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalConfiguration

`func (o *StatusLayoutUpdate) GetApprovalConfiguration() ApprovalConfiguration`

GetApprovalConfiguration returns the ApprovalConfiguration field if non-nil, zero value otherwise.

### GetApprovalConfigurationOk

`func (o *StatusLayoutUpdate) GetApprovalConfigurationOk() (*ApprovalConfiguration, bool)`

GetApprovalConfigurationOk returns a tuple with the ApprovalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfiguration

`func (o *StatusLayoutUpdate) SetApprovalConfiguration(v ApprovalConfiguration)`

SetApprovalConfiguration sets ApprovalConfiguration field to given value.

### HasApprovalConfiguration

`func (o *StatusLayoutUpdate) HasApprovalConfiguration() bool`

HasApprovalConfiguration returns a boolean if a field has been set.

### SetApprovalConfigurationNil

`func (o *StatusLayoutUpdate) SetApprovalConfigurationNil(b bool)`

 SetApprovalConfigurationNil sets the value for ApprovalConfiguration to be an explicit nil

### UnsetApprovalConfiguration
`func (o *StatusLayoutUpdate) UnsetApprovalConfiguration()`

UnsetApprovalConfiguration ensures that no value is present for ApprovalConfiguration, not even an explicit nil
### GetLayout

`func (o *StatusLayoutUpdate) GetLayout() WorkflowLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *StatusLayoutUpdate) GetLayoutOk() (*WorkflowLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *StatusLayoutUpdate) SetLayout(v WorkflowLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *StatusLayoutUpdate) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *StatusLayoutUpdate) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *StatusLayoutUpdate) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil
### GetProperties

`func (o *StatusLayoutUpdate) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StatusLayoutUpdate) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StatusLayoutUpdate) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.


### GetStatusReference

`func (o *StatusLayoutUpdate) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *StatusLayoutUpdate) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *StatusLayoutUpdate) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


