# ChangedWorklogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPage** | Pointer to **bool** |  | [optional] 
**NextPage** | Pointer to **string** | The URL of the next list of changed worklogs. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of this changed worklogs list. | [optional] [readonly] 
**Since** | Pointer to **int64** | The datetime of the first worklog item in the list. | [optional] [readonly] 
**Until** | Pointer to **int64** | The datetime of the last worklog item in the list. | [optional] [readonly] 
**Values** | Pointer to [**[]ChangedWorklog**](ChangedWorklog.md) | Changed worklog list. | [optional] [readonly] 

## Methods

### NewChangedWorklogs

`func NewChangedWorklogs() *ChangedWorklogs`

NewChangedWorklogs instantiates a new ChangedWorklogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangedWorklogsWithDefaults

`func NewChangedWorklogsWithDefaults() *ChangedWorklogs`

NewChangedWorklogsWithDefaults instantiates a new ChangedWorklogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPage

`func (o *ChangedWorklogs) GetLastPage() bool`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ChangedWorklogs) GetLastPageOk() (*bool, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ChangedWorklogs) SetLastPage(v bool)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *ChangedWorklogs) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetNextPage

`func (o *ChangedWorklogs) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ChangedWorklogs) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ChangedWorklogs) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ChangedWorklogs) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *ChangedWorklogs) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ChangedWorklogs) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ChangedWorklogs) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ChangedWorklogs) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSince

`func (o *ChangedWorklogs) GetSince() int64`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *ChangedWorklogs) GetSinceOk() (*int64, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *ChangedWorklogs) SetSince(v int64)`

SetSince sets Since field to given value.

### HasSince

`func (o *ChangedWorklogs) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetUntil

`func (o *ChangedWorklogs) GetUntil() int64`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *ChangedWorklogs) GetUntilOk() (*int64, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *ChangedWorklogs) SetUntil(v int64)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *ChangedWorklogs) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetValues

`func (o *ChangedWorklogs) GetValues() []ChangedWorklog`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ChangedWorklogs) GetValuesOk() (*[]ChangedWorklog, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ChangedWorklogs) SetValues(v []ChangedWorklog)`

SetValues sets Values field to given value.

### HasValues

`func (o *ChangedWorklogs) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


