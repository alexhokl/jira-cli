# StatusProjectUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**StatusProjectUsagePage**](StatusProjectUsagePage.md) |  | [optional] 
**StatusId** | Pointer to **string** | The status ID. | [optional] 

## Methods

### NewStatusProjectUsageDTO

`func NewStatusProjectUsageDTO() *StatusProjectUsageDTO`

NewStatusProjectUsageDTO instantiates a new StatusProjectUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusProjectUsageDTOWithDefaults

`func NewStatusProjectUsageDTOWithDefaults() *StatusProjectUsageDTO`

NewStatusProjectUsageDTOWithDefaults instantiates a new StatusProjectUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *StatusProjectUsageDTO) GetProjects() StatusProjectUsagePage`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *StatusProjectUsageDTO) GetProjectsOk() (*StatusProjectUsagePage, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *StatusProjectUsageDTO) SetProjects(v StatusProjectUsagePage)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *StatusProjectUsageDTO) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetStatusId

`func (o *StatusProjectUsageDTO) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *StatusProjectUsageDTO) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *StatusProjectUsageDTO) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *StatusProjectUsageDTO) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


