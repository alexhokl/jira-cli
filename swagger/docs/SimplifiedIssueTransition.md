# SimplifiedIssueTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to [**IssueTransitionStatus**](IssueTransitionStatus.md) | The issue status change of the transition. | [optional] [readonly] 
**TransitionId** | Pointer to **int32** | The unique ID of the transition. | [optional] [readonly] 
**TransitionName** | Pointer to **string** | The name of the transition. | [optional] [readonly] 

## Methods

### NewSimplifiedIssueTransition

`func NewSimplifiedIssueTransition() *SimplifiedIssueTransition`

NewSimplifiedIssueTransition instantiates a new SimplifiedIssueTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedIssueTransitionWithDefaults

`func NewSimplifiedIssueTransitionWithDefaults() *SimplifiedIssueTransition`

NewSimplifiedIssueTransitionWithDefaults instantiates a new SimplifiedIssueTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SimplifiedIssueTransition) GetTo() IssueTransitionStatus`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SimplifiedIssueTransition) GetToOk() (*IssueTransitionStatus, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SimplifiedIssueTransition) SetTo(v IssueTransitionStatus)`

SetTo sets To field to given value.

### HasTo

`func (o *SimplifiedIssueTransition) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransitionId

`func (o *SimplifiedIssueTransition) GetTransitionId() int32`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *SimplifiedIssueTransition) GetTransitionIdOk() (*int32, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *SimplifiedIssueTransition) SetTransitionId(v int32)`

SetTransitionId sets TransitionId field to given value.

### HasTransitionId

`func (o *SimplifiedIssueTransition) HasTransitionId() bool`

HasTransitionId returns a boolean if a field has been set.

### GetTransitionName

`func (o *SimplifiedIssueTransition) GetTransitionName() string`

GetTransitionName returns the TransitionName field if non-nil, zero value otherwise.

### GetTransitionNameOk

`func (o *SimplifiedIssueTransition) GetTransitionNameOk() (*string, bool)`

GetTransitionNameOk returns a tuple with the TransitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionName

`func (o *SimplifiedIssueTransition) SetTransitionName(v string)`

SetTransitionName sets TransitionName field to given value.

### HasTransitionName

`func (o *SimplifiedIssueTransition) HasTransitionName() bool`

HasTransitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


