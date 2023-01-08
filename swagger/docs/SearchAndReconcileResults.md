# SearchAndReconcileResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Indicates whether this is the last page of the paginated response. | [optional] [readonly] 
**Issues** | Pointer to [**[]IssueBean**](IssueBean.md) | The list of issues found by the search or reconsiliation. | [optional] [readonly] 
**Names** | Pointer to **map[string]string** | The ID and name of each field in the search results. | [optional] [readonly] 
**NextPageToken** | Pointer to **string** | Continuation token to fetch the next page. If this result represents the last or the only page this token will be null. This token will expire in 7 days. | [optional] [readonly] 
**Schema** | Pointer to [**map[string]JsonTypeBean**](JsonTypeBean.md) | The schema describing the field types in the search results. | [optional] [readonly] 

## Methods

### NewSearchAndReconcileResults

`func NewSearchAndReconcileResults() *SearchAndReconcileResults`

NewSearchAndReconcileResults instantiates a new SearchAndReconcileResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAndReconcileResultsWithDefaults

`func NewSearchAndReconcileResultsWithDefaults() *SearchAndReconcileResults`

NewSearchAndReconcileResultsWithDefaults instantiates a new SearchAndReconcileResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *SearchAndReconcileResults) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *SearchAndReconcileResults) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *SearchAndReconcileResults) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *SearchAndReconcileResults) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetIssues

`func (o *SearchAndReconcileResults) GetIssues() []IssueBean`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *SearchAndReconcileResults) GetIssuesOk() (*[]IssueBean, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *SearchAndReconcileResults) SetIssues(v []IssueBean)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *SearchAndReconcileResults) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetNames

`func (o *SearchAndReconcileResults) GetNames() map[string]string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *SearchAndReconcileResults) GetNamesOk() (*map[string]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *SearchAndReconcileResults) SetNames(v map[string]string)`

SetNames sets Names field to given value.

### HasNames

`func (o *SearchAndReconcileResults) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNextPageToken

`func (o *SearchAndReconcileResults) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SearchAndReconcileResults) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SearchAndReconcileResults) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SearchAndReconcileResults) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSchema

`func (o *SearchAndReconcileResults) GetSchema() map[string]JsonTypeBean`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SearchAndReconcileResults) GetSchemaOk() (*map[string]JsonTypeBean, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SearchAndReconcileResults) SetSchema(v map[string]JsonTypeBean)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SearchAndReconcileResults) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


