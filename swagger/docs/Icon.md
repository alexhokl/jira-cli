# Icon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The URL of the tooltip, used only for a status icon. If not set, the status icon in Jira is not clickable. | [optional] 
**Title** | Pointer to **string** | The title of the icon. This is used as follows:   *  For a status icon it is used as a tooltip on the icon. If not set, the status icon doesn&#39;t display a tooltip in Jira.  *  For the remote object icon it is used in conjunction with the application name to display a tooltip for the link&#39;s icon. The tooltip takes the format \&quot;\\[application name\\] icon title\&quot;. Blank itemsare excluded from the tooltip title. If both items are blank, the icon tooltop displays as \&quot;Web Link\&quot;. | [optional] 
**Url16x16** | Pointer to **string** | The URL of an icon that displays at 16x16 pixel in Jira. | [optional] 

## Methods

### NewIcon

`func NewIcon() *Icon`

NewIcon instantiates a new Icon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconWithDefaults

`func NewIconWithDefaults() *Icon`

NewIconWithDefaults instantiates a new Icon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *Icon) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Icon) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Icon) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Icon) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetTitle

`func (o *Icon) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Icon) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Icon) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Icon) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl16x16

`func (o *Icon) GetUrl16x16() string`

GetUrl16x16 returns the Url16x16 field if non-nil, zero value otherwise.

### GetUrl16x16Ok

`func (o *Icon) GetUrl16x16Ok() (*string, bool)`

GetUrl16x16Ok returns a tuple with the Url16x16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl16x16

`func (o *Icon) SetUrl16x16(v string)`

SetUrl16x16 sets Url16x16 field to given value.

### HasUrl16x16

`func (o *Icon) HasUrl16x16() bool`

HasUrl16x16 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


