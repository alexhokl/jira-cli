# ApprovalConfigurationPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **string** | The active approval configuration. | [optional] 
**TransitionApproved** | Pointer to **string** | The transition ID for approved state. | [optional] 
**TransitionRejected** | Pointer to **string** | The transition ID for rejected state. | [optional] 

## Methods

### NewApprovalConfigurationPreview

`func NewApprovalConfigurationPreview() *ApprovalConfigurationPreview`

NewApprovalConfigurationPreview instantiates a new ApprovalConfigurationPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalConfigurationPreviewWithDefaults

`func NewApprovalConfigurationPreviewWithDefaults() *ApprovalConfigurationPreview`

NewApprovalConfigurationPreviewWithDefaults instantiates a new ApprovalConfigurationPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApprovalConfigurationPreview) GetActive() string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApprovalConfigurationPreview) GetActiveOk() (*string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApprovalConfigurationPreview) SetActive(v string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApprovalConfigurationPreview) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTransitionApproved

`func (o *ApprovalConfigurationPreview) GetTransitionApproved() string`

GetTransitionApproved returns the TransitionApproved field if non-nil, zero value otherwise.

### GetTransitionApprovedOk

`func (o *ApprovalConfigurationPreview) GetTransitionApprovedOk() (*string, bool)`

GetTransitionApprovedOk returns a tuple with the TransitionApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionApproved

`func (o *ApprovalConfigurationPreview) SetTransitionApproved(v string)`

SetTransitionApproved sets TransitionApproved field to given value.

### HasTransitionApproved

`func (o *ApprovalConfigurationPreview) HasTransitionApproved() bool`

HasTransitionApproved returns a boolean if a field has been set.

### GetTransitionRejected

`func (o *ApprovalConfigurationPreview) GetTransitionRejected() string`

GetTransitionRejected returns the TransitionRejected field if non-nil, zero value otherwise.

### GetTransitionRejectedOk

`func (o *ApprovalConfigurationPreview) GetTransitionRejectedOk() (*string, bool)`

GetTransitionRejectedOk returns a tuple with the TransitionRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionRejected

`func (o *ApprovalConfigurationPreview) SetTransitionRejected(v string)`

SetTransitionRejected sets TransitionRejected field to given value.

### HasTransitionRejected

`func (o *ApprovalConfigurationPreview) HasTransitionRejected() bool`

HasTransitionRejected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


