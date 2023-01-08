# IssueFieldOptionCreateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**IssueFieldOptionConfiguration**](IssueFieldOptionConfiguration.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** | The properties of the option as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/) are defined in the descriptor for the issue field module. | [optional] 
**Value** | **string** | The option&#39;s name, which is displayed in Jira. | 

## Methods

### NewIssueFieldOptionCreateBean

`func NewIssueFieldOptionCreateBean(value string, ) *IssueFieldOptionCreateBean`

NewIssueFieldOptionCreateBean instantiates a new IssueFieldOptionCreateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldOptionCreateBeanWithDefaults

`func NewIssueFieldOptionCreateBeanWithDefaults() *IssueFieldOptionCreateBean`

NewIssueFieldOptionCreateBeanWithDefaults instantiates a new IssueFieldOptionCreateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *IssueFieldOptionCreateBean) GetConfig() IssueFieldOptionConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IssueFieldOptionCreateBean) GetConfigOk() (*IssueFieldOptionConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IssueFieldOptionCreateBean) SetConfig(v IssueFieldOptionConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IssueFieldOptionCreateBean) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetProperties

`func (o *IssueFieldOptionCreateBean) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueFieldOptionCreateBean) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueFieldOptionCreateBean) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueFieldOptionCreateBean) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetValue

`func (o *IssueFieldOptionCreateBean) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IssueFieldOptionCreateBean) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IssueFieldOptionCreateBean) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


