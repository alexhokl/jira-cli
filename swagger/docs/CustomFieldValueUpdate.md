# CustomFieldValueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIds** | **[]int64** | The list of issue IDs. | 
**Value** | **interface{}** | The value for the custom field. The value must be compatible with the [custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#data-types) as follows:   *  &#x60;string&#x60; the value must be a string.  *  &#x60;number&#x60; the value must be a number.  *  &#x60;datetime&#x60; the value must be a string that represents a date in the ISO format or the simplified extended ISO format. For example, &#x60;\&quot;2023-01-18T12:00:00-03:00\&quot;&#x60; or &#x60;\&quot;2023-01-18T12:00:00.000Z\&quot;&#x60;. However, the milliseconds part is ignored.  *  &#x60;user&#x60; the value must be an object that contains the &#x60;accountId&#x60; field.  *  &#x60;group&#x60; the value must be an object that contains the group &#x60;name&#x60; or &#x60;groupId&#x60; field. Because group names can change, we recommend using &#x60;groupId&#x60;.  A list of appropriate values must be provided if the field is of the &#x60;list&#x60; [collection type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#collection-types). | 

## Methods

### NewCustomFieldValueUpdate

`func NewCustomFieldValueUpdate(issueIds []int64, value interface{}, ) *CustomFieldValueUpdate`

NewCustomFieldValueUpdate instantiates a new CustomFieldValueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValueUpdateWithDefaults

`func NewCustomFieldValueUpdateWithDefaults() *CustomFieldValueUpdate`

NewCustomFieldValueUpdateWithDefaults instantiates a new CustomFieldValueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueIds

`func (o *CustomFieldValueUpdate) GetIssueIds() []int64`

GetIssueIds returns the IssueIds field if non-nil, zero value otherwise.

### GetIssueIdsOk

`func (o *CustomFieldValueUpdate) GetIssueIdsOk() (*[]int64, bool)`

GetIssueIdsOk returns a tuple with the IssueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIds

`func (o *CustomFieldValueUpdate) SetIssueIds(v []int64)`

SetIssueIds sets IssueIds field to given value.


### GetValue

`func (o *CustomFieldValueUpdate) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldValueUpdate) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldValueUpdate) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *CustomFieldValueUpdate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomFieldValueUpdate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


