# BulkIssuePropertyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | EXPERIMENTAL. The Jira expression to calculate the value of the property. The value of the expression must be an object that can be converted to JSON, such as a number, boolean, string, list, or map. The context variables available to the expression are &#x60;issue&#x60; and &#x60;user&#x60;. Issues for which the expression returns a value whose JSON representation is longer than 32768 characters are ignored. | [optional] 
**Filter** | Pointer to [**IssueFilterForBulkPropertySet**](IssueFilterForBulkPropertySet.md) | The bulk operation filter. | [optional] 
**Value** | Pointer to **interface{}** | The value of the property. The value must be a [valid](https://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters. | [optional] 

## Methods

### NewBulkIssuePropertyUpdateRequest

`func NewBulkIssuePropertyUpdateRequest() *BulkIssuePropertyUpdateRequest`

NewBulkIssuePropertyUpdateRequest instantiates a new BulkIssuePropertyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuePropertyUpdateRequestWithDefaults

`func NewBulkIssuePropertyUpdateRequestWithDefaults() *BulkIssuePropertyUpdateRequest`

NewBulkIssuePropertyUpdateRequestWithDefaults instantiates a new BulkIssuePropertyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *BulkIssuePropertyUpdateRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *BulkIssuePropertyUpdateRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *BulkIssuePropertyUpdateRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *BulkIssuePropertyUpdateRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFilter

`func (o *BulkIssuePropertyUpdateRequest) GetFilter() IssueFilterForBulkPropertySet`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BulkIssuePropertyUpdateRequest) GetFilterOk() (*IssueFilterForBulkPropertySet, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BulkIssuePropertyUpdateRequest) SetFilter(v IssueFilterForBulkPropertySet)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BulkIssuePropertyUpdateRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetValue

`func (o *BulkIssuePropertyUpdateRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkIssuePropertyUpdateRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkIssuePropertyUpdateRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BulkIssuePropertyUpdateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BulkIssuePropertyUpdateRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BulkIssuePropertyUpdateRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


