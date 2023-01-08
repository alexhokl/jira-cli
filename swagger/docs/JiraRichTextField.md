# JiraRichTextField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**RichText** | [**JiraRichTextInput**](JiraRichTextInput.md) |  | 

## Methods

### NewJiraRichTextField

`func NewJiraRichTextField(fieldId string, richText JiraRichTextInput, ) *JiraRichTextField`

NewJiraRichTextField instantiates a new JiraRichTextField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraRichTextFieldWithDefaults

`func NewJiraRichTextFieldWithDefaults() *JiraRichTextField`

NewJiraRichTextFieldWithDefaults instantiates a new JiraRichTextField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraRichTextField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraRichTextField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraRichTextField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetRichText

`func (o *JiraRichTextField) GetRichText() JiraRichTextInput`

GetRichText returns the RichText field if non-nil, zero value otherwise.

### GetRichTextOk

`func (o *JiraRichTextField) GetRichTextOk() (*JiraRichTextInput, bool)`

GetRichTextOk returns a tuple with the RichText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichText

`func (o *JiraRichTextField) SetRichText(v JiraRichTextInput)`

SetRichText sets RichText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


