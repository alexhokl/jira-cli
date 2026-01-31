# ChangeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The name of the field changed. | [optional] [readonly] 
**FieldId** | Pointer to **string** | The ID of the field changed. | [optional] [readonly] 
**Fieldtype** | Pointer to **string** | The type of the field changed. | [optional] [readonly] 
**From** | Pointer to **string** | The details of the original value. | [optional] [readonly] 
**FromString** | Pointer to **string** | The details of the original value as a string. | [optional] [readonly] 
**To** | Pointer to **string** | The details of the new value. | [optional] [readonly] 
**ToString** | Pointer to **string** | The details of the new value as a string. | [optional] [readonly] 

## Methods

### NewChangeDetails

`func NewChangeDetails() *ChangeDetails`

NewChangeDetails instantiates a new ChangeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeDetailsWithDefaults

`func NewChangeDetailsWithDefaults() *ChangeDetails`

NewChangeDetailsWithDefaults instantiates a new ChangeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ChangeDetails) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ChangeDetails) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ChangeDetails) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ChangeDetails) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFieldId

`func (o *ChangeDetails) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *ChangeDetails) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *ChangeDetails) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *ChangeDetails) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetFieldtype

`func (o *ChangeDetails) GetFieldtype() string`

GetFieldtype returns the Fieldtype field if non-nil, zero value otherwise.

### GetFieldtypeOk

`func (o *ChangeDetails) GetFieldtypeOk() (*string, bool)`

GetFieldtypeOk returns a tuple with the Fieldtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldtype

`func (o *ChangeDetails) SetFieldtype(v string)`

SetFieldtype sets Fieldtype field to given value.

### HasFieldtype

`func (o *ChangeDetails) HasFieldtype() bool`

HasFieldtype returns a boolean if a field has been set.

### GetFrom

`func (o *ChangeDetails) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChangeDetails) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChangeDetails) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ChangeDetails) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromString

`func (o *ChangeDetails) GetFromString() string`

GetFromString returns the FromString field if non-nil, zero value otherwise.

### GetFromStringOk

`func (o *ChangeDetails) GetFromStringOk() (*string, bool)`

GetFromStringOk returns a tuple with the FromString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromString

`func (o *ChangeDetails) SetFromString(v string)`

SetFromString sets FromString field to given value.

### HasFromString

`func (o *ChangeDetails) HasFromString() bool`

HasFromString returns a boolean if a field has been set.

### GetTo

`func (o *ChangeDetails) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ChangeDetails) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ChangeDetails) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ChangeDetails) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToString

`func (o *ChangeDetails) GetToString() string`

GetToString returns the ToString field if non-nil, zero value otherwise.

### GetToStringOk

`func (o *ChangeDetails) GetToStringOk() (*string, bool)`

GetToStringOk returns a tuple with the ToString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToString

`func (o *ChangeDetails) SetToString(v string)`

SetToString sets ToString field to given value.

### HasToString

`func (o *ChangeDetails) HasToString() bool`

HasToString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


