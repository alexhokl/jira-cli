# ComponentIssuesCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCount** | Pointer to **int64** | The count of issues assigned to a component. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL for this count of issues for a component. | [optional] [readonly] 

## Methods

### NewComponentIssuesCount

`func NewComponentIssuesCount() *ComponentIssuesCount`

NewComponentIssuesCount instantiates a new ComponentIssuesCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentIssuesCountWithDefaults

`func NewComponentIssuesCountWithDefaults() *ComponentIssuesCount`

NewComponentIssuesCountWithDefaults instantiates a new ComponentIssuesCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCount

`func (o *ComponentIssuesCount) GetIssueCount() int64`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *ComponentIssuesCount) GetIssueCountOk() (*int64, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *ComponentIssuesCount) SetIssueCount(v int64)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *ComponentIssuesCount) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetSelf

`func (o *ComponentIssuesCount) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ComponentIssuesCount) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ComponentIssuesCount) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ComponentIssuesCount) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


