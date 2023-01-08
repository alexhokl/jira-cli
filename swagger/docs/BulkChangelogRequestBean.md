# BulkChangelogRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldIds** | Pointer to **[]string** | List of field IDs to filter changelogs | [optional] 
**IssueIdsOrKeys** | **[]string** | List of issue IDs/keys to fetch changelogs for | 
**MaxResults** | Pointer to **int32** | The maximum number of items to return per page | [optional] [default to 1000]
**NextPageToken** | Pointer to **string** | The cursor for pagination | [optional] 

## Methods

### NewBulkChangelogRequestBean

`func NewBulkChangelogRequestBean(issueIdsOrKeys []string, ) *BulkChangelogRequestBean`

NewBulkChangelogRequestBean instantiates a new BulkChangelogRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkChangelogRequestBeanWithDefaults

`func NewBulkChangelogRequestBeanWithDefaults() *BulkChangelogRequestBean`

NewBulkChangelogRequestBeanWithDefaults instantiates a new BulkChangelogRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldIds

`func (o *BulkChangelogRequestBean) GetFieldIds() []string`

GetFieldIds returns the FieldIds field if non-nil, zero value otherwise.

### GetFieldIdsOk

`func (o *BulkChangelogRequestBean) GetFieldIdsOk() (*[]string, bool)`

GetFieldIdsOk returns a tuple with the FieldIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldIds

`func (o *BulkChangelogRequestBean) SetFieldIds(v []string)`

SetFieldIds sets FieldIds field to given value.

### HasFieldIds

`func (o *BulkChangelogRequestBean) HasFieldIds() bool`

HasFieldIds returns a boolean if a field has been set.

### GetIssueIdsOrKeys

`func (o *BulkChangelogRequestBean) GetIssueIdsOrKeys() []string`

GetIssueIdsOrKeys returns the IssueIdsOrKeys field if non-nil, zero value otherwise.

### GetIssueIdsOrKeysOk

`func (o *BulkChangelogRequestBean) GetIssueIdsOrKeysOk() (*[]string, bool)`

GetIssueIdsOrKeysOk returns a tuple with the IssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueIdsOrKeys

`func (o *BulkChangelogRequestBean) SetIssueIdsOrKeys(v []string)`

SetIssueIdsOrKeys sets IssueIdsOrKeys field to given value.


### GetMaxResults

`func (o *BulkChangelogRequestBean) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *BulkChangelogRequestBean) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *BulkChangelogRequestBean) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *BulkChangelogRequestBean) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPageToken

`func (o *BulkChangelogRequestBean) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BulkChangelogRequestBean) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BulkChangelogRequestBean) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BulkChangelogRequestBean) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


