# IssueArchivalSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**Errors**](Errors.md) |  | [optional] 
**NumberOfIssuesUpdated** | Pointer to **int64** |  | [optional] 

## Methods

### NewIssueArchivalSyncResponse

`func NewIssueArchivalSyncResponse() *IssueArchivalSyncResponse`

NewIssueArchivalSyncResponse instantiates a new IssueArchivalSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueArchivalSyncResponseWithDefaults

`func NewIssueArchivalSyncResponseWithDefaults() *IssueArchivalSyncResponse`

NewIssueArchivalSyncResponseWithDefaults instantiates a new IssueArchivalSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *IssueArchivalSyncResponse) GetErrors() Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IssueArchivalSyncResponse) GetErrorsOk() (*Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IssueArchivalSyncResponse) SetErrors(v Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IssueArchivalSyncResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetNumberOfIssuesUpdated

`func (o *IssueArchivalSyncResponse) GetNumberOfIssuesUpdated() int64`

GetNumberOfIssuesUpdated returns the NumberOfIssuesUpdated field if non-nil, zero value otherwise.

### GetNumberOfIssuesUpdatedOk

`func (o *IssueArchivalSyncResponse) GetNumberOfIssuesUpdatedOk() (*int64, bool)`

GetNumberOfIssuesUpdatedOk returns a tuple with the NumberOfIssuesUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfIssuesUpdated

`func (o *IssueArchivalSyncResponse) SetNumberOfIssuesUpdated(v int64)`

SetNumberOfIssuesUpdated sets NumberOfIssuesUpdated field to given value.

### HasNumberOfIssuesUpdated

`func (o *IssueArchivalSyncResponse) HasNumberOfIssuesUpdated() bool`

HasNumberOfIssuesUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


