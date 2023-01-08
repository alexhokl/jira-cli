# ScreenScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen scheme. | [optional] 
**Id** | Pointer to **int64** | The ID of the screen scheme. | [optional] 
**IssueTypeScreenSchemes** | Pointer to [**PageBeanIssueTypeScreenScheme**](PageBeanIssueTypeScreenScheme.md) | Details of the issue type screen schemes associated with the screen scheme. | [optional] 
**Name** | Pointer to **string** | The name of the screen scheme. | [optional] 
**Screens** | Pointer to [**ScreenTypes**](ScreenTypes.md) | The IDs of the screens for the screen types of the screen scheme. | [optional] 

## Methods

### NewScreenScheme

`func NewScreenScheme() *ScreenScheme`

NewScreenScheme instantiates a new ScreenScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenSchemeWithDefaults

`func NewScreenSchemeWithDefaults() *ScreenScheme`

NewScreenSchemeWithDefaults instantiates a new ScreenScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScreenScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ScreenScheme) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScreenScheme) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScreenScheme) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ScreenScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueTypeScreenSchemes

`func (o *ScreenScheme) GetIssueTypeScreenSchemes() PageBeanIssueTypeScreenScheme`

GetIssueTypeScreenSchemes returns the IssueTypeScreenSchemes field if non-nil, zero value otherwise.

### GetIssueTypeScreenSchemesOk

`func (o *ScreenScheme) GetIssueTypeScreenSchemesOk() (*PageBeanIssueTypeScreenScheme, bool)`

GetIssueTypeScreenSchemesOk returns a tuple with the IssueTypeScreenSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeScreenSchemes

`func (o *ScreenScheme) SetIssueTypeScreenSchemes(v PageBeanIssueTypeScreenScheme)`

SetIssueTypeScreenSchemes sets IssueTypeScreenSchemes field to given value.

### HasIssueTypeScreenSchemes

`func (o *ScreenScheme) HasIssueTypeScreenSchemes() bool`

HasIssueTypeScreenSchemes returns a boolean if a field has been set.

### GetName

`func (o *ScreenScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScreenScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScreens

`func (o *ScreenScheme) GetScreens() ScreenTypes`

GetScreens returns the Screens field if non-nil, zero value otherwise.

### GetScreensOk

`func (o *ScreenScheme) GetScreensOk() (*ScreenTypes, bool)`

GetScreensOk returns a tuple with the Screens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreens

`func (o *ScreenScheme) SetScreens(v ScreenTypes)`

SetScreens sets Screens field to given value.

### HasScreens

`func (o *ScreenScheme) HasScreens() bool`

HasScreens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


