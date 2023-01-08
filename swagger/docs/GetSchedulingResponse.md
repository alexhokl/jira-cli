# GetSchedulingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | **string** | The dependencies for the plan. This is \&quot;Sequential\&quot; or \&quot;Concurrent\&quot;. | 
**EndDate** | [**GetDateFieldResponse**](GetDateFieldResponse.md) | The end date field for the plan. | 
**Estimation** | **string** | The estimation unit for the plan. This is \&quot;StoryPoints\&quot;, \&quot;Days\&quot; or \&quot;Hours\&quot;. | 
**InferredDates** | **string** | The inferred dates for the plan. This is \&quot;None\&quot;, \&quot;SprintDates\&quot; or \&quot;ReleaseDates\&quot;. | 
**StartDate** | [**GetDateFieldResponse**](GetDateFieldResponse.md) | The start date field for the plan. | 

## Methods

### NewGetSchedulingResponse

`func NewGetSchedulingResponse(dependencies string, endDate GetDateFieldResponse, estimation string, inferredDates string, startDate GetDateFieldResponse, ) *GetSchedulingResponse`

NewGetSchedulingResponse instantiates a new GetSchedulingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchedulingResponseWithDefaults

`func NewGetSchedulingResponseWithDefaults() *GetSchedulingResponse`

NewGetSchedulingResponseWithDefaults instantiates a new GetSchedulingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *GetSchedulingResponse) GetDependencies() string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *GetSchedulingResponse) GetDependenciesOk() (*string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *GetSchedulingResponse) SetDependencies(v string)`

SetDependencies sets Dependencies field to given value.


### GetEndDate

`func (o *GetSchedulingResponse) GetEndDate() GetDateFieldResponse`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetSchedulingResponse) GetEndDateOk() (*GetDateFieldResponse, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetSchedulingResponse) SetEndDate(v GetDateFieldResponse)`

SetEndDate sets EndDate field to given value.


### GetEstimation

`func (o *GetSchedulingResponse) GetEstimation() string`

GetEstimation returns the Estimation field if non-nil, zero value otherwise.

### GetEstimationOk

`func (o *GetSchedulingResponse) GetEstimationOk() (*string, bool)`

GetEstimationOk returns a tuple with the Estimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimation

`func (o *GetSchedulingResponse) SetEstimation(v string)`

SetEstimation sets Estimation field to given value.


### GetInferredDates

`func (o *GetSchedulingResponse) GetInferredDates() string`

GetInferredDates returns the InferredDates field if non-nil, zero value otherwise.

### GetInferredDatesOk

`func (o *GetSchedulingResponse) GetInferredDatesOk() (*string, bool)`

GetInferredDatesOk returns a tuple with the InferredDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredDates

`func (o *GetSchedulingResponse) SetInferredDates(v string)`

SetInferredDates sets InferredDates field to given value.


### GetStartDate

`func (o *GetSchedulingResponse) GetStartDate() GetDateFieldResponse`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetSchedulingResponse) GetStartDateOk() (*GetDateFieldResponse, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetSchedulingResponse) SetStartDate(v GetDateFieldResponse)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


