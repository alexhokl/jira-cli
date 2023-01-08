# ScreenableField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the screen tab field. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the screen tab field. Required on create and update. The maximum length is 255 characters. | [optional] 

## Methods

### NewScreenableField

`func NewScreenableField() *ScreenableField`

NewScreenableField instantiates a new ScreenableField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenableFieldWithDefaults

`func NewScreenableFieldWithDefaults() *ScreenableField`

NewScreenableFieldWithDefaults instantiates a new ScreenableField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScreenableField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScreenableField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScreenableField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScreenableField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScreenableField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenableField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenableField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScreenableField) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


