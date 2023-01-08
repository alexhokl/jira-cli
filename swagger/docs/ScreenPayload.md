# ScreenPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the screen | [optional] 
**Name** | Pointer to **string** | The name of the screen | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Tabs** | Pointer to [**[]TabPayload**](TabPayload.md) | The tabs of the screen. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/\\#api-rest-api-3-screens-screenid-tabs-tabid-fields-post | [optional] 

## Methods

### NewScreenPayload

`func NewScreenPayload() *ScreenPayload`

NewScreenPayload instantiates a new ScreenPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenPayloadWithDefaults

`func NewScreenPayloadWithDefaults() *ScreenPayload`

NewScreenPayloadWithDefaults instantiates a new ScreenPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScreenPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScreenPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScreenPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScreenPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScreenPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScreenPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScreenPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScreenPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *ScreenPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *ScreenPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *ScreenPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *ScreenPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetTabs

`func (o *ScreenPayload) GetTabs() []TabPayload`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *ScreenPayload) GetTabsOk() (*[]TabPayload, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabs

`func (o *ScreenPayload) SetTabs(v []TabPayload)`

SetTabs sets Tabs field to given value.

### HasTabs

`func (o *ScreenPayload) HasTabs() bool`

HasTabs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


