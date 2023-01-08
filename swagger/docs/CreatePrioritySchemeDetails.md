# CreatePrioritySchemeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPriorityId** | **int64** | The ID of the default priority for the priority scheme. | 
**Description** | Pointer to **string** | The description of the priority scheme. | [optional] 
**Mappings** | Pointer to [**PriorityMapping**](PriorityMapping.md) | Instructions to migrate the priorities of issues.  &#x60;in&#x60; mappings are used to migrate the priorities of issues to priorities used within the priority scheme.  &#x60;out&#x60; mappings are used to migrate the priorities of issues to priorities not used within the priority scheme.   *  When **priorities** are **added** to the new priority scheme, no mapping needs to be provided as the new priorities are not used by any issues.  *  When **priorities** are **removed** from the new priority scheme, no mapping needs to be provided as the removed priorities are not used by any issues.  *  When **projects** are **added** to the priority scheme, the priorities of issues in those projects might need to be migrated to new priorities used by the priority scheme. This can occur when the current scheme does not use all the priorities in the project(s)&#39; priority scheme(s).           *  An &#x60;in&#x60; mapping must be provided for each of these priorities.  *  When **projects** are **removed** from the priority scheme, no mapping needs to be provided as the removed projects are not using the priorities of the new priority scheme.  For more information on &#x60;in&#x60; and &#x60;out&#x60; mappings, see the child properties documentation for the &#x60;PriorityMapping&#x60; object below. | [optional] 
**Name** | **string** | The name of the priority scheme. Must be unique. | 
**PriorityIds** | **[]int64** | The IDs of priorities in the scheme. | 
**ProjectIds** | Pointer to **[]int64** | The IDs of projects that will use the priority scheme. | [optional] 

## Methods

### NewCreatePrioritySchemeDetails

`func NewCreatePrioritySchemeDetails(defaultPriorityId int64, name string, priorityIds []int64, ) *CreatePrioritySchemeDetails`

NewCreatePrioritySchemeDetails instantiates a new CreatePrioritySchemeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrioritySchemeDetailsWithDefaults

`func NewCreatePrioritySchemeDetailsWithDefaults() *CreatePrioritySchemeDetails`

NewCreatePrioritySchemeDetailsWithDefaults instantiates a new CreatePrioritySchemeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPriorityId

`func (o *CreatePrioritySchemeDetails) GetDefaultPriorityId() int64`

GetDefaultPriorityId returns the DefaultPriorityId field if non-nil, zero value otherwise.

### GetDefaultPriorityIdOk

`func (o *CreatePrioritySchemeDetails) GetDefaultPriorityIdOk() (*int64, bool)`

GetDefaultPriorityIdOk returns a tuple with the DefaultPriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriorityId

`func (o *CreatePrioritySchemeDetails) SetDefaultPriorityId(v int64)`

SetDefaultPriorityId sets DefaultPriorityId field to given value.


### GetDescription

`func (o *CreatePrioritySchemeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePrioritySchemeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePrioritySchemeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePrioritySchemeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMappings

`func (o *CreatePrioritySchemeDetails) GetMappings() PriorityMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *CreatePrioritySchemeDetails) GetMappingsOk() (*PriorityMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *CreatePrioritySchemeDetails) SetMappings(v PriorityMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *CreatePrioritySchemeDetails) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetName

`func (o *CreatePrioritySchemeDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePrioritySchemeDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePrioritySchemeDetails) SetName(v string)`

SetName sets Name field to given value.


### GetPriorityIds

`func (o *CreatePrioritySchemeDetails) GetPriorityIds() []int64`

GetPriorityIds returns the PriorityIds field if non-nil, zero value otherwise.

### GetPriorityIdsOk

`func (o *CreatePrioritySchemeDetails) GetPriorityIdsOk() (*[]int64, bool)`

GetPriorityIdsOk returns a tuple with the PriorityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityIds

`func (o *CreatePrioritySchemeDetails) SetPriorityIds(v []int64)`

SetPriorityIds sets PriorityIds field to given value.


### GetProjectIds

`func (o *CreatePrioritySchemeDetails) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreatePrioritySchemeDetails) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreatePrioritySchemeDetails) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreatePrioritySchemeDetails) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


