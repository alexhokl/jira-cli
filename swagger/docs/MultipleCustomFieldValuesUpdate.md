# MultipleCustomFieldValuesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomField** | **string** | The ID or key of the custom field. For example, &#x60;customfield_10010&#x60;. | 
**IssueIds** | **[]int64** | The list of issue IDs. | 
**Value** | **interface{}** | The value for the custom field. The value must be compatible with the [custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#data-types) as follows:   *  &#x60;string&#x60; the value must be a string.  *  &#x60;number&#x60; the value must be a number.  *  &#x60;datetime&#x60; the value must be a string that represents a date in the ISO format or the simplified extended ISO format. For example, &#x60;\&quot;2023-01-18T12:00:00-03:00\&quot;&#x60; or &#x60;\&quot;2023-01-18T12:00:00.000Z\&quot;&#x60;. However, the milliseconds part is ignored.  *  &#x60;user&#x60; the value must be an object that contains the &#x60;accountId&#x60; field.  *  &#x60;group&#x60; the value must be an object that contains the group &#x60;name&#x60; or &#x60;groupId&#x60; field. Because group names can change, we recommend using &#x60;groupId&#x60;.  A list of appropriate values must be provided if the field is of the &#x60;list&#x60; [collection type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#collection-types). | 

## Methods

### NewMultipleCustomFieldValuesUpdate

`func NewMultipleCustomFieldValuesUpdate(customField string, issueIds []int64, value interface{}, ) *MultipleCustomFieldValuesUpdate`

NewMultipleCustomFieldValuesUpdate instantiates a new MultipleCustomFieldValuesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleCustomFieldValuesUpdateWithDefaults

`func NewMultipleCustomFieldValuesUpdateWithDefaults() *MultipleCustomFieldValuesUpdate`

NewMultipleCustomFieldValuesUpdateWithDefaults instantiates a new MultipleCustomFieldValuesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomField

`func (o *MultipleCustomFieldValuesUpdate) GetCustomField() string`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *MultipleCustomFieldValuesUpdate) GetCustomFieldOk() (*string, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *MultipleCustomFieldValuesUpdate) SetCustomField(v string)`

SetCustomField sets CustomField field to given value.


### GetIssueIds

`func (o *MultipleCustomFieldValuesUpdate) GetIssueIds() []int64`

GetIssueIds returns the IssueIds field if non-nil, zero value otherwise.

### GetIssueIdsOk

`func (o *MultipleCustomFieldValuesUpdate) GetIssueIdsOk() (*[]int64, bool)`

GetIssueIdsOk returns a tuple with the IssueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIds

`func (o *MultipleCustomFieldValuesUpdate) SetIssueIds(v []int64)`

SetIssueIds sets IssueIds field to given value.


### GetValue

`func (o *MultipleCustomFieldValuesUpdate) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MultipleCustomFieldValuesUpdate) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MultipleCustomFieldValuesUpdate) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *MultipleCustomFieldValuesUpdate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MultipleCustomFieldValuesUpdate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


