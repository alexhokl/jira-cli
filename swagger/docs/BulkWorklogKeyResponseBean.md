# BulkWorklogKeyResponseBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Worklogs** | Pointer to [**[]WorklogKeyResult**](WorklogKeyResult.md) | A list of successfully retrieved worklogs with their issue and worklog IDs. | [optional] 

## Methods

### NewBulkWorklogKeyResponseBean

`func NewBulkWorklogKeyResponseBean() *BulkWorklogKeyResponseBean`

NewBulkWorklogKeyResponseBean instantiates a new BulkWorklogKeyResponseBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWorklogKeyResponseBeanWithDefaults

`func NewBulkWorklogKeyResponseBeanWithDefaults() *BulkWorklogKeyResponseBean`

NewBulkWorklogKeyResponseBeanWithDefaults instantiates a new BulkWorklogKeyResponseBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorklogs

`func (o *BulkWorklogKeyResponseBean) GetWorklogs() []WorklogKeyResult`

GetWorklogs returns the Worklogs field if non-nil, zero value otherwise.

### GetWorklogsOk

`func (o *BulkWorklogKeyResponseBean) GetWorklogsOk() (*[]WorklogKeyResult, bool)`

GetWorklogsOk returns a tuple with the Worklogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorklogs

`func (o *BulkWorklogKeyResponseBean) SetWorklogs(v []WorklogKeyResult)`

SetWorklogs sets Worklogs field to given value.

### HasWorklogs

`func (o *BulkWorklogKeyResponseBean) HasWorklogs() bool`

HasWorklogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


