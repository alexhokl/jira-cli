# ProjectUsagePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | Page token for the next page of project usages. | [optional] 
**Values** | Pointer to [**[]ProjectUsage**](ProjectUsage.md) | The list of projects. | [optional] 

## Methods

### NewProjectUsagePage

`func NewProjectUsagePage() *ProjectUsagePage`

NewProjectUsagePage instantiates a new ProjectUsagePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUsagePageWithDefaults

`func NewProjectUsagePageWithDefaults() *ProjectUsagePage`

NewProjectUsagePageWithDefaults instantiates a new ProjectUsagePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ProjectUsagePage) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ProjectUsagePage) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ProjectUsagePage) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ProjectUsagePage) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetValues

`func (o *ProjectUsagePage) GetValues() []ProjectUsage`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProjectUsagePage) GetValuesOk() (*[]ProjectUsage, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProjectUsagePage) SetValues(v []ProjectUsage)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProjectUsagePage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


