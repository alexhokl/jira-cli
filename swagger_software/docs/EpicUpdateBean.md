# EpicUpdateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to [**PartiallyUpdateEpicRequestColor**](PartiallyUpdateEpicRequestColor.md) |  | [optional] 
**Done** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewEpicUpdateBean

`func NewEpicUpdateBean() *EpicUpdateBean`

NewEpicUpdateBean instantiates a new EpicUpdateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicUpdateBeanWithDefaults

`func NewEpicUpdateBeanWithDefaults() *EpicUpdateBean`

NewEpicUpdateBeanWithDefaults instantiates a new EpicUpdateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *EpicUpdateBean) GetColor() PartiallyUpdateEpicRequestColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EpicUpdateBean) GetColorOk() (*PartiallyUpdateEpicRequestColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EpicUpdateBean) SetColor(v PartiallyUpdateEpicRequestColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *EpicUpdateBean) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDone

`func (o *EpicUpdateBean) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *EpicUpdateBean) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *EpicUpdateBean) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *EpicUpdateBean) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetName

`func (o *EpicUpdateBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpicUpdateBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpicUpdateBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EpicUpdateBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummary

`func (o *EpicUpdateBean) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EpicUpdateBean) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EpicUpdateBean) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EpicUpdateBean) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


