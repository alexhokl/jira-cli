# StatusWorkflowUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **string** | The status ID. | [optional] 
**Workflows** | Pointer to [**StatusWorkflowUsagePage**](StatusWorkflowUsagePage.md) |  | [optional] 

## Methods

### NewStatusWorkflowUsageDTO

`func NewStatusWorkflowUsageDTO() *StatusWorkflowUsageDTO`

NewStatusWorkflowUsageDTO instantiates a new StatusWorkflowUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWorkflowUsageDTOWithDefaults

`func NewStatusWorkflowUsageDTOWithDefaults() *StatusWorkflowUsageDTO`

NewStatusWorkflowUsageDTOWithDefaults instantiates a new StatusWorkflowUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *StatusWorkflowUsageDTO) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *StatusWorkflowUsageDTO) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *StatusWorkflowUsageDTO) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *StatusWorkflowUsageDTO) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetWorkflows

`func (o *StatusWorkflowUsageDTO) GetWorkflows() StatusWorkflowUsagePage`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *StatusWorkflowUsageDTO) GetWorkflowsOk() (*StatusWorkflowUsagePage, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *StatusWorkflowUsageDTO) SetWorkflows(v StatusWorkflowUsagePage)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *StatusWorkflowUsageDTO) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


