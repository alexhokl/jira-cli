# FieldLastUsed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Last used value type:   *  *TRACKED*: field is tracked and a last used date is available.  *  *NOT\\_TRACKED*: field is not tracked, last used date is not available.  *  *NO\\_INFORMATION*: field is tracked, but no last used date is available. | [optional] 
**Value** | Pointer to **time.Time** | The date when the value of the field last changed. | [optional] 

## Methods

### NewFieldLastUsed

`func NewFieldLastUsed() *FieldLastUsed`

NewFieldLastUsed instantiates a new FieldLastUsed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldLastUsedWithDefaults

`func NewFieldLastUsedWithDefaults() *FieldLastUsed`

NewFieldLastUsedWithDefaults instantiates a new FieldLastUsed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldLastUsed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldLastUsed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldLastUsed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldLastUsed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *FieldLastUsed) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldLastUsed) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldLastUsed) SetValue(v time.Time)`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldLastUsed) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


