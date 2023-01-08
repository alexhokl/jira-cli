# StatusProjectIssueTypeUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypes** | Pointer to [**StatusProjectIssueTypeUsagePage**](StatusProjectIssueTypeUsagePage.md) |  | [optional] 
**ProjectId** | Pointer to **string** | The project ID. | [optional] 
**StatusId** | Pointer to **string** | The status ID. | [optional] 

## Methods

### NewStatusProjectIssueTypeUsageDTO

`func NewStatusProjectIssueTypeUsageDTO() *StatusProjectIssueTypeUsageDTO`

NewStatusProjectIssueTypeUsageDTO instantiates a new StatusProjectIssueTypeUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusProjectIssueTypeUsageDTOWithDefaults

`func NewStatusProjectIssueTypeUsageDTOWithDefaults() *StatusProjectIssueTypeUsageDTO`

NewStatusProjectIssueTypeUsageDTOWithDefaults instantiates a new StatusProjectIssueTypeUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypes

`func (o *StatusProjectIssueTypeUsageDTO) GetIssueTypes() StatusProjectIssueTypeUsagePage`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *StatusProjectIssueTypeUsageDTO) GetIssueTypesOk() (*StatusProjectIssueTypeUsagePage, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *StatusProjectIssueTypeUsageDTO) SetIssueTypes(v StatusProjectIssueTypeUsagePage)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *StatusProjectIssueTypeUsageDTO) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetProjectId

`func (o *StatusProjectIssueTypeUsageDTO) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StatusProjectIssueTypeUsageDTO) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StatusProjectIssueTypeUsageDTO) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *StatusProjectIssueTypeUsageDTO) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatusId

`func (o *StatusProjectIssueTypeUsageDTO) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *StatusProjectIssueTypeUsageDTO) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *StatusProjectIssueTypeUsageDTO) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *StatusProjectIssueTypeUsageDTO) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


