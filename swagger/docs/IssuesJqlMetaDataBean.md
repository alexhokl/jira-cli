# IssuesJqlMetaDataBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The number of issues that were loaded in this evaluation. | 
**MaxResults** | **int32** | The maximum number of issues that could be loaded in this evaluation. | 
**StartAt** | **int64** | The index of the first issue. | 
**TotalCount** | **int64** | The total number of issues the JQL returned. | 
**ValidationWarnings** | Pointer to **[]string** | Any warnings related to the JQL query. Present only if the validation mode was set to &#x60;warn&#x60;. | [optional] 

## Methods

### NewIssuesJqlMetaDataBean

`func NewIssuesJqlMetaDataBean(count int32, maxResults int32, startAt int64, totalCount int64, ) *IssuesJqlMetaDataBean`

NewIssuesJqlMetaDataBean instantiates a new IssuesJqlMetaDataBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesJqlMetaDataBeanWithDefaults

`func NewIssuesJqlMetaDataBeanWithDefaults() *IssuesJqlMetaDataBean`

NewIssuesJqlMetaDataBeanWithDefaults instantiates a new IssuesJqlMetaDataBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IssuesJqlMetaDataBean) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IssuesJqlMetaDataBean) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IssuesJqlMetaDataBean) SetCount(v int32)`

SetCount sets Count field to given value.


### GetMaxResults

`func (o *IssuesJqlMetaDataBean) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *IssuesJqlMetaDataBean) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *IssuesJqlMetaDataBean) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.


### GetStartAt

`func (o *IssuesJqlMetaDataBean) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *IssuesJqlMetaDataBean) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *IssuesJqlMetaDataBean) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.


### GetTotalCount

`func (o *IssuesJqlMetaDataBean) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IssuesJqlMetaDataBean) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IssuesJqlMetaDataBean) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetValidationWarnings

`func (o *IssuesJqlMetaDataBean) GetValidationWarnings() []string`

GetValidationWarnings returns the ValidationWarnings field if non-nil, zero value otherwise.

### GetValidationWarningsOk

`func (o *IssuesJqlMetaDataBean) GetValidationWarningsOk() (*[]string, bool)`

GetValidationWarningsOk returns a tuple with the ValidationWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationWarnings

`func (o *IssuesJqlMetaDataBean) SetValidationWarnings(v []string)`

SetValidationWarnings sets ValidationWarnings field to given value.

### HasValidationWarnings

`func (o *IssuesJqlMetaDataBean) HasValidationWarnings() bool`

HasValidationWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


