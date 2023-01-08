# SecuritySchemeWithProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLevel** | Pointer to **int64** | The default level ID of the issue security scheme. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the issue security scheme. | [optional] [readonly] 
**Id** | **int64** | The ID of the issue security scheme. | [readonly] 
**Name** | **string** | The name of the issue security scheme. | [readonly] 
**ProjectIds** | Pointer to **[]int64** | The list of project IDs associated with the issue security scheme. | [optional] [readonly] 
**Self** | **string** | The URL of the issue security scheme. | [readonly] 

## Methods

### NewSecuritySchemeWithProjects

`func NewSecuritySchemeWithProjects(id int64, name string, self string, ) *SecuritySchemeWithProjects`

NewSecuritySchemeWithProjects instantiates a new SecuritySchemeWithProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeWithProjectsWithDefaults

`func NewSecuritySchemeWithProjectsWithDefaults() *SecuritySchemeWithProjects`

NewSecuritySchemeWithProjectsWithDefaults instantiates a new SecuritySchemeWithProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLevel

`func (o *SecuritySchemeWithProjects) GetDefaultLevel() int64`

GetDefaultLevel returns the DefaultLevel field if non-nil, zero value otherwise.

### GetDefaultLevelOk

`func (o *SecuritySchemeWithProjects) GetDefaultLevelOk() (*int64, bool)`

GetDefaultLevelOk returns a tuple with the DefaultLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLevel

`func (o *SecuritySchemeWithProjects) SetDefaultLevel(v int64)`

SetDefaultLevel sets DefaultLevel field to given value.

### HasDefaultLevel

`func (o *SecuritySchemeWithProjects) HasDefaultLevel() bool`

HasDefaultLevel returns a boolean if a field has been set.

### GetDescription

`func (o *SecuritySchemeWithProjects) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySchemeWithProjects) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySchemeWithProjects) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecuritySchemeWithProjects) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SecuritySchemeWithProjects) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecuritySchemeWithProjects) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecuritySchemeWithProjects) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *SecuritySchemeWithProjects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySchemeWithProjects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySchemeWithProjects) SetName(v string)`

SetName sets Name field to given value.


### GetProjectIds

`func (o *SecuritySchemeWithProjects) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SecuritySchemeWithProjects) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SecuritySchemeWithProjects) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SecuritySchemeWithProjects) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetSelf

`func (o *SecuritySchemeWithProjects) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SecuritySchemeWithProjects) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SecuritySchemeWithProjects) SetSelf(v string)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


