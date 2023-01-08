# ScreenWithTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the screen. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the screen. | [optional] [readonly] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the screen. | [optional] 
**Tab** | Pointer to [**ScreenableTab**](ScreenableTab.md) | The tab for the screen. | [optional] 

## Methods

### NewScreenWithTab

`func NewScreenWithTab() *ScreenWithTab`

NewScreenWithTab instantiates a new ScreenWithTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenWithTabWithDefaults

`func NewScreenWithTabWithDefaults() *ScreenWithTab`

NewScreenWithTabWithDefaults instantiates a new ScreenWithTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScreenWithTab) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenWithTab) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenWithTab) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenWithTab) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ScreenWithTab) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScreenWithTab) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScreenWithTab) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ScreenWithTab) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScreenWithTab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenWithTab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenWithTab) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScreenWithTab) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *ScreenWithTab) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ScreenWithTab) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ScreenWithTab) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ScreenWithTab) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTab

`func (o *ScreenWithTab) GetTab() ScreenableTab`

GetTab returns the Tab field if non-nil, zero value otherwise.

### GetTabOk

`func (o *ScreenWithTab) GetTabOk() (*ScreenableTab, bool)`

GetTabOk returns a tuple with the Tab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTab

`func (o *ScreenWithTab) SetTab(v ScreenableTab)`

SetTab sets Tab field to given value.

### HasTab

`func (o *ScreenWithTab) HasTab() bool`

HasTab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


