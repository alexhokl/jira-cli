# TimeTrackingProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key for the time tracking provider. For example, *JIRA*. | 
**Name** | Pointer to **string** | The name of the time tracking provider. For example, *JIRA provided time tracking*. | [optional] 
**Url** | Pointer to **string** | The URL of the configuration page for the time tracking provider app. For example, *_/example/config/url*. This property is only returned if the &#x60;adminPageKey&#x60; property is set in the module descriptor of the time tracking provider app. | [optional] [readonly] 

## Methods

### NewTimeTrackingProvider

`func NewTimeTrackingProvider(key string, ) *TimeTrackingProvider`

NewTimeTrackingProvider instantiates a new TimeTrackingProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeTrackingProviderWithDefaults

`func NewTimeTrackingProviderWithDefaults() *TimeTrackingProvider`

NewTimeTrackingProviderWithDefaults instantiates a new TimeTrackingProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TimeTrackingProvider) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TimeTrackingProvider) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TimeTrackingProvider) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *TimeTrackingProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeTrackingProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeTrackingProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeTrackingProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *TimeTrackingProvider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TimeTrackingProvider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TimeTrackingProvider) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TimeTrackingProvider) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


