# GetPermissionHolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The permission holder type. This is \&quot;Group\&quot; or \&quot;AccountId\&quot;. | 
**Value** | **string** | The permission holder value. This is a group name if the type is \&quot;Group\&quot; or an account ID if the type is \&quot;AccountId\&quot;. | 

## Methods

### NewGetPermissionHolderResponse

`func NewGetPermissionHolderResponse(type_ string, value string, ) *GetPermissionHolderResponse`

NewGetPermissionHolderResponse instantiates a new GetPermissionHolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPermissionHolderResponseWithDefaults

`func NewGetPermissionHolderResponseWithDefaults() *GetPermissionHolderResponse`

NewGetPermissionHolderResponseWithDefaults instantiates a new GetPermissionHolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetPermissionHolderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPermissionHolderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPermissionHolderResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GetPermissionHolderResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetPermissionHolderResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetPermissionHolderResponse) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


