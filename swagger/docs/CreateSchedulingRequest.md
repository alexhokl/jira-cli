# CreateSchedulingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | Pointer to **string** | The dependencies for the plan. This must be \&quot;Sequential\&quot; or \&quot;Concurrent\&quot;. | [optional] 
**EndDate** | Pointer to [**CreateDateFieldRequest**](CreateDateFieldRequest.md) | The end date field for the plan. | [optional] 
**Estimation** | **string** | The estimation unit for the plan. This must be \&quot;StoryPoints\&quot;, \&quot;Days\&quot; or \&quot;Hours\&quot;. | 
**InferredDates** | Pointer to **string** | The inferred dates for the plan. This must be \&quot;None\&quot;, \&quot;SprintDates\&quot; or \&quot;ReleaseDates\&quot;. | [optional] 
**StartDate** | Pointer to [**CreateDateFieldRequest**](CreateDateFieldRequest.md) | The start date field for the plan. | [optional] 

## Methods

### NewCreateSchedulingRequest

`func NewCreateSchedulingRequest(estimation string, ) *CreateSchedulingRequest`

NewCreateSchedulingRequest instantiates a new CreateSchedulingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSchedulingRequestWithDefaults

`func NewCreateSchedulingRequestWithDefaults() *CreateSchedulingRequest`

NewCreateSchedulingRequestWithDefaults instantiates a new CreateSchedulingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *CreateSchedulingRequest) GetDependencies() string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *CreateSchedulingRequest) GetDependenciesOk() (*string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *CreateSchedulingRequest) SetDependencies(v string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *CreateSchedulingRequest) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateSchedulingRequest) GetEndDate() CreateDateFieldRequest`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateSchedulingRequest) GetEndDateOk() (*CreateDateFieldRequest, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateSchedulingRequest) SetEndDate(v CreateDateFieldRequest)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateSchedulingRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEstimation

`func (o *CreateSchedulingRequest) GetEstimation() string`

GetEstimation returns the Estimation field if non-nil, zero value otherwise.

### GetEstimationOk

`func (o *CreateSchedulingRequest) GetEstimationOk() (*string, bool)`

GetEstimationOk returns a tuple with the Estimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimation

`func (o *CreateSchedulingRequest) SetEstimation(v string)`

SetEstimation sets Estimation field to given value.


### GetInferredDates

`func (o *CreateSchedulingRequest) GetInferredDates() string`

GetInferredDates returns the InferredDates field if non-nil, zero value otherwise.

### GetInferredDatesOk

`func (o *CreateSchedulingRequest) GetInferredDatesOk() (*string, bool)`

GetInferredDatesOk returns a tuple with the InferredDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredDates

`func (o *CreateSchedulingRequest) SetInferredDates(v string)`

SetInferredDates sets InferredDates field to given value.

### HasInferredDates

`func (o *CreateSchedulingRequest) HasInferredDates() bool`

HasInferredDates returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateSchedulingRequest) GetStartDate() CreateDateFieldRequest`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateSchedulingRequest) GetStartDateOk() (*CreateDateFieldRequest, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateSchedulingRequest) SetStartDate(v CreateDateFieldRequest)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateSchedulingRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


