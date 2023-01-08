# BoardColumnPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumIssueConstraint** | Pointer to **int64** | The maximum issue constraint for the column | [optional] 
**MinimumIssueConstraint** | Pointer to **int64** | The minimum issue constraint for the column | [optional] 
**Name** | Pointer to **string** | The name of the column | [optional] 
**StatusIds** | Pointer to [**[]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | The status IDs for the column | [optional] 

## Methods

### NewBoardColumnPayload

`func NewBoardColumnPayload() *BoardColumnPayload`

NewBoardColumnPayload instantiates a new BoardColumnPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardColumnPayloadWithDefaults

`func NewBoardColumnPayloadWithDefaults() *BoardColumnPayload`

NewBoardColumnPayloadWithDefaults instantiates a new BoardColumnPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumIssueConstraint

`func (o *BoardColumnPayload) GetMaximumIssueConstraint() int64`

GetMaximumIssueConstraint returns the MaximumIssueConstraint field if non-nil, zero value otherwise.

### GetMaximumIssueConstraintOk

`func (o *BoardColumnPayload) GetMaximumIssueConstraintOk() (*int64, bool)`

GetMaximumIssueConstraintOk returns a tuple with the MaximumIssueConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumIssueConstraint

`func (o *BoardColumnPayload) SetMaximumIssueConstraint(v int64)`

SetMaximumIssueConstraint sets MaximumIssueConstraint field to given value.

### HasMaximumIssueConstraint

`func (o *BoardColumnPayload) HasMaximumIssueConstraint() bool`

HasMaximumIssueConstraint returns a boolean if a field has been set.

### GetMinimumIssueConstraint

`func (o *BoardColumnPayload) GetMinimumIssueConstraint() int64`

GetMinimumIssueConstraint returns the MinimumIssueConstraint field if non-nil, zero value otherwise.

### GetMinimumIssueConstraintOk

`func (o *BoardColumnPayload) GetMinimumIssueConstraintOk() (*int64, bool)`

GetMinimumIssueConstraintOk returns a tuple with the MinimumIssueConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumIssueConstraint

`func (o *BoardColumnPayload) SetMinimumIssueConstraint(v int64)`

SetMinimumIssueConstraint sets MinimumIssueConstraint field to given value.

### HasMinimumIssueConstraint

`func (o *BoardColumnPayload) HasMinimumIssueConstraint() bool`

HasMinimumIssueConstraint returns a boolean if a field has been set.

### GetName

`func (o *BoardColumnPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardColumnPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardColumnPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardColumnPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusIds

`func (o *BoardColumnPayload) GetStatusIds() []ProjectCreateResourceIdentifier`

GetStatusIds returns the StatusIds field if non-nil, zero value otherwise.

### GetStatusIdsOk

`func (o *BoardColumnPayload) GetStatusIdsOk() (*[]ProjectCreateResourceIdentifier, bool)`

GetStatusIdsOk returns a tuple with the StatusIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIds

`func (o *BoardColumnPayload) SetStatusIds(v []ProjectCreateResourceIdentifier)`

SetStatusIds sets StatusIds field to given value.

### HasStatusIds

`func (o *BoardColumnPayload) HasStatusIds() bool`

HasStatusIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


