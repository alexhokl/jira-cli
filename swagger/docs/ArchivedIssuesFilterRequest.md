# ArchivedIssuesFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedBy** | Pointer to **[]string** | List archived issues archived by a specified account ID. | [optional] 
**ArchivedDateRange** | Pointer to [**DateRangeFilterRequest**](DateRangeFilterRequest.md) |  | [optional] 
**IssueTypes** | Pointer to **[]string** | List archived issues with a specified issue type ID. | [optional] 
**Projects** | Pointer to **[]string** | List archived issues with a specified project key. | [optional] 
**Reporters** | Pointer to **[]string** | List archived issues where the reporter is a specified account ID. | [optional] 

## Methods

### NewArchivedIssuesFilterRequest

`func NewArchivedIssuesFilterRequest() *ArchivedIssuesFilterRequest`

NewArchivedIssuesFilterRequest instantiates a new ArchivedIssuesFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivedIssuesFilterRequestWithDefaults

`func NewArchivedIssuesFilterRequestWithDefaults() *ArchivedIssuesFilterRequest`

NewArchivedIssuesFilterRequestWithDefaults instantiates a new ArchivedIssuesFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedBy

`func (o *ArchivedIssuesFilterRequest) GetArchivedBy() []string`

GetArchivedBy returns the ArchivedBy field if non-nil, zero value otherwise.

### GetArchivedByOk

`func (o *ArchivedIssuesFilterRequest) GetArchivedByOk() (*[]string, bool)`

GetArchivedByOk returns a tuple with the ArchivedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedBy

`func (o *ArchivedIssuesFilterRequest) SetArchivedBy(v []string)`

SetArchivedBy sets ArchivedBy field to given value.

### HasArchivedBy

`func (o *ArchivedIssuesFilterRequest) HasArchivedBy() bool`

HasArchivedBy returns a boolean if a field has been set.

### GetArchivedDateRange

`func (o *ArchivedIssuesFilterRequest) GetArchivedDateRange() DateRangeFilterRequest`

GetArchivedDateRange returns the ArchivedDateRange field if non-nil, zero value otherwise.

### GetArchivedDateRangeOk

`func (o *ArchivedIssuesFilterRequest) GetArchivedDateRangeOk() (*DateRangeFilterRequest, bool)`

GetArchivedDateRangeOk returns a tuple with the ArchivedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDateRange

`func (o *ArchivedIssuesFilterRequest) SetArchivedDateRange(v DateRangeFilterRequest)`

SetArchivedDateRange sets ArchivedDateRange field to given value.

### HasArchivedDateRange

`func (o *ArchivedIssuesFilterRequest) HasArchivedDateRange() bool`

HasArchivedDateRange returns a boolean if a field has been set.

### GetIssueTypes

`func (o *ArchivedIssuesFilterRequest) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *ArchivedIssuesFilterRequest) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *ArchivedIssuesFilterRequest) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *ArchivedIssuesFilterRequest) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetProjects

`func (o *ArchivedIssuesFilterRequest) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ArchivedIssuesFilterRequest) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ArchivedIssuesFilterRequest) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ArchivedIssuesFilterRequest) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetReporters

`func (o *ArchivedIssuesFilterRequest) GetReporters() []string`

GetReporters returns the Reporters field if non-nil, zero value otherwise.

### GetReportersOk

`func (o *ArchivedIssuesFilterRequest) GetReportersOk() (*[]string, bool)`

GetReportersOk returns a tuple with the Reporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporters

`func (o *ArchivedIssuesFilterRequest) SetReporters(v []string)`

SetReporters sets Reporters field to given value.

### HasReporters

`func (o *ArchivedIssuesFilterRequest) HasReporters() bool`

HasReporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


