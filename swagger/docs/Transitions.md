# Transitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional transitions details in the response. | [optional] [readonly] 
**Transitions** | Pointer to [**[]IssueTransition**](IssueTransition.md) | List of issue transitions. | [optional] [readonly] 

## Methods

### NewTransitions

`func NewTransitions() *Transitions`

NewTransitions instantiates a new Transitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionsWithDefaults

`func NewTransitionsWithDefaults() *Transitions`

NewTransitionsWithDefaults instantiates a new Transitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *Transitions) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *Transitions) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *Transitions) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *Transitions) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetTransitions

`func (o *Transitions) GetTransitions() []IssueTransition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *Transitions) GetTransitionsOk() (*[]IssueTransition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *Transitions) SetTransitions(v []IssueTransition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *Transitions) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


