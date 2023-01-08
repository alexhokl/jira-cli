# BulkWorklogKeyRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**[]WorklogCompositeKey**](WorklogCompositeKey.md) | A list of issue and worklog ID pairs. | [optional] 

## Methods

### NewBulkWorklogKeyRequestBean

`func NewBulkWorklogKeyRequestBean() *BulkWorklogKeyRequestBean`

NewBulkWorklogKeyRequestBean instantiates a new BulkWorklogKeyRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWorklogKeyRequestBeanWithDefaults

`func NewBulkWorklogKeyRequestBeanWithDefaults() *BulkWorklogKeyRequestBean`

NewBulkWorklogKeyRequestBeanWithDefaults instantiates a new BulkWorklogKeyRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *BulkWorklogKeyRequestBean) GetRequests() []WorklogCompositeKey`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BulkWorklogKeyRequestBean) GetRequestsOk() (*[]WorklogCompositeKey, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BulkWorklogKeyRequestBean) SetRequests(v []WorklogCompositeKey)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BulkWorklogKeyRequestBean) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


