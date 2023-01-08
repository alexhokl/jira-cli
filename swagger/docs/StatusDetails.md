# StatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the status. | [optional] [readonly] 
**IconUrl** | Pointer to **string** | The URL of the icon used to represent the status. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the status. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the status. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the field. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the status. | [optional] [readonly] 
**StatusCategory** | Pointer to [**StatusCategory**](StatusCategory.md) | The category assigned to the status. | [optional] [readonly] 

## Methods

### NewStatusDetails

`func NewStatusDetails() *StatusDetails`

NewStatusDetails instantiates a new StatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDetailsWithDefaults

`func NewStatusDetailsWithDefaults() *StatusDetails`

NewStatusDetailsWithDefaults instantiates a new StatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StatusDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *StatusDetails) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *StatusDetails) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *StatusDetails) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *StatusDetails) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *StatusDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatusDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatusDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StatusDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StatusDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatusDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *StatusDetails) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StatusDetails) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StatusDetails) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *StatusDetails) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSelf

`func (o *StatusDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *StatusDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *StatusDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *StatusDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStatusCategory

`func (o *StatusDetails) GetStatusCategory() StatusCategory`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *StatusDetails) GetStatusCategoryOk() (*StatusCategory, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *StatusDetails) SetStatusCategory(v StatusCategory)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *StatusDetails) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


