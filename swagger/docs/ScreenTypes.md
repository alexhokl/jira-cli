# ScreenTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **int64** | The ID of the create screen. | [optional] 
**Default** | **int64** | The ID of the default screen. Required when creating a screen scheme. | 
**Edit** | Pointer to **int64** | The ID of the edit screen. | [optional] 
**View** | Pointer to **int64** | The ID of the view screen. | [optional] 

## Methods

### NewScreenTypes

`func NewScreenTypes(default_ int64, ) *ScreenTypes`

NewScreenTypes instantiates a new ScreenTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenTypesWithDefaults

`func NewScreenTypesWithDefaults() *ScreenTypes`

NewScreenTypesWithDefaults instantiates a new ScreenTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *ScreenTypes) GetCreate() int64`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ScreenTypes) GetCreateOk() (*int64, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ScreenTypes) SetCreate(v int64)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ScreenTypes) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *ScreenTypes) GetDefault() int64`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ScreenTypes) GetDefaultOk() (*int64, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ScreenTypes) SetDefault(v int64)`

SetDefault sets Default field to given value.


### GetEdit

`func (o *ScreenTypes) GetEdit() int64`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *ScreenTypes) GetEditOk() (*int64, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *ScreenTypes) SetEdit(v int64)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *ScreenTypes) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetView

`func (o *ScreenTypes) GetView() int64`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ScreenTypes) GetViewOk() (*int64, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ScreenTypes) SetView(v int64)`

SetView sets View field to given value.

### HasView

`func (o *ScreenTypes) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


