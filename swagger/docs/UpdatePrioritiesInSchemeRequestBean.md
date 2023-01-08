# UpdatePrioritiesInSchemeRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to [**PrioritySchemeChangesWithoutMappings**](PrioritySchemeChangesWithoutMappings.md) | Priorities to add to a scheme | [optional] 
**Remove** | Pointer to [**PrioritySchemeChangesWithoutMappings**](PrioritySchemeChangesWithoutMappings.md) | Priorities to remove from a scheme | [optional] 

## Methods

### NewUpdatePrioritiesInSchemeRequestBean

`func NewUpdatePrioritiesInSchemeRequestBean() *UpdatePrioritiesInSchemeRequestBean`

NewUpdatePrioritiesInSchemeRequestBean instantiates a new UpdatePrioritiesInSchemeRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrioritiesInSchemeRequestBeanWithDefaults

`func NewUpdatePrioritiesInSchemeRequestBeanWithDefaults() *UpdatePrioritiesInSchemeRequestBean`

NewUpdatePrioritiesInSchemeRequestBeanWithDefaults instantiates a new UpdatePrioritiesInSchemeRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *UpdatePrioritiesInSchemeRequestBean) GetAdd() PrioritySchemeChangesWithoutMappings`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *UpdatePrioritiesInSchemeRequestBean) GetAddOk() (*PrioritySchemeChangesWithoutMappings, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *UpdatePrioritiesInSchemeRequestBean) SetAdd(v PrioritySchemeChangesWithoutMappings)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *UpdatePrioritiesInSchemeRequestBean) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetRemove

`func (o *UpdatePrioritiesInSchemeRequestBean) GetRemove() PrioritySchemeChangesWithoutMappings`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *UpdatePrioritiesInSchemeRequestBean) GetRemoveOk() (*PrioritySchemeChangesWithoutMappings, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *UpdatePrioritiesInSchemeRequestBean) SetRemove(v PrioritySchemeChangesWithoutMappings)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *UpdatePrioritiesInSchemeRequestBean) HasRemove() bool`

HasRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


