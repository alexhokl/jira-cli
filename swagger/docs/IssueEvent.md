# IssueEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the event. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the event. | [optional] [readonly] 

## Methods

### NewIssueEvent

`func NewIssueEvent() *IssueEvent`

NewIssueEvent instantiates a new IssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEventWithDefaults

`func NewIssueEventWithDefaults() *IssueEvent`

NewIssueEventWithDefaults instantiates a new IssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueEvent) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


