# CreatePermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | [**CreatePermissionHolderRequest**](CreatePermissionHolderRequest.md) | The permission holder. | 
**Type** | **string** | The permission type. This must be \&quot;View\&quot; or \&quot;Edit\&quot;. | 

## Methods

### NewCreatePermissionRequest

`func NewCreatePermissionRequest(holder CreatePermissionHolderRequest, type_ string, ) *CreatePermissionRequest`

NewCreatePermissionRequest instantiates a new CreatePermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermissionRequestWithDefaults

`func NewCreatePermissionRequestWithDefaults() *CreatePermissionRequest`

NewCreatePermissionRequestWithDefaults instantiates a new CreatePermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *CreatePermissionRequest) GetHolder() CreatePermissionHolderRequest`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *CreatePermissionRequest) GetHolderOk() (*CreatePermissionHolderRequest, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *CreatePermissionRequest) SetHolder(v CreatePermissionHolderRequest)`

SetHolder sets Holder field to given value.


### GetType

`func (o *CreatePermissionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePermissionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePermissionRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


