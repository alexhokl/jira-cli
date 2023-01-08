# DeleteAndReplaceVersionBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldReplacementList** | Pointer to [**[]CustomFieldReplacement**](CustomFieldReplacement.md) | An array of custom field IDs (&#x60;customFieldId&#x60;) and version IDs (&#x60;moveTo&#x60;) to update when the fields contain the deleted version. | [optional] 
**MoveAffectedIssuesTo** | Pointer to **int64** | The ID of the version to update &#x60;affectedVersion&#x60; to when the field contains the deleted version. | [optional] 
**MoveFixIssuesTo** | Pointer to **int64** | The ID of the version to update &#x60;fixVersion&#x60; to when the field contains the deleted version. | [optional] 

## Methods

### NewDeleteAndReplaceVersionBean

`func NewDeleteAndReplaceVersionBean() *DeleteAndReplaceVersionBean`

NewDeleteAndReplaceVersionBean instantiates a new DeleteAndReplaceVersionBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAndReplaceVersionBeanWithDefaults

`func NewDeleteAndReplaceVersionBeanWithDefaults() *DeleteAndReplaceVersionBean`

NewDeleteAndReplaceVersionBeanWithDefaults instantiates a new DeleteAndReplaceVersionBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldReplacementList

`func (o *DeleteAndReplaceVersionBean) GetCustomFieldReplacementList() []CustomFieldReplacement`

GetCustomFieldReplacementList returns the CustomFieldReplacementList field if non-nil, zero value otherwise.

### GetCustomFieldReplacementListOk

`func (o *DeleteAndReplaceVersionBean) GetCustomFieldReplacementListOk() (*[]CustomFieldReplacement, bool)`

GetCustomFieldReplacementListOk returns a tuple with the CustomFieldReplacementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldReplacementList

`func (o *DeleteAndReplaceVersionBean) SetCustomFieldReplacementList(v []CustomFieldReplacement)`

SetCustomFieldReplacementList sets CustomFieldReplacementList field to given value.

### HasCustomFieldReplacementList

`func (o *DeleteAndReplaceVersionBean) HasCustomFieldReplacementList() bool`

HasCustomFieldReplacementList returns a boolean if a field has been set.

### GetMoveAffectedIssuesTo

`func (o *DeleteAndReplaceVersionBean) GetMoveAffectedIssuesTo() int64`

GetMoveAffectedIssuesTo returns the MoveAffectedIssuesTo field if non-nil, zero value otherwise.

### GetMoveAffectedIssuesToOk

`func (o *DeleteAndReplaceVersionBean) GetMoveAffectedIssuesToOk() (*int64, bool)`

GetMoveAffectedIssuesToOk returns a tuple with the MoveAffectedIssuesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAffectedIssuesTo

`func (o *DeleteAndReplaceVersionBean) SetMoveAffectedIssuesTo(v int64)`

SetMoveAffectedIssuesTo sets MoveAffectedIssuesTo field to given value.

### HasMoveAffectedIssuesTo

`func (o *DeleteAndReplaceVersionBean) HasMoveAffectedIssuesTo() bool`

HasMoveAffectedIssuesTo returns a boolean if a field has been set.

### GetMoveFixIssuesTo

`func (o *DeleteAndReplaceVersionBean) GetMoveFixIssuesTo() int64`

GetMoveFixIssuesTo returns the MoveFixIssuesTo field if non-nil, zero value otherwise.

### GetMoveFixIssuesToOk

`func (o *DeleteAndReplaceVersionBean) GetMoveFixIssuesToOk() (*int64, bool)`

GetMoveFixIssuesToOk returns a tuple with the MoveFixIssuesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFixIssuesTo

`func (o *DeleteAndReplaceVersionBean) SetMoveFixIssuesTo(v int64)`

SetMoveFixIssuesTo sets MoveFixIssuesTo field to given value.

### HasMoveFixIssuesTo

`func (o *DeleteAndReplaceVersionBean) HasMoveFixIssuesTo() bool`

HasMoveFixIssuesTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


