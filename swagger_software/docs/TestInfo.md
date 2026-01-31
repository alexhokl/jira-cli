# TestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNumber** | **int64** | The total number of tests considered during a build.  | 
**NumberPassed** | **int64** | The number of tests that passed during a build.  | 
**NumberFailed** | **int64** | The number of tests that failed during a build.  | 
**NumberSkipped** | Pointer to **int64** | The number of tests that were skipped during a build.  | [optional] [default to 0]

## Methods

### NewTestInfo

`func NewTestInfo(totalNumber int64, numberPassed int64, numberFailed int64, ) *TestInfo`

NewTestInfo instantiates a new TestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoWithDefaults

`func NewTestInfoWithDefaults() *TestInfo`

NewTestInfoWithDefaults instantiates a new TestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNumber

`func (o *TestInfo) GetTotalNumber() int64`

GetTotalNumber returns the TotalNumber field if non-nil, zero value otherwise.

### GetTotalNumberOk

`func (o *TestInfo) GetTotalNumberOk() (*int64, bool)`

GetTotalNumberOk returns a tuple with the TotalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumber

`func (o *TestInfo) SetTotalNumber(v int64)`

SetTotalNumber sets TotalNumber field to given value.


### GetNumberPassed

`func (o *TestInfo) GetNumberPassed() int64`

GetNumberPassed returns the NumberPassed field if non-nil, zero value otherwise.

### GetNumberPassedOk

`func (o *TestInfo) GetNumberPassedOk() (*int64, bool)`

GetNumberPassedOk returns a tuple with the NumberPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberPassed

`func (o *TestInfo) SetNumberPassed(v int64)`

SetNumberPassed sets NumberPassed field to given value.


### GetNumberFailed

`func (o *TestInfo) GetNumberFailed() int64`

GetNumberFailed returns the NumberFailed field if non-nil, zero value otherwise.

### GetNumberFailedOk

`func (o *TestInfo) GetNumberFailedOk() (*int64, bool)`

GetNumberFailedOk returns a tuple with the NumberFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFailed

`func (o *TestInfo) SetNumberFailed(v int64)`

SetNumberFailed sets NumberFailed field to given value.


### GetNumberSkipped

`func (o *TestInfo) GetNumberSkipped() int64`

GetNumberSkipped returns the NumberSkipped field if non-nil, zero value otherwise.

### GetNumberSkippedOk

`func (o *TestInfo) GetNumberSkippedOk() (*int64, bool)`

GetNumberSkippedOk returns a tuple with the NumberSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberSkipped

`func (o *TestInfo) SetNumberSkipped(v int64)`

SetNumberSkipped sets NumberSkipped field to given value.

### HasNumberSkipped

`func (o *TestInfo) HasNumberSkipped() bool`

HasNumberSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


