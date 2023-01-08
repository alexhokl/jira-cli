# BulkChangelogResponseBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueChangeLogs** | Pointer to [**[]IssueChangeLog**](IssueChangeLog.md) | The list of issues changelogs. | [optional] [readonly] 
**NextPageToken** | Pointer to **string** | Continuation token to fetch the next page. If this result represents the last or the only page, this token will be null. | [optional] [readonly] 

## Methods

### NewBulkChangelogResponseBean

`func NewBulkChangelogResponseBean() *BulkChangelogResponseBean`

NewBulkChangelogResponseBean instantiates a new BulkChangelogResponseBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkChangelogResponseBeanWithDefaults

`func NewBulkChangelogResponseBeanWithDefaults() *BulkChangelogResponseBean`

NewBulkChangelogResponseBeanWithDefaults instantiates a new BulkChangelogResponseBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueChangeLogs

`func (o *BulkChangelogResponseBean) GetIssueChangeLogs() []IssueChangeLog`

GetIssueChangeLogs returns the IssueChangeLogs field if non-nil, zero value otherwise.

### GetIssueChangeLogsOk

`func (o *BulkChangelogResponseBean) GetIssueChangeLogsOk() (*[]IssueChangeLog, bool)`

GetIssueChangeLogsOk returns a tuple with the IssueChangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueChangeLogs

`func (o *BulkChangelogResponseBean) SetIssueChangeLogs(v []IssueChangeLog)`

SetIssueChangeLogs sets IssueChangeLogs field to given value.

### HasIssueChangeLogs

`func (o *BulkChangelogResponseBean) HasIssueChangeLogs() bool`

HasIssueChangeLogs returns a boolean if a field has been set.

### GetNextPageToken

`func (o *BulkChangelogResponseBean) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BulkChangelogResponseBean) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BulkChangelogResponseBean) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BulkChangelogResponseBean) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


