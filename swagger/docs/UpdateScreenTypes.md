# UpdateScreenTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **string** | The ID of the create screen. To remove the screen association, pass a null. | [optional] 
**Default** | Pointer to **string** | The ID of the default screen. When specified, must include a screen ID as a default screen is required. | [optional] 
**Edit** | Pointer to **string** | The ID of the edit screen. To remove the screen association, pass a null. | [optional] 
**View** | Pointer to **string** | The ID of the view screen. To remove the screen association, pass a null. | [optional] 

## Methods

### NewUpdateScreenTypes

`func NewUpdateScreenTypes() *UpdateScreenTypes`

NewUpdateScreenTypes instantiates a new UpdateScreenTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScreenTypesWithDefaults

`func NewUpdateScreenTypesWithDefaults() *UpdateScreenTypes`

NewUpdateScreenTypesWithDefaults instantiates a new UpdateScreenTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *UpdateScreenTypes) GetCreate() string`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UpdateScreenTypes) GetCreateOk() (*string, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UpdateScreenTypes) SetCreate(v string)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UpdateScreenTypes) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDefault

`func (o *UpdateScreenTypes) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UpdateScreenTypes) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UpdateScreenTypes) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UpdateScreenTypes) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEdit

`func (o *UpdateScreenTypes) GetEdit() string`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *UpdateScreenTypes) GetEditOk() (*string, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *UpdateScreenTypes) SetEdit(v string)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *UpdateScreenTypes) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetView

`func (o *UpdateScreenTypes) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *UpdateScreenTypes) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *UpdateScreenTypes) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *UpdateScreenTypes) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


