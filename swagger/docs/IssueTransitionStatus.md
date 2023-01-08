# IssueTransitionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusId** | Pointer to **int32** | The unique ID of the status. | [optional] [readonly] 
**StatusName** | Pointer to **string** | The name of the status. | [optional] [readonly] 

## Methods

### NewIssueTransitionStatus

`func NewIssueTransitionStatus() *IssueTransitionStatus`

NewIssueTransitionStatus instantiates a new IssueTransitionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTransitionStatusWithDefaults

`func NewIssueTransitionStatusWithDefaults() *IssueTransitionStatus`

NewIssueTransitionStatusWithDefaults instantiates a new IssueTransitionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusId

`func (o *IssueTransitionStatus) GetStatusId() int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *IssueTransitionStatus) GetStatusIdOk() (*int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *IssueTransitionStatus) SetStatusId(v int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *IssueTransitionStatus) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStatusName

`func (o *IssueTransitionStatus) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *IssueTransitionStatus) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *IssueTransitionStatus) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *IssueTransitionStatus) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


