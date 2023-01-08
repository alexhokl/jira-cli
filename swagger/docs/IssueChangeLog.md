# IssueChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeHistories** | Pointer to [**[]Changelog**](Changelog.md) | List of changelogs that belongs to given issueId. | [optional] [readonly] 
**IssueId** | Pointer to **string** | The ID of the issue. | [optional] [readonly] 

## Methods

### NewIssueChangeLog

`func NewIssueChangeLog() *IssueChangeLog`

NewIssueChangeLog instantiates a new IssueChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueChangeLogWithDefaults

`func NewIssueChangeLogWithDefaults() *IssueChangeLog`

NewIssueChangeLogWithDefaults instantiates a new IssueChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeHistories

`func (o *IssueChangeLog) GetChangeHistories() []Changelog`

GetChangeHistories returns the ChangeHistories field if non-nil, zero value otherwise.

### GetChangeHistoriesOk

`func (o *IssueChangeLog) GetChangeHistoriesOk() (*[]Changelog, bool)`

GetChangeHistoriesOk returns a tuple with the ChangeHistories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeHistories

`func (o *IssueChangeLog) SetChangeHistories(v []Changelog)`

SetChangeHistories sets ChangeHistories field to given value.

### HasChangeHistories

`func (o *IssueChangeLog) HasChangeHistories() bool`

HasChangeHistories returns a boolean if a field has been set.

### GetIssueId

`func (o *IssueChangeLog) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *IssueChangeLog) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *IssueChangeLog) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *IssueChangeLog) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


