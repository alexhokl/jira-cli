# DateRangeFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateAfter** | **string** | List issues archived after a specified date, passed in the YYYY-MM-DD format. | 
**DateBefore** | **string** | List issues archived before a specified date provided in the YYYY-MM-DD format. | 

## Methods

### NewDateRangeFilterRequest

`func NewDateRangeFilterRequest(dateAfter string, dateBefore string, ) *DateRangeFilterRequest`

NewDateRangeFilterRequest instantiates a new DateRangeFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeFilterRequestWithDefaults

`func NewDateRangeFilterRequestWithDefaults() *DateRangeFilterRequest`

NewDateRangeFilterRequestWithDefaults instantiates a new DateRangeFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateAfter

`func (o *DateRangeFilterRequest) GetDateAfter() string`

GetDateAfter returns the DateAfter field if non-nil, zero value otherwise.

### GetDateAfterOk

`func (o *DateRangeFilterRequest) GetDateAfterOk() (*string, bool)`

GetDateAfterOk returns a tuple with the DateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAfter

`func (o *DateRangeFilterRequest) SetDateAfter(v string)`

SetDateAfter sets DateAfter field to given value.


### GetDateBefore

`func (o *DateRangeFilterRequest) GetDateBefore() string`

GetDateBefore returns the DateBefore field if non-nil, zero value otherwise.

### GetDateBeforeOk

`func (o *DateRangeFilterRequest) GetDateBeforeOk() (*string, bool)`

GetDateBeforeOk returns a tuple with the DateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBefore

`func (o *DateRangeFilterRequest) SetDateBefore(v string)`

SetDateBefore sets DateBefore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


