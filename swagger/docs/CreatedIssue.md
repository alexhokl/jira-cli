# CreatedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the created issue or subtask. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the created issue or subtask. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the created issue or subtask. | [optional] [readonly] 
**Transition** | Pointer to [**NestedResponse**](NestedResponse.md) | The response code and messages related to any requested transition. | [optional] [readonly] 
**Watchers** | Pointer to [**NestedResponse**](NestedResponse.md) | The response code and messages related to any requested watchers. | [optional] [readonly] 

## Methods

### NewCreatedIssue

`func NewCreatedIssue() *CreatedIssue`

NewCreatedIssue instantiates a new CreatedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedIssueWithDefaults

`func NewCreatedIssueWithDefaults() *CreatedIssue`

NewCreatedIssueWithDefaults instantiates a new CreatedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedIssue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreatedIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CreatedIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreatedIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreatedIssue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreatedIssue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSelf

`func (o *CreatedIssue) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CreatedIssue) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CreatedIssue) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CreatedIssue) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTransition

`func (o *CreatedIssue) GetTransition() NestedResponse`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *CreatedIssue) GetTransitionOk() (*NestedResponse, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *CreatedIssue) SetTransition(v NestedResponse)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *CreatedIssue) HasTransition() bool`

HasTransition returns a boolean if a field has been set.

### GetWatchers

`func (o *CreatedIssue) GetWatchers() NestedResponse`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *CreatedIssue) GetWatchersOk() (*NestedResponse, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *CreatedIssue) SetWatchers(v NestedResponse)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *CreatedIssue) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


