# RemoveOptionFromIssuesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**SimpleErrorCollection**](SimpleErrorCollection.md) | A collection of errors related to unchanged issues. The collection size is limited, which means not all errors may be returned. | [optional] 
**ModifiedIssues** | Pointer to **[]int64** | The IDs of the modified issues. | [optional] 
**UnmodifiedIssues** | Pointer to **[]int64** | The IDs of the unchanged issues, those issues where errors prevent modification. | [optional] 

## Methods

### NewRemoveOptionFromIssuesResult

`func NewRemoveOptionFromIssuesResult() *RemoveOptionFromIssuesResult`

NewRemoveOptionFromIssuesResult instantiates a new RemoveOptionFromIssuesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveOptionFromIssuesResultWithDefaults

`func NewRemoveOptionFromIssuesResultWithDefaults() *RemoveOptionFromIssuesResult`

NewRemoveOptionFromIssuesResultWithDefaults instantiates a new RemoveOptionFromIssuesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *RemoveOptionFromIssuesResult) GetErrors() SimpleErrorCollection`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RemoveOptionFromIssuesResult) GetErrorsOk() (*SimpleErrorCollection, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RemoveOptionFromIssuesResult) SetErrors(v SimpleErrorCollection)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RemoveOptionFromIssuesResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetModifiedIssues

`func (o *RemoveOptionFromIssuesResult) GetModifiedIssues() []int64`

GetModifiedIssues returns the ModifiedIssues field if non-nil, zero value otherwise.

### GetModifiedIssuesOk

`func (o *RemoveOptionFromIssuesResult) GetModifiedIssuesOk() (*[]int64, bool)`

GetModifiedIssuesOk returns a tuple with the ModifiedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedIssues

`func (o *RemoveOptionFromIssuesResult) SetModifiedIssues(v []int64)`

SetModifiedIssues sets ModifiedIssues field to given value.

### HasModifiedIssues

`func (o *RemoveOptionFromIssuesResult) HasModifiedIssues() bool`

HasModifiedIssues returns a boolean if a field has been set.

### GetUnmodifiedIssues

`func (o *RemoveOptionFromIssuesResult) GetUnmodifiedIssues() []int64`

GetUnmodifiedIssues returns the UnmodifiedIssues field if non-nil, zero value otherwise.

### GetUnmodifiedIssuesOk

`func (o *RemoveOptionFromIssuesResult) GetUnmodifiedIssuesOk() (*[]int64, bool)`

GetUnmodifiedIssuesOk returns a tuple with the UnmodifiedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmodifiedIssues

`func (o *RemoveOptionFromIssuesResult) SetUnmodifiedIssues(v []int64)`

SetUnmodifiedIssues sets UnmodifiedIssues field to given value.

### HasUnmodifiedIssues

`func (o *RemoveOptionFromIssuesResult) HasUnmodifiedIssues() bool`

HasUnmodifiedIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


