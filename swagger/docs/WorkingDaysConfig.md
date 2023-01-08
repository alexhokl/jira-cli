# WorkingDaysConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Friday** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Monday** | Pointer to **bool** |  | [optional] 
**NonWorkingDays** | Pointer to [**[]NonWorkingDay**](NonWorkingDay.md) |  | [optional] 
**Saturday** | Pointer to **bool** |  | [optional] 
**Sunday** | Pointer to **bool** |  | [optional] 
**Thursday** | Pointer to **bool** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 
**Tuesday** | Pointer to **bool** |  | [optional] 
**Wednesday** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkingDaysConfig

`func NewWorkingDaysConfig() *WorkingDaysConfig`

NewWorkingDaysConfig instantiates a new WorkingDaysConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkingDaysConfigWithDefaults

`func NewWorkingDaysConfigWithDefaults() *WorkingDaysConfig`

NewWorkingDaysConfigWithDefaults instantiates a new WorkingDaysConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriday

`func (o *WorkingDaysConfig) GetFriday() bool`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *WorkingDaysConfig) GetFridayOk() (*bool, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *WorkingDaysConfig) SetFriday(v bool)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *WorkingDaysConfig) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetId

`func (o *WorkingDaysConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkingDaysConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkingDaysConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WorkingDaysConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMonday

`func (o *WorkingDaysConfig) GetMonday() bool`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *WorkingDaysConfig) GetMondayOk() (*bool, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *WorkingDaysConfig) SetMonday(v bool)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *WorkingDaysConfig) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetNonWorkingDays

`func (o *WorkingDaysConfig) GetNonWorkingDays() []NonWorkingDay`

GetNonWorkingDays returns the NonWorkingDays field if non-nil, zero value otherwise.

### GetNonWorkingDaysOk

`func (o *WorkingDaysConfig) GetNonWorkingDaysOk() (*[]NonWorkingDay, bool)`

GetNonWorkingDaysOk returns a tuple with the NonWorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonWorkingDays

`func (o *WorkingDaysConfig) SetNonWorkingDays(v []NonWorkingDay)`

SetNonWorkingDays sets NonWorkingDays field to given value.

### HasNonWorkingDays

`func (o *WorkingDaysConfig) HasNonWorkingDays() bool`

HasNonWorkingDays returns a boolean if a field has been set.

### GetSaturday

`func (o *WorkingDaysConfig) GetSaturday() bool`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *WorkingDaysConfig) GetSaturdayOk() (*bool, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *WorkingDaysConfig) SetSaturday(v bool)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *WorkingDaysConfig) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *WorkingDaysConfig) GetSunday() bool`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *WorkingDaysConfig) GetSundayOk() (*bool, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *WorkingDaysConfig) SetSunday(v bool)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *WorkingDaysConfig) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### GetThursday

`func (o *WorkingDaysConfig) GetThursday() bool`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *WorkingDaysConfig) GetThursdayOk() (*bool, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *WorkingDaysConfig) SetThursday(v bool)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *WorkingDaysConfig) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetTimezoneId

`func (o *WorkingDaysConfig) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *WorkingDaysConfig) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *WorkingDaysConfig) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *WorkingDaysConfig) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetTuesday

`func (o *WorkingDaysConfig) GetTuesday() bool`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *WorkingDaysConfig) GetTuesdayOk() (*bool, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *WorkingDaysConfig) SetTuesday(v bool)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *WorkingDaysConfig) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *WorkingDaysConfig) GetWednesday() bool`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *WorkingDaysConfig) GetWednesdayOk() (*bool, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *WorkingDaysConfig) SetWednesday(v bool)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *WorkingDaysConfig) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


