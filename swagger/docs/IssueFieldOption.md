# IssueFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**IssueFieldOptionConfiguration**](IssueFieldOptionConfiguration.md) |  | [optional] 
**Id** | **int64** | The unique identifier for the option. This is only unique within the select field&#39;s set of options. | 
**Properties** | Pointer to **map[string]interface{}** | The properties of the object, as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see [Issue Field Option Property Index](https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/)) are defined in the descriptor for the issue field module. | [optional] 
**Value** | **string** | The option&#39;s name, which is displayed in Jira. | 

## Methods

### NewIssueFieldOption

`func NewIssueFieldOption(id int64, value string, ) *IssueFieldOption`

NewIssueFieldOption instantiates a new IssueFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldOptionWithDefaults

`func NewIssueFieldOptionWithDefaults() *IssueFieldOption`

NewIssueFieldOptionWithDefaults instantiates a new IssueFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *IssueFieldOption) GetConfig() IssueFieldOptionConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IssueFieldOption) GetConfigOk() (*IssueFieldOptionConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IssueFieldOption) SetConfig(v IssueFieldOptionConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IssueFieldOption) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *IssueFieldOption) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueFieldOption) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueFieldOption) SetId(v int64)`

SetId sets Id field to given value.


### GetProperties

`func (o *IssueFieldOption) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IssueFieldOption) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IssueFieldOption) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IssueFieldOption) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetValue

`func (o *IssueFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IssueFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IssueFieldOption) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


