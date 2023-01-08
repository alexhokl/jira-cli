# DefaultShareScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The scope of the default sharing for new filters and dashboards:   *  &#x60;AUTHENTICATED&#x60; Shared with all logged-in users.  *  &#x60;GLOBAL&#x60; Shared with all logged-in users. This shows as &#x60;AUTHENTICATED&#x60; in the response.  *  &#x60;PRIVATE&#x60; Not shared with any users. | 

## Methods

### NewDefaultShareScope

`func NewDefaultShareScope(scope string, ) *DefaultShareScope`

NewDefaultShareScope instantiates a new DefaultShareScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultShareScopeWithDefaults

`func NewDefaultShareScopeWithDefaults() *DefaultShareScope`

NewDefaultShareScopeWithDefaults instantiates a new DefaultShareScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *DefaultShareScope) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DefaultShareScope) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DefaultShareScope) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


