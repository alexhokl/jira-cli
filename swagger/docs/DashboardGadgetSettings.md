# DashboardGadgetSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | The color of the gadget. Should be one of &#x60;blue&#x60;, &#x60;red&#x60;, &#x60;yellow&#x60;, &#x60;green&#x60;, &#x60;cyan&#x60;, &#x60;purple&#x60;, &#x60;gray&#x60;, or &#x60;white&#x60;. | [optional] 
**IgnoreUriAndModuleKeyValidation** | Pointer to **bool** | Whether to ignore the validation of module key and URI. For example, when a gadget is created that is a part of an application that isn&#39;t installed. | [optional] 
**ModuleKey** | Pointer to **string** | The module key of the gadget type. Can&#39;t be provided with &#x60;uri&#x60;. | [optional] 
**Position** | Pointer to [**DashboardGadgetPosition**](DashboardGadgetPosition.md) | The position of the gadget. When the gadget is placed into the position, other gadgets in the same column are moved down to accommodate it. | [optional] 
**Title** | Pointer to **string** | The title of the gadget. | [optional] 
**Uri** | Pointer to **string** | The URI of the gadget type. Can&#39;t be provided with &#x60;moduleKey&#x60;. | [optional] 

## Methods

### NewDashboardGadgetSettings

`func NewDashboardGadgetSettings() *DashboardGadgetSettings`

NewDashboardGadgetSettings instantiates a new DashboardGadgetSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardGadgetSettingsWithDefaults

`func NewDashboardGadgetSettingsWithDefaults() *DashboardGadgetSettings`

NewDashboardGadgetSettingsWithDefaults instantiates a new DashboardGadgetSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *DashboardGadgetSettings) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DashboardGadgetSettings) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DashboardGadgetSettings) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DashboardGadgetSettings) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIgnoreUriAndModuleKeyValidation

`func (o *DashboardGadgetSettings) GetIgnoreUriAndModuleKeyValidation() bool`

GetIgnoreUriAndModuleKeyValidation returns the IgnoreUriAndModuleKeyValidation field if non-nil, zero value otherwise.

### GetIgnoreUriAndModuleKeyValidationOk

`func (o *DashboardGadgetSettings) GetIgnoreUriAndModuleKeyValidationOk() (*bool, bool)`

GetIgnoreUriAndModuleKeyValidationOk returns a tuple with the IgnoreUriAndModuleKeyValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUriAndModuleKeyValidation

`func (o *DashboardGadgetSettings) SetIgnoreUriAndModuleKeyValidation(v bool)`

SetIgnoreUriAndModuleKeyValidation sets IgnoreUriAndModuleKeyValidation field to given value.

### HasIgnoreUriAndModuleKeyValidation

`func (o *DashboardGadgetSettings) HasIgnoreUriAndModuleKeyValidation() bool`

HasIgnoreUriAndModuleKeyValidation returns a boolean if a field has been set.

### GetModuleKey

`func (o *DashboardGadgetSettings) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *DashboardGadgetSettings) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *DashboardGadgetSettings) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *DashboardGadgetSettings) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.

### GetPosition

`func (o *DashboardGadgetSettings) GetPosition() DashboardGadgetPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DashboardGadgetSettings) GetPositionOk() (*DashboardGadgetPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DashboardGadgetSettings) SetPosition(v DashboardGadgetPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DashboardGadgetSettings) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardGadgetSettings) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardGadgetSettings) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardGadgetSettings) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DashboardGadgetSettings) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUri

`func (o *DashboardGadgetSettings) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *DashboardGadgetSettings) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *DashboardGadgetSettings) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *DashboardGadgetSettings) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


