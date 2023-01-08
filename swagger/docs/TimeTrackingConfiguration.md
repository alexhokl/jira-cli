# TimeTrackingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultUnit** | **string** | The default unit of time applied to logged time. | 
**TimeFormat** | **string** | The format that will appear on an issue&#39;s *Time Spent* field. | 
**WorkingDaysPerWeek** | **float64** | The number of days in a working week. | 
**WorkingHoursPerDay** | **float64** | The number of hours in a working day. | 

## Methods

### NewTimeTrackingConfiguration

`func NewTimeTrackingConfiguration(defaultUnit string, timeFormat string, workingDaysPerWeek float64, workingHoursPerDay float64, ) *TimeTrackingConfiguration`

NewTimeTrackingConfiguration instantiates a new TimeTrackingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeTrackingConfigurationWithDefaults

`func NewTimeTrackingConfigurationWithDefaults() *TimeTrackingConfiguration`

NewTimeTrackingConfigurationWithDefaults instantiates a new TimeTrackingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultUnit

`func (o *TimeTrackingConfiguration) GetDefaultUnit() string`

GetDefaultUnit returns the DefaultUnit field if non-nil, zero value otherwise.

### GetDefaultUnitOk

`func (o *TimeTrackingConfiguration) GetDefaultUnitOk() (*string, bool)`

GetDefaultUnitOk returns a tuple with the DefaultUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUnit

`func (o *TimeTrackingConfiguration) SetDefaultUnit(v string)`

SetDefaultUnit sets DefaultUnit field to given value.


### GetTimeFormat

`func (o *TimeTrackingConfiguration) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *TimeTrackingConfiguration) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *TimeTrackingConfiguration) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.


### GetWorkingDaysPerWeek

`func (o *TimeTrackingConfiguration) GetWorkingDaysPerWeek() float64`

GetWorkingDaysPerWeek returns the WorkingDaysPerWeek field if non-nil, zero value otherwise.

### GetWorkingDaysPerWeekOk

`func (o *TimeTrackingConfiguration) GetWorkingDaysPerWeekOk() (*float64, bool)`

GetWorkingDaysPerWeekOk returns a tuple with the WorkingDaysPerWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDaysPerWeek

`func (o *TimeTrackingConfiguration) SetWorkingDaysPerWeek(v float64)`

SetWorkingDaysPerWeek sets WorkingDaysPerWeek field to given value.


### GetWorkingHoursPerDay

`func (o *TimeTrackingConfiguration) GetWorkingHoursPerDay() float64`

GetWorkingHoursPerDay returns the WorkingHoursPerDay field if non-nil, zero value otherwise.

### GetWorkingHoursPerDayOk

`func (o *TimeTrackingConfiguration) GetWorkingHoursPerDayOk() (*float64, bool)`

GetWorkingHoursPerDayOk returns a tuple with the WorkingHoursPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHoursPerDay

`func (o *TimeTrackingConfiguration) SetWorkingHoursPerDay(v float64)`

SetWorkingHoursPerDay sets WorkingHoursPerDay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


