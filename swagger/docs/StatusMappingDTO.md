# StatusMappingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeId** | **string** | The issue type for the status mapping. | 
**ProjectId** | **string** | The project for the status mapping. | 
**StatusMigrations** | [**[]StatusMigration**](StatusMigration.md) | The list of old and new status ID mappings for the specified project and issue type. | 

## Methods

### NewStatusMappingDTO

`func NewStatusMappingDTO(issueTypeId string, projectId string, statusMigrations []StatusMigration, ) *StatusMappingDTO`

NewStatusMappingDTO instantiates a new StatusMappingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusMappingDTOWithDefaults

`func NewStatusMappingDTOWithDefaults() *StatusMappingDTO`

NewStatusMappingDTOWithDefaults instantiates a new StatusMappingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeId

`func (o *StatusMappingDTO) GetIssueTypeId() string`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *StatusMappingDTO) GetIssueTypeIdOk() (*string, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *StatusMappingDTO) SetIssueTypeId(v string)`

SetIssueTypeId sets IssueTypeId field to given value.


### GetProjectId

`func (o *StatusMappingDTO) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StatusMappingDTO) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StatusMappingDTO) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStatusMigrations

`func (o *StatusMappingDTO) GetStatusMigrations() []StatusMigration`

GetStatusMigrations returns the StatusMigrations field if non-nil, zero value otherwise.

### GetStatusMigrationsOk

`func (o *StatusMappingDTO) GetStatusMigrationsOk() (*[]StatusMigration, bool)`

GetStatusMigrationsOk returns a tuple with the StatusMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMigrations

`func (o *StatusMappingDTO) SetStatusMigrations(v []StatusMigration)`

SetStatusMigrations sets StatusMigrations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


