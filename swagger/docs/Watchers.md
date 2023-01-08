# Watchers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsWatching** | Pointer to **bool** | Whether the calling user is watching this issue. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of these issue watcher details. | [optional] [readonly] 
**WatchCount** | Pointer to **int32** | The number of users watching this issue. | [optional] [readonly] 
**Watchers** | Pointer to [**[]UserDetails**](UserDetails.md) | Details of the users watching this issue. | [optional] [readonly] 

## Methods

### NewWatchers

`func NewWatchers() *Watchers`

NewWatchers instantiates a new Watchers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchersWithDefaults

`func NewWatchersWithDefaults() *Watchers`

NewWatchersWithDefaults instantiates a new Watchers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsWatching

`func (o *Watchers) GetIsWatching() bool`

GetIsWatching returns the IsWatching field if non-nil, zero value otherwise.

### GetIsWatchingOk

`func (o *Watchers) GetIsWatchingOk() (*bool, bool)`

GetIsWatchingOk returns a tuple with the IsWatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatching

`func (o *Watchers) SetIsWatching(v bool)`

SetIsWatching sets IsWatching field to given value.

### HasIsWatching

`func (o *Watchers) HasIsWatching() bool`

HasIsWatching returns a boolean if a field has been set.

### GetSelf

`func (o *Watchers) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Watchers) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Watchers) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Watchers) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetWatchCount

`func (o *Watchers) GetWatchCount() int32`

GetWatchCount returns the WatchCount field if non-nil, zero value otherwise.

### GetWatchCountOk

`func (o *Watchers) GetWatchCountOk() (*int32, bool)`

GetWatchCountOk returns a tuple with the WatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchCount

`func (o *Watchers) SetWatchCount(v int32)`

SetWatchCount sets WatchCount field to given value.

### HasWatchCount

`func (o *Watchers) HasWatchCount() bool`

HasWatchCount returns a boolean if a field has been set.

### GetWatchers

`func (o *Watchers) GetWatchers() []UserDetails`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *Watchers) GetWatchersOk() (*[]UserDetails, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *Watchers) SetWatchers(v []UserDetails)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *Watchers) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


