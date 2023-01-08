# BulkTransitionSubmitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedIssueIdsOrKeys** | **[]string** | List of all the issue IDs or keys that are to be bulk transitioned. | 
**TransitionId** | **string** | The ID of the transition that is to be performed on the issues. | 

## Methods

### NewBulkTransitionSubmitInput

`func NewBulkTransitionSubmitInput(selectedIssueIdsOrKeys []string, transitionId string, ) *BulkTransitionSubmitInput`

NewBulkTransitionSubmitInput instantiates a new BulkTransitionSubmitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTransitionSubmitInputWithDefaults

`func NewBulkTransitionSubmitInputWithDefaults() *BulkTransitionSubmitInput`

NewBulkTransitionSubmitInputWithDefaults instantiates a new BulkTransitionSubmitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedIssueIdsOrKeys

`func (o *BulkTransitionSubmitInput) GetSelectedIssueIdsOrKeys() []string`

GetSelectedIssueIdsOrKeys returns the SelectedIssueIdsOrKeys field if non-nil, zero value otherwise.

### GetSelectedIssueIdsOrKeysOk

`func (o *BulkTransitionSubmitInput) GetSelectedIssueIdsOrKeysOk() (*[]string, bool)`

GetSelectedIssueIdsOrKeysOk returns a tuple with the SelectedIssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedIssueIdsOrKeys

`func (o *BulkTransitionSubmitInput) SetSelectedIssueIdsOrKeys(v []string)`

SetSelectedIssueIdsOrKeys sets SelectedIssueIdsOrKeys field to given value.


### GetTransitionId

`func (o *BulkTransitionSubmitInput) GetTransitionId() string`

GetTransitionId returns the TransitionId field if non-nil, zero value otherwise.

### GetTransitionIdOk

`func (o *BulkTransitionSubmitInput) GetTransitionIdOk() (*string, bool)`

GetTransitionIdOk returns a tuple with the TransitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionId

`func (o *BulkTransitionSubmitInput) SetTransitionId(v string)`

SetTransitionId sets TransitionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


