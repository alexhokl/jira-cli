# GetPlanResponseForPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The plan ID. | 
**IssueSources** | Pointer to [**[]GetIssueSourceResponse**](GetIssueSourceResponse.md) | The issue sources included in the plan. | [optional] 
**Name** | **string** | The plan name. | 
**ScenarioId** | **string** | Default scenario ID. | 
**Status** | **string** | The plan status. This is \&quot;Active\&quot;, \&quot;Trashed\&quot; or \&quot;Archived\&quot;. | 

## Methods

### NewGetPlanResponseForPage

`func NewGetPlanResponseForPage(id string, name string, scenarioId string, status string, ) *GetPlanResponseForPage`

NewGetPlanResponseForPage instantiates a new GetPlanResponseForPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlanResponseForPageWithDefaults

`func NewGetPlanResponseForPageWithDefaults() *GetPlanResponseForPage`

NewGetPlanResponseForPageWithDefaults instantiates a new GetPlanResponseForPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPlanResponseForPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlanResponseForPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlanResponseForPage) SetId(v string)`

SetId sets Id field to given value.


### GetIssueSources

`func (o *GetPlanResponseForPage) GetIssueSources() []GetIssueSourceResponse`

GetIssueSources returns the IssueSources field if non-nil, zero value otherwise.

### GetIssueSourcesOk

`func (o *GetPlanResponseForPage) GetIssueSourcesOk() (*[]GetIssueSourceResponse, bool)`

GetIssueSourcesOk returns a tuple with the IssueSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSources

`func (o *GetPlanResponseForPage) SetIssueSources(v []GetIssueSourceResponse)`

SetIssueSources sets IssueSources field to given value.

### HasIssueSources

`func (o *GetPlanResponseForPage) HasIssueSources() bool`

HasIssueSources returns a boolean if a field has been set.

### GetName

`func (o *GetPlanResponseForPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPlanResponseForPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPlanResponseForPage) SetName(v string)`

SetName sets Name field to given value.


### GetScenarioId

`func (o *GetPlanResponseForPage) GetScenarioId() string`

GetScenarioId returns the ScenarioId field if non-nil, zero value otherwise.

### GetScenarioIdOk

`func (o *GetPlanResponseForPage) GetScenarioIdOk() (*string, bool)`

GetScenarioIdOk returns a tuple with the ScenarioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioId

`func (o *GetPlanResponseForPage) SetScenarioId(v string)`

SetScenarioId sets ScenarioId field to given value.


### GetStatus

`func (o *GetPlanResponseForPage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlanResponseForPage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlanResponseForPage) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


