# EnvironmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment. | 
**Type** | Pointer to **string** | The &#39;type&#39; or &#39;category&#39; of environment this environment belongs to. | [optional] 

## Methods

### NewEnvironmentDetails

`func NewEnvironmentDetails(name string, ) *EnvironmentDetails`

NewEnvironmentDetails instantiates a new EnvironmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDetailsWithDefaults

`func NewEnvironmentDetailsWithDefaults() *EnvironmentDetails`

NewEnvironmentDetailsWithDefaults instantiates a new EnvironmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentDetails) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EnvironmentDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


