# MultiIssueEntityProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | Pointer to [**[]IssueEntityPropertiesForMultiUpdate**](IssueEntityPropertiesForMultiUpdate.md) | A list of issue IDs and their respective properties. | [optional] 

## Methods

### NewMultiIssueEntityProperties

`func NewMultiIssueEntityProperties() *MultiIssueEntityProperties`

NewMultiIssueEntityProperties instantiates a new MultiIssueEntityProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiIssueEntityPropertiesWithDefaults

`func NewMultiIssueEntityPropertiesWithDefaults() *MultiIssueEntityProperties`

NewMultiIssueEntityPropertiesWithDefaults instantiates a new MultiIssueEntityProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *MultiIssueEntityProperties) GetIssues() []IssueEntityPropertiesForMultiUpdate`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *MultiIssueEntityProperties) GetIssuesOk() (*[]IssueEntityPropertiesForMultiUpdate, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *MultiIssueEntityProperties) SetIssues(v []IssueEntityPropertiesForMultiUpdate)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *MultiIssueEntityProperties) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


