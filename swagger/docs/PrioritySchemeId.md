# PrioritySchemeId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the priority scheme. | [optional] [readonly] 
**Task** | Pointer to [**TaskProgressBeanJsonNode**](TaskProgressBeanJsonNode.md) | The in-progress issue migration task. | [optional] [readonly] 

## Methods

### NewPrioritySchemeId

`func NewPrioritySchemeId() *PrioritySchemeId`

NewPrioritySchemeId instantiates a new PrioritySchemeId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritySchemeIdWithDefaults

`func NewPrioritySchemeIdWithDefaults() *PrioritySchemeId`

NewPrioritySchemeIdWithDefaults instantiates a new PrioritySchemeId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrioritySchemeId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrioritySchemeId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrioritySchemeId) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrioritySchemeId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *PrioritySchemeId) GetTask() TaskProgressBeanJsonNode`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PrioritySchemeId) GetTaskOk() (*TaskProgressBeanJsonNode, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PrioritySchemeId) SetTask(v TaskProgressBeanJsonNode)`

SetTask sets Task field to given value.

### HasTask

`func (o *PrioritySchemeId) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


