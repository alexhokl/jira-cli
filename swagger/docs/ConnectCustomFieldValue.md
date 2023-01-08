# ConnectCustomFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of custom field. | 
**FieldID** | **int32** | The custom field ID. | 
**IssueID** | **int32** | The issue ID. | 
**Number** | Pointer to **float32** | The value of number type custom field when &#x60;_type&#x60; is &#x60;NumberIssueField&#x60;. | [optional] 
**OptionID** | Pointer to **string** | The value of single select and multiselect custom field type when &#x60;_type&#x60; is &#x60;SingleSelectIssueField&#x60; or &#x60;MultiSelectIssueField&#x60;. | [optional] 
**RichText** | Pointer to **string** | The value of richText type custom field when &#x60;_type&#x60; is &#x60;RichTextIssueField&#x60;. | [optional] 
**String** | Pointer to **string** | The value of string type custom field when &#x60;_type&#x60; is &#x60;StringIssueField&#x60;. | [optional] 
**Text** | Pointer to **string** | The value of of text custom field type when &#x60;_type&#x60; is &#x60;TextIssueField&#x60;. | [optional] 

## Methods

### NewConnectCustomFieldValue

`func NewConnectCustomFieldValue(type_ string, fieldID int32, issueID int32, ) *ConnectCustomFieldValue`

NewConnectCustomFieldValue instantiates a new ConnectCustomFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectCustomFieldValueWithDefaults

`func NewConnectCustomFieldValueWithDefaults() *ConnectCustomFieldValue`

NewConnectCustomFieldValueWithDefaults instantiates a new ConnectCustomFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectCustomFieldValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectCustomFieldValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectCustomFieldValue) SetType(v string)`

SetType sets Type field to given value.


### GetFieldID

`func (o *ConnectCustomFieldValue) GetFieldID() int32`

GetFieldID returns the FieldID field if non-nil, zero value otherwise.

### GetFieldIDOk

`func (o *ConnectCustomFieldValue) GetFieldIDOk() (*int32, bool)`

GetFieldIDOk returns a tuple with the FieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldID

`func (o *ConnectCustomFieldValue) SetFieldID(v int32)`

SetFieldID sets FieldID field to given value.


### GetIssueID

`func (o *ConnectCustomFieldValue) GetIssueID() int32`

GetIssueID returns the IssueID field if non-nil, zero value otherwise.

### GetIssueIDOk

`func (o *ConnectCustomFieldValue) GetIssueIDOk() (*int32, bool)`

GetIssueIDOk returns a tuple with the IssueID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueID

`func (o *ConnectCustomFieldValue) SetIssueID(v int32)`

SetIssueID sets IssueID field to given value.


### GetNumber

`func (o *ConnectCustomFieldValue) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ConnectCustomFieldValue) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ConnectCustomFieldValue) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ConnectCustomFieldValue) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOptionID

`func (o *ConnectCustomFieldValue) GetOptionID() string`

GetOptionID returns the OptionID field if non-nil, zero value otherwise.

### GetOptionIDOk

`func (o *ConnectCustomFieldValue) GetOptionIDOk() (*string, bool)`

GetOptionIDOk returns a tuple with the OptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionID

`func (o *ConnectCustomFieldValue) SetOptionID(v string)`

SetOptionID sets OptionID field to given value.

### HasOptionID

`func (o *ConnectCustomFieldValue) HasOptionID() bool`

HasOptionID returns a boolean if a field has been set.

### GetRichText

`func (o *ConnectCustomFieldValue) GetRichText() string`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *ConnectCustomFieldValue) GetRichTextOk() (*string, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *ConnectCustomFieldValue) SetRichText(v string)`

SetRichText sets RichText field to given value.

### HasRichText

`func (o *ConnectCustomFieldValue) HasRichText() bool`

HasRichText returns a boolean if a field has been set.

### GetString

`func (o *ConnectCustomFieldValue) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *ConnectCustomFieldValue) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *ConnectCustomFieldValue) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *ConnectCustomFieldValue) HasString() bool`

HasString returns a boolean if a field has been set.

### GetText

`func (o *ConnectCustomFieldValue) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConnectCustomFieldValue) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConnectCustomFieldValue) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConnectCustomFieldValue) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


