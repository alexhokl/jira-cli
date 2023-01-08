# UpdatePrioritySchemeResponseBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityScheme** | Pointer to [**PrioritySchemeWithPaginatedPrioritiesAndProjects**](PrioritySchemeWithPaginatedPrioritiesAndProjects.md) |  | [optional] 
**Task** | Pointer to [**TaskProgressBeanJsonNode**](TaskProgressBeanJsonNode.md) | The in-progress issue migration task. | [optional] [readonly] 

## Methods

### NewUpdatePrioritySchemeResponseBean

`func NewUpdatePrioritySchemeResponseBean() *UpdatePrioritySchemeResponseBean`

NewUpdatePrioritySchemeResponseBean instantiates a new UpdatePrioritySchemeResponseBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrioritySchemeResponseBeanWithDefaults

`func NewUpdatePrioritySchemeResponseBeanWithDefaults() *UpdatePrioritySchemeResponseBean`

NewUpdatePrioritySchemeResponseBeanWithDefaults instantiates a new UpdatePrioritySchemeResponseBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityScheme

`func (o *UpdatePrioritySchemeResponseBean) GetPriorityScheme() PrioritySchemeWithPaginatedPrioritiesAndProjects`

GetPriorityScheme returns the PriorityScheme field if non-nil, zero value otherwise.

### GetPrioritySchemeOk

`func (o *UpdatePrioritySchemeResponseBean) GetPrioritySchemeOk() (*PrioritySchemeWithPaginatedPrioritiesAndProjects, bool)`

GetPrioritySchemeOk returns a tuple with the PriorityScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityScheme

`func (o *UpdatePrioritySchemeResponseBean) SetPriorityScheme(v PrioritySchemeWithPaginatedPrioritiesAndProjects)`

SetPriorityScheme sets PriorityScheme field to given value.

### HasPriorityScheme

`func (o *UpdatePrioritySchemeResponseBean) HasPriorityScheme() bool`

HasPriorityScheme returns a boolean if a field has been set.

### GetTask

`func (o *UpdatePrioritySchemeResponseBean) GetTask() TaskProgressBeanJsonNode`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UpdatePrioritySchemeResponseBean) GetTaskOk() (*TaskProgressBeanJsonNode, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UpdatePrioritySchemeResponseBean) SetTask(v TaskProgressBeanJsonNode)`

SetTask sets Task field to given value.

### HasTask

`func (o *UpdatePrioritySchemeResponseBean) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


