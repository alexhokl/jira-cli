# IconBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The URL of the tooltip, used only for a status icon. | [optional] 
**Title** | Pointer to **string** | The title of the icon, for use as a tooltip on the icon. | [optional] 
**Url16x16** | Pointer to **string** | The URL of a 16x16 pixel icon. | [optional] 

## Methods

### NewIconBean

`func NewIconBean() *IconBean`

NewIconBean instantiates a new IconBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconBeanWithDefaults

`func NewIconBeanWithDefaults() *IconBean`

NewIconBeanWithDefaults instantiates a new IconBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *IconBean) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IconBean) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IconBean) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *IconBean) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetTitle

`func (o *IconBean) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IconBean) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IconBean) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IconBean) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl16x16

`func (o *IconBean) GetUrl16x16() string`

GetUrl16x16 returns the Url16x16 field if non-nil, zero value otherwise.

### GetUrl16x16Ok

`func (o *IconBean) GetUrl16x16Ok() (*string, bool)`

GetUrl16x16Ok returns a tuple with the Url16x16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl16x16

`func (o *IconBean) SetUrl16x16(v string)`

SetUrl16x16 sets Url16x16 field to given value.

### HasUrl16x16

`func (o *IconBean) HasUrl16x16() bool`

HasUrl16x16 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


