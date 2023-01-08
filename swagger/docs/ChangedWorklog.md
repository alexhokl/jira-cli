# ChangedWorklog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**[]EntityProperty**](EntityProperty.md) | Details of properties associated with the change. | [optional] [readonly] 
**UpdatedTime** | Pointer to **int64** | The datetime of the change. | [optional] [readonly] 
**WorklogId** | Pointer to **int64** | The ID of the worklog. | [optional] [readonly] 

## Methods

### NewChangedWorklog

`func NewChangedWorklog() *ChangedWorklog`

NewChangedWorklog instantiates a new ChangedWorklog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangedWorklogWithDefaults

`func NewChangedWorklogWithDefaults() *ChangedWorklog`

NewChangedWorklogWithDefaults instantiates a new ChangedWorklog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ChangedWorklog) GetProperties() []EntityProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ChangedWorklog) GetPropertiesOk() (*[]EntityProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ChangedWorklog) SetProperties(v []EntityProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ChangedWorklog) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ChangedWorklog) GetUpdatedTime() int64`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ChangedWorklog) GetUpdatedTimeOk() (*int64, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ChangedWorklog) SetUpdatedTime(v int64)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ChangedWorklog) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetWorklogId

`func (o *ChangedWorklog) GetWorklogId() int64`

GetWorklogId returns the WorklogId field if non-nil, zero value otherwise.

### GetWorklogIdOk

`func (o *ChangedWorklog) GetWorklogIdOk() (*int64, bool)`

GetWorklogIdOk returns a tuple with the WorklogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorklogId

`func (o *ChangedWorklog) SetWorklogId(v int64)`

SetWorklogId sets WorklogId field to given value.

### HasWorklogId

`func (o *ChangedWorklog) HasWorklogId() bool`

HasWorklogId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


