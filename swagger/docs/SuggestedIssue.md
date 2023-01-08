# SuggestedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the issue. | [optional] [readonly] 
**Img** | Pointer to **string** | The URL of the issue type&#39;s avatar. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the issue. | [optional] [readonly] 
**KeyHtml** | Pointer to **string** | The key of the issue in HTML format. | [optional] [readonly] 
**Summary** | Pointer to **string** | The phrase containing the query string in HTML format, with the string highlighted with HTML bold tags. | [optional] [readonly] 
**SummaryText** | Pointer to **string** | The phrase containing the query string, as plain text. | [optional] [readonly] 

## Methods

### NewSuggestedIssue

`func NewSuggestedIssue() *SuggestedIssue`

NewSuggestedIssue instantiates a new SuggestedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedIssueWithDefaults

`func NewSuggestedIssueWithDefaults() *SuggestedIssue`

NewSuggestedIssueWithDefaults instantiates a new SuggestedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuggestedIssue) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuggestedIssue) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuggestedIssue) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SuggestedIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImg

`func (o *SuggestedIssue) GetImg() string`

GetImg returns the Img field if non-nil, zero value otherwise.

### GetImgOk

`func (o *SuggestedIssue) GetImgOk() (*string, bool)`

GetImgOk returns a tuple with the Img field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImg

`func (o *SuggestedIssue) SetImg(v string)`

SetImg sets Img field to given value.

### HasImg

`func (o *SuggestedIssue) HasImg() bool`

HasImg returns a boolean if a field has been set.

### GetKey

`func (o *SuggestedIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SuggestedIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SuggestedIssue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SuggestedIssue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyHtml

`func (o *SuggestedIssue) GetKeyHtml() string`

GetKeyHtml returns the KeyHtml field if non-nil, zero value otherwise.

### GetKeyHtmlOk

`func (o *SuggestedIssue) GetKeyHtmlOk() (*string, bool)`

GetKeyHtmlOk returns a tuple with the KeyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyHtml

`func (o *SuggestedIssue) SetKeyHtml(v string)`

SetKeyHtml sets KeyHtml field to given value.

### HasKeyHtml

`func (o *SuggestedIssue) HasKeyHtml() bool`

HasKeyHtml returns a boolean if a field has been set.

### GetSummary

`func (o *SuggestedIssue) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SuggestedIssue) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SuggestedIssue) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SuggestedIssue) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetSummaryText

`func (o *SuggestedIssue) GetSummaryText() string`

GetSummaryText returns the SummaryText field if non-nil, zero value otherwise.

### GetSummaryTextOk

`func (o *SuggestedIssue) GetSummaryTextOk() (*string, bool)`

GetSummaryTextOk returns a tuple with the SummaryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryText

`func (o *SuggestedIssue) SetSummaryText(v string)`

SetSummaryText sets SummaryText field to given value.

### HasSummaryText

`func (o *SuggestedIssue) HasSummaryText() bool`

HasSummaryText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


