# StatusWorkflowUsagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Page token for the next page of issue type usages. | [optional] 
**Values** | Pointer to [**[]StatusWorkflowUsageWorkflow**](StatusWorkflowUsageWorkflow.md) | The list of statuses. | [optional] 

## Methods

### NewStatusWorkflowUsagePage

`func NewStatusWorkflowUsagePage() *StatusWorkflowUsagePage`

NewStatusWorkflowUsagePage instantiates a new StatusWorkflowUsagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWorkflowUsagePageWithDefaults

`func NewStatusWorkflowUsagePageWithDefaults() *StatusWorkflowUsagePage`

NewStatusWorkflowUsagePageWithDefaults instantiates a new StatusWorkflowUsagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *StatusWorkflowUsagePage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *StatusWorkflowUsagePage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *StatusWorkflowUsagePage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *StatusWorkflowUsagePage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetValues

`func (o *StatusWorkflowUsagePage) GetValues() []StatusWorkflowUsageWorkflow`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StatusWorkflowUsagePage) GetValuesOk() (*[]StatusWorkflowUsageWorkflow, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StatusWorkflowUsagePage) SetValues(v []StatusWorkflowUsageWorkflow)`

SetValues sets Values field to given value.

### HasValues

`func (o *StatusWorkflowUsagePage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


