# WorkflowSchemeUsagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Token for the next page of issue type usages. | [optional] 
**Values** | Pointer to [**[]WorkflowSchemeUsage**](WorkflowSchemeUsage.md) | The list of workflow schemes. | [optional] 

## Methods

### NewWorkflowSchemeUsagePage

`func NewWorkflowSchemeUsagePage() *WorkflowSchemeUsagePage`

NewWorkflowSchemeUsagePage instantiates a new WorkflowSchemeUsagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeUsagePageWithDefaults

`func NewWorkflowSchemeUsagePageWithDefaults() *WorkflowSchemeUsagePage`

NewWorkflowSchemeUsagePageWithDefaults instantiates a new WorkflowSchemeUsagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *WorkflowSchemeUsagePage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WorkflowSchemeUsagePage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WorkflowSchemeUsagePage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *WorkflowSchemeUsagePage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetValues

`func (o *WorkflowSchemeUsagePage) GetValues() []WorkflowSchemeUsage`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkflowSchemeUsagePage) GetValuesOk() (*[]WorkflowSchemeUsage, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkflowSchemeUsagePage) SetValues(v []WorkflowSchemeUsage)`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkflowSchemeUsagePage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


