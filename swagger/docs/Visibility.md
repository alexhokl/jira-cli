# Visibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableString** | The ID of the group or the name of the role that visibility of this item is restricted to. | [optional] 
**Type** | Pointer to **string** | Whether visibility of this item is restricted to a group or role. | [optional] 
**Value** | Pointer to **string** | The name of the group or role that visibility of this item is restricted to. Please note that the name of a group is mutable, to reliably identify a group use &#x60;identifier&#x60;. | [optional] 

## Methods

### NewVisibility

`func NewVisibility() *Visibility`

NewVisibility instantiates a new Visibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisibilityWithDefaults

`func NewVisibilityWithDefaults() *Visibility`

NewVisibilityWithDefaults instantiates a new Visibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *Visibility) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Visibility) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Visibility) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Visibility) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Visibility) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Visibility) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetType

`func (o *Visibility) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Visibility) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Visibility) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Visibility) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Visibility) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Visibility) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Visibility) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Visibility) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


