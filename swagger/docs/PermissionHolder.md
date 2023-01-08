# PermissionHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional permission holder details in the response. | [optional] [readonly] 
**Parameter** | Pointer to **string** | As a group&#39;s name can change, use of &#x60;value&#x60; is recommended. The identifier associated withthe &#x60;type&#x60; value that defines the holder of the permission. | [optional] 
**Type** | **string** | The type of permission holder. | 
**Value** | Pointer to **string** | The identifier associated with the &#x60;type&#x60; value that defines the holder of the permission. | [optional] 

## Methods

### NewPermissionHolder

`func NewPermissionHolder(type_ string, ) *PermissionHolder`

NewPermissionHolder instantiates a new PermissionHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionHolderWithDefaults

`func NewPermissionHolderWithDefaults() *PermissionHolder`

NewPermissionHolderWithDefaults instantiates a new PermissionHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *PermissionHolder) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *PermissionHolder) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *PermissionHolder) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *PermissionHolder) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetParameter

`func (o *PermissionHolder) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *PermissionHolder) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *PermissionHolder) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *PermissionHolder) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetType

`func (o *PermissionHolder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PermissionHolder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PermissionHolder) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *PermissionHolder) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PermissionHolder) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PermissionHolder) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PermissionHolder) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


