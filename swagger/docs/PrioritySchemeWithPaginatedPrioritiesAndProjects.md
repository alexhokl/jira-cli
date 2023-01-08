# PrioritySchemeWithPaginatedPrioritiesAndProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** |  | [optional] 
**DefaultPriorityId** | Pointer to **string** | The ID of the default issue priority. | [optional] 
**Description** | Pointer to **string** | The description of the priority scheme | [optional] 
**Id** | **string** | The ID of the priority scheme. | 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | **string** | The name of the priority scheme | 
**Priorities** | Pointer to [**PageBeanPriorityWithSequence**](PageBeanPriorityWithSequence.md) | The paginated list of priorities. | [optional] 
**Projects** | Pointer to [**PageBeanProjectDetails**](PageBeanProjectDetails.md) | The paginated list of projects. | [optional] 
**Self** | Pointer to **string** | The URL of the priority scheme. | [optional] 

## Methods

### NewPrioritySchemeWithPaginatedPrioritiesAndProjects

`func NewPrioritySchemeWithPaginatedPrioritiesAndProjects(id string, name string, ) *PrioritySchemeWithPaginatedPrioritiesAndProjects`

NewPrioritySchemeWithPaginatedPrioritiesAndProjects instantiates a new PrioritySchemeWithPaginatedPrioritiesAndProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritySchemeWithPaginatedPrioritiesAndProjectsWithDefaults

`func NewPrioritySchemeWithPaginatedPrioritiesAndProjectsWithDefaults() *PrioritySchemeWithPaginatedPrioritiesAndProjects`

NewPrioritySchemeWithPaginatedPrioritiesAndProjectsWithDefaults instantiates a new PrioritySchemeWithPaginatedPrioritiesAndProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultPriorityId

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDefaultPriorityId() string`

GetDefaultPriorityId returns the DefaultPriorityId field if non-nil, zero value otherwise.

### GetDefaultPriorityIdOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDefaultPriorityIdOk() (*string, bool)`

GetDefaultPriorityIdOk returns a tuple with the DefaultPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriorityId

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetDefaultPriorityId(v string)`

SetDefaultPriorityId sets DefaultPriorityId field to given value.

### HasDefaultPriorityId

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasDefaultPriorityId() bool`

HasDefaultPriorityId returns a boolean if a field has been set.

### GetDescription

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetName(v string)`

SetName sets Name field to given value.


### GetPriorities

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetPriorities() PageBeanPriorityWithSequence`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetPrioritiesOk() (*PageBeanPriorityWithSequence, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetPriorities(v PageBeanPriorityWithSequence)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### GetProjects

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetProjects() PageBeanProjectDetails`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetProjectsOk() (*PageBeanProjectDetails, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetProjects(v PageBeanProjectDetails)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSelf

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PrioritySchemeWithPaginatedPrioritiesAndProjects) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


