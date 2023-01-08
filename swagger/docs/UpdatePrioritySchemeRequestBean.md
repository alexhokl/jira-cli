# UpdatePrioritySchemeRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPriorityId** | Pointer to **int64** | The default priority of the scheme. | [optional] 
**Description** | Pointer to **string** | The description of the priority scheme. | [optional] 
**Mappings** | Pointer to [**PriorityMapping**](PriorityMapping.md) | Instructions to migrate the priorities of issues.  &#x60;in&#x60; mappings are used to migrate the priorities of issues to priorities used within the priority scheme.  &#x60;out&#x60; mappings are used to migrate the priorities of issues to priorities not used within the priority scheme.   *  When **priorities** are **added** to the priority scheme, no mapping needs to be provided as the new priorities are not used by any issues.  *  When **priorities** are **removed** from the priority scheme, issues that are using those priorities must be migrated to new priorities used by the priority scheme.           *  An &#x60;in&#x60; mapping must be provided for each of these priorities.  *  When **projects** are **added** to the priority scheme, the priorities of issues in those projects might need to be migrated to new priorities used by the priority scheme. This can occur when the current scheme does not use all the priorities in the project(s)&#39; priority scheme(s).           *  An &#x60;in&#x60; mapping must be provided for each of these priorities.  *  When **projects** are **removed** from the priority scheme, the priorities of issues in those projects might need to be migrated to new priorities within the **Default Priority Scheme** that are not used by the priority scheme. This can occur when the **Default Priority Scheme** does not use all the priorities within the current scheme.           *  An &#x60;out&#x60; mapping must be provided for each of these priorities.  For more information on &#x60;in&#x60; and &#x60;out&#x60; mappings, see the child properties documentation for the &#x60;PriorityMapping&#x60; object below. | [optional] 
**Name** | Pointer to **string** | The name of the priority scheme. Must be unique. | [optional] 
**Priorities** | Pointer to [**UpdatePrioritiesInSchemeRequestBean**](UpdatePrioritiesInSchemeRequestBean.md) | The priorities in the scheme. | [optional] 
**Projects** | Pointer to [**UpdateProjectsInSchemeRequestBean**](UpdateProjectsInSchemeRequestBean.md) | The projects in the scheme. | [optional] 

## Methods

### NewUpdatePrioritySchemeRequestBean

`func NewUpdatePrioritySchemeRequestBean() *UpdatePrioritySchemeRequestBean`

NewUpdatePrioritySchemeRequestBean instantiates a new UpdatePrioritySchemeRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrioritySchemeRequestBeanWithDefaults

`func NewUpdatePrioritySchemeRequestBeanWithDefaults() *UpdatePrioritySchemeRequestBean`

NewUpdatePrioritySchemeRequestBeanWithDefaults instantiates a new UpdatePrioritySchemeRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPriorityId

`func (o *UpdatePrioritySchemeRequestBean) GetDefaultPriorityId() int64`

GetDefaultPriorityId returns the DefaultPriorityId field if non-nil, zero value otherwise.

### GetDefaultPriorityIdOk

`func (o *UpdatePrioritySchemeRequestBean) GetDefaultPriorityIdOk() (*int64, bool)`

GetDefaultPriorityIdOk returns a tuple with the DefaultPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriorityId

`func (o *UpdatePrioritySchemeRequestBean) SetDefaultPriorityId(v int64)`

SetDefaultPriorityId sets DefaultPriorityId field to given value.

### HasDefaultPriorityId

`func (o *UpdatePrioritySchemeRequestBean) HasDefaultPriorityId() bool`

HasDefaultPriorityId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePrioritySchemeRequestBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePrioritySchemeRequestBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePrioritySchemeRequestBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePrioritySchemeRequestBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMappings

`func (o *UpdatePrioritySchemeRequestBean) GetMappings() PriorityMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *UpdatePrioritySchemeRequestBean) GetMappingsOk() (*PriorityMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *UpdatePrioritySchemeRequestBean) SetMappings(v PriorityMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *UpdatePrioritySchemeRequestBean) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetName

`func (o *UpdatePrioritySchemeRequestBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePrioritySchemeRequestBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePrioritySchemeRequestBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePrioritySchemeRequestBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriorities

`func (o *UpdatePrioritySchemeRequestBean) GetPriorities() UpdatePrioritiesInSchemeRequestBean`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *UpdatePrioritySchemeRequestBean) GetPrioritiesOk() (*UpdatePrioritiesInSchemeRequestBean, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *UpdatePrioritySchemeRequestBean) SetPriorities(v UpdatePrioritiesInSchemeRequestBean)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *UpdatePrioritySchemeRequestBean) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### GetProjects

`func (o *UpdatePrioritySchemeRequestBean) GetProjects() UpdateProjectsInSchemeRequestBean`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UpdatePrioritySchemeRequestBean) GetProjectsOk() (*UpdateProjectsInSchemeRequestBean, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UpdatePrioritySchemeRequestBean) SetProjects(v UpdateProjectsInSchemeRequestBean)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UpdatePrioritySchemeRequestBean) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


