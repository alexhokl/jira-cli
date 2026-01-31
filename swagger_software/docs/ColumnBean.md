# ColumnBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **int32** |  | [optional] 
**Min** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Statuses** | Pointer to [**[]GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner**](GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner.md) |  | [optional] 

## Methods

### NewColumnBean

`func NewColumnBean() *ColumnBean`

NewColumnBean instantiates a new ColumnBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnBeanWithDefaults

`func NewColumnBeanWithDefaults() *ColumnBean`

NewColumnBeanWithDefaults instantiates a new ColumnBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *ColumnBean) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ColumnBean) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ColumnBean) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ColumnBean) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *ColumnBean) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ColumnBean) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ColumnBean) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *ColumnBean) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *ColumnBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ColumnBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatuses

`func (o *ColumnBean) GetStatuses() []GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ColumnBean) GetStatusesOk() (*[]GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ColumnBean) SetStatuses(v []GetConfiguration200ResponseColumnConfigColumnsInnerStatusesInner)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ColumnBean) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


