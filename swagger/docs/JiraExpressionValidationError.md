# JiraExpressionValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** | The text column in which the error occurred. | [optional] 
**Expression** | Pointer to **string** | The part of the expression in which the error occurred. | [optional] 
**Line** | Pointer to **int32** | The text line in which the error occurred. | [optional] 
**Message** | **string** | Details about the error. | 
**Type** | **string** | The error type. | 

## Methods

### NewJiraExpressionValidationError

`func NewJiraExpressionValidationError(message string, type_ string, ) *JiraExpressionValidationError`

NewJiraExpressionValidationError instantiates a new JiraExpressionValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionValidationErrorWithDefaults

`func NewJiraExpressionValidationErrorWithDefaults() *JiraExpressionValidationError`

NewJiraExpressionValidationErrorWithDefaults instantiates a new JiraExpressionValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *JiraExpressionValidationError) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *JiraExpressionValidationError) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *JiraExpressionValidationError) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *JiraExpressionValidationError) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetExpression

`func (o *JiraExpressionValidationError) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *JiraExpressionValidationError) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *JiraExpressionValidationError) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *JiraExpressionValidationError) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetLine

`func (o *JiraExpressionValidationError) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *JiraExpressionValidationError) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *JiraExpressionValidationError) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *JiraExpressionValidationError) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetMessage

`func (o *JiraExpressionValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JiraExpressionValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JiraExpressionValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *JiraExpressionValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JiraExpressionValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JiraExpressionValidationError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


