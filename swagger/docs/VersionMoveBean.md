# VersionMoveBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The URL (self link) of the version after which to place the moved version. Cannot be used with &#x60;position&#x60;. | [optional] 
**Position** | Pointer to **string** | An absolute position in which to place the moved version. Cannot be used with &#x60;after&#x60;. | [optional] 

## Methods

### NewVersionMoveBean

`func NewVersionMoveBean() *VersionMoveBean`

NewVersionMoveBean instantiates a new VersionMoveBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionMoveBeanWithDefaults

`func NewVersionMoveBeanWithDefaults() *VersionMoveBean`

NewVersionMoveBeanWithDefaults instantiates a new VersionMoveBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *VersionMoveBean) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *VersionMoveBean) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *VersionMoveBean) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *VersionMoveBean) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetPosition

`func (o *VersionMoveBean) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VersionMoveBean) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VersionMoveBean) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VersionMoveBean) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


