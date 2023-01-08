# HealthCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Jira health check item. | [optional] 
**Name** | Pointer to **string** | The name of the Jira health check item. | [optional] 
**Passed** | Pointer to **bool** | Whether the Jira health check item passed or failed. | [optional] 

## Methods

### NewHealthCheckResult

`func NewHealthCheckResult() *HealthCheckResult`

NewHealthCheckResult instantiates a new HealthCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResultWithDefaults

`func NewHealthCheckResultWithDefaults() *HealthCheckResult`

NewHealthCheckResultWithDefaults instantiates a new HealthCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HealthCheckResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheckResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheckResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheckResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HealthCheckResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthCheckResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassed

`func (o *HealthCheckResult) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *HealthCheckResult) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *HealthCheckResult) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *HealthCheckResult) HasPassed() bool`

HasPassed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


