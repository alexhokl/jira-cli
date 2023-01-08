# StatusProjectUsagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Page token for the next page of issue type usages. | [optional] 
**Values** | Pointer to [**[]StatusProjectUsage**](StatusProjectUsage.md) | The list of projects. | [optional] 

## Methods

### NewStatusProjectUsagePage

`func NewStatusProjectUsagePage() *StatusProjectUsagePage`

NewStatusProjectUsagePage instantiates a new StatusProjectUsagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusProjectUsagePageWithDefaults

`func NewStatusProjectUsagePageWithDefaults() *StatusProjectUsagePage`

NewStatusProjectUsagePageWithDefaults instantiates a new StatusProjectUsagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *StatusProjectUsagePage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *StatusProjectUsagePage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *StatusProjectUsagePage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *StatusProjectUsagePage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetValues

`func (o *StatusProjectUsagePage) GetValues() []StatusProjectUsage`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StatusProjectUsagePage) GetValuesOk() (*[]StatusProjectUsage, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StatusProjectUsagePage) SetValues(v []StatusProjectUsage)`

SetValues sets Values field to given value.

### HasValues

`func (o *StatusProjectUsagePage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


