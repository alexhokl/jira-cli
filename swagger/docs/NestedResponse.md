# NestedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCollection** | Pointer to [**ErrorCollection**](ErrorCollection.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**WarningCollection** | Pointer to [**WarningCollection**](WarningCollection.md) |  | [optional] 

## Methods

### NewNestedResponse

`func NewNestedResponse() *NestedResponse`

NewNestedResponse instantiates a new NestedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedResponseWithDefaults

`func NewNestedResponseWithDefaults() *NestedResponse`

NewNestedResponseWithDefaults instantiates a new NestedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCollection

`func (o *NestedResponse) GetErrorCollection() ErrorCollection`

GetErrorCollection returns the ErrorCollection field if non-nil, zero value otherwise.

### GetErrorCollectionOk

`func (o *NestedResponse) GetErrorCollectionOk() (*ErrorCollection, bool)`

GetErrorCollectionOk returns a tuple with the ErrorCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCollection

`func (o *NestedResponse) SetErrorCollection(v ErrorCollection)`

SetErrorCollection sets ErrorCollection field to given value.

### HasErrorCollection

`func (o *NestedResponse) HasErrorCollection() bool`

HasErrorCollection returns a boolean if a field has been set.

### GetStatus

`func (o *NestedResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NestedResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NestedResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NestedResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarningCollection

`func (o *NestedResponse) GetWarningCollection() WarningCollection`

GetWarningCollection returns the WarningCollection field if non-nil, zero value otherwise.

### GetWarningCollectionOk

`func (o *NestedResponse) GetWarningCollectionOk() (*WarningCollection, bool)`

GetWarningCollectionOk returns a tuple with the WarningCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCollection

`func (o *NestedResponse) SetWarningCollection(v WarningCollection)`

SetWarningCollection sets WarningCollection field to given value.

### HasWarningCollection

`func (o *NestedResponse) HasWarningCollection() bool`

HasWarningCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


