# DashboardGadget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | The color of the gadget. Should be one of &#x60;blue&#x60;, &#x60;red&#x60;, &#x60;yellow&#x60;, &#x60;green&#x60;, &#x60;cyan&#x60;, &#x60;purple&#x60;, &#x60;gray&#x60;, or &#x60;white&#x60;. | [readonly] 
**Id** | **int64** | The ID of the gadget instance. | [readonly] 
**ModuleKey** | Pointer to **string** | The module key of the gadget type. | [optional] [readonly] 
**Position** | [**DashboardGadgetPosition**](DashboardGadgetPosition.md) | The position of the gadget. | [readonly] 
**Title** | **string** | The title of the gadget. | [readonly] 
**Uri** | Pointer to **string** | The URI of the gadget type. | [optional] [readonly] 

## Methods

### NewDashboardGadget

`func NewDashboardGadget(color string, id int64, position DashboardGadgetPosition, title string, ) *DashboardGadget`

NewDashboardGadget instantiates a new DashboardGadget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardGadgetWithDefaults

`func NewDashboardGadgetWithDefaults() *DashboardGadget`

NewDashboardGadgetWithDefaults instantiates a new DashboardGadget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *DashboardGadget) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DashboardGadget) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DashboardGadget) SetColor(v string)`

SetColor sets Color field to given value.


### GetId

`func (o *DashboardGadget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardGadget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardGadget) SetId(v int64)`

SetId sets Id field to given value.


### GetModuleKey

`func (o *DashboardGadget) GetModuleKey() string`

GetModuleKey returns the ModuleKey field if non-nil, zero value otherwise.

### GetModuleKeyOk

`func (o *DashboardGadget) GetModuleKeyOk() (*string, bool)`

GetModuleKeyOk returns a tuple with the ModuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleKey

`func (o *DashboardGadget) SetModuleKey(v string)`

SetModuleKey sets ModuleKey field to given value.

### HasModuleKey

`func (o *DashboardGadget) HasModuleKey() bool`

HasModuleKey returns a boolean if a field has been set.

### GetPosition

`func (o *DashboardGadget) GetPosition() DashboardGadgetPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DashboardGadget) GetPositionOk() (*DashboardGadgetPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DashboardGadget) SetPosition(v DashboardGadgetPosition)`

SetPosition sets Position field to given value.


### GetTitle

`func (o *DashboardGadget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardGadget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardGadget) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUri

`func (o *DashboardGadget) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *DashboardGadget) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *DashboardGadget) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *DashboardGadget) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


