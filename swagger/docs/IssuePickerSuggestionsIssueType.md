# IssuePickerSuggestionsIssueType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the type of issues suggested for use in auto-completion. | [optional] [readonly] 
**Issues** | Pointer to [**[]SuggestedIssue**](SuggestedIssue.md) | A list of issues suggested for use in auto-completion. | [optional] [readonly] 
**Label** | Pointer to **string** | The label of the type of issues suggested for use in auto-completion. | [optional] [readonly] 
**Msg** | Pointer to **string** | If no issue suggestions are found, returns a message indicating no suggestions were found, | [optional] [readonly] 
**Sub** | Pointer to **string** | If issue suggestions are found, returns a message indicating the number of issues suggestions found and returned. | [optional] [readonly] 

## Methods

### NewIssuePickerSuggestionsIssueType

`func NewIssuePickerSuggestionsIssueType() *IssuePickerSuggestionsIssueType`

NewIssuePickerSuggestionsIssueType instantiates a new IssuePickerSuggestionsIssueType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuePickerSuggestionsIssueTypeWithDefaults

`func NewIssuePickerSuggestionsIssueTypeWithDefaults() *IssuePickerSuggestionsIssueType`

NewIssuePickerSuggestionsIssueTypeWithDefaults instantiates a new IssuePickerSuggestionsIssueType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssuePickerSuggestionsIssueType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssuePickerSuggestionsIssueType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssuePickerSuggestionsIssueType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssuePickerSuggestionsIssueType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssues

`func (o *IssuePickerSuggestionsIssueType) GetIssues() []SuggestedIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *IssuePickerSuggestionsIssueType) GetIssuesOk() (*[]SuggestedIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *IssuePickerSuggestionsIssueType) SetIssues(v []SuggestedIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *IssuePickerSuggestionsIssueType) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetLabel

`func (o *IssuePickerSuggestionsIssueType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IssuePickerSuggestionsIssueType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IssuePickerSuggestionsIssueType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IssuePickerSuggestionsIssueType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMsg

`func (o *IssuePickerSuggestionsIssueType) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *IssuePickerSuggestionsIssueType) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *IssuePickerSuggestionsIssueType) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *IssuePickerSuggestionsIssueType) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSub

`func (o *IssuePickerSuggestionsIssueType) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IssuePickerSuggestionsIssueType) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IssuePickerSuggestionsIssueType) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IssuePickerSuggestionsIssueType) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


