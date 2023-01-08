# SuggestedMappingsRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] 
**Priorities** | Pointer to [**SuggestedMappingsForPrioritiesRequestBean**](SuggestedMappingsForPrioritiesRequestBean.md) | The priority changes in the scheme. | [optional] 
**Projects** | Pointer to [**SuggestedMappingsForProjectsRequestBean**](SuggestedMappingsForProjectsRequestBean.md) | The project changes in the scheme. | [optional] 
**SchemeId** | Pointer to **int64** | The id of the priority scheme. | [optional] 
**StartAt** | Pointer to **int64** | The index of the first item returned on the page. | [optional] 

## Methods

### NewSuggestedMappingsRequestBean

`func NewSuggestedMappingsRequestBean() *SuggestedMappingsRequestBean`

NewSuggestedMappingsRequestBean instantiates a new SuggestedMappingsRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedMappingsRequestBeanWithDefaults

`func NewSuggestedMappingsRequestBeanWithDefaults() *SuggestedMappingsRequestBean`

NewSuggestedMappingsRequestBeanWithDefaults instantiates a new SuggestedMappingsRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *SuggestedMappingsRequestBean) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SuggestedMappingsRequestBean) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SuggestedMappingsRequestBean) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SuggestedMappingsRequestBean) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetPriorities

`func (o *SuggestedMappingsRequestBean) GetPriorities() SuggestedMappingsForPrioritiesRequestBean`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *SuggestedMappingsRequestBean) GetPrioritiesOk() (*SuggestedMappingsForPrioritiesRequestBean, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *SuggestedMappingsRequestBean) SetPriorities(v SuggestedMappingsForPrioritiesRequestBean)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *SuggestedMappingsRequestBean) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### GetProjects

`func (o *SuggestedMappingsRequestBean) GetProjects() SuggestedMappingsForProjectsRequestBean`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *SuggestedMappingsRequestBean) GetProjectsOk() (*SuggestedMappingsForProjectsRequestBean, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *SuggestedMappingsRequestBean) SetProjects(v SuggestedMappingsForProjectsRequestBean)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *SuggestedMappingsRequestBean) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetSchemeId

`func (o *SuggestedMappingsRequestBean) GetSchemeId() int64`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *SuggestedMappingsRequestBean) GetSchemeIdOk() (*int64, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *SuggestedMappingsRequestBean) SetSchemeId(v int64)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *SuggestedMappingsRequestBean) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.

### GetStartAt

`func (o *SuggestedMappingsRequestBean) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *SuggestedMappingsRequestBean) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *SuggestedMappingsRequestBean) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *SuggestedMappingsRequestBean) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


