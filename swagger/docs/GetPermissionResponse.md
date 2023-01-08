# GetPermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | [**GetPermissionHolderResponse**](GetPermissionHolderResponse.md) | The permission holder. | 
**Type** | **string** | The permission type. This is \&quot;View\&quot; or \&quot;Edit\&quot;. | 

## Methods

### NewGetPermissionResponse

`func NewGetPermissionResponse(holder GetPermissionHolderResponse, type_ string, ) *GetPermissionResponse`

NewGetPermissionResponse instantiates a new GetPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPermissionResponseWithDefaults

`func NewGetPermissionResponseWithDefaults() *GetPermissionResponse`

NewGetPermissionResponseWithDefaults instantiates a new GetPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *GetPermissionResponse) GetHolder() GetPermissionHolderResponse`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *GetPermissionResponse) GetHolderOk() (*GetPermissionHolderResponse, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *GetPermissionResponse) SetHolder(v GetPermissionHolderResponse)`

SetHolder sets Holder field to given value.


### GetType

`func (o *GetPermissionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPermissionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPermissionResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


