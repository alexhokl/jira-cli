# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional search result details in the response. | [optional] [readonly] 
**Issues** | Pointer to [**[]IssueBean**](IssueBean.md) | The list of issues found by the search. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**Names** | Pointer to **map[string]string** | The ID and name of each field in the search results. | [optional] [readonly] 
**Schema** | Pointer to [**map[string]IssueBeanSchemaValue**](IssueBeanSchemaValue.md) | The schema describing the field types in the search results. | [optional] [readonly] 
**StartAt** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of results on the page. | [optional] [readonly] 
**WarningMessages** | Pointer to **[]string** | Any warnings related to the JQL query. | [optional] [readonly] 

## Methods

### NewSearchResults

`func NewSearchResults() *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *SearchResults) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *SearchResults) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *SearchResults) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *SearchResults) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetIssues

`func (o *SearchResults) GetIssues() []IssueBean`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *SearchResults) GetIssuesOk() (*[]IssueBean, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *SearchResults) SetIssues(v []IssueBean)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *SearchResults) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetMaxResults

`func (o *SearchResults) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchResults) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchResults) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchResults) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNames

`func (o *SearchResults) GetNames() map[string]string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *SearchResults) GetNamesOk() (*map[string]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *SearchResults) SetNames(v map[string]string)`

SetNames sets Names field to given value.

### HasNames

`func (o *SearchResults) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetSchema

`func (o *SearchResults) GetSchema() map[string]IssueBeanSchemaValue`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SearchResults) GetSchemaOk() (*map[string]IssueBeanSchemaValue, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SearchResults) SetSchema(v map[string]IssueBeanSchemaValue)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SearchResults) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetStartAt

`func (o *SearchResults) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *SearchResults) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *SearchResults) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *SearchResults) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *SearchResults) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResults) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResults) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchResults) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarningMessages

`func (o *SearchResults) GetWarningMessages() []string`

GetWarningMessages returns the WarningMessages field if non-nil, zero value otherwise.

### GetWarningMessagesOk

`func (o *SearchResults) GetWarningMessagesOk() (*[]string, bool)`

GetWarningMessagesOk returns a tuple with the WarningMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningMessages

`func (o *SearchResults) SetWarningMessages(v []string)`

SetWarningMessages sets WarningMessages field to given value.

### HasWarningMessages

`func (o *SearchResults) HasWarningMessages() bool`

HasWarningMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


