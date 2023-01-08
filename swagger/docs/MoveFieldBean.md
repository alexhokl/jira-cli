# MoveFieldBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The ID of the screen tab field after which to place the moved screen tab field. Required if &#x60;position&#x60; isn&#39;t provided. | [optional] 
**Position** | Pointer to **string** | The named position to which the screen tab field should be moved. Required if &#x60;after&#x60; isn&#39;t provided. | [optional] 

## Methods

### NewMoveFieldBean

`func NewMoveFieldBean() *MoveFieldBean`

NewMoveFieldBean instantiates a new MoveFieldBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveFieldBeanWithDefaults

`func NewMoveFieldBeanWithDefaults() *MoveFieldBean`

NewMoveFieldBeanWithDefaults instantiates a new MoveFieldBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *MoveFieldBean) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *MoveFieldBean) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *MoveFieldBean) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *MoveFieldBean) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetPosition

`func (o *MoveFieldBean) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MoveFieldBean) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MoveFieldBean) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MoveFieldBean) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


