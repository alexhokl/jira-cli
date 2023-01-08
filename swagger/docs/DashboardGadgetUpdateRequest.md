# DashboardGadgetUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | The color of the gadget. Should be one of &#x60;blue&#x60;, &#x60;red&#x60;, &#x60;yellow&#x60;, &#x60;green&#x60;, &#x60;cyan&#x60;, &#x60;purple&#x60;, &#x60;gray&#x60;, or &#x60;white&#x60;. | [optional] 
**Position** | Pointer to [**DashboardGadgetPosition**](DashboardGadgetPosition.md) | The position of the gadget. | [optional] 
**Title** | Pointer to **string** | The title of the gadget. | [optional] 

## Methods

### NewDashboardGadgetUpdateRequest

`func NewDashboardGadgetUpdateRequest() *DashboardGadgetUpdateRequest`

NewDashboardGadgetUpdateRequest instantiates a new DashboardGadgetUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardGadgetUpdateRequestWithDefaults

`func NewDashboardGadgetUpdateRequestWithDefaults() *DashboardGadgetUpdateRequest`

NewDashboardGadgetUpdateRequestWithDefaults instantiates a new DashboardGadgetUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *DashboardGadgetUpdateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DashboardGadgetUpdateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DashboardGadgetUpdateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DashboardGadgetUpdateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPosition

`func (o *DashboardGadgetUpdateRequest) GetPosition() DashboardGadgetPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DashboardGadgetUpdateRequest) GetPositionOk() (*DashboardGadgetPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DashboardGadgetUpdateRequest) SetPosition(v DashboardGadgetPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DashboardGadgetUpdateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardGadgetUpdateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardGadgetUpdateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardGadgetUpdateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DashboardGadgetUpdateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


